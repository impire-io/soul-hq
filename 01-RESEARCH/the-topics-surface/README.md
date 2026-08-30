# The topics surface — what does the shell owe a realm where topics are many and some are machinery?

**Component:** shell
**State:** active
**Started:** 2026-08-30

## Abstract

The shell's conversations screen renders `topic.Board` raw: a flat rail
showing name and a one-word state, ordered lexicographically by path —
effectively arbitrary — with no activity ordering, no grouping beyond
the archived fold, and no use of the organizing material the record
already carries (`Tags`, `Parent`, the curator-maintained `dormant`
state, `Discover`). Design 0003 §7 [O2] anticipated the scaling
question; the declare surface made it acute: every declared agent
auto-starts a home topic (0009 §7) that lands on the human list as an
indistinguishable row, as does the placements topic. Yet the home topic
is not a conversation at all — it is the agent's operational room
(placement work item, non-record wake outcomes, failure self-reports,
the narrow credential's only writable subject). This topic investigates
what the conversations screen becomes when it leans into the record's
own idea of topics — legible at scale, with the machinery re-homed onto
an agent detail surface where a person can still talk to the agent — and
what, if anything, the record itself must grow to support it. The
operator's direction (2026-08-30): internal topics hidden from the
conversations list, available from the agent's detail.

## The question

Can the shell partition and order a large board into a legible topics
surface — human conversations leading, agent homes and system topics
re-homed onto an agent detail surface that still lets a person talk to
the agent — using only readings the record already serves, with any gap
named as a precise upstream ask rather than a shell-side store?

## Pre-registered bars

Written before any experiment runs. The standing house rules bind every
bar: no shell-owned store (0001 §4), no per-topic reads on the 1s tick
(0003 §5, §8), every act on the session's own admission (0001 §6).

- **Bar 1 — home-ness is derivable, not stored.** On a rig realm with
  ≥5 declared agents — at least one declaration re-pointed to a
  different home after first submit, and one whose home IS the
  placements topic (spec 011 allows it) — the shell classifies every
  board entry (human conversation / agent home / placements) correctly
  at every render, derived solely from the placements-topic reading and
  the configured placements name. Protocol: shell e2e against a real
  realm; pass = zero misclassifications across renders, zero new
  per-topic reads on the tick, zero shell-held state (restart the
  shell, the partition survives).
- **Bar 2 — the list orders by activity at scale.** A seeded realm with
  200 topics renders the conversations list ordered by last activity,
  the ordering read costing one bounded call per tick — never O(topics)
  OPS reads. Protocol: bench the tick against the seeded realm; pass =
  the board+ordering read completes in <100ms sustained at 200 topics.
  If the record cannot serve last-activity today, the bar is answered
  by a core spike measuring the same numbers (e.g. `BoardEntry`
  carrying last-activity) and the upstream ask is named with the spike's
  figures — a shell-side workaround that violates the tick rule is a
  FAIL, not a pass.
- **Bar 3 — talking on the detail actually wakes the agent.** A turn
  posted from the agent detail surface (carrying the agent's mention —
  mention is the default-on wake) wakes a declared agent exactly once,
  and the answer lands on the home topic visible on that same surface.
  Protocol: declared-agent rig; pass = one outcome (idempotent under
  replay, including across a wrap restart), rendered on the detail
  without a manual refresh. A no-dispatcher realm shows honest words,
  never a spinner (0009 §3's rule extends).
- **Bar 4 — no mention goes dark.** With home topics hidden from the
  conversations rail, a mention of the signed-in person landing on a
  hidden topic still surfaces within one tick, and the surfaced mark
  leads somewhere a click can follow. Protocol: e2e posting a mention
  into a hidden home; pass = the spine tally counts it and a visible
  path (agent row badge or equivalent) reaches the message.

## Reversal condition

Two observable readings would reverse the direction this topic assumes:

- **Hiding mislabels a lived surface**: if on the byon realm (or any
  realm with active people) human-authored turns recur in agent home
  topics beyond replies-to-the-agent — people deliberately using homes
  as ordinary conversations — then homes are conversations after all,
  and the direction reverses from hiding toward marking (a labeled
  group, not a removed row).
- **Derivation misses too much**: if realms accumulate undeclared
  machinery topics the placements reading cannot see (personal
  paste-block agents' homes, future system topics) to the point that
  the human list stays cluttered despite Bar 1 passing, then role
  derivation is insufficient and record-side vocabulary (tags or a
  kind) returns to the table as its own research topic.

## Verdict

<Empty until graduation. Filled by /research-graduate: PASS/FAIL per bar with the
honest numbers, each load-bearing claim tagged [measured] / [mechanism-argument]
/ [judgment].>
