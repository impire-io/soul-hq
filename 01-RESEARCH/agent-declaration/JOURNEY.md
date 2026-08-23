# agent-declaration — journey

Topic opened 2026-08-23. The investigation appends here as it happens.

## 2026-08-23 — hypothesis before any experiment

Recon of the shipped surfaces is running (workloads: declaration, runner,
wrap engine, natstest, minter; core: ops, artefacts, work, notify, registry;
imps: schedule channels; identity: mint.ephemeral, embed seam). No spike has
run yet. The rig hypothesis, stated first so the experiments can refute it:

- **One rig, four bars, one variable at a time.** A separate scratch Go
  module (session scratchpad, never committed here) wires the shipped
  packages against an embedded operator-mode NATS server. Bar 1 runs first
  and alone: the wake dispatcher is scratch code, the four sources are
  shipped primitives (notify subscription for mentions, ops consumer for
  topic ops, JetStream message scheduling for ticks, plain subscription for
  external subjects), the outcome op rides wrap's deterministic-id shape
  (UUIDv5 over trigger op id + persona). Expected failure mode, named now:
  the schedule source — if the pinned nats-server predates message
  scheduling, the tick channel needs either a server bump (allowed; not a
  wire change) or a scratch scheduler (which would taint the bar and must
  be recorded as such, not papered over).
- **Bar 2 rides Bar 1's rig** with one change: the agent's instructions
  move from rig constant to a stage-1 artefact; the harness is the scripted
  hermetic one (wrap's own test pattern), so "runs revision A/B" is
  observable in the outcome op's content.
- **Bar 3 swaps the credential path** from workloads' own minter to
  identity's D28 `mint.ephemeral` via the embed seam, with one granted and
  one ungranted tool; refusal must come from the transport/grant layer, not
  rig code.
- **Bar 4 is an enumeration first, a drive second** — every mutation of
  the declare flow named against a shipped surface before the rig performs
  it; any gap is the finding, not a thing to shim silently.

Reversal watch carried from the README: a mutation the rig needs that never
appears as an op; a pass that requires branching on human vs machine.
