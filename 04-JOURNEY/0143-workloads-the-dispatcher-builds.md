# Episode 0143 — The dispatcher builds: submit-and-forget is real (2026-08-28)

Design 0007's standing serve arm landed the day after its graduation —
built in parallel with the inference plane's founding, by a build agent
against the design and reviewed line-by-line before the merge
(`soulstream-workloads` `438a4c3`, `specs/011-dispatcher`, the
`dispatcher` package + a 603-line integration suite). The room now does
what the 0140 focus asked: hand the realm a declaration and the realm
serves the agent from then on.

What was built, every part a shipped mechanism composed: the placement
topic watched **live** with materialise-poll as catch-up only (the
research spike's poll deliberately not copied — design 0007 §9's build
requirement paid); open placements raced through the ordinary claim
path; owned placements **resumed from the log** on start with no new
op; agent placements served through the specs/009 wake engine — the
0006 budget riding its admission untouched; probes answered and peers
swept on fleet's own reclaim discipline, so a dispatcher node and a
runner node share one realm without either knowing about the other.

The design's open decisions, taken and recorded in the spec:

- **§2 the serve seam → option (b)**: the dispatcher owns its claim
  path. `TryPlace` hardwires `Runner.Launch` — correct for backend
  workloads — while the probe/sweep halves need no Runner at all [the
  graduation research's measured shape, now the shipped one].
  *Correction to this episode's first form (same day): "fleet
  untouched" overstated — fleet gained exactly one additive change,
  the unexported `placementOf` exported as `fleet.DeclarationOf`, so
  the placement wire format keeps exactly one definition rather than
  a drifting copy. No behavior change; the reclaim discipline and the
  Runner path are untouched.*
- **`Servable` is the self-selection line**: an agent with a wake set
  is engine-served; everything else stays the Runner path's — pinned
  by its own standing test (the sixth, beyond the asked-for bars).
- **Drain and crash are different ends, as ceremony**: `Drain` stops
  taking work, cancels engines, and waits so an in-flight failure
  lands the agent's own self-report; a hard stop posts nothing and
  the successor serves that wake exactly once.
- **Credentials stay out by construction**: the `ConnectAgent` hook is
  the entire seam — design 0007 §5's lane, wired at the product's
  founding ceremony, not here.
- **One addition the build taught**: a node-local `RaceBackoff` —
  without it, a declaration a node cannot serve becomes a
  claim/abandon spin on the record. It delays a decision, never makes
  one.

Acceptance, verified independently after the agent's build (my own
runs, 3× `-race` across the suite in 48.4s) [measured]:
submit-and-forget with restart resume and no re-claim (3.65s); a
hard-killed dispatcher posting nothing and the restart serving exactly
once (2.93s); two nodes racing and failing over with the wake in the
window answered exactly once (2.68s); the declared budget halting the
uncooperative cycle through the dispatcher path (2.35s) and the
legitimate delegation clean under defaults (2.33s); the runner path
untouched (1.22s). `make check` green on the merge.

The build's review also paid a debt in shipped code: **`fleet.Node`'s
owned set was an unsynchronised map** — written by `TryPlace`/`Release`
on the caller's goroutine, read by the probe callback on the
connection's; today's tests never interleave the accesses, so `-race`
had never fired. Fixed the same hour (`a5c7d0a`), guarded exactly as
the dispatcher guards its own served set. Five more findings entered
design 0007's ledger (§5/§6/§9): the tool door's credential is
per-agent while the template is per-node — the engine seam needs a
per-persona answer before a real harness is wired under the
dispatcher; the drain ceremony's self-report is a `Retries:1` property;
an engine that stops on its own converges on reclaim through the
probe-silence path [the build's reading, adopted]; `artifact` is
required by the schema but meaningless for engine-served agents; and
no retirement path exists — nothing ever un-places a placement.

What remains before a person can submit-and-forget end to end: the
product wiring — the founding mints the engine-credential lane behind
`ConnectAgent` (design 0007 §5's [O], now including the tool door's
per-persona half), a `dispatcher` plane in the house, and the
declaration's `inference` block closing against the plane's catalogue
(episode 0142, design 0001 §5) — one spec in `soulstream`, next.

Reversal condition: none — records a completed build against a
graduated design; the deviations taken are recorded above and in the
spec, none contradicting a measured bar.

Trail: `soulstream-workloads` branch `011-dispatcher` merged `438a4c3`
(spec `d7e80fb`, feature `d84bf17`, tests `6d77946`); design
[`0007-agents-as-infrastructure.md`](../02-DESIGN/soulstream-workloads/0007-agents-as-infrastructure.md)
(§2 now resolved in code); graduation [episode
0141](0141-ecosystem-agents-as-infrastructure.md).
