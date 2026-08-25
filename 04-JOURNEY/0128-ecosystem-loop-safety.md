# Episode 0128 — Loop safety: the room budgets the cascade (2026-08-25)

The colony gate, opened and closed in one day. The question, pre-registered
with four bars before any run: when agents wake agents, can the room bound
every cascade — cycles and fan-out — at wake admission, using only what the
record already carries, with zero cooperation from the inhabitants, while
human-rooted delegation of legitimate depth completes untouched? A
scratchpad rig (embedded JetStream, real `realm.Client`s, real
`wrap.Wrapper` instances driven by script `Invoker` closures, plus a
topic-wake colony loop built to design 0005 §3's author filter) answered
all four bars **PASS [measured]**, `-race` throughout, none amended:

- **The chain is readable from the record** — two always-mentioning agents
  under three concurrent human roots produced 421 turns in 5.0s (84.2
  wakes/s unenforced), and the record-only walker — the `WakeOpID` UUIDv5
  binding checked against every candidate, never stream order (0082's
  lesson) — resolved **421/421** to a human root, zero ambiguity, the
  three cascades cleanly separated (exactly 3 turns per depth, 1–141).
- **A budget at admission halts the uncooperative cycle** — a depth budget
  D=4 stopped the ping-pong at **exactly 4 outcomes** (the bound
  pre-computed before the run) with one loud, op-less refusal; the
  authorship-window budget (K own turns per topic per window) stopped an
  id-evading cascade at **exactly 2K**. No halt used the script's
  cooperation — the scripts were written to never stop.
- **Legitimate delegation survives the same gate** — human-rooted
  owner→A→B→A under the composed budget: exactly 3 outcomes, **zero**
  refusals, depths {1,2,3}, all chained to root. The budget is not
  "agents never wake agents".
- **Fan-out is bounded, and the danger is a number now** — three
  topic-wake agents, unenforced: **434 ops in 0.34s (1,264.7 ops/s) from
  ONE human op**, doubling per generation (3→6→12→24→48→96→192). Window
  K=3: **exactly N·K=9 ops**, 12 loud refusals; a fresh probe op drew +3
  refusals and zero new ops — parked and loud, never wedged.

Refuted along the way, both by measurement: **budget-in-the-harness-slot**
(a refusing invoker became a 312-turn failure ping-pong, 156 self-reports —
the wrapper's own outcome contract turns refusal into testimony that taps
the asker; 0083's runaway reborn *through* the budget), and **depth alone
as enforcement** (self-posted outcomes carry no UUIDv5 binding: 404/404
unresolvable to any root, and a D=4 gate watched 393 turns pass in 3s
without firing once — the pre-named `correlated_self_post` hazard is
total, not partial).

What it decided: **the composed budget at wake admission** — the
authorship-window as the enforcement floor (authorship is mechanical, the
client's persona; body text and arbitrary op ids cannot forge it), the
provable-chain depth as the tight bound and the operator's diagnostic
[measured]; refusals **op-less and loud**, before the outcome obligation
attaches [measured]; budget exhaustion is a **delay, not a loss** — the
mention stays in the inbox window, the outcome id stays deterministic, a
slid window re-admits; the record is the position, unchanged
[mechanism-argument]. No core wire change: everything computes from the
materialised view and `WakeOpID`. Named improvement, not a prerequisite:
if the MCP arm ever posts outcomes under the wake's deterministic id,
depth becomes reliable for self-posts too. Design
[`0006-loop-safety.md`](../02-DESIGN/soulstream-workloads/0006-loop-safety.md)
carries the functional spec; the sequencing rule of 0005 §7 resolves —
topic-wake colonies unblock when the budget ships in the admission path,
not on further research.

Reversal condition: a production harness whose outcomes neither ride the
deterministic id nor land as the persona's own authored ops (observable:
an outcome invisible to both the depth walk and the authorship count)
evades the window floor and reopens the topic; an admission point that
cannot read the topic view it budgets on under a real backend (observable:
a wake admitted without a materialised view) moves the mechanism or
reopens the topic.

Trail: research topic `01-RESEARCH/loop-safety/` (pre-registered bars,
substrate reading, per-run raw numbers; folder removed at graduation, full
history in git); topic commits `baedabc` (pre-registration), `38fa319`
(substrate + hazard pre-naming), `838af57` (the runs); rig in the session
scratchpad per how-we-work (`scratch/loop-safety-rig`, `-race -count=1`
green 45.9s); design
[`0006-loop-safety.md`](../02-DESIGN/soulstream-workloads/0006-loop-safety.md);
exhibits [episode 0082](0082-ecosystem-agent-participation.md) G7,
[episode 0083](0083-workloads-the-waker-lands.md),
[episode 0126](0126-ecosystem-agent-declaration.md); context
[`0004-wrap.md`](../02-DESIGN/soulstream-workloads/0004-wrap.md) §10,
[`0005-agent-declaration.md`](../02-DESIGN/soulstream-workloads/0005-agent-declaration.md)
§3/§7.
