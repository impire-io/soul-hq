# Episode 0124 — The first hour and the presence lease (2026-08-23 → 2026-08-24)

The operator named the next gap in the 0116 pattern, one level up: not
a feature without a human end but **the product without a first hour**
— people install soulstream, sign in, and stand in an empty realm not
knowing what to do next. The survey found the furnishing acts all
standing and unsequenced: the agents screen already hands one
paste-able block that runs unchanged under POSIX shells and fish
[measured, code trace: `modules/agents/render.go`, `renderRunIt`],
Home already carries disconnected hints ("Set one up", the start
card), and the wrap announces nothing on start [measured, code trace:
`cmd/soulstream/wrap.go` — no profile op], so the realm's first
evidence of a wrapped agent is its first answer. Two designs were
drafted across the two days and decided together on 2026-08-24: shell
[`0008-the-first-hour.md`](../02-DESIGN/soulstream-shell/0008-the-first-hour.md)
and core extension
[`presence.md`](../02-DESIGN/soulstream-core/extensions/presence.md).

The decisions. **0008: guidance is a reading, never a store** — the
first-steps card on Home derives entirely from realm state per render
(the catalog's display-grade demotion applied to guidance), with
three refusals standing on house rules: no onboarding store (0001
§4), no per-person tour (D26 — a second person joining a furnished
realm sees the furnished realm), no wizard. Beside it, the **arrival
principle** (an act ends when the realm's own evidence shows the
thing arrived, live, on the acting screen) and the empty-state rule
(an empty screen a furnishing act would fill offers that act).
**presence.md: the who-is-around face** — a `soulstream-presence` KV,
one key per persona, additive value (`status`, `since`, optional
`doing`, kind-shaped extras); a lease renewed on a cadence (30s/90s
until the dogfood says otherwise), a farewell written on clean stop,
and staleness read as gone: **departure is derived, never merely
announced** [mechanism-argument: a crashed writer cannot say goodbye
— workloads 0003 §3's undecidable-stop posture, applied]. The store
forgets nothing (no TTL-delete — expiry destroys the evidence that
distinguishes *left* from *vanished* and takes "last seen" with it);
each thing writes its own key on its own admission, no collector;
the wrap must, services should, workload tools and adapters may.
**Upstream ask #3** (continuing shell 0001's numbering): the wrap
publishes its profile on start and holds the lease.

Refuted and refined. Lifecycle *events* as the mechanism were argued
and refused — announcements are manners, silence is the truth-teller
[mechanism-argument]. The operator's teach-back ("a heartbeat
protocol we capture and put in a KV with TTL") surfaced two
corrections absorbed into the doc before deciding: no capturer (a
collector would be a new privileged side-channel), and no TTL (the
reader judges staleness; the store never deletes). One inherited rule
was sharpened in the writing: the thin convention's
attention-aware-agents example collides with a flat "no machine reads
it", so advisory now reads **courtesy, never correctness** — a field
rides if a wrong or stale value costs nothing but grace [judgment,
recorded in the doc]. The payload line keeps the box closed: words a
person reads, never numbers a machine steers by — capacity stays the
fleet's, history stays refused as the query layer the protocol
declines.

What it opened: the fresh-eyes install (0008 [O1] — the step ordering
is asserted, not measured; one install by someone who is not the
operator, chafe log open from minute one); upstream ask #3's build in
soulstream; the cadence numbers against NGS budgets (presence [O1]);
humans' sessions and write scoping later ([O2]/[O3]). A standing face
also answers shell 0005 [O2] — running-state on workload tool rows —
without waiting for the record's declaration vocabulary.

Reversal condition: if a fresh-eyes first hour still ends lost with
the card standing (observable: the minute-one chafe log), stored or
stateful guidance is reconsidered as a new decision, not a quiet
patch; and if any consumer demonstrably needs to *depend* on presence
for correctness (observable: a design asking to gate an act, route,
or schedule on the face), that need gets its own mechanism — the face
is never quietly promoted to authority.

Trail: designs
[`02-DESIGN/soulstream-shell/0008-the-first-hour.md`](../02-DESIGN/soulstream-shell/0008-the-first-hour.md)
and
[`02-DESIGN/soulstream-core/extensions/presence.md`](../02-DESIGN/soulstream-core/extensions/presence.md)
(grown from the thin paragraph in
[`extensions/library-and-adapters.md`](../02-DESIGN/soulstream-core/extensions/library-and-adapters.md),
which now points at it); both design READMEs; this commit.
