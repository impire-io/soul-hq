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
