# Episode 0117 — The store shows what it holds (2026-08-19)

Design [0004](../02-DESIGN/soulstream-shell/0004-the-storage-explorer.md)
built the same day it was written, as its own module. The shell's Storage
readout said `412 ops · 3.1 MB` and nothing else; from today it also says
which ones. Both stores under a plain name and the name the server answers
to, a page of messages newest first over a validated subject pattern, and
one message whole — every `Soulstream-` header verbatim, the payload as it
is, and the canonical form beside it, which is what a signature is actually
over and the thing people get wrong about this record. Verdicts run through
core's own `VerifyRecord` against a keyring built from the persona
directory: no second verdict vocabulary on the screens.

**Zero upstream additions** [measured, code trace]: `realm.Client.JetStream()`,
`record.Parse` and `Record.Canonical` are already public on the pinned core
(v0.11.1). The one addition anywhere is in the shell's own support layer —
`KeyringFor(personas…)`, for a reader holding ops rather than a materialised
conversation.

**The lane decision held, and got a real subject.** The explorer reads on
the signed-in person's own admission rather than the surface's shared read
lane. The design's first draft argued that from `SOULSTREAM.>` containing
one person's notifications — and that subject space does not exist: the
op-log captures `SOULSTREAM.TOPICS.>`, the persona inboxes are their own
stream, and the service lane is captured by nothing at all [measured, code
trace: `realm/spec.go`]. Corrected in the design, the argument came out
stronger: the inbox store is per-person by construction, so the lane is
doing work today rather than only the day a deployment narrows its grant.
What the screen still says out loud is that it narrows nothing *here* —
the product grants every admitted persona the whole subject space, and a
comfortable lie about that would be the worst possible copy on the one
screen people reach when they have stopped believing what they were told.
A test pins the sentence and refuses four ways of implying otherwise.

**What it refuses, mechanically:** no act, no delete, no persistent index,
and no search — the query layer is what the protocol deliberately does not
have, and a debugging screen is the obvious place to smuggle one in. The
screen says so in words and a test fails the build if an act appears on it.

Three things the build decided that the design had left open or wrong. The
walk is **backwards by sequence with a stated cap** — the client publishes
no reverse-filtered read, and a compacting store leaves its live messages
near the tail, so a page of messages is usually a page of sequences; a read
that stops at the cap says how far it looked, because a silent stop reads as
"that is everything". **Following is a mode, not a toggle** (`?follow=1`
renders the tail's SSE init and the key is a plain link), so pausing is a
navigation: no script, survives reload, bookmarkable. And **subject matching
is written in the module** and table-tested, because the server's own filter
reads forward from a sequence while this screen reads backward from the
newest.

Measured on the standing gate against a real realm: the work item written
moments earlier is in the list with its verdict, a subject pattern narrows
it truthfully, a pattern that is not one is refused where it was typed, one
message opens whole, the overview's readout offers the way in through the
frame (`shell.Link`, no import), and the tail refuses an unauthenticated
caller. Measured live in a browser: **zero page overflow at 1000 px and
390 px** with the table scrolling inside its own wrapper, and a message
written while the tail was running arrived at the top within one tick,
verified [measured].

Reversal condition: design 0004 §2's stands — an operator who cannot debug a
real failure because their own admission cannot read what broke (observable:
a recurring failure class whose evidence sits on subjects the operator's
persona is refused) brings back a deliberately named operator lane as a
numbered decision, never a quiet fallback to the shared read lane. New from
the build: if the sequence walk's cap proves to be the common case rather
than the edge on a real deployment (observable: the "stopped after examining"
note appearing on ordinary reads), the walk needs a server-side reverse read
and that is an upstream ask, not a bigger number.

Trail: design
[`0004-the-storage-explorer.md`](../02-DESIGN/soulstream-shell/0004-the-storage-explorer.md)
(§8 as-built); shell `0daa95f` on main, unreleased — tagging is the
operator's; episode [0116](0116-ecosystem-what-shipped-without-a-human-end.md)
(the evaluation that asked for it), [0078](0078-shell-the-module-contract.md)
(the contract it plugs into), [0080](0080-shell-one-instrument-any-width.md)
(the width guard it was measured against).
