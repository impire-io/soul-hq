# soulstream-identity — external tools: the catalog and the forwarding door (D39–D41)

*Graduated from research topic `agent-external-tools` ([episode
0118](../../04-JOURNEY/0118-ecosystem-agent-external-tools.md); Bars 2
and 3 measured PASS on the full composition in consumer position, Bar 1
measured as the baseline this design's build must beat, Bar 4 the
operator's ratified statement). The problem it answers: a person adds
an external tool — a remote MCP server, a third-party API — to their
soulstream, and their agent uses it such that no outbound credential
ever exists in anything the agent can read, the remote sees the
calling person, and the catalog does not become a per-user store.
Decisions continue the global numbering: D39–D41. The record-side
discovery convention lives in
[`soulstream-core/extensions/tool-catalog.md`](../soulstream-core/extensions/tool-catalog.md);
this document owns the custody and the door.*

## D39 — The catalog is two layers: one discovery face, two custodies

"Which tools does this soulstream have?" has one answer surface: a
**record-borne catalog entry for every tool, uniformly** — the ones
this deployment runs (workload declarations, whose personas are
already in the registry) and the ones nobody here runs (remote
resources). The entry shape and its store are the record's convention
(the core extension doc); this decision fixes the split of authority:

- The record entry is **display/discovery-grade** — the A10 pattern
  (names demoted to display-layer resolution) applied to tools. It is
  never consulted for authority.
- The **identity plane custodies what needs a safe**: the full
  resource configuration with its client secret, in the D36 sealed
  secret store, under the plane's own ops (D40). The plane **never
  reads the record** — the cycle guard's zero edges stay zero
  [measured: zero soulstream-core edges in the module graph].
- The **writer writes both halves**: whoever adds a tool (the shell's
  module, an operator script) performs the plane op and writes the
  record entry. Drift is bounded honestly: a record entry whose plane
  resource is absent fails at link time as "this tool isn't serving" —
  the cross-link failure mode.
- **D26 stands un-bent**: entries are realm-level declarations, no
  per-user rows anywhere. Linking remains each persona's own act on
  their own prefix; the catalog names doors, never whose key opens one.

**Reversal condition**: if the two halves drift in practice faster
than link-time failure surfaces it (observable: recurring
tool-isn't-serving reports against entries believed live), the write
becomes a single transactional surface as a new D-decision — not a
quiet merge of the two stores.

## D40 — `resources.*`: the catalog's plane half is an op family

The remote-resource half of the catalog becomes runtime state behind a
management op family on the sealed surface, beside the vault's own:

| Op | Body → reply | Notes |
|---|---|---|
| `resources.add` | full resource declaration → `{}` | validates as the broker's `Resource.Validate` does; the client secret goes into the D36 store, never into the reply or any log |
| `resources.remove` | `{name}` → `{}` | standing grants against the resource keep their custody; the next ceremony refuses by name |
| `resources.list` | `{}` → public halves only | never the secret; the display face's source of truth for the writer |

- Management-gated like `guardrail.load` (the permission template
  decides who), and **evaluated by the guardrail like any op** — "who
  may add a tool" is exactly a D37 question, and the chokepoint
  already stands there.
- The broker's resource map becomes mutable under a lock, merged from
  two sources at equal rank: the statically declared list
  (`GrantResources` — kept, for deployments that refuse runtime
  mutation) and the store-held entries.
- **Acceptance criterion (Bar 1's pass, priced by its baseline)**: a
  resource added through the op serves its first link ceremony with
  zero process or plane restarts, while a 5ms-cadence probe on a
  pre-existing resource's `grants.access` shows no failed access and
  no gap beyond 50ms. The measured restart baseline it must beat:
  9.2–16.9ms max gap, 1–2 failed accesses [measured].

## D41 — The forwarding door: adapter position, endpoints only, nothing held

The door is how an agent reaches external tools: an MCP server the
agent's harness already speaks to, grown a forwarding half. Its
invariants, each measured in the graduation rig:

1. **Adapter position, never core.** Forwarding needs the identity
   client, and soulstream-core imports nothing of this ecosystem — so
   the forwarding half lives with whoever composes a door (the product
   binary's stdio verb; soulstream-mcp for remote clients when its
   demand arrives). `mcpserver.NewServer` returns the SDK's own
   server, so forwarding tools register beside the record's tools with
   zero core changes [measured].
2. **One door, N targets, endpoints only.** The agent's configuration
   names exactly one door. The door speaks MCP-client to its targets
   and re-exposes their tools; it **never spawns a process** — a tool
   this deployment runs is a workload (`role: tool`,
   `lifecycle: service`), run by workloads and reached as an endpoint.
   Tool-name prefixing across targets is user-visible and is decided
   at build, not silently.
3. **No outbound token at rest, in config, or in anything the agent
   reads.** Authority is fetched per call: the agent's own admission →
   `grants.access` (own, or on-behalf with the subject's D33
   delegation under standing C4 consent) → the token lives for the
   call and dies with it. The standing gate is Bar 2's scan: every
   file a wrapped run leaves on disk, the harness's own environment
   and argv, grepped for every token the provider ever minted, zero
   hits with a fired positive control [measured, 5/5]. The agent's own
   NATS credential and the delegation may appear in its configuration
   — they are the agent's identity and its authority artifact, not the
   remote's secret; the design says this out loud.
4. **The remote sees the calling person** [measured, 6/6: zero
   cross-attributions between two humans' agents; a stolen delegation
   refused as an actor mismatch; revocation refusing one agent's next
   call — which never reaches the remote — while the other's serves].
5. **The door refuses in words**: a refused access returns as an
   error-marked tool result naming the refusal, and the refused call
   never touches the remote — kept deliberately from the rig.

Measured cost of the whole path (consent-consulted on-behalf access +
remote MCP session + call): 1.8–3.4ms in-process [measured]; a real
wrapped wake end to end, person → record → harness → door → remote →
record: ~2s [measured].

## Configuration surface

- The plane: the D40 ops beside the existing `GrantResources` static
  list; the D36 store already present; the represented-user template
  unchanged (linking and access ride the existing `grants.>` tail).
- The door: which targets to expose, resolved from the catalog's
  discovery face; the subject/delegation it acts under. Where the
  delegation lives in a wrapped agent's configuration is stated
  deployment surface, not an accident.
- The shell: an admin surface for `resources.*` and the per-person
  linking ceremony (the browser half the broker always needed); a
  module design of its own at build time.

## As built (2026-08-21, episode [0120](../../04-JOURNEY/0120-ecosystem-the-tools-arc-builds.md))

Landed as core v0.12.0/v0.12.1 (toolcatalog), identity v0.10.0
(`resources.*`), workloads v0.7.0 (the lane's declared door
environment), soulstream v0.13.0-rc.9 (the door). The deltas worth
recording:

- **The resource record rests whole in the grants custody domain**, not
  the D36 bucket this document drafted: one sealed record
  (`resource/<name>`, public half beside its secret, same first key)
  in the store the broker already owns — D39's own anti-split-brain
  rule applied at rest, and one custody domain fewer in the path. D36
  remains the general per-persona store.
- **The grants half went always-on in embed**: a deployment declaring
  nothing statically still serves `resources.add` and comes alive
  without a restart — the enable-on-non-empty rule died with the
  static catalog.
- **The remote's service endpoint lives on the catalog entry**, not
  here: building the door found it living nowhere (the resource holds
  only OAuth endpoints), and reachability is the catalog's to say for
  both kinds. Core v0.12.1 carries the correction.
- **Startup collisions refuse loudly**: a name both declared and
  stored refuses the broker's construction; a link completed after its
  resource was retired refuses by name rather than dialing a
  zero-valued provider.
- **D40's acceptance bar, measured on the built op** [measured,
  consumer position]: an add under a 5ms-cadence probe on a
  pre-existing resource — 0 failed accesses, max gap 5.7ms against the
  50ms bar, the added resource serving its first ceremony 2.1ms after
  the op returned, zero restarts. Plus persistence across a plane
  restart, removal leaving standing custody untouched, the
  declared-name refusal, the transport killing a represented user's
  `resources.add`, and the secret sealed at rest with a fired control.
- **D41's build decisions**: tool names always prefix
  (`<entry>_<tool>` — [O1] closed); discovery runs at door startup,
  authority fetched for the listing exactly as for a call, per-entry
  failures degrading to stderr notes ([O2] closed as
  startup-discovery; lazy per-person listing returns if real catalogs
  chafe); workload targets ride no bearer (in-deployment authority is
  the guardrail's, D34 lane 4); the account segment derives from the
  connection's own server-asserted grants unless declared; an entry
  named to shadow the record's own tools is refused with a note.

## Open [O]

- **[O3] Tool-use visibility on the record**: catalog entries give
  future vocabulary a name to bind to ("scribe used github, on Daan's
  behalf"); deliberately not designed here.
- **[O4] The shell's halves** — ✅ closed: the tools module (shell
  design [`0005`](../soulstream-shell/0005-the-tools-module.md), shell
  v0.11.0-rc.2, episode
  [0122](../../04-JOURNEY/0122-ecosystem-the-shell-arc-lands.md)).
- **[O5] Workload-target authority**: today the door sends no bearer
  in-deployment; if a workload tool needs to tell callers apart, that
  is the guardrail at the door's own chokepoint (D37's third), not a
  token.
