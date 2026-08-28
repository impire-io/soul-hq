# Episode 0141 — Agents as infrastructure: five bars in a day, the dispatcher is a composition (2026-08-27 → 2026-08-28)

The research topic the 0140 focus opened — can a declared agent be
submitted to the realm and served by a standing dispatcher, on nobody's
laptop? — graduated to design the morning after it was pre-registered:
**all five bars PASS**, measured in consumer-position spikes pinning
the rc.2 tags (core v0.14.0-rc.1, workloads v0.8.0-rc.2, identity
v0.12.0), each run three consecutive times under `-race` [measured].

The bars, with their numbers:

- **Submit-and-forget** — the submitter's connection closed before the
  dispatcher ever saw the item; submit→claimed+served 159.7–162.1ms; a
  fresh dispatcher instance with zero local state resumed from the log
  alone (exactly one live claim across the restart), deduped the served
  wake and answered the missed one in 162.9–214.5ms. The hard-kill arm:
  zero outcomes while dead, invocation at-least-once (2), outcome
  exactly-once (1), served 208ms after restart [measured].
- **Placement decides which node** — contested split 2/2 every run;
  live owners never reclaimed under continuous sweeps; a hard-killed
  node's agents reclaimed `claim,abandon,claim` and a wake posted into
  the failover window answered by the survivor in 1.051–1.064s with
  exactly one outcome; zero probe ops on the stream [measured].
- **The budget gates the dispatcher's admission** — the declared budget
  rode submission→`DeclaredConfig`→admission: the uncooperative cycle
  halted at exactly MaxHops=4 with a loud op-less refusal; owner→A→B→A
  under defaults completed 3 outcomes, zero refusals — design 0006 §6's
  by-construction claim now measured through the submission path
  [measured].
- **The harness thinks on custodied credentials** — the discovery came
  from reading before running: **neither canonical scope carries a
  secrets tail** (the agent scope reaches no identity-plane subject at
  all), so custody is structural, and all three probes died at the
  server as permissions violations [measured]. The dispatcher resolved
  the declared provider's secret at wake time in 1.98ms from its own
  D36 tree; a real subprocess read it from its environment — the same
  non-interactive lane real harnesses take with env API keys
  [mechanism-argument]; `wrap.Template.Env` is the already-shipped
  injection seam; the record census found the value nowhere.
- **The shell drives the loop as a pure consumer** — the whole
  declare→submit→served→answered→read-back loop on a minted canonical
  persona-scope session admission in 561ms end to end; the
  pure-consumer line compiler-enforced (the spike module sits outside
  every repo namespace); the agent-scope credential's submission never
  reached the record [measured].

Nothing was refuted, and that is itself the finding: **the standing
dispatcher is a composition of shipped mechanisms, not new mechanism**
— submission is `fleet.Submit`, placement is the 0003 claim path,
serving is the wrap engine, the budget travels with the engine, and
custody is the identity plane's existing structure. Neither reversal
reading fired: env-injected keys are the non-interactive lane, and no
store beside the log was needed anywhere. Two sharpenings worth
keeping: a *graceful* stop posts the harness's failure self-report (the
engine's contract) while a *crash* leaves the record clean for the
successor — the dispatcher's stop ceremony must choose drain
deliberately; and restart-dedup and failover-dedup are one mechanism,
the deterministic outcome id.

What it opened, now in the graduated design
([`0007-agents-as-infrastructure.md`](../02-DESIGN/soulstream-workloads/0007-agents-as-infrastructure.md))
as its [O] ledger for the operator: the serve seam's shape (fleet's
`TryPlace` hardwires `Runner.Launch`), the declaration's first-class
`inference` block (the spike rode `args`), the founding's role-key
naming for the engine-credential mint, the provider-secret naming
convention and the grants-broker boundary (a person's own provider
account), and the shell's declare-surface module design at its build.

Reversal condition: the design's own — a build that cannot hold
exactly-once from the log and durable consumers alone, or a harness
population that abandons env-credential lanes, reopens the question;
the research's two pre-registered readings stand as the observables.

Trail: research pre-registration and journal
(`01-RESEARCH/agents-as-infrastructure/`, removed at graduation — git
history at `7f895c1`…`4b7b272` keeps the full trail including the
verdict); design
[`0007-agents-as-infrastructure.md`](../02-DESIGN/soulstream-workloads/0007-agents-as-infrastructure.md);
the focus [episode 0140](0140-ecosystem-the-focus-agents-as-infrastructure.md);
the fired reversal clause design 0004 §9; spikes in the session
scratchpad (consumer-position module `aai-bar1`, seven tests).
