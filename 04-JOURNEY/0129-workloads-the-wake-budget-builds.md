# Episode 0129 — The wake budget builds: the colony gate ships (2026-08-25)

The build, same day as its research (episode
[0128](0128-ecosystem-loop-safety.md)) and its design
([`0006-loop-safety.md`](../02-DESIGN/soulstream-workloads/0006-loop-safety.md)):
`specs/008-loop-safety` in soulstream-workloads, merged to main
(`af49b80`). The wrap package now evaluates the composed wake budget at
admission — after the self-skip and the outcome-existence pre-check,
before the outcome obligation attaches: the authorship-window floor (K
own turns per topic per W) and the provable-chain depth bound (D hops
over the `WakeOpID` binding), both computed from the topic view the wake
already reads. A refusal posts nothing and logs one `wake_refused` line
with the numbers; the ancestry walker (`ParentOf`, `ChainToRoot`,
`ProvableHops`, ambiguity reported never absorbed) ships in the pure
correlate half as the gate's engine and the operator's diagnostic. No
core change, no new dependency, no state beside the log.

The rig's discriminating cases came along as the feature's own suite,
all green under `-race` on the embedded server `[measured]`: the
uncooperative two-agent cycle halts at **exactly D=4** outcomes with
loud op-less refusals and no refusal testimony in the topic; the
id-evading self-post cycle (the MCP arm) halts **within 2K=6** on the
window floor; the human-rooted owner→A→B→A delegation completes with
**exactly 3 outcomes and zero refusals** under the shipped defaults; and
`Unbudgeted: true` reproduces the unbounded wrapper (past 12 turns in
0.58s) with the standing logged once per wrapper. The existing wake
suite passes byte-identically under a zero budget. Full gate green:
`make fmt && make test && make lint`, 0 issues.

Defaults are **on** — D=4, K=8, W=10m (design 0006 §3 `[judgment]`,
generous against every measured legitimate flow, orders of magnitude
under the measured danger). Opting out is the explicit `Unbudgeted`
config, never a zero-value reading. With this, design 0005 §7's
sequencing rule is satisfiable in code: a topic-wake colony deployment
carries this budget at its dispatcher's admission seam — the gate is
built where mention-wake lives today, and the seam is named for the
dispatcher when demand builds it.

Reversal condition: as episode 0128 and design 0006 — an outcome
invisible to both the depth walk and the authorship count evades the
floor and reopens the topic; an admission point that cannot materialise
the view it budgets on moves the mechanism or reopens the topic. For
the build itself: none beyond those — records a completed, measured
implementation.

Trail: `soulstream-workloads/specs/008-loop-safety/` (spec, plan,
research, data model, library contract, quickstart, tasks — all 16
done); branch `008-loop-safety` (`cb6e51c`, merged `af49b80`); design
[`0006-loop-safety.md`](../02-DESIGN/soulstream-workloads/0006-loop-safety.md);
research [episode 0128](0128-ecosystem-loop-safety.md); wrap context
[episode 0085](0085-workloads-wrap-run-your-agent-where-you-are.md).
