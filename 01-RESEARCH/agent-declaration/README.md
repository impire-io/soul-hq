# Can an agent be declared — who it is, what it may do, what wakes it — as a record, and run by the existing runtime with no new machinery?

**Component:** ecosystem
**State:** active
**Started:** 2026-08-23

## Abstract

The operator's ask: declare a new agent from instructions — who I am, what I
can do, what I need to act on — and deploy it, all from the shell. Almost
every word of that sentence already ships: persona and attestation (the
registry; the shell's Agents module, episode 0079), capabilities (grants,
the tool catalog, D28 scoped minting), deployment (declarations, the fleet
claim path M3.1, three isolation backends), and mention-wakes (wrap, design
0004). What has never been designed is the seam between them: the workload
declaration has **no trigger vocabulary** (named unbuilt in the roadmap,
unblocked since the claim path landed), instructions have **no declared
home**, and wrap §10 leaves "a registration pointing at an artifact rather
than a host command" and "registrations-as-shell-objects" open. The sibling
imps repo has a shipped channel vocabulary (external subjects, topic
op-logs, schedule ticks) and a measured three-way boundary — the record, the
room, the inhabitant (imps journey 0007) — recorded only outside this hq. A
decisive pass unlocks instruction-declared agents end to end and closes the
wrap §10 questions; a fail stops a stack-shaped drift (an agent platform
beside the record) before it starts.

Constraint carried in, not up for re-decision: "imp" never names an identity
kind — personas stay the one noun (rationale.md's standing rejection, D27,
the registry's no-`kind` rule). If the word survives anywhere, it names a
program shape, never a persona taxonomy.

## The question

Does "declare an agent from instructions" decompose **entirely into existing
vocabulary** — persona for who, grants for what it may do, a stage-1 topic
artefact for the instructions, the workload declaration plus a new *wake*
vocabulary for what it acts on, and shell-side composition for the verb —
with no new machinery beside the log (S2), no core wire or stream change
(work.md's acid test), and no behaviour branching on human vs machine
(workloads constitution II)?

## Pre-registered bars

Rig for all bars: the workloads hermetic-gate style — embedded NATS in
operator mode, a real declared agent, spike scripts in the session
scratchpad, conclusions and raw numbers appended to `JOURNEY.md`.

- **Bar 1 — Wake is vocabulary, not machinery.** A declaration extended
  with a `wake` section covering the four named sources — a mention of the
  persona, an op on a followed topic, a schedule tick, an external subject —
  runs on the rig: one agent, byte-identical declaration, at least one wake
  from each source, every wake leaving exactly one outcome op under a
  deterministic id, and a restart mid-backlog answering nothing twice
  (wrap §3/§6's position-is-the-record property preserved). *Pass:* 4/4
  sources fire using only shipped NATS/JetStream primitives (consumers,
  subscriptions, message scheduling) and additive op vocabulary. *Fail:*
  any source needs a change to the core wire format or stream layout.
- **Bar 2 — Instructions are an artefact of the record.** The agent's
  instructions live as a stage-1 artefact in a topic; the registration
  references the artefact, never a host path; the runtime materializes it
  per wake and keeps no durable copy. *Pass:* wake N runs revision A; the
  artefact is revised through ordinary ops; wake N+1 runs revision B with no
  redeploy; both wakes attributable in the log; killing the node between
  wakes loses scratch only, never history. *Fail:* correctness requires the
  runtime or wrapper to hold an authoritative copy — observable as a wake
  answering from state that appears nowhere in the topic (a constitution-I
  breach).
- **Bar 3 — Capability is the identity plane, unchanged.** The declared
  agent's allowed surface resolves entirely to shipped machinery: D28
  `mint.ephemeral` with scoped templates for transport, grants and the tool
  catalog for outbound calls. *Pass:* the rig agent successfully calls one
  granted tool and is refused one ungranted tool at the transport/grant
  layer, with zero new authorization code and no new policy store. *Fail:*
  the spike needs any new policy surface.
- **Bar 4 — The shell verb is composition.** Enumerate the
  declare-from-shell flow — name and shown-as into persona plus attestation;
  instructions into artefact ops; capabilities into grants; wake into a
  registration — and drive the full flow rig-side through shipped surfaces
  (the shell UI itself is out of scope for this topic). *Pass:* every
  mutation the flow performs is an existing op or API call, visible in the
  log. *Fail:* any step has no shipped surface; the missing surface is
  named, and per soulstream constitution I it lands upstream first.

## Named successors, not this topic

- **Agent-wakes-agent loop safety** — wrap §10's stated prerequisite before
  any colony-shaped deployment; two measured runaway exhibits already exist
  (episodes 0082, 0083). This topic's bars are single-agent by design; the
  successor must be open before any deployment where declared agents wake
  each other.
- **Runtime join/leave** — imps design 0003's unfired reversal condition
  ("the first real colony scenario in which an imp's topic set must change
  without a restart") is watched here, not decided here.

## Reversal condition

Written now, phrased as observable readings:

- A bar passes **only** by adding a second coordination store — an agent
  registry outside the record, durable runtime-held state, or a new
  privileged endpoint — observable as a mutation in the rig that never
  appears as an op. Then instruction-declared agents remain wrap-side
  host-command registrations, the trigger vocabulary stays unbuilt, and
  this direction is recorded as refuted with the raw rig evidence.
- Any bar passes only by branching behaviour on whether the persona is
  human or machine (workloads constitution II violated in the spike's
  code). Same consequence: the direction is abandoned, not patched.

## Verdict

*Empty until graduation. Filled by `/research-graduate`: PASS/FAIL per bar
with the honest numbers, each load-bearing claim tagged [measured] /
[mechanism-argument] / [judgment].*
