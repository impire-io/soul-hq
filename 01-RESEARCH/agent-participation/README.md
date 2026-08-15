# What wakes an agent — and how is an answer guaranteed back?

**Component:** ecosystem
**State:** active
**Started:** 2026-08-15

## Abstract

[Episode 0079](../../04-JOURNEY/0079-shell-agents-join-the-stream.md) landed
two of the three things agent participation needs: an agent has a persona
with a revocable credential and a countersigned operator, and anything
MCP-speaking can act through the stdio door. What is missing is the third:
nothing *wakes* an agent that has no resident loop — a mention reaches its
inbox and sits there. The direction under test: waking is a
**soulstream-workloads** concern (a notify message triggers a workload that
invokes the harness headlessly), the reply obligation belongs to the
**runner** — which consumes the harness's typed terminal event and posts the
outcome — while the MCP tools remain enrichment the model may use, never the
channel the conversation depends on. If the bars hold, every harness
(claude-code, codex, a langchain service) plugs in through an invocation
template — configuration, not per-framework integration — and an agent is
addressable like a person whether or not a process exists. If they fail, the
verdict tells us where the waker must actually live and which harnesses need
real adapters.

## The question

**When a persona that is a program is mentioned, what wakes it — and how is
exactly one answer guaranteed back into the topic, regardless of what the
model did during the run?**

The scope is the wake path: notify subject → workload trigger → headless
invocation → outcome op. Acting (the MCP door,
[library-and-adapters](../../02-DESIGN/soulstream-core/extensions/library-and-adapters.md))
and identity ([the agent design](../../02-DESIGN/soulstream-identity/agent.md),
episode 0079) are treated as landed inputs. Resident agents that own their
loop and subscribe themselves are the degenerate case — they need nothing
from this topic and appear only as a control.

## Pre-registered bars

Written before any experiment. Each bar is measured against a rig wiring
real component releases (core, identity, workloads, the stdio MCP adapter);
the harness under test is an **unmodified** claude-code in headless mode
unless the bar names otherwise.

- **Bar 1 — a mention wakes a program that isn't running.** A registered
  agent persona is @-mentioned in a topic. A workload-triggered runner —
  holding a durable consumer on the persona's `NOTIFY` subject and nothing
  else resident — invokes the harness headlessly with the anchoring context;
  the runner posts the harness's **typed terminal event** text as a turn.
  **Pass:** the reply lands in the same topic attributed to the agent's
  persona, end to end; and in a discriminating trial with the `post_turn`
  tool removed from the harness's MCP surface, the reply still lands — the
  answer provably does not depend on the model choosing to call a tool.
- **Bar 2 — exactly one outcome op, under faults.** Three injections, each
  repeated: (a) the harness process is killed mid-run; (b) the harness hangs
  past the run timeout and redelivery runs to max-deliver; (c) the harness
  posts its reply through `post_turn` mid-run and its terminal text is only
  a report about having replied. **Pass:** every wake terminates in exactly
  one outcome op in the topic — an attributed failure turn for (a) and (b),
  the mid-run turn alone with zero duplicates for (c) — counted as ops by
  that persona since the anchoring op; and the notify message is acked only
  after its outcome op exists, so no trial leaves a dangling wake.
- **Bar 3 — the address outlives the process, and revocation bites the
  wake.** With no process running, at least three mentions accumulate on the
  durable consumer and the next wake(s) drain the backlog — none lost, each
  yielding its outcome op. The runner obtains a **per-run ephemeral
  credential** for the persona (the `mint.ephemeral` lane, TTL bounded by
  the run timeout) — no long-lived agent credential at rest in workload
  configuration. **Pass:** after the operator revokes the agent in the
  shell, the next wake is refused within the token lane's recorded bound and
  produces no reply, while the persona remains mentionable and its history
  stays attributed; re-granting re-admits the wake, not merely offers to.
- **Bar 4 — the template generalizes.** A second, structurally different
  harness (codex headless, or a minimal script-harness emitting the same
  event shape) passes Bar 1's flow given **only** a different invocation
  template — command line plus terminal-event mapping. **Pass:** the diff of
  runner code between the two harnesses is empty; the template is
  configuration, or this bar fails.

## Decisions that are not bars

Load-bearing calls this topic must take that no experiment can settle —
resolved by design argument, tagged `[mechanism-argument]` or `[judgment]`,
never `[measured]`. Recorded so the outputs stay honest about which are
evidence and which are choices.

| Marker | Decision | Character |
|---|---|---|
| G1 | The waker's home: a soulstream-workloads trigger feature, or its own component | structural |
| G2 | Who authors the failure turn — the agent's persona (runner as credential custodian, the adapter pattern) or the runner's own voice mentioning the agent (the no-attribution-laundering rule pulls this way) | attribution |
| G3 | Harness narration (the "let me look into this" mid-run messages) relays as ephemeral presence, never as turns | convention |
| G4 | Subagents carry no personas; they surface through the parent's voice; the criterion for when one graduates to its own persona with its own `operated_by` | judgment |
| G5 | External agent protocols (Synadia's included) are adapters if ever needed, never the contract — Soulstream's own wire stays the seam | direction |
| G6 | The invocation-template schema: what a registered agent's "how to run it" record carries (command, terminal-event mapping, timeout, redelivery budget) and where it lives in the Agents module | structural |
| G7 | Loop safety once agents can wake agents — mention-storm budget and debounce, and whether it lives in the waker | structural, likely successor topic |

## Reversal condition

Written now. **If real harnesses cannot be bracketed** — observable as a
mainstream harness whose headless mode yields no machine-readable terminal
event and which cannot pass Bar 4 by configuration alone — then the
runner-owned reply contract fails for that class, and per-harness adapters
reopen as real integrations rather than templates.

**If the waker cannot live in workloads** — observable as the rig growing a
standing bespoke scheduler beside soulstream-workloads, re-implementing
trigger, ack-after-outcome, or retry that the workload plane should own —
then G1 reverses and the waker becomes its own component decision.

**If the harness ecosystem consolidates on an external agent protocol** —
observable as two of the named harnesses shipping native support for the
same third-party protocol — the cheaper door is an adapter to that protocol,
and the invocation-template direction is re-argued rather than assumed.

## Verdict

Graduated 2026-08-15 `--to design` → [`02-DESIGN/soulstream-workloads/0004-the-waker.md`](../../02-DESIGN/soulstream-workloads/0004-the-waker.md), per G1 (the waker is the workload plane's trigger arm).

- **Bar 1 — PASS** `[measured]`. A mention posted while no agent
  process existed woke unmodified headless claude-code through the
  runner spike (10.3s run); the typed terminal `result` landed as
  exactly one `turn.post` authored by the agent persona.
  Discriminating trial: with `post_turn` removed from the MCP surface
  the reply still landed (4.4s, correct content) — the answer does
  not depend on the model calling a tool.
- **Bar 2 — PASS** `[measured]`. SIGKILL mid-run and hang-to-
  max-deliver each ended in exactly **one** attributed failure turn
  (delivery 2/2 named in the body); the MCP-mid-run trial ended in
  exactly **one** turn (the model's own, runner acked without
  posting). Consumer state after all trials: 0 unprocessed, 0
  redelivered — no dangling wakes. Refinement forced by Bar 3: the
  invariant is "every **admitted** wake ends in exactly one outcome
  op"; refused wakes produce no op by design.
- **Bar 3 — PASS** `[measured]`, on the full product stack (callout
  always on). Three mentions accumulated and drained (after fixing a
  real spike bug: correlation must diff before/after run snapshots,
  not anchor ordering — the first attempt silently swallowed two
  replies). Revocation refused the next wake in **2ms**, server-
  enforced, with the persona still mentionable and history attributed;
  re-grant re-admitted the *same* naked mention (delivery 2 answered).
  `mint.ephemeral` per-run credentials: 5s-TTL cred admitted at t=0,
  refused at t≈12s; a 150s-TTL cred carried a full wake as the
  harness's only credential. Caveat: minting is operator-gated — the
  runner mints, the agent cannot.
- **Bar 4 — PASS** `[measured]`, with its pre-registered fallback. A
  second harness speaking codex-cli 0.14's captured `exec --json`
  grammar ran through the byte-identical runner binary on a
  template-only change (dot-path terminal mapping). Live codex was
  blocked by expired machine auth (typed error event captured — even
  its failure path is machine-readable); the live rerun is one
  command once the operator re-authenticates.
