# 04-JOURNEY — the narrative record

What was built, what was measured, what was believed and then refuted, and
what each episode taught — for the whole ecosystem, in one chronological
sequence. The design docs in `../02-DESIGN/` and the frozen specs in the
component repos say what each system *is*; these episodes say how we *got
here* — including the reversals, because a refuted assumption is as
load-bearing as the shipped code.

> **Keeping this log alive:** whenever a feature lands, a research
> investigation concludes, or a load-bearing decision is made, add a numbered
> episode with `/journey-log` (research topics get theirs via
> `/research-graduate`). Episodes are named `NNNN-<component>-<slug>.md` —
> one shared sequence; component tags are single-word (`core`,
> `workloads`, `identity`, `idp`, `shell`, `mcp`, `cli`, `soulstream`
> for the product, `ecosystem` for cross-cutting). Episodes ≤ 0069 keep
> the pre-rename tags — the [naming map](#the-naming-map-2026-08-13)
> resolves them. Follow
> [`TEMPLATE.md`](TEMPLATE.md) — including its required Reversal-condition
> line and evidence-class tags. Honesty rules apply here as everywhere:
> record what actually happened, including failures, reversals, and findings
> that contradicted expectations. This duty is anchored in
> [`../00-GENESIS/how-we-work.md`](../00-GENESIS/how-we-work.md); the
> numbering and index are enforced by `internal/hqlint`.

Episodes `0001`–`0049` were carried from the five per-project journals when
the hqs merged (2026-08-02, episode
[0050](0050-ecosystem-one-hq.md)) and renumbered into this shared sequence —
ordered by the original commit time of each episode, links retargeted,
prose untouched. The [numbering map](#pre-merge-numbering-map) below
resolves every pre-merge reference (frozen specs and old commit messages
cite the per-project numbers).

## Where things stand, per component

The per-component summaries below were carried whole at the hq merge and are refreshed with every new episode, as before.

### soulstream-core — the record (as of 2026-08-25; named soulstream until episode 0069)

**Sealed topics BUILT — the locked binder** ([episode
0131](0131-core-sealed-topics-build.md); `specs/021-sealed-topics`,
merged `724f10d`, pre-released the same day in `v0.14.0-rc.1`,
episode [0132](0132-soulstream-the-rc-carries-both-builds.md)): the
design-validated extension is real —
`sealed.op`/`sealed.epoch` with epoch and nonce inside the signed
payload, the AAD splice matrix refusing every arm [measured], the Bar 3
rollup re-carrying the wrapped-key table (a fresh member reads from the
post-purge baseline [measured]), endorsed X25519 sealing keys in the
registry (strict decode ships first — Bar 2), the `Unwrapper` custody
seam mirroring `Signer` for D9's custodian, curator blind and memory
silent on purpose. The build caught one design blind spot: sealed topic
ids derive from the word `sealed`, never the display name they would
have leaked. Eg-walker is now the last chafe-gated unbuilt core design.

**The system stream lands — schedules ride the substrate** ([episode
0130](0130-ecosystem-the-agent-declaration-builds.md); feature
`020-system-stream`, merged `f0a09f2`, pre-released the same day in
`v0.14.0-rc.1`): `SOULSTREAM_SYSTEM`
provisioned create-or-report with `AllowMsgSchedules` + `AllowMsgTTL`
and the `SCHEDULES`/`TICKS` subject helpers — non-record plumbing by
taxonomy, server-generated, unsigned, never authoritative. The design
0005 bet held [measured]: a `@every 1s` registration produced a server
tick with no consumer running, and per-message TTL expiry is real.

**Presence BUILT — v0.13.0, the same day it was decided** ([episode
0125](0125-ecosystem-the-presence-lease-builds.md)): the `presence`
package in the tool catalog's shape plus the new ground — `State`
carries the KV entry's own timestamp as the renewal moment,
`Read(now)` derives *present / left / last-seen* as the reader's
judgment, `Hold` is the whole lease in one call, and renewal is a
plain put (after a crash-and-restart a stale revision must never make
the fresh writer lose). The lease-honesty triple, additive round-trip,
and the unchanged op-log census are standing tests [measured]; the
docs page ships with it. The wrap is its first writer (soulstream
branch `011-presence-lease`, episode 0125).

**Presence is a decided extension — the who-is-around face** ([episode
0124](0124-ecosystem-the-first-hour-and-the-presence-lease.md);
[`extensions/presence.md`](../02-DESIGN/soulstream-core/extensions/presence.md),
grown from the thin paragraph in `library-and-adapters.md`): a
`soulstream-presence` KV holding one lease per persona for everything
that runs — renewed on a cadence, a farewell as manners, staleness
read as gone, because **departure is derived, never merely announced**
[mechanism-argument: a crashed writer cannot say goodbye]. The store
forgets nothing (no TTL-delete — *left* and *vanished* stay
distinguishable), each thing writes its own key on its own admission
(no collector), and the payload line keeps it presence rather than
telemetry: words a person reads, never numbers a machine steers by —
advisory sharpened to *courtesy, never correctness*. The wrap is the
first mandatory writer (shell 0008's upstream ask #3); unbuilt, not
chafe-gated.

**The canonical form broke clean — v0.11.0/v0.11.1** ([episode
0112](0112-ecosystem-the-canonical-form-breaks-clean.md)): every
signature binds the realm's cryptographic identity (A10 — born at
first provision, first wins) and every record carries the required
acting credential (E3), custodian-verified on the sign.record lane.
Reads never hard-fail: v1 records parse and verify as the named
`legacy-shape`. The whole ecosystem re-pinned in one pass; three real
defects surfaced and were fixed, every other consumer compiled
unchanged.

**F1 and C4 are built — v0.9.0 and v0.10.0, the same evening they were
designed** ([episodes 0108](0108-ecosystem-the-key-becomes-resolvable.md)/
[0109](0109-ecosystem-consent-enters-the-record.md)):
`registry.EnsureSigningKey` wired into all three signer consumers
(unknown-key stops being the shipped default), and the grant vocabulary
with its projection, dual attribution, rollup baking — **Bar 4 measured
PASS on the full composition** in identity's embedgate.

**The record's tenancy half is designed — and two canonical-form
amendments are decided** ([episode
0107](0107-ecosystem-platform-tenancy-guardrails.md);
[`extensions/tenancy.md`](../02-DESIGN/soulstream-core/extensions/tenancy.md)):
the grant-record vocabulary (`grant.issue`/`grant.revoke` + dual
exercise attribution, Bar 4 its build gate, the projection duty
answering 0029's objection), the registry's ensure-signing-key act
(F1 — confirmed by code trace as the shipped default: no reader
consults `keys.public`, no signer consumer publishes a profile), the
self-declared responds-when-addressed field — and the operator's two
pre-v1 clean breaks, landing with their builds: **A10** the account
key replaces the realm name in the canonical signed form, **E3** a
required acting-credential field, custodian-verified on the
`sign.record` lane, testimony-grade on self-custody.

**The node module is consumable** ([episode
0010](0048-soulstream-the-node-becomes-consumable.md)): its landing-day
`replace => ../` dropped the same day v0.7.0 made it unnecessary — the
module now pins the tag and downstream compositions can require it
(soulnode's Phase 2 front door is the first). Full node suite measured
green against the tag; co-development rides an untracked `go.work`
(soulrealm 0011's discipline, adopted).

The reference library has shipped the MVP and most of day-2 — **`v0.7.0`**
(2026-08-02) is current: foundation + op-log engine, CLI + MCP clients,
signing, rollup, scatter-gather discovery, the curator, work stages 1–2,
distribution, config-file identity, persona accountability
([episode 0001](0007-soulstream-genesis-and-the-reference-library.md), the founding
retrospective), the **memory convention**
([episode 0003](0009-soulstream-memory-convention-and-exhibits.md)): collective search as
graded testimony, portable self-authenticating exhibits, a public witness
surface — with the first archivist a separate repository,
[impire-io/soulstream-archivist](https://github.com/impire-io/soulstream-archivist),
**verified live against the NGS realm** (2026-07-26) — and **provisioning
byte limits** ([episode 0004](0010-soulstream-provisioning-byte-limits.md)): limit-
enforced accounts provision out of the box. Just landed: the **signer seam**
([episode 0006](0026-soulstream-the-signer-seam.md)) — signing delegated through
`identity.Signer` to an external custodian (SoulIdentity's M2 wiring
point), local keys the first implementation, a failing signer failing the
publish loudly — hardened for release the same day
([episode 0007](0027-soulstream-dx-hardening-and-the-cycle-guard.md)): typed-nil
signers refused at `Connect`, responder callbacks carrying the error
instead of a `-1` sentinel, and the cycle-guard dependency rule (neither
core repo imports the other; consumers wire the structural interface)
recorded on both sides — shipped together as **`v0.6.0`**. Then the **remote
MCP node** ([episode 0009](0047-soulstream-remote-mcp-node-built.md), **`v0.7.0`**): a
URL into the realm for clients that cannot install anything — credential-free
bearer passthrough onto per-principal callout-admitted connections, the tool
surface now a public embeddable `mcpserver` (SoulNode's fourth upstream ask,
plus `soulstream_whoami`), sign-in via an external OIDC authorization server
only (soulfold the intended default, the AS-facing contract proven to be the
interface), all five user stories measured on an in-process admission edge,
and a trust model that closes the prototype's forged-hint DoS. The **two-week
dogfood run started 2026-07-27** (protocol:
[`../03-IMPLEMENTATION/DOGFOOD.md`](../03-IMPLEMENTATION/DOGFOOD.md))
— daan, smith, and scribe on the NGS realm, the archivist keeping; its chafe
log feeds the eg-walker and sealed-topics gates. The central architectural
bet — leaderless coordination, no coordinator and no consensus — stands,
un-refuted.

**The project's working structure lives in `hq/`**
([episode 0002](0008-soulstream-adopting-the-hq-way.md)): GENESIS holds the vision, the
constitution (v1.1.0, wired into every spec-kit plan via the
`.specify/memory/constitution.md` symlink, now carrying the anti-drift Working
Agreement), and how-we-work; research runs a `/research-start` →
`/research-graduate` lifecycle; this journal is numbered episodes with the
structure enforced by `internal/hqlint` on the standard gate. The first
research topic has concluded: **sealed-topics**
([episode 0005](0011-soulstream-sealed-topics.md)) ran its four pre-registered bars in
a day and graduated to design — the sealed design survives the shipped
substrate, with amendments now folded into
[`../02-DESIGN/extensions/sealed-topics.md`](../02-DESIGN/soulstream-core/extensions/sealed-topics.md)
(the `{"ct"}` payload wrapper, signature-covered epoch/nonce, signing-chain
endorsement of sealing keys, key-carrying sealed baselines); the build's
priority stays gated on the dogfood chafe log (to 2026-08-10). A second
research topic has concluded: **remote-mcp-node**
([episode 0008](0038-soulstream-remote-mcp-node.md)) graduated to design after Bars 1–3
and its reversal-condition measurement PASSED on a live Synadia Cloud BYON —
a credential-less MCP node that passes the caller's bearer through to auth
callout, admitting a no-install client as a real, signed realm member. Bar 4
found the one gap that shaped the build: Claude Desktop's hosted connector
authenticates by OAuth only (no static-header lane), which the node answers
with an external OIDC authorization server — that topic is now **built and
shipped** as feature 018 ([episode 0009](0047-soulstream-remote-mcp-node-built.md)). What
is *not* yet built is the rest of the forward plan in
[`../03-IMPLEMENTATION/ROADMAP.md`](../03-IMPLEMENTATION/ROADMAP.md):
eg-walker live co-editing, sealed topics, and a browser/WebSocket client.

### soulstream-workloads — the room (as of 2026-08-28; named soulrealm until episode 0069)

**The dispatcher is DESIGNED — five bars measured in one day**
([episode 0141](0141-ecosystem-agents-as-infrastructure.md); design
[`0007-agents-as-infrastructure.md`](../02-DESIGN/soulstream-workloads/0007-agents-as-infrastructure.md)):
the `agents-as-infrastructure` research graduated the morning after
its pre-registration — submit-and-forget from the log alone
(submit→served ~160ms, restart resumes with no re-claim, hard-kill
serves exactly once), placement by the 0003 claim path with the
serving handoff (failover answered in ~1.05s, restart-dedup and
failover-dedup one mechanism), the declared budget riding
submission→admission (cycle halts at its bound, delegation clean),
custody structural (no canonical scope carries a secrets tail; the
provider secret resolves at wake time in 1.98ms and injects through
the shipped `Template.Env`), and the whole loop drivable from a
session admission in 561ms [measured, all 3× -race]. The finding: the
dispatcher is a **composition of shipped mechanisms, not new
mechanism**. Unbuilt; §9's [O]s (serve seam, `inference` block schema,
founding role naming) go to spec.

**The focus turned to the room's missing daemon — agents as
infrastructure** ([episode
0140](0140-ecosystem-the-focus-agents-as-infrastructure.md)): the
operator's demand-driven direction the evening rc.2 shipped — the next
builds are the **standing dispatcher** (submit-and-forget declared
agents: design 0004 §9's recorded reversal condition firing by its own
words) and the shell's declare surface, with **inference providers and
models** named part of the storyline (a dispatcher-served agent has no
signed-in person to wrap). The research ran the same evening and
graduated the next morning (episode 0141, above).

**Capability minting is BUILT — the declaration's names become the
credential** ([episode 0137](0137-ecosystem-capability-minting.md);
`specs/010-capability-minting`, merged `14e95a0` + `2465dba`,
pre-released in `v0.8.0-rc.2`, episode
[0139](0139-soulstream-the-rc2-carries-the-tenants-and-the-capabilities.md)):
the schema-only `capabilities` block reaches the minter
seam — `Scope` carries the selectors, `MintTags` renders the
`tool:`/`topic:`/`persona:` vocabulary in one refusing surface (a
subject-grammar value never leaves), the local minter narrows
`SOULSTREAM.SVC.>` to exactly the declared tools (empty list grants
none), and the operator rig measures granted-answers /
ungranted-refused with zero responder deliveries [measured]. The
load-bearing find rode the product wiring: the planned D28 vault-role
import would trip the binding-resolved ambiguity refusal and break the
token lane, so the scoped lane ships **local-first** —
`ScopedSigningKeyMinter`, the identical D28 claim shape signed by a
deployment-held role seed (open amendment in spec 010); the op lane
stays the fleet-era path gated on the token lane's named-role answer.
The repo still imports nothing of the identity plane.

**Agent declaration is BUILT — the record declares, the room runs**
([episode 0130](0130-ecosystem-the-agent-declaration-builds.md);
`specs/009-agent-declaration`, merged `09c446f`): the declaration grows
`instructions` (materialised at the lineage tip every wake,
digest-checked, scratch-only — a revision reached the next wake with no
redeploy [measured]), `capabilities` (schema this slice; the D28 mint
lane is the named follow-on `capability-minting`), the 0006 `budget`
block, the `soulstream://` record-form artifact (digest-tamper refuses
and abandons [measured]), and all four `wake` kinds through wrap's
five-step admission — exactly-once across an engine restart for every
stream-backed kind, subject honestly at-most-once, an uncooperative
topic-wake cycle halted at 2K=8 with op-less loud refusals [measured].
The enforcement-read gap `[O]` RESOLVED: **runtime-side reads** — the
engine's credential reads the record; the agent scope stays
`$JS.API.INFO` (stream-wide JS read tails would breach own-prefix
confinement; founding-time scope templates make widening a
per-deployment migration — the byon rc.10 lesson). Topic wakes ship
*with* their colony gate: the budget sits in the dispatcher's admission
path, satisfying design 0006 §6 by construction.

**Loop safety is ANSWERED and BUILT — the colony gate ships, same day
as its research** ([episodes 0128](0128-ecosystem-loop-safety.md)/[0129](0129-workloads-the-wake-budget-builds.md);
design
[`0006-loop-safety.md`](../02-DESIGN/soulstream-workloads/0006-loop-safety.md);
`specs/008-loop-safety`, merged `af49b80`):
all four pre-registered bars PASS [measured], same day as its opening. The
danger is a number now — an unenforced three-agent colony amplified one
human op into 434 ops in 0.34s (1,264.7 ops/s), doubling per generation —
and the composed wake budget at admission (authorship-window floor,
unforgeable because authorship is mechanical; provable-chain depth bound
over the `WakeOpID` UUIDv5 binding) halted every measured cascade at its
pre-computed bound: the uncooperative ping-pong at exactly D, the
id-evading self-post cascade at exactly 2K, the colony at exactly N·K —
while a legitimate owner→A→B→A delegation completed with zero refusals.
Refuted by measurement: a budget in the harness slot (refusal became a
312-turn failure ping-pong — the gate sits at admission, op-less and
loud, or nowhere) and depth alone (404/404 self-posted outcomes invisible
to record-only ancestry). The build landed the same day (episode 0129):
wrap evaluates the budget at admission with defaults on (D=4, K=8/10m),
`Unbudgeted` the explicit opt-out, the walker shipped as diagnostic —
the rig's cases green as the feature's own `-race` suite. Colonies now
wait only on a dispatcher that carries this gate at its admission seam.

**Agent declaration is DESIGNED — all four research bars measured**
([episode 0126](0126-ecosystem-agent-declaration.md); design
[`0005-agent-declaration.md`](../02-DESIGN/soulstream-workloads/0005-agent-declaration.md)):
the declaration grows `wake` (four kinds with honest delivery classes;
the trigger vocabulary the roadmap named), `instructions` (a stage-1
artefact materialised per wake — revision reaches a running agent with
no redeploy), and `capabilities` (names resolved by D28 tag-template
mints — the ecosystem's first measured `{{tag()}}` template); schedules
ride the one additive `SOULSTREAM_SYSTEM` stream; the declare flow is
composition of shipped surfaces end to end. Built two days later
([episode 0130](0130-ecosystem-the-agent-declaration-builds.md)) with
the enforcement-read gap resolved runtime-side.

**The fleet lands — v0.6.0, placement IS work.claim** ([episode
0113](0113-workloads-placement-is-work-claim.md)): a submission is an
ordinary work item, every idle node races, the log decides — measured
on two nodes, four contested placements each run by exactly one node,
a live owner never reclaimed and a silent one reclaimed as an ordinary
abandon with no double close, no probe traffic on the stream, and the
runner/declaration/backend seams untouched. Named, not built: a
long-running serve loop.

**The template grew `mcp_args` — a subcommand can be the tool door**
([episode 0089](0089-ecosystem-wrap-in-the-house.md);
`specs/007-mcp-args/`, **v0.4.0**): additive through preset and the
per-run MCP config, absent when empty, so the product binary can point
a wrapped harness at its own `soulstream mcp` verb — no second binary
on the agent's machine.

**The waker became wrap the day it landed — run your agent where you
are** ([episode 0085](0085-workloads-wrap-run-your-agent-where-you-are.md)):
the operator retired the byname and repositioned the front door as
personal — `soulstream wrap --harness claude` (core v0.8.4's new
external-subcommand seam) wraps the assistant already signed in on your
machine, dissolving the provider-login question outright. The central
daemon is cut with its reversal recorded (design 0004 §9): the wrapper
holds only the agent's credential and no consumer state — outcomes
publish under deterministic ids, so **the record is the position** —
catch-up from the bounded inbox, live from a raw subscription, restarts
answer nothing twice, failures are the agent's own self-report. Real
claude answered through it in 19s, and its reply's self-mention
exercised the loop guard unprompted.

**The waker was built first — agents are addressable, not only runnable**
([episode 0083](0083-workloads-the-waker-lands.md), same day as its
research): `soulstream-workloads waker serve` holds a durable consumer per
registered agent and turns a mention into one harness invocation with
exactly one outcome op per admitted wake; `make test-wake` woke a real
`claude -p` (mention → attributed reply, 6.6s). The build rode core
v0.8.3's new `PostTurnIdempotent` and caught two design-level bugs the
research rig never hit — the self-wake loop (failure testimony tapping the
agent re-woke it; now it taps only the asker) and the outcome-id collision
(one mention tapping two agents deduped their outcomes; the wake id now
hashes notify op **and** persona) — both corrected in design 0004 and
regression-guarded.

**The waker was designed first — all four research bars measured PASS**
([episode 0082](0082-ecosystem-agent-participation.md)): a mention wakes a
headless harness through a durable consumer and an invocation template,
every admitted wake ends in exactly one outcome op (proven under kill,
timeout, and mid-run MCP posting), revocation refuses the next wake in
~2ms while the persona stays addressable, and a second harness costs a
template, not code. Design
[`0004-wrap.md`](../02-DESIGN/soulstream-workloads/0004-wrap.md):
the workload plane's trigger arm, with support-layer standing (the waker
mints for the agent; the agent cannot mint for itself). Loop safety
(agent-wakes-agent) was named **[O]** here and left as a successor topic —
answered in [episode 0128](0128-ecosystem-loop-safety.md).

**Soulstream is pinned at v0.6.0 — the dev replace is gone** ([episode
0011](0037-soulrealm-pinned-to-the-record.md)): soulnode's composition research named
the filesystem `replace` a release blocker for any consumer wanting to pin
soulrealm; soulstream's tag proved current (main = v0.6.0 + docs only
[measured]) and the whole gate runs green against it. Co-development now
rides an untracked `go.work`; soulstream changes arrive by tag bump.

**The project was founded** ([episode 0001](0001-soulrealm-genesis.md)): soulrealm is the
runtime companion to soulstream — soulstream records, soulrealm runs. The hq is
bootstrapped from the sibling project's proven structure, with a constitution
whose non-negotiable article is the **substrate boundary** (soulrealm never
becomes a store of record; everything worth keeping flows back to topics as
ops).

**The substrate question is decided** ([episode
0002](0002-soulrealm-the-substrate-decision.md)): after a live NEX spike and a source
read, **soulrealm builds its own runtime — NEX as influence, not dependency —
with the soulstream op-log as the single control plane.** Measured: role
(`agent`/`tool`) is orthogonal to lifecycle (`service`/`function`/`job`); NEX
issues scoped per-workload identity for free and is embeddable via public
options (so a fork was never forced). The `[judgment]` call to rebuild rather
than embed turned on **not running a second control plane** (`$NEX.control.*`)
beside the op-log — recorded after teach-back, with the embed case argued at
full strength first. That opened design doc
[`0001-soulrealm-runtime.md`](../02-DESIGN/soulstream-workloads/0001-soulrealm-runtime.md): the
single-plane runtime, the role×lifecycle model, a realm-semantic per-workload
minter, lifecycle-as-ops, and pluggable backends, with an honest NEX influence
ledger and its `[O]` open sub-questions.

**The first slice is specced** ([`specs/001-launch-an-agent/spec.md`](../../soulstream-workloads/specs/001-launch-an-agent/spec.md)):
declare an `agent`/`service`, mint a persona-scoped credential, launch it
native, post a turn as its persona, lifecycle visible as ops — no second
control plane. Minimal spec-kit scaffolding is in place (`.specify/` with the
constitution symlinked). The signing story is resolved soulrealm-held.

**Scope is soulstream-only** ([episode 0003](0003-soulrealm-soulstream-only-scope.md)):
soulrealm depends on soulstream and nothing else — no Impire-platform services
(identity, tenancy, vault) for now. The minter stays a seam for a future
external authority, but none is designed in.

**M1.1 is implemented** ([episode 0004](0004-soulrealm-the-first-agent-runs.md)): the Go
module `github.com/impire-io/soulrealm` exists, and an agent launched by
soulrealm posts a turn attributed to its persona while its lifecycle shows up
as `work.open/claim/done` on the topic — proven end-to-end (SC-001, SC-002),
whole gate green. The plan's bet held: **no new soulstream vocabulary** —
soulrealm is the work.md "runner". Six packages (declaration, minter,
backend/native, runner, two cmds), pure logic split from I/O so most tests need
no server; the native backend proves it does not leak soulrealm's secrets into
a workload. All five success criteria met (SC-003 enforcement via an in-process
operator-mode server).

**M1.2 is done** ([episode 0005](0005-soulrealm-a-tool-answers.md)): soulrealm launches a
`tool` service and an agent discovers it by name and calls it (uppercase round
trip). Added the tool role, role-aware scopes, and the runner's launch/stop
(services don't self-exit). A measured lesson landed a boundary: tool
request-reply is transient, so it rides soulrealm's own `SOULREALM.SVC.*`
subjects, not the stored `SOULSTREAM.>` stream (which would ack and race the
reply). SC-001/002/003 proven end-to-end.

**The hq is now aligned with its own contract** ([episode
0006](0006-soulrealm-hq-alignment.md)): the "hq structural lint" the constitution and
skills had cited as the enforcement backbone is finally built —
`internal/hqlint`, a test-only Go package that rides `make test` and the commit
gate. Along the way: README/CLAUDE status corrected to Phase 1 (M1.1 + M1.2
done); specs 001/002 marked Shipped with 002's spec-kit short-circuit recorded
honestly; the full spec-kit flow vendored from pra (which also fixed a
plan/tasks template that had baked in pra's constitution principles); and
Article VI clarified (constitution 0.1.1).

**M1.3 is done — Phase 1 is complete** ([episode 0007](0020-soulrealm-a-second-wall.md)):
the byte-identical M1.1/M1.2 declarations run inside **microsandbox**
microVMs (an open amendment: the roadmap had said Docker/Firecracker), with
the identical op mapping on the topic and a measured isolation boundary — a
host path readable natively is denied in-guest. The `msb` CLI is supervised
as a child process (no CGO SDK); loopback NATS is rewritten to the guest's
host alias under a host-only network policy; backend selection is node-side
(`SOULREALM_BACKEND`), and the declaration still cannot name a backend. The
default suite stays hermetic (stub CLI); `make test-msb` boots real microVMs.
A refuted diagnosis is on record (":ro unsupported" was really msb failing on
symlinked mount sources), plus a named limitation (remote NATS needs the
`public` profile — Fleet-era).

**The Kubernetes gate is met** ([episode 0008](0024-soulrealm-kubernetes-backend.md)):
the `kubernetes-backend` research topic pre-registered four bars and four
spikes measured all four PASS — a prototype backend behind the *unchanged*
seam ran the byte-identical M1.1/M1.2 declarations as pods on kind (agent,
tool round trip, crash → abandon), and a scope probe inside a pod against
Synadia NGS was denied out-of-scope with its credential delivered as a
Secret. One expectation inverted: on non-loopback NATS, Kubernetes is
*ahead* of the microVM backend (ordinary pod egress vs msb's Fleet-era
`public` profile). Named honestly: pods are weaker isolation than microVMs —
the case is adoption, not isolation. Opened design
[`0002-kubernetes-backend.md`](../02-DESIGN/soulstream-workloads/0002-kubernetes-backend.md).

**M2.1 is done — Phase 2 is complete** ([episode
0009](0028-soulrealm-a-third-wall-lands.md)): `backend/k8s` landed through the full
spec-kit flow — one runner-supervised pod per workload, credential as a
Secret that never touches host disk, artifact as a per-run OCI image on the
CA-trusted base pushed digest-pinned to the operator's registry. Five e2e
scenarios green on kind + a local registry in ~26 s (`make test-k8s`);
default gate hermetic; `make test-msb` still green. Two recorded decisions
from planning: the artifact channel **reversed** from the draft's node HTTP
to an OCI-registry interface (maintainer decision — an open amendment to
design 0002's candidates), and client-go was chosen after teach-back, kept
entirely inside `backend/k8s`. **Next:** research gates for the later
horizons — Fleet, sandboxes (stage 5), tool ecosystem — and the `nats://`
artifact-addressing question at the artifact-registry milestone.

**The Fleet gate is met** ([episode 0010](0033-soulrealm-fleet.md)): the `fleet`
research topic pre-registered three bars and four spikes measured all three
PASS — placement **is** `work.claim` (exactly-one-launch 120/120 contested
rounds, replay-reconstructible, zero transient signaling), node death
reclaims via *projection nominates → probe vetoes → ordinary `work.abandon`
decides* (10/10 kills per variant within bound; the probe eliminates the
live-silent false positive at zero cost on true deaths), and a node without
the signing seed launches a scope-enforced workload (expiry floor measured
at 10 ms). Two open reversals on the record: Bar 2 amended pre-run when
work.md's timeout-by-projection surfaced, and spike 3's judgment against
scoped signing keys fell to spike 4's tag-template measurement — the minter
role dissolves into the identity plane (`soulidentity`), amending episode
0003's soulstream-only scope. Opened design
[`0003-fleet.md`](../02-DESIGN/soulstream-workloads/0003-fleet.md); roadmap Phase 3 (Fleet) is
unblocked. **Next:** the spec-kit pass for the first Fleet milestone; the
soulidentity tags-on-mint addition gates the preferred minting path.

### soulstream-identity — the name (as of 2026-08-27; named soulidentity until episode 0069)

**The sealed record gains its custodian — D9 builds** ([episode
0138](0138-identity-the-sealed-record-gains-its-custodian.md); design
[`sealing-keys.md`](../02-DESIGN/soulstream-identity/sealing-keys.md)
D50–D53, `specs/005-sealing-custody` merged `72bd164`, released in
`v0.12.0`, episode
[0139](0139-soulstream-the-rc2-carries-the-tenants-and-the-capabilities.md)):
the `persona-sealing-key` vault kind with D26's first-touch clause
realized, one `seal.unwrap` op releasing artifacts and never material,
`keys.public` growing the `sealing/` grammar (one directory door), the
persona template growing exactly one tail, and the client's
`PersonaUnwrapper` satisfying core's `Unwrapper` seam structurally.
The D9 gate measures the no-oracle line as a number: three sealed
messages across two epochs materialise at exactly 2 unwraps — one per
epoch — with the sealed mention body opening through the same op,
structure-only and foreign-owner negatives, and no path from any op to
a sealing seed [measured]. Named [O]s: rotation (old wraps make it
never replace-and-forget), batch unwrap, product wiring, the scope
re-render note for pre-capability accounts.

**The agent scope ships — one template, resolved per mint** ([episode
0137](0137-ecosystem-capability-minting.md); `specs/004-agent-scope`,
merged `e032687`, released in `v0.12.0`): the canonical agent capability template
exported beside the persona scope (`AgentScopePubAllow/SubAllow`,
`{{tag(topic)}}`/`{{tag(tool)}}`, notify by `{{name()}}` so
reachability cannot drift from attribution), and research 0126's rig
reborn as a standing e2e that measured the arc's two open server
behaviors: multi-value tag expansion (both tagged tools answer through
one credential) and the zero-matching-tag line drop (a tool-less mint
admits and reaches nothing) [measured]. No op, mint, or vault change —
D28 already stamps tags; the tag-policy watch now has its first
consumer.

**The BYON live run: two defects caught, then the provider arm
measures sound** ([episode 0136](0136-identity-the-byon-live-run.md),
identity `a0545c8`/`31279c6`, released in `v0.12.0`): the one bound episode 0135
left — the provider arm compile-proven, unmeasured — was paid in three
runs against the real Synadia Cloud BYON system. Run one drew 400: the
D47 amend patched `allowed_accounts` alone, and the JWT law refuses
accounts without users — fixed as a whole-object read-union-write
(`auth_users`/`xkey` carried forward, correct under merge or replace
patch semantics) with a userless AUTH refused by name. Run two held
the coupling live but exposed the probe outliving `t.Fatalf` and
re-creating its account mid-cleanup (now joined before any name is
freed), and taught that a BYON system's view carries no client URL —
the operator's fact, passed explicitly. Run three PASSED [measured]:
births 4.36s/3.29s against Bar 2's 5s bound, the scoped round trip
alive in the newborn account, AUTH carrying both tenants on read-back,
suspend/resume through the provider, the pre-existing account probed
12 times uninterrupted, zero leftovers. Every residue of the 0133
verdict is closed — both D47 arms measured, local and BYON.

**The platform-account topology measured sound, designed as D46–D49**
([episode 0133](0133-ecosystem-platform-account-topology.md)): the
operator's realms-are-accounts question ran as two-day research — all
four bars measured on live operator-mode servers, the reversal
condition never fired. The export seam (`account_token_position`
enforces at import-*definition* time — a tenant cannot even construct
a route to another tenant's surface) carries the D15 proof across
account boundaries by configuration alone; tenant birth is one fast
act (395–776µs) but as-built births an unusable tenant — no AUTH
`allowed_accounts` coupling, and a plain signing key that leaves every
minted user **admitted but inert** (0110's serves-immediately claim
refuted through admission); isolation holds through shared services
under two named disciplines (per-tenant connections; the connection
decides the tenant, never the payload); and the multi-tenant human
works today on the token lane while the OIDC ambiguity refusal stands
by operator decision. Per-tenant persona custody decided (each
tenant's persona keys on its own JetStream — a shared vault silently
shadows same-named personas across tenants, measured). Design:
[`platform-topology.md`](../02-DESIGN/soulstream-identity/platform-topology.md).
**D47 landed the same day** ([episode
0134](0134-identity-tenants-born-admissible.md), identity `447ec6b`,
released in `v0.12.0`): the tenant signing key is a scoped signer carrying the
canonical persona template (exported once from `client`, prefix-aware
— the ceremony adopts the same source on its next touch), and creation
amends AUTH `allowed_accounts` (idempotent, fail-closed between acts;
empty `AuthAccount` skips honestly). Store → **usable** admission
2.77ms with the out-of-scope publish refused [measured]. **Both
residues closed the same day** ([episode
0135](0135-ecosystem-the-residues-close.md)): the client mirrors the
op family and the e2e proves D47 through the real op family (create →
usable admission 11.2ms, `df8e4a3`); the provider arm carries both
halves (`f6a1a33` — the group's scope, the jwt_settings coupling; the
live BYON run followed the same day, episode 0136); and the house grew
tenants (soulstream spec 012): dir resolver persisting runtime
tenants, `SystemConn` wired, `soulstream account …` the hand — create
→ usable admission 8.8ms through the running house, surviving
restart [measured]. Remaining behind the focus gate: D46 export
configuration, D48 per-tenant persona custody, D49's disciplines with
the first shared-service build.

**Both graduated designs BUILT, two days after the asks** ([episodes
0120](0120-ecosystem-the-tools-arc-builds.md)/[0121](0121-ecosystem-the-approvals-loop-closes.md)):
the tools arc across four repos — core v0.12.0/v0.12.1 (the
`toolcatalog` discovery face, plus the correction the door forced:
reachability is not ceremony), identity v0.10.0 (`resources.*` runtime
custody, measured beating its own baseline: 0 failed accesses and a
5.7ms max gap under the probe, first ceremony 2.1ms after add, zero
restarts), workloads v0.7.0 (the lane's declared door environment),
soulstream v0.13.0-rc.9 (the forwarding door: one door, many targets,
per-call authority, refusals in words) — and the approvals loop closed
in identity v0.11.0 (durable tickets with witnessed expiry, the
public presentation/deny/status/pending/list ends, per-rule approvers)
with the house's guardrail turned on. What remains of both designs is
their shell halves, named [O]s in the design docs.

**The guardrail's human end measured to where it stops, then designed
shut** ([episode 0119](0119-ecosystem-guardrail-human-end.md)): D38's
loop is missing **two** ends, not one — the emit half works today and
the approval mints publicly as a plain D33 delegation, but the client
has no `approvals.present`, so a correct approval leaves the retry
still deferring [measured, 6/6]. The carrier question dissolved (the
plane can neither write the record nor push where anyone may listen —
it keeps readable state and the originating adapter carries the news),
a doc/code drift in D38's hash sketch was caught and corrected, and
the operator's ticket-lifecycle direction — durable tickets with their
own TTL, expiry as a recorded and notified outcome, async by
construction — became design
[`approvals.md`](../02-DESIGN/soulstream-identity/approvals.md)
(D42–D45). Builds behind the focus gate.

**External tools measured and designed** ([episode
0118](0118-ecosystem-agent-external-tools.md)): the research opened by
0116 graduated in two days with its riskiest claim measured on the
full composition — a wrapped agent's forwarding door, holding no
token, fetching authority per call, and **the remote seeing the
calling person** (6/6, zero cross-attributions, stolen delegations
refused, revocation biting one agent while the other kept serving);
Bar 2's custody scan over a real wrap wake found zero provider tokens
anywhere the agent reads (5/5). The graduated designs:
[`external-tools.md`](../02-DESIGN/soulstream-identity/external-tools.md)
(D39 the two-layer catalog, D40 `resources.*`, D41 the door
invariants) and the record's
[`tool-catalog.md`](../02-DESIGN/soulstream-core/extensions/tool-catalog.md)
discovery convention — the operator's own questions reshaped the
catalog twice (the run-your-own/remote split; one discovery face over
two custodies). Builds behind the 0071 focus gate.

**Built, and not yet reachable by a person** ([episode
0116](0116-ecosystem-what-shipped-without-a-human-end.md)): the
evaluation found both of the plane's newest capabilities complete and
unwired at the product. The grants broker is off in the house — the
`grants.*` ops enable on a non-empty `GrantResources` and `soulstream`
declares none — and the guardrail is default-off with no read op for
its standing rules, no feed of its decisions, and no home for D38's
middle clause (a human's yes). Neither is a defect in the plane; both
are missing ends. Two research topics carried them; both graduated
2026-08-21 (episodes 0118 and 0119).

**The provider arm closes — v0.9.0** ([episode
0114](0114-identity-the-provider-arm-closes.md)): D35's second
authority backend measured live on Synadia Cloud — an account born at
runtime in ~51s (the local arm's 1.69ms, four orders of magnitude
apart: propagation, not our path), a principal admitted through the
real cloud, suspend/resume landing, the pre-existing account
uninterrupted across 22 probes. The tenancy topic has **no open
residue left**. Two defects the run forced: the control plane 5xxs on
a just-created account (retry discipline added) and the SDK carries
auth on the context.

**The tenancy set is BUILT — v0.4.0/v0.5.0/v0.6.0 in one overnight arc**
([episode 0110](0110-identity-the-tenancy-set-builds.md)): the general
secret store on the extracted sealed-CAS pattern (D36), the guardrail
evaluator at the op-path chokepoint with fail-closed discipline and
approvals as one-shot subject-signed delegations (D37/D38), and runtime
tenancy — an account born, suspended, resumed on a live dir-resolver in
1.69ms store→admitted, the new tenant mintable the moment the op
returns (D35). Residues named: the ProviderAPI authority arm (BYON, an
operator act), rate counters and approver policy by demonstrated need.

**Tenancy and guardrails designed — D35–D38** ([episode
0107](0107-ecosystem-platform-tenancy-guardrails.md);
[`tenancy.md`](../02-DESIGN/soulstream-identity/tenancy.md)): the
`accounts.*` lifecycle op family with pluggable authority (local
operator key | provider API — runtime birth measured at 543–774µs,
zero restarts, no usable half-account across ~2,800 probes/run), the
general secret store as the third custody domain, the guardrail
evaluator (CEL; allow-path p99 206–220µs at 100 rules; the mandated
belt-and-braces discipline after the 622ms cost-limit scare)
unskippable at capability chokepoints, and approvals as one-shot D33
delegations. All eleven open decisions taken by the operator
2026-08-18; builds behind the 0071 focus gate; the BYON provider arm
of Bar 1 the one named residue.

**The grants broker lands — v0.3.0** ([episode
0105](0105-identity-the-grants-broker-lands.md); realizes D30–D33,
`specs/003-grants-broker/`): the overnight slice reviewed sound and
completed in the morning — the `grants.*` op family on the
principal-scoped surface, custody in its own sealed CAS bucket, derived
short-lived tokens the only return, subject-signed delegations with the
not-before check the review added. The repo's own gate now carries the
transport clause in consumer position [measured]: the imposter's publish
dies at the server and the delivery log shows the victim's subject
served exactly twice, by her own two calls. The CLI grew the ceremony
(`grant link|access|ls|revoke`); the runbook find of record: GitHub
answers form-encoded without `Accept: application/json` — caught
writing SC-005's runbook, before it bit. **SC-005 closed the same
evening** [measured]: the operator's real-GitHub walk rotated the line
live and refused after revoke — 0104's Bar 2 residue closes (run
recorded in the feature's quickstart). Still open, named: the
wrapper's overlay seam and the fold's exchange grant.

**Outbound grants designed — every remote call carries the calling
user** ([episode 0104](0104-ecosystem-outbound-identity-grants.md); D30–D34 in
[`grants.md`](../02-DESIGN/soulstream-identity/grants.md)): the
overnight research topic `outbound-identity-grants` measured all three
pre-registered bars on real components and graduated. The inversion:
no outbound credential ever lives in agent, workload, or MCP-client
config — a `grants.*` op family on the principal-scoped surface
custodies per-persona OAuth grants in a **second sealed custody
domain** (the key vault's records are immutable by design and refuse
rotation — measured, not assumed), returns only derived short-lived
access tokens (Article I's line, D32), and honors on-behalf-of only
with a minted, bounded, **subject-signed** delegation whose actor is
the server-proven principal — a stolen delegation refuses as an actor
mismatch [measured]. Bar 1 ran the fold-as-AS composition multi-user
through the real node: two passkey personas, zero cross-attributions,
dead bearer refused in 12.6ms [measured]. Cross-component demands now
on the record: the fold's token-lifetime knob (revocation propagates
in `exp` + callout TTL) and its RFC 8693 exchange gap; the wrapper's
per-run overlay seam returns to carry delegations. Build gated —
see the roadmap's M6.

**The default IdP is a sibling — soulfold, the refusal holds**
([episode 0019](0039-soulidentity-soulfold-the-default-idp.md)): the operator's
default-IdP question — should SoulIdentity, bearing the name, also be
the passkey-first OIDC provider deployments get out of the box — was
answered by holding the vision's refusal: identity truth stays in the
deployment's IAM, and the default IAM becomes **soulfold**
(`github.com/impire-io/soulfold`), a NATS-native (JetStream KV),
embeddable, passkey-first OIDC provider the callout issuer treats
identically to Entra — the D23 seam only, no side-channel, no shared
store, no precedence [mechanism-argument]. "Default" is distribution
wiring (`--oidc-issuer`), replaceable by any OIDC provider by config.
Pocket-id itself was ruled out as the bundle not for being Node (it is
Go since v1) but for being an application with an SQL store, not an
embeddable library. Named, not built: D23 multi-issuer dispatch for
deployments running soulfold beside a second issuer.

**The embed seam — D29, M2's second consumer-proven addition** ([episode
0018](0036-soulidentity-the-embed-seam.md); D29 in
[`../02-DESIGN/agent.md`](../02-DESIGN/soulstream-identity/agent.md), feature
`specs/002-embed-seam/`): soulnode's composition research measured the
wall — provisioning is public, the serve assembly was `cmdServe`-only,
and two consumers had already ridden the module-namespace dodge to reach
`internal/`. The public `embed` package now exposes the one seam:
`Run(ctx, Options)`, value-only options, custody unchanged (D13),
provisioning still wire-only; the daemon is its first consumer (one
assembly, two entrypoints). The proof is compiler-grade:
`e2e/embedgate/`'s module path sits outside the repo namespace, so an
`internal/` import cannot compile; the gate runs the M4 admission shape
through `embed.Run` [measured]. A drain lesson strengthened the contract:
`Run` flush-confirms its unsubscribes, so returned means silent. Existing
gates unchanged, uncached-green [measured].

**Role selection by declared name — D28, the first consumer-proven
M2 addition** ([episode 0017](0035-soulidentity-role-selection-by-name.md); D28 in
[`../02-DESIGN/agent.md`](../02-DESIGN/soulstream-identity/agent.md), answering
soulidentity#1): soulrealm's fleet needs one scoped signing key per role
on one realm account — the exact observable D5's amendment reversal
condition watched. The `mint.ephemeral` op issues an ephemeral scoped
user JWT against a **named role** for a **caller-supplied public key**
(no seed in either direction; the response is the JWT alone), with
**tags** in the user claims for scoped templates to resolve and a
required TTL (D22's bound). The nouns, corrected by the operator the
same day: a **team is the account, the tenant** (0013's "teams are
accounts"); the declared signing key is a **role** — wire field, client,
CLI, and internals renamed before any consumer wired in. Proven in the
M3-gate e2e [measured]: the by-name-minted JWT admits to the
operator-mode realm; with a second role imported the binding path refuses
as ambiguous while both roles stay reachable by name. Named, not built:
the token lane's own named-role answer (node enrollment), and per-role
tag policy.

**One noun: persona** ([episode 0016](0032-soulidentity-one-noun-persona.md); D27 in
[`../02-DESIGN/agent.md`](../02-DESIGN/soulstream-identity/agent.md), constitution 1.3.1):
persona == identity — the ecosystem's one noun for the represented
subject, adopted from soulstream's fixed terminology; *principal* is the
server-proven (account, user) a connection speaks as (D15's term), and
"identity" survives only in the product name. The persona is born at
first encounter (D26); the vault is where its durable artifacts live. A
vocabulary pass over vision, constitution, missions, and load-bearing
comments — no type, op, or JSON field changed [measured: gate green].

**The vault is the directory — ephemeral users, keys on first touch**
([episode 0015](0031-soulidentity-the-vault-is-the-directory.md); D26 in
[`../02-DESIGN/nats-surface.md`](../02-DESIGN/soulstream-identity/nats-surface.md)): episode
0014's persona-directory trust path was refuted by the operator the same
day — identity truth lives in the IAM, users are ephemeral, and no
per-user act may exist anywhere. A persona key is a capability artifact,
not identity (an access token carries no user key and cannot sign
records): the caller's own key **materializes inside the vault on first
touch**, owner-bound, and `keys.public` is the **open directory read** —
readers build verification keyrings from the identity plane; no profile
store. The sealed-communication key follows the same pattern when D9
lands.

**M2's first gate criterion is measured — the seam carries a real record**
([episode 0014](0030-soulidentity-the-cross-service-proof.md), re-proven on the D26
shape in 0015): a Soulstream record signed through the running
SoulIdentity service verifies in a real realm. The proof lives in `e2e/`
— a consumer-position module importing both repos (the cycle guard's
shape), riding `make test`, now with **zero per-user acts**: one team
key, one minted credential, the key materializing at signer construction,
the reader's keyring one `keys.public` answer; announce, baseline, and
turn read `SigVerified` — `unknown-key` without the keyring, the negative
control [measured]. What remains of M2 is the **node half** (one pooled
connection per user, no node-held creds), in soulstream's remote MCP node
feature — which publishes nothing per user; it builds keyrings from the
identity plane.

**The identity registry is dissolved — authorization lives in the ACLs
and the bindings** ([episode 0013](0029-soulidentity-the-registry-dissolves.md); D25 in
[`../02-DESIGN/nats-surface.md`](../02-DESIGN/soulstream-identity/nats-surface.md), amending
D2/D5/D6/D18/D22/D24): the operator's one-identity-one-persona question
unwound the ledger field by field. Persona keys carry their owner
(account, user) at import and `sign.record`/`keys.public` check the
caller against the binding; every mint resolves its signing key by the
target account's D24 team binding (ambiguity refuses); the management ops
are gated by the server's own permission enforcement on the op tail —
`requireAdmin`, the `admin` flag, `identities.*`, and `internal/registry`
are deleted; mint is an operator op (self-mint died with the row that
authorized it). The token store is the one registry standing. All three
e2e gates re-proven on the new shape [measured], including the new
op-tail proof (a represented user publishing a management op on its own
prefix: server-refused, zero service decisions) and the revocation bound
(re-admitted 5.25 s after connect at a 5 s TTL). The client now carries
the M2 seam surface: `PersonaSigner` satisfies soulstream's
`identity.Signer` structurally, exercised in the M3 rig. Next: **M2 —
consumers wire in** (the cross-repo gate proof, then the remote MCP
node).

**The Entra/OIDC lane is open — role == team, no mapping store**
([episode 0012](0025-soulidentity-entra-role-claim-lane.md); D23–D24 in
[`../02-DESIGN/auth-callout.md`](../02-DESIGN/soulstream-identity/auth-callout.md), built as
spec-driven feature `specs/001-entra-oidc-backend/` through the newly
enabled speckit flow): an external client presents an Entra access token
instead of an API token; the issuer validates it against one pinned
issuer/audience via JWKS discovery (fail-closed, key rollover without
restart) and authorizes by the `roles` claim — the role value IS the team
name, resolving directly against the vault's account signing keys, which
now carry their account binding at import. D22's sketched rule table was
refuted before it was built; no catalog, no per-user entries; admin and
personas never come from claims. Gate met on the stub rig [measured]:
sealed-leg admission with full attribution, the nine-row refusal matrix,
`sit_` lane untouched, and the accepted revocation asymmetry demonstrated
(token lifetime + one TTL; cached token re-admitted 5.2 s after connect
at a 5 s TTL). Real-tenant verification is a documented manual runbook.
Next in the execution order remains **M2 — consumers wire in**.

**The subject space gained its ecosystem namespace**
([episode 0011](0023-soulidentity-the-shared-subject-prefix.md), D14 amended at the
operator's direction): the root is `<prefix>.soulidentity` with the prefix
shared across all soulstream components (`--prefix` /
`SOULSTREAM_PREFIX`), empty by default. Environments coexist in one realm,
and the account token sits at declared position `P+2`, so a cross-account
export (`account_token_position`) extends D15's principal proof by
configuration alone. The M3 gate e2e now runs fully prefixed
(`prod.soulstream.soulidentity.>`) [measured]; the honest cost — prefix
mismatch is silent timeouts — is mitigated by the startup root log and the
shared environment variable.

**Milestone 4 — auth callout, the front door — is shipped**
([episode 0010](0022-soulidentity-m4-auth-callout-ships.md); design
[`../02-DESIGN/auth-callout.md`](../02-DESIGN/soulstream-identity/auth-callout.md), D19–D22,
researched in [episode 0008](0019-soulidentity-sentinel-credential-flow.md) and
[episode 0009](0021-soulidentity-claims-mapping-shape.md)): an external-identity client
holds a sentinel creds file (minted by the admin-gated `sentinel.mint` op,
public by design) and its API token; the issuer — one process, two
connections, the callout subscription in the dedicated AUTH account —
digest-validates against the token store (records name an identity, carry
no policy), authorizes via the registry row, and mints a TTL-bounded
scoped JWT for the server-assigned key with the vault's role keys. Gate
met [measured]: admission with server-enforced scope and full audit
attribution; bypass-lane connections produce zero callout decisions;
invalid and revoked tokens refused; the D21 xkey leg proven (sealed
requests and responses). Token management is four admin-gated surface ops.
Entra/OIDC arrives later as validator configuration on the same D22
interface. Open: `ngs-capabilities` (blocked on operator access to the
Synadia account; gates only the NGS deployment class), and next in the
execution order, **M2 — consumers wire in**.

**Milestone 3 — the NATS-native rebuild — is shipped**
([episode 0007](0018-soulidentity-m3-the-nats-native-rebuild.md); design
[`../02-DESIGN/nats-surface.md`](../02-DESIGN/soulstream-identity/nats-surface.md), D14–D18):
the service answers over sealed NATS request/reply on the caller's own
subject prefix — the principal proven by the server's publish-permission
enforcement (D15, via the scoped-key template
`soulidentity.{{account-subject()}}.{{name()}}.>`) — and the vault seals
into NATS KV with both xkeys deployment-supplied. Act-as (D6) is enforced
and audited against real principals; management is admin-gated in the
registry (D18, the socket trust model's successor). All four gate criteria
measured in the e2e proof: unauthorized act-as refused and logged, wire and
store ciphertext-only (positive-control-verified), cross-prefix requests
refused by the server itself. The socket agent, `NATSOption`, file
keystore, and `sign/nonce` are deleted. Next in the execution order: M4
(auth callout, the front door), then M2 (consumers wire in). Design and
review that preceded the build:
[episode 0005](0016-soulidentity-the-nats-surface-design.md),
[episode 0006](0017-soulidentity-design-review-amendments.md).

**The first-key story is decided — M3's research gate is open**
([episode 0004](0015-soulidentity-first-key-story.md), D13): the unwrapping xkey for the
KV backend's envelope encryption is a local root secret on the service host
— decided as a `0600` file, then amended at design review to
deployment-supplied environment configuration
([episode 0006](0017-soulidentity-design-review-amendments.md)) — named honestly as
plaintext in the same trust class as the creds file, with the envelope's
real gain being that broker disks,
replicas, and backups never hold plaintext seeds. All three pre-registered
bars passed [measured]: the sealed round-trip survives broker+service
restart unattended, the store holds ciphertext only
(positive-control-verified), and the from-nothing bootstrap is two operator
acts plus one automatic service act.

**The mission was re-centered, twice in one day**
([episode 0002](0013-soulidentity-the-identity-plane-re-centering.md) then
[episode 0003](0014-soulidentity-nats-only-and-the-connection-ladder.md); constitution
1.1.0 → 1.2.0): SoulIdentity is the identity plane of the Soulstream
ecosystem — the representation of identity for humans and agents, delivered
as a **NATS-only** service with xkey-sealed E2E request/reply (D11, D12).
There is no socket: a presented creds file bypasses SoulIdentity entirely
(self-custody, server-verified natively), everything else arrives through
auth callout, authorized by the declared registry or by validated claims in
the presented token (D2). NATS KV with xkey envelope encryption is the
vault's initial backend. The milestone-1 socket surface, `NATSOption` seam,
and file keystore are transitional until the NATS-native rebuild (M3).

**Milestone 1 — the walking skeleton — is shipped**
([episode 0001](0012-soulidentity-genesis-and-the-walking-skeleton.md)): vault, declared
identities, the agent on a Unix socket, scoped minting, and the `client`
package with `NATSOption`, proven end to end against an operator-mode NATS
server — mint through the agent, nonce signed in the vault, scope enforced by
the server, no seed ever in the client process [measured]. The design is
twelve numbered decisions in [`../02-DESIGN/agent.md`](../02-DESIGN/soulstream-identity/agent.md);
the plan is [`../03-IMPLEMENTATION/ROADMAP.md`](../03-IMPLEMENTATION/ROADMAP.md)
(execution order M3 → M4 → M2: the NATS-native rebuild, then auth callout as
the front door, then consumers wire in over the NATS surface). The first
release tag arrived 2026-08-02 — `v0.1.0`, for soulfold's pin; open
questions before their milestones (NGS callout
capabilities, the sentinel-credential flow, the first-key story, the
claims-mapping shape, service round-trip latency) are named on the roadmap.

### soulstream — the product, the house (as of 2026-08-27; named soulnode until episode 0069)

**v0.14.0-rc.2 — the 08-27 arc becomes installable the evening it
landed** ([episode
0139](0139-soulstream-the-rc2-carries-the-tenants-and-the-capabilities.md)):
the pins move from the day's pseudo-versions to named tags (`8f72f66`,
the idp `v0.8.1` patch taken along) — identity `v0.12.0` (D47 + agent
scope + sealing custodian), workloads `v0.8.0-rc.2` (capability
minting, gate proven `GOWORK=off` against the pinned core), core
unchanged at rc.1 — carrying specs 012 and 013. Verified from the
outside [measured]: prerelease with all four tarballs, the tap at
`0.14.0-rc.2`, the darwin_arm64 binary printing its version and
answering `soulstream account`. The declaration story is now
brew-installable end to end; that an agent still runs only where a
person runs wrap became the next focus the same evening ([episode
0140](0140-ecosystem-the-focus-agents-as-infrastructure.md)).

**Capability minting in the house — spec 013, released in
v0.14.0-rc.2 (episode 0139)**
([episode 0137](0137-ecosystem-capability-minting.md), merged
`baf5a7e`): the founding grows the **agent capability key** — a scoped
signer on the realm account rendering `client.AgentScope*`, beside the
plain workload key — and the runtime plane routes capability-bearing
declarations through the scoped lane on the state-held seed. Measured
(TestM14): the granted scope-probe completes `done`; granting a
different tool leaves the probe's own subject server-denied and the
run abandons — the narrowing bites through the full authority chain
[measured]. Legacy realms and BYON realms without the agent scope
refuse capability declarations by name before any op publishes
(pre-v1 clean break); the identity vault gains **no** second role, so
the token lane is untouched (the RoleForAccount finding, spec 013).
The D47 persona-scope adoption rode along: the ceremony renders
`client.PersonaScope*` from the one exported source. The standing
exception (workloads/identity pins bump after their mains push, the
0089 precedent) closed inside the rc.2 cut (episode 0139).

**The house grew tenants — spec 012, released in v0.14.0-rc.2
(episode 0139)** ([episode
0135](0135-ecosystem-the-residues-close.md)): the platform topology's
product half, taken as a full cycle at the operator's direction. The
embedded server's resolver became a **dir resolver** under
`<state>/resolver` (runtime tenants persist; seeded create-if-absent
so AUTH's runtime amendments survive restart — and the dir resolver
demanded the operator's claims, synthesized in memory from the
ceremony seed), the identity plane's `SystemConn` is wired (SYS user
minted in memory — no new artifact, day-2 realms gain tenancy on
their next `up` via the operator-key ensure), and `soulstream account
create|list|show|suspend|resume` is the hand. Gate measured: create →
usable token-lane admission **8.8ms** through the running house, the
tenant admitting after stop-and-start, M1.1 green on the dir
resolver. BYO stays honestly off (design 0003 — no operator material
on that side). Design 0001 §3 amended in the same change.

**v0.14.0-rc.1 — the rc carries both builds, brew-installable**
([episode 0132](0132-soulstream-the-rc-carries-both-builds.md)): pins
to core `v0.14.0-rc.1` (system stream + sealed topics) and workloads
`v0.8.0-rc.1` (the agent declaration), the wrap verb grows
`--declaration` (the same thin-main wiring as `soulstream-wrap`), and
`init`/`up` provisions `SOULSTREAM_SYSTEM` for free. The tap formula
moved with the release — `brew install impire-io/tap/soulstream`
serves the rc; the darwin_arm64 tarball's binary verified from the
outside [measured]. One defect died on the way: the spec-011 rig's
fixed-port bind (8378 race under package parallelism) — the flake
fixed on main with the foldplane pattern. The byon soak stays on the
v0.13.0-rc line; adopting the rc there is the operator's own act.

**The wrap announces itself and lights its lamp — branch
`011-presence-lease`** ([episode
0125](0125-ecosystem-the-presence-lease-builds.md); spec
`specs/011-presence-lease/`, the first spec-kit feature since 010):
before answering mentions the wrap ensures a directory floor
(lookup-first, because `registry.Publish` replaces metadata; no
signing key, because the lane holds none — said honestly) and holds
the presence lease beside the run loop, the farewell waited for
before the connection closes. Advisory throughout: a lease failure is
a log line, never a refusal to answer. The live rig measures the
whole story through real admission — the persona scope's existing
tails let a wrap create and write the presence bucket with no
ceremony change [measured]. Merged to main 2026-08-25 (`b86ec57`), and
the house now composes shell v0.11.0-rc.3 — the first hour's screens
(episode [0127](0127-shell-the-first-hour-builds.md), pin `68acda6`);
the live run on byon is the quickstart's pending human act.

**byon runs rc.8, adopted rather than re-founded** ([episode
0115](0115-soulstream-byon-adopts-the-new-form.md)): the canonical
break's refusal got precise — `soulstream adopt` reads the realm's own
op-log, adopts an empty one and refuses a populated one by count. The
live realm migrated with its passkeys, accounts, and vault intact, and
a signed turn on it now reads wire-v2, acting-named, verified
[measured]. F1 closed itself in production the moment rc.8 started: the
archivist's profile appeared in a personas bucket empty for three days.

**v0.13.0-rc.2 — first contact hardened, the tap feeds itself**
([episode 0100](0100-soulstream-v0-13-0-rc-2.md)): the six
first-contact fixes released; goreleaser's first pipeline-driven
formula push landed in the tap (the HOMEBREW_TAP_TOKEN act closed) and
`brew upgrade` moved a real machine rc.1 → rc.2 [measured].

**First contact: the BYON founding** ([episode
0099](0099-soulstream-the-byon-founding.md)): realm `byon` founded
live on the Impire DEV Synadia Cloud BYON — callout admitting on
Synadia's infrastructure, garbage refused, every plane serving, the
custody audit clean, the pre-existing wiring byte-identical. First
contact forced six fixes (client-dir collision, system-by-id,
awaiting-state resume, seeds-persist-immediately, lossy-channel
retries, read-first callout ops) — all on main with replaying tests,
awaiting v0.13.0-rc.2 behind the `HOMEBREW_TAP_TOKEN` secret. The
unsealed-callout caveat is now measured (`sealed_requests=false`);
the private-link idle-watchdog bug is Synadia's, evidence handed over.

**brew install soulstream** ([episode
0098](0098-soulstream-the-tap-opens.md)):
`brew install impire-io/tap/soulstream` installs the current candidate
on macOS and Linux — the tap repo is live with the v0.13.0-rc.1
formula (installed and `brew test`ed on a clean PATH [measured]), and
goreleaser owns the formula from the next tag. Formula over cask
(Linuxbrew matters; the goreleaser deprecation is watched). The
`HOMEBREW_TAP_TOKEN` act closed at rc.2 (episode 0100): the pipeline
now feeds the tap itself.

**v0.13.0-rc.1 — bring your own server** ([episode
0097](0097-soulstream-v0-13-0-rc-1.md)): the BYO arc released — same
pins as v0.12.0-rc.1, one new capability: founding on a server
soulstream does not run, both flavours. Artifact round-tripped and
verified. Known-open on the candidate: the live Synadia Cloud runbook
and the platform-xkey sealing caveat.

**BYO NATS — designed and shipped the same day** ([episodes
0095](0095-soulstream-byo-nats-designed.md)/[0096](0096-soulstream-byo-nats-ships.md);
design [0003](../02-DESIGN/soulstream/0003-byo-nats.md) resolving
composition 0001 §4's [O]; soulstream `specs/010`): two flavours behind
`byo.flavour` — the self-hosted kit (exact `nsc` commands and config
fragments, applied by the operator, verified behaviourally with
refusals naming the unapplied item, closed by a callout smoke round)
and Synadia Cloud BYON (the account half driven through the
control-plane API — byon-setup graduated, idempotent, lost-seed
refused). Operator mode required; conf-auth and NGS shared refused by
name — `ngs-capabilities` closed unopened. The custody rule holds by
test: no operator or account master key ever travels; self-hosted, no
seed crosses the boundary in either direction. Measured end to end
against a config-file operator-mode rig playing the operator; the live
Synadia founding is the quickstart's manual runbook.

**v0.12.0-rc.1 — the clean-break candidate** ([episode
0094](0094-soulstream-v0-12-0-rc-1.md), pinning core v0.8.4 /
workloads v0.4.0 / shell v0.6.0 / idp v0.5.0): everything episodes
0089–0093 landed, released — one binary and one paste for agents, one
canon on every page, one console for administration, one vocabulary
with no fallbacks. A realm founded on an earlier candidate is refused
with the hand-migration named; artifact round-tripped and verified.

**One console, one vocabulary** ([episodes
0091](0091-ecosystem-the-shell-is-the-console.md)/[0092](0092-ecosystem-the-names-say-what-they-do.md);
`specs/009-one-console-one-vocabulary/`): the planes are named by
function — `planes.signin`, `planes.mcp`, `--signin-listen`,
`--mcp-listen` — and pre-v1 renames are clean breaks (episode 0093):
no aliases, no fallbacks; a byname-era realm is refused by name with
the hand-migration in the refusal; the
bundled sign-in plane serves its admin API and not its console
(`/admin` → 404, pinned); `up` prints functional labels; docs and the
site mirror the real output. Pins idp v0.5.0 + shell v0.6.0.

**Wrap in the house — one binary, one paste** ([episode
0089](0089-ecosystem-wrap-in-the-house.md);
`specs/008-wrap-in-the-house/`, design
[0002](../02-DESIGN/soulstream/0002-wrap-in-the-house.md)): the product
binary answers `soulstream wrap` and `soulstream mcp` natively over the
libraries it already pins — no Go toolchain, no PATH assembly on an
agent's machine — and the getting-started teaches download-and-paste
with `go install` gone. Proven live end to end: the Agents screen's
paste block, unedited, under fish, answered a mention posted while the
wrapper was off (one reply, real claude, 6.7 s); a revoked credential
refused loudly. Pins shell v0.5.0 + workloads v0.4.0 (go.sum resolves
once the tags are pushed).

**v0.11.0-rc.2 — wrap ships** ([episode
0086](0086-soulstream-v0-11-0-rc-2.md)): the second candidate pins core
v0.8.4, workloads v0.3.0 and shell v0.4.3 — a person on this RC creates
an agent in the shell, pastes one block, and either talks through their
assistant (the MCP door) or lets it answer mentions for them
(`soulstream wrap`), from their own machine with their own logins.

**The helm plane — the cockpit joins the bundle** ([episode
0068](0068-soulnode-the-helm-plane.md); `specs/007-the-helm-plane/`):
`planes.helm` runs soulhelm v0.1.0 through its public embed seam — on
by default beside the fold, absent-block-means-disabled on old state
dirs, `helm console` the fourth URL at `up`. Decision of record:
`SessionIssuer()` — enabling the helm switches the identity plane's
OIDC admission lane on in local mode with the bundled fold as AS (an
explicit external AS wins; ephemeral listeners resolve to no issuer
and the helm skips loudly). Named, not built: a dedicated scoped helm
ceremony user (the plane hands the ops lane today).

**The front of house — URLs, the console, no collision**
([episode 0012](0062-soulnode-the-front-of-house.md); `v0.3.0`): `up`
now logs the MCP door, the fold sign-in, and the admin console URLs;
the fold is on by default (`init && up` lands a person at a passkey
prompt with an admin console); the door and fold can't share a
listener (refused by name; public mode needs two fronted routes); and
the bare-IP RP-id footgun is caught at load (issuer defaults to
localhost). Two bugs the live browser run surfaced were fixed
upstream and pinned: the invite link dead-ended (soulfold grew a
standalone `/enroll`, v0.3.1) and a phantom duplicate user (CreateUser
made index-first, v0.3.2). Verified the whole human path in a real
Chromium: click the enrol link → create a passkey → land in the
console.

**v0.1.0 released — the house gets a shipping label**
([episode 0009](0058-soulnode-the-release-pipeline.md)): CI (first
runner run green, folded-realm gate included) and the tag-triggered
goreleaser release, archivist pattern; `v0.1.0` published with
linux/darwin binaries, an artifact round-tripped, checksum-verified,
answering `soulnode version` → `0.1.0`. No windows on purpose (the
ceremony's owner-only modes). The operator made soulrealm public —
the whole consumed stack now fetches with zero credentials, and the
"private-module credential" blocker is gone. **Next:** day-2 items,
soulfold M3 upstream, or Phase 3's tsnet gate.

**The folded realm — one binary, real people**
([episode 0008](0057-soulnode-the-folded-realm.md);
`specs/006-the-fold-plane/`): `planes.fold` runs the bundled
passkey-first OIDC provider (soulfold by tag, public embed seam) on
the node's own JetStream, and public door mode defaults its AS at it.
Measured end to end with zero external services: DCR → passkey
sign-in (first touch enrolls) → bearer MCP session at the door →
`whoami` names the fold identity; founding token coexists; old state
dirs load unchanged. Startup order recorded as load-bearing (fold
before identity plane). Two upstream tags rode along (soulfold v0.1.1,
v0.1.2). **Next:** release/distribution polish, the fold's M3
upstream, or Phase 3's tsnet gate.

**The public door opens — Phase 2 completes**
([episode 0007](0055-soulnode-the-public-door.md);
`specs/005-public-door/`): `planes.door` grew
`public_url`/`auth_issuer`/`auth_audience` additively and the node now
serves the whole hosted-client OAuth story — a scripted client knowing
only the door URL walks challenge → resource metadata → DCR →
code+PKCE → token, and the bearer forms an MCP session whose realm
identity is the token's subject (lane=oidc, role=realm in the audit).
Founding token coexists; undeclared roles and garbage refuse; the
listener stays loopback with HTTPS as deployment fronting. Measured
against the upstream AS contract's stand-in — AS-agnostic, soulfold
the intended default (its own admission proof is episode 0054). Bonus
root-cause fix: listen port 0 no longer collides with nats-server's
default-port reading. **Next:** the bundled-fold wiring (soulfold M5),
day-2, or Phase 3's tsnet gate when fronting measures insufficient.

**The door is open — Phase 2's local mode is done** ([episode
0006](0049-soulnode-the-door-opens.md); `specs/004-the-front-door/`): upstream's
remote MCP node (soulstream 018, v0.7.0; made consumable in its journey
0010) runs as the door plane — `planes.door`, loopback HTTP, static
bearers through the realm's own callout, custodying nothing. Measured: a
real MCP client with the founding token forms a session, lists the tool
surface, and `whoami` names the realm-admitted owner; garbage bearers
refuse; the state dir is untouched by the door. The vision sentence
executes: `init && up`, paste the token, work. Public mode waits on
soulfold upstream. The pseudo-version exception closed 2026-08-02: all
four upstreams tagged (soulrealm v0.1.0, archivist v0.2.0,
soulstream/node v0.7.0, soulidentity v0.1.0) and soulnode pins them by
tag. **Next:** public door mode when soulfold lands, or Phase 3's tsnet
gate when Phase 2 fronting is measured insufficient.

**Phase 1 is complete — an agent runs** ([episode
0005](0045-soulnode-an-agent-runs.md); `specs/003-an-agent-runs/`): `soulnode
workload start` is the invocation-scoped runtime plane — upstream's own
`agent-echo`, declared unchanged, launches with a minted TTL-bounded
credential, posts as its persona, and its lifecycle lands as a completed
work item owned by `runner`; nothing credential-shaped lingers after end
of life. The ceremony carries the two-keys split (plain workload minting
key beside the scoped admission key; inventory 20). The composition
caught a consumer-proven upstream bug on its first enforcing run —
agents lacked `$JS.API.INFO` for the realm client's availability probe —
fixed upstream first (soulrealm `3fee11f`), pin bumped. Three
pseudo-version pins await upstream tags. **Next:** Phase 2 — the front
door, gated on soulstream's `018-remote-mcp-node` (in flight upstream).

**M1.2 is done — the realm remembers** ([episode
0004](0044-soulnode-the-realm-remembers.md); `specs/002-realm-joins/`): `init`
founds the record substrate; the memory plane runs in-process (public
`keeper`/`archive` on a realm client signing through the identity plane —
persona key vault-held, nothing on disk). Measured: the owner's full path
through admission (post a turn → archivist keeps it, author `owner` →
memory answers with attribution and a citation) rides `make test` in
~5 s; restart continuity exactly-once; the disabled-plane arm clean.
`config.json` carries `realm` (default `home`) and the first plane block.
**Next:** M1.3 — an agent runs (the runtime plane; carries the known
workload-signing-key design question).

**M1.1 is done — first boot is real** ([episode
0003](0042-soulnode-first-boot-is-real.md); `specs/001-init-and-up/`): `soulnode
init && soulnode up` founds and runs a realm — the whole ceremony
persisted into one state directory (17 artifacts, owner-only modes), the
founding acts through public surfaces, the first token printed once, the
embedded operator-mode server + identity plane on ordinary loopback
connections. Measured: init in 0.15 s; the found→admit→refuse→revoke→
restart e2e rides `make test` in ~1 s; re-init is a verified no-op. One
refuted assumption on record (refuse-on-0755 became tighten-then-verify).
Dependency exception tracked: soulidentity pinned at a pseudo-version
until it tags. **Next:** M1.2 — the realm joins (provisioning + the
archivist plane).

**The composition gate is met — Phase 1 is unblocked** ([episode
0002](0040-soulnode-the-composition-gate.md)): the `single-binary-composition` topic
measured all three pre-registered bars PASS — embedded admission parity
(4 ms token-lane admission, server-asserted principal, refusals audited),
the embed surface per component (three of four upstream asks *delivered*
before graduation: soulidentity's public `embed.Run`, the archivist's
public `keeper`/`archive`, soulrealm pinned to tagged soulstream v0.6.0;
the fourth — soulstream's public MCP surface — held for the maintainer's
open `remote-mcp-node` topic, gating only Phase 2), and the first-boot
ceremony fully generated by code from an empty directory. The maintainer's
transport call superseded the proposed split: **everything through
loopback** — decomposition is configuration, never architecture. The
constitution ratified at **1.0.0** with Article III reworded accordingly.
Design [`0001-soulnode-composition.md`](../02-DESIGN/soulstream/0001-soulnode-composition.md)
is open. **Next:** the spec-kit pass for M1.1 (`soulnode init` + the
server and identity plane).

**The project was founded** ([episode 0001](0034-soulnode-genesis.md)): SoulNode is
the single-binary distribution of the Soulstream ecosystem — embedded NATS,
SoulIdentity, the archivist, soulrealm, and an MCP front door in one process
on a machine the user owns, with the first boot (`soulnode init`) treated as
the product. The hq is bootstrapped from soulrealm's proven structure. The
constitution's load-bearing articles are **composition, not invention**
(SoulNode wires public tagged surfaces; domain logic lands upstream) and
**same shape as any deployment** (operator-mode NATS + auth-callout locally,
no dev fork). Feasibility entered measured; so did the obstacles
(SoulIdentity's serve path was `internal/`, soulrealm's `replace`
directive) — both since dissolved by the upstream landings above.

### soulstream-idp — the fold (as of 2026-08-18; named soulfold until episode 0069)

**The token-lifetime knob** ([episode
0106](0106-idp-the-token-lifetime-knob.md); **v0.7.0**): the first of
0104's two due demands closed — the access-token lifetime is a
deployment knob through the whole seam
(`embed.Options.AccessTokenLifetime`, `--access-token-lifetime`), zero
keeping the default hour; the e2e proves both arms (a 5-minute fold
issues 5-minute tokens through the full ceremony) [measured]. Named,
not built: per-client lifetimes, the product-side plumbing, and the
second due demand — the RFC 8693 exchange grant, spec first.

**The refresh grant lands** ([episode
0103](0103-ecosystem-the-session-outlives-its-token.md)): refresh
tokens leave "not in M1's scope" — `offline_access` mints a rotating
one beside the access token, single-use by the codes' own CAS flip on
a digest-keyed record (D12), 30 days idle bounds any one token,
revocation deletes it, and the M1 gate walks the scenario: refresh
without a ceremony, rotation, the spent token refused. Grown for the
shell's sessions, which died with their one-hour bearer. Shipped as
**v0.6.0**, pinned on soulstream main.

**The console becomes the standalone deployment's surface** ([episode
0091](0091-ecosystem-the-shell-is-the-console.md); **v0.5.0**, design
D31): the admin API always serves; whether the HTML console does is a
serve/embed option — on standalone, off inside the soulstream product,
where the shell administers through the same published API. The page
copy and README shed the bynames in the same pass (episode 0092).

**The fold wears the ecosystem's canon** ([episode
0090](0090-idp-the-fold-wears-the-canon.md); **v0.4.2**, design D30):
sign-in, enrolment, and the admin console dropped the dark
violet theme of the design system's earlier iteration for the
Soulsystem light canon the shell vendors — one rewritten style block,
same pages, same classes, the CRT surface kept for the shown-once
token alone. The door now reads as the same product as the console it
opens onto.

**The admin console — a browser and your passkey**
([episode 0011](0061-soulfold-the-admin-console.md); `v0.3.0`): the
admin surface split (D25) — the JSON API moved to `/api/admin`, and a
server-rendered console landed at `/admin`, authenticated by a
passkey session gated on the `admin` group. From a browser: create
people, mint their enrolment invites (shown once), set groups,
disable accounts, register clients — CSRF-guarded, no server-box
access. Verified in a real Chromium session with a CDP virtual
authenticator, which caught a genuine constraint the Go test could
not: WebAuthn refuses a bare-IP issuer host (use `localhost` or the
fronted name, never `127.0.0.1`). Browser-session logic extracted to
`internal/websession`.

**M3 ships — the fold is complete**
([episode 0010](0060-soulfold-m3-the-lifecycle.md);
`specs/005-the-lifecycle/`, **v0.2.0**): users, groups→roles, invites,
client registration, and the `/admin` JSON surface on the graduated
lifecycle design — enrollment everywhere by single-use digest-stored
invite, first-touch deleted, the four research bars riding `make test`
permanently. The M3 gate measured: four acts to a signed-in admin;
membership changes in the next token; admin-surface authz. Soulnode
composed it the same hour: the founding invite prints once beside the
founding token, and the owner's bearer carries admin (fold surface) +
realm (callout admission) with the two authorities cleanly separate.
**Every milestone on the fold's roadmap is shipped.**

**The bootstrap story concludes — invitation is the only door**
([episode 0009](0059-soulfold-bootstrap-story.md)): M3's gating
research passed all four bars against the fold itself — a fresh fold
to a signed-in admin in four counted acts, invites exactly-once
(25/25 races) and digest-stored with positive controls, the opened
store's 29 artifacts all refusing as enrollment, and the pocket-id
audit fixing M3's scope (its open `/setup` bootstrap refused — the
operator-act invite is stronger trust at the same act count).
First-touch enrollment is deleted, not disabled. Third design doc:
[lifecycle](../02-DESIGN/soulstream-idp/lifecycle.md) (D20–D24). **The M3
build lands next.**

**M5 ships — the embed seam**
([episode 0008](0056-soulfold-m5-the-embed-seam.md);
`specs/004-embed-seam/`): the public `embed.Run(ctx, Options)`
assembly (D29 pattern) with the AS-contract half the bundled story
needs — RFC 7591 DCR and the fixed token audience, both opt-in. The
compiler-proof consumer gate (`e2e/embedgate`, module path outside the
namespace) runs discovery → DCR → passkey sign-in → audience/
roles-bearing tokens; the M4 rig's fold half now rides the seam;
`soulfold serve` is the seam's first consumer; `authtest` went public
for consumers' own gates. **The public-door path is complete on the
fold's side — M3 (lifecycle + bootstrap research) is all that
remains.**

**M4 ships — the fold joins the fleet**
([episode 0007](0054-soulfold-m4-the-fold-in-the-fleet.md);
`specs/003-fold-in-the-fleet/`, the `e2e/` rig module): a browser
user's passkey sign-in becomes a NATS admission — fold-issued access
token + public sentinel through soulidentity's auth callout at its
published tag v0.1.0, the role value resolving against the declared
binding, scope server-enforced, refusals audited; and the identical
rig passes with an Entra-shaped stub (constitution II demonstrated).
The fold's tokens speak Entra's claim vocabulary (oid /
preferred_username / roles) because the seam's verifier keys subjects
by oid. **Soulnode's public door is unblocked; M5 (embed) next.**

**M2 ships — the refusal becomes behavior**
([episode 0006](0053-soulfold-m2-passkeys.md); `specs/002-passkeys/`):
WebAuthn ceremonies from go-webauthn are the only way through — the
username-form login is gone, first touch registers a passkey, second
touch asserts it, all proven in `make test` by a virtual authenticator
doing real ES256 ceremonies through the full HTTP flow (stock go-oidc
RP, restart mid-flow, forged POSTs inert). The D14 origin matrix
refuses all four foreign shapes server-side; no credential secret in
the store, positive-control-verified. Pending human act: the
physical-authenticator runbook (feature quickstart). First-touch
enrollment is the loud interim until M3's bootstrap research. **Next:
M4 — the callout admission proof.**

**M1 ships — the OP skeleton is real**
([episode 0005](0052-soulfold-m1-the-op-skeleton.md);
`specs/001-op-skeleton/`): discovery, JWKS, and the code+PKCE flow on
the sealed store, all measured in `make test` against an embedded
nats-server — a stock go-oidc RP signs in end to end (tokens verifying
against published JWKS, subject the seeded user), a full restart
between the login POST and the token exchange is invisible, forged
POSTs (missing/wrong CSRF, foreign Origin) leave the store's revision
unmoved, a rotation runs under a never-restarted verifier at 61/61
verifications, and the custody scan stays clean with its positive
control. The page inventory is exactly {login, error}. **Next: M2
(passkeys), then M4 before M3** per the operator's public-door
priority.

**The envelope is decided — M1 is unblocked**
([episode 0004](0051-soulfold-kv-encryption-at-rest.md)): the last
M1-gating research concluded with all four pre-registered bars passing
[measured] — the store's proven mechanics survive xkey sealing
byte-for-byte (restart 7/7, matrix 24/24, CAS 8,000/8,000,
redemption 100/100), a stopped-store scan with positive controls
recovers nothing (and forced the D6 amendment digesting the username
index key), filestore encryption was measured leaving the NATS surface
plaintext (the decisive asymmetry), and the cost landed at +44 B / ~57 µs
per record and +1.19 ms on a real end-to-end sign-in. The store design
grew D16–D19; every research gate M1 names is now concluded to design —
the build is next.

**The sign-in surface is decided**
([episode 0003](0046-soulfold-session-and-ui-shape.md)): the second M1-gating
research concluded with all four bars passing [measured] — the whole
flow is two server-rendered pages and zero JavaScript (15/15 rig
checks with a stock RP), a process restart mid-flow is invisible, the
CSRF posture rejects every forged shape with zero state change, and
the WebAuthn origin/RP-ID matrix (10/10) fixed the naming rules:
renaming the public host invalidates enrolled passkeys. Second design
doc: [session-and-ui](../02-DESIGN/soulstream-idp/session-and-ui.md) (D9–D15). M1's
two named research topics are done; one new topic was opened at the
operator's direction before the build starts — KV entry protection at
rest (xkeys), since the record envelope sits inside the store-shape
one-way door.

**The store is decided** ([episode 0002](0043-soulfold-kv-schema-and-key-lifecycle.md)):
M1's gating research concluded with all four pre-registered bars passing
[measured] — restart round-trip 6/6 byte-identical, additive decode
25/25 (and the cross-version RMW trap measured into design rule D3),
CAS with zero lost updates and exactly-once code redemption 100/100,
and a full JWKS rotation with 0 failures in 466 verifications under a
never-restarted go-oidc verifier. The fold's first design doc landed:
[store-and-key-lifecycle](../02-DESIGN/soulstream-idp/store-and-key-lifecycle.md)
(D1–D8, including RS256 for the seam). M1's remaining research is the
session and UI shape; the build follows it.

**Genesis — the fold is founded from a refusal**
([episode 0001](0041-soulfold-genesis-the-fold.md)): soulidentity's default-IdP
question (its journey 0019) resolved with the identity plane's refusal
holding — the ecosystem's default IAM is this sibling project: a
NATS-native (JetStream KV), embeddable, passkey-first OIDC provider that
consumers reach exclusively through standard OIDC, indistinguishable
from Entra by design. Build-vs-adopt was examined (pocket-id is an
application, not an embeddable library; no embeddable Go passkey-IdP
library exists [judgment]); the constitution (1.0.0) fixes the founding
constraints — passkeys, not passwords; indistinguishable by design —
and the roadmap sequences M1 (the OP skeleton) behind its KV-schema and
key-lifecycle research. No product code exists yet.

### soulstream-shell — the shell (as of 2026-08-27; named soulhelm until episode 0069)

**The declare surface is the named next module — behind the
dispatcher's submit op** ([episode
0140](0140-ecosystem-the-focus-agents-as-infrastructure.md)): the
agents-as-infrastructure focus's human end — authoring and submitting
agent declarations from the shell. Its design waits for research
`agents-as-infrastructure` to fix the submit op's shape; Bar 5 of that
topic exists to prove the whole declare→submit→served→answer loop
reachable from the shell's pure-consumer position before the module
design is written (the one-way door holds by measurement, not hope).

**The first hour BUILT — v0.11.0-rc.3, composed into the house the
same day** ([episode 0127](0127-shell-the-first-hour-builds.md)):
Home's first-steps card derives from realm state at every render — the
no-store rule a pure function under unit test, the whole lifecycle
walked by an unseeded e2e (four pending steps, each act flips its own,
the card gone with nothing to dismiss) [measured]. The roster's Around
column reads the presence face in the person's words (in / left {when}
/ seen {when}, the honest dash, a 5s live channel), the paste card
names the next step, and the tools/people/apps empty states now offer
their act. The support layer's last-seen refusal rewritten honestly:
no store of its own — the realm grew a face to read. Remaining: the
byon soak and the fresh-eyes install ([O1] — the ordering is still
asserted, not measured).

**The first hour is designed, ahead of its build** ([episode
0124](0124-ecosystem-the-first-hour-and-the-presence-lease.md); design
[`0008`](../02-DESIGN/soulstream-shell/0008-the-first-hour.md)): the
operator named the 0116 disease one level up — the product without a
first hour, every screen assuming the person knows why they are
there. The answer keeps the house rules: **guidance is a reading,
never a store** — a first-steps card on Home derived entirely from
realm state per render (no onboarding store, no per-person tour, no
wizard; D26 and 0001 §4 upheld), the **arrival principle** (an act
ends when the realm's own evidence shows the thing arrived, live, on
the acting screen), and empty states that offer their act. The agent
path's floor is honest to what exists — the paste block already
ships; the wrap announces nothing, so first evidence is the first
answer [measured, code trace] — and **upstream ask #3** asks the wrap
for a profile on start plus a presence lease (the core extension
decided in the same episode). Open and named: the fresh-eyes install
that would test the step ordering ([O1]).

**The sheet shape: the cognitive-load pass** ([episode
0123](0123-shell-the-sheet-shape.md); design
[`0007`](../02-DESIGN/soulstream-shell/0007-the-sheet-shape.md)): the
operator's read of the landed screens, answered in one pass. The three
list pages lead with their tables and each add-form waits in a shared
slide-over (the rail drawer's mechanics on the frame's own panel
signal, at every width) with a result line of its own; the tools form
branches on Kind — four everyday fields, the seven provider fields
folded under "Provider sign-in"; every destructive act stands behind a
question in the archive confirm's shape behind one-word keys (Disable,
Enable, Revoke, Remove — the sentence rides the hover); rows speak the
person's words (new / going on / quiet / closed / archived, a work
mark's sentence agreeing with its stamped strip, the status act
answering in a sentence); raw ids ride hovers; the storage panel leads
with payload and verdict, Headers and Signed bytes folded; and below
1180px the details column becomes a drawer on its own signal instead
of vanishing with the close/archive acts. The standing gates were
re-pointed, not weakened — the banned-word gate refused `.archfold`'s
class name on a sheet (hence `.stow`), and the e2e tripped on exactly
the copy the pass exists to change. On branch `ux-cognitive-load`,
gate green, unreleased.

**Tools and approvals join the spine — v0.11.0-rc.2** ([episode
0122](0122-ecosystem-the-shell-arc-lands.md); designs
[`0005`](../02-DESIGN/soulstream-shell/0005-the-tools-module.md)/[`0006`](../02-DESIGN/soulstream-shell/0006-the-approvals-module.md)):
the tools screen — the catalog with each person's own standing,
connect/disconnect as the person's own OAuth ceremony through the
module's callback, the admin's both-halves add with the secret crossing
once — and the approvals screen — the guardrail's tickets with their
count on the spine, one tap minting the person's own signed answer and
delivering it to the originator's tail, mint and delivery kept apart in
the design's own words. Activation by the deployment's declared
GuardrailOn fact. The consumer-position gate walks both against the
published product; composed into the house as soulstream v0.13.0-rc.11.
Episode 0116's three asks are all screens now.

**The store shows what it holds** ([episode
0117](0117-shell-the-store-shows-what-it-holds.md); design
[`0004`](../02-DESIGN/soulstream-shell/0004-the-storage-explorer.md),
built the same day it was written): a module of its own for reading the
stores — both of them under a plain name and the name the server
answers to, a page of messages newest first over a validated subject
pattern, one message whole (headers verbatim, payload as it is, the
canonical form beside it), and a live tail that is a mode of the screen
rather than a toggle on it. Verdicts run through core's own
`VerifyRecord`; zero upstream additions. Reads ride the signed-in
person's own admission rather than the shared read lane, and the screen
says out loud that this narrows nothing here — the product grants every
persona the whole subject space, and a test refuses four ways of
implying otherwise. No act, no delete, no index, no search: the query
layer is what the protocol deliberately lacks. Shipped as
**v0.11.0-rc.1**, a marked candidate. The design's own first
draft was corrected by the build — `SOULSTREAM.>` is not a store; the
op-log captures `SOULSTREAM.TOPICS.>`, the inboxes are their own
stream, and the service lane is kept nowhere — which made the custody
argument stronger, not weaker. Measured live at 1000 px and 390 px:
zero page overflow, the table scrolling inside its own wrapper.

**The evaluation names what the screens lack** ([episode
0116](0116-ecosystem-what-shipped-without-a-human-end.md)): living on
the candidate produced three asks, and the survey found the shell owns
one of them outright — the storage explorer above. The other two,
external tools and the guardrail's human end, became research topics,
both with a shell module at the end of them.

**The session renews its bearer — and ends when it no longer can**
([episode 0103](0103-ecosystem-the-session-outlives-its-token.md)):
the shell held sign-in's one-hour access token for the life of an
unexpiring session, and a console left open rotted an hour in — NATS
reconnects re-presenting the dead bearer, the admin lane 401ing, the
screen still saying signed in. The session now custodies the grant:
`offline_access` at sign-in, renewal through the idp's new refresh
grant behind the same session (the NATS admission by token handler on
every reconnect, the admin lane per call), and a session that cannot
produce a living credential closes itself into the sign-in card —
drained exactly once, table-tested. Shipped as **v0.9.0** (the
`Session.Bearer` field became the renewing method — the pre-v1 clean
break), pinned on soulstream main.

**Conversations begin and end in the shell** ([episode
0102](0102-shell-conversations-begin-and-end.md); design
[`0003`](../02-DESIGN/soulstream-shell/0003-conversation-lifecycle.md)):
starting from a fold in the rail and a card on Home (one act, the
person lands in the new conversation), the close-then-archive ladder
in the details panel — Close while live, Archive only once closed and
behind a two-step ask, no reopen because the record has none —
archived conversations folded at the foot of the list with the toggle
surviving the stream's morphs, the composer yielding where the record
refuses the write, and copy that stays honest through the record's
half-successes, all table-tested. Zero upstream additions; the e2e
gate walks the whole ladder against a real realm.

**The fronted console signs in** ([episode
0101](0101-shell-the-fronted-console-signs-in.md); **v0.7.0**):
`Options.PublicURL` — a console fronted by TLS termination or a
tailnet serve registers an OAuth callback the visitor's browser can
actually reach, instead of its own bound loopback address. Found live
on the byon deployment's first passkey sign-in; measured fixed on the
same deployment (product wiring `planes.shell.public_url`, shipped in
soulstream v0.13.0-rc.4 after rc.3's pass-through gap — the plane gate
now asserts the live /login redirect). Pinned by soulstream.

**The shell is the console** ([episode
0091](0091-ecosystem-the-shell-is-the-console.md); **v0.6.0**): the
People & sign-in screen grew to the whole published admin contract —
add a person, edit groups in place, shut out and let back in, and the
apps that sign people in — drawn on the spine only for sessions whose
own token carries the admin role, with the sign-in service's verified
refusals as the authority. The idp's HTML console unmounts in the
bundled product (D31); administration has one home.

**The agents screen breathes, and the wrap leads** ([episode
0089](0089-ecosystem-wrap-in-the-house.md); **v0.5.0**): the credential
card now opens with one portable paste block — writes the creds file,
runs `soulstream wrap --harness claude`, POSIX and fish alike — with a
Copy key, and folds the hard paths beneath as quiet cards, every one
spelling the door as the product binary. The screen took the canon's
rhythm (sections, styled folds, no stray note, no doubled operator
ids), and the overview grew an Agents card pointing the way to a first
agent. The pinned agents e2e passed with the new markup.

**The credential screen says where it goes — and now offers the wrap**
([episodes 0084](0084-shell-where-this-goes.md)/[0085](0085-workloads-wrap-run-your-agent-where-you-are.md)):
the shown-once agent credential
(episode 0079's block) now carries per-assistant set-up folds — Claude
Code (the block is already the exact `.mcp.json`), codex (the same five
values as a filled TOML table), and anything else that speaks MCP
(pi.dev, opencode, …) in plain words — plus the one `soulstream-mcp`
install line. Same reveal, same shown-once lifetime; the pinned 0079
e2e passed untouched.

**Founded, shipped, and composed in one day** ([episode
0067](0067-soulhelm-founding-and-first-light.md); research graduated
the same morning, [episode 0066](0066-ecosystem-soulsystem-cockpit.md)):
**v0.1.0** is the observe surface (board, topics with earned signature
verdicts, storage, plane health) rendered in the cassette-light token
source over Datastar SSE, with fold sessions (RFC 7591 DCR +
code+PKCE, memory-only, each session admitting as itself through the
OIDC callout lane) and the first act — `work.open` signed as the
signed-in principal. The surface is closed until sign-in (an open
scope amendment: sessions moved into the founding release rather than
ship an unauthenticated realm viewer). The consumer-position e2e
(module path outside the namespace, upstreams at published tags) walks
the whole human ceremony in ~4 s inside `make test`. Soulnode composed
it the same day as `planes.helm` ([episode
0068](0068-soulnode-the-helm-plane.md)).

**The direction changed the same day** — the operator reframed the
component as **the shell**: a pure modular frame containing zero module
logic, agnostic from Soulstream by contract, every human surface — the
observe core, the fold's administration, the agent designer/manager to
come — a module plugging in through one exported contract, with the
shell providing composition (registration, activation, navigation,
sessions, cross-linking). Research topic:
[`shell-module-contract`](0078-shell-the-module-contract.md),
four pre-registered bars. And the ecosystem naming re-centering renames
the component **soulstream-shell** ([episode
0069](0069-ecosystem-one-name-soulstream.md); bare *soulshell* was
decided and superseded the same day), executed in the sweep ([episode
0070](0070-ecosystem-the-rename-sweep.md), v0.2.0).

**The focus (2026-08-14,** [episode
0071](0071-ecosystem-the-focus.md)**):** participation enters scope —
the mission is the **usable cockpit**: view topics, collaborate
directly (post, reply, comment, open topics), mention notifications,
riding the backend-held session admission already measured. The
shell-module-contract bars ride that build, with the collaboration
module as Bar 1's second module. **The composer landed the same
morning** ([episode 0072](0072-shell-the-composer.md)) — post and
anchored reply as the session's own principal, two pre-existing
rendering bugs fixed — and **the chat shape the same day** ([episode
0073](0073-shell-the-chat-shape.md)): conversations rail, own-message
bubbles, threaded replies, docked composer, plus the `make screens`
helper whose first real screenshots caught the retired *soulsystem*
wordmark. **The screen review landed the same morning** ([episode
0074](0074-shell-the-spine-and-the-details.md)): the icon rail with
Home ("Your realm at a glance"), the capped centered column, and the
details panel — People, Status, Waiting on — all verified on live
screenshots. **Mentions landed the same day** ([episode
0075](0075-shell-mentions-and-a-word-retired.md)): the session follows
its own inbox, marks clear by reading, the mentioning message stays
highlighted off the record; "Daan" replaces the raw id; "realm"
retired from every screen with a mechanical assertion. **The named gap
closed the same day** ([episode
0076](0076-shell-a-name-that-taps-somebody.md)): core v0.8.1's
explicit-mentions arm (the first consumer-proven addition since the
focus) plus the shell's @-picker — the body keeps what was typed, who
it taps rides beside it, and C4's reversal condition was tested and
survived with its cost named. **The design pass held the canon**
([episode 0077](0077-shell-the-canon-held.md)): the vendored design
contract, hardware treatment everywhere, and the channel semantics
made honest — amber for a voice that answers for itself, teal for a
voice somebody else answers for (`operated_by`, accountability not
species), Scribe the first machine voice on screen as a peer. **The
research graduated** ([episode 0078](0078-shell-the-module-contract.md),
design [`0002-the-module-shape.md`](../02-DESIGN/soulstream-shell/0002-the-module-shape.md)):
all four bars PASS — the pure shell (compiler-guarded), the support
layer, three modules (overview, conversations, People & sign-in — the
last activating only when the deployment administers its own people),
cross-linking without imports, and an outside-namespace probe module.
**Agents joined the stream the same evening** ([episode
0079](0079-shell-agents-join-the-stream.md)): mint a credential in the
shell, paste it into the stdio door's config, and the agent
collaborates as a named, accountable, revocable teal voice — core
v0.8.2's revocable lane, shell v0.4.0's Agents module, soulstream
**v0.10.0**. **The frame went responsive the same night** ([episode
0080](0080-shell-one-instrument-any-width.md)): zero horizontal
overflow measured at 1000 px and 390 px across every screen, hermetic
guards in the gate — shell v0.4.1, soulstream **v0.10.1**. **The
owner can no longer be locked out** ([episode
0081](0081-idp-the-last-administrator-stays.md)): the idp enforces
"the admin group never empties" atomically under CAS (raced 25×8,
one admin standing every round), the shell reflects it — idp v0.4.1,
shell v0.4.2, and the evaluation candidate tagged as the operator
asked: **soulstream v0.11.0-rc.1**, a marked pre-release. **Next:**
the operator lives on the candidate; what chafes decides.

## Episode index

| # | Component | Episode |
|---|---|---|
| 0001 | soulrealm | [Genesis: soulrealm gets an HQ (2026-07-22)](0001-soulrealm-genesis.md) |
| 0002 | soulrealm | [The substrate decision: a from-scratch, NEX-influenced runtime (2026-07-22)](0002-soulrealm-the-substrate-decision.md) |
| 0003 | soulrealm | [Soulstream-only scope: the platform waits (2026-07-22)](0003-soulrealm-soulstream-only-scope.md) |
| 0004 | soulrealm | [The first agent runs (2026-07-22)](0004-soulrealm-the-first-agent-runs.md) |
| 0005 | soulrealm | [A tool answers (2026-07-22)](0005-soulrealm-a-tool-answers.md) |
| 0006 | soulrealm | [HQ alignment: the lint gets built (2026-07-24)](0006-soulrealm-hq-alignment.md) |
| 0007 | soulstream | [Genesis to v0.3: the protocol, and the library that proves it (2026-07-11 → 2026-07-24)](0007-soulstream-genesis-and-the-reference-library.md) |
| 0008 | soulstream | [Adopting the hq way: the process gets a constitution (2026-07-24)](0008-soulstream-adopting-the-hq-way.md) |
| 0009 | soulstream | [The memory convention: the realm learns to be asked (2026-07-25)](0009-soulstream-memory-convention-and-exhibits.md) |
| 0010 | soulstream | [Provisioning byte limits: the strict landlord gets a one-command realm (2026-07-27)](0010-soulstream-provisioning-byte-limits.md) |
| 0011 | soulstream | [Sealed topics survive the substrate: four bars, one encoding amendment (2026-07-27)](0011-soulstream-sealed-topics.md) |
| 0012 | soulidentity | [Genesis: the design thread and the walking skeleton (2026-07-28)](0012-soulidentity-genesis-and-the-walking-skeleton.md) |
| 0013 | soulidentity | [The identity-plane re-centering (2026-07-28)](0013-soulidentity-the-identity-plane-re-centering.md) |
| 0014 | soulidentity | [NATS-only and the connection ladder (2026-07-28)](0014-soulidentity-nats-only-and-the-connection-ladder.md) |
| 0015 | soulidentity | [The first-key story: a local file, named honestly (2026-07-28)](0015-soulidentity-first-key-story.md) |
| 0016 | soulidentity | [The NATS-surface design: the principal is the subject (2026-07-28)](0016-soulidentity-the-nats-surface-design.md) |
| 0017 | soulidentity | [Design review: seeds from the environment, no v1, D15 taught back (2026-07-28)](0017-soulidentity-design-review-amendments.md) |
| 0018 | soulidentity | [M3: the NATS-native rebuild ships (2026-07-28)](0018-soulidentity-m3-the-nats-native-rebuild.md) |
| 0019 | soulidentity | [The sentinel-credential flow: URL + token is enough (2026-07-28)](0019-soulidentity-sentinel-credential-flow.md) |
| 0020 | soulrealm | [A second wall: the microsandbox backend (2026-07-28)](0020-soulrealm-a-second-wall.md) |
| 0021 | soulidentity | [The claims-mapping shape: one pipeline, policy never in the credential store (2026-07-28)](0021-soulidentity-claims-mapping-shape.md) |
| 0022 | soulidentity | [M4: auth callout ships, the front door opens (2026-07-28)](0022-soulidentity-m4-auth-callout-ships.md) |
| 0023 | soulidentity | [The shared subject prefix: one namespace for the ecosystem (2026-07-28)](0023-soulidentity-the-shared-subject-prefix.md) |
| 0024 | soulrealm | [A third wall on rented ground: Kubernetes as a backend (2026-07-28 → 2026-07-29)](0024-soulrealm-kubernetes-backend.md) |
| 0025 | soulidentity | [The Entra lane: role == team, no mapping store (2026-07-29)](0025-soulidentity-entra-role-claim-lane.md) |
| 0026 | soulstream | [The signer seam: signing learns to be delegated (2026-07-29)](0026-soulstream-the-signer-seam.md) |
| 0027 | soulstream | [DX hardening: the seam's two sharp edges, and the cycle guard (2026-07-29)](0027-soulstream-dx-hardening-and-the-cycle-guard.md) |
| 0028 | soulrealm | [A third wall lands: the Kubernetes backend ships (2026-07-29)](0028-soulrealm-a-third-wall-lands.md) |
| 0029 | soulidentity | [The registry dissolves: authorization in the ACLs and the bindings (2026-07-29)](0029-soulidentity-the-registry-dissolves.md) |
| 0030 | soulidentity | [The cross-service proof: the seam carries a real record (2026-07-29)](0030-soulidentity-the-cross-service-proof.md) |
| 0031 | soulidentity | [The vault is the directory: ephemeral users, keys on first touch (2026-07-29)](0031-soulidentity-the-vault-is-the-directory.md) |
| 0032 | soulidentity | [One noun: persona (2026-07-29)](0032-soulidentity-one-noun-persona.md) |
| 0033 | soulrealm | [Fleet: the log nominates, evidence vetoes, the log decides (2026-07-29 → 2026-07-31)](0033-soulrealm-fleet.md) |
| 0034 | soulnode | [Genesis: SoulNode gets an HQ (2026-07-31)](0034-soulnode-genesis.md) |
| 0035 | soulidentity | [Role selection by declared name: the ephemeral mint op (2026-07-31)](0035-soulidentity-role-selection-by-name.md) |
| 0036 | soulidentity | [The embed seam: the serve assembly becomes public (2026-08-01)](0036-soulidentity-the-embed-seam.md) |
| 0037 | soulrealm | [Pinned to the record: the soulstream replace drops (2026-08-01)](0037-soulrealm-pinned-to-the-record.md) |
| 0038 | soulstream | [The remote MCP node: a URL into the realm, proven on the BYON (2026-07-30 → 08-01)](0038-soulstream-remote-mcp-node.md) |
| 0039 | soulidentity | [Soulfold: the default IdP is a sibling, the refusal holds (2026-08-02)](0039-soulidentity-soulfold-the-default-idp.md) |
| 0040 | soulnode | [The composition gate: three bars PASS, the ecosystem opens its seams (2026-07-31 → 2026-08-02)](0040-soulnode-the-composition-gate.md) |
| 0041 | soulfold | [Genesis: the fold is founded from a refusal (2026-08-02)](0041-soulfold-genesis-the-fold.md) |
| 0042 | soulnode | [First boot is real: init and up land (2026-08-02)](0042-soulnode-first-boot-is-real.md) |
| 0043 | soulfold | [The store is decided: four bars, four passes (2026-08-02)](0043-soulfold-kv-schema-and-key-lifecycle.md) |
| 0044 | soulnode | [The realm remembers: provisioning and the memory plane land (2026-08-02)](0044-soulnode-the-realm-remembers.md) |
| 0045 | soulnode | [An agent runs: Phase 1 is complete (2026-08-02)](0045-soulnode-an-agent-runs.md) |
| 0046 | soulfold | [Two pages, zero scripts, and a name that becomes a door (2026-08-02)](0046-soulfold-session-and-ui-shape.md) |
| 0047 | soulstream | [The remote MCP node, built: the door that holds nothing (2026-08-02)](0047-soulstream-remote-mcp-node-built.md) |
| 0048 | soulstream | [The node becomes consumable: the replace drops (2026-08-02)](0048-soulstream-the-node-becomes-consumable.md) |
| 0049 | soulnode | [The door opens: the MCP front door lands (2026-08-02)](0049-soulnode-the-door-opens.md) |
| 0050 | ecosystem | [One hq: the five headquarters merge](0050-ecosystem-one-hq.md) |
| 0051 | soulfold | [The envelope is decided: xkeys seal the store (2026-08-02)](0051-soulfold-kv-encryption-at-rest.md) |
| 0052 | soulfold | [M1 ships: the OP skeleton on the sealed store (2026-08-02)](0052-soulfold-m1-the-op-skeleton.md) |
| 0053 | soulfold | [M2 ships: passkeys, the refusal becomes behavior (2026-08-02)](0053-soulfold-m2-passkeys.md) |
| 0054 | soulfold | [M4 ships: the fold joins the fleet (2026-08-02)](0054-soulfold-m4-the-fold-in-the-fleet.md) |
| 0055 | soulnode | [The public door opens (2026-08-03)](0055-soulnode-the-public-door.md) |
| 0056 | soulfold | [M5 ships: the fold learns to live in your house (2026-08-03)](0056-soulfold-m5-the-embed-seam.md) |
| 0057 | soulnode | [The folded realm: one binary, real people (2026-08-03)](0057-soulnode-the-folded-realm.md) |
| 0058 | soulnode | [v0.1.0: the house gets a shipping label (2026-08-03)](0058-soulnode-the-release-pipeline.md) |
| 0059 | soulfold | [The bootstrap story: invitation is the only door (2026-08-03)](0059-soulfold-bootstrap-story.md) |
| 0060 | soulfold | [M3 ships: the fold is complete (2026-08-03)](0060-soulfold-m3-the-lifecycle.md) |
| 0061 | soulfold | [The admin console: a browser and your passkey (2026-08-03)](0061-soulfold-the-admin-console.md) |
| 0062 | soulnode | [The front of house: URLs, the console, no collision (2026-08-03)](0062-soulnode-the-front-of-house.md) |
| 0063 | soulfold | [The console gets a face (2026-08-03)](0063-soulfold-the-console-gets-a-face.md) |
| 0064 | ecosystem | [The platform turn: tenancy, guardrails, and eight decisions (2026-08-04)](0064-ecosystem-the-platform-turn.md) |
| 0065 | ecosystem | [The ecosystem goes fair-code (2026-08-08)](0065-ecosystem-fair-code.md) |
| 0066 | ecosystem | [The helm: the cockpit earns its design (2026-08-13)](0066-ecosystem-soulsystem-cockpit.md) |
| 0067 | soulhelm | [Founding and first light: the helm is real (2026-08-13)](0067-soulhelm-founding-and-first-light.md) |
| 0068 | soulnode | [The helm plane: the cockpit joins the bundle (2026-08-13)](0068-soulnode-the-helm-plane.md) |
| 0069 | ecosystem | [One name: Soulstream (2026-08-13)](0069-ecosystem-one-name-soulstream.md) |
| 0070 | ecosystem | [The rename sweep: eight repos, one evening (2026-08-13)](0070-ecosystem-the-rename-sweep.md) |
| 0071 | ecosystem | [The focus: a product, not a platform (2026-08-14)](0071-ecosystem-the-focus.md) |
| 0072 | shell | [The composer: the shell stops being a window (2026-08-14)](0072-shell-the-composer.md) |
| 0073 | shell | [The chat shape: a rail, a conversation, a docked composer (2026-08-14)](0073-shell-the-chat-shape.md) |
| 0074 | shell | [The spine and the details: the operator's screen review lands (2026-08-14)](0074-shell-the-spine-and-the-details.md) |
| 0075 | shell | [Mentions land, and a word retires (2026-08-14)](0075-shell-mentions-and-a-word-retired.md) |
| 0076 | shell | [Type @, pick a person: a name that taps somebody (2026-08-14)](0076-shell-a-name-that-taps-somebody.md) |
| 0077 | shell | [The canon, held: two channels at equal weight (2026-08-14)](0077-shell-the-canon-held.md) |
| 0078 | shell | [The module contract: four bars, one shell (2026-08-13 → 2026-08-14)](0078-shell-the-module-contract.md) |
| 0079 | shell | [Agents join the stream (2026-08-14)](0079-shell-agents-join-the-stream.md) |
| 0080 | shell | [One instrument, any width (2026-08-14)](0080-shell-one-instrument-any-width.md) |
| 0081 | idp | [The last administrator stays (2026-08-14)](0081-idp-the-last-administrator-stays.md) |
| 0082 | ecosystem | [What wakes an agent (2026-08-15)](0082-ecosystem-agent-participation.md) |
| 0083 | workloads | [The waker lands (2026-08-15)](0083-workloads-the-waker-lands.md) |
| 0084 | shell | [Where this goes (2026-08-15)](0084-shell-where-this-goes.md) |
| 0085 | workloads | [Wrap: run your agent where you are (2026-08-15)](0085-workloads-wrap-run-your-agent-where-you-are.md) |
| 0086 | soulstream | [v0.11.0-rc.2: wrap ships (2026-08-15)](0086-soulstream-v0-11-0-rc-2.md) |
| 0087 | ecosystem | [The docs catch up, and the site learns to onboard (2026-08-15)](0087-ecosystem-the-docs-catch-up.md) |
| 0088 | ecosystem | [The site speaks the new names (2026-08-15)](0088-ecosystem-the-site-speaks-the-new-names.md) |
| 0089 | ecosystem | [Wrap in the house: one binary, one paste (2026-08-15)](0089-ecosystem-wrap-in-the-house.md) |
| 0090 | idp | [The fold wears the ecosystem's canon (2026-08-15)](0090-idp-the-fold-wears-the-canon.md) |
| 0091 | ecosystem | [The shell is the console (2026-08-15)](0091-ecosystem-the-shell-is-the-console.md) |
| 0092 | ecosystem | [The names say what they do (2026-08-15)](0092-ecosystem-the-names-say-what-they-do.md) |
| 0093 | soulstream | [Pre-v1 renames are clean breaks (2026-08-16)](0093-soulstream-pre-v1-renames-are-clean-breaks.md) |
| 0094 | soulstream | [v0.12.0-rc.1: the clean-break candidate (2026-08-16)](0094-soulstream-v0-12-0-rc-1.md) |
| 0095 | soulstream | [BYO NATS designed: founding on a server we don't run (2026-08-16)](0095-soulstream-byo-nats-designed.md) |
| 0096 | soulstream | [BYO NATS ships: the kit, the probes, the driver (2026-08-16)](0096-soulstream-byo-nats-ships.md) |
| 0097 | soulstream | [v0.13.0-rc.1: bring your own server (2026-08-16)](0097-soulstream-v0-13-0-rc-1.md) |
| 0098 | soulstream | [brew install soulstream: the tap opens (2026-08-16)](0098-soulstream-the-tap-opens.md) |
| 0099 | soulstream | [First contact: a realm founded on a live Synadia Cloud BYON (2026-08-16)](0099-soulstream-the-byon-founding.md) |
| 0100 | soulstream | [v0.13.0-rc.2: first contact hardened, and the tap feeds itself (2026-08-16)](0100-soulstream-v0-13-0-rc-2.md) |
| 0101 | shell | [The fronted console signs in: PublicURL crosses the seam (2026-08-16 → 08-17)](0101-shell-the-fronted-console-signs-in.md) |
| 0102 | shell | [Conversations begin and end in the shell (2026-08-17)](0102-shell-conversations-begin-and-end.md) |
| 0103 | ecosystem | [The session outlives its token: the refresh grant crosses the seam (2026-08-17)](0103-ecosystem-the-session-outlives-its-token.md) |
| 0104 | ecosystem | [Outbound identity: every remote call carries the calling user (2026-08-17)](0104-ecosystem-outbound-identity-grants.md) |
| 0105 | identity | [The grants broker lands: outbound identity in custody (2026-08-18)](0105-identity-the-grants-broker-lands.md) |
| 0106 | idp | [The token-lifetime knob: revocation propagation gets a dial (2026-08-18)](0106-idp-the-token-lifetime-knob.md) |
| 0107 | ecosystem | [Tenancy and guardrails: the platform question answered (2026-08-04 → 2026-08-18)](0107-ecosystem-platform-tenancy-guardrails.md) |
| 0108 | ecosystem | [F1 closes: every signing persona becomes readable (2026-08-18)](0108-ecosystem-the-key-becomes-resolvable.md) |
| 0109 | ecosystem | [Consent enters the record, and Bar 4 finally measures (2026-08-18)](0109-ecosystem-consent-enters-the-record.md) |
| 0110 | identity | [The tenancy set builds: secrets, the guardrail, approvals, accounts (2026-08-19)](0110-identity-the-tenancy-set-builds.md) |
| 0111 | ecosystem | [One session, several audiences: the exchange grant lands on both ends (2026-08-19)](0111-ecosystem-one-session-several-audiences.md) |
| 0112 | ecosystem | [The canonical form breaks clean: the realm's key signs, the record names the hand (2026-08-19)](0112-ecosystem-the-canonical-form-breaks-clean.md) |
| 0113 | workloads | [Placement is work.claim: the fleet lands (2026-08-19)](0113-workloads-placement-is-work-claim.md) |
| 0114 | identity | [The provider arm closes: an account born on someone else's operator key (2026-08-19)](0114-identity-the-provider-arm-closes.md) |
| 0115 | soulstream | [byon adopts the new form: the migration that refused to be a re-founding (2026-08-19)](0115-soulstream-byon-adopts-the-new-form.md) |
| 0116 | ecosystem | [What shipped without a human end (2026-08-19)](0116-ecosystem-what-shipped-without-a-human-end.md) |
| 0117 | shell | [The store shows what it holds (2026-08-19)](0117-shell-the-store-shows-what-it-holds.md) |
| 0118 | ecosystem | [External tools: the agent reaches out, the person is who the remote sees (2026-08-19 → 2026-08-21)](0118-ecosystem-agent-external-tools.md) |
| 0119 | ecosystem | [The guardrail's human end: the loop measured to where it stops, then designed shut (2026-08-19 → 2026-08-21)](0119-ecosystem-guardrail-human-end.md) |
| 0120 | ecosystem | [The tools arc builds: catalog, custody, door (2026-08-21)](0120-ecosystem-the-tools-arc-builds.md) |
| 0121 | ecosystem | [The approvals loop closes: tickets, presentation, policy (2026-08-21)](0121-ecosystem-the-approvals-loop-closes.md) |
| 0122 | ecosystem | [The shell arc lands: both human ends on the spine (2026-08-21)](0122-ecosystem-the-shell-arc-lands.md) |
| 0123 | shell | [The sheet shape: tables lead, forms slide over, keys ask first (2026-08-23)](0123-shell-the-sheet-shape.md) |
| 0124 | ecosystem | [The first hour and the presence lease (2026-08-23 → 2026-08-24)](0124-ecosystem-the-first-hour-and-the-presence-lease.md) |
| 0125 | ecosystem | [The presence lease builds: the lamp lights (2026-08-24)](0125-ecosystem-the-presence-lease-builds.md) |
| 0126 | ecosystem | [Agent declaration: the record declares, the room runs (2026-08-23 → 2026-08-25)](0126-ecosystem-agent-declaration.md) |
| 0127 | shell | [The first hour builds: the card derives, the roster breathes (2026-08-25)](0127-shell-the-first-hour-builds.md) |
| 0128 | ecosystem | [Loop safety: the room budgets the cascade (2026-08-25)](0128-ecosystem-loop-safety.md) |
| 0129 | workloads | [The wake budget builds: the colony gate ships (2026-08-25)](0129-workloads-the-wake-budget-builds.md) |
| 0130 | ecosystem | [The agent declaration builds: the record declares, the room runs (2026-08-25)](0130-ecosystem-the-agent-declaration-builds.md) |
| 0131 | core | [Sealed topics build: the locked binder (2026-08-25)](0131-core-sealed-topics-build.md) |
| 0132 | soulstream | [The rc carries both builds: v0.14.0-rc.1 across the stack (2026-08-25)](0132-soulstream-the-rc-carries-both-builds.md) |
| 0133 | ecosystem | [The platform-account topology: measured sound, two fixes short (2026-08-26 → 2026-08-27)](0133-ecosystem-platform-account-topology.md) |
| 0134 | identity | [Tenants are born admissible: D47 lands (2026-08-27)](0134-identity-tenants-born-admissible.md) |
| 0135 | ecosystem | [The residues close: tenancy reaches the hand, both arms carry D47 (2026-08-27)](0135-ecosystem-the-residues-close.md) |
| 0136 | identity | [The BYON live run: two defects caught, then the provider arm measures sound (2026-08-27)](0136-identity-the-byon-live-run.md) |
| 0137 | ecosystem | [Capability minting: the declaration's names become the credential (2026-08-27)](0137-ecosystem-capability-minting.md) |
| 0138 | identity | [The sealed record gains its custodian: D9 builds (2026-08-27)](0138-identity-the-sealed-record-gains-its-custodian.md) |
| 0139 | soulstream | [The rc.2 carries the tenants and the capabilities: v0.14.0-rc.2, identity v0.12.0, workloads v0.8.0-rc.2 (2026-08-27)](0139-soulstream-the-rc2-carries-the-tenants-and-the-capabilities.md) |
| 0140 | ecosystem | [The focus: agents as infrastructure (2026-08-27)](0140-ecosystem-the-focus-agents-as-infrastructure.md) |
| 0141 | ecosystem | [Agents as infrastructure: five bars in a day, the dispatcher is a composition (2026-08-27 → 2026-08-28)](0141-ecosystem-agents-as-infrastructure.md) |

## The naming map (2026-08-13)

The ecosystem naming re-centering (episodes
[0069](0069-ecosystem-one-name-soulstream.md) /
[0070](0070-ecosystem-the-rename-sweep.md)) renamed every project under
the one brand. Episodes 0001–0069 keep their filenames and component
tags — the record is append-only; this map resolves them. Future
episodes use the short tag.

| Old name (tag in episodes ≤ 0069) | Repository today | Episode tag now |
|---|---|---|
| soulstream (the record library) | soulstream-core | core |
| soulrealm | soulstream-workloads | workloads |
| soulidentity | soulstream-identity | identity |
| soulnode | soulstream (the product) | soulstream |
| soulfold | soulstream-idp | idp |
| soulhelm | soulstream-shell | shell |
| soulstream/node (nested module) | soulstream-mcp | mcp |
| soulstream-archivist | soulstream-archivist (unchanged) | — |
| — | soulstream-cli (to be founded) | cli |

## Pre-merge numbering map

Old (per-project) episode → new shared number. Frozen specs, old commit messages, and archived documents cite the old numbers.

| Component | Old # | New # | Episode |
|---|---|---|---|
| soulrealm | 0001 | 0001 | [0001-soulrealm-genesis.md](0001-soulrealm-genesis.md) |
| soulrealm | 0002 | 0002 | [0002-soulrealm-the-substrate-decision.md](0002-soulrealm-the-substrate-decision.md) |
| soulrealm | 0003 | 0003 | [0003-soulrealm-soulstream-only-scope.md](0003-soulrealm-soulstream-only-scope.md) |
| soulrealm | 0004 | 0004 | [0004-soulrealm-the-first-agent-runs.md](0004-soulrealm-the-first-agent-runs.md) |
| soulrealm | 0005 | 0005 | [0005-soulrealm-a-tool-answers.md](0005-soulrealm-a-tool-answers.md) |
| soulrealm | 0006 | 0006 | [0006-soulrealm-hq-alignment.md](0006-soulrealm-hq-alignment.md) |
| soulstream | 0001 | 0007 | [0007-soulstream-genesis-and-the-reference-library.md](0007-soulstream-genesis-and-the-reference-library.md) |
| soulstream | 0002 | 0008 | [0008-soulstream-adopting-the-hq-way.md](0008-soulstream-adopting-the-hq-way.md) |
| soulstream | 0003 | 0009 | [0009-soulstream-memory-convention-and-exhibits.md](0009-soulstream-memory-convention-and-exhibits.md) |
| soulstream | 0004 | 0010 | [0010-soulstream-provisioning-byte-limits.md](0010-soulstream-provisioning-byte-limits.md) |
| soulstream | 0005 | 0011 | [0011-soulstream-sealed-topics.md](0011-soulstream-sealed-topics.md) |
| soulidentity | 0001 | 0012 | [0012-soulidentity-genesis-and-the-walking-skeleton.md](0012-soulidentity-genesis-and-the-walking-skeleton.md) |
| soulidentity | 0002 | 0013 | [0013-soulidentity-the-identity-plane-re-centering.md](0013-soulidentity-the-identity-plane-re-centering.md) |
| soulidentity | 0003 | 0014 | [0014-soulidentity-nats-only-and-the-connection-ladder.md](0014-soulidentity-nats-only-and-the-connection-ladder.md) |
| soulidentity | 0004 | 0015 | [0015-soulidentity-first-key-story.md](0015-soulidentity-first-key-story.md) |
| soulidentity | 0005 | 0016 | [0016-soulidentity-the-nats-surface-design.md](0016-soulidentity-the-nats-surface-design.md) |
| soulidentity | 0006 | 0017 | [0017-soulidentity-design-review-amendments.md](0017-soulidentity-design-review-amendments.md) |
| soulidentity | 0007 | 0018 | [0018-soulidentity-m3-the-nats-native-rebuild.md](0018-soulidentity-m3-the-nats-native-rebuild.md) |
| soulidentity | 0008 | 0019 | [0019-soulidentity-sentinel-credential-flow.md](0019-soulidentity-sentinel-credential-flow.md) |
| soulrealm | 0007 | 0020 | [0020-soulrealm-a-second-wall.md](0020-soulrealm-a-second-wall.md) |
| soulidentity | 0009 | 0021 | [0021-soulidentity-claims-mapping-shape.md](0021-soulidentity-claims-mapping-shape.md) |
| soulidentity | 0010 | 0022 | [0022-soulidentity-m4-auth-callout-ships.md](0022-soulidentity-m4-auth-callout-ships.md) |
| soulidentity | 0011 | 0023 | [0023-soulidentity-the-shared-subject-prefix.md](0023-soulidentity-the-shared-subject-prefix.md) |
| soulrealm | 0008 | 0024 | [0024-soulrealm-kubernetes-backend.md](0024-soulrealm-kubernetes-backend.md) |
| soulidentity | 0012 | 0025 | [0025-soulidentity-entra-role-claim-lane.md](0025-soulidentity-entra-role-claim-lane.md) |
| soulstream | 0006 | 0026 | [0026-soulstream-the-signer-seam.md](0026-soulstream-the-signer-seam.md) |
| soulstream | 0007 | 0027 | [0027-soulstream-dx-hardening-and-the-cycle-guard.md](0027-soulstream-dx-hardening-and-the-cycle-guard.md) |
| soulrealm | 0009 | 0028 | [0028-soulrealm-a-third-wall-lands.md](0028-soulrealm-a-third-wall-lands.md) |
| soulidentity | 0013 | 0029 | [0029-soulidentity-the-registry-dissolves.md](0029-soulidentity-the-registry-dissolves.md) |
| soulidentity | 0014 | 0030 | [0030-soulidentity-the-cross-service-proof.md](0030-soulidentity-the-cross-service-proof.md) |
| soulidentity | 0015 | 0031 | [0031-soulidentity-the-vault-is-the-directory.md](0031-soulidentity-the-vault-is-the-directory.md) |
| soulidentity | 0016 | 0032 | [0032-soulidentity-one-noun-persona.md](0032-soulidentity-one-noun-persona.md) |
| soulrealm | 0010 | 0033 | [0033-soulrealm-fleet.md](0033-soulrealm-fleet.md) |
| soulnode | 0001 | 0034 | [0034-soulnode-genesis.md](0034-soulnode-genesis.md) |
| soulidentity | 0017 | 0035 | [0035-soulidentity-role-selection-by-name.md](0035-soulidentity-role-selection-by-name.md) |
| soulidentity | 0018 | 0036 | [0036-soulidentity-the-embed-seam.md](0036-soulidentity-the-embed-seam.md) |
| soulrealm | 0011 | 0037 | [0037-soulrealm-pinned-to-the-record.md](0037-soulrealm-pinned-to-the-record.md) |
| soulstream | 0008 | 0038 | [0038-soulstream-remote-mcp-node.md](0038-soulstream-remote-mcp-node.md) |
| soulidentity | 0019 | 0039 | [0039-soulidentity-soulfold-the-default-idp.md](0039-soulidentity-soulfold-the-default-idp.md) |
| soulnode | 0002 | 0040 | [0040-soulnode-the-composition-gate.md](0040-soulnode-the-composition-gate.md) |
| soulfold | 0001 | 0041 | [0041-soulfold-genesis-the-fold.md](0041-soulfold-genesis-the-fold.md) |
| soulnode | 0003 | 0042 | [0042-soulnode-first-boot-is-real.md](0042-soulnode-first-boot-is-real.md) |
| soulfold | 0002 | 0043 | [0043-soulfold-kv-schema-and-key-lifecycle.md](0043-soulfold-kv-schema-and-key-lifecycle.md) |
| soulnode | 0004 | 0044 | [0044-soulnode-the-realm-remembers.md](0044-soulnode-the-realm-remembers.md) |
| soulnode | 0005 | 0045 | [0045-soulnode-an-agent-runs.md](0045-soulnode-an-agent-runs.md) |
| soulfold | 0003 | 0046 | [0046-soulfold-session-and-ui-shape.md](0046-soulfold-session-and-ui-shape.md) |
| soulstream | 0009 | 0047 | [0047-soulstream-remote-mcp-node-built.md](0047-soulstream-remote-mcp-node-built.md) |
| soulstream | 0010 | 0048 | [0048-soulstream-the-node-becomes-consumable.md](0048-soulstream-the-node-becomes-consumable.md) |
| soulnode | 0006 | 0049 | [0049-soulnode-the-door-opens.md](0049-soulnode-the-door-opens.md) |
