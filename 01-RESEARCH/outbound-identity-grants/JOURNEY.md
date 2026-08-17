# outbound-identity-grants — investigation journey (started 2026-08-17)

## 2026-08-17 — the terrain, recovered before any new rig

Read back through episodes 0038/0047 (remote node researched, then built
as soulstream 018, v0.7.0), 0054–0057 (fold M4/M5, the public door, the
folded realm): far more of lane 1 exists than the topic assumed. The
door already walks a hosted client through 401 → RFC 9728 → discovery →
DCR → auth-code+PKCE → token; the fold already admits a passkey user
through soulstream-identity's callout on a real operator-mode server;
the bundled composition (fold plane + door defaulting at it) shipped
with a one-human gate. What is genuinely unmeasured is exactly Bar 1's
content: **two personas through that composition, each admitted as
itself, zero cross-attribution**, plus refusal-on-expiry/revocation
timing through the door. The fold's user ids are persona-shaped by
construction since v0.1.2 (`u-hex` — the folded-realm gate caught the
underscore refusal live).

## 2026-08-17 — first probe: the idp's OP surface, dumped live [measured]

An earlier session's scratchpad rig (`idprig`, embed.Run + authtest
virtual passkey + scripted PKCE dance) re-run against the local
soulstream-idp main (v0.6.0 + the 0103 refresh grant):

- Access token claims: `oid`=`u-4ef68342345d030f` (legal persona slug),
  `preferred_username`, `roles: [engineering, admin]`, `aud` = [client
  id, the fixed `TokenAudience`]. The Entra vocabulary the callout OIDC
  lane (D23/D24) reads — confirmed present, nothing to add for Bar 1.
- Refresh grant: rotating `srt_…` returned; replay of the spent token →
  400 `invalid_grant`. Matches episode 0103's gate.
- **RFC 8693 token exchange: `unsupported_grant_type`** — the idp
  cannot today re-mint a token for a *different* audience from an
  existing session. Audience is deployment-fixed (`TokenAudience`,
  M5's opt-in). Load-bearing for the design: one deployment = one
  resource audience; a persona reaching *several* remote realms/APIs
  from one session needs either a dance per audience or token exchange
  growing in the idp. Named as the likely first implementation slice.
- DCR: 201 with `token_endpoint_auth_method: none` (public client);
  revocation endpoint present; userinfo returns `sub` only;
  introspection requires client credentials.

## 2026-08-17 — Bar 2, provider half: the stand-in dances [measured]

Dex v2.44.0 in docker (memory storage, `reuseInterval: 1s`, static
users alice/bob, one static broker client). The rig's `LinkDance`
scripts the full authorization-code + PKCE ceremony with plain
net/http (login form POST; `skipApprovalScreen`); no browser.

- Linking dance completes: token set with `offline_access` refresh
  token (3/3 runs while iterating).
- Rotation: 3/3 consecutive refreshes each returned a new token.
  One Dex semantic learned the honest way: a refresh inside the 1s
  `reuseInterval` — measured from issuance — is answered as a *retry*,
  same token back (the first run read that as "no rotation"; rig
  timing, not provider behavior; test now clears the window first).
- The spent token refused after the interval: 400
  `invalid_request — Refresh token is invalid or has already been
  claimed by another client`.

Dex's discovery also advertises RFC 8693 token exchange — the
stand-in supports what the idp doesn't yet; useful contrast for the
design's exchange story.

## 2026-08-17 — Bar 2, custody half: CAS survives the stampede [measured]

The broker prototype's pure-logic core (`grantstore.go`: CAS store with
JetStream-KV Update(rev) semantics + the Access loop): 8 concurrent
Access calls on one grant, under `-race`, against a fake with Dex's
measured rotation semantics — 8/8 succeed and the stored refresh token
is still redeemable afterward, in both reuse-window regimes (0s,
500ms). The loser of a CAS race hands out its still-valid access token
and writes nothing — the winner's successor is the line.

Composed against real Dex: linking dance seeds the store, two Access
calls across a persisted rotation both serve, revocation (custody
delete) refuses the next access with grant-not-found. One design fact
worth carrying: the crash window between redeem-at-provider and
CAS-write-of-successor is the one place rotation can still lose the
line; the Access loop writes the successor before returning the access
token to shrink it, and providers' reuse intervals exist for exactly
this — the design should name the residual window honestly.

Still open on Bar 2: the principal-enforcement clause (B refused A's
grant by the identity plane's subject space, never by broker logic) —
that is the NATS-surface half; and the pre-registered real-provider
morning step.

## 2026-08-17 — Bar 3, the delegation core: minted, bounded, audited [measured]

`delegation.go` + its test, `-race` clean: the runner (holding the
signing key the agent never sees — 0082's custody line) mints a
delegation only against standing consent covering every requested
resource and scope; the delegating broker redeems on-behalf-of only
when the delegation verifies, is live, names the server-asserted
caller as actor and the subject as subject, and the resource is within
bounds. Measured: 1 allowed; 4 refusals (no delegation, expired,
wrong caller presenting a stolen delegation, out-of-bounds resource) —
each refusal audited naming both personas (5/5 entries carry both);
consent revocation refuses the next mint with the audit trail intact;
an unconsented resource never mints.

The transport binding still to measure on the NATS rig: `caller` here
is a function argument; on the real surface it must be the D15
subject-token principal, so an agent cannot *claim* an actor. Same
mechanism as Bar 2's principal clause — one rig covers both.

## 2026-08-17 — four repos mapped; three findings reshape the topic

Parallel sweeps of soulstream-identity, -idp, -mcp/-core, -workloads
(full maps in the session transcript). The load-bearing facts:

1. **Bar 1 has a ready rig.** The remote node is shipped code
   (soulstream-mcp v0.1.0), and its `rigtest/` package stands up the
   whole admission edge in-process: NATS with callout, identity plane
   via `embed.Run`, an AS stub written from the frozen 018 AS
   contract. The idp's own `e2e/fold_in_fleet_test.go` already proves
   fold-token-through-callout for one user. Bar 1 = rigtest with the
   AS stub swapped for the real idp (embed seam + authtest virtual
   passkeys), two users, attribution + refusal timing.
2. **The vault cannot custody a rotating secret as it stands** —
   records are immutable by design (Create-only store seam, KV
   create-only writes, "keys are immutable; import under a new name"),
   kinds are a closed enum, and `derive()` requires a public half. A
   grants store therefore needs a numbered decision (a D10
   conversation), and a broker op that *returns* access tokens
   collides head-on with Article I's custody-without-possession — it
   must be argued as its own decision, not ridden in quietly beside
   `ExportSeed`. The callout token store is unsealed (digests only) —
   wrong home for credentials.
3. **The 0082 runner is gone.** The waker landed 2026-08-15 13:01 and
   was cut the same afternoon ("the daemon gives way to the wrapper");
   today's `wrap` holds one long-lived persona credential from env and
   mints nothing — the `mintRunCredential` path is recoverable only
   from history (workloads `715f683`). The asker reaches the run as
   text (`Author` in the notify payload), nothing credential-shaped.
   So Bar 3's minting *home* is an open design question the rig
   deliberately does not decide: candidates are the wrapper regrowing
   a mint lane (0004-wrap.md's written reversal condition), or the
   **subject-signed delegation** — the asking persona signs the
   delegation with its own persona key at mention time, verifiable by
   anyone through the `keys.public` directory (D26), no new trust
   root. The rig measures the delegation *mechanism*; the episode
   argues the home.

## 2026-08-17 — Bar 1: PASS [measured, local rig per the registered protocol]

The rig: the real soulstream-idp through its embed seam (seeded users,
virtual-passkey ceremonies — the fold's own e2e idiom), the real
identity plane (embed.Run, OIDC lane pointed at the fold), the real
remote node (soulstream-mcp v0.1.0) on the operator-mode AUTH topology
replicated from rigtest. No stand-ins anywhere in the trust path.

- **Two personas, each admitted as itself, 2/2**: ada →
  `u-cd00f20fc32ba2a3`, grace → `u-aed28616fbe46d78`; `whoami` names
  each token's own oid; one topic with a turn from each, and the
  independent reader attributes both correctly — **0
  cross-attributions**. No issuer-claim-profile change was needed: the
  fold already speaks the validator's vocabulary (0054's decision
  doing its job).
- **Liveness across bumps**: 5s callout TTL, 10/10 writes across
  2.5×TTL — every server bump re-admitted the living bearer silently.
- **Dead bearer refused in 12.6ms** (bound: 2×TTL = 10s) — the node's
  freshest-wins pool re-probes a changed bearer immediately; refusal
  is effectively instant, not TTL-paced.
- **The revocation asymmetry, measured**: revoking the refresh token
  at the fold did NOT refuse the still-valid access token (admitted
  1.5×TTL later) — a callout-class verifier checks signature and
  claims, so revocation bites at access-token `exp` (1h, package
  constant). The 0103 watch item ("per-client token-lifetime knobs
  stay deferred until a milestone demands them") — this composition
  is the milestone demanding it: a bounded outbound revocation story
  needs a short, configurable access-token lifetime at the fold.

BYON confirmation (authorized, not required by the registered
protocol) left as an optional strengthening step.

## 2026-08-17 — Bars 2 and 3, transport halves: PASS [measured, -race]

Operator-mode server enforcing a scoped-signer template that grants
exactly `identity.{{account-subject()}}.{{name()}}.grants.>` on the
caller's own prefix (the D15/D25 mechanism verbatim; one bug earned:
a user signed by a signing key must carry `IssuerAccount`). Four rows:

- alice reaches her own grant — token served.
- **bob publishing to alice's grants subject is refused by the
  server** (`Permissions Violation for Publish`, timeout to bob) —
  the broker service's delivery log shows alice's subject reached it
  exactly once (her own request). The refusal happens in the
  transport, with nothing for a prober to learn.
- the agent presents a valid delegation on its OWN subject —
  on-behalf access to alice's grant allowed, bounded.
- bob presents the *stolen* (validly signed, unexpired) delegation on
  his own subject — refused as an actor mismatch, because the caller
  is the server-proven subject-token principal, never a claim. Both
  on-behalf rows audited naming both personas (1 allowed / 1 refused).

With this, Bar 2's clauses are all measured except the pre-registered
real-provider morning step, and Bar 3's mechanism is fully measured
(pure-logic clauses + the transport actor binding). Full rig suite
green under -race in one run.

Smaller facts that will matter in the design: the identity callout
pins exactly one issuer and hardcodes Entra claim names (fine for the
bars — Dex is outbound-only, it never touches the callout); the idp
has no RFC 8693 / client-credentials / per-request audience and
silently drops unknown scopes (per-remote audiences = the exchange
gap, now measured from both ends); idp revocation deletes the session
record but a JWT stays signature-valid until `exp` for a
callout-class verifier — effective revocation is refresh-revocation
plus ≤1h drain; `mint.ephemeral`'s `tags` ride into user claims —
inert today, a natural carrier for delegation facts; core's client
uses static `nats.Token` while the node uses `TokenHandler` (the
refresh seam a broker consumer would need); and **no outbound HTTP
machinery exists anywhere in core or workloads** — the broker's
provider-facing half is genuinely new surface, which is exactly why
it should live behind the identity plane rather than in every agent.
