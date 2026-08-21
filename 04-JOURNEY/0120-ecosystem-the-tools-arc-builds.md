# Episode 0120 — The tools arc builds: catalog, custody, door (2026-08-21)

Episode 0118's design became running code across four repos in one arc,
in the graduation's own build order — most enabling first, most
irreversible last:

**core v0.12.0 — the discovery face.** The `toolcatalog` package:
`soulstream-tools` KV created on first write, entries additive by
construction (unknown fields survive a round trip byte-for-byte, an
unfamiliar kind reads present-but-unsupported — the record's own rule),
write-time validation only, absence a normal state everywhere
[measured: the extension's three acceptance criteria as package tests].
**v0.12.1 the same hour** — the build's first correction: the remote's
own service URL lived *nowhere* (the catalog refused endpoints on
remote entries; the plane's resource holds only OAuth endpoints), so
`endpoint` became reachability-for-both-kinds. The anti-split-brain
rule was always about the ceremony's half — client ids, secrets, AS
endpoints — and those stay refused.

**identity v0.10.0 — the custody half runs at runtime (D40).**
`resources.add|remove|list` on the sealed surface; the broker's catalog
mutable under one lock, the declared list merged with sealed store-held
records (the whole declaration one record, secret beside public half —
D39's rule applied at rest, which is also where the draft's
"D36-bucket" wording died: the grants domain custodies its own);
grants **always-on** in embed, so a deployment declaring nothing still
comes alive through the op. **The acceptance bar measured on the built
op** [measured, consumer position]: a resource added under a
5ms-cadence probe — **0 failed accesses, max gap 5.7ms** (bar: 50ms),
**first ceremony 2.1ms after the op returned, zero restarts** — beating
the graduated baseline (9.2–16.9ms restart gaps) rather than merely
meeting the bar. Plus persistence across a plane restart, removal
leaving standing custody untouched, the declared-name refusal, the
transport killing a represented user's `resources.add` before the
service hears it, and the secret sealed at rest with a fired control.

**workloads v0.7.0 — the lane carries a door's outbound identity.**
`Lane.MCPExtraEnv`: the subject a personal agent acts for and the
delegation authorizing it ride into the door's environment *by
declaration* — wrap scrubs the host's `SOULSTREAM_*` on purpose, so
nothing a door needs arrives by inheritance.

**soulstream v0.13.0-rc.9 — the door forwards (D41).** The stdio `mcp`
verb reads the catalog and re-exposes each target's tools beside the
record's own — `<entry>_<tool>`, one door, many targets, whose tool is
whose legible in any transcript. Endpoints only, never a process;
authority fetched per call (own grant, or on-behalf under the declared
delegation) and gone when the call returns; a refused fetch answers the
agent in words and **never touches the target**; per-entry failures
degrade to stderr notes; a realm with no catalog leaves the door
byte-identical to before [measured: package tests over real MCP
transports — per-call fetch counts, distinct bearers per call, the
no-bearer workload lane, refusal-in-words, entry-by-entry degradation].
The account segment derives from the connection's own server-asserted
grants unless declared.

Refuted or corrected along the way: the design's D36 placement for
resource custody (the grants domain owns its own, one record); the
catalog's no-endpoints-on-remote rule (reachability is not ceremony);
and the door rig's assumption that `{{MCP_CONFIG}}` was environment
(it is argv template vocabulary — recorded in 0118's research, honored
in the build).

What remains open, named in the design: the shell's halves — the
`resources.*` admin surface and the per-person linking ceremony
(browser-bound by nature) — and workload-target authority when a tool
must tell callers apart (the guardrail at the door's own chokepoint,
not a token).

Reversal condition: D39's stands (recurring tool-isn't-serving drift
between the halves forces one transactional write surface); D41's
stands (a required remote whose sessions outlive any fetchable token
reopens the no-custody invariant by name). New from the build: if
startup discovery proves too static for real catalogs (observable:
doors restarted to see new tools), lazy per-person listing returns as
the named alternative it was.

Trail: designs
[`external-tools.md`](../02-DESIGN/soulstream-identity/external-tools.md)
(as-built §) and
[`tool-catalog.md`](../02-DESIGN/soulstream-core/extensions/tool-catalog.md)
(the endpoint correction); core `37fea05` (v0.12.0) + `47e6b80`
(v0.12.1), identity `047c5f5` (v0.10.0), workloads `c66d070` (v0.7.0),
soulstream `c347847` (v0.13.0-rc.9); episode
[0118](0118-ecosystem-agent-external-tools.md) (the graduation this
builds).
