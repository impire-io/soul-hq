# Episode 0071 — The focus: a product, not a platform (2026-08-14)

The operator's challenge series — three questions put to the
constellation the day after it renamed — ended in a direction decision:
**we have been tempted into building platform depth that is not a
focused product.** The record supports the charge `[measured, from the
record itself]`: three workload backends and a fleet design with zero
external users; an entire passkey IdP dragged in because hosted MCP
connectors authenticate by OAuth only; the dogfood chafe gate lapsed
2026-08-10 with no verdict episode while building continued; and the
shell was declared "pretty useless" by the operator the day it shipped
its bars.

The product thesis, recorded: **Soulstream is a place for humans and
AI, and each connects the way that is native to them.** Humans connect
through a web surface — the shell. Agents connect through MCP — and
**stdio MCP is the choice of record for this iteration** `[judgment]`:
technically sufficient, at the accepted cost that online platforms
(hosted claude.ai connectors) cannot connect. That parks the whole
remote-OAuth chain the record shows was the idp's reason to exist.

The decisions:

- **Participation enters the shell's scope** — openly overturning
  design 0001 §1's "a cockpit, not a client" and dissolving [O4]. The
  mission milestone is **a usable cockpit**: view topics, collaborate
  directly (post, reply, comment, open topics), and mention
  notifications, live. Mechanism honesty: participation rides the
  backend-held per-session admission that Bar 2 and the first act
  already measured — the browser-native client and the WebSocket
  upstream ask stay parked; they were only ever needed for a
  NATS-in-the-browser client, which this is not.
- **Slimming is freeze-and-focus, not demolition** `[judgment]`: the
  inventory is thin, green, and cheap to hold; the expensive thing was
  the building reflex. Frozen behind real-demand gates: soulstream-mcp
  (built, tagged v0.1.0, waiting for the online-platform need — the
  stdio-only choice's reversal is cheap by construction), the idp's
  day-2, workloads' fleet build and any further backends, and sealed
  topics — **e2e encryption is explicitly out of iteration one** (it
  was never built; the deferral is now a decision, not a lapsed gate).
- **The identity plane stays what the product already needs** — token
  admission, workload minting, custody, the working passkey sign-in —
  and stops growing.
- The `shell-module-contract` bars ride the real build: the second
  module proving Bar 1 becomes the **collaboration surface** (the
  value), not fold administration (which waits as a later module) — a
  pre-experiment bar amendment, recorded in the topic journey.

Reversal condition: stdio-only reopens the day an online-platform
connection is a real need — observable as a named user who cannot
connect — and soulstream-mcp is already built for that day.
Participation-in-shell reverses if the custody gate cannot hold under
posting — observable as the custody scan failing or a design forced to
hold credentials in the browser.

Trail: episodes [0069](0069-ecosystem-one-name-soulstream.md) /
[0070](0070-ecosystem-the-rename-sweep.md); design 0001 §1/§9 amended
in place
([`0001-soulhelm-the-helm.md`](../02-DESIGN/soulstream-shell/0001-soulhelm-the-helm.md));
the re-scoped
[`shell-module-contract`](../01-RESEARCH/shell-module-contract/README.md)
topic; the roadmap's focus note.
