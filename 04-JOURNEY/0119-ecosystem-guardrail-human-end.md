# Episode 0119 — The guardrail's human end: the loop measured to where it stops, then designed shut (2026-08-19 → 2026-08-21)

The question, opened by episode 0116's evaluation beside its sibling
0118: how does a deferred guardrail decision reach a human, get
answered, and release the refused invocation — without a second control
plane, without the shell keeping rule or decision state, and without
the approval artifact carrying the arguments it binds. Five
pre-registered bars; graduated to design with the honest verdict the
topic existed to earn: **three bars FAIL by construction with their
gaps made mechanical** — a topic opened to find missing ends, finding
exactly them.

**The central finding, measured** (rig in the session scratchpad,
consumer position, identity v0.9.0 by tag, deliberately no core
requirement; 6/6 runs): D38's loop is missing **two** ends, not one.
The emit half works today — a real op tripping a defer refuses with the
rule and a deterministic, argument-sensitive invocation id, everything
the approving side needs to bind to. The mint half is public — an
approval is exactly a D33 delegation naming `invocation:<id>`, one
`MintDelegation` away. And then nothing: the public client (27
methods) has **no `approvals.present`**, the sealed envelope is
internal, and a correctly minted, verifiable approval leaves the retry
still deferring. Bar 1 FAIL by construction, proven mechanically rather
than argued. Bars 4 and 5 likewise: no read op for rules or decisions
exists (the dispatch serves exactly two ops), and any
directory-resolvable persona's signature approves, with no policy to
refuse one.

**Bar 3 — PASS on the artifact as minted** [measured, 6/6]: the
deferred op's arguments carried a planted secret; the minted artifact
carries the invocation hash and the parties and nothing of the
arguments; positive control fired. The wire and at-rest arms wait for
a presentable loop — nothing can go on a wire yet — and become the
build's gate.

**Bar 2 — resolved, and the carrier question dissolved** [measured +
mechanism-argument]: zero soulstream-core edges in identity's entire
module graph, and the plane can neither write the record (the cycle
guard) nor push on its own prefix (a persona may subscribe only to
`_INBOX.>` and `SOULSTREAM.>` — nobody may listen where the plane may
publish). So the plane sends nothing: it keeps readable state, and
**the adapter that originated the call carries the news** — the door
to its agent, the shell to its human. Nothing ever crosses the line,
and Bar 2's re-measurement after any build stays trivial.

Refuted or corrected along the way: the assumption that the record's
`PERSONA.NOTIFY` was the natural carrier (it would breach the wall,
and turns out to be unnecessary); and a **doc/code drift in D38
itself** — the design said the invocation hash includes time; the
implementation rightly excludes it (a retry could never match a
time-salted name), measured via same-id-on-retry, and the design text
corrected at graduation.

**The operator directed the lifecycle** (2026-08-21), and the design
carries it whole: tickets become durable state with a TTL of their own
(pending → approved → spent, with denied and **expired as a
first-class, recorded, notified outcome**); two clocks kept distinct
(the ticket's human-window vs the approval's minutes-long retry-window
starting at the yes); **async by construction** — no call ever waits
for a human, the refusal is structured, the retry is the execution;
and notification named as correctness, not courtesy — "approved, retry
now" must land inside the one-shot's window.

What it opened: the build items in the graduated design
[`soulstream-identity/approvals.md`](../02-DESIGN/soulstream-identity/approvals.md)
— D42 the ticket, D43 the loop's public ends (`PresentApproval`,
`approvals.status|pending`, `guardrail.list`), D44 the notification
decomposition, D45 per-rule approver policy — plus the shell's
approvals screen as a module design at build time, all behind the 0071
focus gate. The rig is the draft of the standing gate the build
inherits.

Reversal condition: the design's own, carried from the topic and
softened by visible expiry — if real tickets recurringly expire before
any human sees them despite recorded, notified expiry (observable: the
expired/answered ratio staying high on a real deployment), the answer
is pre-authorization policy with the one-shot ticket as the exception,
not a longer clock and not a better screen.

Trail: design
[`approvals.md`](../02-DESIGN/soulstream-identity/approvals.md)
(D42–D45, and the D38 textual correction); the concluded topic
(removed on graduation, full history in git — opened `7fe0011`, the
rig and the two-missing-ends finding `8017189`, the operator's ticket
lifecycle `b06b152`, verdicts `4435093`); episodes
[0116](0116-ecosystem-what-shipped-without-a-human-end.md) (the ask),
[0110](0110-identity-the-tenancy-set-builds.md) (the headless build
this gives ends to),
[0107](0107-ecosystem-platform-tenancy-guardrails.md) (D37/D38's
design), [0027](0027-soulstream-dx-hardening-and-the-cycle-guard.md)
(the guard the carrier answer honors), and
[0118](0118-ecosystem-agent-external-tools.md) (the sibling whose door
is this loop's other adapter).
