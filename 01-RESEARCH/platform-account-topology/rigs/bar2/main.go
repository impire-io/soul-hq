// Bar 2 rig — one-act tenant birth and admission, the coupling questions.
//
// Reproduces the EXACT JWT shapes the identity code produces, so the server
// behaviour measured here is the behaviour the product would get:
//   - tenant account JWT: localoperator.go buildJWT — SigningKeys.Add (PLAIN,
//     not AddScopedSigner), JS unlimited, Conn=-1.
//   - minted user: mint.go claims()+ephemeral — SetScoped(true),
//     IssuerAccount=tenant, signed by the tenant's signing key.
//
// Two server-side unknowns the recon flagged, measured in isolation:
//   Q1 (scoped-on-plain-key): does the server admit a SetScoped(true) user
//       signed by a PLAIN signing key, and what permissions does it get?
//   Q2 (allowed_accounts): can a callout-issued user land in a freshly-created
//       tenant that AUTH.allowed_accounts does not list — and does amending
//       allowed_accounts fix it, at what cost?
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/nats-io/jwt/v2"
	natsserver "github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nkeys"
)

func die(ctx string, err error) {
	if err != nil {
		fmt.Printf("FATAL %s: %v\n", ctx, err)
		os.Exit(1)
	}
}

func acct() (nkeys.KeyPair, string) {
	kp, _ := nkeys.CreateAccount()
	pub, _ := kp.PublicKey()
	return kp, pub
}

// tenantAccountJWT mirrors localoperator.go buildJWT. scoped=false is the
// AS-BUILT shape (plain SigningKeys.Add); scoped=true is the control that
// adds the persona scope template as a scoped signer.
func tenantAccountJWT(op nkeys.KeyPair, acctPub, name, signingPub string, scoped bool) string {
	ac := jwt.NewAccountClaims(acctPub)
	ac.Name = name
	if scoped {
		scope := jwt.NewUserScope()
		scope.Key = signingPub
		scope.Role = "soulstream-user"
		scope.Template = jwt.UserPermissionLimits{Permissions: jwt.Permissions{
			Pub: jwt.Permission{Allow: jwt.StringList{
				"identity.status", "identity.xkey",
				"identity.{{account-subject()}}.{{name()}}.sign.record",
				"SOULSTREAM.>", "$JS.API.>", "$KV.>", "$O.>", "$SYS.REQ.USER.INFO",
			}},
			Sub: jwt.Permission{Allow: jwt.StringList{"_INBOX.>", "SOULSTREAM.>"}},
		}}
		ac.SigningKeys.AddScopedSigner(scope)
	} else {
		ac.SigningKeys.Add(signingPub)
	}
	ac.Limits.JetStreamLimits = jwt.JetStreamLimits{MemoryStorage: -1, DiskStorage: -1, Streams: -1, Consumer: -1}
	ac.Limits.Conn = -1
	tok, err := ac.Encode(op)
	die("tenant account jwt "+name, err)
	return tok
}

// mintedUserJWT mirrors mint.go claims()+ephemeral: SetScoped(true),
// IssuerAccount=account, signed by the account's signing key.
func mintedUserJWT(signingKP nkeys.KeyPair, account, user, userPub string) string {
	uc := jwt.NewUserClaims(userPub)
	uc.Name = user
	uc.IssuerAccount = account
	uc.SetScoped(true)
	uc.Expires = time.Now().Add(15 * time.Minute).Unix()
	tok, err := uc.Encode(signingKP)
	die("minted user jwt", err)
	return tok
}

func credsFile(jwtStr string, seed []byte) string {
	c, err := jwt.FormatUserConfig(jwtStr, seed)
	die("format creds", err)
	f, err := os.CreateTemp("", "rig-*.creds")
	die("temp creds", err)
	_, _ = f.Write(c)
	_ = f.Close()
	return f.Name()
}

func landAccount(sys *nats.Conn, token string) error {
	msg, err := sys.Request("$SYS.REQ.CLAIMS.UPDATE", []byte(token), 5*time.Second)
	if err != nil {
		return err
	}
	var resp struct {
		Error *struct {
			Description string `json:"description"`
		} `json:"error"`
	}
	if err := json.Unmarshal(msg.Data, &resp); err == nil && resp.Error != nil {
		return fmt.Errorf("resolver refused: %s", resp.Error.Description)
	}
	return nil
}

func main() {
	fmt.Println("=== Bar 2: one-act tenant birth and admission (the coupling questions) ===")

	opKP, _ := nkeys.CreateOperator()
	opPub, _ := opKP.PublicKey()

	// SYS.
	sysKP, sysPub := acct()
	sysClaims := jwt.NewAccountClaims(sysPub)
	sysClaims.Name = "SYS"
	sysJWT, err := sysClaims.Encode(opKP)
	die("sys jwt", err)

	// AUTH: callout enabled. allowed_accounts starts WITHOUT the tenant
	// (the tenant is born later — the coupling gap under test). We add a
	// pre-existing SVC account so allowed_accounts is non-empty and realistic.
	_, svcPub := acct()
	authKP, authPub := acct()
	authSKKP, authSKPub := acct()
	authSKSeedB, _ := authSKKP.Seed()
	authSKSeed := string(authSKSeedB)
	issuerUserKP, _ := nkeys.CreateUser()
	issuerUserPub, _ := issuerUserKP.PublicKey()
	issuerUserSeed, _ := issuerUserKP.Seed()

	buildAuthJWT := func(allowed ...string) string {
		ac := jwt.NewAccountClaims(authPub)
		ac.Name = "AUTH"
		ac.SigningKeys.Add(authSKPub)
		ac.EnableExternalAuthorization(issuerUserPub)
		for _, a := range allowed {
			ac.Authorization.AllowedAccounts.Add(a)
		}
		// No XKey: unsealed callout for rig simplicity (sealing is orthogonal
		// to the coupling questions; Bar 1 already exercises sealed transport).
		tok, err := ac.Encode(opKP)
		die("auth jwt", err)
		return tok
	}
	authJWT := buildAuthJWT(svcPub) // tenant NOT listed yet

	// SVC account (pre-existing) — where the service/vault would live.
	svcClaims := jwt.NewAccountClaims(svcPub)
	svcClaims.Name = "SVC"
	svcClaims.Limits.JetStreamLimits = jwt.JetStreamLimits{MemoryStorage: -1, DiskStorage: -1, Streams: -1, Consumer: -1}
	svcJWT, err := svcClaims.Encode(opKP)
	die("svc jwt", err)

	// Resolver + server.
	res, err := natsserver.NewDirAccResolver(mustTemp(), 1000, time.Minute, natsserver.NoDelete)
	die("resolver", err)
	for pub, j := range map[string]string{sysPub: sysJWT, authPub: authJWT, svcPub: svcJWT} {
		die("store "+pub, res.Store(pub, j))
	}
	oc := jwt.NewOperatorClaims(opPub)
	oc.Name = "op"
	oc.SystemAccount = sysPub
	opJWT, err := oc.Encode(opKP)
	die("op jwt", err)
	trusted, err := jwt.DecodeOperatorClaims(opJWT)
	die("decode op", err)
	srv, err := natsserver.NewServer(&natsserver.Options{
		Host: "127.0.0.1", Port: -1,
		TrustedOperators: []*jwt.OperatorClaims{trusted},
		SystemAccount:    sysPub,
		AccountResolver:  res,
		NoLog:            true, NoSigs: true,
	})
	die("server", err)
	go srv.Start()
	if !srv.ReadyForConnections(10 * time.Second) {
		die("ready", fmt.Errorf("not ready"))
	}
	defer srv.Shutdown()
	url := srv.ClientURL()

	// SYS connection (the one act's channel).
	sysUserKP, _ := nkeys.CreateUser()
	sysUserPub, _ := sysUserKP.PublicKey()
	sysUserSeed, _ := sysUserKP.Seed()
	sysUserJWT, err := jwt.NewUserClaims(sysUserPub).Encode(sysKP)
	die("sys user jwt", err)
	sysConn, err := nats.Connect(url, nats.UserCredentials(credsFile(sysUserJWT, sysUserSeed)))
	die("sys conn", err)
	defer sysConn.Close()

	// --- BIRTH: create tenant "acme" as one act (mirror localoperator). ---
	acmeKP, acmePub := acct()          // the tenant account identity
	acmeSKKP, acmeSKPub := acct()      // its signing key (plain)
	acmeSKSeed, _ := acmeSKKP.Seed()
	_ = acmeSKSeed
	born := time.Now()
	die("land acme", landAccount(sysConn, tenantAccountJWT(opKP, acmePub, "acme", acmeSKPub, false)))
	landDur := time.Since(born)
	fmt.Printf("\n[birth] tenant acme landed as one act in %v\n", landDur)
	_ = acmeKP

	// ================= Q1: scoped-on-plain-key, in isolation =================
	// Direct connect (no callout, no allowed_accounts gate) with a user minted
	// exactly as mint.ephemeral does, signed by acme's PLAIN signing key.
	fmt.Println("\n--- Q1: does the server admit a SetScoped(true) user on a PLAIN signing key? ---")
	q1UserKP, _ := nkeys.CreateUser()
	q1UserPub, _ := q1UserKP.PublicKey()
	q1UserSeed, _ := q1UserKP.Seed()
	q1JWT := mintedUserJWT(acmeSKKP, acmePub, "alice", q1UserPub)
	q1Creds := credsFile(q1JWT, q1UserSeed)
	q1Conn, q1Err := nats.Connect(url, nats.UserCredentials(q1Creds),
		nats.RetryOnFailedConnect(false), nats.MaxReconnects(0))
	if q1Err != nil {
		fmt.Printf("[Q1] admission REFUSED on plain signing key: %v\n", q1Err)
	} else {
		fmt.Printf("[Q1] admission SUCCEEDED on plain signing key\n")
		q1Conn.Close()
		probePerms(url, "Q1/plain", q1Creds)
	}

	// CONTROL: the same, but acme's signing key is a SCOPED signer.
	fmt.Println("\n--- Q1 control: the same user on a SCOPED signing key (persona template) ---")
	acme2KP, acme2Pub := acct()
	acme2SKKP, acme2SKPub := acct()
	die("land acme2", landAccount(sysConn, tenantAccountJWT(opKP, acme2Pub, "acme2", acme2SKPub, true)))
	_ = acme2KP
	c2UserKP, _ := nkeys.CreateUser()
	c2UserPub, _ := c2UserKP.PublicKey()
	c2UserSeed, _ := c2UserKP.Seed()
	c2JWT := mintedUserJWT(acme2SKKP, acme2Pub, "alice", c2UserPub)
	c2Creds := credsFile(c2JWT, c2UserSeed)
	c2Conn, c2Err := nats.Connect(url, nats.UserCredentials(c2Creds),
		nats.RetryOnFailedConnect(false), nats.MaxReconnects(0))
	if c2Err != nil {
		fmt.Printf("[Q1c] admission REFUSED on scoped signing key: %v\n", c2Err)
	} else {
		fmt.Printf("[Q1c] admission SUCCEEDED on scoped signing key\n")
		c2Conn.Close()
		probePerms(url, "Q1c/scoped", c2Creds)
	}

	// ================= Q2: allowed_accounts coupling via callout =================
	fmt.Println("\n--- Q2: callout admission into acme (tenant NOT in allowed_accounts) ---")
	// Sentinel: a bearer, deny-all user in AUTH (D19). Hand-built.
	sentinelKP, _ := nkeys.CreateUser()
	sentinelPub, _ := sentinelKP.PublicKey()
	sentinelSeed, _ := sentinelKP.Seed()
	sc := jwt.NewUserClaims(sentinelPub)
	sc.Name = "sentinel"
	sc.BearerToken = true
	sc.Permissions.Pub.Deny = jwt.StringList{">"}
	sc.Permissions.Sub.Deny = jwt.StringList{">"}
	sentinelJWT, err := sc.Encode(authKP)
	die("sentinel jwt", err)
	sentinelCreds := credsFile(sentinelJWT, sentinelSeed)

	// The callout issuer: a tiny stand-in that mints for the server-assigned
	// key exactly as mint.ForKey→ephemeral does, using acme's signing key.
	issuerCreds := credsFile(mustIssuerJWT(authKP, issuerUserPub), issuerUserSeed)
	issuerConn, err := nats.Connect(url, nats.UserCredentials(issuerCreds))
	die("issuer conn", err)
	defer issuerConn.Close()
	// The issuer knows both tenants' signing keys: acme (plain), acme2 (scoped).
	startIssuer(issuerConn, authSKSeed, map[string]tokenBinding{
		"acme-token":  {signer: acmeSKKP, account: acmePub, user: "alice"},
		"acme2-token": {signer: acme2SKKP, account: acme2Pub, user: "alice"},
	})

	admitTest := func(token string) (bool, time.Duration, error, string) {
		t0 := time.Now()
		c, e := nats.Connect(url, nats.UserCredentials(sentinelCreds), nats.Token(token),
			nats.RetryOnFailedConnect(false), nats.MaxReconnects(0))
		d := time.Since(t0)
		if e != nil {
			return false, d, e, ""
		}
		c.Close()
		return true, d, nil, ""
	}

	ok, d, e, _ := admitTest("acme-token")
	if ok {
		fmt.Printf("[Q2a] UNEXPECTED — admitted into acme without allowed_accounts (%v)\n", d)
	} else {
		fmt.Printf("[Q2a] REFUSED as predicted — tenant not in allowed_accounts: %v\n", e)
	}

	// Amend AUTH: add BOTH tenants to allowed_accounts, re-land as one act.
	fmt.Println("\n--- Q2: amend AUTH.allowed_accounts += acme, acme2; re-land; retry ---")
	amend0 := time.Now()
	die("re-land auth", landAccount(sysConn, buildAuthJWT(svcPub, acmePub, acme2Pub)))
	amendDur := time.Since(amend0)
	time.Sleep(200 * time.Millisecond) // let the resolver update propagate

	// Q2b — acme (PLAIN signing key): connection admitted, but the user is
	// inert (Q1's defect compounding through the callout path).
	ok2, d2, _, _ := admitTest("acme-token")
	if ok2 {
		fmt.Printf("[Q2b] acme (plain key): transport ADMITTED after amend (%v) — but the minted user is:\n", d2)
		// Build a direct-connect creds mirroring what the callout minted, to probe.
		pk, _ := nkeys.CreateUser()
		pkPub, _ := pk.PublicKey()
		pkSeed, _ := pk.Seed()
		probePerms(url, "Q2b/acme-plain", credsFile(mintedUserJWT(acmeSKKP, acmePub, "alice", pkPub), pkSeed))
	} else {
		fmt.Printf("[Q2b] acme still refused after amend\n")
	}

	// Q2c — acme2 (SCOPED signing key) + in allowed_accounts: the combined fix.
	fmt.Println("\n--- Q2c: acme2 (SCOPED signing key) + in allowed_accounts — the combined fix ---")
	ok3, d3, _, _ := admitTest("acme2-token")
	if ok3 {
		fmt.Printf("[Q2c] acme2 ADMITTED after amend (%v) — and the minted user is:\n", d3)
		sk, _ := nkeys.CreateUser()
		skPub, _ := sk.PublicKey()
		skSeed, _ := sk.Seed()
		probePerms(url, "Q2c/acme2-scoped", credsFile(mintedUserJWT(acme2SKKP, acme2Pub, "alice", skPub), skSeed))
	} else {
		fmt.Printf("[Q2c] acme2 refused after amend\n")
	}

	fmt.Printf("\n[timing] tenant birth (one act): %v | AUTH amend land: %v\n", landDur, amendDur)
}

// probePerms reports what a user can actually do, using a FRESH connection per
// subject (a permissions violation degrades a connection, so reuse would lie).
// Each connection installs an async error handler capturing the violation.
func probePerms(url, label, creds string) {
	// First: is the connection even usable? A SetScoped(true) user with no
	// matching scoped-signer template inherits the scoped sentinel limits
	// (0 subscriptions / 0 payload) — admitted but INERT. Detect that once.
	probe := make(chan error, 1)
	c, err := nats.Connect(url, nats.UserCredentials(creds),
		nats.RetryOnFailedConnect(false), nats.MaxReconnects(0),
		nats.ErrorHandler(func(_ *nats.Conn, _ *nats.Subscription, e error) {
			select {
			case probe <- e:
			default:
			}
		}))
	if err != nil {
		fmt.Printf("    [%s] connection UNUSABLE — %v (scoped-empty limits: admitted but inert)\n", label, err)
		return
	}
	// Try one real subscription; 0-sub scoped users fail here.
	if _, serr := c.SubscribeSync("_INBOX.probe"); serr != nil {
		c.Close()
		fmt.Printf("    [%s] INERT — cannot subscribe: %v (scoped-empty: 0 subs / 0 payload, admitted but useless)\n", label, serr)
		return
	}
	c.Close()

	testPub := func(subj string) string {
		violation := make(chan error, 1)
		cc, err := nats.Connect(url, nats.UserCredentials(creds),
			nats.RetryOnFailedConnect(false), nats.MaxReconnects(0),
			nats.ErrorHandler(func(_ *nats.Conn, _ *nats.Subscription, e error) {
				select {
				case violation <- e:
				default:
				}
			}))
		if err != nil {
			return "connect err: " + err.Error()
		}
		defer cc.Close()
		if err := cc.Publish(subj, []byte("x")); err != nil {
			return "pub err: " + err.Error()
		}
		_ = cc.FlushTimeout(time.Second)
		select {
		case e := <-violation:
			return "DENIED (" + e.Error() + ")"
		case <-time.After(300 * time.Millisecond):
			return "ALLOWED"
		}
	}
	fmt.Printf("    [%s] pub SOULSTREAM.TOPICS.x → %s\n", label, testPub("SOULSTREAM.TOPICS.x"))
	fmt.Printf("    [%s] pub random.foreign.x    → %s\n", label, testPub("random.foreign.x"))
	fmt.Printf("    [%s] pub $JS.API.STREAM.NAMES → %s\n", label, testPub("$JS.API.STREAM.NAMES"))
}

func mustIssuerJWT(authKP nkeys.KeyPair, issuerUserPub string) string {
	tok, err := jwt.NewUserClaims(issuerUserPub).Encode(authKP)
	die("issuer user jwt", err)
	return tok
}

type tokenBinding struct {
	signer  nkeys.KeyPair
	account string
	user    string
}

// startIssuer subscribes the callout responder: for a known token it mints
// (account,user) into the target account using its signing key — exactly the
// shape mint.ForKey→ephemeral produces (SetScoped(true), IssuerAccount set).
func startIssuer(nc *nats.Conn, authSKSeed string, bindings map[string]tokenBinding) {
	authSK, err := nkeys.FromSeed([]byte(authSKSeed))
	die("auth sk", err)
	_, err = nc.Subscribe("$SYS.REQ.USER.AUTH", func(m *nats.Msg) {
		req, err := jwt.DecodeAuthorizationRequestClaims(string(m.Data))
		if err != nil {
			return
		}
		resp := jwt.NewAuthorizationResponseClaims(req.UserNkey)
		resp.Audience = req.Server.ID
		if b, ok := bindings[req.ConnectOptions.Token]; ok {
			uc := jwt.NewUserClaims(req.UserNkey)
			uc.Name = b.user
			uc.IssuerAccount = b.account
			uc.SetScoped(true)
			uc.Expires = time.Now().Add(15 * time.Minute).Unix()
			if userJWT, err := uc.Encode(b.signer); err == nil {
				resp.Jwt = userJWT
			} else {
				resp.Error = err.Error()
			}
		} else {
			resp.Error = "credential rejected"
		}
		tok, err := resp.Encode(authSK)
		if err != nil {
			return
		}
		_ = m.Respond([]byte(tok))
	})
	die("issuer subscribe", err)
	die("issuer flush", nc.Flush())
}

func mustTemp() string {
	d, err := os.MkdirTemp("", "rig-resolver-*")
	die("temp dir", err)
	return d
}
