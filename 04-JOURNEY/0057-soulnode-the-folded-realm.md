# Episode 0057 — The folded realm: one binary, real people (2026-08-03)

The night's arc closes where it aimed: **a SoulNode now carries its own
sign-in**. `planes.fold` (opt-in) runs the deployment's passkey-first
OIDC provider — soulfold, by tag, through its public embed seam —
inside the node's process, storing on the node's JetStream over its own
bypass-lane creds with its seal seed under `<state>/fold/`; and public
door mode, given no authorization server, **defaults at the bundled
fold** — soulfold M5's distribution story, composed
(`specs/006-the-fold-plane/`, branch `006-the-fold-plane`).

The gate, in `make test` [measured]:

- **The whole human chain, zero external services**: the door's
  resource metadata names the bundled fold; a scripted browser
  registers by DCR, signs the founding persona in by passkey ceremony
  (first touch enrolls, exactly the fold's M2 story), and the access
  token opens an MCP session at the door — `whoami` naming the
  person's fold identity, the audit attributing lane=oidc role=realm.
- **The default wiring measured**: fold enabled + `public_url` set +
  nothing else → `auth_issuer`/`auth_audience` resolve to the fold's
  issuer and `soulnode-<realm>`; the identity plane's OIDC lane
  validates against exactly that pair.
- **Old realms stay honest**: a pre-fold state dir loads with the
  plane off (the config block's absence means disabled — a pointer,
  deliberately, so nothing sprouts on upgrade); enabling it there
  refuses by name. The founding inventory grew to 21 artifacts (the
  fold's creds are founding matter).
- **Ordering is load-bearing and recorded**: the fold serves before
  the identity plane, whose OIDC validator discovers its issuer at
  startup — a bootstrap circle broken by sequence, not by retry.

Two consumer-proven additions landed upstream on the way, each tagged:
`embed.Options.NATSCreds` (soulfold v0.1.1 — an operator-mode parent's
store connection is authenticated, and the seam only spoke a bare URL)
and **persona-shaped user ids** (v0.1.2 — `u-hex`: the OIDC lane names
the downstream persona by the token's oid, and soulstream's identity
grammar refuses underscores; the folded-realm gate caught it live).

What it opened: the vision sentence has its public half — `init && up`,
front the door, and *other people* sign in with a passkey and work in
the realm. What stands between this and handing it to someone: the
fold's M3 (real lifecycle behind the first-touch-enrollment interim,
with the bootstrap-story research), release/distribution polish, and
the physical-authenticator runbook (a human act, pending).

Reversal condition: none — records a completed composition; the
fold-by-default-at-init question is named for the release story, not
decided tonight.

Trail: `specs/006-the-fold-plane/`; soulfold `v0.1.1`/`v0.1.2` and its
episode [0056](0056-soulfold-m5-the-embed-seam.md); the
`006-the-fold-plane` merge in the soulnode repo.
