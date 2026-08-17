# Episode 0102 — Conversations begin and end in the shell (2026-08-17)

**The gap, named by the operator:** the shell could read conversations
and post into them, but not start one and not end one — every
conversation on its screens was seeded by another surface (MCP, CLI, a
test rig), a fresh realm dead-ended at "No conversations yet.", and a
finished conversation sat in the rail forever. The record had carried
all three acts since before the shell existed — `StartTopic`,
`Handle.Close`, `Handle.Archive` in the core version the shell already
pins, so the whole build needed **zero upstream additions and no
version bump** [measured: the change-set touches no go.mod].

**Designed first, built the same day.** Design
[`0003-conversation-lifecycle.md`](../02-DESIGN/soulstream-shell/0003-conversation-lifecycle.md)
recorded the operator's calls: starting lives in *both* places a
person looks (a fold at the head of the rail, a card on Home — one act
behind both); ending is a **ladder** — a live conversation offers
Close beside its status, only a closed one offers Archive, behind a
two-step ask that says what archive means; archived conversations rest
under a collapsed fold at the foot of the list; the composer yields
only where the record would refuse the write (archived), and stays on
closed — closing is social, the record warns and accepts
[mechanism-argument: core's own lifecycle contract]. No reopen is
offered anywhere, because the protocol has none — the fold ignores a
transition back to active [mechanism-argument, read off core's apply].

**The build** (soulstream-shell, 17 files, +908/−51): every act rides
the session's own admitted connection — class (a) of design 0001 §4,
unchanged. Two rendering mechanisms were bench-verified before use
[measured: read off the vendored bundle, then proven in a live
browser]: an act whose outcome is *navigation* answers as a
`text/javascript` response the bundle executes (`shell.Script` —
starting lands the person in the new conversation), and the archived
fold inside the stream-owned rail keeps its toggle across the
one-second morphs via `data-preserve-attr="open"` — hand-opened, it
survived 4+ ticks in the `make screens` browser pass; served open only
when the person is looking at an archived conversation. The lifecycle
acts answer into `#convo-life`, a dock of their own beside the
stream-owned details panel — 0001 §5's mutation-targets lesson,
applied again.

**Truthful copy under half-success, pinned by tests:** the record can
close a conversation and still hand back an error (the tidy-up behind
the close is best-effort) — the words say closed and never "failed";
archiving twice answers "Already archived", not an error; a lost final
compaction says archived *and* how to finish; a post into a
conversation someone archived mid-session is refused in the composer's
own words. All four branches are table-tested [measured], and the
whole ladder — start → announce on the record → first message wakes it
→ close → a further post still lands → ask → archive → post refused,
composer yielded, rail folded, default landing skips it → archive
again answers "already" — is walked by the standing e2e gate against a
real realm (suite 16–20 s on the dev machine) [measured].

**One invariant deliberately relaxed:** the details panel's "no
buttons" test assertion. Its spirit — nothing on the panel pretends to
be a way somewhere — survives as "no unresolved anchors; every button
is an act"; the lifecycle acts are the panel's first controls. Recorded
here because a written invariant changed shape, not silently.

Reversal condition: episode 0071's stands for the participation scope
(the custody gate under posting). Specific to this build: **closed
conversations keep their composer** ([O1] in design 0003) — reverses
if closed-topic posts recur in realms with active people, read from
the record.

Trail: design
[`0003-conversation-lifecycle.md`](../02-DESIGN/soulstream-shell/0003-conversation-lifecycle.md)
(hq `9b3aef4`); soulstream-shell `1ad66ec` (feature, tests, gate);
screenshots verified live over `make screens` (start fold, the ladder
in the details panel, the life dock, the archived fold, the Home
card).
