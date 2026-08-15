# agent-participation — investigation journey

Topic opened 2026-08-15.

## 2026-08-15 — Harness terminal events, measured before any rig

The reply contract stands on harnesses emitting a machine-readable
terminal event; probed both installed harnesses directly, no realm
involved.

- **claude-code 2.1.220** `[measured]`: `claude -p --output-format
  stream-json` refuses without `--verbose` (exact error recorded);
  with it, a one-prompt run emits 13 typed JSONL events (`system`,
  `assistant`, `result`) and the stream's last line is
  `{"type":"result","subtype":"success","result":"pong"}` — the
  terminal event is typed, distinct, and carries the final text.
  Gotcha for the runner: without stdin redirected it waits 3s for
  piped input (`< /dev/null` required).
- **codex-cli 0.14.0** `[measured]`: `codex exec --json
  --skip-git-repo-check` emits JSONL with a `msg.type` grammar
  (config echo, `prompt`, `background_event`, …). The run could not
  complete: the machine's codex auth is expired — ten
  `Failed to refresh token: 401 Unauthorized` retries, then a **typed
  terminal error event** `{"msg":{"type":"error","message":…}}`. Even
  the failure path is machine-readable, which is the property Bar 4's
  template mapping needs. Live codex stays blocked until the operator
  runs `codex login`; Bar 4 proceeds on its pre-registered fallback —
  a minimal script-harness — built to codex's *real captured* event
  grammar (stricter than the registered "same event shape": a
  different shape genuinely exercises the template's terminal-event
  mapping). A live codex re-run is a one-command repeat if auth
  returns.

## 2026-08-15 — Recon findings that shaped the rig

- The stdio adapter lives in **soulstream-core** (`cmd/soulstream-mcp`),
  not the soulstream-mcp repo (that is the remote HTTP door); the
  machine's installed copy was v0.6.0 — stale against v0.8.2 — so the
  rig builds both CLI and adapter from source `[measured]`.
- **No trigger machinery exists in soulstream-workloads** — workloads
  start one way (`workload start <declaration>`), the declaration has
  no trigger vocabulary, and `function`/`job` lifecycles are reserved
  words that validation rejects. The waker is new machinery, and G1 is
  a real decision, not an extension question `[measured]`.
- Core's own `topic.FollowInbox` is an **ephemeral ordered consumer**
  — the waker must own its durable consumer on the separate
  `SOULSTREAM_NOTIFY` stream directly. A workload-minted agent scope
  (`$JS.API.INFO` only) could not even create it: the waker needs
  standing no workload gets today `[measured]`, which is itself G1
  evidence.
- `mint.ephemeral` is **operator-gated**: the caller needs an
  ops-lane permission the represented agent never has — the runner
  mints for the agent; the agent cannot mint for itself `[measured,
  identity e2e]`. The waker is a privileged support-layer component,
  the same class as the shell's in episode 0079.
- The 0079 credential block is a literal `.mcp.json`; the adapter is
  configured purely by env (`SOULSTREAM_URL/CREDS/TOKEN/REALM/
  PERSONA`), lane fields never come from a config file, and an agent
  configured only by the block posts **unsigned** (the pinned standing
  defect) — the rig closes that by also passing
  `SOULSTREAM_KEY_FILE` `[measured]`.

## 2026-08-15 — Bar 1: PASS (both trials)

Rig: isolated `nats-server -js` on :4333 (context `agentrig`, realm
`agentrig`, never touching the operator's live context), core CLI +
stdio adapter built from soulstream-core @ v0.8.2, personas `operator`
and `clerk` with local signing keys. Runner spike: 380-line Go
program, durable consumer `waker-clerk` on
`SOULSTREAM.PERSONA.NOTIFY.clerk` (AckExplicit, ack only after the
outcome op exists), per-wake run dir with generated `.mcp.json`,
harness invoked by template, terminal event extracted by dot-path
mapping, reply posted via the reference CLI as the agent persona.
`SOULSTREAM_*` env scrubbed from every child.

- **Trial 1 (the wake path)** `[measured]`: a mention posted while no
  agent process existed sat in the notify stream, the runner woke
  `claude -p` (haiku, strict MCP config, 24-tool adapter surface),
  the harness ran 10.3s, and the terminal `result` text landed as a
  `turn.post` authored `clerk` — exactly one reply op. The reply
  itself @-mentioned the asker, firing a fresh `mention.notify`: G7's
  loop-safety concern observed live on the first wake.
- **Trial 2 (discriminating: `post_turn` removed)** `[measured]`:
  same config plus `--disallowedTools
  mcp__soulstream__soulstream_post_turn`; harness ran 4.4s, answered
  correctly (17×23=391), reply landed as clerk — the answer provably
  does not depend on the model choosing to call a tool.

Harness-side gotchas recorded for the invocation-template schema
(G6): `-p` + `stream-json` requires `--verbose`; stdin must be
`/dev/null`; `--strict-mcp-config` keeps the operator's own plugin
MCP servers out of the wake; `--allowedTools "mcp__soulstream__*"`
wildcard works.

## 2026-08-15 — Bar 2: PASS (all three faults)

Same rig, runner policy: on harness failure, retry (nak, 2s) until
`delivery == max_deliver` (2 in these trials), then post the
attributed failure turn and ack; ack never precedes the outcome op.

- **(a) SIGKILL mid-run** `[measured]`: process group killed 3s into
  each delivery; delivery 1 → retry, delivery 2 → failure turn
  ("harness died (signal: killed) after 3.007s (delivery 2/2)").
  Clerk ops in topic: **1**.
- **(b) hang past run timeout** `[measured]`: 5s timeout against a
  harness needing ~10s; two timeouts, then the failure turn. Clerk
  ops in topic: **1**.
- **(c) model replies via MCP mid-run** `[measured]`: prompt
  instructed the model to use `post_turn` and end with "I have posted
  my reply"; the runner's correlation check (ops by the persona with
  `stream_seq` past the anchor) found the MCP-posted turn and acked
  **without posting** — outcome kind `correlated_mcp_post`. Clerk ops
  in topic: **1**; the answer, not the self-referential terminal
  text, is what the topic holds.
- Consumer state after all trials: 0 unprocessed, 0 redelivered
  pending — no dangling wakes `[measured]`.
- Spike nit: the failure-turn wording puts the asker's handle where
  the agent's belongs; accidentally right as a *mechanism* (the
  mention notifies the asker of the failure), wrong as a sentence.
  G2's "who authors the failure turn" question stands.

## 2026-08-15 — Bar 4: PASS (template-only second harness)

Second harness: a 15-line shell script emitting **codex-cli 0.14's
`exec --json` grammar** (nested `msg.type` events; grammar captured
from the real binary earlier — auth-expiry blocked live codex). Only
the invocation template changed: command line plus terminal mapping
`msg.type == task_complete` / text at `msg.last_agent_message`, no
MCP block. Runner binary byte-identical (same md5 as the Bar 1 runs)
— the code diff between harnesses is empty by construction
`[measured]`. Wake → reply landed as clerk, one op. The dot-path
terminal mapping is what makes a template configuration rather than
code. Caveat, stated plainly: this proves the runner is
harness-agnostic against a *faithful stand-in*; the live-codex rerun
is a one-command repeat once the operator re-authenticates codex.

## 2026-08-15 — Bar 3: PASS (full product stack), one spike bug found and fixed

Stack: `soulstream init && up` (v0.11.0-rc.1 tree) on isolated ports —
operator-mode NATS :4433, callout always on, identity plane embedded.
Agent `wakebot` registered exactly as episode 0079's shell does it:
`CreateToken(realmPub, handle, …)` over the ops lane via the public
identity client (a 60-line `agentctl` spike). The runner grew three
full-stack concerns: its own creds for the durable consumer (ops
lane), a **pre-wake admission probe** — connect as the agent with
sentinel + token before spending a harness run — and CLI posting
through the agent's own token-lane context, so a revoked agent cannot
post even by way of the runner.

- **Backlog** `[measured]`: 3 mentions accumulated in
  `SOULSTREAM_NOTIFY` while nothing ran, then drained — 3 distinct
  replies. **The first drain attempt found a spike bug**: with several
  mentions in one topic, correlation-by-stream-seq let wake 1's reply
  masquerade as the answer to mentions 2 and 3 (both acked, no
  replies — the unanswered ops stay visible in the trial topic as the
  bug's honest cost). Fix: correlate by **before/after snapshot diff
  of the run**, not by anchor ordering. Re-run passed. The
  invocation-context snapshot the runner already takes is exactly the
  right baseline — a lesson the design doc must carry.
- **Revocation bites the wake in 2ms** `[measured]`: `RevokeToken` →
  next mention's admission probe refused (`Authorization Violation`,
  wake at .4329s refused at .4346s), no harness run, **no reply
  possible** — and the persona stayed mentionable (the mention posted
  fine) with all history attributed. The wake-path revocation bound is
  *next wake*, tighter than 0079's general bound (open connections
  end at JWT expiry) because per-wake connections don't outlive the
  run `[mechanism-argument]`.
- **Re-grant re-admits, not merely offers** `[measured]`: after
  re-minting, the *same* naked mention redelivered (delivery 2) and
  received its reply. A revocation window is a delay, not a loss.
- **Per-run ephemeral lane** `[measured]`: `mint.ephemeral` (role
  `realm`, caller-generated key, no vault entry) — a 5s-TTL credential
  admitted at t=0 and refused outright at t≈12s; a 150s-TTL credential
  carried a complete wake end-to-end as the harness's only credential
  (no token in its env). Caveats stated: the minting caller must hold
  operator-grade creds (the runner mints; the agent cannot), and the
  spike pre-mints manually — folding the mint into the runner's
  per-wake flow is one more step, not new machinery.
- **Invariant refinement the trials forced**: a *refused* wake
  produces no op at all — the agent cannot speak, so nothing is
  posted and the mention waits. Bar 2's invariant is precisely "every
  **admitted** wake ends in exactly one outcome op"; who tells the
  topic the agent is gone is G2's question, now with measured stakes.
- Wry note for calibration: the final reply said "the bar3 realm" —
  the realm is `agentfull`. Attribution, custody, and the envelope
  were exactly right; the *content* was the model's hallucination.
  The wake path guarantees the reply arrives and is honestly authored,
  never that it is smart.

## 2026-08-15 — The decisions that are not bars, resolved

- **G1 — the waker's home** `[judgment, on measured scope facts]`: the
  waker joins **soulstream-workloads as its trigger arm**, not a new
  component. The fleet design's op-log-triggered launch (0003-fleet)
  is the same shape — a standing observer converting stream events
  into launches — and `LifecycleFunction` is the reserved word waiting
  for exactly this. The measured caveat carries into the design: the
  waker's standing is privileged (durable consumer on the notify
  stream, ops-lane minting, posting) — support-layer class, like the
  shell's, never a workload's own scope. The reversal condition's
  observable (a bespoke scheduler beside workloads) describes the
  *spike*, as expected for a rig; it fires for the product only if the
  trigger arm ships outside the workload plane.
- **G2 — failure authorship** `[mechanism-argument]`: the runner posts
  the agent's *reply* as the agent — custodian of output, the adapter
  pattern, the text is the agent's own. But *failure* is the runner's
  testimony **about** the agent, spoken in the runner's own persona
  (the waker is a persona too — D27 makes jobs first-class),
  mentioning the agent and the asker. Forced by measurement: a revoked
  agent cannot speak at all (2ms refusal, no op possible), so the
  agent's voice can never be the failure channel — and a runner that
  ghostwrites failure as the agent would launder attribution the
  moment the agent is gone.
- **G3 — narration is presence, never turns** `[judgment]`: the
  mid-run assistant events (observed in every run's event log) are the
  agent's typing indicator; relaying them as ephemeral presence is an
  affordance, posting them would pollute the op-log. Unmeasured —
  presence relay was not built in the spike.
- **G4 — subagents carry no personas** `[judgment]`: claude-code
  spawned whatever it spawned inside its runs; nothing surfaced or
  needed to. Internal machinery speaks through the parent's voice; a
  subagent graduates to its own persona (with its own `operated_by`)
  exactly when it becomes independently addressable — mentioned,
  woken, revoked on its own.
- **G5 — external agent protocols stay adapters** `[mechanism-
  argument]`: a structurally different harness grammar cost one
  template, zero code. The wire (notify subjects, `turn.post`) is the
  contract; any third-party agent protocol that matters later is one
  more template or bridge, never the seam.
- **G6 — the invocation template, measured shape** `[measured]`: a
  registered agent's "how to run it" record needs: command argv with
  `{{PROMPT}}`/`{{MCP_CONFIG}}`/`{{RUN_DIR}}` placeholders; the
  terminal-event mapping (dot-path type field, terminal value, text
  field, optional status/success); the MCP env block; run timeout;
  retry budget (max-deliver). Harness idioms live in the template,
  not the runner: claude needs `--verbose` with stream-json, stdin
  from `/dev/null`, `--strict-mcp-config`, tool allow/deny flags.
- **G7 — loop safety is real and unbuilt** `[measured trigger,
  successor topic]`: the very first wake's reply @-mentioned the
  asker and fired a fresh notify. Two woken agents mentioning each
  other would ping-pong on this machinery exactly as built. The waker
  is the natural budget point (every wake passes through it);
  debounce/budget design leaves as a successor topic, pre-registered
  before any agent-wakes-agent deployment.
