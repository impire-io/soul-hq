# JOURNEY — guardrail-human-end (opened 2026-08-19)

The investigation, as it happens. Appended to, never rewritten.

## 2026-08-19 — opened

Opened from the operator's evaluation of the running system: "and what
about the boundaries/guardrails? I would expect something for that as
well, no?"

The survey that produced the bars, before any experiment:

- The evaluator is real and complete at the custodian chokepoint
  (identity v0.5.0, episode 0110): CEL, 10k cost limit, interrupt every
  100 steps, 25 ms deadline, first match decides, an erroring rule
  fails closed, an invalid rule refuses the whole load while the
  running set stays whole.
- It is off unless asked for: `-guardrail` defaults false, and a
  non-empty starting rule set enables it implicitly
  (`cmd/soulstream-identity/main.go`, `embed/embed.go`). The product
  sets neither.
- The op dispatch serves exactly two guardrail-adjacent ops —
  `guardrail.load` and `approvals.present` — and the client package
  exposes neither. There is **no read op** for the standing rules and
  no feed of decisions; the audit exists only as log lines.
- D37 named three chokepoints and only the first is built: the
  custodian's op path, then "the waker's wake decision and the door's
  tool-call forwarding". The third is the seam this topic shares with
  [`agent-external-tools`](../agent-external-tools/README.md) — the
  catalog says which tools exist, the guardrail says which calls are
  allowed. The two topics stay separate so neither holds the other
  hostage; the seam is named here so graduation does not forget it.
- Bar 2 exists because the obvious carrier is the wrong one to assume:
  routing a defer through the record would make the identity plane
  depend on soulstream-core, which is exactly what the cycle guard
  (episode 0027) and 0107's Bar 5 keep at zero.

## 2026-08-21 — the rig: the loop measured to the exact line where it stops

Rig in the session scratchpad (`guardrailrig/`, consumer position:
identity v0.9.0 by tag, zero replaces — and deliberately **no
soulstream-core requirement**, because the rig lives on the identity
side of Bar 2's line). The embedgate ceremony, the plane with
`EnableGuardrail` and one data rule (`action == "secrets.put"` →
defer), a caller on the ops lane, an approver admitted through callout.
6/6 runs green.

**What measured, and how far** [all measured]:

- **The emit half works today.** A real op tripping the defer refuses
  with the rule name and the invocation id in the refusal string — the
  caller holds everything the approving side needs to bind to, and the
  deferral is observable in the audit. D38's "the guardrail's defer
  emits the request" is true at the caller already; no plane-side
  carrier exists or is needed for this half.
- **The id is bindable.** The same invocation retried yields the same
  id; different arguments yield a different id. Which surfaced a
  **doc/code drift**: D38's sketch says the invocation is named by
  `hash(principal, action, args, time)` — the implementation hashes
  principal, action, and the raw argument bytes and **excludes time**,
  and the code is right: a retry at a later moment could never match a
  time-salted hash. The design doc's sketch needs the word removed.
- **The mint half is public.** The approval is exactly a D33 delegation
  naming `invocation:<id>`, minted through `MintDelegation` — the
  approver's one `sign.record` round trip, persona key materializing on
  first touch. No new vocabulary was needed, which is B10 paying out.
- **Bar 3 — PASS on the artifact as minted** (6/6): the deferred op's
  arguments carried a planted secret; the minted artifact carries the
  invocation hash and the parties and **nothing of the arguments**,
  positive control fired. Honest scope: the wire and at-rest arms of
  Bar 3's protocol cannot run until an artifact can actually be
  presented — nothing to put on a wire yet.
- **Bar 1 — FAIL by construction, both halves measured to the gap.**
  The public client (27 methods) has **no `approvals.present`**, and
  the sealed envelope is internal — so a correctly minted, correctly
  bound, verifiable approval **cannot be presented by any consumer**.
  The rig proves it mechanically: after minting, the retry still
  defers. D38's loop is missing not one end but **two** — the human
  surface the topic opened for, and the client method underneath it.
- **Bar 2 — the cycle-guard half measured**: `go mod graph` over the
  rig shows **zero soulstream-core edges** in identity v0.9.0's entire
  transitive graph. And the carrier question dissolves the way the
  external-tools door question did: the defer already reaches *the
  caller*; the person-facing carrier belongs to **composition** — a
  surface that already holds both a session and the identity client
  (the shell) reads pending defers from the plane and puts them in
  front of a person. The plane publishes nothing to the record; the
  candidate design is a read op, not a carrier. Bar 2's re-measurement
  after any build stays trivial because nothing crosses the line.

**Bars 4 and 5 — FAIL by construction, gaps named** [code trace,
re-confirmed in consumer position]: the op dispatch serves exactly
`guardrail.load` and `approvals.present`; there is no read op for the
standing rules and no feed of decisions (the audit is log lines), so
nothing can show a person the rules without keeping a second copy —
and any directory-resolvable persona's signature approves, with no
policy layer to refuse one, and no way to demonstrate refusal
end-to-end until presentation exists.

**What the topic now points to, small and named**: (1) a
`PresentApproval` client method — the loop's missing link, one method
over an op that already exists; (2) a pending-defers/standing-rules
read surface (Bar 4's op); (3) the shell module that puts a defer in
front of a person and turns their yes into the mint+present pair;
(4) per-rule approver policy (Bar 5, already 0110's named follow-up).
The reversal condition stands untested: whether real deferrals survive
their minutes-long TTL until a human reaches a browser is measurable
only with the surface built.

## 2026-08-21 — the operator directs the ticket lifecycle

The operator, on the ELI5: tickets need a TTL of their own; a
synchronous call must be handleable asynchronously (a human takes an
arbitrary time, possibly longer than any TTL); an expired ticket must
notify the originator; and a *deferred* ticket must notify the
originator that a human is now required. Folded into the design the
topic graduates to:

- **The ticket becomes stateful.** Today the defer is stateless — a
  refusal string, no pending record anywhere. The operator's TTL
  requirement confirms the pending store must exist (it was already
  half of Bar 4's read-op gap). States: pending → approved → spent,
  with pending → **expired** as a first-class, recorded, notified
  outcome — never a silent death. Denied joins as the human's no.
- **Two clocks, deliberately distinct.** The ticket TTL is the human's
  window (long); the approval TTL is the retry's window (minutes,
  one-shot, already built, and its clock starts at the yes — so human
  latency never eats the retry's window). Nothing changes in the
  built half.
- **Async by construction, as an invariant**: no call ever waits for a
  human. The op returns refused-with-ticket immediately; the retry is
  the execution. The refusal becomes structured (ticket id, TTL,
  "a human has been asked") rather than prose, so the originating side
  can act on it programmatically — a door answers its agent's tool
  call promptly as pending, never a hung MCP call.
- **Notification decomposes under the wall, twice over**: the plane
  can neither write the record (cycle guard) nor push on its own
  prefix (a persona's sub permissions are `_INBOX.>` and
  `SOULSTREAM.>` — nobody may listen where the plane may publish)
  [measured: the template in both rigs]. So the plane owns the truth —
  the ticket store, every transition, a `status` read op on the
  principal's own tail (the same template addition `grants.>` was) —
  and **the adapter that originated the call carries the news**: the
  door tells its agent, the shell shows its human. Composition carries
  in both directions; the plane stays core-free and push-free.
- **Notification is correctness, not courtesy** [mechanism-argument]:
  the approval TTL starts at the yes, so an originator polling slowly
  would always miss its window. "Approved — retry now" reaching the
  originator promptly is what makes the one-shot approval usable, and
  belongs in the design's acceptance criteria.
- **Open, named for the design doc**: restart semantics. Approvals are
  in-memory and fail closed on restart — defensible at minutes-scale.
  Tickets at human-scale TTLs are harder: a restart would expire every
  pending ask unwitnessed. Durable (the plane has sealed stores) vs
  in-memory-with-honest-loss is a D-decision to take at graduation.

The reversal condition softens accordingly: expiry being visible,
recorded, and notified makes longer ticket TTLs tolerable, so
"deferrals die before a human sees them" now argues for tuning the
ticket clock and pre-authorization policy — not against the surface.
