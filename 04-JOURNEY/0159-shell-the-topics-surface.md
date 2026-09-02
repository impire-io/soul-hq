# Episode 0159 — The topics surface: the machinery leaves the rail (2026-08-30 → 2026-09-02)

Research topic `the-topics-surface` opened on the operator's direction
— the conversations list unmanageable at scale, an agent's home not a
conversation — asked one question: can the shell partition and order a
large board into a legible topics surface, the machinery re-homed onto
an agent detail a person can still talk through, from readings the
record already serves? Four pre-registered bars; three days from
question to shipped surface, with one core feature grown on the way.

**The verdicts** (full numbers in the topic README, git history):

- **Bar 1, the partition — PASS** [measured]: *machinery = placements
  topic ∪ every declared home*, a role read from the declarations at
  every ask, never stored. 206/206 rig entries classified with the
  additive re-point (submission is additive; "re-point" is not
  vocabulary — both homes stay machinery) and the
  home-on-the-placements case; restart-identical; e2e'd against a real
  realm with zero misclassified rows.
- **Bar 2, activity order at scale — PASS by the registered escape**
  [measured]: the shipped `Board` failed on loopback alone (~105ms at
  200×20, ≥3 sequential round trips per topic), so the bar's own
  answer path fired — upstream ask #4, built the same day as core
  spec 022 (episode [0158](0158-core-the-followed-board.md)): 7ms
  cold, 0.2ms warm, zero round trips. The tick renders from memory.
- **Bar 3, the talk wakes — PASS, composed of two measured halves**:
  the e2e proves the shell's half (the detail composer's turn lands
  authored-by-the-person and carrying the agent's mention — the wake's
  own trigger — rendering live, the no-dispatcher realm answered in
  words) [measured]; the engine's exactly-once half is episode 0130's
  ground [measured]. The composed loop on a serving realm is byon's to
  witness once a harness serves there [mechanism-argument until then].
- **Bar 4, no mention goes dark — PASS** [measured]: a mention into a
  hidden room reaches the spine tally within a tick, the rail grows a
  pointer whose click lands on the agent's detail, the row carries the
  mark, the room shows the message.

**What shipped** (design
[`0012`](../02-DESIGN/soulstream-shell/0012-the-topics-surface.md),
shell `v0.11.0-rc.8`, composed as soulstream `v0.14.0-rc.10`): the
support layer's watch — core's followed board plus one placements
follow, both memory rebuilt on connect, zero JetStream reads on the
render tick; the rail and Home list machinery-free and newest-activity
first with the archived fold standing; deep-opened rooms honest about
whose they are; the agent detail at `/agents/room` — identity,
presence, the declaration's facts with their delivery classes, and the
room's thread with the mentioning composer.

**As-built deltas, ledgered in 0012 §8**: the partition rides a
*second* follower rather than §2's "one pass" — the board projection
deliberately retains no op bodies (core FR-3) and a declaration IS a
work-item body [mechanism-argument]; the thread rendering moved to the
support layer (`soulstream/threadview.go`, the `UnreadMark` precedent)
so both screens share one message rendering and no module imports
another — the purity gates passed untouched; the rail grew the
hidden-rooms pointer line 0012 §4 implied but never drew.

**What it opened**: the open thread's per-tick materialise (a
follow-topic watcher is the natural next mechanism); the visual pass
on the new markup (canon classes, no bespoke CSS — the operator's
call, calm-pass style); 0012 §8's named opens (sub-topics, tags, the
undeclared-machinery watch, the thousands-of-topics screen).

Reversal condition: as registered on the topic — if byon shows people
deliberately conversing in agent homes, hiding reverses toward
marking; if undeclared machinery accumulates past the partition's
reach, record-side vocabulary returns as its own topic.

Trail: research `01-RESEARCH/the-topics-surface/` (removed at
graduation; full history in git), designs
[`soulstream-shell/0012`](../02-DESIGN/soulstream-shell/0012-the-topics-surface.md)
and
[`soulstream-core/extensions/followed-board.md`](../02-DESIGN/soulstream-core/extensions/followed-board.md);
shell branch `0012-topics-surface` (`acceead`, merged `12af8db`), tag
`v0.11.0-rc.8`; soulstream pin `eac5d2d`, tag `v0.14.0-rc.10`;
episodes [0158](0158-core-the-followed-board.md) (the core half).
