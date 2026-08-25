# loop-safety — investigation journey

Opened 2026-08-25. The investigation appends here as it happens.

## 2026-08-25 — the substrate, read before any run

What the record carries today, from the shipped code (no experiment yet):

- **The wake→trigger binding exists.** A `mention.notify` carries the
  mentioning op's id, its author, and the topic
  (`soulstream-core/topic/mention.go:57`, `notify.go:24`). The outcome op
  publishes under `WakeOpID = UUIDv5(mentionOpID + "/" + persona)`
  (`soulstream-workloads/wrap/correlate.go:22`) — one-way, but
  *verifiable*: given a candidate trigger and the outcome's author, the
  binding checks. Record-only ancestry is a scan, not an inversion.
- **Exhibit 1's mechanism confirmed in code.** `PostTurnIdempotent` runs
  `mergeMentions` over the reply body (`topic/post.go:76`): an agent reply
  that *says* `@asker` fires a fresh notify even though the wrapper passes
  `mentions=nil`. The cascade edge is body text, not an API choice.
- **The shipped guards, located.** Self-skip (`wrap/wake.go:42`), failure
  taps only the asker (`wake.go:113`), outcome-existence pre-check
  (`wake.go:54`). All local, as the pre-registration states.
- **A pre-named Bar 1 hazard.** The `correlated_self_post` path
  (`wake.go:87`): an agent that posts during its run via its own client
  (the MCP arm) lands an outcome under an *arbitrary* id — no UUIDv5
  binding, so the record-only chain roots there. Registered before the
  run: Bar 1 is expected to hold for wrapper-posted outcomes and break at
  self-posts; whether a depth budget is thereby evadable is a Bar 2
  sub-experiment.
- **The admission point exists.** `handleWake` reads the topic *before*
  invoking — a gate can sit on that read. And a gate placed *in the
  harness slot* instead would convert refusals into self-reports that tap
  the asker — the 0083 runaway shape reborn through the budget itself.
  Pre-registered as a discriminating sub-experiment (gate placement).
- **Window budgets are record-computable.** `Contribution` carries
  `Timestamp`, `Mentions`, `Author` (`topic/view.go:40`) — a "my own
  outcomes in this topic, recently" budget needs no state beside the log,
  and authorship is mechanical (the client's persona), so it cannot be
  evaded by body text.

Mechanism candidates going into the rig, to be discriminated, not
assumed: (a) provable-chain depth via the UUIDv5 walk; (b) local
authorship-window budget (K own outcomes per topic per window W). The
rig: script-harness `Invoker` closures over the real `wrap.Wrapper`
(the injection seam is shipped — `wrap.go:32`), embedded JetStream
(natstest's shape), scratchpad module `scratch/loop-safety-rig`.

## 2026-08-25 — the runs: all four bars measured

Rig: embedded JetStream, real `realm.Client`s, real `wrap.Wrapper` for
cascade generation, a gated wake loop (identical shape ± the admission
gate) for enforcement runs, a topic-wake colony loop per design 0005 §3
for Bar 4. Every run `-race`; full suite `-race -count=1` green in
45.9s. Raw numbers as logged:

- **Bar 1 PASS [measured].** Two real wrappers, always-mention scripts,
  three human roots posted back-to-back: **421 agent turns in 5.0s
  (84.2 wakes/s, the unenforced danger number)** — and the record-only
  walker resolved **421/421** to a human root, **0 ambiguous**, with the
  three concurrent cascades cleanly separated (exactly 3 turns per
  depth, depths 1–141). The 0082 correlation lesson holds: id-binding,
  never stream order.
- **Bar 1 hazard confirmed, and it is total [measured].** The same
  cascade with self-posted outcomes (the `correlated_self_post` / MCP
  arm): **404 agent turns, 0/404 resolvable**. Record-only ancestry is
  blind to every outcome that does not ride the deterministic id.
- **Gate placement [measured].** A "budget" living in the harness slot
  (the invoker refusing, `Retries=1`): the wrapper's own contract turns
  every refusal into a self-report tapping the asker — **312 agent
  turns in 4s, 156 of them self-reports**, the 0083 failure ping-pong
  reborn *through* the budget. The gate must sit at admission, before
  the outcome obligation attaches; a refusal must be op-less.
- **Bar 2 PASS [measured].** `depthGate(D=4)` at admission, both agents
  maximally uncooperative, wrapper-posted outcomes: **exactly 4 agent
  outcomes** (the pre-computed bound), **1 loud op-less refusal**,
  settled, depths {1,2,3,4}, all chains resolved. Zero cooperation
  needed.
- **Bar 2 evasion [measured].** `depthGate(4)` with self-posted
  outcomes: provable hops stay 0, the gate never fires — **393 turns in
  3s, 0 refusals**. Depth alone is not enforcement.
- **Bar 2 window PASS [measured].** `windowGate(K=3, 1min)` against the
  same evading cascade: **halts at exactly 6 = 2K**, 1 refusal,
  settled. Authorship is mechanical (the client's persona) — body text
  and arbitrary op ids cannot forge it.
- **Bar 3 PASS [measured].** Composed gate (depth 5 + window 5/min),
  human-rooted delegation owner→A→B→A: **exactly 3 outcomes, 0
  refusals**, depths {1,2,3}, all resolved to root. The budget is not
  "agents never wake agents".
- **Bar 4 PASS [measured].** Colony of three topic-wake agents (author
  filter per design 0005 §3). Unenforced: **434 agent ops in 0.34s —
  1,264.7 ops/s from ONE human op**, generation growth 3, 6, 12, 24,
  48, 96, 192 (doubling per generation; the exponential is real, not
  hypothesized). Enforced `windowGate(K=3, 1min)`: **exactly 9 = N·K
  ops**, 12 loud refusals at settle; a fresh human probe op drew +3
  refusals and **zero** new ops — parked and loud, never wedged, never
  silent.

What the numbers decide (each tagged):

1. **The mechanism is a composed budget** — the authorship-window as
   the enforcement floor (unforgeable, halts every cascade shape
   measured), the provable-chain depth as the tight bound and the
   operator's diagnostic on chains that ride the deterministic id
   [measured: each alone fails a case the other covers].
2. **The gate sits at wake admission, refusals op-less and loud** —
   never in the harness slot [measured: the 312-turn failure
   ping-pong], never as a posted op [mechanism-argument: an op is a
   wake source; 0083].
3. **Budget exhaustion is a delay, not a loss** [mechanism-argument]:
   nothing is acked away — the mention stays in the inbox window, the
   outcome id stays deterministic, a slid window re-admits; the
   record is the position, unchanged.
4. **The self-post hole is structural and the design must not lean on
   provability** [measured: 404/404 and 393-past-the-gate]. Named
   improvement, not a prerequisite: if the MCP arm ever posts outcomes
   under the wake's deterministic id, depth becomes reliable there too.
5. Reversal condition standing: neither observable fired — ancestry
   reconstructs from the record alone where the id rides (Bar 1), and
   no halt depended on harness behavior (Bars 2/4 used maximally
   uncooperative and id-evading scripts).
