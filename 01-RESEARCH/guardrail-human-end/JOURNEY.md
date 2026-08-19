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
