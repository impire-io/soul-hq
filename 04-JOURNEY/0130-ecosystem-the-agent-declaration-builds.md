# Episode 0130 — The agent declaration builds: the record declares, the room runs (2026-08-25)

The build of design
[`0005-agent-declaration.md`](../02-DESIGN/soulstream-workloads/0005-agent-declaration.md),
two days after its research graduated (episode
[0126](0126-ecosystem-agent-declaration.md)), in two halves landed in
dependency order. **Core grew the one additive stream** (feature
`020-system-stream`, merged `f0a09f2`, unreleased — tagging stays a
human act): `SOULSTREAM_SYSTEM` provisioned create-or-report beside its
siblings with `AllowMsgSchedules` + `AllowMsgTTL`, the
`SCHEDULES.<persona>.<name>` / `TICKS.<persona>.<name>` subject helpers,
conformance and report coverage. The substrate holds the design's bet
`[measured]`: a registration carrying `Nats-Schedule: @every 1s` made
the embedded nats-server 2.14.3 append a tick on the TICKS subject
~1.3s later **with no consumer running**, and a `Nats-TTL: 1s` message
was readable immediately and gone within ~1s — real expiry, not
acceptance-only.

**Workloads then built the declaration and the engine** (feature
`009-agent-declaration`, merged `09c446f`, whole gate green including
`-race`). The declaration grows `instructions` (a stage-1 artefact
reference, materialised at the lineage **tip** on every wake,
digest-checked, scratch-only — a revision through ordinary ops reached
the next wake with no redeploy and no durable copy `[measured]`),
`capabilities` (names-not-grants, schema and strict validation this
slice), `wake` (all four kinds with their delivery classes normative),
the 0006 `budget` block, and the record-form artifact
`soulstream://<topic-path>/<artefact-name>` (resolved to the tip,
digest-verified into run scratch; a tampered digest refuses and the
work item abandons `[measured]`). Wrap's engine generalized to every
kind through the unchanged five-step admission order — self-skip,
outcome-existence, budget, invoke, discharge — with trigger identities
through the frozen `WakeOpID` namespace: mention = notify op id
(existing outcomes unbroken), topic = the triggering op id, schedule =
the tick's stream sequence, subject = SHA-256 of the payload.
`soulstream-wrap --declaration <file>` is the host; without it the
wrapper is byte-for-byte yesterday's.

The design's acceptance bars, measured on the embedded server: four
kinds fired from one declaration and every stream-backed wake answered
**exactly once across an engine restart** (stream count 1 and attempt
count 1 per trigger), the subject kind's loss-when-down asserted as the
documented behavior it is; an op authored by the declared persona never
woke it; an uncooperative topic-wake cycle halted at the window bound
**2K = 8** with op-less loud refusals — design 0006 §6's requirement is
satisfied by construction, the budget sits in the dispatcher's
admission path, so topic wakes ship *with* their colony gate; schedule
re-publish replaced without doubling the rate, purge deregistered, and
`Nats-Schedule-TTL` bounded the backlog live `[measured]`.

**The enforcement-read gap `[O]` is RESOLVED: runtime-side reads**
`[judgment, recorded]`. The engine performs every record-position read
(catch-up, materialisation, outcome-existence, budget computation,
ticks, instructions) on its own connection; the declared agent's minted
scope stays exactly the shipped `$JS.API.INFO`. Two reasons, both from
the field: the JS read tails are stream-wide, so widening the agent
template would breach own-prefix confinement; and scoped templates are
written at account founding, so widening is a per-deployment
control-plane migration — the byon rc.10 deployment duty, already paid
once, not to be minted again. Capabilities resolution (the D28
`mint.ephemeral` tag lane) is the named follow-on `capability-minting`;
this repo deliberately keeps no identity-plane dependency until that
demand.

Build findings propagated into design 0005: topic-wake `types` are
engine-restricted to the replayable set (`turn.post`, `comment.add`,
`comment.reply`, `attachment.add`); a mention and a topic wake on the
same op collapse to one outcome slot (the deterministic id makes the
collapse honest); a schedule wake parked at shutdown answers on the
next start.

Reversal condition: for runtime-side reads — a wake host that cannot
hold a credential able to read the record it dispatches (observable: a
dispatcher forced to proxy reads through a second service, or an agent
needing record reads mid-run that the prompt cannot carry) reopens the
widened-scope arm. For the build: none beyond design 0005/0006's own —
records a completed, measured implementation.

Trail: `soulstream-core/specs/020-system-stream/` (merged `f0a09f2`);
`soulstream-workloads/specs/009-agent-declaration/` (spec, plan, tasks,
`contracts/wake-kinds.md`; branch merged `09c446f`); designs
[`0005-agent-declaration.md`](../02-DESIGN/soulstream-workloads/0005-agent-declaration.md)
(status + §5 resolution + build findings, this change) and
[`0006-loop-safety.md`](../02-DESIGN/soulstream-workloads/0006-loop-safety.md);
research [episode 0126](0126-ecosystem-agent-declaration.md); budget
build [episode 0129](0129-workloads-the-wake-budget-builds.md).
