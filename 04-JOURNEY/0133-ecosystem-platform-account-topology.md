# Episode 0133 — The platform-account topology: measured sound, two fixes short (2026-08-26 → 2026-08-27)

The operator's question — can the Soulstream services live in one
dedicated platform account, serving many per-tenant accounts? — opened
as research topic `platform-account-topology` and closed in two days
with all four pre-registered bars measured on live operator-mode
servers. The answer is **yes, the topology fits the core invariants**:
the topic's reversal condition (a forced wire change, a mutual import
across the cycle guard, a privileged identity tier, or an export path
that loses the D15 proof) never fired. The seam episode 0064 named S1
around — a realm *is* an account — extends to a hosting shape by
configuration, not architecture.

**The bars, measured** (rigs preserved in the topic through its life;
full history in git):

- **Bar 1 — the export seam: PASS** [measured, 3 runs]. A platform
  account exporting `identity.*.>` with `account_token_position = P+2`
  is protected at **import-definition time**: a tenant can only import
  its *own* account key at the token position (confirmed against
  nats-server's own protected-import test), so it cannot construct a
  route to another tenant's surface at all. Refusal grade no-responder,
  zero service-side decisions, and a negative control proving the
  position — not some adjacent configuration — is the load-bearing
  mechanism (without it, a malicious import of a foreign key routes).
- **Bar 2 — tenant birth: CONDITIONAL PASS** [measured, 5 runs]. The
  one-act birth is 395–776µs (corroborating 0110's 1.69ms full-engine
  figure) and the *fixed* path admits a callout user in 2–3.5ms — but
  as-built, `accounts.create` births an unusable tenant, twice over:
  it never adds the tenant to AUTH `allowed_accounts` (admission draws
  `Authorization Violation`), and it installs the signing key **plain**
  while every mint issues `SetScoped(true)` users, so the minted user
  inherits the 0-subscription/0-payload scoped-sentinel limits —
  **admitted but inert**. 0110's "the mint path serves the new tenant
  the moment the op returns" is thereby refuted through admission
  [measured]. Both fixes measured sufficient.
- **Bar 3 — isolation through shared services: PASS** [measured, 3
  runs, real provisioned `SOULSTREAM` streams]. Under per-tenant
  connections, a tenant-A principal cannot read or write tenant B's
  data (unreachable by construction — cross-account JetStream has no
  addressing), cannot reach the service's B-connection, and a
  client-supplied tenant claim is ignored: the connection decides,
  never the payload. Every service act lands in exactly one tenant's
  stream.
- **Bar 4 — the multi-tenant human: PASS with findings** [measured, 3
  runs]. The token lane serves multi-tenant humans today — two tokens,
  both admitted, server-isolated — so 0064-S2 costs nothing there. The
  OIDC lane's D24 ambiguity refusal is deterministic and
  order-independent (a length count; fuzzed across orderings). And the
  vault keys persona keys as `persona/<user>` with no account
  component, so a shared platform vault **silently shadows** same-named
  personas across tenants — exactly the silent shadowing the bar ruled
  out; failed as-built, resolution required.

**Refuted along the way**: 0110's serves-immediately claim (above); and
the assumption that `account_token_position` is a runtime check — it is
an import-*definition* constraint, which is stronger (the foreign route
never exists).

**The decisions — taken by the operator 2026-08-27**: per-tenant
persona custody (each tenant's persona-key bucket on its own JetStream,
reached through the service's per-tenant connection; the platform vault
keeps platform custody only — the ground: persona information is not
used outside its tenant), and the OIDC two-tenant refusal **stands**
(constitution III — no tenant-selection machinery until a real blocked
consumer demands it; the token lane is the multi-tenant human's lane).

What it opened: design
[`platform-topology.md`](../02-DESIGN/soulstream-identity/platform-topology.md)
(D46–D49) — the export seam as deployment shape, the two tenant-birth
fixes as D35 amendments, per-tenant persona custody, and the
shared-service disciplines — all behind the 0071 focus gate, built when
the product demands multi-tenancy. The product-side residue is named
there: `SystemConn` unwired, no client/CLI surface for `accounts.*`.

Reversal condition: the design's own conditions stand (D48's
cross-tenant-resolution consumer; D49's blocked OIDC human). For the
topology itself, re-stated from the topic: a deployment class where the
export path cannot preserve the D15 proof (observable: a tenant
admitted to a subject carrying a foreign account token, or reachability
requiring wildcard grants) returns services to per-tenant accounts,
D14's prefix answer standing.

Trail: design
[`02-DESIGN/soulstream-identity/platform-topology.md`](../02-DESIGN/soulstream-identity/platform-topology.md);
the concluded topic (removed on graduation, full history in git —
opened `1acf493`, bars `c6b8558`/`fc5e976`/`68b1602`/`b349e1d`, verdict
`1880e85`); episodes [0064](0064-ecosystem-the-platform-turn.md) (S1/S2,
the direction), [0107](0107-ecosystem-platform-tenancy-guardrails.md)
(the tenancy design this extends),
[0110](0110-identity-the-tenancy-set-builds.md) (the `accounts.*` build
whose admission claim this corrected),
[0112](0112-ecosystem-the-canonical-form-breaks-clean.md) (A10, the
account key in the record).
