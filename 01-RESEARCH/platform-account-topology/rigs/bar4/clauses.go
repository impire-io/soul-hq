package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/nats-io/jwt/v2"
	natsserver "github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nkeys"
)

// ---------------------------------------------------------------------------
// C1 — token lane: one human, two tenants, two tokens; admits both; isolated.
// ---------------------------------------------------------------------------
func clauseC1() {
	fmt.Println("\n--- C1: one human holds two tokens into tenants A and B ---")

	opKP, _ := nkeys.CreateOperator()
	opPub, _ := opKP.PublicKey()
	sysKP, sysPub := acct()
	sysJWT := mustEnc(named(jwt.NewAccountClaims(sysPub), "SYS"), opKP)

	authKP, authPub := acct()
	authSKKP, authSKPub := acct()
	authSKSeed, _ := authSKKP.Seed()
	issuerUserKP, _ := nkeys.CreateUser()
	issuerUserPub, _ := issuerUserKP.PublicKey()
	issuerUserSeed, _ := issuerUserKP.Seed()

	aKP, aPub := acct()
	aSKKP, aSKPub := acct()
	bKP, bPub := acct()
	bSKKP, bSKPub := acct()
	_ = aKP
	_ = bKP

	authClaims := jwt.NewAccountClaims(authPub)
	authClaims.Name = "AUTH"
	authClaims.SigningKeys.Add(authSKPub)
	authClaims.EnableExternalAuthorization(issuerUserPub)
	authClaims.Authorization.AllowedAccounts.Add(aPub, bPub) // both tenants (Bar 2 fix)
	authJWT := mustEnc(authClaims, opKP)

	aJWT := scopedTenantJWT(opKP, aPub, "tenant-a", aSKPub)
	bJWT := scopedTenantJWT(opKP, bPub, "tenant-b", bSKPub)

	res, err := natsserver.NewDirAccResolver(mustTemp(), 1000, time.Minute, natsserver.NoDelete)
	die("resolver", err)
	for pub, j := range map[string]string{sysPub: sysJWT, authPub: authJWT, aPub: aJWT, bPub: bJWT} {
		die("store", res.Store(pub, j))
	}
	srv := startServer(opKP, opPub, sysPub, res)
	defer srv.Shutdown()
	url := srv.ClientURL()

	sysConn := sysConnect(url, sysKP)
	defer sysConn.Close()

	// The issuer knows both tokens: the ONE human, two tenants.
	issuerConn, err := nats.Connect(url, nats.UserCredentials(credsFile(mustEnc(jwt.NewUserClaims(issuerUserPub), authKP), issuerUserSeed)))
	die("issuer conn", err)
	defer issuerConn.Close()
	startIssuer(issuerConn, string(authSKSeed), map[string]tokenBinding{
		"human-token-a": {signer: aSKKP, account: aPub, user: "daan"},
		"human-token-b": {signer: bSKKP, account: bPub, user: "daan"},
	})

	sentinelCreds := sentinel(authKP)

	admit := func(token string) (*nats.Conn, error) {
		return nats.Connect(url, nats.UserCredentials(sentinelCreds), nats.Token(token),
			nats.RetryOnFailedConnect(false), nats.MaxReconnects(0), nats.Name("probe"))
	}

	ca, ea := admit("human-token-a")
	cb, eb := admit("human-token-b")
	switch {
	case ea != nil:
		fmt.Printf("[C1] FAIL — human refused into tenant A: %v\n", ea)
	case eb != nil:
		fmt.Printf("[C1] FAIL — human refused into tenant B: %v\n", eb)
	default:
		fmt.Printf("[C1] PASS — the same human admitted into BOTH tenants with two tokens\n")
	}
	// Isolation: the A-identity cannot act in B's account. Both share the
	// subject name SOULSTREAM.>, but they are DIFFERENT accounts' streams;
	// prove A cannot reach B by having B subscribe and A publish — no delivery.
	if ea == nil && eb == nil {
		gotCross := crossReach(ca, cb, "SOULSTREAM.TOPICS.cross")
		if gotCross {
			fmt.Printf("[C1] FAIL — tenant A's identity reached tenant B's subject (isolation broken)\n")
		} else {
			fmt.Printf("[C1] PASS — tenant A's identity cannot reach tenant B's account subjects (server-isolated)\n")
		}
	}
	if ca != nil {
		ca.Close()
	}
	if cb != nil {
		cb.Close()
	}
}

// crossReach returns true if a message published by conn `from` is delivered
// to a subscriber on conn `to` (i.e. cross-account leakage).
func crossReach(from, to *nats.Conn, subj string) bool {
	sub, err := to.SubscribeSync(subj)
	if err != nil {
		// to cannot even subscribe here (scoped) — treat as no leak.
		return false
	}
	_ = to.Flush()
	_ = from.Publish(subj, []byte("leak"))
	_ = from.Flush()
	_, err = sub.NextMsg(500 * time.Millisecond)
	return err == nil
}

// ---------------------------------------------------------------------------
// C2 — OIDC ambiguity: roleFor reproduced verbatim (issuer.go:241-267),
//      fuzzed over orderings to demonstrate order-independence.
// ---------------------------------------------------------------------------

// roleForRepro is issuer.go:241-267 reproduced verbatim, with the vault lookup
// replaced by a set membership over `declaredRoles` (an account signing key
// with an account binding). authKeyName is the issuer's own key (never a role).
func roleForRepro(roles []string, declaredRoles map[string]bool, authKeyName string) (string, error) {
	if len(roles) == 0 {
		return "", fmt.Errorf("token carries no roles claim")
	}
	var declared []string
	for _, role := range roles {
		if role == authKeyName {
			continue
		}
		if !declaredRoles[role] {
			continue // undeclared values are inert
		}
		declared = append(declared, role)
	}
	switch len(declared) {
	case 0:
		return "", fmt.Errorf("no declared role among roles %v", roles)
	case 1:
		return declared[0], nil
	default:
		return "", fmt.Errorf("ambiguous roles %v", declared)
	}
}

func clauseC2() {
	fmt.Println("\n--- C2: OIDC two-tenant roles claim refuses, order-independently ---")
	roleA, roleB := "SAROLEAAAA", "SAROLEBBBB"
	declared := map[string]bool{roleA: true, roleB: true}
	authKey := "auth/issuer"

	// A subject assigned into BOTH tenants: both orders + with noise values.
	cases := [][]string{
		{roleA, roleB},
		{roleB, roleA},
		{"undeclared-x", roleA, roleB},
		{roleB, "undeclared-y", roleA},
		{roleA, authKey, roleB},
	}
	allAmbiguous := true
	for _, c := range cases {
		_, err := roleForRepro(c, declared, authKey)
		amb := err != nil && strings.Contains(err.Error(), "ambiguous")
		if !amb {
			allAmbiguous = false
		}
		fmt.Printf("    roles=%v → %v\n", c, errStr(err))
	}
	// Single-tenant and none, for contrast.
	oneOK := func() bool { r, e := roleForRepro([]string{"undeclared", roleA}, declared, authKey); return e == nil && r == roleA }()
	noneRefused := func() bool { _, e := roleForRepro([]string{"undeclared"}, declared, authKey); return e != nil }()

	if allAmbiguous && oneOK && noneRefused {
		fmt.Printf("[C2] PASS — two declared roles ALWAYS refuse as ambiguous (every order); exactly one admits; none refuses. Order cannot decide (a length count).\n")
	} else {
		fmt.Printf("[C2] FAIL — allAmbiguous=%v oneOK=%v noneRefused=%v\n", allAmbiguous, oneOK, noneRefused)
	}
}

// ---------------------------------------------------------------------------
// C3 — persona-name scope: vault keys persona keys as "persona/<user>" with NO
//      account component, so a shared vault makes names GLOBAL across tenants.
// ---------------------------------------------------------------------------

// personaEntry mirrors vault.Entry's owner binding.
type personaEntry struct{ account, user string }

// sharedVaultGenerate reproduces vault.GeneratePersonaKey's collision rule
// against a SHARED map keyed exactly as the code keys it: "persona/<user>",
// no account. Returns (owner, error): a same-name second owner refuses.
func sharedVaultGenerate(store map[string]personaEntry, account, user string) (personaEntry, error) {
	key := "persona/" + user // service.go:524 — PersonaKeyPrefix + user, no account
	if e, ok := store[key]; ok {
		if e.account != account || e.user != user {
			return personaEntry{}, fmt.Errorf("vault: %s exists with another owner", key)
		}
		return e, nil
	}
	e := personaEntry{account: account, user: user}
	store[key] = e
	return e, nil
}

func clauseC3() {
	fmt.Println("\n--- C3: persona-name scope in a SHARED platform vault ---")
	accountA, accountB := "AACCOUNTA", "ABCCOUNTB"
	shared := map[string]personaEntry{}

	// Tenant A's "daan" signs first — materializes persona/daan bound to A.
	_, e1 := sharedVaultGenerate(shared, accountA, "daan")
	// Tenant B's "daan" tries to sign — SAME vault key persona/daan.
	_, e2 := sharedVaultGenerate(shared, accountB, "daan")

	if e1 == nil && e2 != nil {
		fmt.Printf("[C3] MEASURED — shared vault: tenant B's same-named user is REFUSED (\"%v\"). Persona names are GLOBAL across tenants (first-owner-wins, D26).\n", e2)
		fmt.Printf("    This is silent to the second tenant (service.go maps it to \"has no persona key\", anti-probing D26) — NOT loud, NOT naming the first owner.\n")
		fmt.Printf("    Decision for the operator — resolve one of:\n")
		fmt.Printf("      (a) per-tenant vault buckets → persona names account-scoped, no collision; or\n")
		fmt.Printf("      (b) key persona keys as persona/<account>/<user> → one shared vault, account-scoped names (a D26 clean-break rename).\n")
	} else {
		fmt.Printf("[C3] unexpected: e1=%v e2=%v\n", e1, e2)
	}

	// Control: with account-scoped keys, both sign.
	scoped := map[string]personaEntry{}
	scopedGen := func(account, user string) error {
		key := "persona/" + account + "/" + user
		if _, ok := scoped[key]; ok {
			return fmt.Errorf("exists")
		}
		scoped[key] = personaEntry{account, user}
		return nil
	}
	c1 := scopedGen(accountA, "daan")
	c2 := scopedGen(accountB, "daan")
	if c1 == nil && c2 == nil {
		fmt.Printf("[C3] CONTROL — with persona/<account>/<user> keys, BOTH tenants' \"daan\" sign; no collision (option b works).\n")
	}
}

// --- small helpers ---

func named(c *jwt.AccountClaims, n string) *jwt.AccountClaims { c.Name = n; return c }

func mustEnc(claims interface{ Encode(nkeys.KeyPair) (string, error) }, kp nkeys.KeyPair) string {
	tok, err := claims.Encode(kp)
	die("encode", err)
	return tok
}

func startServer(opKP nkeys.KeyPair, opPub, sysPub string, res natsserver.AccountResolver) *natsserver.Server {
	oc := jwt.NewOperatorClaims(opPub)
	oc.Name = "op"
	oc.SystemAccount = sysPub
	opJWT := mustEnc(oc, opKP)
	trusted, err := jwt.DecodeOperatorClaims(opJWT)
	die("decode op", err)
	srv, err := natsserver.NewServer(&natsserver.Options{
		Host: "127.0.0.1", Port: -1,
		TrustedOperators: []*jwt.OperatorClaims{trusted},
		SystemAccount:    sysPub, AccountResolver: res,
		NoLog: true, NoSigs: true,
	})
	die("server", err)
	go srv.Start()
	if !srv.ReadyForConnections(10 * time.Second) {
		die("ready", fmt.Errorf("not ready"))
	}
	return srv
}

func sysConnect(url string, sysKP nkeys.KeyPair) *nats.Conn {
	uk, _ := nkeys.CreateUser()
	upub, _ := uk.PublicKey()
	useed, _ := uk.Seed()
	nc, err := nats.Connect(url, nats.UserCredentials(credsFile(mustEnc(jwt.NewUserClaims(upub), sysKP), useed)))
	die("sys conn", err)
	return nc
}

func sentinel(authKP nkeys.KeyPair) string {
	kp, _ := nkeys.CreateUser()
	pub, _ := kp.PublicKey()
	seed, _ := kp.Seed()
	sc := jwt.NewUserClaims(pub)
	sc.Name = "sentinel"
	sc.BearerToken = true
	sc.Permissions.Pub.Deny = jwt.StringList{">"}
	sc.Permissions.Sub.Deny = jwt.StringList{">"}
	return credsFile(mustEnc(sc, authKP), seed)
}

func errStr(err error) string {
	if err == nil {
		return "ADMITTED"
	}
	return "refused: " + err.Error()
}
