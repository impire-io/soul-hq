# Episode 0085 — Wrap: run your agent where you are (2026-08-15)

Hours after the waker landed, the operator reshaped it with two sentences:
"waker" is not a good name, and the easiest attach path is personal — run
the agent where you want and wrap it. Both held up better than the built
thing. The byname failed the constitution's own plain-word test; and the
personal wrapper **dissolves the provider-login problem the same message
had raised** (episode 0083's open question): the harness runs on the
machine where the person already ran `claude login` — their config, their
models, their keys — so nothing needs delivering at all `[judgment ratified
by mechanics]`.

The reshape shipped across three repos the same day:

- **Core v0.8.4** (specs/019): the CLI grows **external subcommands** —
  the git convention: an unknown verb execs `soulstream-<verb>` from PATH
  with the resolved identity projected into its environment; built-ins
  always win; a typo still gets the usage error. `soulstream wrap` is the
  seam's named first occupant `[measured, hermetic stub test]`.
- **Workloads** (specs/006, `waker/` → `wrap/`): the central daemon is
  **cut** — durable consumers, operator standing, dialer lanes, minting,
  the second persona, all deleted, its reversal condition recorded in
  design 0004 §9 (agents-as-infrastructure or the fleet's claim path
  brings a serve arm back over the same engine). What replaced it needs
  none of that standing: **the record is the position**. Every outcome
  publishes under the deterministic wake id, so the wrapper — holding
  only the agent's own credential block — catches up by reading its
  bounded inbox and answering whatever has no outcome op, then follows a
  raw subscription live. Restarts answer nothing twice `[measured]`;
  faults become the agent's own **self-report** (the same principal that
  would have replied — the ghostwriting objection doesn't apply to
  self-report), tapping only the asker `[measured]`.
- **Shell**: the credential screen's set-up guidance grew its fourth
  fold — export the block's five values, `soulstream wrap --harness
  claude`, done.

Measured on the way through: the full hermetic suite reshaped and green
(backlog + live + restart-no-duplicates, three fault classes, the
operator-mode refusal, the second grammar by template); `make test-wrap`
woke a **real claude** through the wrapper in 19s — and its reply
@-mentioned itself ("@clerk here"), firing a self-notify the measured
self-loop guard absorbed unprompted. A live footnote for the folds: the
preset's tool door resolves `soulstream-mcp` from PATH, and a stale one
degrades the *enrichment* (the model told us its topic read failed) while
the reply still lands — the reply obligation working exactly as designed
`[measured]`.

What was given up, honestly: the daemon's at-least-once redelivery
machinery (JetStream ack/nak) is gone; the wrapper's guarantee leans on
the bounded inbox window (newest 100 per persona — protocol) plus
deterministic ids. Mentions that fall out of the window while a wrapper
is off are lost to the wrapper (the ops stay in topic history). Named,
not hidden — and the cut is a net deletion: no consumer rights, no
support-layer credential, no per-agent provider-key gap (the harness
inherits the host's login state; the template's new `env` block covers
the per-agent key when that isn't the lane).

Reversal condition: carried in design 0004 §9 — a realm operating agents
as infrastructure (observable: several wrappers under a process
supervisor and an operator asking for one process, or fleet placement
landing) returns the serve arm over the same engine.

Trail: design [`0004-wrap.md`](../02-DESIGN/soulstream-workloads/0004-wrap.md)
(rewritten; file renamed with references repaired); core `320ac33` + tag
`v0.8.4` ([`specs/019-external-subcommands/`](../../soulstream-core/specs/019-external-subcommands/));
workloads `38a33b0`→`7bbd441` (merge,
[`specs/006-wrap/`](../../soulstream-workloads/specs/006-wrap/)); shell
`4d0d6ed`; episodes [0083](0083-workloads-the-waker-lands.md) (what the
daemon proved) and [0084](0084-shell-where-this-goes.md) (the folds this
extends).
