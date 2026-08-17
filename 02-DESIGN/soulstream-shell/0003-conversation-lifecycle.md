# 0003 — soulstream-shell: conversation lifecycle

**Status:** built — decided and landed 2026-08-17 (episode
[0102](../../04-JOURNEY/0102-shell-conversations-begin-and-end.md)).
This document fills in the participation scope 0001 §1 gained from
episode [0071](../../04-JOURNEY/0071-ecosystem-the-focus.md): starting,
closing and archiving conversations from the shell. Every act here is
class (a) of [0001 §4](0001-soulhelm-the-helm.md) — the enumeration is
unchanged; this is a feature inside it, not a design change.

## §1 The gap

The shell reads conversations and posts into them, but cannot start
one and cannot end one. Every conversation on its screens was seeded by
another surface — the MCP door, the CLI, a test rig. A fresh realm
greets its first person with "No conversations yet." and no way
forward; a finished conversation sits in the list forever. The record
has carried all three acts since before the shell existed
(`topic.StartTopic`, `Handle.Close`, `Handle.Archive` — present in the
core version the shell already pins, so no upstream ask and no version
bump) [measured: verified against the pinned tag].

## §2 Mechanism — three acts, one lane

All three acts ride the signed-in person's own realm client (0001 §6:
delegated authority, never borrowed identity), signed by their persona
key, attributed on the record. The shell's shared read lane never
mutates; no new privileged side-channel, no shell-owned store.

- **Start** takes a name (required) and an optional line about the
  subject; nothing else at birth — no tags, no parent, no first
  message. The conversation is born `proposed` and turns `active` on
  the first message [mechanism-argument: that is the record's own
  lifecycle fold].
- **Close** materialises first — the archived guard reads the handle's
  last-observed state, so an unmaterialised close could hit an archived
  conversation [measured: core reads `h.lifecycle`, set only by
  observation] — then posts the transition.
- **Archive** is the record's loud, final act: transition, final
  compaction, withdrawn-attachment sweep; afterwards writes are refused
  outright. Nothing archives on a person's behalf (GENESIS: never a
  silent janitor).

## §3 The decided surface

- **Starting** lives in both places a person would look: a fold at the
  head of the conversation list, and a card on Home in the shape the
  admin screen already uses. The empty states ("No conversations
  yet.") point at it. Success lands the person in the new
  conversation; failure speaks beside the form.
- **Ending is a ladder**: a live conversation offers *Close*; only a
  closed one offers *Archive*, behind a two-step confirm that says
  what archive means — kept for reading, closed to writing, no way
  back. The ladder mirrors the record's own story (closing is social,
  archiving is final) and keeps the terminal act out of casual reach
  [judgment]. No reopen is offered anywhere: the protocol has no
  reopen — closed is one-way [mechanism-argument: the lifecycle fold
  ignores transitions to `active`].
- **The list**: live and closed conversations show as today, each with
  its state in plain words; archived ones move under a collapsed
  "Archived" fold at the foot — still readable, one click away, never
  deleted. The default conversation a person lands on skips archived.
- **The composer**: an archived conversation shows a quiet note in the
  composer's place — the record would refuse the write, so the surface
  does not offer it. A closed conversation keeps its composer: closing
  is a social convention the record warns about but accepts, and the
  state stands plainly in the details panel [mechanism-argument].

## §4 Truthful copy under partial failure (normative)

The record's lifecycle calls can half-succeed, and the surface must
not lie in either direction:

1. Close can return an error *with the close standing* (the tidy-up
   compaction behind it failed). The copy says closed, and says the
   tidy-up did not finish — never "failed to close".
2. Archiving an already-archived conversation reports *already done*,
   never an error.
3. Archive can land the transition and lose the final compaction race;
   the copy says archived and says how to finish (archive again).
4. Posting into a conversation someone else archived mid-session is
   answered with the archived note, not a raw error.

## §5 Rendering notes (bench-verified, binding on the build)

- The shell's SSE lane carries element patches only; an act whose
  outcome is *navigation* (starting a conversation) answers with a
  script response, which the vendored bundle executes [measured: read
  off the bundle]. Everything else stays element patches.
- 0001 §5's lesson holds: mutation results get their own targets. The
  lifecycle acts render inside the stream-owned details panel but
  answer into an act-owned sibling element the stream never writes.
- The archived fold lives inside the stream-owned list and keeps its
  open/closed state across morphs via the bundle's attribute-preserve
  mark [measured: read off the bundle]; the server serves it open only
  when the open conversation is itself archived.
- The board tick gains no per-topic reads: the list's partition uses
  the lifecycle the board already carries.

## §6 Acceptance criteria

1. The standing e2e gate walks the whole ladder against a real realm:
   start (lands on the record, named, attributed to the session's
   principal) → close (state closed; a further post still lands —
   closing is social) → archive confirm (the "no way back" step) →
   archive (state archived; a further post is refused; the composer
   yields; the list folds it away; the default landing skips it) →
   archive again (answered "already", not "error").
2. Every act rides the session principal's own client — the class-(a)
   row of 0001 §4's standing e2e stays green, and the custody scan
   (0001 §8.2) stays clean.
3. On-screen copy stays in the plain register (0001 §7): people read
   "conversation", never a vocabulary byname.
4. §4's four truthfulness branches each have a test pinning their
   copy.

## §7 Open questions [O]

- **[O1]** Closed conversations keep their composer. Reverses if
  people in practice keep writing into closed conversations expecting
  them to wake — evidence would be closed-topic posts recurring in
  realms with active people.
- **[O2]** The archived fold is a foot-of-list answer. A realm with
  hundreds of archived conversations may want a screen of its own;
  that is a new decision, not a stretch of this one.

## §8 As built (2026-08-17, episode 0102)

Landed as designed; the deltas worth recording:

- Whether a conversation is archived is read once at the page serve
  (one read-lane materialise), never on the tick — the board tick
  gained no per-topic reads, as §5 requires.
- The details panel's standing test invariant "no buttons" was relaxed
  to "no unresolved anchors; every button is an act" to seat the
  lifecycle acts — the spirit (nothing pretends to be a way somewhere)
  survives, per-state honesty moved into its own test.
- The e2e gate walks the whole §6.1 ladder against a real realm; the
  two morph claims of §5 (script execution, the fold's preserved
  toggle) were proven in a live browser over `make screens`.
