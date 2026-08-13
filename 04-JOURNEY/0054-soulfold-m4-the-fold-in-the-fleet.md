# Episode 0054 — M4 ships: the fold joins the fleet (2026-08-02)

The consumer-position proof the fold was founded toward: **a browser
user's passkey sign-in becomes a NATS admission**. The rig
(`e2e/` in the soulfold repo, its own module) runs the whole chain —
authorize → passkey ceremony → code → access token → sentinel + token
connect → soulidentity's auth callout → minted scoped credential — on
a real operator-mode server, with soulidentity imported **at its
published tag v0.1.0** through its public embed seam and configured
with exactly what any deployment gives it: an issuer URL and an
audience (`specs/003-fold-in-the-fleet/`, sequenced before M3 by the
operator's public-door priority, episode 0052).

The gate, both arms in `make test` [measured]:

- **The fold arm**: a user enrolled by a real WebAuthn ceremony signs
  in, and their access token admits the connection; the token's role
  value (`engineering`) resolves against the deployment's declared
  role binding (the ENG account's scoped signing key imported under
  that name); the server enforces the role's template — in-scope round
  trip green, out-of-scope publish drawing a permissions violation;
  the audit attributes lane=oidc, the role, and the issuer. Zero
  per-person acts anywhere: the only declared facts are the role
  bindings.
- **The stub arm**: the *identical* gate with the fold swapped for an
  Entra-shaped stub issuer — same discovery, same JWKS shape, same
  claim vocabulary. Indistinguishability demonstrated, not asserted
  (constitution II).
- **The refusal**: on both arms, a token whose role names nothing
  declared refuses, and the refusal is in the audit.

One decision surfaced by the seam and recorded in the spec: **the
fold's access tokens speak Entra's claim vocabulary** — `oid`,
`preferred_username`, `roles` — because the verifier of record
(soulidentity's D23/D24 lane) keys subjects by `oid` with no `sub`
fallback, and constitution II's test is literally "unable to tell the
fold from Entra" [mechanism-argument]. `User.Roles` arrived additively
(store D2), seeded for now, group-derived when M3 lands.

What it opened: **soulnode's public door mode is unblocked** — the
external AS its OAuth story waits on now demonstrably admits users
into a NATS deployment; and M5's embed lift has its consumer already
running (the rig's fold half switches from `internal/serve` to the
public seam when it exists).

Reversal condition: none — records a completed proof; if soulidentity
ever grows a `sub` fallback or multi-issuer dispatch (its named
extension), the claim-vocabulary decision may relax additively, never
by removing a claim.

Trail: `specs/003-fold-in-the-fleet/` and the `e2e/` module in the
soulfold repo; soulidentity's design
[auth-callout](../02-DESIGN/soulstream-identity/auth-callout.md) (D23–D24,
unchanged — that is the point); the `003-fold-in-the-fleet` merge.
