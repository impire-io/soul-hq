# Episode 0083 — The waker lands (2026-08-15)

M3.2 built, the same day its research graduated: `soulstream-workloads waker serve`
is the workload plane's trigger arm — one durable AckExplicit consumer per
registered agent on the notify stream, an admission probe per wake,
harness-agnostic invocation templates, and the runner-owned reply
obligation. The spec-kit flow ran whole (specs/005: spec → plan with
Constitution Check → research D1–D8 → tasks T001–T020 → build), the full
gate is green with nothing skipped, and `make test-wake` woke a **real**
`claude -p` through the built waker — mention → attributed reply in 6.6s
`[measured]`. The hermetic default suite proves the protocol with a
scripted harness speaking both measured grammars; ten packages, zero lint
issues.

Cross-repo prerequisite, landed first: **core v0.8.3** exports
`PostTurnIdempotent` — the retry-with-same-id duty the library always
documented, now a public arm riding rollup's existing preset-id machinery
(additive; two tests prove a same-id repost leaves one op). Workloads'
go.mod pins v0.8.3 and adds `soulstream-identity/client` (the ephemeral
lane's minter, behind a narrow interface wired only in `cmd/`) and
`google/uuid`. The graph-pruning gotcha is on the record in tasks T002:
workspace builds read the pinned version's go.mod, so the pin flipped only
in the landing commit `[measured]`.

The build's gate caught two design-level bugs the research rig never hit —
both now regression-guarded:

- **The self-wake loop** `[measured]`: the graduated design had the failure
  turn "mentioning the agent and the asker" — but tapping the agent
  notifies it, and a notify to a registered agent *is a wake*. The very
  first fault trial ran away: failure turn → notify → wake → failure turn,
  forever, each outcome op becoming the next wake's trigger. Corrected in
  design 0004 §7: failure **names** the agent, **taps only the asker**;
  and an agent's own mention of itself never wakes it. G7's loop-safety
  successor topic just acquired its second measured exhibit.
- **The outcome-id collision** `[measured]`: the wake op id hashed only the
  notify op — but one mention can tap several registered agents, and two
  agents' outcomes deduped into a single turn (the multi-agent gate test).
  A wake is one delivery *to one agent*: the id is now UUIDv5 of notify op
  **and persona** (design 0004 §6).

Shape notes for the record: the waker is a sibling of the runner (`waker/`
beside `runner/` — the launch arm and the trigger arm, both thin
orchestration over injected narrow interfaces with pure logic split out);
authorship is mechanical (two live realm clients — core stamps the author
from the client's persona, so cross-authorship has no code path); refused
wakes are op-less by design and loud in the log — `log/slog` enters the
repo as a stated precedent (research D8), the first logging in it. The
declaration is untouched: registrations are waker configuration until the
fleet's claim path gives trigger vocabulary a second consumer (research
D2, reversal observable recorded).

Reversal condition: none — records a completed build; the design's
conditions (episode 0082) move with design 0004.

Trail: design [`0004-the-waker.md`](../02-DESIGN/soulstream-workloads/0004-the-waker.md)
(amended §5/§6/§7 with the landed corrections);
[`specs/005-the-waker/`](../../soulstream-workloads/specs/005-the-waker/) frozen;
workloads `7688e3d`→`944f424` (merge), core `4d50a8e` + tag `v0.8.3`.
Push order: core (+ tag) before workloads, then `go mod tidy` completes
go.sum for module-mode CI.
