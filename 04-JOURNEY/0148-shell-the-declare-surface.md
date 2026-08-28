# Episode 0148 — The declare surface: an agent placed from a browser (2026-08-28)

The arc's last human chapter — designed in the morning (design
[`0009`](../02-DESIGN/soulstream-shell/0009-the-declare-surface.md)),
built by a parallel agent against it, reviewed and merged the same
afternoon (`soulstream-shell` `bf0258f`, ~2,280 insertions), and
composed into the house the same hour (`soulstream` `dd8d6b1`). A
person can now do from the browser what episode 0147 gave the CLI:
author a declaration, submit it on their own admission, and watch it
come alive.

What the agents sheet grew, each line a design rule kept:

- **The form's output IS the declaration** — the same JSON the CLI
  takes, shown folded "as JSON"; parsing and validation are
  upstream's (`declaration.Parse`/`Validate`, one definition of the
  wire format), refusals surfaced in their own words, the
  credential-shaped-model refusal included. No second schema exists
  anywhere.
- **Submit is the person's own act** — `fleet.Submit` on the session's
  admission, because a placement is an ordinary work item any persona
  may open and the surface acts as nobody (the Bar 5 loop, now with a
  screen on it). Close the tab; the agent still arrives.
- **Arrival is the record's evidence** (0008's principle): the acting
  screen renders `open → claimed by <node>` live; a deployment with no
  dispatcher shows the open placement with honest words instead of a
  spinner.
- **The lists are readings** — placements and the models picker
  (catalogue names straight from realm KV, no inference dependency)
  reconstructed per render, no store; **nothing offers to retire**
  what no vocabulary can retire (design 0007 §9's open, respected on
  screen); **no secret passes through the shell** — the paste-able
  `soulstream provider set` command stands in (D36 is caller-own,
  spec 014's wall 3 carried).

The composition: the shell's embed surface grew two **declared
facts** — `PlacementsTopic` (the topic's NAME when the dispatcher
plane is on; the shell resolves it against the board itself, because
reading must not write) and `CapabilityRole` (`ceremony.AgentRole`,
the one name the shell cannot derive from anything it reads) — and
the house's helm plane hands both, absent-means-absent, the same
declaration pattern `AdminBase` and `GuardrailOn` ride.

Gates [measured, my verification]: the shell's full `make check`
green; module and support suites 3× `-race`; the e2e declare gate —
the browser-grade walk from form to claimed placement on a real
realm — in 6.38s, with the absent arm (no dispatcher declared, the
lane hidden with its words) at 0.61s; the house's gate green on the
composition. One new dependency entered the shell, named openly:
workloads' `declaration` and `fleet` packages at the rc.3 tag.

The storyline the operator opened on Wednesday is now complete at
every altitude: the substrate serves (0143), the mind thinks
(0144–0146), the house wires (0147), and the screen declares (this
episode). What remains is by-demand and named: catalogue writing from
the shell, per-wake keys, retirement vocabulary, the live provider
arms, and the fresh-eyes install that now has one more thing to see.

Reversal condition: none — records a completed build against design
0009; the e2e gate and the purity gates are its standing tripwires.

Trail: `soulstream-shell` `bf0258f` (`aba17d3` the lane, `71c4fa8`
the e2e, `1742fe1` the status refactor); `soulstream` `dd8d6b1` (the
declared facts); design
[`0009-the-declare-surface.md`](../02-DESIGN/soulstream-shell/0009-the-declare-surface.md);
the arc [episodes 0140](0140-ecosystem-the-focus-agents-as-infrastructure.md)–[0147](0147-soulstream-the-thinking-house.md).
