# Episode 0063 — The console gets a face (2026-08-03)

The admin console (episode 0061) worked but wore the plainest possible
markup. The operator's nudge — *this can look a lot better* — was
right, so the fold's whole browser surface got a real visual language
(`internal/webstyle`; soulfold **v0.3.3**).

- **One theme, shared.** `internal/webstyle` carries the fold's dark
  design as a single CSS string (no asset pipeline — constitution
  III), descended from the impire.io / soulsystem system: dark world,
  mono eyebrows, a violet accent with the rose edge that marks the
  fold (the door). The sign-in, enrolment, error, and admin pages all
  render through it, so they read as one product.
- **The console is now a designed surface**: a brand bar with the
  `admin` pill and the signed-in identity, a real people table with
  status pills and hover, group chips (admin highlighted), two-column
  cards for the create/groups forms, and an OAuth-clients table — all
  on the panel/glow treatment the website uses. Focus rings, pill
  buttons, sensible spacing.
- **Behaviour is byte-identical**: same routes, same passkey-session
  auth, same CSRF-guarded POSTs, same handlers. Only the markup and a
  new style package changed; the console, login, enrol, and e2e gates
  stay green, and the D9 page inventory is unchanged.

Verified in a real browser at desktop and mobile widths: the login
card, the enrolment page, and the full dashboard (people, groups,
clients) against a live fold.

Reversal condition: none — a presentation change over unchanged
behaviour.

Trail: `internal/webstyle`, the restyled `internal/adminui` and
`internal/ui` templates; soulfold `v0.3.3`; builds on episode
[0061](0061-soulfold-the-admin-console.md).
