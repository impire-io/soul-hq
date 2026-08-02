# Episode 0056 — M5 ships: the fold learns to live in your house (2026-08-03)

The last milestone on the public-door path: the fold's serve assembly
became the public `embed` package — `Run(ctx, Options)`, value-only
options, the ecosystem's D29 pattern — so the single-binary
distribution can run the fold in-process
(`specs/004-embed-seam/`, branch `004-embed-seam`). With it landed the
authorization-server half the bundled story needs: **RFC 7591 dynamic
client registration** (opt-in `EnableDCR`; discovery grows
`registration_endpoint` by wrapping the certified library's document)
and the **fixed token audience** (opt-in `TokenAudience`, joining
every token's `aud` beside the client id) — exactly what soulstream
018's AS-facing contract demands of whatever stands behind a public
door.

The gate, in `make test` [measured]:

- **The compiler-proof consumer position**: `e2e/embedgate` (module
  `soulfold.invalid/embedgate` — a path outside the fold's namespace,
  so any `internal/` import is a compile error, not a review finding)
  embeds and runs the fold: discovery answers a stock library and
  advertises DCR; a hosted client registers itself; the seeded user
  signs in by passkey ceremony; the access token verifies against the
  fold's JWKS carrying the fixed audience, `oid`, and `roles`; ctx
  cancellation returns Run cleanly.
- **The seam's second consumer, same behavior**: the M4 admission
  rig's fold half switched from `internal/serve` to `embed.Run` — as
  episode 0054 said it would — and the both-arms admission gate stays
  green.
- **One assembly, two entrypoints**: `soulfold serve` now runs through
  the seam itself (the daemon is the seam's first consumer), growing
  `--token-audience` and `--enable-dcr`.

Two supporting moves recorded in the spec: `authtest` went public
(consumers proving a bundled sign-in need the virtual authenticator —
soulnode's fold plane is next in line), and founding stand-ins ride
`Options.SeedUsers`/`SeedClients` (idempotent create-only) until M3's
lifecycle replaces seeding.

What it opened: **everything the bundled distribution needs now
exists** — soulnode can run the fold as a plane and default
`planes.door.auth_issuer` at it, which is that repo's own feature to
land. On the fold's side the public-door path is complete: M1, M2, M4,
M5 shipped in one day; M3 (lifecycle + the bootstrap research) is all
that remains, last by the operator's stated priority.

Reversal condition: none — records a completed build; DCR and the
fixed audience stay opt-in until a measured deployment shape demands a
different default.

Trail: `specs/004-embed-seam/`; the public `embed` and `authtest`
packages; `e2e/embedgate/`; the `004-embed-seam` merge in the soulfold
repo.
