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

## 2026-08-19 — seam reconnaissance: Bar 2 kills the obvious answer

Before any rig, a code trace over the two seams the answer has to live
in. It cost an hour and eliminated the implementation everybody would
have written first, which is what the method is for.

**The obvious answer, stated so it can be refuted.** A tool call needs a
token; `wrap` already renders the harness's MCP client configuration
**per run** — `writeMCPConfig`, mode 0600, into the run dir, with the
comment "it may carry a credential" already on it
(`wrap/harness.go`) [measured, code trace]. So: fetch `grants.access`
at wake, write the token into that per-run config beside the soulstream
door, let the harness call the remote itself. Short-lived, per-wake,
never persisted between runs.

**Bar 2 refuses it, and Bar 2 is right.** The harness runs with
`cmd.Dir = spec.RunDir` and `MCP_CONFIG` pointing into that same
directory [measured, code trace] — the agent reads that file, by
construction. A token written there is a credential in something the
agent can read, which is the *exact* wording of the bar, and the bar was
pre-registered as pass/fail with no acceptable location. A per-run file
is a better place than a checked-in `.mcp.json`; it is not zero, and the
bar does not grade on a curve.

Worth naming precisely, because the two are easy to conflate: the
agent's **own** NATS credential already lives on disk where the agent
can read it, and that is fine — it is the agent's identity, the thing it
is supposed to hold. Bar 2 is about the **remote system's** token, which
is somebody else's. `sanitizedEnv` already strips every `SOULSTREAM_*`
from the inherited environment so nothing of the host's leaks in
(`wrap/harness.go`) — the discipline this bar extends is one the repo
already keeps.

**What that leaves.** If the token cannot reach the agent, the agent
cannot call the remote directly — so the call has to go through
something that holds no token and fetches one per invocation on the
caller's own admission. That is a **door that forwards tool calls**, not
an injector. Which lands precisely on the chokepoint D37 named and did
not build: "the waker's wake decision and **the door's tool-call
forwarding**"
([`tenancy.md`](../../02-DESIGN/soulstream-identity/tenancy.md)). The
seam this topic shares with
[`guardrail-human-end`](../guardrail-human-end/README.md) is therefore
not a coincidence to note at graduation — it is where the answer lives.

**And the catalog question got concrete.** The broker holds
`resources map[string]Resource`, built once in `New(...)` and never
written again — no mutex, because nothing mutates it [measured, code
trace: `internal/grants/grants.go`]. So a runtime catalog is not a
configuration knob away; it is either a mutable, locked map plus a new
op (the route that strains D26's "no per-user rows", which Bar 4 must
argue rather than assume), or the catalog lives somewhere that is not
the broker and resources stay declared configuration.

**Consequences for the bars, recorded before the next step rather than
after it:** Bar 2 is now expected to pass *by construction* under the
forwarding-door shape and to fail under any injection shape, which
makes it a design discriminator rather than a measurement — it will
still be run as written, against a real agent, because a bar that is
argued and not measured is not a bar. Bar 1 is unchanged. Bar 3 is
unchanged and becomes the interesting one: a forwarding door has to
carry the *calling person* to the remote, not the agent, and that is
`AccessOnBehalf` with a bounded delegation, which already exists.

**Open, for the operator:** the shape this points at is a door with an
outbound half, and soulstream-mcp is parked (episode 0071) with stdio
MCP the choice of record for this iteration. Whether the forwarding door
is the stdio door growing an outbound side, or the parked remote door
waking up, is a scope decision this topic should put to the operator
before it spends a rig on either.
