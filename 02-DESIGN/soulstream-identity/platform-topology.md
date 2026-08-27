# soulstream-identity — the platform-account topology (D46–D49)

*Graduated from research topic `platform-account-topology` (episode
[0133](../../04-JOURNEY/0133-ecosystem-platform-account-topology.md);
all four bars measured 2026-08-26, the two reserved decisions taken by
the operator 2026-08-27). Extends [`tenancy.md`](tenancy.md) (D35) with
the hosting shape the operator asked for: one dedicated
platform-services account beside per-tenant accounts. Everything here
sits behind the 0071 focus gate — designed now so the demand has an
answer, built when the product demands multi-tenancy. Decisions
continue the global numbering: D46–D49.*

## D46 — The platform-services account and the export seam

A deployment MAY host the identity surface (and its sibling platform
services) in a dedicated platform account, serving many tenant
accounts. The mechanism is D14/D15's already-designed extension path,
adopted here as the deployment shape:

- The platform account **exports** the surface as a service:
  `<prefix>.identity.*.>` with `account_token_position = P+2`
  (1-based; P = prefix token count — position 2 at the bare default).
- Each tenant account **imports** the surface naming **its own
  account public key** at the token position.

The protection is enforced at **import-definition time** [measured,
Bar 1]: the server refuses any import whose account-token position
carries a key other than the importing account's own, so a tenant
cannot even construct a route to another tenant's surface subjects —
isolation is structural, not a runtime check. Zero service-side
decisions are added; the service's own D15 subject check is unchanged.
Stated honestly: the refusal grade is **no-responder (503)** — the
structural absence of a route — not a permissions error naming the
boundary.

**Acceptance criteria**: the Bar 1 matrix re-measured on the real
surface — a tenant's own-key round trip through the import; a
foreign-key attempt refused by the server with the delivery-log proof
(the surface never sees a foreign account token); a malicious import
of another tenant's key refused at account resolution; the
without-position negative control documented in the deployment notes
as the misconfiguration to refuse.

## D47 — Tenant birth completes admission (D35 amended)

Bar 2 measured that `accounts.create` as built births an unusable
tenant, twice over. Two amendments to the D35 authority, both
measured sufficient, neither a wire change:

1. **The tenant signing key is installed as a scoped signer carrying
   the deployment's persona template** — never a plain
   `SigningKeys.Add`. Every mint issues `SetScoped(true)` users
   (D5: permissions live in the signing key's scope, never the JWT);
   a scoped user on a plain key inherits the scoped-sentinel limits
   (0 subscriptions / 0 payload) and is **admitted but inert**
   [measured]. The template is the same persona scope the founding
   ceremony renders (one source, so the two cannot drift — the
   spec-010 SC-004 discipline extended to tenancy).
2. **`accounts.create` amends AUTH `allowed_accounts` with the new
   tenant** in the same op — a second one-act resolver push on the
   system connection (1.1–1.5ms measured) — or the design of the
   deployment must state explicitly where that coupling is performed.
   Without it the callout issuer may not place users into the new
   tenant (D21's explicit list, never `*`), and the token/OIDC lanes
   refuse with `Authorization Violation` [measured].

With both, create → first working token-lane admission is single-digit
milliseconds against Bar 2's 5s bound [measured: birth 395–776µs,
admission 2.0–3.5ms].

*As built (2026-08-27, identity `447ec6b`, episode
[0134](../../04-JOURNEY/0134-identity-tenants-born-admissible.md)):*
both amendments landed in the LocalOperator authority. The one-source
template is `client.PersonaScopePubAllow`/`PersonaScopeSubAllow`
(prefix-aware; the embed seam passes the deployment's prefix) — the
founding ceremony adopts the same export on its next touch. The AUTH
amend is lookup → add → re-land complete (idempotent; tenant lands
first so the between-acts window fails closed), and an empty
`AuthAccount` skips the coupling — no callout half, no admission list
to maintain (its reversal condition lives in the episode). Live-test
proof: store → **usable** admission 2.77ms, out-of-scope publish
drawing the permissions violation, the tenant read back from AUTH's
stored JWT [measured]. The ProviderAPI arm's D47 parity remains A8's
named BYON residue.

**The product-side residue — closed 2026-08-27** (episode
[0135](../../04-JOURNEY/0135-ecosystem-the-residues-close.md), spec
`012-tenants-in-the-house`): the node wires `SystemConn` (a SYS user
minted in memory), the client mirrors the op family
(`AccountCreate/Resolve/Accounts/Suspend/Resume`, identity `df8e4a3`),
`soulstream account …` is the hand, and the embedded server's dir
resolver persists runtime tenants across restarts with a
create-if-absent seed protecting AUTH's runtime amendments. The
op-family e2e measures this design's acceptance criterion end to end:
create → usable token-lane admission **11.2ms** through the sealed
surface and callout [measured]. The provider arm carries both D47
halves too (`f6a1a33`: the group's `Scope`, the `jwt_settings`
authorization read-union-write); its live BYON run (2026-08-27,
episode 0136) caught two defects before passing — the authorization
amend must write the object **whole** (`auth_users`/`xkey` carried
forward beside the unioned `allowed_accounts`, valid under merge or
replace patch semantics; a userless AUTH is refused by name, since
the JWT law forbids accounts without users) — then measured sound:
births 3.3–4.4s inside Bar 2's 5s bound, the scoped round trip alive
in the newborn account, AUTH carrying both tenants on read-back
[measured]. A BYON system's control-plane view exposes no
client-reachable URL; the round-trip clause takes the deployment's
own URL as operator input.

**Acceptance criteria**: Bar 2 re-measured through the real op family —
`accounts.create` then token-lane admission with a **usable** user
(subscribes, publishes its persona scope, foreign subjects refused as a
permissions violation) within 5s, ≥6 runs reported as spreads; an
existing tenant's continuous probe uninterrupted; zero usable
half-tenants across pre-creation probes.

## D48 — Per-tenant persona custody (operator decision, 2026-08-27)

Persona-key custody is **per tenant**: each tenant's persona keys live
in a vault bucket on **that tenant's own JetStream**, reached through
the platform service's per-tenant connection (D49). The platform
vault retains **platform custody only** — the operator key, the AUTH
signing key, the tenant signing keys (`team/<name>`). The sealed-store
discipline is unchanged (D31: a domain per bucket, sealed to the
deployment's first key, CAS on write).

The ground, as decided: persona information is not used outside its
tenant — persona keys and their owner bindings serve attribution
*within* a tenant. What follows:

- **Persona names are account-scoped by construction.** `persona/daan`
  in tenant A and in tenant B are different keys in different buckets.
  The silent cross-tenant shadowing Bar 4 measured on a shared vault
  (first-owner-wins with an anti-probing refusal that names no owner)
  is eliminated rather than loudened — there is no collision to
  report.
- **D26's directory read becomes a per-tenant read**: `keys.public`
  resolves personas of the caller's own tenant. This is the
  decision's rationale made mechanical, not a restriction bolted on.
- **A tenant's custody travels with its account**: suspension freezes
  it with the account, and a future tenant export/move carries the
  persona keys because they live in the tenant's own JetStream.

**Reversal condition**: a consumer needing cross-tenant persona
resolution (observable: a verification blocked on a foreign tenant's
key, recorded as an issue) reopens as a *directory federation*
question — never by returning persona keys to a shared vault.

## D49 — Shared-service disciplines, and the OIDC multi-tenant stance

Two disciplines every shared platform service MUST follow — both are
existing invariants extended, not new machinery [measured, Bar 3]:

- **Per-tenant connections.** A shared service reaches each tenant
  through a connection scoped to that tenant's account (its own
  service user there, or a D47-minted credential). One connection
  never spans tenants; the service shares process and code, never
  credentials or reach.
- **The connection decides the tenant — never the payload.** The
  acting tenant is derived from which server-proven connection
  delivered the request (D15 applied to services); a request argument
  naming a tenant is ignored. A service that honored a payload tenant
  claim would be the one measured way to break isolation across an
  otherwise-sound topology [measured: the claim probe, ignored].

Under these, cross-tenant reads and writes are unreachable by
construction and every service act lands in exactly one tenant's
stream with its attribution intact.

**The OIDC multi-tenant stance** (operator decision, 2026-08-27):
D24's ambiguity refusal **stands** — a subject whose roles claim names
two declared tenants is refused, deterministically and
order-independently (a length count over declared matches; fuzzed
[measured]). No tenant-selection machinery is built before a real
consumer demands it (constitution III); the token lane already serves
the multi-tenant human — two tokens, both admitted, server-isolated
[measured, Bar 4 C1].

**Reversal condition**: a real OIDC human assigned into two tenants
and blocked at connection (observable: recorded as an issue) reopens
tenant selection as a new D-decision — the candidate shapes named at
decision time are a per-tenant audience and a connect-time account
hint; neither is designed further until then.

## Build order (behind the 0071 focus gate)

D47 first — it closes a live correctness gap in the shipped
`accounts.*` family (an inert tenant is a defect whether or not the
platform topology is ever deployed). Then D48 (per-tenant custody —
touches the vault's bucket layout, cheapest before more personas
exist), then D46 (export configuration — deployment shape, no code),
D49's disciplines landing with the first shared-service build. The
product's `SystemConn` wiring and an `accounts.*` client surface ride
whichever product cycle adopts multi-tenancy.
