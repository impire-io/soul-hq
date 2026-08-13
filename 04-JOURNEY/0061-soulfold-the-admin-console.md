# Episode 0061 — The admin console: a browser and your passkey (2026-08-03)

M3 shipped the lifecycle with a JSON admin API (D24) — deliberately no
pages. The operator's push made the gap plain: an admin wants to
click, not curl, and expected a URL to open. So the admin surface
split in two (D25): the JSON API moved to `/api/admin/*` for
automation, and a **server-rendered console landed at `/admin`** for a
person with a browser (`specs/005-the-lifecycle/`, branch
`006-admin-console`, soulfold **v0.3.0**).

What it is, and what held [measured]:

- **Passkey-session auth, no bearer.** A visitor with no session sees
  a login page whose one script runs a WebAuthn *assertion* — a
  session-only ceremony, no relying party, no auth request. On success
  the fold sets the `sf_session` (D11) and the console renders only if
  that user is active and in the `admin` group. This reuses D9's
  single-script exception and adds no new page kind to the sign-in
  flow.
- **The console drives the lifecycle**: list and create people, mint
  their enrolment invites (shown once in the flash, D21), move them
  between groups, disable/enable accounts, register and delete OAuth
  clients — every change CSRF-guarded (the token minted into the
  browser-session record, D13) and landed by POST/redirect/GET.
- **Refusals**: a non-admin's valid passkey is refused at the console
  door; an unauthenticated request never sees the dashboard; a forged
  CSRF changes nothing. The machine API keeps its own admin-role
  bearer check, now under `/api/admin`.
- Verified two ways: a Go gate that drives the whole flow through a
  virtual authenticator against the real HTTP handlers, and a real
  Chromium session with a CDP virtual authenticator — enrol by invite,
  sign in, and operate the dashboard.

The browser run caught what the Go authenticator could not: **WebAuthn
refuses a bare IP as the RP ID.** `127.0.0.1` as the issuer host makes
every ceremony fail with "invalid domain" in a real browser; the test
authenticator doesn't enforce the rule. Recorded as a deployment
constraint (use `localhost` locally, the fronted name in production —
never an IP issuer), surfaced in the quickstart and carried into the
soulnode wiring.

Refactor of record: the browser-session logic (the `sf_session` cookie
+ record) moved to a shared `internal/websession` so the sign-in flow
and the console mint and read it the same way; `BrowserSession` gained
a CSRF field (additive, store D2).

What it opened: soulnode should now log the console's URL beside the
door's and guard the two listeners from colliding — the next episode.

Reversal condition: none — a completed build against the graduated
lifecycle design, extended by D25 which carries its own reasoning.

Trail: `specs/005-the-lifecycle/` (quickstart's day-to-day is now the
console); design [lifecycle](../02-DESIGN/soulstream-idp/lifecycle.md) (D25,
D24 amended); soulfold `v0.3.0`; the `006-admin-console` merge.
