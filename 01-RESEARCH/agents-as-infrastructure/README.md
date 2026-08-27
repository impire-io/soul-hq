# Agents as infrastructure — can a declared agent be submitted to the realm and served by a standing dispatcher, on nobody's laptop?

**Component:** ecosystem
**State:** active
**Started:** 2026-08-27

## Abstract

The declaration story is whole as of v0.14.0-rc.2 — wake, instructions,
capabilities, budget, all in one JSON contract — but an agent still runs
only where a person runs `soulstream wrap`: it lives and dies with a
terminal. The demand signal has arrived from real users (the operator's
report, 2026-08-27): what people miss is **submit-and-forget** — hand
the realm a declaration and the realm serves the agent's wakes from then
on — and a **shell surface for declaring agents**. This is design
[`0004-wrap.md`](../../02-DESIGN/soulstream-workloads/0004-wrap.md) §9's
recorded reversal condition firing by its own words: *a realm operating
agents as infrastructure — nobody's laptop, centrally credentialed —
brings the serve arm back as a fleet-era feature over the same engine,
with design 0003's placement answering which node runs a wake.* A
server-side harness also forces the question wrap dissolved by wrapping
the person's own signed-in assistant: **inference providers and
models** — what serves the agent's thinking when no person is logged in
on the machine. A decisive answer unlocks the dispatcher build, the
shell's declare surface designed against its submit op, and an honest
provider/model story (or its named successor topic).

## The question

What is the shape of the standing dispatcher that serves declared
agents as infrastructure — submission, placement, admission, and the
harness/provider story — such that submit-and-forget holds without a
store of record beside the log (the record is the position), without
bypassing the 0006 budget at admission, and without the declaration
ever carrying a credential?

## Pre-registered bars

Written before any experiment runs. Protocols name real rigs; a bar
passes only on measured evidence.

- **Bar 1 — submit-and-forget.** A declaration submitted to the realm
  as ordinary ops — with the submitting process exited before any
  trigger fires — is served by a standing dispatcher: a mention of the
  declared persona produces exactly one attributed outcome op. Killed
  and restarted mid-backlog, the dispatcher answers nothing twice and
  loses no stream-backed wake (exactly-once across restart, the 0130
  standard). Protocol: a spike over the wrap engine's admission path
  with a durable consumer per submitted agent, against a real NATS
  rig; kill/restart between trigger and outcome; replay-reconstruct
  the story from the log alone.
- **Bar 2 — placement decides which node.** With two dispatcher nodes
  on one realm, every submitted agent is served by exactly one node,
  decided by design 0003's claim path (submission as an ordinary work
  item — no auction, no coordinator, no new vocabulary unless measured
  necessary and named openly); killing the serving node reclaims its
  agents to the survivor as ordinary abandon→claim with no double
  serve — a wake posted during failover produces exactly one outcome.
  Protocol: two dispatcher processes, contested submissions, one node
  killed; ownership replay-reconstructible; zero probe traffic on the
  stream (the fleet's M3.1 standard).
- **Bar 3 — the budget gates the dispatcher's admission.** Design
  [`0006-loop-safety.md`](../../02-DESIGN/soulstream-workloads/0006-loop-safety.md)'s
  composed wake budget enforced at the dispatcher's admission seam —
  §6's by-construction claim measured, not assumed: the 0128 rig's
  uncooperative cases (ping-pong, id-evading self-post cascade, the
  colony) halt at their pre-computed bounds (D, 2K, N·K) through the
  dispatcher path with op-less loud refusals, while the legitimate
  owner→A→B→A delegation completes with zero refusals. Protocol:
  re-run the 0128 rig cases with the dispatcher in wrap's place.
- **Bar 4 — the harness thinks on custodied credentials.** A
  dispatcher-served agent invokes a harness whose model/provider
  access comes from identity-plane custody, never from a person's
  interactive login and never from the declaration (0005 §5's names,
  not grants): the declaration names its model/provider requirement;
  the dispatcher resolves the provider credential at wake time from
  the identity plane (secret store or grants broker — the research
  decides which lane and why); a custody probe shows the credential
  unreachable from the agent's own scope and absent from everything
  the agent reads (the 0118 Bar 2 scan standard); one wake completes
  end to end against a real provider or its measured stand-in.
  Protocol: a spike wiring identity-plane custody into the
  dispatcher's harness invocation, plus the custody scan.
- **Bar 5 — the shell can drive the loop as a pure consumer.** The
  declare→submit→served→answer loop is drivable from consumer
  position through surfaces the shell already reaches — public ops
  and published client packages, a session-scoped admission plus
  existing admin/approval lanes only where the design says so by
  name — with no shell-only upstream surface and no `internal/`
  import (the shell's one-way door). The refused arm measured too: a
  submission from a persona without the required standing refuses
  loudly. Protocol: the Bar 1 flow re-run in consumer position under
  the shell's constraints.

## Reversal condition

Two observable readings, either of which reverses the direction this
topic assumes (the serve arm returns as infrastructure):

- **The harness cannot be centrally credentialed.** Every viable
  harness for dispatcher-served agents demands interactive,
  per-person authentication that no custodied credential satisfies
  (observable: Bar 4's spike finds no lane — API key, exchanged
  token, or grant-derived credential — that a real harness accepts
  non-interactively). Then personal wrap remains the product truth,
  agents ride people's machines, and this topic graduates abandoned
  with the finding recorded.
- **Submit-and-forget demands a store beside the record.** The
  dispatcher cannot hold exactly-once across restart from the log and
  its durable consumers alone (observable: Bar 1's restart run loses
  or duplicates outcomes unless a dispatcher-owned store of record is
  added). Then the record-is-the-position discipline and this shape
  are in genuine conflict, and the conflict goes back to design
  before any build.

If the inference-provider/model question outgrows Bar 4 — more than
one lane, per-provider ceremony, model routing — it becomes a named
successor topic rather than silently widening this one (one question
per topic).

## Verdict

<Empty until graduation.>
