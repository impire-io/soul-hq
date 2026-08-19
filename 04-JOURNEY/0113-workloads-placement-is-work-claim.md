# Episode 0113 — Placement is work.claim: the fleet lands (2026-08-19)

Design [0003-fleet.md](../02-DESIGN/soulstream-workloads/0003-fleet.md)'s
central bet, built and measured a fortnight after its research gate
(episode 0033): **placement IS `work.claim`** — no auction, no
coordinator, no second control plane. soulstream-workloads **v0.6.0**.

A submission is an ordinary work item carrying the declaration; every
idle node races to claim it; the log decides. The load-bearing detail
the build made explicit: *publishing a claim never means winning it* —
a node re-reads the projection after claiming and launches only if the
projection names it owner, so the first claim in stream order wins and
every loser's claim stays visible as void. Reclaim is the researched
three-step, unchanged: the projection nominates a silent owner, a
**transient probe** on soulstream-workloads' own service space vetoes a
live one, and an ordinary `work.abandon` reopens the item.

Measured on two nodes against one realm [measured, §8's criteria]: four
contested placements, each run by **exactly one** node with exactly one
live claim apiece, reconstructible from replay alone; a live owner past
the reclaim bound never reclaimed (the probe's veto, at zero cost on
the record); a silent owner reclaimed and re-placed, its timeline
reading `claim,abandon,claim` with **no double close** and no
resurrected workload without a fresh claim; no probe traffic anywhere
on the stream.

The seam discipline held literally: the winning node calls
`Runner.Launch` exactly as a single node does, so the run's own
lifecycle stays the runner's work item and placement is the fleet's —
runner, declaration, and backend are untouched, and today's suites run
unchanged (§8.5, measured by the whole gate staying green). §8.4's
seedless scoped launch is the minter's existing measured property
(episode 0033's spike 4, the delegated-minting fallback), inherited
rather than re-measured here — named so the inheritance is visible.

What is *not* built, named: the declaration trigger vocabulary (a node
still places what a caller submits; research D2 left this waiting on
exactly this claim path, and it can now proceed), and a long-running
`fleet serve` loop — `TryPlace`/`Sweep` are the primitives, their
scheduling deliberately the caller's until a deployment shape demands
one.

Reversal condition: if contested placement ever produces two live
claims on one item (observable: the §8.1 assertion failing on any
deployment), placement-as-claim fails and the design's rejected
auction alternative reopens — the whole bet rests on the log being the
only arbiter.

Trail: design
[`0003-fleet.md`](../02-DESIGN/soulstream-workloads/0003-fleet.md) §8;
workloads `3f6793e` (v0.6.0, `fleet/` + `integration/fleet_test.go`);
episode [0033](0033-soulrealm-fleet.md) (the research gate this
finally builds).
