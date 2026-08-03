# Enrollment and the lifecycle

**Graduated from research:** bootstrap-story, 2026-08-03 —
[episode 0059](../../04-JOURNEY/0059-soulfold-bootstrap-story.md).
**Realized by:** M3 (the lifecycle) on the
[roadmap](../../03-IMPLEMENTATION/ROADMAP.md). The store these records
ride on is [store-and-key-lifecycle](store-and-key-lifecycle.md)
(D1–D8, D16–D19); the ceremonies they gate are
[session-and-ui](session-and-ui.md) (D9–D15).

Who may enroll, who decides, and what an admin is: the founding
refusal (passkeys, not passwords) extends to enrollment — possession
of an invite, never presence in the store or arrival at a page, is the
only enrollment right. Every mechanism below passed a pre-registered
bar in the graduating research [measured]; the acceptance criteria at
the end are those bars restated as M3 gate tests.

## Decisions

### D20 — Enrollment rides invites; there is no open lane

A passkey registration ceremony begins only against a live invite for
exactly that user. A user with no credentials and no invite cannot
begin any ceremony; a user with credentials asserts. First-touch
enrollment (M2's loud interim) is deleted, not disabled — no mode,
no flag, no fallback.

Reasoning: enrollment is the trust decision; presence in the store
must never make it (Bar 1: a seeded, credential-less user's begin
refuses at HTTP [measured]). Recovery and multi-device are the same
mechanism — a fresh invite adds a credential to an already-enrolled
user.

### D21 — An invite is a bearer secret: digest-stored, single-use, expiring

The token (`sfi_…`) is shown exactly once at mint. The store keys the
record by `invite.<hex(sha256(token))[:32]>` (D12's rule) holding only
the target user, `consumed`, and `expires_at` (D5 authoritative). The
consume is a D4 CAS flip executed *before* the credential binds — the
race loser is refused structurally, and a ceremony that fails after
consumption burns its invite (the recovery is a fresh invite, never a
reusable one).

Reasoning: exactly-once 25/25 races of 8; expired/forged/replayed all
refuse with zero state change; the bearer appears nowhere in keys or
opened records while the digest key is scan-findable (the positive
control) [measured].

### D22 — The first invite is an operator act; every later one is an admin op

Possession of the deployment's state mints the bootstrap invite
(`soulfold invite` on the state dir; `embed.Options.InviteSink` for
embedding parents) — the same trust that founded the fold, mirroring
soulidentity's first-key story. The from-nothing ceremony is four
counted acts (serve, seed admin, mint invite, one browser ceremony
that enrolls and signs in), collapsing to init/up + one browser act in
the bundled shape [measured]. After bootstrap, invites come from the
admin surface (D24). The audited alternative — pocket-id's open
`/setup`, first-arriver-wins — is refused: same act count, weaker
trust.

### D23 — Groups are names; the roles claim is membership

A group record is a name and nothing else — it never carries
permissions (constitution II; the consumer's side declares what a role
grants). The user record's `groups` field is the lived membership; the
roles claim is its names, plus the pre-M3 `roles` field read-only
forever (D2: additive, never removed). A membership change surfaces in
the next issued token [measured] — revocation propagates at token
lifetime, exactly like the callout's bound.

### D24 — The admin surface is a JSON API authenticated by the fold's own tokens

`/admin/*` — users (create/list/disable, set groups), groups
(create/list), invites (mint — the one response that ever carries a
bearer), clients (register/list/delete). No pages: D9's inventory
stays {login, error}. The bearer is a fold-issued access token whose
roles claim names `admin`, verified against the fold's own published
keys, issuer, and expiry — the fold trusts exactly what it tells
everyone else to trust. Non-admin tokens and bare requests refuse
[measured].

The audit table that fixed this scope (Bar 4, against pocket-id's
surface): **needed** — users, groups→roles, invites, client
registration, the JSON API; **deferred** — per-client allowed-groups,
queryable audit store, API keys, invite revocation (TTL bounds it);
**refused** — open registration, open `/setup`, LDAP sync (the fold
stands where Entra stands; it fronts no directory), SMTP, branding,
custom claims.

## Acceptance criteria (the M3 gate inherits these)

1. From-nothing bootstrap: a fresh fold to a signed-in admin (roles ∋
   admin, JWKS-verified) in the four counted acts; the bootstrap
   invite refuses replay with the user record unmoved.
2. No open lane: a credential-less user without an invite cannot begin
   a ceremony, at the HTTP surface.
3. Invite honesty: exactly-once consumption under racing consumers;
   expired/forged/replayed refuse with zero state change; the bearer
   digest-stored (positive-control-verified).
4. Membership propagation: an admin-surface group change surfaces in
   the next issued token.
5. Admin surface authz: admin-role bearers pass; non-admin bearers and
   bare requests refuse.
