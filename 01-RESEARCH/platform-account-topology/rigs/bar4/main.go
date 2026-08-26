// Bar 4 rig — the multi-tenant human, expressible and loud. Three clauses:
//
//	C1 (token lane): one human, two API tokens into tenants A and B, admits
//	    into both; the two admitted identities are isolated (neither reaches
//	    the other's account subjects) — cross-tenant action refused by the server.
//
//	C2 (OIDC ambiguity): D24's roleFor refuses a subject whose roles claim names
//	    two declared tenants — deterministically, regardless of claim order.
//	    roleFor is reproduced verbatim from issuer.go:241-267 and fuzzed over
//	    orderings; the code-trace argument (a length count cannot depend on
//	    order) is the mechanism, the fuzz is the demonstration.
//
//	C3 (persona-name scope): the identity vault keys persona keys as
//	    "persona/<user>" with NO account component (service.go:524 +
//	    vault.GeneratePersonaKey), so a SHARED platform vault makes persona
//	    names GLOBAL across tenants — first-owner-wins, the second tenant's
//	    same-named user cannot sign. Reproduced behaviorally; the two
//	    resolution options are stated for the operator.
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/nats-io/jwt/v2"
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

// scopedTenantJWT builds a tenant account whose signing key IS a scoped signer
// carrying the persona template — Bar 2's measured combined-fix shape.
func scopedTenantJWT(op nkeys.KeyPair, acctPub, name, signingPub string) string {
	ac := jwt.NewAccountClaims(acctPub)
	ac.Name = name
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
	ac.Limits.JetStreamLimits = jwt.JetStreamLimits{MemoryStorage: -1, DiskStorage: -1, Streams: -1, Consumer: -1}
	ac.Limits.Conn = -1
	tok, err := ac.Encode(op)
	die("tenant jwt "+name, err)
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

type tokenBinding struct {
	signer  nkeys.KeyPair
	account string
	user    string
}

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

func landAccount(sys *nats.Conn, token string) error {
	msg, err := sys.Request("$SYS.REQ.CLAIMS.UPDATE", []byte(token), 5*time.Second)
	if err != nil {
		return err
	}
	var resp struct {
		Error *struct{ Description string } `json:"error"`
	}
	if err := json.Unmarshal(msg.Data, &resp); err == nil && resp.Error != nil {
		return fmt.Errorf("resolver refused: %s", resp.Error.Description)
	}
	return nil
}

func mustTemp() string {
	d, err := os.MkdirTemp("", "rig-resolver-*")
	die("temp dir", err)
	return d
}

func main() {
	fmt.Println("=== Bar 4: the multi-tenant human ===")
	clauseC1()
	clauseC2()
	clauseC3()
}
