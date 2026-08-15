# 0004 — The waker: notify-triggered invocation

*Graduated from research `agent-participation` (episode
[0082](../../04-JOURNEY/0082-ecosystem-agent-participation.md)); all four
pre-registered bars measured PASS on a rig wiring core v0.8.2, the identity
plane, and the full product stack. **Built as M3.2** (episode
[0083](../../04-JOURNEY/0083-workloads-the-waker-lands.md),
[`specs/005-the-waker/`](../../../soulstream-workloads/specs/005-the-waker/)) —
amendments from the build are marked "landed correction". Status tags:
**[V]** unless marked.*

The waker is the workload plane's **trigger arm**: the standing component
that converts a message on a persona's notify subject into one invocation
of a harness, and guarantees the topic exactly one outcome op per admitted
wake. It is the machinery that makes an agent addressable like a person
whether or not a process exists. Nothing here changes core's wire: the
waker is a consumer and a client, never new vocabulary.

## 1. Vocabulary

- **Wake** — one delivery of one notify message to the waker, ending in
  exactly one of: an outcome op in the topic (admitted wake), or no op and
  a redelivery (refused or retried wake).
- **Invocation template** — the per-harness configuration that makes the
  waker harness-agnostic: command argv plus terminal-event mapping (§5).
- **Outcome op** — the one `turn.post` an admitted wake leaves behind: the
  agent's reply, or the waker's failure testimony.
- **Admission probe** — a pre-wake connection as the agent itself; refusal
  refuses the wake before a harness run is spent.

## 2. Standing **[V]**

The waker's standing is **support-layer, not workload-scope**: it MUST hold
(a) a durable consumer on the `SOULSTREAM_NOTIFY` stream, which no
workload-minted agent scope can create (`$JS.API.INFO` is an agent's whole
JetStream surface), and (b) an ops-lane credential able to call the
identity plane's token and mint ops — `mint.ephemeral` is operator-gated:
**the waker mints for the agent; the agent cannot mint for itself**. One
waker process serves one realm and any number of registered agents; each
agent gets its own durable consumer (`waker-<persona>`) filtered to
`SOULSTREAM.PERSONA.NOTIFY.<persona>`.

## 3. The consumer **[V]**

AckExplicit; `AckWait` MUST exceed the run timeout with margin; server-side
`MaxDeliver` unlimited — the retry budget is waker policy (§4), so
exhaustion produces a failure turn, never a silent server-side drop. The
notify message is acked **only after** the wake's outcome is decided:
outcome op posted, correlation found, or refusal naked for redelivery.
Mentions accumulate while nothing runs (the notify stream retains per
persona); a starting waker drains the backlog in stream order.

## 4. The wake protocol **[V]**

Per delivery:

1. **Probe admission** (when the agent is token-lane registered): connect
   as the agent (sentinel + token). Refusal → log, nak with delay, stop —
   no harness, no op. Measured bound: revocation refuses the next wake in
   ~2ms, and a re-granted agent answers the *same* pending mention on
   redelivery. A refused wake MUST NOT produce an agent-authored op — the
   agent cannot speak; see §7 for the waker's own testimony.
2. **Materialize context**: read the topic, take a **before snapshot** of
   its contributions, and carry the anchoring op's author and body into the
   prompt.
3. **Invoke** the harness per template: fresh run dir, generated
   `.mcp.json` bound to the agent's credential, environment scrubbed of
   every `SOULSTREAM_*` inherited variable, stdin closed, process-group
   kill on timeout.
4. **Extract the terminal event** by the template mapping. The event
   stream decides, never the exit code and never prose.
5. **Discharge the reply obligation** (§6).

## 5. The invocation template **[V]**

A registered agent's record carries, beside handle/shown-as/operator
(episode 0079), its "how to run it":

```json
{
  "command": ["claude", "-p", "{{PROMPT}}", "--output-format", "stream-json",
              "--verbose", "--mcp-config", "{{MCP_CONFIG}}",
              "--strict-mcp-config", "--allowedTools", "mcp__soulstream__*"],
  "terminal": {
    "type_field": "type",          // dot-paths; codex needs "msg.type"
    "terminal_value": "result",
    "text_field": "result",
    "status_field": "subtype",
    "success_value": "success"
  },
  "mcp_env": { "SOULSTREAM_URL": "…", "SOULSTREAM_CREDS": "…",
               "SOULSTREAM_TOKEN": "…", "SOULSTREAM_REALM": "…",
               "SOULSTREAM_PERSONA": "…" },
  "run_timeout_seconds": 150,
  "max_deliver": 2
}
```

Placeholders `{{PROMPT}}`, `{{TOPIC}}`, `{{MCP_CONFIG}}`, `{{RUN_DIR}}`
(landed correction: `{{TOPIC}}` joined for harnesses that take the topic as
an argument). The template's home is the **waker configuration file**
(research D2: registrations are operator configuration until the fleet's
claim path gives the declaration a second consumer). Harness idioms live in
the template, never in waker code — measured: claude-code and a
codex-grammar harness ran through a byte-identical waker on template-only
changes. A harness whose headless mode has no machine-readable terminal
event MUST be refused a template (or wrapped and named degraded) — that is
the per-harness contract, and the reversal condition if a mainstream
harness cannot meet it.

## 6. The reply obligation **[V]**

The runner owns the reply; MCP tools are enrichment. Every outcome
publishes under the wake's one deterministic op id — **UUIDv5 of the notify
op id and the agent persona** (landed correction: one mention can tap
several registered agents; a wake is one delivery *to one agent* — hashing
the notify op alone made two agents' outcomes dedupe into a single turn
[measured, the build's multi-agent gate test]). The id doubles as
`Nats-Msg-Id` (core v0.8.3's `PostTurnIdempotent`), so same-wake reposts
dedupe inside the record's 2-minute window; beyond it, a redelivery
pre-check (materialise, find the id) closes the crash-after-post window.
After the run, take an **after snapshot** and correlate by **set difference
of the persona's turns between snapshots** — MUST NOT correlate by anchor
stream order (measured failure: with several mentions in one topic, an
earlier wake's reply masquerades as a later mention's answer and swallows
replies).

- Harness posted during the run → ack, post nothing.
- Terminal success with text → post the text as the agent, ack.
- Failure with retry budget left → nak with delay.
- Failure at budget → post the failure turn (§7), ack.

Every admitted wake therefore ends in exactly one outcome op — measured
under SIGKILL, timeout-to-budget, and mid-run MCP posting.

## 7. Authorship **[V]**

The waker posts the agent's *reply* as the agent — custodian of output,
the adapter pattern; the text is the agent's own, signed with the agent's
key where one is wired (`SOULSTREAM_KEY_FILE`; an agent configured only by
the 0079 block posts unsigned — the pinned standing defect). *Failure* is
the waker's testimony **about** the agent and is authored by the waker's
own persona (the waker is a persona — one noun, D27), **naming the agent
in the body and tapping only the asker** (landed correction: the graduated
design said "mentioning the agent and the asker"; tapping the agent
notifies it, and a notify to a registered agent is a wake — the build's
hermetic gate measured the loop, failure turn → notify → wake → failure
turn, forever). A companion guard landed with it: a mention **authored by
the agent itself** never wakes it. Forced by measurement: a revoked agent
cannot speak, so the agent's voice can never be the failure channel, and
ghostwriting failure would launder attribution exactly when it matters.
*(The spike posted failures as the agent; the built waker enforces the
corrected policy — authorship is mechanical, two live clients, no
switchable author.)*

## 8. Credentials **[V]**

Two lanes, both measured:

- **Token lane (default)**: the wake rides the agent's registration
  (sentinel + `sit_` token); every connection gets a fresh callout-issued
  JWT, so the wake-path revocation bound is *next wake* — tighter than the
  general bound (open connections end at JWT expiry) because per-wake
  connections do not outlive the run.
- **Ephemeral lane**: the waker calls `mint.ephemeral` per wake (role by
  name, caller-generated key, TTL ≥ run timeout, no vault entry) and hands
  the harness a credential that expires with the run. Nothing long-lived
  reaches harness configuration; the registration (revocable token record)
  stays the fact the admission probe checks.

## 9. Narration and presence **[D]**

Mid-run assistant events are the agent's typing indicator. The waker MAY
relay them as ephemeral presence (outside the captured prefix, per the
presence convention) and MUST NOT post them as turns. Deliberate mid-run
`post_turn` calls by the model are real ops and are honored by §6.

## 10. Open, named **[O]**

- **Loop safety**: the first measured wake's reply @-mentioned its asker
  and fired a fresh notify; two woken agents can ping-pong on this
  machinery exactly as built. The waker is the budget point (every wake
  passes through it); debounce/budget design is a successor research
  topic, pre-registered before any agent-wakes-agent deployment.
- **Trigger vocabulary in the declaration**: `LifecycleFunction` is the
  reserved word this design makes real; how a workload declaration names
  its notify trigger (and how fleet placement claims a wake) is specified
  at spec-kit time against design 0003's claim path.
- **Where the waker's own persona registers**: as an operated agent
  (episode 0079's ceremony) or as founding furniture — a product
  composition question.

## 11. Acceptance criteria

1. A mention of a registered agent, posted while no agent process exists,
   yields exactly one reply turn authored by the agent, produced through a
   headless harness the waker invoked — including with `post_turn` absent
   from the tool surface.
2. Kill and hang faults end in exactly one waker-authored failure turn at
   retry budget; a mid-run MCP reply is never duplicated; no dangling
   deliveries after any trial.
3. With the waker down, mentions accumulate and later drain completely;
   revocation refuses the next wake with no op posted while the persona
   stays mentionable; re-grant answers the same pending mention.
4. A second harness passes criterion 1 by template change alone — the
   waker binary byte-identical.
