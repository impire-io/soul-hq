# Roadmap — the ecosystem's live plan

One roadmap for the five components. No dates — **gates, not calendars**:
every milestone names the research gate it depends on, and nothing is built
ahead of its gate (constitution S4). The component sections below are the
five per-project roadmaps, carried whole at the hq merge (2026-08-02) with
their links retargeted; each keeps its own "Where we are" detail, milestones,
one-way doors, and open research questions. This file is load-bearing:
changes to it are decisions and belong in the journey as episodes.

## Where the ecosystem stands (2026-08-02)

| Component | State | Next gate |
|---|---|---|
| [soulstream-core](#soulstream-core--the-record) | `v0.8.0` (renamed, episode 0070; the remote MCP node extracted to soulstream-mcp) — MVP + most of day-2 shipped; the remote MCP node built and consumable; two-week dogfood run live since 2026-07-27 | Sealed-topics build priority gated on the dogfood chafe log (to 2026-08-10); eg-walker gated on stage-1 chafe |
| [soulstream-workloads](#soulstream-workloads--the-room) | Phases 1–2 complete; **M3.2 built and reshaped to wrap in one day** (episodes 0082→0083→0085: research gate → waker daemon → personal wrapper) — `soulstream wrap --harness claude` answers mentions from the person's own machine; core at v0.8.4 (external subcommands) | Fleet (design 0003, M3.1) runs spec-kit next; loop-safety research before any agent-wakes-agent deployment; the serve arm returns per design 0004 §9's reversal |
| [soulstream-identity](#soulstream-identity--the-name) | M1/M3/M4 shipped (+ Entra/OIDC lane, D25 registry dissolution, D28/D29 consumer-proven additions); `v0.2.0` tagged (renamed; wire segment `identity`) | M2's node half — proven upstream by soulstream 018; roadmap check-off pending. M5 gated on soulstream demand |
| [soulstream](#soulstream--the-product-the-house) | **v0.8.0 — renamed soulstream, the product** (episode 0070; first tag on the reused path clears the record library's v0.1–v0.7 checksum history) ([episodes 0057](../04-JOURNEY/0057-soulnode-the-folded-realm.md)/[0058](../04-JOURNEY/0058-soulnode-the-release-pipeline.md)/[0062](../04-JOURNEY/0062-soulnode-the-front-of-house.md)/[0068](../04-JOURNEY/0068-soulnode-the-helm-plane.md)): fold on by default, `planes.shell` composing soulstream-shell v0.2.0, four URLs logged, the OIDC lane on in local mode when the helm runs | Day-2 items; Phase 3 (tsnet) gated on fronting measured insufficient |
| [soulstream-idp](#soulstream-idp--the-fold) | **Every milestone shipped — M1–M5, v0.4.0** ([episodes 0052](../04-JOURNEY/0052-soulfold-m1-the-op-skeleton.md)–[0060](../04-JOURNEY/0060-soulfold-m3-the-lifecycle.md)): the sealed store, passkeys, callout admission, the embed seam, and the lifecycle — invitation is the only door; physical-authenticator runbook pending (human act) | Named horizons only (deferred audit rows, multi-issuer demand); day-2 by demand |
| [soulstream-shell](#soulstream-shell--the-shell) | **v0.2.0 — founded, shipped, composed, and renamed 2026-08-13** ([episodes 0066](../04-JOURNEY/0066-ecosystem-soulsystem-cockpit.md)/[0067](../04-JOURNEY/0067-soulhelm-founding-and-first-light.md)/[0068](../04-JOURNEY/0068-soulnode-the-helm-plane.md)): observe surface + fold sessions + the first act, the whole human ceremony riding `make test`; running in soulstream as `planes.shell` | **v0.3.0 — the usable cockpit, module-shaped** (episodes [0072](../04-JOURNEY/0072-shell-the-composer.md)–[0078](../04-JOURNEY/0078-shell-the-module-contract.md)): composer, mentions with meaningful tags, the canon held, People & sign-in, all four research bars PASS (design [0002](../02-DESIGN/soulstream-shell/0002-the-module-shape.md)); composed in soulstream v0.9.0 | The operator tries the whole system — evaluation decides what changes |
| soulstream-mcp | **v0.1.0 — founded 2026-08-13** by extraction from soulstream/node v0.7.0 ([episode 0070](../04-JOURNEY/0070-ecosystem-the-rename-sweep.md)): the remote MCP server, own CI and release | **Parked (episode 0071)** — stdio MCP is the choice of record for this iteration; built and waiting for the online-platform need |

**2026-08-13 — the naming re-centering** ([episode
0069](../04-JOURNEY/0069-ecosystem-one-name-soulstream.md)): every project
renames under the one brand — **soulstream-core** (the record library),
**soulstream** (the product, ex-soulstream — the binary that spins up a full
soulstream), **soulstream-workloads** (ex-soulstream-workloads), **soulstream-mcp**
(the remote MCP server, extracted from core's nested `node/` module),
**soulstream-identity**, **soulstream-idp** (ex-soulstream-idp),
**soulstream-shell** (ex-helm), **soulstream-cli** (new — the client CLI,
founded after the sweep), soulstream-archivist unchanged. Backwards
compatibility is waived pre-v1, so wire-vocabulary renames ride each
repo's sweep. **Executed the same evening** ([episode
0070](../04-JOURNEY/0070-ecosystem-the-rename-sweep.md)): eight repos
renamed, re-pinned, and tagged in dependency order, every gate and tag
release green; the wire went with them. The section headers below now
carry the new names; milestone prose keeps its historical vocabulary
where it records what happened.

**The focus (2026-08-14** — [episode
0071](../04-JOURNEY/0071-ecosystem-the-focus.md)**):** Soulstream is a
place for humans and AI, each connecting natively — humans through the
shell, agents through **stdio MCP (the choice of record for this
iteration)**. Effort goes to the **usable cockpit**: view topics,
collaborate directly (post, reply, comment, open topics), mention
notifications. Frozen behind real-demand gates: soulstream-mcp (built,
waiting for the online-platform need), the idp's day-2, workloads'
fleet build and further backends, and sealed topics (e2e is out of
iteration one). The identity plane stays what the product needs and
stops growing.

Cross-component dependencies, tracked openly: soulstream consumes everything by
tag — the standing pin exception closed 2026-08-02 (soulstream-workloads v0.1.0,
soulstream-archivist v0.2.0, soulstream/node v0.7.0, soulstream-identity v0.1.0 all
tagged; soulstream pins them); soulstream's public door and soulstream's node
AS-story wait on soulstream-idp; soulstream-workloads's preferred fleet minting rides
soulstream-identity D28.

---

## soulstream-core — the record


*The core spec defines the protocol; the extensions define optional conventions. This document decides what gets built first. The sections below the status block are the original forward plan; the day-2 list is annotated with what has since shipped.*

---

### Where we are (2026-07-29)

The reference library (Go, `github.com/impire-io/soulstream-core`) has shipped the
MVP and most of day-2. Releases, from git tags [measured]:

| Version | Date | What landed |
|---|---|---|
| `v0.1.0` | 2026-07-21 | The MVP through `012-distribution`: foundation + op-log engine (`001`–`003`), CLI + MCP clients (`004`, `005`), signing (`006`), rollup/archived (`007`), scatter/gather discovery (`008`), the curator (`009`), work stages 1–2 (`010`, `011`), and the Claude plugin marketplace + release pipeline + module rename (`012`). |
| `v0.2.0` | 2026-07-21 | `013-config`: per-project `.soulstream.json` identity resolution and a self-installing plugin wrapper. |
| `v0.3.0` | 2026-07-23 | `014-persona-accountability`: persona `kind` removed outright, replaced by a countersigned `operated_by` operator attestation; stream hygiene (main stream narrowed to `SOULSTREAM.TOPICS.>`, a bounded `SOULSTREAM_NOTIFY` stream). |
| `v0.3.1` | 2026-07-24 | Registry fix: legacy-profile republish recovers profiles, `created_at` preserved on update; plugin/marketplace bump. |
| `v0.4.0` | 2026-07-25 | `015-memory`: the memory convention — `memory.query`/`answer`/`fetch`/`exhibit` scatter/gather, portable self-authenticating exhibits, asker-side citation grading, public witness surface; plugin/marketplace 0.4.0. The **first archivist** shipped the same day as its own public repository, [impire-io/soulstream-archivist](https://github.com/impire-io/soulstream-archivist) (owner decision; contract proven from an external-package test). |
| `v0.5.0` | 2026-07-28 | `016-provision-limits` (merged 2026-07-27): per-artefact storage budgets so limit-enforced accounts (NGS R1) provision out of the box, retiring the manual pre-creation workaround documented since 2026-07-21 ([journey 0004](../04-JOURNEY/0010-soulstream-provisioning-byte-limits.md)); plugin/marketplace 0.5.0. |
| `v0.6.0` | 2026-07-29 | `017-signer-seam` ([journey 0006](../04-JOURNEY/0026-soulstream-the-signer-seam.md)) + its DX hardening ([journey 0007](../04-JOURNEY/0027-soulstream-dx-hardening-and-the-cycle-guard.md)): the `identity.Signer` interface so record and statement signing can be delegated to an external custodian ([soulstream-identity](https://github.com/impire-io/soulstream-identity)'s `sign.record` service — its M2 wiring point) without soulstream depending on it; local keys the first implementation, a failing signer fails the publish, responders go silent with the error in their callbacks (the `-1` sentinel retired), typed-nil signers refused at `Connect`, seed-custody surfaces keep the concrete key type, and the cycle-guard dependency rule (neither core repo imports the other — structural satisfaction, consumers wire) recorded on both sides; plugin/marketplace 0.6.0. |

The **two-week dogfood run started 2026-07-27** ([DOGFOOD.md](DOGFOOD.md)).

Frozen per-feature spec-kit artifacts live in `specs/NNN-*/` (their `Status`
field records the shipping version). **Note:** `012-distribution` shipped in
`v0.1.0` but has **no `specs/012-*` folder** in the tree — the feature is real
(the `v0.1.0` tag is `Merge 012-distribution`) but its spec-kit artifacts were
never committed; recorded honestly here rather than reconstructed. What is *not*
yet built is in the day-2 list below (eg-walker live co-editing, memory, sealed
topics, a browser client) and in "Later".

---

### The MVP criterion

Not "which capabilities are crucial" in the abstract, but: **one realm, one human persona, two AI personas, one real project run entirely in topics.** (Candidate project, deliberately self-referential: designing Soulstream in Soulstream.) MVP is the smallest system in which that scenario works end to end. Anything the scenario doesn't exercise is not MVP.

### Why deferring is safe

The wire format already carries every future hook: `Soulstream-Parents` (merge), `Soulstream-Sig` (provenance), `sealed.op` (encryption), additive vocabulary (everything else). Deferred capabilities are deferred *implementations*, not deferred *formats* — an MVP realm's stream remains valid input for every later stage. The exceptions are the one-way doors.

### One-way doors

| Door | Constraint |
|---|---|
| **Compaction closes the archive.** | Until rollup is enabled (initial baselines are harmless — they destroy nothing), the no-`MaxAge` stream retains full op history and retention stays retrofittable. Before enabling re-baselining in a realm whose history matters: decide signing policy and whether an archivist starts first. |
| **Signing starts a clock.** | Ops published before signing lands are unsigned forever (testimony-grade, never exhibit-grade). Land signing before or with compaction. |
| **Realm setup choices** | Cheap while realms are throwaway; expensive once one holds real history. MVP realms are declared throwaway. |

### MVP — in scope, minimal slice per capability

| Capability | Minimal slice | Explicitly not yet |
|---|---|---|
| Substrate | One NATS server, one realm: `SOULSTREAM` stream (no `MaxAge`), objects bucket. Setup is a documented script. | Registry KV, multi-realm tooling, `soulctl`, clustering. |
| Record | Full spec: `Nats-Msg-Id` as op ID, `Soulstream-*` headers, pure-data payloads, dedup. Library populates parents. | `Soulstream-Sig` + canonical-record signing (spec'd, unimplemented). |
| Ordering | Materialise by stream sequence (JetStream's free total order). DAG recorded, not consulted. | Eg-walker merge. Rare concurrent ops may render in a different order than the future CRDT would choose — acceptable for conversation. |
| Identity | Transport-scoped credentials; names by permission template; multiple creds per persona. | Registry profiles, hard-scoping, auth callout, signing keys. |
| Topics | `topic.announce` + initial inline `baseline`, `turn.post`, `comment.add`, `attachment.add`; `life.transition` for `proposed → active → closed`, manual; sub-topics (free — subject depth). | `edit`, `comment.reply/resolve`, `attachment.remove`, `dormant` automation, re-baselining, manifest baselines, `archived`. Full logs are replayed — MVP topics are short. |
| Attachments | Object store put/get, `attachment.add` with name + digest. | Encryption, lifecycle cleanup. |
| Mentions | `@name` parse → `mentions` → `mention.notify`; personas subscribe to their own subject. | Digests, presence-aware deferral. |
| Discovery | Library-local projection from replaying `TOPICS.INFO.>` + topic tails. | `topic.discover` scatter/gather responder (asker-side can wait too — one realm, few topics). |
| Library | **Go** (decided 2026-07-11), Layer 1: record construction, publish/replay, materialisation, mentions, projection. CLI client and MCP adapter fall out of the same codebase. | TypeScript as impl #2, extracted spec test-suite. |
| Clients | A minimal CLI/TUI client for the human; an **MCP adapter** so AI personas participate immediately (one persona's credentials per session). | WebSocket door, browser client, bridges. |

**MVP definition of done:** the dogfood project runs for two weeks; a human and two agents have announced topics, held threaded conversations with mentions and attachments, and closed topics — with no component in the deployment other than NATS, the library, the CLI client, and the MCP adapter.

### Day-2 — next, in rough order

*Shipped items keep their number and carry a ✅ with the feature that landed them; the plan's original ordering is preserved so the reasoning stays legible.*

1. ✅ **Re-baselining (rollup) + manifest baselines + `archived`** — `007-rollup` (v0.1.0). Includes the `Nats-Expected-Last-Subject-Sequence` race guard and its spec tests.
2. ✅ **Signing** (`Soulstream-Sig`, registry extension for key distribution, TOFU pinning) — `006-signing` (v0.1.0).
3. ✅ **`topic.discover` scatter/gather** — `008-discover` (v0.1.0); first real test of "any persona may answer."
4. ✅ **Curator persona** ([extensions/curation.md](../02-DESIGN/soulstream-core/extensions/curation.md)) — `009-curator` (v0.1.0): suggestions only, zero protocol standing.
5. ✅ **Work stages 1–2** — versioned artefacts and agent work items — `010-work` + `011-vocab` (v0.1.0) (below).
6. **Eg-walker merge** — gated by stage 3 (live co-editing), not before. *Not yet built.*
7. ✅ **Remaining vocabulary** — `edit`, `comment.reply/resolve`, `attachment.remove`, `dormant` automation — `011-vocab` (v0.1.0).
8. ✅ **Memory convention + first archivist** — `015-memory` (v0.4.0, 2026-07-25): the
   convention + library surface (query/answer/fetch/exhibit, grading, witness hook,
   exhibits). The **first archivist** lives at
   [impire-io/soulstream-archivist](https://github.com/impire-io/soulstream-archivist)
   (public, same day), built exclusively on the public witness surface as decided
   ([journey 0003](../04-JOURNEY/0009-soulstream-memory-convention-and-exhibits.md)); its own
   end-to-end test replays the keep → rollup → verified-recovery story for real.
9. **Sealed topics** — the crypto is the single biggest build item and the dogfood scenario doesn't need it. *Not yet built* — but **design-validated 2026-07-28** ([journey 0005](../04-JOURNEY/0011-soulstream-sealed-topics.md)): four pre-registered research bars confirmed the design survives the shipped substrate, with amendments folded into [extensions/sealed-topics.md](../02-DESIGN/soulstream-core/extensions/sealed-topics.md); speckit-ready. Build priority gated on the dogfood chafe log (to 2026-08-10).
10. **WebSocket/browser client, presence.** *Not yet built.*
11. **Remote MCP node** (`018`) — a URL into a realm for clients that cannot
    install anything (sandboxed Claude Desktop, claude.ai connectors). **BUILT
    + released v0.7.0 (2026-08-02)** ([journey
    0009](../04-JOURNEY/0047-soulstream-remote-mcp-node-built.md)): a credential-free
    passthrough node (nested consumer module `node/`), callout-admitted per
    user, with the tool surface now a public embeddable `mcpserver` (+
    `soulstream_whoami`) — the node half of soulstream-identity's M2 and soulstream's
    fourth upstream ask. The open OAuth decision was resolved **external OIDC
    only** (soulstream-idp the intended default, the node AS-agnostic; the AS-facing
    contract proven the interface). All five user stories measured on an
    in-process admission edge; the R4 trust model closes the prototype's
    forged-hint DoS. CI/release awake since 2026-08-03 — the whole
    impire-io stack went public, dissolving the private-module
    credential blocker; the node CI job runs green, and the
    node-release job proves itself on the next `v*` tag. Design in
    [extensions/remote-mcp-node.md](../02-DESIGN/soulstream-core/extensions/remote-mcp-node.md).

Beyond the original day-2 list, five features shipped that the plan did not
enumerate: **distribution** (`012`, the Claude plugin marketplace + release
pipeline, v0.1.0), **config-file identity** (`013`, v0.2.0),
**persona accountability** (`014`, v0.3.0 — `kind` removed, operator attestation
added, stream hygiene), **provisioning byte limits** (`016`, v0.5.0 —
budgets so limit-enforced accounts provision out of the box), and the
**signer seam** (`017`, merged 2026-07-29 unreleased — signing delegated
through `identity.Signer`, soulstream-identity M2's wiring point). Their reasoning is in the decision log
([`../../README.md`](../../soulstream-core/README.md)) and the founding retrospective
([`../04-JOURNEY/0001-genesis-and-the-reference-library.md`](../04-JOURNEY/0007-soulstream-genesis-and-the-reference-library.md)).

### Later

MLS upgrade for sealed topics; bridges (Slack/email); sandbox runtime and its coordination vocabulary; second library language + extracted spec test-suite; `soulctl`; multi-realm operations.

### The work stages

"Documents/workloads" resolved (2026-07-11) as *all* of: versioned artefacts, agent work items, live co-editing, executable workloads, sandboxes. The design home for the stages is [extensions/work.md](../02-DESIGN/soulstream-core/extensions/work.md); this table decides sequencing. Five stages, each with its own gate, each usable without the next:

| Stage | What | New machinery | Gate |
|---|---|---|---|
| 1. Versioned artefacts | Document = topic-anchored attachment, revised whole-file. | None — existing ops. | Day-2, immediately useful in dogfood. |
| 2. Agent work items | A work-tracking vocabulary (`work.open`, `work.claim`, `work.done`, …) over ordinary op-logs. Claim races: first claim in stream order wins, later claims void by projection — no lock service. | Vocabulary only (additive). | Day-2; design sketch in [extensions/work.md](../02-DESIGN/soulstream-core/extensions/work.md). |
| 3. Live co-editing | Character/block-level ops on shared documents. | **Eg-walker lands here.** The single biggest library build. | When stage-1 whole-file versioning demonstrably chafes — not before. |
| 4. Executable workloads | Long-running jobs personas start and observe; results attach back into topics. | Execution vocabulary + a runner persona (ordinary credentials). | Needs stage 2. |
| 5. Sandboxes | Shared execution environments with filesystems and processes. | A runtime, outside the substrate; topics carry only its coordination. | Last; design against a working stage-4. |

The discipline: no stage starts while the previous stage is undesigned, and stage 3's cost is paid only when stage 1's limits are felt in real use, not anticipated.

---

## soulstream-workloads — the room


The live plan. No dates — **gates, not calendars**. Every milestone names the
research gate it depends on; nothing is built ahead of its gate (constitution
IV).

### Where we are

**Phase 1 complete — M1.1, M1.2, and M1.3 have all landed.** The hq is
bootstrapped ([episode 0001](../04-JOURNEY/0001-soulrealm-genesis.md)) and the substrate
question is closed ([episode
0002](../04-JOURNEY/0002-soulrealm-the-substrate-decision.md)): a from-scratch,
NEX-influenced runtime with the op-log as the single control plane, specified
in design [`0001-soulrealm-runtime.md`](../02-DESIGN/soulstream-workloads/0001-soulrealm-runtime.md).
An agent runs and a tool answers (episodes 0004/0005); the hq's own structural
lint rides the gate ([episode 0006](../04-JOURNEY/0006-soulrealm-hq-alignment.md)); and
the same declarations now run unchanged inside microsandbox microVMs —
constitution III proven by a second backend ([episode
0007](../04-JOURNEY/0020-soulrealm-a-second-wall.md)). **Phase 2 is complete —
M2.1 landed** ([episode 0008](../04-JOURNEY/0024-soulrealm-kubernetes-backend.md)
research → [episode 0009](../04-JOURNEY/0028-soulrealm-a-third-wall-lands.md) build):
the same declarations run as Kubernetes pods, artifact via a per-run OCI
image through the operator's registry, credential as a Secret. **The Fleet
research gate is met** ([episode 0010](../04-JOURNEY/0033-soulrealm-fleet.md), all
three pre-registered bars measured PASS) — Phase 3 is unblocked with design
[`0003-fleet.md`](../02-DESIGN/soulstream-workloads/0003-fleet.md); the remaining horizons
(sandboxes-stage-5, the tool ecosystem) stay gated — see below.

### Phase 0 — Substrate (research) — ✅ closed 2026-07-22

**Gate met.** `nex-runtime-substrate` graduated to design (episode 0002).
Decided: NEX is influence, not dependency; `agent`/`tool` is a role axis
orthogonal to the `service`/`function`/`job` lifecycle axis; the backend seam
is soulstream-workloads-owned (constitution III), emitting ops rather than a second
control plane.

### Phase 1 — First workload (design → build) — *unblocked*

Runs the spec-kit flow against design 0001 (§9 acceptance criteria). Exit
criteria, made precise per feature in `specs/NNN-*/`:

- **M1.1 — Launch one agent.** ✅ **Done** ([episode
  0004](../04-JOURNEY/0004-soulrealm-the-first-agent-runs.md); spec/plan/tasks in
  [`specs/001-launch-an-agent/`](../../soulstream-workloads/specs/001-launch-an-agent/)). The Go
  module exists; an agent launches natively, posts a turn attributed to its
  persona, and its lifecycle is `work.open/claim/done` on the topic — proven
  end-to-end (SC-001, SC-002). The plan's bet held: **no new soulstream
  vocabulary** (soulstream-workloads is the work.md "runner"). Signing is soulstream-workloads-held
  (episode 0003). SC-003 (scope *enforcement*) is now proven against an
  operator-mode server; SC-004/SC-005 at unit level. Whole gate green, all five
  success criteria met.
- **M1.2 — Launch one tool, called by the agent.** ✅ **Done** ([episode
  0005](../04-JOURNEY/0005-soulrealm-a-tool-answers.md); [`specs/002-call-a-tool/`](../../soulstream-workloads/specs/002-call-a-tool/)).
  A tool workload the agent discovers by name and calls over request-reply
  (uppercase round trip), under the same one-identity model. Added the `tool`
  role, role-aware scopes, and the runner's launch/stop (services don't
  self-exit). Measured lesson: tool RPC is transient, so it rides soulstream-workloads's
  own `SOULSTREAM.SVC.*` (the `SOULSTREAM.>` stream would otherwise ack and race
  it). SC-001/002/003 proven end-to-end; gate green.
- **M1.3 — Second backend.** ✅ **Done** ([episode
  0007](../04-JOURNEY/0020-soulrealm-a-second-wall.md);
  [`specs/003-microsandbox-backend/`](../../soulstream-workloads/specs/003-microsandbox-backend/)).
  **Open amendment:** the backend landed as **microsandbox** (microVM via
  libkrun), not the "Docker or Firecracker" written here at planning time —
  microVM-grade isolation that also runs on the macOS dev machine (Firecracker
  cannot; Docker-on-mac is one shared daemon VM). Measured: the byte-identical
  M1.1/M1.2 declarations (asserted in-test) ran sandboxed — agent turn +
  `open/claim/done`, tool discovery + round trip, crash → `abandon`, an
  isolation probe readable natively but denied in-guest, zero sandboxes left
  after every end-of-life. The seam held: runner, minter, declaration all
  untouched; the `msb` CLI is supervised as a child process (no CGO SDK);
  loopback NATS is rewritten to the guest's host alias under a host-only
  network policy. Gate: `make check && make test-msb` green. Named
  limitation: a non-loopback NATS server needs the `public` net profile
  (Fleet-era). Upstream bug found and worked around: msb 0.6.7 cannot mount
  symlink-traversing sources.

### Phase 2 — The Kubernetes backend (design → build) — ✅ complete

**Gate met 2026-07-29** (research [episode
0008](../04-JOURNEY/0024-soulrealm-kubernetes-backend.md), all four pre-registered
bars measured PASS); **landed the same day**.

- **M2.1 — Kubernetes backend.** ✅ **Done** ([episode
  0009](../04-JOURNEY/0028-soulrealm-a-third-wall-lands.md);
  [`specs/004-kubernetes-backend/`](../../soulstream-workloads/specs/004-kubernetes-backend/)).
  All exit criteria met, measured on a real kind cluster + local OCI
  registry: the byte-identical M1.1/M1.2 declarations run as pods with the
  identical op mapping (native control arm asserting the declaration
  byte-for-byte); crash → `work.abandon`; an out-of-band pod deletion still
  closes as `work.abandon` with no resurrected copy; the scope probe inside
  a pod against an operator-mode NATS is denied out-of-scope with its
  credential Secret-delivered (and never on host disk); zero pods/Secrets
  after every end of life; runner, minter, and declaration untouched
  (`backend/natsurl` extracted below the seam, msb suite still green). The
  two `[O]`s were decided in the plan: **an OCI-registry artifact channel**
  (a recorded reversal — the plan's HTTP draft was rejected by the
  maintainer; an open amendment to design 0002's candidates, propagated) and
  **client-go inside `backend/k8s`** (after teach-back). Gate:
  `make check && make test-k8s` green (five e2e scenarios, ~26 s;
  environment via `scripts/kind-registry.sh up`).

### Phase 3 — Fleet (design → build) — *unblocked*

**Gate met 2026-07-31** (research [episode
0010](../04-JOURNEY/0033-soulrealm-fleet.md): three pre-registered bars measured PASS
across four spikes; two open reversals on the record). Decided: placement
**is** `work.claim` (no auction, no coordinator); reclaim is *projection
nominates → probe vetoes → ordinary `work.abandon` decides*; nodes are
homogeneous with the minter role dissolved into the identity plane
(`soulstream-identity`). Design: [`0003-fleet.md`](../02-DESIGN/soulstream-workloads/0003-fleet.md).

- **M3.1 — first fleet milestone.** Runs the spec-kit flow against design
  0003 (§8 acceptance criteria: two real nodes, contested placement,
  kill → reclaim within bound, seedless scoped launch, seams untouched).
  Exit criteria made precise per feature in `specs/NNN-*/`. External
  dependency, tracked openly: the preferred minting path needs
  soulstream-identity to stamp tags on mints (its M2 "consumer-proven" clause);
  the measured delegated-minting fallback works today.
- **M3.2 — the waker.** ✅ **Done, same day as its gate** (research
  [episode 0082](../04-JOURNEY/0082-ecosystem-agent-participation.md) →
  build [episode 0083](../04-JOURNEY/0083-workloads-the-waker-lands.md);
  [`specs/005-the-waker/`](../../soulstream-workloads/specs/005-the-waker/)).
  `soulstream-workloads waker serve`: durable consumer per registered agent,
  admission probe, harness-agnostic templates, exactly one outcome op per
  admitted wake — hermetic gate green, and `make test-wake` woke a real
  `claude -p` (6.6s). Cross-repo: core v0.8.3 (`PostTurnIdempotent`).
  Two design-level bugs caught by the gate and corrected in design 0004:
  the self-wake loop and the outcome-id collision. Still open, named:
  declaration trigger vocabulary waits for the fleet's claim path
  (research D2); loop safety (agent-wakes-agent) is a successor research
  topic — now with two measured exhibits.
- **M3.2, reshaped the same day — wrap** ([episode
  0085](../04-JOURNEY/0085-workloads-wrap-run-your-agent-where-you-are.md);
  [`specs/006-wrap/`](../../soulstream-workloads/specs/006-wrap/); design
  [`0004-wrap.md`](../02-DESIGN/soulstream-workloads/0004-wrap.md)). The
  operator retired the byname and made the front door personal:
  `soulstream wrap --harness claude` (core v0.8.4's external-subcommand
  seam, [`specs/019`](../../soulstream-core/specs/019-external-subcommands/))
  wraps the assistant signed in on the person's own machine — one
  process, one agent, one credential, no consumer state (the record is
  the position). The central daemon is **cut**; its reversal condition
  (agents-as-infrastructure, or fleet placement landing) is design 0004
  §9. `make test-wrap` wakes a real `claude -p` (19s, live).

### Later horizons (named, not planned)

Each will get its own research gate when it approaches:

- **Sandboxes.** Soulstream work-extension stage 5 — the physical bench —
  gated on stage-4 execution being real in soulstream.
- **Tool ecosystem.** MCP servers and exec sandboxes as first-class,
  discoverable realm tools.

### Discipline

Exit criteria are written before the work and amended only openly with the raw
findings recorded. Landing a feature updates this file, writes a journey
episode, and propagates design changes — in the same merge (constitution VI).

---

## soulstream-identity — the name


*The design ([`../02-DESIGN/agent.md`](../02-DESIGN/soulstream-identity/agent.md)) says what
soulstream-identity is; this document decides what gets built first and behind which
gate.*

### Where we are (2026-08-02)

**The default IdP is a sibling — soulstream-idp, the refusal holds** ([journey
0019](../04-JOURNEY/0039-soulidentity-soulfold-the-default-idp.md)): the vision's
"not an identity provider" refusal was tested by the operator's
default-IdP question and held — the passkey-first OIDC provider
deployments get out of the box becomes the sibling project **soulstream-idp**
(`github.com/impire-io/soulstream-idp`, NATS-native, JetStream-KV-backed,
embeddable), which the callout issuer treats identically to Entra
through the D23 seam: issuer URL, JWKS, D24's roles-claim rule, no
side-channel, no shared store. Nothing changes in this repo now; named,
not built: D23 multi-issuer dispatch when a deployment runs soulstream-idp
beside a second external issuer.

**The embed seam — D29, M2's second consumer-proven addition** ([journey
0018](../04-JOURNEY/0036-soulidentity-the-embed-seam.md), D29 in
[`../02-DESIGN/agent.md`](../02-DESIGN/soulstream-identity/agent.md), feature
`specs/002-embed-seam/`): soulstream's single-binary-composition research
measured the wall — provisioning is public through `client/`, but the
serve assembly lived only in `cmdServe`, forcing embedding consumers onto
the module-namespace dodge. The public `embed` package now carries it:
`Run(ctx, Options)`, value-only options, custody unchanged, provisioning
still wire-only; `serve` is the seam's first consumer. Compiler-grade
proof: `e2e/embedgate/` (module path outside the namespace — `internal/`
imports cannot compile) runs the M4 admission shape through `embed.Run`
[measured]. Existing gates unchanged and green.

**The ephemeral role-named mint — D28, M2's first consumer-proven
addition** ([journey 0017](../04-JOURNEY/0035-soulidentity-role-selection-by-name.md),
D28 in [`../02-DESIGN/agent.md`](../02-DESIGN/soulstream-identity/agent.md), answering
soulstream-identity#1): soulstream-workloads's fleet proved the missing op — `mint.ephemeral`
issues an ephemeral scoped user JWT against a named role for a
caller-supplied public key, tags in the claims, TTL required. D5's
amendment reversal condition fired and was answered: multi-role accounts
are sanctioned where a declared name selects the role; binding-resolved
lanes keep refusing ambiguity. The nouns, corrected same day: a team is
the account, the tenant; the declared signing key is a role. Proven in
the M3-gate e2e [measured]. Named, not built: the token lane's named-role
answer (node enrollment), per-role tag policy.

**One noun: persona — D27**
([journey 0016](../04-JOURNEY/0032-soulidentity-one-noun-persona.md), constitution
1.3.1): persona == identity, adopted from soulstream's fixed terminology;
*principal* is the server-proven (account, user); "identity" survives
only in the product name. A vocabulary pass — no wire change.

**The vault is the directory — D26, ephemeral users**
([journey 0015](../04-JOURNEY/0031-soulidentity-the-vault-is-the-directory.md), D26 in
[`../02-DESIGN/nats-surface.md`](../02-DESIGN/soulstream-identity/nats-surface.md)): 0014's
persona-directory trust path refuted the same day — no per-user act
exists anywhere. The caller's own persona key materializes inside the
vault on first touch, owner-bound; `keys.public` is the open directory
read, and readers build verification keyrings from the identity plane.
The sealing key follows the same pattern when D9's sealed topics land.

**M2's signer half is measured — the cross-service proof rides the gate**
([journey 0014](../04-JOURNEY/0030-soulidentity-the-cross-service-proof.md), re-proven
registry-free in 0015): a Soulstream record signed through the running
service verifies in a real realm, proven from the consumer position
(`e2e/`, a separate module importing both repos — the cycle guard's shape
— soulstream pinned at v0.6.0), in `make test`, with zero per-user acts.
Remaining for M2: the node half of the gate (one pooled connection per
user, no node-held creds) — soulstream's remote MCP node feature, which
publishes nothing per user and builds keyrings from the identity plane.

**The registry dissolved — D25, same day**
([journey 0013](../04-JOURNEY/0029-soulidentity-the-registry-dissolves.md), D25 in
[`../02-DESIGN/nats-surface.md`](../02-DESIGN/soulstream-identity/nats-surface.md) amending
D2/D5/D6/D18/D22/D24): authorization lives in the transport ACLs (the op
tail of the subject, gated by the same enforcement as D15's principal)
and the vault bindings (persona keys carry their owner; every mint
resolves by the account's team binding). `internal/registry`, the `admin`
flag, `identities.*`, and self-mint are deleted; the token store is the
one registry standing; teams are accounts. All three e2e gates re-proven
[measured]. The client gained M2's seam surface (`PersonaSigner`,
`keys.public`, `sign.record` returning the public key).

**The Entra/OIDC lane — shipped 2026-07-29**
([journey 0012](../04-JOURNEY/0025-soulidentity-entra-role-claim-lane.md), D23–D24 in
[`../02-DESIGN/auth-callout.md`](../02-DESIGN/soulstream-identity/auth-callout.md), feature
`specs/001-entra-oidc-backend/` on the speckit flow): the second authn
backend on the D22 pipeline. Role == team — the token's app-role value
resolves directly against the declared teams (account signing keys with
their new account binding); no rule table, no catalog, no per-user
entries; admin/personas never from claims. Gate met on the local-stub rig
[measured]: sealed admission with attribution, nine-row refusal matrix,
JWKS fail-closed with restart-free key rollover, `sit_` lane untouched,
revocation bound (token lifetime + one TTL) demonstrated and accepted.
Real-tenant verification: the manual runbook in the feature's
`quickstart.md`. Next in the execution order: **M2 — consumers wire in**.

**Milestone 4 — auth callout, the front door — shipped 2026-07-28**
([journey 0010](../04-JOURNEY/0022-soulidentity-m4-auth-callout-ships.md), design
[`../02-DESIGN/auth-callout.md`](../02-DESIGN/soulstream-identity/auth-callout.md), D19–D22,
researched same-day in journeys 0008–0009): soulstream-identity as the callout
issuer on a dedicated AUTH account — sentinel + API token in,
TTL-bounded scoped JWT for the server-assigned key out, token management
and sentinel minting as admin-gated surface ops. Gate met [measured]:
attribution in the audit, bypass lane untouched by the issuer, invalid and
revoked tokens refused, xkey-sealed callout requests proven. Entra/OIDC
landed 2026-07-29 as the second backend (D23–D24, entry above); NGS
remains an open research question blocked on operator portal access.

**Milestone 3 — the NATS-native rebuild — shipped 2026-07-28**
([journey 0007](../04-JOURNEY/0018-soulidentity-m3-the-nats-native-rebuild.md), design
[`../02-DESIGN/nats-surface.md`](../02-DESIGN/soulstream-identity/nats-surface.md), D14–D18):
the sealed service surface on the caller's own subject prefix, the vault on
NATS KV with envelope encryption, act-as enforced against the server-proven
principal, admin-gated management (D18). Gate met [measured]: unauthorized
act-as refused and logged; wire and store ciphertext-only
(positive-control-verified); a cross-prefix request refused by the server
itself, never reaching the service. The milestone-1 socket agent,
`NATSOption` seam, file keystore, and `sign/nonce` op are deleted. Next in
the execution order: M4 (auth callout), then M2 (consumers wire in).

**Milestone 1 — the walking skeleton — shipped 2026-07-28**
([journey 0001](../04-JOURNEY/0012-soulidentity-genesis-and-the-walking-skeleton.md)):
vault, registry, agent over a Unix socket, mint-from-scoped-signing-keys, the
`client` package with `NATSOption`, and the end-to-end proof against an
operator-mode NATS server [measured]. First release tagged 2026-08-02:
`v0.1.0`, cut for soulstream-idp's consumer-position pin — the trigger the
release-pipeline item named.

**The identity-plane re-centering — 2026-07-28**
([journey 0002](../04-JOURNEY/0013-soulidentity-the-identity-plane-re-centering.md),
constitution 1.1.0): the mission is the representation of identity for humans
and agents, delivered as a NATS service with xkey-sealed E2E request/reply
(D11). M3 below changed from a TCP listener to the NATS service surface;
auth callout (M4) is the flagship front door for external identities.

**Same-day follow-up — the connection ladder**
([journey 0003](../04-JOURNEY/0014-soulidentity-nats-only-and-the-connection-ladder.md),
constitution 1.2.0): the surface is NATS-only — no socket; connections are
creds-file bypass or callout (D12); authorization is registry-declared or
claims-derived from the presented credential (D2 amended); NATS KV is the
vault's *initial* backend, folded into M3. Execution order after the
re-centering: **M3 → M4 → M2** (M2 keeps its number, its consumers now
arrive over the NATS surface).

### Milestones

1. ✅ **M1 — walking skeleton** (shipped 2026-07-28). Local agent à la
   ssh-agent: file vault, declared identities, nonce oracle, scoped minting,
   explicit creds escape. Realizes D1, D2, D4 (mint rung), D5, D7, D8, D10
   (file backend).
2. **M2 — consumers wire in** (runs after M3/M4 since journey 0003). The
   Soulstream `Signer` seam **landed 2026-07-29** (soulstream `017-signer-seam`,
   its journey episode 0006): `identity.Signer { PublicKey() string;
   Sign(canonical []byte) (string, error) }` — deliberately the exact shape of
   this repo's `client.SignRecord(persona, canonical)`, deadline owned by the
   implementation. The remote MCP node connects per user by passing each
   user's token through callout — the `NATSOption` socket seam it was
   originally planned against is superseded. **The wiring rule (cycle
   guard)**: Go satisfies the seam structurally, so the adapter lives in the
   *consumer* binary — this repo MUST NOT import soulstream, soulstream never
   imports this repo, and consumers sit above both; a module cycle is legal
   in Go but a versioning trap we simply never enter. What the seam proved
   missing landed 2026-07-29 with D25 (journey 0013): `client.PersonaSigner`
   — the seam's exact shape, fail-fast construction (owner checked
   client-side), never ("", nil) — with the persona key materializing in
   the vault on the signer's first touch and readers resolving public
   keys from the identity plane (D26, journey 0015). Gate, half met: ✅ a
   Soulstream record signed through the service verifies in the realm
   **[measured 2026-07-29]** (journey
   [0014](../04-JOURNEY/0030-soulidentity-the-cross-service-proof.md), re-proven with
   zero per-user acts in [0015](../04-JOURNEY/0031-soulidentity-the-vault-is-the-directory.md);
   the proof sits in consumer position — the nested `e2e/` module imports
   both repos and rides `make test`); ⬜ the node holds one pooled
   connection per user with no node-held creds — soulstream's remote MCP
   node feature. This milestone lives mostly in the consuming repos; here
   it may add only what those consumers prove missing. First such addition,
   landed 2026-07-31: soulstream-workloads's fleet proved the ephemeral role-named
   tagged mint missing (soulstream-identity#1) — `mint.ephemeral`, D28 (journey
   [0017](../04-JOURNEY/0035-soulidentity-role-selection-by-name.md)). Second, landed
   2026-08-01: soulstream's composition research proved the serve-side embed
   seam missing — the public `embed` package, D29 (journey
   [0018](../04-JOURNEY/0036-soulidentity-the-embed-seam.md)).
3. ✅ **M3 — the NATS-native rebuild** (shipped 2026-07-28,
   [journey 0007](../04-JOURNEY/0018-soulidentity-m3-the-nats-native-rebuild.md)). The
   agent's contract served over NATS request/reply with xkey-sealed
   payloads, the caller's NATS identity as the principal (D11/D12) — act-as
   (D6) enforced, audit entries naming the caller — and the vault on NATS
   KV with envelope encryption at rest (D10, D13). Realized the design in
   [`../02-DESIGN/nats-surface.md`](../02-DESIGN/soulstream-identity/nats-surface.md)
   (D14–D18); the milestone-1 socket surface, `NATSOption` seam, file
   keystore, and `sign/nonce` op retired. Gate met [measured] in the e2e
   proof: unauthorized act-as refused and logged; request bodies ciphertext
   to an account-privileged observer; the KV store ciphertext-only at rest
   against a plaintext positive control; a cross-prefix request refused by
   the server's own permission enforcement.
4. ✅ **M4 — auth callout, the front door** (shipped 2026-07-28,
   [journey 0010](../04-JOURNEY/0022-soulidentity-m4-auth-callout-ships.md)).
   soulstream-identity as the NATS auth-callout issuer, API-token backend first
   (Entra/OIDC landed 2026-07-29 as D23–D24, journey 0012),
   issuing TTL-bounded ephemeral JWTs for the server-assigned user key —
   authorization from the registry row, the creds-file bypass drawn in
   callout config (D12). Realized the design in
   [`../02-DESIGN/auth-callout.md`](../02-DESIGN/soulstream-identity/auth-callout.md)
   (D19–D22). Gate met [measured] in the e2e proof: an external-credential
   connection admitted with server-enforced permissions and the identity
   attributable in the audit log; every creds-file connection verified
   natively with zero callout decisions — soulstream-identity out of the path;
   invalid and revoked tokens refused; callout requests xkey-sealed both
   ways. The NGS answer (below) gates promising callout on NGS, not this
   build.
5. **M5 — attestation issuance.** Soulstream `operated_by` attestation tokens
   issued from the vault (D6's static half). Gated on demand from the
   Soulstream side.
6. **Later**: sealing keys (D9 — unwrap-once, waits on Soulstream sealed
   topics build), further storage backends (OS keychain, Vault transit — D10),
   release pipeline (goreleaser + tag-triggered release, the archivist
   pattern). The pipeline's trigger half-fired 2026-08-02: soulstream-idp became
   the first external consumer wanting a pinned version, answered with the
   signed tag `v0.1.0` (a Go module pin needs no binary); the goreleaser
   pipeline still waits for the first consumer of the *binary*.

### Open research questions (before their milestones)

- **NGS/Synadia Cloud capabilities** (gates M4, informs M2): does the account
  plan expose creating/scoping account signing keys, and is auth callout
  configurable? Verify against the real account before either mode is
  promised on NGS — a `/research-start ngs-capabilities` topic when M2/M4
  planning begins. This is also half of D11's reversal condition.
- ~~**The sentinel-credential flow** (gated M4)~~ — answered 2026-07-28
  ([journey 0008](../04-JOURNEY/0019-soulidentity-sentinel-credential-flow.md), D19–D21
  in [`../02-DESIGN/auth-callout.md`](../02-DESIGN/soulstream-identity/auth-callout.md)): the
  client holds URL + external token only (`default_sentinel`), or a public
  bearer deny-all sentinel creds file besides; the issued JWT is for the
  server-assigned ephemeral key, minted by the vault's role keys;
  everything fails closed [measured]. D11's reversal condition is half
  resolved; the NGS half remains.
- ~~**The first-key story** (gated M3)~~ — answered 2026-07-28
  ([journey 0004](../04-JOURNEY/0015-soulidentity-first-key-story.md), D13): a `0600`
  local file on the service host, minted at first start; bootstrap is two
  operator acts + one automatic service act [measured].
- ~~**The claims-mapping shape** (gated M4)~~ — answered 2026-07-28
  ([journey 0009](../04-JOURNEY/0021-soulidentity-claims-mapping-shape.md), D22):
  validate → authorize → mint; the token record names an identity and
  carries no policy; Entra later is validator config + the D2
  claims-derived rules on the same interface; the issued-JWT TTL is the
  revocation propagation bound [measured]. D12's watch stays armed inside
  D22's reversal condition.
- **Service round-trip latency under real load** (informs M2/M3): signing
  and mint requests ride NATS request/reply, and callout sits on the connect
  path for represented users; the MCP node's real usage will measure both.
  The reversal condition in journey 0001 names what happens if the
  per-operation assumption fails.

### One-way doors

| Door | Constraint |
|---|---|
| **Custody boundary.** | Once consumers rely on "seeds never leave", any API that returns key material — however convenient — is a constitution-I amendment, not a feature. |
| **Wire contract.** | The agent's JSON surface is mirrored in `client/`; the payload shapes survive the transport swap to NATS subjects (M3) — changes after M2 must stay compatible or version the subject space. |
| **Vault record shape.** | Sealed `stored{}` records decode additively; a *required* binding on an existing kind (as D25 did to persona keys) means unbound records fail closed until re-imported — that re-import is the migration story, stated per change. The registry file door closed with the registry (D25, journey 0013). |

---

## soulstream — the product, the house


The live plan. No dates — **gates, not calendars**. Every milestone names the
research gate it depends on; nothing is built ahead of its gate (constitution
IV).

### Where we are

**Phase 1 is complete** (episodes
[0003](../04-JOURNEY/0042-soulnode-first-boot-is-real.md) /
[0004](../04-JOURNEY/0044-soulnode-the-realm-remembers.md) /
[0005](../04-JOURNEY/0045-soulnode-an-agent-runs.md), all 2026-08-02): `soulstream
init` founds a realm in ~0.15 s and `soulstream up` runs it — embedded
operator-mode server, identity plane, memory plane on ordinary loopback
connections; `soulstream workload start` runs a declared agent with a
minted credential under full enforcement. Every §9 exit criterion of
design [`0001-soulnode-composition.md`](../02-DESIGN/soulstream/0001-soulnode-composition.md)
is measured green in `make test`. The standing pseudo-version exception
**closed 2026-08-02**: soulstream-workloads v0.1.0, soulstream-archivist v0.2.0,
soulstream/node v0.7.0, and soulstream-identity v0.1.0 are all tagged and
pinned by tag. **Next:** Phase 2 — the front door — gated on soulstream's
`018-remote-mcp-node` cycle (in flight upstream, carrying the fourth
embed ask).

### Phase 0 — Composition (research) — ✅ closed 2026-08-02

**Gate met.** `single-binary-composition` graduated to design (episode
0002). Decided: five planes in one process, every plane on an ordinary
loopback NATS connection (constitution III as ratified); the first-boot
ceremony is code, provisioning through public surfaces only; the
in-process pipe transport is a finding of record (fixed ~10 s mute
refusals — candidate upstream issue), not the product shape.

### Phase 1 — The bundle (design → build) — *unblocked*

Runs the spec-kit flow against design 0001 (§9 acceptance criteria). Exit
criteria made precise per feature in `specs/NNN-*/`:

- **M1.1 — The server and the identity plane.** ✅ **Done** ([episode
  0003](../04-JOURNEY/0042-soulnode-first-boot-is-real.md);
  `specs/001-init-and-up/`). Measured: `init` founds a realm in 0.15 s
  (17 artifacts, owner-only modes, token printed once, re-init a
  verified no-op); the found→admit→refuse→revoke→restart e2e rides
  `make test` in ~1 s; admission matches the research exactly
  (server-asserted persona, own-prefix confinement, audited refusals).
  One open exception tracked: soulstream-identity pinned at a pseudo-version of
  main until it tags.
- **M1.2 — The realm joins.** ✅ **Done** ([episode
  0004](../04-JOURNEY/0044-soulnode-the-realm-remembers.md);
  `specs/002-realm-joins/`). Measured: the owner's full admission path —
  post → kept (author `owner`) → memory answers with attribution and a
  citation — in ~5 s inside `make test`; restart exactly-once; the
  disabled-plane arm clean; the archivist's persona key vault-held.
  Second pseudo-version pin tracked (archivist, above its v0.1.0).
- **M1.3 — An agent runs.** ✅ **Done — Phase 1 complete** ([episode
  0005](../04-JOURNEY/0045-soulnode-an-agent-runs.md);
  `specs/003-an-agent-runs/`). Measured: upstream's `agent-echo`,
  declared unchanged, runs with a minted TTL-bounded credential under
  full enforcement — turn authored by its persona, lifecycle a completed
  work item owned by `runner`, everything kept, nothing
  credential-shaped lingering. The two-keys split landed in the ceremony
  (plain workload minting key beside the scoped admission key). One
  consumer-proven upstream fix on the first enforcing run (soulstream-workloads
  `3fee11f`: agents need `$JS.API.INFO`). Third pseudo-version pin
  tracked (soulstream-workloads, no tags upstream).

External dependency, resolved 2026-08-02: soulstream-workloads tagged `v0.1.0`
and soulstream pins it by tag, alongside soulstream-archivist `v0.2.0`,
soulstream/node `v0.7.0`, and soulstream-identity `v0.1.0`.

### Phase 2 — The front door — ✅ local mode done 2026-08-02

**Gate met the same day**: upstream 018 landed and tagged (soulstream
v0.7.0), its node module made consumable (soulstream journey 0010).

- **The door plane.** ✅ **Done** ([episode
  0006](../04-JOURNEY/0049-soulnode-the-door-opens.md);
  `specs/004-the-front-door/`). Measured: MCP client + founding token →
  session, tools, realm-admitted `whoami`; garbage refused; the door
  custodies nothing (state dir untouched); disabled arm identical to
  Phase 1. The fourth pseudo-version pin (`soulstream/node`) closed
  2026-08-02 with the `node/v0.7.0` tag.
- **Public mode.** ✅ **Done** ([episode
  0055](../04-JOURNEY/0055-soulnode-the-public-door.md);
  `specs/005-public-door/`). `planes.door` grew
  `public_url`/`auth_issuer`/`auth_audience` additively (a package
  deal, validated); the OIDC lane rides the identity plane. Measured:
  the full hosted-client walk — 401 challenge → resource metadata →
  DCR → code+PKCE → bearer MCP session with the token's subject as the
  realm identity — against the upstream AS contract's stand-in;
  founding token coexists; undeclared roles and garbage refuse. HTTPS
  is deployment fronting (`tailscale serve` before the loopback door);
  the bundled-fold default (`auth_issuer` at the in-process fold) is
  soulstream-idp M5's wiring. One consumer-caught fix: listen port 0 now
  means "any free port" to the embedded server too, not nats-server's
  default 4222.

### Phase 3 — The tailnet inside — gated

An embedded tailnet node (tsnet) behind a flag: a stable MagicDNS name and
HTTPS certs with zero host setup. Built only if Phase 2 measures the
host-`tailscale serve` path insufficient for the audience — the dependency
is heavy, and the gate exists to keep it honest.

### The fold plane — ✅ landed 2026-08-03

**The bundled sign-in** ([episode
0057](../04-JOURNEY/0057-soulnode-the-folded-realm.md);
`specs/006-the-fold-plane/`): `planes.fold` (opt-in) runs soulstream-idp by
tag through its public embed seam on the node's own JetStream; public
door mode defaults its AS at the bundled fold. Measured: the whole
human chain — DCR, passkey sign-in at the bundled fold, bearer MCP
session at the door, `whoami` naming the fold identity — with zero
external services; old state dirs load unchanged (absent block =
disabled); the fold serves before the identity plane (its OIDC
validator discovers at startup). Two consumer-proven upstream
additions tagged on the way (soulstream-idp v0.1.1 NATSCreds, v0.1.2
persona-shaped ids).

### The release pipeline — ✅ landed 2026-08-03, v0.1.0 released

**CI + the tag-triggered release** ([episode
0058](../04-JOURNEY/0058-soulnode-the-release-pipeline.md)): the
archivist pattern composed — gofmt/build/test/lint on every push
(first runner run green, folded-realm gate included), goreleaser
publishing linux/darwin × amd64/arm64 with checksums and the stamped
version. Proven live: `v0.1.0` is
[released](https://github.com/impire-io/soulstream/releases/tag/v0.1.0),
an artifact round-tripped and verified. No windows on purpose (the
ceremony's owner-only modes). Zero credentials anywhere: the operator
made soulstream-workloads public, so the whole consumed stack fetches openly —
the "private-module credential" blocker dissolved.

### Later horizons (named, not planned)

Each will get its own research gate when it approaches:

- **BYO NATS.** Design 0001 §4 carries the [O]: the ceremony subset
  against a user-supplied server. Ships behind its own pass, not with the
  bundle.
- **Day 2.** Upgrade in place, backup/restore of the state dir, moving a
  realm to a new machine as a copy.
- **Multi-node.** Deferred to soulstream-workloads's Fleet work; soulstream stays
  single-node until the upstream node supervisor exists and a second node
  is a measured need.

### Discipline

Exit criteria are written before the work and amended only openly with the
raw findings recorded. Landing a feature updates this file, writes a
journey episode, and propagates design changes — in the same merge
(constitution VI).

---

## soulstream-idp — the fold


*The design docs in [`../02-DESIGN/`](../02-DESIGN/soulstream-idp/README.md) will say what
soulstream-idp is; this document decides what gets built first and behind which
gate. Every milestone's design arrives by research graduation before its
build starts — a capability that isn't decided yet is a research topic, not
a task.*

### Where we are (2026-08-02)

**Every research gate M1 names is concluded** — the store
([journey 0002](../04-JOURNEY/0043-soulfold-kv-schema-and-key-lifecycle.md) →
[store-and-key-lifecycle](../02-DESIGN/soulstream-idp/store-and-key-lifecycle.md),
D1–D8), the sign-in surface
([journey 0003](../04-JOURNEY/0046-soulfold-session-and-ui-shape.md) →
[session-and-ui](../02-DESIGN/soulstream-idp/session-and-ui.md), D9–D15), and
the envelope
([journey 0004](../04-JOURNEY/0051-soulfold-kv-encryption-at-rest.md) →
store-and-key-lifecycle D16–D19, amending D6): records are sealed
app-layer with the deployment's xkey, seed custodied outside the store,
username index digested, filestore encryption defense-in-depth only —
four bars, four passes, cost +1.19 ms on a real sign-in [measured].
**M1 shipped the same day** ([journey
0005](../04-JOURNEY/0052-soulfold-m1-the-op-skeleton.md),
`specs/001-op-skeleton/`): the OP skeleton is real — a stock go-oidc RP
signs in end to end on the sealed store, restarts (full and mid-flow)
are invisible, forged POSTs change nothing, and a full key rotation
runs under a never-restarted verifier with zero failures [measured].
Genesis: [journey 0001](../04-JOURNEY/0041-soulfold-genesis-the-fold.md).
**Next: M2 (passkeys), then M4 before M3** — the operator's
public-door-path priority (2026-08-02), openly reordering the milestone
list below: the door needs the admission proof (M4) sooner than the
admin lifecycle (M3).

### Milestones

1. ✅ **M1 — the OP skeleton** (shipped 2026-08-02, [journey
   0005](../04-JOURNEY/0052-soulfold-m1-the-op-skeleton.md),
   `specs/001-op-skeleton/`). Discovery, JWKS, and the
   authorization-code flow with PKCE served from the certified OP
   library (`zitadel/oidc`), storage on JetStream KV under the D16
   envelope, a seeded user and client standing in for the ceremonies.
   **Gate met [measured]**: a stock `go-oidc` RP completes sign-in
   against the running fold with an embedded nats-server as the store;
   the issued tokens (ID + JWT access, D15) verify against published
   JWKS; the fold survives restart — including mid-flow — with its
   state in KV; forged POSTs rejected with zero state change; a full
   key rotation under a never-restarted verifier at zero failures. Research before build: the KV schema and the signing-key
   lifecycle — **done** ([design](../02-DESIGN/soulstream-idp/store-and-key-lifecycle.md),
   [journey 0002](../04-JOURNEY/0043-soulfold-kv-schema-and-key-lifecycle.md));
   the session and UI shape — **done**
   ([design](../02-DESIGN/soulstream-idp/session-and-ui.md),
   [journey 0003](../04-JOURNEY/0046-soulfold-session-and-ui-shape.md));
   KV entry protection at rest — **done**
   ([design](../02-DESIGN/soulstream-idp/store-and-key-lifecycle.md) D16–D19,
   [journey 0004](../04-JOURNEY/0051-soulfold-kv-encryption-at-rest.md)).
2. ✅ **M2 — passkeys** (shipped 2026-08-02, [journey
   0006](../04-JOURNEY/0053-soulfold-m2-passkeys.md),
   `specs/002-passkeys/`). WebAuthn registration and login ceremonies
   (`go-webauthn`) replaced the seeded stub; the passkey-only rule
   (constitution I) is enforced behavior — the form POST is gone.
   **Gate met [measured]**: register-then-login proven at library
   level in `make test` (virtual authenticator doing real ES256
   ceremonies), the e2e sign-in authenticating only by ceremony, the
   D14 origin matrix refused server-side (4/4 foreign shapes), no
   credential secret in the store positive-control-verified; the
   real-authenticator browser runbook documented in the feature's
   `quickstart.md` — **running it is a pending human act**. Interim
   honesty: first-touch enrollment stands in for M3's researched
   bootstrap story.
3. ✅ **M3 — the lifecycle** (shipped 2026-08-03 — after M4/M5 per the
   operator's public-door priority — [journey
   0010](../04-JOURNEY/0060-soulfold-m3-the-lifecycle.md),
   `specs/005-the-lifecycle/`, **v0.2.0**). Users, groups (names =
   roles-claim values), client registration, invites, and the `/admin`
   JSON surface, on the graduated lifecycle design (D20–D24; the
   bootstrap-story research ran first, journey 0009). **Gate met
   [measured]**: from-nothing bootstrap to a signed-in admin in four
   counted acts; group membership changes surface in the next issued
   token; invites exactly-once and digest-stored; first-touch
   enrollment deleted; every prior milestone gate green on the invite
   mechanism. soulstream composed it the same hour (the founding invite
   printed once beside the founding token).
4. ✅ **M4 — the fold in the fleet** (shipped 2026-08-02 — before M3,
   the operator's public-door priority — [journey
   0007](../04-JOURNEY/0054-soulfold-m4-the-fold-in-the-fleet.md),
   `specs/003-fold-in-the-fleet/`, the `e2e/` rig module). **Gate met
   [measured]**: a passkey user's fold-issued access token admits
   through soulstream-identity's callout — imported at its published tag
   v0.1.0 via the public embed seam, configured with issuer URL +
   audience only — the role value resolving against the declared
   binding, scope server-enforced, refusals audited; and the identical
   rig passes with an Entra-shaped stub issuer
   (indistinguishability). Decision of record: the fold's access
   tokens speak Entra's claim vocabulary (oid / preferred_username /
   roles) because the seam's verifier keys subjects by oid
   [mechanism-argument].
5. ✅ **M5 — the embed seam and the default wiring** (shipped
   2026-08-03, [journey
   0008](../04-JOURNEY/0056-soulfold-m5-the-embed-seam.md),
   `specs/004-embed-seam/`). The public `embed.Run(ctx, Options)`
   assembly (D29 pattern), plus the AS-contract half the bundled story
   needs: RFC 7591 DCR (opt-in, discovery grows
   `registration_endpoint`) and the fixed token audience (opt-in,
   joining every `aud`). **Gate met [measured]**: `e2e/embedgate`
   (module path outside the namespace — `internal/` imports cannot
   compile) embeds and runs the fold through discovery, DCR, a passkey
   sign-in, and audience/roles-bearing tokens; the M4 rig's fold half
   now rides the seam; `soulstream-idp serve` is the seam's first consumer.
   The soulstream-side default wiring (`planes.door.auth_issuer` at the
   bundled fold) is that repo's feature to land.

### Open research questions (before their milestones)

- ~~**The KV schema and key lifecycle** (gates M1)~~ — concluded
  2026-08-02, all bars passed: see
  [store-and-key-lifecycle](../02-DESIGN/soulstream-idp/store-and-key-lifecycle.md) and
  [journey 0002](../04-JOURNEY/0043-soulfold-kv-schema-and-key-lifecycle.md).
- ~~**The session and UI shape** (gates M1/M2)~~ — concluded
  2026-08-02, all bars passed: see
  [session-and-ui](../02-DESIGN/soulstream-idp/session-and-ui.md) and
  [journey 0003](../04-JOURNEY/0046-soulfold-session-and-ui-shape.md).
- ~~**KV entry protection at rest** (gates M1)~~ — concluded
  2026-08-02, all bars passed: app-layer xkey sealing ships (D16–D19,
  D6 amended — username index digested); filestore encryption is
  defense-in-depth, not a substitute. See
  [store-and-key-lifecycle](../02-DESIGN/soulstream-idp/store-and-key-lifecycle.md)
  and [journey 0004](../04-JOURNEY/0051-soulfold-kv-encryption-at-rest.md).
- ~~**The bootstrap story** (gates M3)~~ — concluded 2026-08-03, all
  four bars passed: invitation is the only enrollment right (D20–D24
  in [lifecycle](../02-DESIGN/soulstream-idp/lifecycle.md)); first-touch
  deleted; pocket-id's open `/setup` refused. See
  [journey 0009](../04-JOURNEY/0059-soulfold-bootstrap-story.md).

### One-way doors

| Door | Constraint |
|---|---|
| **The seam.** | Once soulstream-identity's distribution defaults to the fold, any Soulstream-only claim, endpoint, or side-channel is a constitution-II amendment, not a feature — it would collapse the ecosystem's two planes into one. |
| **Passkeys only.** | Once users enroll, a password lane cannot be added without redefining Principle I; there is no quiet path to "temporary passwords". |
| **Store shape.** | KV records must decode additively once M1 lands; a breaking record change is a stated migration, never a silent re-read. |

---

## soulstream-shell — the shell

*The human cockpit — observe and configure the whole soulsystem from a
browser, beside the MCP door. Design:
[`0001-soulhelm-the-helm.md`](../02-DESIGN/soulstream-shell/0001-soulhelm-the-helm.md).*

### Where we are (2026-08-13, evening)

**v0.1.0 is shipped and composed** ([episode
0067](../04-JOURNEY/0067-soulhelm-founding-and-first-light.md)):
founded from the morning's graduated research ([episode
0066](../04-JOURNEY/0066-ecosystem-soulsystem-cockpit.md)), the helm
serves the observe surface with fold sessions and the first act, the
whole human ceremony riding its consumer-position e2e in ~4 s, and
soulstream runs it as `planes.helm` ([episode
0068](../04-JOURNEY/0068-soulnode-the-helm-plane.md)).

**Direction change, same day:** the component is reframed as **the
shell** — a pure modular frame with zero module logic, agnostic from
Soulstream by contract; every human surface is a module plugging in
through one exported contract. The
[`shell-module-contract`](../02-DESIGN/soulstream-shell/0002-the-module-shape.md)
research (four pre-registered bars) now gates M2, which arrives
module-shaped. The component renames **soulstream-shell** in the naming
re-centering's sweep ([episode
0069](../04-JOURNEY/0069-ecosystem-one-name-soulstream.md)).

### Milestones

1. ✅ **M0 — founding** (2026-08-13, episode 0067):
   [impire-io/soulstream-shell](https://github.com/impire-io/soulstream-shell), the
   founding articles in the README, hq wiring complete (component tag,
   constellation, this section).
2. ✅ **M1 — the observe surface + sessions** (v0.1.0, same day). One
   open scope amendment recorded in episode 0067: sessions (drafted as
   M2's first half) moved into the founding release — the surface never
   ships unauthenticated. Acceptance gates 1, 2, and 4 of design 0001
   §8 are standing tests; the first class-(a) act covers gate 3's
   record-op arm.
3. **M2 — the configure surfaces.** Design 0001 §4: classes (b) and (c)
   in the UI (tokens/people/clients; plane toggles with restart
   semantics). Class-(b) standalone authority may trail the tenancy
   topic's grant answer ([O2]); as a soulstream plane the ops lane
   carries it today.
4. **Later, gated:** the participant client (its own research topic —
   [O4], with upstream ask #1's WebSocket listener); a dedicated scoped
   helm ceremony user in soulstream (hardening).

### Open research questions (before their milestones)

- **[O1/ask #1]** soulstream's embedded WebSocket listener — gates the
  participant-client horizon, not the helm.
- **[O2/ask #2]** the class-(b) grant lane for a standalone helm —
  tracked in [`platform-tenancy-guardrails`](../01-RESEARCH/platform-tenancy-guardrails/README.md).
- **[O3]** the persona-id → display-name mapping source — decided at
  M1 build, recorded in design 0001.

### One-way doors

| Door | Constraint |
|---|---|
| **Pure consumer.** | The day the helm needs an `internal/` import or a surface that exists only for it, it stops being a component and the ambition returns to the existing surfaces — that is the graduated reversal condition, kept live. |
| **Custodies nothing.** | Sessions in memory; any durable helm store of record is a constitution amendment, not a feature. |

