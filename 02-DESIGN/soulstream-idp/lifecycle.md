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
(`soulstream-idp invite` on the state dir; `embed.Options.InviteSink` for
embedding parents) — the same trust that founded the fold, mirroring
soulstream-identity's first-key story. The from-nothing ceremony is four
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

### D25a — an invite link enrols standalone, at `/enroll`

The enrolment invite (D21) is delivered as `<issuer>/enroll?invite=…`
— a **standalone page**, not `/login`. `/login` only exists inside a
relying party's OIDC request (it needs an auth-request id); an invite
link a person clicks has none, so it lands on `/enroll`, which runs a
session-only registration ceremony (`Begin(username, "", invite)` →
create → `Finish`) and shows a confirmation. The invite is the whole
capability; a cross-origin submission is refused (D13's Origin wall),
and the ceremony consumes the invite exactly as any registration does
(D21). After enrolling, the user signs in through any RP with their
passkey; an admin's convenience session lands them in the console.

Reasoning: without this the printed invite URL dead-ends at "missing
auth request" — enrolment would only work if some RP happened to
forward the invite through `/login`, which no invite link can assume
[measured: the bare `/login?invite=` path errors; `/enroll?invite=`
completes].

### D25 — the admin surface is two halves: a machine API and a human console

D24's JSON API moves to `/api/admin/*` and gains a sibling: a
server-rendered **console at `/admin`** for a person with a browser.
The split is by consumer, not by capability — both call the one
`lifecycle.Service`.

- **The console authenticates by passkey session, not by bearer.** A
  visitor with no session sees a login page whose one script runs a
  WebAuthn *assertion* (a session-only ceremony — no relying party, no
  auth request); on success the fold sets an `sf_session` naming the
  user and the console renders only if that user is in the `admin`
  group and active. This reuses D11's browser session and D9's
  single-script exception; it adds no new page kind to the sign-in
  flow (login/error), because the console is its own surface.
- **State-changing POSTs carry the session's CSRF token** (minted into
  the browser-session record, D13's rule) and land back on the
  dashboard (POST/redirect/GET). `SameSite=Lax` is the outer wall.
- **The one bearer a console response ever carries is the enrolment
  invite**, shown once in the flash after minting (D21).
- **Custody and refusals are unchanged**: a non-admin's valid passkey
  is refused at the console door; an unauthenticated request never
  sees the dashboard; the machine API keeps its own bearer check.

Reasoning: an operator wants to click, not curl — the fold is a human
tool, and Entra (the thing it stands in for, constitution II) has a
portal too [judgment]. Server-rendered, no SPA, no framework
(constitution III): the console is HTML the fold prints, gated on the
same passkey the rest of the system already trusts. Measured: an
admin enrols, signs into `/admin` with their passkey, and creates a
user / mints an invite / sets groups / disables an account from the
browser; a non-admin passkey is refused; a forged CSRF changes
nothing [measured].

Deployment note (soulstream / any two-service host): the console shares
the fold's listener with the OIDC endpoints, so it is reachable
wherever the fold's issuer is. It does **not** share the MCP door's
listener — those are separate ports/services and must front to
distinct public routes (soulstream logs both URLs and refuses a shared
listen address).

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
   bare requests refuse (the machine API); and the human console
   (D25): an admin's passkey session reaches the dashboard, a
   non-admin's does not, an unauthenticated request sees only the
   login page, and a forged CSRF changes nothing.
