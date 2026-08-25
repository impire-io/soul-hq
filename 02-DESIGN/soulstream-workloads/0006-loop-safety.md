# 0006 — Loop safety: the wake budget

**Graduated from** research topic `loop-safety`
([episode 0128](../../04-JOURNEY/0128-ecosystem-loop-safety.md)); all four
bars measured. This document specifies the budget every wake admission
enforces — the mechanism that makes agent-wakes-agent deployments (and
therefore topic-wake colonies, design
[`0005`](0005-agent-declaration.md) §7) safe to run.

## 1. What must exist [V]

A **wake budget**, enforced at wake admission — in `wrap`'s
`handleWake` before the harness is invoked, and at the same seam in any
future wake dispatcher (0005's topic/schedule/subject kinds). Two
composed parts, both computed from the materialised topic view alone (no
state beside the log, no core wire change):

- **The window floor** — refuse a wake when the persona already authored
  `K` `turn.post` ops in this topic within window `W`. Authorship is
  mechanical (the client's persona signs), so this floor cannot be evaded
  by body text, forged ids, or MCP self-posts `[V — the rig's id-evading
  cascade halted at exactly 2K]`.
- **The depth bound** — refuse a wake whose outcome would sit more than
  `D` provable hops from a chain root, where a hop is the `WakeOpID`
  UUIDv5 binding (`UUIDv5(triggerOpID + "/" + persona)`) checked against
  candidate ops in the view — never stream order `[V — 421/421 chains
  resolved, halt at exactly D outcomes]`.

The window is the enforcement floor; the depth bound is the tight control
and the diagnostic on chains that ride the deterministic id. Neither
alone suffices `[V — depth evaded by self-posts (393 turns past a D=4
gate); the window is coarser than depth on provable chains]`.

## 2. Refusal semantics [V]

- **Op-less.** A refused wake posts nothing — no outcome, no testimony.
  A refusal that posts an op is a wake source and re-arms the loop
  `[V — the harness-slot variant produced a 312-turn failure ping-pong]`.
- **Loud.** One structured log line per refusal
  (`wake_refused`, with persona, trigger op, and the legible reason
  including the numbers: hops vs D, count vs K/W) — the 0083 precedent.
- **A delay, not a loss.** Nothing is acked away: the mention stays in
  the bounded inbox, the outcome id stays deterministic, and a slid
  window re-admits on a later catch-up. The record is the position,
  unchanged.
- **Placement.** The gate runs after the self-skip and the
  outcome-existence pre-check, before the harness invocation — before
  the outcome obligation attaches. Never in the harness slot; never as
  a retry-exhaustion failure.

## 3. Configuration surface [D]

Wrap config and CLI flags, with the same knobs growing into 0005's
declaration `wake` section as a `budget` block (schema growth, never
machinery):

```
budget:
  max_hops: 4          # D — provable-chain depth bound; 0 disables
  window:
    max: 8             # K — own turn.posts per topic per window; 0 disables
    per: 10m           # W
```

Defaults on, with the values above `[judgment — generous against every
legitimate flow measured, orders of magnitude under the danger numbers]`.
Disabling either knob is explicit configuration; disabling both is the
operator saying "unbudgeted", and the wrapper logs that standing once at
startup.

## 4. The walker [V]

The ancestry walk is part of the shipped surface, not rig code: from any
op, resolve the provable parent (the unique candidate whose `WakeOpID`
binding matches), chain to root, report depth and the chain itself. It
serves the depth gate and the operator's question "why was this wake
refused / where did this cascade come from". Ambiguity (two parent
matches) is a reported error, never absorbed `[V — zero ambiguous matches
across 421 resolved chains]`.

## 5. Acceptance criteria

1. An uncooperative two-agent mention cycle (scripts that always mention
   each other) halts at exactly `D` outcomes with one loud op-less
   refusal; no harness cooperation involved.
2. The same cycle with outcomes posted by the agents' own clients
   (arbitrary ids) halts within the window bound `2K`; refusals loud,
   op-less.
3. A human-rooted delegation chain shorter than `D` completes with zero
   refusals; every outcome chains to the root through the walker.
4. A topic-wake colony of `N` agents halts within `N·K` outcomes; a
   fresh op after the halt draws refusals, not silence and not a new
   cascade; shutdown stays clean.
5. A refusal never posts an op of any kind, and a budget set to 0/0
   reproduces today's behavior byte-for-byte.

## 6. Sequencing (resolves 0005 §7)

Mention-wake `wrap` grows the budget first. Topic-wake colonies unblock
when this budget ships **in the admission path of whatever dispatches
them** — the gate is the build, not further research. The reconciler and
any agent-wakes-agent deployment (0126's "doubly" note) inherit the same
requirement.

## 7. Open, named [O]

- **Deterministic ids on the MCP arm**: if `post_turn` through the tool
  door learns the wake's outcome id (the runner passes it; the post is
  idempotent under it), depth becomes reliable for self-posts too and
  the window floor stops being the only guard there. Improvement, not
  prerequisite.
- **Per-pair budgets** (A-wakes-B rate distinct from A's topic total) —
  by demand, if a real colony chafes on the topic-global window.
- **Refusal surfacing in the shell** (a parked-wake face) — a product
  composition question, not this mechanism's.

Reversal condition: as episode 0128 — an outcome invisible to both the
depth walk and the authorship count (neither deterministic-id nor
persona-authored) evades the floor and reopens the topic; an admission
point that cannot materialise the view it budgets on moves the mechanism
or reopens the topic.
