# When agents wake agents, what bounds the cascade — can the room enforce a budget from the record alone?

**Component:** ecosystem
**State:** active
**Started:** 2026-08-25

## Abstract

Loop safety is the named successor topic with three measured exhibits: the
very first wake's reply @-mentioned its asker and fired a fresh notify
(episode 0082, G7); the waker's first fault trial ran away — failure turn →
notify → wake → failure turn, forever (episode 0083); and the
agent-declaration rig showed the topic-wake self-loop appears on day one
unless the persona's own ops are excluded (episode 0126). Every shipped
guard is local and self-referential — self-mention never wakes (design 0004
§7), failure taps only the asker (0004 §7), the topic-wake author filter
(design 0005 §3), the per-wake retry budget (0082 G6) — and none can see a
cycle of length ≥ 2 or fan-out amplification. Design 0005 §7 makes this
topic the colony gate: a prerequisite for any deployment where declared
agents wake each other. A decisive answer unlocks topic-wake colonies; a
decisive failure keeps agent-wakes-agent deployments barred with the bar
stated instead of assumed.

## The question

Can the room bound every agent-wakes-agent cascade — cycles and fan-out —
at wake admission, using only what the record already carries, with zero
cooperation from the inhabitants, while human-rooted delegation chains of
legitimate depth complete untouched?

The mechanism (causal-depth budget, hop counter carried on the notify,
per-pair rate limit, or another shape) is the investigation's output, not
its input; the bars below test the property, not the candidate.

## Pre-registered bars

Protocol common to all bars: a scratch rig in the session scratchpad
(embedded server, real component libraries — the agent-declaration rig's
lineage), two declared agents A and B that mention-wake each other, plus a
colony-shaped topic-wake variant for Bar 4. Harnesses are script harnesses
driven by prompt/template only — no rig-specific code in the waker path
beyond the mechanism under test.

- **Bar 1 — the causal chain is readable from the record.** For every wake
  in the two-agent rig seeded by one human-authored root mention, the
  admission point can reconstruct the wake's ancestry — triggering notify
  op → the outcome op that produced it → that wake's trigger, back to the
  human root — from the record alone. Pass: every wake's chain resolves to
  the root, unambiguously, with several concurrent mentions in flight (the
  0082 correlation lesson: no stream-order anchoring). Fail: any wake whose
  trigger cannot be identified without runner-local state.
- **Bar 2 — the budget halts the uncooperative cycle.** With both harnesses
  scripted to always @-mention the other (maximally uncooperative
  inhabitants), the room-side mechanism halts the A↔B cascade: total wakes
  from the root stay ≤ a bound pre-computed from the configured budget
  before the run, every refusal is op-less and loud (0083's precedent), and
  the parked state is observable to the operator. Fail: wakes continue past
  the bound, or the halt required the inhabitant's cooperation.
- **Bar 3 — legitimate delegation survives the same gate.** A human-rooted
  A→B→reply chain within the configured budget completes under identical
  enforcement: every admitted wake ends in exactly one attributed outcome
  op, and no refusal fires. This bar discriminates a budget from the blunt
  alternative — "agents never wake agents" — which would pass Bar 2 and
  fail here.
- **Bar 4 — fan-out is bounded, and the unbounded danger is measured
  first.** In the colony-shaped rig (one op tapping N ≥ 2 topic-wake agents
  whose outcomes tap others), first record the unenforced growth — the
  danger as a number, not a hypothesis — then, with enforcement on, total
  wakes from one root stay ≤ the pre-computed function of the budget and
  the run ends parked-and-loud, not wedged. Fail: growth exceeds the bound,
  or the rig deadlocks instead of refusing.

## Reversal condition

If causal ancestry cannot be reconstructed from the record alone
(observable: Bar 1 fails — a wake whose trigger op is unresolvable without
runner-local state), budget-as-record-fact is the wrong shape and loop
safety must live in stateful admission infrastructure beside the log —
which reopens the thin-waker standing (episode 0082, G1) rather than
patching it. If any bar's halt turns out to depend on harness behavior
(observable: a prompt-only change to the script harness reopens a halted
loop), the mechanism is not enforcement and the direction is abandoned,
not weakened.

## Verdict

<Empty until graduation. Filled by /research-graduate: PASS/FAIL per bar with the
honest numbers, each load-bearing claim tagged [measured] / [mechanism-argument]
/ [judgment].>
