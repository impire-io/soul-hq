# How does a person give an agent an external tool, so that every call carries the calling person?

**Component:** ecosystem
**State:** active
**Started:** 2026-08-19

## Abstract

The outbound half of identity already exists and is unreachable. The
grants broker ships in soulstream-identity v0.9.0 — `grants.link.start`
/ `link.complete` / `access` / `list` / `revoke`, four honest lanes
(D34), custody sealed in its own domain (D31), on-behalf-of only as a
bounded delegation against standing consent (D33/C4) — and the product
wires none of it: `soulstream` declares no `GrantResources`, so the
`grants.*` ops are off in the house today [measured, code trace]. The
piece that would make it usable is not the broker, it is the two ends
nobody has designed: **where the tool catalog lives**, and **how a
running agent gets a token at call time without one ever landing in an
`.mcp.json`**.

Answering this closes the industry default the grants design exists to
refuse — one credential in the agent's config, every persona collapsed
into a single remote user — for the case people actually have: a
person, an assistant, and a remote MCP server. Answering it wrongly
either puts a secret back in a config file or turns the resource
catalog into the per-user store D26 refused.

## The question

How does a person add an external tool — a remote MCP server or a
third-party API — to their soulstream and let their agent use it, such
that (a) no outbound credential ever exists in agent, workload, or MCP
client configuration, (b) the remote sees the *calling person*, not the
agent and not a shared service account, and (c) the resource catalog
does not become a per-user store?

## Pre-registered bars

- **Bar 1 — a tool becomes usable without a restart, and without
  disturbing the ones already there.** Protocol: the identity plane at
  its published tag plus a candidate catalog mechanism, against a real
  third-party AS; a resource is added while a continuous probe exercises
  a pre-existing resource's `grants.access` at 5 ms cadence. **Pass:**
  `grants.link.start` serves the new resource with zero process
  restarts, and the probe's maximum gap stays under 50 ms with zero
  failed accesses. **Fail** sends the answer to "adding a tool is an
  operator config act; the shell surfaces linking only" — which is a
  result, not a defeat.
- **Bar 2 — no outbound credential in anything the agent can read.**
  Protocol: a real agent run through `soulstream wrap` against a remote
  MCP server that requires OAuth; after a successful tool call, a
  positive-control grep for the access token *and* the refresh token
  over the harness's configuration files, the wrap invocation's
  environment, the workload's disk, and the process's own argv.
  **Pass:** zero hits, with the planted control string found — the D13
  idiom. Any hit fails the bar outright; there is no acceptable
  location.
- **Bar 3 — the remote sees the caller.** Protocol: two enrolled humans,
  each with an agent, one remote that logs its authenticated subject;
  each agent makes the same tool call on its person's behalf.
  **Pass:** the remote's log attributes each call to that person's own
  remote subject with **zero cross-attributions**; revoking one
  person's grant refuses their next call while the other's keeps
  serving; every on-behalf access audits both personas (the D33 shape).
- **Bar 4 — the catalog's shape is decided by what Bar 1 needed.** Not
  an experiment: the topic must state, in one paragraph and before
  graduating, where the catalog lives (composition config, identity
  plane state, or the record as a registry extension), which component
  owns it, and what it costs. **Pass** requires the statement to name
  the losing options and why they lost.

## Reversal condition

The topic assumes the catalog can become runtime data without breaking
D26's spirit ("declared configuration, no per-user rows"). What would
change our minds, phrased as an observable: **if adding a tool cannot
be expressed without a row keyed by the person who added it** — i.e.
the mechanism that makes Bar 1 pass is a per-user store by another
name — then the answer is that resources stay declared configuration
and the shell surfaces only the linking ceremony, which needs a browser
and has no other home. That outcome closes this topic honestly with a
much smaller shell module and no identity-plane change at all.

## Verdict

<Empty until graduation.>
