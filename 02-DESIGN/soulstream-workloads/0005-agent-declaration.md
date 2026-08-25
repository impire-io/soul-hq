# 0005 — Agent declaration: wake, instructions, capabilities

*Graduated from research `agent-declaration` (episode
[0126](../../04-JOURNEY/0126-ecosystem-agent-declaration.md)); every
mechanism below marked **[V]** was measured on a live rig, race-checked.
This document grows the workload declaration so an agent can be declared
— who it is, what instructs it, what it may reach, what wakes it — as a
record, and run by the existing runtime. It fills the roadmap's named
hole (the declaration trigger vocabulary) and closes wrap
[`0004`](0004-wrap.md) §10's harness-as-workload and
registrations-as-shell-objects questions. **BUILT 2026-08-25** ([episode
0130](../../04-JOURNEY/0130-ecosystem-the-agent-declaration-builds.md);
workloads `specs/009-agent-declaration`, core `specs/020-system-stream`)
— §5's `[O]` resolved **runtime-side reads**; capabilities resolution is
the named follow-on `capability-minting`.*

## 1. What this adds — and what it refuses

Growth is schema and provisioning, never machinery: new fields on the
declaration, one additive stream, and a wake engine that generalizes
wrap's measured protocol. It refuses, by constitution and by measured
reversal-watch: any store of agent state beside the record **[V —
neither rig reversal fired]**, any behaviour branching on human vs
machine, any backend detail in a declaration, and "imp" (or any word)
as an identity kind — a declared agent is a persona like every other.

## 2. The declaration, extended

The strict-decode rule stands: unknown fields refuse the document.
Existing fields (`role`, `lifecycle`, `persona`, `topic`, `artifact`,
`args`) are unchanged. New, all optional except as noted:

| Field | Shape | Meaning |
|---|---|---|
| `artifact` | grows a **record form**: `soulstream://<topic-path>/<artefact-name>` beside `file://` | the executable *or* instructions lineage the registration points at — never a host path for declared agents **[V]** |
| `instructions` | `{topic, artefact}` | a stage-1 artefact reference; the runtime materialises the lineage **tip** at every wake, digest-checked, and MUST NOT hold a durable copy — revision reaches a running agent with no redeploy **[V]** |
| `capabilities` | `{role, tools[]}` | **names, not grants** **[V]**: `role` is the vault entry name of a scoped signing key (identity D28's selector); `tools` ride as mint tags for the account's scoped template to resolve. The declaration cannot widen anything. |
| `wake` | `[{kind, …}]` | what wakes the agent — §3 |

Wake entry kinds, each carrying its **delivery class** as a normative
fact readers and shells MUST surface:

| kind | fields | delivery **[V]** |
|---|---|---|
| `mention` | — | replay-exact (notify stream, inbox-window bounded) |
| `topic` | `path`, `types[]` (default `turn.post`; the shipped engine restricts to the replayable set — `turn.post`, `comment.add`, `comment.reply`, `attachment.add`) | replay-exact (ops stream) |
| `schedule` | `name`, `pattern` (`@every` / `@at` / 6-field cron), `ttl` (optional) | replay-exact, TTL-bounded backlog (per-message `Nats-TTL`) |
| `subject` | `subject` | **at-most-once** — a wake arriving while the agent is down is lost; declaring it is declaring that honestly |

## 3. Wake semantics (wrap's protocol, generalized) [V]

- **One wake, one outcome.** Outcome op id = UUIDv5(persona, trigger
  identity); trigger identity is the notify op id (mention), the op id
  (topic), the tick's stream sequence (schedule), or a content digest
  (subject). Posted with `PostTurnIdempotent`.
- **Position is the record.** No durable consumer, no dispatcher
  state: the outcome-existence check against the topic decides; a
  restart mid-backlog answers nothing twice — measured at both the
  stream and the post-attempt level.
- **Self-wake exclusion is normative.** A `topic` wake MUST exclude
  ops authored by the declared persona — the loop appears on day one
  without it **[V]**. Budgets beyond this guard are design
  [`0006-loop-safety.md`](0006-loop-safety.md) (graduated, episode
  0128), a prerequisite for colonies (§7).
- Non-record wakes (schedule, subject) land their outcomes on the
  declared home topic; record wakes answer where they were triggered.
- **Build facts (episode 0130):** a mention and a topic wake on the
  same op collapse to one outcome slot — the deterministic id makes
  the collapse honest, stated not hidden; a schedule wake parked at
  shutdown answers on the next start (replay-exact within its TTL).

## 4. SOULSTREAM_SYSTEM — the one additive stream [V]

Schedules ride JetStream message scheduling: registering a schedule is
one headered message; the server appends ticks whether or not any
consumer runs. Neither realm stream enables it, so provisioning grows
**one** stream:

- Name `SOULSTREAM_SYSTEM`, capture `SOULSTREAM.SYSTEM.>`,
  `AllowMsgSchedules` + `AllowMsgTTL` on. Families nest under the
  prefix: `SOULSTREAM.SYSTEM.SCHEDULES.<persona>.<name>` (the
  registrations), `SOULSTREAM.SYSTEM.TICKS.<persona>.<name>` (the
  server's ticks); later system families join the same home — streams
  consolidate, they do not proliferate per feature. KV faces keep the
  bucket-per-face pattern (`soulstream-personas`, `soulstream-tools`,
  `soulstream-presence`); no system KV exists.
- The taxonomy, normative: `SOULSTREAM` is the record (signed ops
  only); `SOULSTREAM_NOTIFY` is record-grade; `SOULSTREAM_SYSTEM` is
  non-record plumbing — server-generated, unsigned, replayable, never
  authoritative. **Drift guard:** a system family whose messages start
  needing signatures or authorship has outgrown the stream and belongs
  in the record as ops. Trip-wire: a consumer materialising system
  messages into durable state no op carries.
- Reconciling a declaration's `wake.schedule` entries = publishing the
  headered message (re-publish replaces; purge deregisters). The
  agent's cron is declarative state carried by the substrate.

## 5. Capability enforcement — the identity plane, unchanged [V]

The runtime resolves `capabilities` entirely through shipped
machinery: the workload generates its user keypair locally; the
runtime (or its minter seam) calls D28 `mint.ephemeral` with the
declared role name, the persona as the user, and `tool:<name>` tags;
the account's scoped template (e.g. `SOULSTREAM.SVC.{{tag(tool)}}`)
is the entire policy, enforced at the transport. Measured: granted
tool answers; ungranted tool times out with a permissions violation on
the agent's own connection and zero responder deliveries. Who may
declare which tags stays the identity plane's named tag-policy item.

**The enforcement-read gap — RESOLVED: runtime-side reads** (build
decision, [episode
0130](../../04-JOURNEY/0130-ecosystem-the-agent-declaration-builds.md)).
The wake engine performs every record-position read (catch-up,
materialisation, outcome-existence, budget computation, ticks,
instructions) on its **own** connection; the declared agent's minted
scope stays the shipped `$JS.API.INFO`. The widened-scope arm was
rejected because the JS read tails are stream-wide — granting them
breaches own-prefix confinement — and because scoped templates are
written at account founding, making template widening a per-deployment
control-plane migration (the byon rc.10 deployment duty). Reversal
condition: a wake host that cannot hold a credential able to read the
record it dispatches reopens the widened-scope arm.

## 6. The soul topic is a guarded surface

Instructions-in-the-record means revision is privilege escalation:
whoever can post to the soul topic reprograms the agent. Normative:
the soul topic's posting rights are scoped (the operator and the
personas they name); readers verify authorship via the registry
keyring; a registration MAY pin a revision digest for high-trust
agents, making drift a refusal instead of a surprise.

## 7. Sequencing rule (from the adversarial pass, standing)

Mention-wake agents ship before topic-wake colonies. The loop-safety
research topic (agent-wakes-agent budgets; two measured exhibits in
episodes 0082/0083, a third in this topic's rig) GRADUATED 2026-08-25
([episode 0128](../../04-JOURNEY/0128-ecosystem-loop-safety.md), design
[`0006-loop-safety.md`](0006-loop-safety.md)): the composed wake budget
at admission. The prerequisite resolves to a **build** requirement: any
deployment where declared agents wake each other carries the 0006
budget in its admission path.

## 8. The declare flow is composition [V]

Every mutation of "declare an agent from the shell" is a shipped
surface; the shell verb sequences, it never stores:

| step | surface |
|---|---|
| name + shown-as | attestation token (operator's key) + profile published by the agent persona — "attested" |
| instructions | `attachment.add` on the soul topic |
| registration | the declaration itself as a record artefact |
| placement | `work.open` (fleet `Submit` is the product surface) |
| launch | `work.claim` by the runtime; boot **from the registration read out of the record** |
| liveness | the presence lease ([`presence.md`](../soulstream-core/extensions/presence.md), decided): profile on start, lease ongoing — a declared agent is the same kind of persona the wrap already is |

## 9. The reconciler's shape

A persona with ordinary scoped credentials, lifecycle as ops — "a
runner is a persona that does things, not a service tier"; no
privileged tier exists to put it in. The bootstrap fixed point stands
apart: one hand-started process per node wears a persona; declaring
the reconciler itself is coherent only in a fleet where another node
supervises it — a later rung, not the floor. Presence is advisory to
it like to everyone: liveness authority stays the fleet's
probe-before-abandon.

## 10. Acceptance criteria

1. Four wake kinds fire from one declaration; every stream-backed wake
   answers exactly once across a restart (stream count and attempt
   count both 1); the subject kind's loss-when-down is documented
   behavior, not a bug.
2. An instructions revision through ordinary ops reaches the next wake
   with no redeploy; a runtime death between wakes loses scratch only.
3. A declared agent reaches exactly its granted tools, refused at the
   transport, with zero authorization code in the runtime.
4. The declare flow's every mutation is an existing op, KV write, or
   identity-plane call; the shell performs none it could not perform
   through the public surfaces.

## Open [O]

- ~~The enforcement-read gap (§5)~~ — RESOLVED runtime-side reads
  (build decision, episode 0130; reversal condition in §5).
- ~~Loop-safety budgets (§7)~~ — built ([episodes
  0128](../../04-JOURNEY/0128-ecosystem-loop-safety.md)/[0129](../../04-JOURNEY/0129-workloads-the-wake-budget-builds.md));
  the shipped engine routes every wake kind through the budget at
  admission, so topic wakes carry their colony gate by construction.
- Capabilities resolution — schema shipped; the D28 `mint.ephemeral`
  tag lane is the named follow-on feature `capability-minting` (the
  repo keeps no identity-plane dependency until that demand).
- Runtime join/leave (a declared agent's topic set changing without a
  restart) — imps design 0003's reversal condition, watched, unfired.
- Descendant tag scoping (`{{tag(topic)}}.>`) — untested, per the
  fleet design's own admission.
