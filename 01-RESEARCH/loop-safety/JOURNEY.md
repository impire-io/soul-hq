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
