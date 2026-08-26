# platform-account-topology — investigation journal (opened 2026-08-26)

## 2026-08-26 — Code recon before any rig runs [measured, code trace]

Read before building, per the method. Five facts the bars will lean on,
each traced to code, none yet run:

1. **`accounts.create` births an island.** The LocalOperator authority
   builds the tenant JWT complete — identity, name, signing key, JS
   limits, conn limit — and pushes it to `$SYS.REQ.CLAIMS.UPDATE`
   (`soulstream-identity/internal/accounts/localoperator.go:50-71`).
   Nothing touches the AUTH account JWT: `allowed_accounts` (explicit,
   never `*` — D21) does not learn the new tenant. Bar 2's coupling
   question is therefore already answered in code — the token/OIDC lane
   cannot place users into a created tenant; the rig will show it live
   and measure the amended path.
2. **Scoped mints onto a plain key — a suspected admission break.**
   Every mint path issues scoped, permission-less user JWTs
   (`SetScoped(true)`, `internal/mint/mint.go` `claims()`); but the
   authority installs the tenant signing key PLAIN
   (`SigningKeys.Add`, no `AddScopedSigner` — `localoperator.go:57`),
   and the tenant JWT carries no persona scope template at all. If the
   server refuses a scoped user on a non-scoped key, "the mint path
   serves the new tenant the moment the op returns" (episode 0110) has
   never been true through admission — the live test connects a
   hand-signed plain user, not a minted one
   (`internal/accounts/accounts_live_test.go:127-143`). To be measured,
   not assumed.
3. **No client surface for tenancy.** `accounts.*` exists only
   service-side (`internal/service/ops_accounts.go`); the public
   `client` package and the CLI expose nothing. The product never wires
   `SystemConn` (`soulstream/node/node.go:234` area), so the ops are
   dark in the house — Bar 2 spikes both.
4. **Persona collision is owner-silent by design.** Cross-account
   same-name signing refuses with a deliberately unprobing message
   ("has no persona key") — first owner wins, the refusal must not
   probe the vault (`internal/service/service.go:512-541`). Bar 4's
   "naming the first owner" clause collides with that anti-probing
   choice; the owner IS resolvable one op away via the open
   `keys.public` directory read. The rig measures what is; the bar
   clause gets judged honestly against it.
5. **D24 ambiguity is order-independent by construction.** `roleFor`
   collects every declared role match before deciding; >1 refuses
   (`internal/callout/issuer.go:241-267`). Bar 4 proves it
   behaviorally, both claim orders.

Rig scaffolding adopted from `accounts_live_test.go` (dir resolver,
operator mode, `$SYS.REQ.CLAIMS.UPDATE`) and `soulstream/ceremony/ceremony.go`
(AUTH/REALM JWT shapes, the persona scope template, `client.Segment =
"identity"` — the account token rides at position 2 bare, P+2 with a
prefix, exactly D14's export grammar).

## 2026-08-26 — Bar 1: cross-account export preserves the principal proof — **PASS** [measured]

Rig `rigs/bar1/` (preserved in-topic). Operator-mode server, dir
resolver, a PLATFORM account exporting the identity surface as
`identity.*.>` with `AccountTokenPosition: 2`, tenants importing it.
The "surface" is a decision-free echo responder — it makes zero
authorization decisions, so any isolation observed is the **server's
alone**, which is precisely the bar's "zero service-side decisions"
clause. Deterministic across 3 runs with fresh random keys.

**The mechanism, established against nats-server's own
`TestJWTAccountProtectedImport`** [mechanism-argument, then measured]:
`account_token_position` enforces at **import-definition** time. An
account may only define an import whose subject carries *its own*
account public key at the token position; the server rejects an import
naming any other key. So a tenant cannot even *construct* a route to
another tenant's surface subject — isolation is structural, not a
runtime check.

- **T1 PASS** — tenant A reached the surface at
  `identity.<A-key>.alice.status`; the responder was delivered exactly
  that subject, A's own key stamped at position 2.
- **T2 PASS** — A naming tenant B's key
  (`identity.<B-key>.bob.status`) got **no-responder (503)**; A's valid
  import (`identity.<A-key>.>`) does not cover B's subject, so no route
  exists and the responder never saw it.
- **T3 PASS** — a malicious tenant (EVIL) that *defines* an import of
  B's key was refused by the server at account resolution; its request
  reached **no-responder**, the responder never saw B's key.
- **T5 CONTROL CONFIRMED** — against a second platform account
  exporting the same surface **without** `AccountTokenPosition`, a
  malicious import of B's key *is* valid and EVIL2 reached that surface
  carrying B's key. This is the exact breach the position prevents —
  proving the position, not some unrelated misconfiguration, is the
  load-bearing control.
- **T4 PASS** — the delivery log across the whole run shows the
  position-protected surface was delivered *only* subjects carrying A's
  own key; no foreign key ever reached it.

**Reading**: Bar 1 met. The D15 principal proof extends across the
account boundary by export configuration alone (`AccountTokenPosition`
= the P+2 D14/D15 named), server-enforced, zero service-side decisions
— the platform-services-account topology's central seam is real and
measured. The one honest narrowing: the rig proves the *transport*
proof in isolation with an echo responder; the real service's own D15
subject check (unchanged since M3) stacks on top and is not re-tested
here. The refusal grade is **no-responder (503)**, not a loud
permissions violation — structural absence of a route, which is the
strongest possible refusal (the caller cannot express the request at
all) but is worth stating plainly: a probing tenant sees a timeout-class
silence, not an error naming the boundary.
