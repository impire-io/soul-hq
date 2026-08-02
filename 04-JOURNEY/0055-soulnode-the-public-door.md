# Episode 0055 — The public door opens (2026-08-03)

Phase 2's second half, gated on soulfold upstream, unblocked hours
earlier by the fold's M4 admission proof (episode 0054): `planes.door`
grew `public_url`, `auth_issuer`, and `auth_audience` additively, and
with them a soulnode serves the whole hosted-client OAuth story
(`specs/005-public-door/`, branch `005-public-door`). The wiring is
deliberately thin — upstream's node already carried public mode
(soulstream 018: RFC 9728 resource metadata, the WWW-Authenticate
challenge, AS-agnostic by contract), and the identity plane's OIDC
lane already validated external issuers (soulidentity D23) — SoulNode
composes them: the config's three new fields feed `PublicURL` /
`AuthIssuer` to the door and `OIDCIssuer` / `OIDCAudience` to the
embed surface. Composition, not invention (constitution I).

The gate, in `make test` [measured]:

- **From nothing to a session**: a scripted hosted client knowing only
  the door URL walks 401 challenge → resource metadata (advertising
  exactly the configured public URL and AS) → AS discovery → dynamic
  client registration → authorization-code + PKCE → token; the bearer
  forms an MCP session and `whoami` names the token's subject; the
  audit attributes lane=oidc and the resolved role (`realm` — the
  founding's own declared role binding).
- **Coexistence and refusals**: the founding token still opens the
  door in public mode; a token naming an undeclared role and a garbage
  bearer form no session.
- **Config honesty**: the three fields are a package deal (partial
  declarations refuse; public mode with the door disabled refuses);
  they survive the Save/Load round-trip; the door listener stays
  loopback — HTTPS is deployment fronting, and Phase 3's tsnet gate is
  untouched.

The AS in the gate is the upstream contract's stand-in (`rigtest`),
which is the point: the door is AS-agnostic, soulfold the intended
default rather than a dependency — the fold's own admission proof is
episode 0054's gate. The full bundled story (the fold in-process,
`auth_issuer` defaulted at it) is soulfold M5's wiring, next.

One consumer-caught bug fixed at the root: `listen: 127.0.0.1:0`
pre-flighted a random port while nats-server read port 0 as its
default 4222 — the probe and the server disagreed, and concurrently
running test packages collided on 4222 [measured: three consecutive
full-parallel runs green after mapping 0 to the server's random-port
spelling].

Reversal condition: none — records a completed build; the tsnet
decision keeps its own measurement gate (built only if host fronting
measures insufficient for the audience).

Trail: `specs/005-public-door/` (spec, quickstart with the fronting
runbook); upstream design
[remote-mcp-node](../02-DESIGN/soulstream/extensions/remote-mcp-node.md);
the `005-public-door` merge in the soulnode repo.
