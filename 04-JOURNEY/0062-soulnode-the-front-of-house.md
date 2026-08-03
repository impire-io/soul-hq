# Episode 0062 — The front of house: URLs, the console, and the door that doesn't collide (2026-08-03)

The operator asked the obvious operator question — *where's the admin
URL, and won't the door and the admin site collide?* — and it exposed
a clutch of gaps between "the parts pass their tests" and "a person can
use it." SoulNode's front of house got wired properly
(`specs/006-the-fold-plane/` extended; branch `009-fold-urls`; soulnode
**v0.3.0**, on soulfold v0.3.2).

What landed:

- **Every URL is logged.** `soulnode up` now prints the MCP door, the
  fold sign-in, and the **admin console** (`<fold>/admin`) — and the
  enrolment invite as a working link. `init` prints the token, the
  invite, and points at `up`.
- **The fold is on by default.** `init && up` now lands a person at a
  passkey prompt with an admin console, nothing else to install — the
  bundled experience the whole system was built toward. Off is a
  config flag; `--fold-listen` sets its port.
- **The two services can't collide.** The door and the fold are
  separate listeners; the ceremony refuses a shared address by name
  (ephemeral `:0` exempt), and the log spells out that public mode
  needs two distinct fronted routes (the door's catch-all would
  otherwise swallow the fold).
- **The bare-IP footgun is caught at load.** The fold issuer is the
  WebAuthn RP id, and a browser refuses a bare IP there — so the local
  default issuer is `localhost` (not `127.0.0.1`), and a bare-IP issuer
  refuses at `Verify` with the reason. This was found the only way it
  could be: driving a real Chromium session, not a Go test
  authenticator.

Two bugs the end-to-end browser run surfaced, both fixed upstream and
pinned:

- **The invite link dead-ended.** `<issuer>/login/?invite=…` needs an
  OIDC request a clicked link doesn't have. soulfold grew a standalone
  `/enroll` (v0.3.1, D25a); soulnode points the printed invite there.
- **A phantom duplicate user.** `up` re-seeds the owner every start,
  and `CreateUser` wrote the record before the username index, so a
  duplicate orphaned a second record — two `owner` rows in the console.
  soulfold made `CreateUser` index-first (v0.3.2); the console shows one
  owner.

Verified the whole human path in a real browser against a live
`soulnode up`: click the enrolment link → create a passkey → land in
the admin console as the owner, with a single clean user row. One
binary, a browser, a passkey — the vision executes.

Reversal condition: none — records a completed build and two fixes; the
fold-on-by-default is a stated default (config turns it off), not a
one-way door.

Trail: `specs/006-the-fold-plane/`; soulnode `v0.3.0`; soulfold
`v0.3.1`/`v0.3.2` and episode
[0061](0061-soulfold-the-admin-console.md); design
[lifecycle](../02-DESIGN/soulfold/lifecycle.md) (D25, D25a); the
`009-fold-urls` merge.
