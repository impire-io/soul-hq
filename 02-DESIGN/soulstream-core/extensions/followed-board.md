# Extension: The Followed Board

**Status:** built — drafted 2026-08-30 as **upstream ask #4** from
shell design
[`soulstream-shell/0012`](../../soulstream-shell/0012-the-topics-surface.md)
(research topic `the-topics-surface`, measured on its rig the same
day) and **built the same day** as core spec `022-followed-board`
(episode
[0158](../../../04-JOURNEY/0158-core-the-followed-board.md); tag
`v0.14.0-rc.2`, pinned into soulstream `v0.14.0-rc.9`) — see §6. The
wire spec is untouched: no new subject, no new op type, no
new stream — this is Layer 1 library growth
([`library-and-adapters.md`](library-and-adapters.md)), and
[`core/03-topics.md`](../core/03-topics.md) already names the local
projection as discovery's first layer. Any bare NATS client could
build the same thing from the core docs alone; the library grows it
once so every watcher stops rebuilding it badly.

## §1 The gap [measured]

`topic.Board` is a one-shot that pays per topic: a
`GetLastMsgForSubject` for the announcement, a second as the ops
probe, an ordered-consumer creation, then that topic's full replay —
≥3 *sequential* round trips per topic, serial across topics. Measured
linear at ~0.52ms/topic on loopback (26→52→105→206ms across
50→100→200→400 topics at 20 ops each); with network RTT r it adds
≥3·N·r. A caller that renders every second — the shell's tick — pays
the whole price every second, and still learns nothing about
*when* anything last happened: the board carries no activity, though
every op already carries its timestamp.

One ordered-consumer pass over `SOULSTREAM.TOPICS.>` replays the same
realm whole in 7.5ms (4.4k messages, ~1.6µs/msg) [measured] — and
that is the *cold* cost of a projection whose warm updates are
single-message folds.

## §2 The capability

The library grows a **followed board**: `topic.FollowBoard` (name
final at spec time), one ordered ephemeral consumer over
`SOULSTREAM.TOPICS.>` — INFO and OPS in one pass — maintaining a live
in-memory projection and serving:

- **Per topic:** path, latest `Announcement`, `Parent`/`ParentKnown`,
  `Lifecycle`, and **`LastActivity`** — the timestamp of the topic's
  latest op.
- **A snapshot read** (the board, ordered stably) and **a change
  signal** (the `Follow` shape: a callback per applied update, so a
  renderer knows when to re-render instead of polling).
- **One fold, one definition:** lifecycle derives through the same
  fold `Materialise` applies — the projection must never grow a second
  `apply`.
- **Sealed honesty unchanged:** a sealed announcement renders as
  `Board` renders it (name empty, sealed flag standing) unless the
  caller supplies the unwrapper, exactly as today.
- **Nothing persists.** The projection is memory; a reconnect rebuilds
  from the log (proven restart-identical on the rig [measured]). The
  ordered consumer's own sequence-checked recreation is the gap guard.
- **Memory holds no op bodies** (amended at build, spec 022 FR-3): the
  projection keeps the per-topic summary — announcements, lifecycle,
  one timestamp — plus a bounded, body-free id-index (contribution
  author/kind/resolved, edit-chain aliases, attachment removal) that
  the three application-gated content types require; bodies are
  dropped as they fold. The original "sized by topic count" wording
  proved unachievable exactly: `edit`, `comment.resolve`, and
  `attachment.remove` count as content only when they apply.

`Board` stays for one-shot callers. Whether it re-derives internally
from the same single-pass fold — retiring its per-topic loop, a
strict improvement — is its own spec decision (§5).

## §3 Who it serves

- **The shell** (design 0012): the conversations list ordered by last
  activity and partitioned, with zero JetStream reads on the render
  tick — the ask's origin.
- **Discover responders**: layer 2 of discovery says any persona
  answering `topic.discover` "maintains a projection" — this is that
  projection, given one good implementation.
- **The curator**: sweeps the board on a cadence; the same projection
  serves it without the per-sweep rebuild.

## §4 Acceptance criteria

1. **Scale:** at 200 topics × 20 ops, cold build under 100ms (the rig
   floor is 7.5ms); a snapshot read costs zero round trips; an op
   landing on the stream reflects in the projection on its own
   delivery, unpolled.
2. **Equivalence:** on the same realm, every projected
   announcement/lifecycle/parent equals `Board`'s answer, and
   `LastActivity` equals the topic's latest op timestamp — a golden
   test over a seeded realm, sealed topics included.
3. **Restart:** tear the consumer down, rebuild, byte-identical
   projection.
4. **Memory:** a realm with 10k ops holds a projection sized by its
   topic count, not its op count.
5. **One fold:** the lifecycle fold has exactly one definition shared
   with `Materialise`; the spec test that seeds every lifecycle
   transition passes through both paths with one answer.

## §5 Open questions [O]

- **[O1]** Manifest baselines: a rollup-compacted topic's baseline
  lives behind an object-store pointer, so a cold build may pay one
  fetch per *compacted* topic (bounded by compactions, not ops).
  Measure when rollup-heavy realms exist; lazy resolution is the
  expected shape.
- **[O2]** `Board` re-derived from the single-pass fold, retiring its
  per-topic loop — a strict improvement to the one-shot too (7.5ms
  against 105ms on the same rig), decided at spec time.
- **[O3]** TypeScript parity: Layer 1 names two targets; the TS
  projection follows by demand, against the shared spec tests.

## §6 As built (2026-08-30, episode 0158)

Landed as designed, with the deltas worth recording:

- **The numbers** [measured, research rig, 5 iters]: cold build
  6.5/7/8.9ms (min/med/max) at 200 topics × 20 ops — the same run
  where one `Board` call cost 103.4/106.8/109.8ms; warm `Entries()`
  0.2ms flat, zero round trips. §4.1's 100ms bound met 14× over.
- **§2's memory bullet amended** (above): no op bodies, plus the
  bounded id-index the application-gated content types require.
- **One fold held whole**: `lifecycleFold` extracted and driven by
  both `apply` and the projection; the existing suite passed
  untouched. The content classifier's projection-side mirror is
  pinned by adversarial equivalence tests — void edits/resolves leave
  a dormant topic dormant, a valid-shaped work ref to a ghost item
  wakes it, rollup checkpoints fold as no-ops, a live-observed
  archive folds transition-then-checkpoint.
- **[O1] exercised at build**: a >128KB rollup produced a manifest
  baseline; the projection resolved it lazily and read equal to
  `Board`.
