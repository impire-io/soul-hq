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

Graduated to design 2026-09-02 (designs
[`soulstream-shell/0012`](../../02-DESIGN/soulstream-shell/0012-the-topics-surface.md)
and [`soulstream-core/extensions/followed-board.md`](../../02-DESIGN/soulstream-core/extensions/followed-board.md),
both built: core `v0.14.0-rc.2`, shell `v0.11.0-rc.8`).

- **Bar 1 — home-ness is derivable, not stored: PASS.** The rule
  *machinery = placements topic ∪ every declared home* classified all
  206 rig entries with zero misses, additive re-point and
  home-on-the-placements included, from one reading, 0.5ms med,
  rebuilt identical from a fresh connection [measured]. The shell e2e
  walks the same two edge cases against a real realm: the rail and the
  Home list settle to the person's conversations alone, zero
  misclassified rows, zero per-tick reads (the partition rides a
  placements-topic follower — memory, rebuilt on connect, never a tick
  read) [measured].
- **Bar 2 — the list orders by activity at scale: PASS, by the
  registered escape.** The shipped `Board` measured over the threshold
  on loopback alone (~105ms at 200×20, ~0.52ms/topic, ≥3 sequential
  round trips per topic [measured]) — so the bar's own answer path
  fired: the upstream ask was named with the spike's numbers rather
  than a shell-side workaround. Core built the ask the same day
  (spec 022): cold build 6.5/7/8.9ms at 200×20 against
  103.4/106.8/109.8ms for one `Board` call on the same run, warm
  `Entries()` 0.2ms, zero round trips [measured]. The shell renders
  the list from the projection with zero JetStream reads on the tick;
  the open thread's per-tick materialise is pre-existing surface
  behavior outside this bar's board+ordering scope, named as a
  follow-topic watcher opportunity [judgment].
- **Bar 3 — talking on the detail wakes the agent: PASS, composed of
  two measured halves.** The shell e2e proves its half: the detail
  composer's turn lands on the home topic authored by the signed-in
  person and carrying the agent's mention — the wake's own trigger —
  and renders on the detail live, no reload; the no-dispatcher realm
  answers in honest words [measured]. The engine half — one outcome
  per mention, idempotent across restart — is episode 0130's measured
  ground [measured, workloads rig]. The composed loop on a serving
  realm is byon's dogfood to witness once a harness serves there
  [mechanism-argument until then].
- **Bar 4 — no mention goes dark: PASS.** A second persona's mention
  into a hidden home reaches the spine tally within one tick, the rail
  grows a pointer whose click lands on `/agents/room?who=…`, the
  agent's row carries the mark, and the room shows the message
  [measured, shell e2e].

Reversal condition stands as registered: byon watches whether people
deliberately converse in agent homes (hiding would reverse toward
marking), and whether undeclared machinery accumulates past the
partition's reach (record-side vocabulary would return).
