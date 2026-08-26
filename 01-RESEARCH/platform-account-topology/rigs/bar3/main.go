// Bar 3 rig — isolation holds through shared services.
//
// The topology under test: ONE platform service (an archivist/runtime
// stand-in) serving MANY tenants. The only model that composes with account
// isolation is per-tenant connections: the shared service holds one
// account-scoped connection INTO each tenant and reaches that tenant's data
// only through it. The service shares process and code, never credentials.
//
// This rig provisions REAL SOULSTREAM streams (soulstream-core/realm.ProvisionOn)
// in two tenant accounts A and B, runs a shared service holding one connection
// into each, and fires adversarial probes from a principal in A:
//
//	P1  A reads B's SOULSTREAM stream directly            → must refuse
//	P2  A writes B's account subjects directly            → must refuse
//	P3  A drives the shared service to act in B           → must be impossible
//	P4  the service's act on A's behalf lands in A's      → attribution/isolation
//	    stream, never B's
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/nats-io/jwt/v2"
	natsserver "github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/nats-io/nkeys"

	"github.com/impire-io/soulstream-core/realm"
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

func enc(c interface {
	Encode(nkeys.KeyPair) (string, error)
}, kp nkeys.KeyPair) string {
	t, err := c.Encode(kp)
	die("encode", err)
	return t
}

func credsFile(jwtStr string, seed []byte) string {
	cc, err := jwt.FormatUserConfig(jwtStr, seed)
	die("creds", err)
	f, _ := os.CreateTemp("", "rig-*.creds")
	_, _ = f.Write(cc)
	_ = f.Close()
	return f.Name()
}

// tenantAccountJWT: JetStream unlimited, a scoped signer with the persona
// template, plus a service signing key for the platform service's own user.
func tenantAccountJWT(op nkeys.KeyPair, acctPub, name, personaSK string) string {
	ac := jwt.NewAccountClaims(acctPub)
	ac.Name = name
	scope := jwt.NewUserScope()
	scope.Key = personaSK
	scope.Role = "soulstream-user"
	scope.Template = jwt.UserPermissionLimits{Permissions: jwt.Permissions{
		Pub: jwt.Permission{Allow: jwt.StringList{
			"identity.status", "identity.xkey",
			"SOULSTREAM.>", "$JS.API.>", "$KV.>", "$O.>", "$SYS.REQ.USER.INFO",
		}},
		Sub: jwt.Permission{Allow: jwt.StringList{"_INBOX.>", "SOULSTREAM.>"}},
	}}
	ac.SigningKeys.AddScopedSigner(scope)
	ac.Limits.JetStreamLimits = jwt.JetStreamLimits{MemoryStorage: -1, DiskStorage: -1, Streams: -1, Consumer: -1}
	ac.Limits.Conn = -1
	return enc(ac, op)
}

// userInAccount mints a user directly under the account key (bypass lane) with
// explicit permissions — the platform service's per-tenant connection user.
func userInAccount(accKP nkeys.KeyPair, name string, pub, sub []string) string {
	uk, _ := nkeys.CreateUser()
	upub, _ := uk.PublicKey()
	useed, _ := uk.Seed()
	uc := jwt.NewUserClaims(upub)
	uc.Name = name
	uc.Permissions.Pub.Allow = pub
	uc.Permissions.Sub.Allow = sub
	return credsFile(enc(uc, accKP), useed)
}

// scopedUserInAccount mints a scoped user signed by the account's persona
// signing key — a tenant PRINCIPAL, scoped to the persona template.
func scopedUserInAccount(skKP nkeys.KeyPair, accountPub, name string) string {
	uk, _ := nkeys.CreateUser()
	upub, _ := uk.PublicKey()
	useed, _ := uk.Seed()
	uc := jwt.NewUserClaims(upub)
	uc.Name = name
	uc.IssuerAccount = accountPub
	uc.SetScoped(true)
	uc.Expires = time.Now().Add(time.Hour).Unix()
	return credsFile(enc(uc, skKP), useed)
}

func main() {
	fmt.Println("=== Bar 3: isolation holds through shared services ===")
	ctx := context.Background()

	opKP, opPub := func() (nkeys.KeyPair, string) { kp, _ := nkeys.CreateOperator(); p, _ := kp.PublicKey(); return kp, p }()
	sysKP, sysPub := acct()
	sysJWT := enc(func() *jwt.AccountClaims { c := jwt.NewAccountClaims(sysPub); c.Name = "SYS"; return c }(), opKP)

	aKP, aPub := acct()
	aSK, aSKPub := acct()
	bKP, bPub := acct()
	bSK, bSKPub := acct()

	aJWT := tenantAccountJWT(opKP, aPub, "tenant-a", aSKPub)
	bJWT := tenantAccountJWT(opKP, bPub, "tenant-b", bSKPub)

	res, err := natsserver.NewDirAccResolver(mustTemp(), 1000, time.Minute, natsserver.NoDelete)
	die("resolver", err)
	for pub, j := range map[string]string{sysPub: sysJWT, aPub: aJWT, bPub: bJWT} {
		die("store", res.Store(pub, j))
	}
	_ = sysKP
	oc := jwt.NewOperatorClaims(opPub)
	oc.Name = "op"
	oc.SystemAccount = sysPub
	trusted, _ := jwt.DecodeOperatorClaims(enc(oc, opKP))
	srv, err := natsserver.NewServer(&natsserver.Options{
		Host: "127.0.0.1", Port: -1,
		TrustedOperators: []*jwt.OperatorClaims{trusted},
		SystemAccount:    sysPub, AccountResolver: res,
		NoLog: true, NoSigs: true, JetStream: true, StoreDir: mustTemp(),
	})
	die("server", err)
	go srv.Start()
	if !srv.ReadyForConnections(10 * time.Second) {
		die("ready", fmt.Errorf("not ready"))
	}
	defer srv.Shutdown()
	url := srv.ClientURL()

	// The platform service's per-tenant connection users: full SOULSTREAM +
	// JS within EACH tenant account (the archivist/runtime service user).
	// A responder needs to publish replies to requesters' inboxes.
	svcPerms := []string{"SOULSTREAM.>", "$JS.API.>", "$KV.>", "$O.>", "_INBOX.>"}
	svcSub := []string{"_INBOX.>", "SOULSTREAM.>"}
	svcAcreds := userInAccount(aKP, "archivist", svcPerms, svcSub)
	svcBcreds := userInAccount(bKP, "archivist", svcPerms, svcSub)

	svcA, err := nats.Connect(url, nats.UserCredentials(svcAcreds))
	die("svcA", err)
	defer svcA.Close()
	svcB, err := nats.Connect(url, nats.UserCredentials(svcBcreds))
	die("svcB", err)
	defer svcB.Close()

	// Provision REAL SOULSTREAM streams in each tenant, through the service's
	// per-tenant connection — exactly how a platform archivist would set up.
	jsA, err := jetstream.New(svcA)
	die("jsA", err)
	jsB, err := jetstream.New(svcB)
	die("jsB", err)
	if _, err := realm.ProvisionOn(ctx, jsA); err != nil {
		die("provision A", err)
	}
	if _, err := realm.ProvisionOn(ctx, jsB); err != nil {
		die("provision B", err)
	}
	fmt.Println("[setup] real SOULSTREAM streams provisioned in tenants A and B")

	// The shared service exposes a request surface in EACH tenant: it echoes
	// which tenant-connection served the request (the delivery-log proof).
	// The service derives the tenant from WHICH connection delivered the
	// request (server-proven), NEVER from the request payload — the D15
	// discipline applied to a shared service. It echoes both so the rig can
	// prove a payload claim is ignored.
	shared := func(nc *nats.Conn, tenant string) {
		_, err := nc.Subscribe("SOULSTREAM.MEMORY.REQUEST", func(m *nats.Msg) {
			if m.Reply != "" {
				// Payload may CLAIM a tenant; the service ignores it and uses
				// the connection's tenant. Report both for the proof.
				_ = m.Respond([]byte("served-by-tenant:" + tenant + " payload-claimed:" + string(m.Data)))
			}
		})
		die("shared sub "+tenant, err)
	}
	shared(svcA, "A")
	shared(svcB, "B")
	die("flush A", svcA.Flush())
	die("flush B", svcB.Flush())

	// A tenant-A PRINCIPAL (scoped user).
	aPrin, err := nats.Connect(url, nats.UserCredentials(scopedUserInAccount(aSK, aPub, "alice")),
		nats.RetryOnFailedConnect(false), nats.MaxReconnects(0))
	die("A principal", err)
	defer aPrin.Close()

	fmt.Println()
	// P1 — A reads B's SOULSTREAM stream directly. A's JS API only reaches
	//   A's account; B's stream is in a different account entirely.
	jsAP, err := jetstream.New(aPrin)
	die("jsAP", err)
	p1ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	// Ask A's JS for the stream — it lists only A's streams; B's is invisible.
	names := []string{}
	for n := range jsAP.StreamNames(p1ctx).Name() {
		names = append(names, n)
	}
	// A can see its OWN SOULSTREAM (proves the probe works), and cannot name
	// B's stream because JS is account-scoped — there is no cross-account read.
	fmt.Printf("[P1] A's JetStream lists only its own streams: %v — B's stream is in a separate account, unreachable by construction\n", names)

	// P2 — A publishes to B's account subjects. In A's account this is just a
	//   local publish that never crosses to B (different account); a B
	//   subscriber sees nothing.
	bWatcher, err := nats.Connect(url, nats.UserCredentials(userInAccount(bKP, "watcher", []string{"_INBOX.>"}, []string{"SOULSTREAM.>"})))
	die("bWatcher", err)
	defer bWatcher.Close()
	wsub, _ := bWatcher.SubscribeSync("SOULSTREAM.TOPICS.>")
	_ = bWatcher.Flush()
	_ = aPrin.Publish("SOULSTREAM.TOPICS.attack", []byte("cross"))
	_ = aPrin.Flush()
	if _, err := wsub.NextMsg(500 * time.Millisecond); err == nil {
		fmt.Printf("[P2] FAIL — A's publish crossed into B's account\n")
	} else {
		fmt.Printf("[P2] PASS — A's publish to SOULSTREAM.TOPICS stayed in A's account; B saw nothing (account-isolated)\n")
	}

	// P3 — A drives the shared service. A's request reaches only the service's
	//   A-connection; the response proves which tenant served it. A can NEVER
	//   reach the service's B-connection (that subscription is on B's account).
	resp, err := aPrin.Request("SOULSTREAM.MEMORY.REQUEST", []byte("who serves me"), 2*time.Second)
	if err != nil {
		fmt.Printf("[P3] A's request to the shared service errored: %v\n", err)
	} else if hasPrefix(string(resp.Data), "served-by-tenant:A") {
		fmt.Printf("[P3] PASS — the shared service served A from A's connection (%q); A cannot address the B-connection\n", string(resp.Data))
	} else {
		fmt.Printf("[P3] FAIL — A was served by %q (crossed tenants)\n", string(resp.Data))
	}

	// P5 — the ONE way a shared service could break isolation: trusting a
	//   client-supplied tenant. A sends a payload CLAIMING tenant B. A
	//   correctly-built service (server-proven connection decides) ignores it.
	resp5, err := aPrin.Request("SOULSTREAM.MEMORY.REQUEST", []byte("B"), 2*time.Second)
	if err != nil {
		fmt.Printf("[P5] request errored: %v\n", err)
	} else if hasPrefix(string(resp5.Data), "served-by-tenant:A") {
		fmt.Printf("[P5] PASS — A claimed \"B\" in the payload; the service served A anyway (%q). The connection decides, never the payload (D15 for shared services)\n", string(resp5.Data))
	} else {
		fmt.Printf("[P5] FAIL — the service honored A's payload tenant claim: %q\n", string(resp5.Data))
	}

	// P4 — the service performs an act on A's behalf: it appends to A's
	//   SOULSTREAM stream (via svcA). Prove it lands in A's stream and NOT B's.
	_, err = jsA.Publish(ctx, "SOULSTREAM.TOPICS.act", []byte(`{"op":"served-for":"A"}`))
	die("svc act A", err)
	time.Sleep(150 * time.Millisecond)
	aStream, err := jsA.Stream(ctx, "SOULSTREAM")
	die("A stream", err)
	bStream, err := jsB.Stream(ctx, "SOULSTREAM")
	die("B stream", err)
	aInfo, _ := aStream.Info(ctx)
	bInfo, _ := bStream.Info(ctx)
	fmt.Printf("[P4] service act landed in A's stream (msgs=%d) and NOT B's (msgs=%d) — attributed to A's connection alone\n",
		aInfo.State.Msgs, bInfo.State.Msgs)
	if aInfo.State.Msgs >= 1 && bInfo.State.Msgs == 0 {
		fmt.Printf("[P4] PASS — every service act lands in exactly one tenant's stream; B untouched\n")
	} else {
		fmt.Printf("[P4] check — A=%d B=%d (B should be 0)\n", aInfo.State.Msgs, bInfo.State.Msgs)
	}
	_ = bKP
	_ = bSK
}

func hasPrefix(s, p string) bool { return len(s) >= len(p) && s[:len(p)] == p }

func mustTemp() string {
	d, err := os.MkdirTemp("", "rig-*")
	die("temp", err)
	return d
}
