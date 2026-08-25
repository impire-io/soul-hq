# Episode 0127 — The first hour builds: the card derives, the roster breathes (2026-08-25)

Design 0008 built, two days after it was decided and one day after its
presence prerequisite (episode
[0125](0125-ecosystem-the-presence-lease-builds.md)). Ships as
**soulstream-shell v0.11.0-rc.3**, composed into soulstream main the
same day.

**The first-steps card** leads Home while the house is young, and §2's
"guidance is a reading, never a store" survived contact with the build
as a *testable property*: the derivation is a pure function
(`deriveSteps(stepFacts)`), so no-store is a unit test rather than a
promise, and the e2e walks the whole lifecycle over real screens — a
fresh house offers four pending steps, each furnishing act (agent,
conversation, tool, person) flips exactly its own step on the next
render, and when everything is done the card is gone with nothing
anywhere to dismiss [measured, the walk passes in 1.9s]. Steps are
offered on evidence, never a guess: an unreadable roster contributes
nothing, and a non-administrator is never shown an act that is not
theirs ([O3] held: a step nobody on screen can take is a wall, not a
door).

**The Around column** reads the presence face through the support
layer — in / left {when} / seen {when} in the person's words, an
honest dash for a voice the realm has never seen, exact moments and
the writer's own line riding hovers, and a 5-second live channel (the
approvals tick's shape) keeping the judgment fresh — so an agent shows
*in* while the person still holds its paste block, which now names the
next step in words: start a conversation and mention it, the answer is
the proof. The support layer's standing last-seen refusal
(`soulstream/agents.go`) is rewritten honestly rather than quietly
contradicted: still no store of this layer's own — the realm grew a
face to *read*. The **empty-state sweep** closed the three bare
sentences the survey found (tools, people, apps): each now says what
the thing is and offers the act, per-audience where the act is not
everyone's.

Two build notes worth the record. The per-module banned-word gates did
their work *before* a line was written: the column's vocabulary was
chosen as the person's words from the start because `persona`,
`realm`, and their kin are refused by standing tests — the gates as
design pressure, not just as police [judgment]. And the e2e's shared
rig seeds two conversations, which would have satisfied the talk step
before the walk began — the first-hour gate runs on an unseeded rig,
which is exactly what "fresh realm" must mean to a test.

What remains is unchanged from 0124/0125: the byon soak carries both
halves live (the lamp from a real wrap, the card on a real Home — the
quickstart's pending human act), and the fresh-eyes install ([O1])
stays the one test nobody has run: the step *ordering* is still
asserted, not measured.

Reversal condition: none — records a completed build/measurement; the
direction decisions and their reversal readings are episode 0124's.

Trail: soulstream-shell branch `first-hour` commit `0a2211d`, merge
`c033c25`, tag **v0.11.0-rc.3**; soulstream pin bump `68acda6` (after
`011-presence-lease` merged as `b86ec57`); design
[`0008-the-first-hour.md`](../02-DESIGN/soulstream-shell/0008-the-first-hour.md).
