# Episode 0082 — What wakes an agent (2026-08-15)

The question, opened and closed in one day against real components: when a
persona that is a program is mentioned, what wakes it — and how is exactly
one answer guaranteed back into the topic, regardless of what the model did
during the run? Episode 0079 had landed identity and acting; this topic
measured the missing third — waking — through a runner spike (a durable
consumer on the persona's `NOTIFY` subject that invokes headless harnesses
via invocation templates and owns the reply). All four pre-registered bars
PASS `[measured]`:

- **The wake path**: a mention posted while no agent process existed woke
  unmodified headless claude-code (10.3s); the typed terminal event's text
  landed as exactly one attributed `turn.post`. With `post_turn` stripped
  from the MCP surface the reply still landed (4.4s) — the conversation's
  integrity provably does not depend on the model choosing to call a tool.
  The reply obligation belongs to the runner; MCP tools are enrichment.
- **Exactly one outcome op under faults**: SIGKILL mid-run and
  hang-to-max-deliver each produced one attributed failure turn; a model
  that replied via MCP mid-run produced one turn with no duplicate; zero
  dangling wakes. Refined by measurement: every **admitted** wake ends in
  exactly one outcome op — a *refused* wake produces no op at all, because
  the agent cannot speak.
- **The address outlives the process**: three mentions accumulated in the
  notify stream and drained on wake; revocation refused the next wake in
  **2ms** server-side while the persona stayed mentionable and its history
  attributed; re-minting re-admitted the *same* pending mention (delivery
  2 answered — a revocation window is a delay, not a loss). Per-run
  `mint.ephemeral` credentials measured: TTL bound enforced (admitted at
  t=0, refused at t≈12s past a 5s TTL), and a 150s-TTL credential carried
  a complete wake as the harness's only credential.
- **The template generalizes**: a second harness speaking codex's captured
  `exec --json` grammar ran through the byte-identical runner binary on a
  template-only change. (Live codex blocked on expired machine auth — its
  typed error event captured; the pre-registered script-harness fallback
  ran, built to the real grammar.)

Refuted along the way: that soulstream-workloads had trigger machinery to
extend (it has none — the declaration has no trigger vocabulary and
`function`/`job` are rejected reserved words `[measured]`); that
correlation can anchor on stream order (with several mentions in one topic,
an earlier wake's reply masquerades as a later mention's answer — the first
drain silently swallowed two replies; the fix is a before/after snapshot
diff of the run `[measured]`); and that the stdio adapter lives in
soulstream-mcp (it is soulstream-core's `cmd/soulstream-mcp`; the machine's
installed copy was two minor versions stale).

The not-bar decisions, argued in the topic journey: **G1** the waker joins
soulstream-workloads as its trigger arm, with support-layer standing no
workload scope gets (`[judgment]` on measured scope facts); **G2** the
runner posts the agent's *reply* as the agent (custodian of output) but
authors *failure* in its own persona — forced by the 2ms refusal, where the
agent's voice does not exist (`[mechanism-argument]`); **G3** harness
narration is presence, never turns; **G4** subagents carry no personas
until independently addressable; **G5** external agent protocols stay
adapters — a different grammar cost one template, zero code; **G6** the
invocation-template schema as measured (command argv with placeholders,
dot-path terminal mapping, MCP env, timeout, retry budget); **G7** loop
safety is real and unbuilt — the very first wake's reply @-mentioned its
asker and fired a fresh notify; agent-wakes-agent budgeting leaves as a
named successor topic, pre-registered before any such deployment.

Also carried out of the trials: `mint.ephemeral` is operator-gated — the
runner mints for the agent, never the agent for itself `[measured]`; and
an agent configured purely by the 0079 credential block posts unsigned
(the pinned standing defect), which the rig closed by passing
`SOULSTREAM_KEY_FILE`.

Reversal condition: if a mainstream harness's headless mode yields no
machine-readable terminal event and cannot pass the template test by
configuration alone (observable: harness-specific code required in the
runner), per-harness adapters reopen as real integrations. If the trigger
arm ships outside the workload plane (observable: a standing bespoke
scheduler beside soulstream-workloads in the product), G1 reopens. If the
harness ecosystem consolidates on a third-party agent protocol (observable:
two named harnesses shipping native support for the same one), the
adapter-to-that-protocol door is re-argued.

Trail: design [`0004-the-waker.md`](../02-DESIGN/soulstream-workloads/0004-the-waker.md);
topic commits `cad444d` (pre-registration), `60310ca` (bars 1/2/4),
`8fe3d9c` (bar 3); rig and runner spike in the session scratchpad per
how-we-work; episode 0079 is the identity/acting substrate this measured
against.
