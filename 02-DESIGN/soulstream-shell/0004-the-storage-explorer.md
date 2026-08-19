# 0004 — soulstream-shell: the storage explorer

**Status:** designed, not built — decided 2026-08-19 (episode
[0116](../../04-JOURNEY/0116-ecosystem-what-shipped-without-a-human-end.md)).
This document fills a hole in [0001 §3](0001-soulhelm-the-helm.md)'s
observe surface: the store is measured but never read. Pure
observation — no mutation class of [0001 §4](0001-soulhelm-the-helm.md)
applies, and none is claimed.

## §1 The gap

The Storage readout says how much the store holds and nothing about
what it holds: `%d ops · %s` plus a VU meter against the roof the store
declares for itself [measured, code trace: the shipped
`modules/overview/render.go` storage card]. That is a fuel gauge. When
something goes wrong — a turn that never appeared, a signature verdict
nobody expected, an agent writing a shape no reader folds — the person
looking at the screen has no way to see the op that caused it. Today
the answer is `nats stream view` at a terminal, with an operator's
credentials, outside the surface that reported the problem.

The record is built to be read this way. Headers carry the record and
the payload is pure data ([`01-protocol.md`](../soulstream-core/core/01-protocol.md)
§The operation record), the subject taxonomy is class-first so the
useful wildcards are cheap (§Subject taxonomy), and the protocol's
stated way to read is a consumer replaying subjects. An explorer is
that consumer with a screen on it.

**Zero upstream additions.** Both surfaces it needs are already public
on the pinned core: `realm.Client.JetStream()` hands the JetStream
context (`realm/connect.go:213`) and `record.Parse(headers, payload)`
turns a message back into a record (`record/record.go:131`), with
`Record.Canonical(realm, topic)` producing the bytes a signature covers
[measured, code trace]. No ask, no version bump.

## §2 The lane (normative)

**The explorer reads on the signed-in person's own admission, never on
the surface's shared read lane.** This amends [0001 §3](0001-soulhelm-the-helm.md)
for this surface only — the board and the topic view keep the backend
read lane they were built on.

The argument is custody, not scoping. The shared read lane exists so
the surface can render a realm's public shape before anybody is signed
in and without acting as anybody; a raw op-log browser is not that. It
is subject-level access to `SOULSTREAM.>`, which includes
`SOULSTREAM.PERSONA.NOTIFY.<persona-id>` — one person's notifications —
and, when sealed topics land, ciphertext whose readership the record
itself decides. A surface that reads those on a lane it was handed for
a different purpose is lending authority, which is exactly what
[0001 §6](0001-soulhelm-the-helm.md)'s "delegated authority, never
borrowed identity" refuses. The session's client already reads the
stream in the shipped build (design [0003 §2](0003-conversation-lifecycle.md)'s
close materialises on it), so the lane costs nothing to use.

**It must be said plainly that this narrows nothing today.** The
product's persona scope grants every admitted persona `SOULSTREAM.>`
and `$JS.API.>` [measured, code trace: `soulstream/ceremony/ceremony.go`,
`scopePubAllow`], so every signed-in person can already read the whole
op-log — which is a fact about the realm, not something the explorer
introduces. The screen therefore says *what your admission can read*
and never implies per-person scoping it does not have. The payoff of
choosing the session lane is in the future tense and worth the choice
anyway: the day a deployment narrows that scope, or sealed topics land,
the explorer follows with no surface change.

**Reversal condition:** if an operator cannot debug a real failure
because their own admission cannot read what broke — observable: a
recurring class of failure whose evidence sits on subjects the
operator's persona is refused — the answer is a deliberately named
operator lane as a numbered decision here, never a quiet fallback to
the shared read lane.

## §3 The decided surface

**Getting there, two ways, no imports.** The module takes a key on the
spine at the foot beside System status, and the Storage card on Home
and on System status offers a way in — asked for through the frame
(`shell/link.go`: identity + screen kind + subject), so the overview
module keeps importing nothing (design [0002 §2](0002-the-module-shape.md)).

**The list.** Ops newest-first, one row each: the stream sequence, the
stream's own timestamp (not the author's claim — the record says
ordering authority is never a clock), the subject, the type, the author
as a display name, and the signature verdict. Filtering is by **NATS
subject pattern**, validated as one and refused as one when malformed —
not a search box (§4). The taxonomy's classes get one-key filters
because they are the questions people actually ask: this conversation's
ops, the whole topic board, one person's notifications, the service
lane.

**One op.** Everything the message is: sequence and stream timestamp,
every `Soulstream-*` header verbatim, the payload, the canonical bytes
the signature covers, and the verdict with its reason. The point of the
screen is that nothing is summarised away — this is where a person goes
when the summary is what they distrust.

**Following.** A live tail with a stated cap and a pause, on the SSE
lane the shell already runs. `SOULSTREAM.>` on a busy realm is a
firehose; the cap is on screen, not implicit.

**Empty and refused states** speak in the plain register: a filter that
matches nothing says so; an admission that is refused the read says
that, and says whose refusal it is.

## §4 What it refuses

- **No store.** Nothing read here outlives the request. The custody
  rule ([0001 §6](0001-soulhelm-the-helm.md)) is unchanged and the
  standing scan covers it.
- **No query layer.** The protocol names its own absence
  ([`01-protocol.md`](../soulstream-core/core/01-protocol.md) §What the
  wire does *not* provide): projections are built by consumers
  replaying subjects. The explorer is one such consumer and must not
  grow a persistent index. Full-text search across the op-log is
  therefore **not offered** — wanting one is a consumer-side index and
  its own decision, not a stretch of this one.
- **No decoding it cannot justify.** A payload renders as text when it
  is valid UTF-8 under a cap; otherwise the screen says its size and
  its type and shows nothing. Object-store attachments are named, not
  fetched and inlined.
- **No acts.** The record is append-only; there is no delete. A surface
  offering one would be lying about the protocol, so the explorer
  offers nothing but reading.
- **No borrowed lane** (§2).

## §5 Rendering notes

- The list is stream-owned (the SSE tick morphs it); the filter form
  and the one-op panel are act-owned siblings the tick never writes —
  [0001 §5](0001-soulhelm-the-helm.md)'s lesson, restated by design
  [0003 §5](0003-conversation-lifecycle.md).
- Signature verdicts use the topic view's existing vocabulary and
  keyring construction. There is no second verdict language on the
  screens.
- The tail's pause survives morphs by the bundle's attribute-preserve
  mark, the mechanism design 0003 §5 already proved on the archived
  fold.

## §6 Acceptance criteria

1. **The lane holds, mechanically.** Every read in the module rides
   the session's client; a standing test asserts the module never
   reaches the support layer's shared read lane, and the custody scan
   stays clean with its positive control firing.
2. **The screen's scope claim is true.** A test pins that no copy
   claims per-person scoping, and the build records once, in §7, what
   the deployment's persona scope actually permits.
3. **Verdicts are earned**, by the same keyring path as the topic view;
   `unknown-key` shows honestly rather than as a failure.
4. **A hostile payload cannot break the page**: a table test over
   non-UTF-8, oversized, and markup-shaped payloads.
5. **Zero upstream additions** — the build uses only `JetStream()`,
   `record.Parse`, and `Record.Canonical` from the pinned core (§1).
6. **Plain register** ([0001 §7](0001-soulhelm-the-helm.md)): the
   screens say storage, message, conversation — never a vocabulary
   byname.
7. **No horizontal overflow at 1000 px and 390 px**, the standing guard
   from episode [0080](../../04-JOURNEY/0080-shell-one-instrument-any-width.md).

## §7 Open questions [O]

- **[O1] The tail's cap** is a number nobody has measured. The first
  build picks one, states it on screen, and records what a real realm's
  rate actually was.
- **[O2] Sealed topics** ([`sealed-topics.md`](../soulstream-core/extensions/sealed-topics.md),
  designed and chafe-gated) turn some payloads into ciphertext. When
  they land the explorer shows that plainly — a sealed op is a readable
  record with an unreadable payload, not an error. No work now.
- **[O3] The narrow-scope question** is soulstream's, not the shell's:
  whether every persona keeps `SOULSTREAM.>` + `$JS.API.>` is the
  product's ceremony decision. §2's lane choice means the explorer
  follows either answer without a surface change.
- **[O4] Admin-only?** Deliberately not gated. Authority comes from the
  transport, and a UI gate over a lane that already permits the read
  would be decoration. If the product narrows the scope (O3), the gate
  arrives for free and in the right place.
