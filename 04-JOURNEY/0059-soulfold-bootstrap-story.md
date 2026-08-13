# Episode 0059 — The bootstrap story: invitation is the only door (2026-08-03)

M3's gating research asked what replaces first-touch enrollment — for
the first admin from nothing, and for everyone after — such that
possession of the fold's store never suffices to enroll. The
hypothesis, pre-registered yesterday with four bars: single-use,
expiring, digest-stored invite tokens as the *only* enrollment right,
with the first one minted by an operator act on the deployment's
state. The rig was the fold itself on its lifecycle branch — the
strongest measurement position: real store, real ceremonies, real HTTP
surface.

All four bars passed, none amended [measured]:

- **The from-nothing ceremony is four counted acts** — serve, seed the
  admin, mint the invite, one browser ceremony that enrolls the
  passkey and signs in — ending in a token whose roles claim carries
  `admin`, verified against published JWKS. The bundled shape
  collapses the operator half into `init`/`up`. The consumed invite
  refuses replay with the record unmoved; a seeded user with no
  invite cannot begin any ceremony — first-touch is deleted, not
  disabled.
- **Invites are D12-honest bearers**: exactly one enrollment per
  invite in 25/25 races of 8 (the CAS consume runs *before* the
  credential binds); expired, forged, and replayed invites refuse with
  zero state change; the bearer appears in no key and no opened
  record while its digest key is scan-findable — the positive control.
- **The store alone cannot enroll**: with the seal seed conceded, all
  29 artifacts recoverable from the opened bucket — keys, fields, and
  digests re-dressed in the token prefix — were presented as invites;
  zero admitted [measured; digest inversion refused by sha256
  preimage, mechanism-argument].
- **The pocket-id audit fixed M3's scope** (docs.pocket-id.org):
  needed — users, groups→roles, invites, client registration, a JSON
  admin API; deferred — per-client allowed-groups, audit store, API
  keys, invite revocation; refused — open registration, LDAP, SMTP,
  branding, custom claims, and pocket-id's own bootstrap (an open
  `/setup` page, first-arriver-wins) — the operator-act invite is
  strictly stronger trust at the same act count [judgment,
  documented].

What it opened: the fold's third design doc,
[lifecycle](../02-DESIGN/soulstream-idp/lifecycle.md) (D20–D24), whose
acceptance criteria are these bars restated as the M3 gate. The M3
build follows immediately — the prototype that measured the bars is
its foundation.

Reversal condition: a supported deployment shape in which the
four-act ceremony demonstrably cannot complete with only the founding
terminal and a browser, or a measured collapse of invite custody to
in-store artifacts — either reopens toward the topic's named fallback
(first-touch as a *stated loopback-only mode*, renamed honestly).

Trail: verdict and topic journey in git history at
`01-RESEARCH/bootstrap-story/` (folder removed by this graduation);
design [lifecycle](../02-DESIGN/soulstream-idp/lifecycle.md); the bar rigs
ride the fold's `make test` (`internal/lifecycle/bars_test.go`,
`internal/serve/bootstrap_test.go`) — permanent gate tests, not
scratchpad.
