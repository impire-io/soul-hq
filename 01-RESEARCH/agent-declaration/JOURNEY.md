# agent-declaration — journey

Topic opened 2026-08-23. The investigation appends here as it happens.

## 2026-08-23 — hypothesis before any experiment

Recon of the shipped surfaces is running (workloads: declaration, runner,
wrap engine, natstest, minter; core: ops, artefacts, work, notify, registry;
imps: schedule channels; identity: mint.ephemeral, embed seam). No spike has
run yet. The rig hypothesis, stated first so the experiments can refute it:

- **One rig, four bars, one variable at a time.** A separate scratch Go
  module (session scratchpad, never committed here) wires the shipped
  packages against an embedded operator-mode NATS server. Bar 1 runs first
  and alone: the wake dispatcher is scratch code, the four sources are
  shipped primitives (notify subscription for mentions, ops consumer for
  topic ops, JetStream message scheduling for ticks, plain subscription for
  external subjects), the outcome op rides wrap's deterministic-id shape
  (UUIDv5 over trigger op id + persona). Expected failure mode, named now:
  the schedule source — if the pinned nats-server predates message
  scheduling, the tick channel needs either a server bump (allowed; not a
  wire change) or a scratch scheduler (which would taint the bar and must
  be recorded as such, not papered over).
- **Bar 2 rides Bar 1's rig** with one change: the agent's instructions
  move from rig constant to a stage-1 artefact; the harness is the scripted
  hermetic one (wrap's own test pattern), so "runs revision A/B" is
  observable in the outcome op's content.
- **Bar 3 swaps the credential path** from workloads' own minter to
  identity's D28 `mint.ephemeral` via the embed seam, with one granted and
  one ungranted tool; refusal must come from the transport/grant layer, not
  rig code.
- **Bar 4 is an enumeration first, a drive second** — every mutation of
  the declare flow named against a shipped surface before the rig performs
  it; any gap is the finding, not a thing to shim silently.

Reversal watch carried from the README: a mutation the rig needs that never
appears as an op; a pass that requires branching on human vs machine.

## 2026-08-23 — Bar 1: PASS [measured]

Rig: scratch module (session scratchpad), embedded open JetStream server,
personas `owner` (human stand-in) and `sprite` (declared agent), three
topics (bench for mentions, watch for topic-ops, sprite-home for
non-record wakes). Declaration carried a four-entry `wake` section:
mention / topic / schedule (`@at` one-shots on an additive schedule
stream) / subject (`signals.demo`). Dispatcher holds no durable state:
every pass re-derives position from the record via throwaway AckNone
consumers — wrap §3's "position is the record", generalized beyond
mentions. Outcome id = UUIDv5(persona, trigger) — the exact shape wrap
already exports as `WakeOpID`.

The run: 4/4 sources fired live; dispatcher stopped; a backlog built
(mention, topic-op, schedule tick landing server-side while nobody ran,
plus one core-NATS signal); fresh dispatcher instance restarted. Every
stream-backed wake in the backlog was answered exactly once. Assertions
held at two levels: stream count per outcome id == 1, and post *attempts*
across both dispatcher lives == 1 (the outcome-existence check did the
work — the duplicate window never had to absorb anything). Single run
5.43s; `-race -count=3` green in 18.6s. Sample ids in the test log
(`mention=8bbe57db…`, `op=dbbdaa6d…`, `tick=372ac296…`, `sig=3588ad30…`).

Findings the design must carry:

1. **Delivery class is per source kind, and must be declared honestly**
   [measured]: the core-subject wake fired while the dispatcher was down
   was *lost* — at-most-once, no backlog, exactly as a plain subscription
   behaves. Stream-backed kinds (mention, topic, schedule) are
   replay-exact. This mirrors imps' SubjectSource/StreamSource split; the
   wake vocabulary should carry the same seam.
2. **Schedules need one additive stream** [measured]: neither SOULSTREAM
   nor SOULSTREAM_NOTIFY enables message schedules; the rig provisioned a
   `SOULSTREAM_SCHEDULES` stream (schedules + ticks subjects,
   AllowMsgSchedules/AllowMsgTTL). Additive beside the notify stream, no
   change to existing streams or the wire — the bar's fail condition was
   not tripped. Server-owned clock confirmed: the one-shot tick landed in
   the stream while no consumer existed.
3. **The self-wake loop appears on day one** [measured]: a watched-topic
   wake answered onto the watched topic re-triggers unless the persona's
   own ops are excluded. The rig's author filter is the minimal guard;
   budgets stay the named successor topic.
4. **The shipped declaration rejects a wake section today**
   [mechanism-argument]: `declaration.Parse` strict-decodes; the
   extension is schema growth in soulstream-workloads (the roadmap's
   named hole), never a core change.
5. **Minted agent credentials cannot re-derive position under
   enforcement** [mechanism-argument, from the shipped scope set]: the
   agent permission set grants `$JS.API.INFO` only, while
   record-as-position scans need stream/consumer read APIs. The wake
   dispatcher as an enforced workload needs either a widened scoped
   template or reads done by the runtime on the workload's behalf —
   named for Bar 3 / the design, not decided here.

## 2026-08-23 — Bar 2: PASS [measured]

Same rig, one wake source (mentions) for crispness. The declaration
gained `instructions: {topic, artefact}` — a reference into the record,
never a host path (wrap §10's "registration pointing at an artifact
rather than a host command", made concrete). The dispatcher materialises
the artefact's tip at every wake — it has no field to cache it in, so
"no durable copy" holds structurally, and the digest is verified against
the object store on each load.

The run: revision A (`ALPHA`) attached as an ordinary `attachment.add`
by the owner; wake 1's outcome body carried ALPHA. Revised to `BRAVO`
through an ordinary anchored revision — **the same running dispatcher**
served BRAVO on wake 2: no redeploy, no restart, the record was the
distribution channel. Dispatcher killed, fresh instance started: wake 3
still BRAVO, and afterwards the soul topic held the full 2-revision
lineage, the object-store tip read BRAVO, and every earlier outcome
still counted exactly 1 — a dispatcher death lost scratch only, never
history. Owner authored the revisions, sprite authored the outcomes —
attribution end to end in one log. 0.93s single run; both bars
`-race -count=2` green in 14.8s.

Finding the design must carry:

6. **The artifact scheme must grow a record form**
   [mechanism-argument]: shipped `declaration.Artifact` validates
   `file://` only. Instructions-declared agents need the registration to
   accept a record reference (topic + artefact lineage, digest-checked
   at materialisation) — schema growth in workloads, same seam as
   finding 4, still nothing core.
