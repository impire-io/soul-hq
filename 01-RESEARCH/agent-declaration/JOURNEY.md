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

## 2026-08-23 — Bar 3: PASS [measured]

The heaviest rig: an operator-mode embedded server (own operator, SYS,
one APP/realm account carrying an unscoped signing key for the bypass
lane and one **scoped** signing key whose template is the entire tool
policy), with the identity plane run **in-process through the D29 embed
seam** and reached only through the public client. The declaration
gained `capabilities: {role, tools}` — a role name and tool names,
no policy.

The chain, every link shipped machinery: the scoped key imported as a
vault entry bound to its account (D24) under the name `agent-role`; the
agent generated its user keypair locally; `mint.ephemeral` (D28) minted
against the **declared role name** with the declaration's tools riding
as tags (`tool:toola`); the account's scoped template —
`SOULSTREAM.SVC.{{tag(tool)}}` — expanded them server-side. Granted
tool: request answered (`ok:toola`). Ungranted tool: requester timed
out, the *agent connection* surfaced "Permissions Violation", and the
responder's delivery counter read **0** — refusal at the transport,
before any service or rig code could have an opinion. The rig contains
zero authorization code and no policy store; the assertion is
structural as registered. 1.64s single run; all three bars
`-race -count=2` green in 17.9s.

Findings the design must carry:

7. **This was the ecosystem's first `{{tag(...)}}` scoped template in
   Go code** [measured]: the mechanism the fleet design (0003 §5) and
   D28 describe worked exactly as written — tags lowercased into the
   user claims, template expanded at authorization, permission-less JWT
   impossible to over-scope. The fleet design's open item now has a
   measured exhibit; descendant scoping (`{{tag(topic)}}.>`) remains
   untested, as its docs already admit.
8. **Capabilities in the declaration are names, not grants**
   [mechanism-argument]: `capabilities.role` selects a declared signing
   key (D28's own selector), `capabilities.tools` become tags the
   template may or may not honor. The declaration cannot widen anything
   — a wake-declared agent whose tags name an unbound tool gets a
   credential that reaches nothing extra. Who may *declare* which tags
   stays the identity plane's named-not-built tag-policy item (D28),
   untouched here.

## 2026-08-23 — Bar 4: PASS [measured]

The composition drive: the whole declare-from-shell flow, run through
shipped surfaces only, the shell UI out of scope as registered.

The enumeration, each step a shipped surface, each mutation in the
record or the realm KV:

| flow step | shipped surface | visible as |
|---|---|---|
| name + shown-as | `registry.NewAttestationToken` (operator's key) + `registry.Publish` by the agent persona itself | profile in `soulstream-personas`, `AttestationStatus == "attested"` |
| instructions | `attachment.add` | stage-1 artefact on the soul topic |
| wake + capabilities | the registration **is a record artefact** (`agent-decl.json`) | `attachment.add`, digest-checked |
| placement | `work.open` (the fleet's shipped pattern; `fleet.Submit` is the product surface) | work item, author `owner` |
| launch | `work.claim` by the agent; dispatcher boots **from the registration artefact read back out of the record** — no local file | claim event, author `sprite` |
| liveness | the mention wake answered under the declared instructions (`COMPOSED` in the outcome body) | `turn.post`, author `sprite` |
| retirement of the declare item | `work.done` | timeline `open:owner → claim:sprite → done:sprite` |

No step lacked a surface; the missing piece is exactly and only the
*verb that sequences them* — which is what "composition" means. Wrap
§10's "registrations-as-shell-objects" resolves as: the registration is
already a record object; the shell renders and sequences, it never
stores. 0.41s single run; the full four-bar suite `-race -count=2`
green in 18.9s.

Addressed head-on, because the reversal condition watches for it: the
one mutation in the whole topic that never appears as an op is the
identity plane's role-key import (Bar 3). That is not a second
coordination store — it is the shipped identity pillar doing exactly
what its constitution demands (keys never in the record; the process
boundary is the custody boundary), it is per-role operator provisioning
rather than per-agent state, and the declaration references it by name
only. Nothing agent-shaped lives outside the record. **Neither reversal
condition fired**: no bar needed a store beside the record, and no code
path anywhere branches on human vs machine — the persona flow Bar 4
drove is byte-for-byte the flow a human persona uses.

## Where this stands

All four pre-registered bars: **PASS**, each `[measured]` on a live rig,
race-checked. The question's answer as measured: yes — "declare an
agent from instructions" decomposes entirely into existing vocabulary,
with the growth confined to (a) a `wake` section + record-form
`artifact` scheme in the workloads declaration (the roadmap's named
hole, schema not machinery), (b) one additive schedule stream beside
the notify stream, and (c) a wider read scope (or runtime-side reads)
for enforced record-position agents. Findings 1–8 above are the design
doc's spine. Named successors stand: agent-wakes-agent loop safety
(before any colony), runtime join/leave (watched, unfired), tag policy
(identity plane's own item).

Graduation (`/research-graduate agent-declaration --to design`) stays
the owner's call after the teach-back the Working Agreement requires —
the adversarial pass belongs in that conversation, not appended here.

## 2026-08-24 — the additive stream is SOULSTREAM_SYSTEM [judgment, owner's call]

The owner named the Bar 1 stream: not a per-purpose
`SOULSTREAM_SCHEDULES`, but **`SOULSTREAM_SYSTEM`** — one additive
stream for the realm's non-record substrate plumbing, one capture
prefix (`SOULSTREAM.SYSTEM.>`), families nested under it
(`SOULSTREAM.SYSTEM.SCHEDULES.<persona>.<name>`,
`SOULSTREAM.SYSTEM.TICKS.<persona>.<name>`, later families beside
them). Rationale: the stream count must not grow per feature; a system
stream serves the next non-record need without another provisioning
artefact.

This sharpens the taxonomy rather than blurring it:

- **SOULSTREAM** — the record. Signed ops only.
- **SOULSTREAM_NOTIFY** — record-grade (mention.notify is a full
  signed record), per-persona bounded. Stays its own stream *because*
  it is record-grade with its own retention shape.
- **SOULSTREAM_SYSTEM** — non-record plumbing: server-generated,
  unsigned, replayable but never authoritative. Schedules and ticks
  today.

The drift guard, written now: SOULSTREAM_SYSTEM must never become a
shadow record. Anything worth keeping still flows to topics as ops
(constitution I); a system family whose messages start needing
signatures or authorship has outgrown the stream and belongs in the
record. Observable trip-wire: a consumer materialising system messages
into durable state that no op carries.

Rig updated and re-measured: all four bars green under `-race` with the
single `SOULSTREAM.SYSTEM.>` capture (9.97s). Per-message `Nats-TTL`
(AllowMsgTTL is on) keeps stale-tick expiry per schedule — the wake
vocabulary's `ttl` knob rides family-level headers, not stream config,
so co-tenancy costs schedules nothing.

## 2026-08-24 — teach-back held; direction survives with precisions

The owner's restatement, held against the evidence:

- **"Declare agents on the soulstream; workloads make them run"** —
  survives [measured, Bars 2/4], with one precision: what was measured
  is declaration-as-record-artefact reaching the runtime through the
  fleet's placement grammar (work item, first-claim-wins, sweeper
  reclaim). A watch-and-converge reconciler was *not* measured.
- **"Not merging imps; agents/imps spawned by workloads"** — survives;
  the triangle stands (record / room / inhabitant), instruction-declared
  agents and coded imps ride the same declaration path.
- **"Heartbeats as KV puts with TTL on a system KV"** — *not a finding*;
  no bar touched liveness. Recorded as a **named follow-on decision**,
  not evidence: it fits the SYSTEM taxonomy (non-record, unsigned,
  expiring) and per-key TTL is real on the pinned server
  [mechanism-argument, unmeasured]; a KV bucket is a stream under the
  hood, so the rationale is "no per-feature proliferation", not "no new
  streams"; TTL expiry emits no tombstone — absence must be *polled*,
  which is the sweeper shape fleet 0003 already has. Interacts with
  fleet's decided probe-before-abandon: a heartbeat KV can be that
  rule's transient evidence, as a fleet-design amendment.

**Opened by the owner: the reconciler's shape.** Should the
declaration-watcher be an agent? The house grammar answers the shape:
a persona with ordinary scoped credentials, lifecycle as ops
(work.md: "a runner is a persona that does things — not a service
tier"; constitution II forbids a privileged tier). The bootstrap fixed
point stands apart: one hand-started process per node *wears* a
persona; declaring the reconciler itself is only coherent in a fleet
where another node supervises it — a later rung, not the floor. The
loop-safety successor applies to it doubly.

**Adversarial pass delivered** (recorded at full strength before any
graduation): (1) wake-vocabulary-belongs-to-the-inhabitant — refuted
because wake determines credential scope and placement, which the
inhabitant must not self-declare; (2) instructions-in-the-record makes
revision a privilege escalation — *stands as a design obligation*: the
soul topic is a guarded surface (posting rights, verified authorship,
optionally a pinned revision digest in the registration); (3) declared
agents make colonies one op away while loop safety is unfinished —
*stands as a sequencing rule*: mention-wake agents ship before
topic-wake colonies.

## 2026-08-24 — convergence with main's presence thread; two notes corrected

Read (not touched) on main's in-flight tree, drafted the same day as
the teach-back: `02-DESIGN/soulstream-core/extensions/presence.md`
(proposed, not decided) plus shell design 0008, the operational thread
making today's *wrap* agents first-class (profile on start, aliveness
ongoing). The convergence is real and the threads compose rather than
collide:

- **The heartbeat follow-on is presence.md.** Yesterday's named
  follow-on decision dissolves into main's draft, which decides two
  points *better than the teach-back wording*: (a) **no TTL** — the
  entry never expires; freshness is the reader's judgment against the
  KV entry's own timestamp, and "last seen" evidence survives. This
  supersedes "the entry eventually disappears" and dissolves this
  journal's no-tombstone caveat entirely — nothing needs to watch for
  deletion because nothing is deleted. (b) **Advisory, never
  authority** — presence informs courtesy, and the fleet never imports
  the reader. My note that a heartbeat KV "can be fleet 0003's
  transient evidence" is hereby **withdrawn**: that coupling is exactly
  what the draft's advisory rule refuses; the fleet keeps
  probe-before-abandon.
- **The SYSTEM consolidation is unaffected and reconciles cleanly**:
  *streams* consolidate under `SOULSTREAM_SYSTEM` (schedules/ticks —
  this topic's measured decision, unchanged); *KV faces* follow the
  house bucket-per-face pattern (`soulstream-personas`,
  `soulstream-tools`, `soulstream-presence`). No "system KV" is needed.
- **Sequencing, in main's own words**: shell 0008 lands presence
  "without waiting for the record's declaration vocabulary" — and this
  topic *is* that vocabulary. Main makes hand-run wrap agents visible
  citizens today; this topic makes agents declarable so the runtime
  spawns them tomorrow. A declared agent inherits profile-on-start and
  the presence lease for free by being the same kind of persona
  (Bar 4 already had it self-publish its profile), and presence's
  "each thing writes its own key, no collector, no privileged
  side-channel" is the same no-privileged-tier grammar the reconciler
  answer used.

Graduation should cite presence.md as the liveness answer (once it is
decided on main) instead of carrying a liveness design of its own.
