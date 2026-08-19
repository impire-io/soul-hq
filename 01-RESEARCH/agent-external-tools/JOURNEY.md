# JOURNEY — agent-external-tools (opened 2026-08-19)

The investigation, as it happens. Appended to, never rewritten.

## 2026-08-19 — opened

Opened from the operator's evaluation of the running system: "with the
new functionality, don't we need the ability to add external tools, like
MCP servers to the shell? Maybe that's a separate module?"

The survey that produced the bars, before any experiment:

- soulstream-identity v0.9.0 ships the whole broker; the client mirrors
  it (`client/grants.go`: `GrantLinkStart`, `GrantLinkComplete`,
  `GrantAccessToken`, `GrantAccessOnBehalf`, `GrantAccessExchange`,
  `Grants`, `GrantRevoke`).
- `soulstream` (the house) declares no `GrantResources`, and
  `GrantResources` is what enables the `grants.*` ops
  (`embed/embed.go`) — so the broker is off in the product today.
- `GrantResource` is a value struct (name, auth/token/revoke URLs,
  client id + secret, scopes, redirect, exchange config), and
  [`grants.md`](../../02-DESIGN/soulstream-identity/grants.md)'s
  configuration surface states the catalog is "declared configuration,
  no per-user rows (D26's spirit: linking is the user's own act on
  their own prefix)". That sentence is the topic's central tension and
  Bar 4 exists to resolve it in the open.
- The linking ceremony is authorization-code + PKCE. It needs a
  browser, and the shell is the only surface in the ecosystem that has
  one.
- The agent-side seam is workloads' `wrap` and its `mcp_args`; nothing
  there fetches a token today.
