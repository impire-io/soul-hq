# 02-DESIGN — the normative design

This is what Soulstream *is*, specified functionally: **what must exist** and
**how each part behaves**. An implementer should be able to build a working
system from these documents without needing undocumented decisions. The reasons
behind the choices are not here — they live in
[`../00-GENESIS/rationale.md`](../../00-GENESIS/rationale.md).

**The spec-kit rule:** every document here is written explicit enough to be the
argument to `/speckit-specify` — the capability, its seams, its configuration
surface, and its acceptance criteria, with no guessing left to the spec writer.
Graduating research enters through `/research-graduate`; behavioral changes made
during implementation propagate back here (see
[`../00-GENESIS/how-we-work.md`](../../99-ARCHIVE/genesis/soulstream/how-we-work.md)).

## core/ — normative; this *is* Soulstream

A realm running only this is a working soulstream.

| Doc | Covers |
|---|---|
| [`core/01-protocol.md`](core/01-protocol.md) | Realms, the stream, the subject taxonomy, the operation record |
| [`core/02-identity.md`](core/02-identity.md) | Credentials, personas, attribution, delegation, notifications |
| [`core/03-topics.md`](core/03-topics.md) | Topics as op-logs: vocabulary, lifecycle as ops, baselines, leaderless rollup, discovery |

## extensions/ — optional conventions

A realm running none of these is still a working soulstream.

| Doc | Covers |
|---|---|
| [`extensions/registry.md`](extensions/registry.md) | Rich persona profiles, operator attestation, key distribution |
| [`extensions/library-and-adapters.md`](extensions/library-and-adapters.md) | The reference library, MCP adapter, WebSocket door, bridges, presence |
| [`extensions/curation.md`](extensions/curation.md) | Curator personas (what the old "steward" became) |
| [`extensions/work.md`](extensions/work.md) | The work stages: versioned artefacts, work items, execution, sandboxes |
| [`extensions/sealed-topics.md`](extensions/sealed-topics.md) | E2E-encrypted topics |
| [`extensions/memory.md`](extensions/memory.md) | Persona memory and collective search |
| [`extensions/tenancy.md`](extensions/tenancy.md) | The record's tenancy half: grant vocabulary, registry additions (F1's ensure-act, E4), and the two decided canonical-form amendments (A10 key-in-record, E3 acting-credential) |
| [`extensions/tool-catalog.md`](extensions/tool-catalog.md) | The tool catalog's discovery face: one realm-readable answer to "which tools are here", uniform across run-your-own and remote — display-grade, never authority; custody lives with soulstream-identity (D39–D41) |
| [`extensions/followed-board.md`](extensions/followed-board.md) | The followed board: one ordered consumer over `SOULSTREAM.TOPICS.>` maintaining a live in-memory board projection — announcement, lifecycle, last activity — serving snapshot reads and a change signal; the wire spec untouched, memory O(topics), one lifecycle fold shared with `Materialise`. Drafted 2026-08-30 as upstream ask #4 from shell design 0012 (the shipped `Board`'s ≥3 round trips per topic measured against a 7.5ms single pass). |
| [`extensions/presence.md`](extensions/presence.md) | The who-is-around face: a presence lease per running thing (wrap must; services should; tools/adapters may) — renewal on a cadence, farewell as manners, silence as the truth-teller, readers judging staleness; words a person reads, never numbers a machine steers by. Decided 2026-08-24 ahead of its build (episode 0124) |

The build order for all of the above is in
[`../03-IMPLEMENTATION/ROADMAP.md`](../../03-IMPLEMENTATION/ROADMAP.md); the frozen
per-feature spec-kit artifacts are in `specs/NNN-*/`.
