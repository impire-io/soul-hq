# Episode 0158 — The followed board: one consumer, the live board (2026-08-30)

The shell's topics-surface research (topic `the-topics-surface`, still
active) measured the shipped `topic.Board` linear in topic count —
~0.52ms/topic on loopback, ≥3 sequential round trips per topic, 105ms
at 200 topics × 20 ops, paid every second by the shell's tick
[measured] — and drafted upstream ask #4 the same day
([`extensions/followed-board.md`](../02-DESIGN/soulstream-core/extensions/followed-board.md)).
The ask built as core spec `022-followed-board` within hours:
`topic.FollowBoard`, ONE ordered consumer over `SOULSTREAM.TOPICS.>`
maintaining the whole board in memory — announcement, parentage,
lifecycle, and the new `BoardEntry.LastActivity` (stream timestamp of
the latest op) — with a snapshot read and a change signal, nothing
persisted, a new follower rebuilding from the log. Shipped numbers on
the research rig [measured]: cold build 6.5/7/8.9ms (min/med/max, 5
iters) at 200×20 against 103.4/106.8/109.8ms for one `Board` call on
the same run — 15×, and 14× inside the acceptance bound; warm
`Entries()` 0.2ms flat, zero round trips.

The lifecycle semantics extracted into one `lifecycleFold` driven by
both `apply` and the projection — the whole existing suite passed
untouched [measured]. One honest amendment surfaced during the build:
three op types (`edit`, `comment.resolve`, `attachment.remove`) count
as content only when their application succeeds, so the projection
carries a body-free id-index (contribution author/kind/resolved,
edit-chain aliases, attachment removal) — the design's "sized by topic
count" memory wording became "no op bodies; summary plus a bounded
id-index," recorded in spec 022 FR-3 and propagated to the design doc
in this change-set. The classifier mirror is pinned by adversarial
equivalence tests [measured]: void edits and resolves leave a dormant
topic dormant, a valid-shaped work ref to a ghost item wakes it,
rollup checkpoints fold as no-ops, baked baselines seed the index,
manifest baselines resolve lazily (one object-store fetch per
compacted topic), a live-observed archive folds transition-then-
checkpoint — every seeded lifecycle equal to `Board` field-for-field.
One test failure during the build was the seeding, not the code: a
sealed op posted onto the dormant fixture correctly woke it in both
folds — the fixture moved, the gate held.

What it opened: `Board` re-deriving from the same single pass (design
[O2] — a strict improvement measured at 15×), TypeScript parity by
demand ([O3]), and the shell's topics-surface build (design 0012),
which this ask unblocks — composed the same day as core
`v0.14.0-rc.2`, pinned into soulstream `v0.14.0-rc.9`.

Reversal condition: none — records a completed build/measurement (the
FR-3 wording amendment stands unless the id-index measurably grows
beyond bounds on a real realm, which would reopen the memory design).

Trail: core `specs/022-followed-board/spec.md`, branch
`022-followed-board` (`bfdf8c1`, merged `4a3860f`), tag
`v0.14.0-rc.2`; soulstream pin `126dc9c`, tag `v0.14.0-rc.9`; hq
design `02-DESIGN/soulstream-core/extensions/followed-board.md`,
research trail `01-RESEARCH/the-topics-surface/JOURNEY.md`.
