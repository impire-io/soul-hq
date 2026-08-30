# the-topics-surface — investigation journey (opened 2026-08-30)

## 2026-08-30 — the mechanism experiments (Bars 1 and 2)

Rig: a scratchpad spike (embedded JetStream server, unsigned realm
client, `Provision`, then the public `topic`/`fleet`/`declaration`
packages only) against the sibling checkouts at main. The shell pins
core `v0.14.0-rc.1`; `topic.Board` is unchanged between that tag and
main at the measured paths. Loopback transport — every number below is
a floor; network RTT multiplies A1 hardest (see the round-trip count).

### Experiment A — what the board costs, and what the candidates cost

Four readings, 5 iterations each, med of min/med/max reported:

| topics × ops | A1 `topic.Board` (shipped) | A2 one-pass full replay | A3 last-per-subject over OPS | A4 `stream.Info` filtered |
|---|---|---|---|---|
| 50 × 20 | 27.2ms | 2.0ms (1.1k msgs) | 0.4ms | 0.1ms |
| 100 × 20 | 51.9ms | 3.9ms (2.2k msgs) | 0.6ms | 0.1ms |
| 200 × 20 | 104.8ms | 7.5ms (4.4k msgs) | 0.9ms | 0.2ms |
| 400 × 20 | 205.7ms | 13.8ms (8.8k msgs) | 1.7ms | 0.4ms |
| 200 × 5 | 94.0ms | 2.5ms | 0.7ms | 0.2ms |
| 200 × 40 | 117.1ms | 14.3ms | 1.3ms | 0.2ms |

**The load-bearing find [measured]:** the shipped board is linear in
topic count at ~0.52ms/topic on loopback (26→52→105→206ms across
50→100→200→400), and only weakly in ops-per-topic (94→117ms for 5→40
ops at 200 topics). The cost is round trips, not bytes: per topic the
board pays a `GetLastMsgForSubject` for the announcement, a second as
the ops-probe, and an ordered-consumer creation before the replay —
≥3 *sequential* round trips per topic, serial across topics (read off
`board.go`/`replay.go`). On a deployment with RTT r that adds ≥3·N·r:
at 1ms RTT and 200 topics, +600ms on every tick
[mechanism-argument from the measured call structure]. **At 200×20 the
shipped board alone already spends ~105ms on loopback — Bar 2's
<100ms threshold is crossed before ordering is even added, and the
shell's tick calls it every second** (and `placementsPath` calls it
again on the declare surfaces).

**The candidates [measured]:** one ordered-consumer pass over
`SOULSTREAM.TOPICS.>` replays the whole realm in 7.5ms at 200×20
(~1.6µs/msg, scaling with total ops, ~2 round trips total) — 14×
cheaper cold than the board it could replace, and it is the *cold*
cost of a watcher-fed projection whose warm tick is ~free. Activity
ordering alone is even cheaper: a `DeliverLastPerSubject` pass over
`SOULSTREAM.TOPICS.OPS.>` returns the last op per topic — each message
carrying its own timestamp — in 0.9ms at 200 topics: last-activity for
the whole board in one bounded call, no core change needed.
`stream.Info` (0.2ms) stands as the change-detector floor.

**The composition find [mechanism-argument]:** the placements topic's
declarations are ops in the same stream, so ONE watcher pass serves
the board, the lifecycles, the last-activity ordering, AND Bar 1's
partition — zero per-topic reads on the tick, no shell store (the
projection is memory, rebuilt from the log on connect, the same shape
as `Follow`). The upstream ask taking form: core grows a followed
board — one ordered consumer over `SOULSTREAM.TOPICS.>` serving live
`BoardEntry` + last-activity — and the shell consumes it.

### Experiment B — the partition (Bar 1)

Seeded on the same rig: a `placements` topic, six declarations through
`fleet.Submit` — agents 1–4 each with a started home, agent-1
submitted a second time pointed at a second home (the "re-point" case:
submission is additive, nothing un-places — there is no re-point
vocabulary, so BOTH homes are placements of record), agent-5 declared
with the placements topic itself as home (spec 011 allows it).

- **B1 [measured]:** the rule *machinery = the placements topic ∪
  every declared home* classified all 206 board entries correctly — 6
  machinery, 200 human, 0 misclassified, from one reading
  (`Materialise` + `fleet.DeclarationOf`), nothing else consulted.
- **B2 [measured]:** that reading costs 0.5ms med (0.5–0.7ms) — the
  added per-tick read if done naively; folded into the watcher
  projection it is zero additional reads.
- **B3 [measured]:** rebuilt from a fresh connection under a different
  persona: identical partition — a pure read, restart-proof by
  construction.

### Bars 3 and 4 — where they stand

Their mechanism halves already stand upstream: the wake engine's
exactly-once outcome for a mention is measured in workloads' own
integration suite (episode 0130 — one outcome slot, idempotent across
restart), and mention slips carry the topic path while hidden topics
stay ON the board (hidden from the rail is not off-board), so the
tray's standing check keeps them alive [mechanism-argument]. The
surface halves — the agent-detail composer that mentions, the badge
path that reaches a hidden home — are e2e criteria against the built
surface and stay open until the design builds.

Spike source: `topics-surface-spike/` in the session scratchpad
(main.go: A1–A4, B1–B3; flags -topics/-ops/-iters).

## 2026-08-30 — the design drafts ahead of graduation

At the operator's direction the findings composed into design
[`0012-the-topics-surface.md`](../../02-DESIGN/soulstream-shell/0012-the-topics-surface.md):
the followed board as upstream ask #4, the partition as a derived
role, the agent detail carrying the room, the four bars restated as
the build's acceptance criteria. The topic stays **active**: bars 3–4
close at build time and graduation composes the episode then.

Upstream ask #4 drafted the same day on the core side:
[`extensions/followed-board.md`](../../02-DESIGN/soulstream-core/extensions/followed-board.md)
— Layer 1 library growth by the presence pattern (a
library-and-adapters bullet grown into its own doc), the wire spec
untouched, the rig's numbers as its §1.
