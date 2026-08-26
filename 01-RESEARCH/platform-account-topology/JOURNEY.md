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

## 2026-08-26 — Bar 2: one-act tenant birth and admission — **CONDITIONAL PASS, two required fixes found** [measured]

Rig `rigs/bar2/` reproduces the exact JWT shapes the identity code
produces (cited to `localoperator.go` buildJWT and `mint.go`
claims()+ephemeral) so the server behaviour it measures is the
behaviour the product would get. It isolates the two unknowns the
recon flagged. 5 runs, spreads reported.

**Timing — the one-act birth is fast and clean.**
- Tenant birth (the `$SYS.REQ.CLAIMS.UPDATE` push, one act): **395µs –
  776µs** [measured, 5 runs] — corroborating episode 0110's full-engine
  1.69ms (which also does the vault import). Well under the 100ms bar.
- AUTH `allowed_accounts` amend land: **1.1 – 1.5ms**.
- Callout admission round trip (combined-fix tenant): **2.0 – 3.5ms**.
- create → first *working* token admission is therefore **single-digit
  ms**, against the 5s bar — enormous margin, *once the two fixes below
  are applied*.

**But as-built, `accounts.create` cannot birth a *usable* tenant.**
Two independent gaps, each measured:

1. **The AUTH `allowed_accounts` coupling (Q2a) — CONFIRMED.** A
   callout-issued user for the freshly-created tenant draws
   `nats: Authorization Violation` at connect, every run: the issuer may
   only place users into accounts AUTH lists (D21), and `accounts.create`
   never amends the AUTH JWT (`localoperator.go` touches only the tenant
   JWT). Fix measured sufficient: add the tenant to `allowed_accounts`
   and re-land the AUTH JWT (a second one-act push, 1.1–1.5ms) — Q2b/Q2c
   then admit.

2. **Scoped-mint-on-plain-key = an INERT user (Q1) — a real defect.**
   The mint issues `SetScoped(true)` users (`mint.go` claims()), but
   `accounts.create` installs the tenant signing key with plain
   `SigningKeys.Add` (`localoperator.go:57`), *not* `AddScopedSigner`
   with a permission template. A scoped user whose signing key carries
   no scope template inherits the scoped **sentinel limits (0
   subscriptions / 0 payload)**: the server *admits the connection* but
   it is **inert** — cannot subscribe or publish anything. The signal
   surfaces either at connect or at first subscribe (both
   `maximum subscriptions exceeded`), same 0-limit cause. So episode
   0110's "the existing mint path serves the new tenant the moment the
   op returns" is **false through admission**: the minted user connects
   dead. Fix measured sufficient (Q1c / Q2c): install the tenant signing
   key as a **scoped signer carrying the persona template** (the same
   `personaScope` template `ceremony.go` gives the founding realm) — the
   minted user then gets the correct persona permissions (realm subjects
   and `$JS.API.>` allowed, foreign subjects denied by permissions
   violation).

**The combined fix works, end to end (Q2c) [measured].** A tenant born
with a scoped-signer signing key *and* added to `allowed_accounts`
admits a callout user in 2–3.5ms whose permissions are exactly the
persona template — realm allowed, foreign denied. So the topology's
tenant-birth path is sound; the current code is two small, local edits
short of delivering it.

**Reading**: Bar 2's mechanism passes with margin, but the pre-registered
"one-act tenant birth through the product" is **not** what ships today —
it is a **three-act** sequence the product does not wire at all
(`SystemConn` absent, no client/CLI surface for `accounts.*`), and even
driven directly the authority produces an inert tenant. The topic's
build list gains two concrete, measured items:
(a) `accounts.create` must add the new tenant to AUTH `allowed_accounts`
    (or the design must state where that coupling is performed), and
(b) the tenant signing key must be installed as a scoped signer with the
    persona template, not plain — otherwise every minted tenant user is
    inert.
Neither is a wire change or a core-invariant breach (the reversal
condition stays unfired); both are localised to the identity plane's
tenancy authority. Flagged for the operator's build-order review.
