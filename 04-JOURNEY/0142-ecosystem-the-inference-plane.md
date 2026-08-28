# Episode 0142 — The inference plane: five bars in a day, and a new component is founded (2026-08-28)

The `inference-plane` research — opened at the operator's direction the
same morning the dispatcher graduated, because the inference story
deserved a plane rather than a config knob — measured **all five
pre-registered bars PASS the day it opened** and graduates to the
founding design of a new component: **soulstream-inference (the mind)**,
the ninth member of the constellation, decided by the operator
2026-08-28.

The bars, with their numbers [all measured, 3× `-race`, consumer
position]:

- **The fleet shape holds** — anycast 40 requests in 4.5ms split
  24/16; micro's `$SRV.INFO` as the resolve surface with model/tags as
  instance metadata; resolve→filter→pin unicast 10/10; an unserved
  capability answering no-responders. The load-bearing find:
  **inference capabilities are tools** — an agent-scope mint tagged
  `tool:infer-chat` reached the plane through the shipped template,
  zero scope widening; and the measured boundary: resolve-and-pin is
  infrastructure's act, agents get anycast.
- **One client loop** — the four reply shapes (streamed, one-shot,
  oversized one-shot auto-streaming, mid-stream error) through one
  collect function whose only branch is the grammar's own tell; the
  request's stream flag never reaches the loop. Invariants refuse as
  protocol errors; sequences strictly consecutive with the
  terminator's count as second check — the enemy is silent
  truncation. Mechanism gifts: the server's 503 control frame
  surfaces as a typed client error in ~1.4ms without entering the
  loop, and `max_payload` charges the serialized header block
  byte-exactly (client-side, survivable — which is why the streaming
  fallback can exist). This bar ran as a parallel agent's spike,
  independently re-run and read line-by-line before recording.
- **The harness thinks credential-free** — the real `claude -p`
  announced in its own words that the env API key takes precedence
  over its login, then completed full round trips through a minimal
  realm door and the plane in 2.21s with the plane's marker verbatim
  in its answer; keyless requests died at the door with zero plane
  deliveries; no provider material anywhere the harness touched.
- **Names, not routes** — re-pointing a virtual name moved traffic
  alpha→beta (its effort default riding along) with zero caller
  change; un-pinned falls back to anycast; the model never in a
  subject, header, or the caller's hands.
- **The record stays the only record** — context assembled fresh from
  the topic each round (the instance saw 1, then 3, then 5 turns);
  the serving instance killed and replaced mid-conversation with the
  next round unchanged; the census clean: the plane created no
  stream, no KV, no bucket. The one deliberate divergence from every
  prior art this shape has: Soulstream already has the conversation
  store, so the plane refuses to grow one — the substrate boundary
  applied to thinking.

None of the three pre-registered reversal readings fired. Like the
dispatcher before it, the plane is **mostly composition**: micro
services + queue groups + the capability-minting scope machinery + D36
custody + one small HTTP door. What is genuinely new is the reply
grammar (now pinned by refusing invariants) and the door.

What it opened: the founding design
([`0001-the-inference-plane.md`](../02-DESIGN/soulstream-inference/0001-the-inference-plane.md))
with its [O] ledger — the catalogue's home, the door-key mint lane and
metering, chunked input behind a real-demand gate, further
capabilities and adapters by demand, realtime media behind its own
future research gate. Design 0007 §3's held inference block resolves
against the catalogue's names once both specs land. The component
repo's founding and first build follow immediately at the operator's
"build this autonomously" direction; the dispatcher's spec pass runs
in parallel on 0007's non-inference halves.

Reversal condition: the design's principles are each falsifiable in
the build — statelessness failing under real context sizes (the wire
discipline's escape hatches proving insufficient), or the tool-space
subject choice colliding with a real deployment's tool vocabulary,
reopens the affected section with the finding recorded; the research's
three readings stand as the observables.

Trail: research pre-registration, journal, and verdict
(`01-RESEARCH/inference-plane/`, removed at graduation — git history
`63123ff`…`91e0dae` keeps the full trail); design
[`0001-the-inference-plane.md`](../02-DESIGN/soulstream-inference/0001-the-inference-plane.md);
the fired successor clause in
[`0007-agents-as-infrastructure.md`](../02-DESIGN/soulstream-workloads/0007-agents-as-infrastructure.md)
§3/§9; spikes in the session scratchpad (`aai-bar1` grown four
inference tests; `infplane-grammar`, nine tests, the parallel agent's
module).
