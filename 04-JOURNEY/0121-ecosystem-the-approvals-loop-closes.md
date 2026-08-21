# Episode 0121 — The approvals loop closes: tickets, presentation, policy (2026-08-21)

Episode 0119's design became identity v0.11.0 the same day, and the
house turned the guardrail on. The loop that research measured as
unclosable — a correctly minted, verifiable approval leaving the retry
still deferring, because nothing public could present it — closes in
consumer position, every arm witnessed.

**The ticket (D42).** A deferral stops being a stateless refusal: it
opens a durable, sealed, TTL-bounded ticket
(`pending → approved → spent`, with `denied` and `expired` as recorded
outcomes — a pending ticket past its window is written expired the
moment any read observes it, never silently dropped). The refusal names
the rule, the invocation, and the window; `ParseDeferral` mirrors it;
no call ever waits for a human — the retry is the execution. The two
clocks stay deliberately distinct: the ticket TTL is the human's window
(1h default, configurable), the approval TTL (minutes, in-memory, D38
unchanged) is the retry's and starts at the yes. A fresh ask on a
terminal ticket reopens a fresh window; a live one never resets.

**The public ends (D43).** `MintApproval` (an approval IS a D33
delegation naming `invocation:<id>` — no second artifact kind),
`PresentApproval` — one method over the op that always existed —
`DenyApproval` (the no this build added: the design's own table named
the denied state with no way to reach it), `ApprovalStatus` (the
originator's own tickets; anyone else's answers not-found
indistinguishably), `PendingApprovals` and `GuardrailRules` (a surface
shows the loop and the standing rules, keeping no copy),
`GuardrailLoad`. The guardrail exemption went prefix-wide: a rule
deferring `approvals.status` would deadlock the loop that resolves
deferrals.

**Policy (D45).** Rules gain an `approvers` clause, consulted against
the *current* set at presentation — policy governs at the moment of the
answer, not the ask. A well-formed yes from outside it refuses by name.
Persona names only for now: group clauses need claims the presentation
lane does not carry — a named follow-up.

**Measured, consumer position** [6/6 clean plus full-gate runs]: defer
(parsed) → pending visible → yes minted and presented → retry served →
spent witnessed; the deny arm with a late yes refused by state; the
rogue approver refused by name; **expiry witnessed** with a
yes-after-the-clock refused; and Bar 3's at-rest arm runnable at last —
the ticket store carries the invocation's name and **never its
arguments**, sealed, plant control fired, the audit silent. The whole
gate: 3.4s.

**The house (soulstream v0.13.0-rc.9)**: `EnableGuardrail` on with an
empty rule set — everything allows, rules load live via
`guardrail.load`, tickets ride beside — so the loop has a live home
rather than a knob nobody turned. The node's own gates ran green over
it, which is the empty set's do-no-harm measured in passing.

What remains open, named in the design: the shell's approvals screen —
where a human actually sees "X wants to do Y" and their tap becomes
mint-plus-deliver — and with it the argument-visibility question
(today an approver sees principal, action, rule, window — never
arguments, the ticket's own privacy line). The reversal condition
0119 set stays live and now measurable: if real tickets recurringly
expire before any human sees them despite witnessed, notified expiry,
the answer is pre-authorization policy, not a longer clock.

Reversal condition: the design's own, carried — plus the build's: if
the prefix-wide guardrail exemption proves too coarse (observable: a
deployment needing to gate `guardrail.load` itself by rule rather than
by template), the exemption narrows to the named deadlock set as a new
D-decision.

Trail: design
[`approvals.md`](../02-DESIGN/soulstream-identity/approvals.md)
(as-built §); identity `4d93cf1` (v0.11.0), soulstream `c347847`
(v0.13.0-rc.9, the house's enable); episodes
[0119](0119-ecosystem-guardrail-human-end.md) (the graduation this
builds), [0110](0110-identity-the-tenancy-set-builds.md) (the headless
build this gives ends to).
