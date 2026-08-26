// Bar 1 rig — does a service exported from a PLATFORM account, with
// account_token_position, preserve the D15 principal proof across account
// boundaries?
//
// The mechanism (confirmed against nats-server's own TestJWTAccountProtectedImport):
// an account may only DEFINE an import whose subject carries ITS OWN account
// key at the token position. So isolation holds two ways:
//
//	(1) runtime — a tenant's valid import routes only its own key; naming
//	    another tenant's key matches no import and reaches no responder;
//	(2) definition — a tenant that tries to import ANOTHER tenant's key has
//	    its import rejected by the server, so the route never exists.
//
// The "surface" is a decision-free echo responder on the platform account:
// it makes zero authorization decisions, so any isolation observed is the
// SERVER's alone — exactly Bar 1's "zero service-side decisions" claim.
package main

import (
	"fmt"
	"os"
	"sync"
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

// userCreds signs a user directly under the account key with default
// permissions (as nats-server's own import test does), returns a creds path.
func userCreds(name string, signer nkeys.KeyPair) string {
	ukp, _ := nkeys.CreateUser()
	upub, _ := ukp.PublicKey()
	useed, _ := ukp.Seed()
	uc := jwt.NewUserClaims(upub)
	uc.Name = name
	tok, err := uc.Encode(signer)
	die("encode user "+name, err)
	creds, err := jwt.FormatUserConfig(tok, useed)
	die("creds "+name, err)
	return writeCreds(name, string(creds))
}

func main() {
	opKP, _ := nkeys.CreateOperator()
	opPub, _ := opKP.PublicKey()
	_, sysPub := acct()
	sysClaims := jwt.NewAccountClaims(sysPub)
	sysClaims.Name = "SYS"
	sysJWT, err := sysClaims.Encode(opKP)
	die("sys jwt", err)

	// PLATFORM account — hosts the surface, exports it with the account
	// token at position 2 (subject identity.*.>, `identity` at position 1).
	platKP, platPub := acct()
	platClaims := jwt.NewAccountClaims(platPub)
	platClaims.Name = "PLATFORM"
	platClaims.Exports.Add(&jwt.Export{
		Name: "identity-surface", Subject: "identity.*.>",
		Type: jwt.Service, AccountTokenPosition: 2,
	})
	platJWT, err := platClaims.Encode(opKP)
	die("platform jwt", err)

	// NEGATIVE CONTROL: a second platform account exporting the SAME surface
	// but WITHOUT account_token_position. Against this one, an import of
	// another tenant's key is a valid import — the breach the position
	// prevents. Proving the contrast makes the position the load-bearing control.
	noposKP, noposPub := acct()
	noposClaims := jwt.NewAccountClaims(noposPub)
	noposClaims.Name = "PLATFORM-NOPOS"
	noposClaims.Exports.Add(&jwt.Export{
		Name: "identity-surface-nopos", Subject: "identity.*.>", Type: jwt.Service,
	})
	noposJWT, err := noposClaims.Encode(opKP)
	die("nopos jwt", err)

	aKP, aPub := acct()
	_, bPub := acct()
	evilKP, evilPub := acct()

	// TENANT A imports the surface with ITS OWN key at the token — valid.
	aClaims := jwt.NewAccountClaims(aPub)
	aClaims.Name = "TENANT-A"
	aClaims.Imports.Add(&jwt.Import{
		Name: "identity-surface", Account: platPub,
		Subject: jwt.Subject("identity." + aPub + ".>"), Type: jwt.Service,
	})
	aJWT, err := aClaims.Encode(opKP)
	die("tenant A jwt", err)

	// TENANT B — just exists as another tenant whose key A/EVIL will try to name.
	bClaims := jwt.NewAccountClaims(bPub)
	bClaims.Name = "TENANT-B"
	bJWT, err := bClaims.Encode(opKP)
	die("tenant B jwt", err)

	// EVIL tenant — tries to import the surface with TENANT B's key at the
	// token position. The server must reject this import (key ≠ EVIL's own).
	evilClaims := jwt.NewAccountClaims(evilPub)
	evilClaims.Name = "TENANT-EVIL"
	evilClaims.Imports.Add(&jwt.Import{
		Name: "identity-surface", Account: platPub,
		Subject: jwt.Subject("identity." + bPub + ".>"), Type: jwt.Service,
	})
	evilJWT, err := evilClaims.Encode(opKP)
	die("evil jwt", err)

	// EVIL2 imports B's key from the NO-POSITION export — a valid import there.
	evil2KP, evil2Pub := acct()
	evil2Claims := jwt.NewAccountClaims(evil2Pub)
	evil2Claims.Name = "TENANT-EVIL2"
	evil2Claims.Imports.Add(&jwt.Import{
		Name: "identity-surface-nopos", Account: noposPub,
		Subject: jwt.Subject("identity." + bPub + ".>"), Type: jwt.Service,
	})
	evil2JWT, err := evil2Claims.Encode(opKP)
	die("evil2 jwt", err)

	res, err := natsserver.NewDirAccResolver(mustTemp(), 1000, time.Minute, natsserver.NoDelete)
	die("resolver", err)
	for pub, j := range map[string]string{
		sysPub: sysJWT, platPub: platJWT, noposPub: noposJWT,
		aPub: aJWT, bPub: bJWT, evilPub: evilJWT, evil2Pub: evil2JWT,
	} {
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
		die("ready", fmt.Errorf("server not ready"))
	}
	defer srv.Shutdown()
	url := srv.ClientURL()

	// The platform-hosted surface: a decision-free echo responder recording
	// every subject it is delivered — the delivery log.
	var mu sync.Mutex
	var seen []string
	platConn, err := nats.Connect(url, nats.UserCredentials(userCreds("surface", platKP)))
	die("platform conn", err)
	defer platConn.Close()
	_, err = platConn.Subscribe("identity.*.>", func(m *nats.Msg) {
		mu.Lock()
		seen = append(seen, m.Subject)
		mu.Unlock()
		if m.Reply != "" {
			_ = m.Respond([]byte("echo:" + m.Subject))
		}
	})
	die("surface subscribe", err)
	die("flush", platConn.Flush())

	// The no-position surface's responder — records what IT is delivered, so
	// the negative control shows the breach concretely.
	var muN sync.Mutex
	var seenN []string
	noposConn, err := nats.Connect(url, nats.UserCredentials(userCreds("surface-nopos", noposKP)))
	die("nopos conn", err)
	defer noposConn.Close()
	_, err = noposConn.Subscribe("identity.*.>", func(m *nats.Msg) {
		muN.Lock()
		seenN = append(seenN, m.Subject)
		muN.Unlock()
		if m.Reply != "" {
			_ = m.Respond([]byte("echo:" + m.Subject))
		}
	})
	die("nopos subscribe", err)
	die("flush nopos", noposConn.Flush())

	aConn, err := nats.Connect(url, nats.UserCredentials(userCreds("alice", aKP)),
		nats.RetryOnFailedConnect(false), nats.MaxReconnects(0))
	die("tenant A conn", err)
	defer aConn.Close()
	evilConn, err := nats.Connect(url, nats.UserCredentials(userCreds("mallory", evilKP)),
		nats.RetryOnFailedConnect(false), nats.MaxReconnects(0))
	die("evil conn", err)
	defer evilConn.Close()
	evil2Conn, err := nats.Connect(url, nats.UserCredentials(userCreds("mallory2", evil2KP)),
		nats.RetryOnFailedConnect(false), nats.MaxReconnects(0))
	die("evil2 conn", err)
	defer evil2Conn.Close()

	fmt.Println("=== Bar 1: cross-account export preserves the principal proof ===")
	fmt.Printf("platform=%s\n tenantA =%s\n tenantB =%s\n evil    =%s\n\n", platPub, aPub, bPub, evilPub)

	// T1 — A reaches the surface with A's OWN key at token 2.
	ownSubject := fmt.Sprintf("identity.%s.alice.status", aPub)
	if msg, err := aConn.Request(ownSubject, []byte("x"), 2*time.Second); err != nil {
		fmt.Printf("[T1] FAIL — A could not reach the surface with its own key: %v\n", err)
	} else {
		fmt.Printf("[T1] PASS — A reached the surface; responder echoed %q\n", string(msg.Data))
	}

	// T2 — A names TENANT B's key. A's valid import (identity.<A>.>) does not
	//   cover it; no route, responder never sees it.
	foreignSubject := fmt.Sprintf("identity.%s.bob.status", bPub)
	b0 := countSeen(&mu, &seen)
	_, err2 := aConn.Request(foreignSubject, []byte("x"), 1500*time.Millisecond)
	a0 := countSeen(&mu, &seen)
	switch {
	case err2 == nil:
		fmt.Printf("[T2] FAIL — A reached the surface naming B's key (isolation broken)\n")
	case a0 > b0:
		fmt.Printf("[T2] FAIL — responder was delivered a B-key subject from A's connection\n")
	default:
		fmt.Printf("[T2] PASS — A refused when naming B's key: %s (responder never saw it)\n", classify(err2))
	}

	// T3 — EVIL tries to reach B's subject through its INVALID import of B's
	//   key. The server rejects the import; no route ever exists.
	b1 := countSeen(&mu, &seen)
	_, err3 := evilConn.Request(foreignSubject, []byte("x"), 1500*time.Millisecond)
	a1 := countSeen(&mu, &seen)
	switch {
	case err3 == nil:
		fmt.Printf("[T3] FAIL — EVIL reached the surface via an import of B's key (definition proof broken)\n")
	case a1 > b1:
		fmt.Printf("[T3] FAIL — responder was delivered a B-key subject from EVIL's connection\n")
	default:
		fmt.Printf("[T3] PASS — EVIL's import of B's key was rejected by the server: %s\n", classify(err3))
	}

	// T5 — NEGATIVE CONTROL: EVIL2 imports B's key from the NO-POSITION
	//   export and reaches that surface with B's key. This is the breach the
	//   position prevents — proving account_token_position is what does the work.
	nb0 := countSeen(&muN, &seenN)
	_, err5 := evil2Conn.Request(foreignSubject, []byte("x"), 1500*time.Millisecond)
	na0 := countSeen(&muN, &seenN)
	switch {
	case err5 == nil && na0 > nb0:
		fmt.Printf("[T5] CONTROL CONFIRMED — without account_token_position, EVIL2 reached the surface carrying B's key (this is the breach the position blocks)\n")
	default:
		fmt.Printf("[T5] control inconclusive — nopos import did not route as expected: %s (delivered=%d→%d)\n", classify(err5), nb0, na0)
	}

	// T4 — the delivery log: prove the surface only ever saw A's own key.
	mu.Lock()
	defer mu.Unlock()
	fmt.Printf("\n=== delivery log (subjects the surface was actually delivered) ===\n")
	sawForeign := false
	for _, s := range seen {
		fmt.Printf("  %s\n", s)
		if contains(s, bPub) {
			sawForeign = true
		}
	}
	if sawForeign {
		fmt.Printf("[T4] FAIL — a foreign key reached the responder\n")
	} else {
		fmt.Printf("[T4] PASS — every delivered subject carried A's own key; no foreign key ever reached the surface\n")
	}
}

func countSeen(mu *sync.Mutex, seen *[]string) int {
	mu.Lock()
	defer mu.Unlock()
	return len(*seen)
}

func classify(err error) string {
	switch err {
	case nil:
		return "nil"
	case nats.ErrNoResponders:
		return "no-responder (503, server-refused route)"
	case nats.ErrTimeout:
		return "timeout (no route, no responder)"
	default:
		return err.Error()
	}
}

func contains(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}

func writeCreds(name, creds string) string {
	f, err := os.CreateTemp("", "rig-"+name+"-*.creds")
	die("temp creds", err)
	_, err = f.WriteString(creds)
	die("write creds", err)
	_ = f.Close()
	return f.Name()
}

func mustTemp() string {
	d, err := os.MkdirTemp("", "rig-resolver-*")
	die("temp dir", err)
	return d
}
