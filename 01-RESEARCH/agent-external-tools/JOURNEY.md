# JOURNEY — agent-external-tools (opened 2026-08-19)

The investigation, as it happens. Appended to, never rewritten.

## 2026-08-19 — opened

Opened from the operator's evaluation of the running system: "with the
new functionality, don't we need the ability to add external tools, like
MCP servers to the shell? Maybe that's a separate module?"

The survey that produced the bars, before any experiment:

- soulstream-identity v0.9.0 ships the whole broker; the client mirrors
  it (`client/grants.go`: `GrantLinkStart`, `GrantLinkComplete`,
  `GrantAccessToken`, `GrantAccessOnBehalf`, `GrantAccessExchange`,
  `Grants`, `GrantRevoke`).
- `soulstream` (the house) declares no `GrantResources`, and
  `GrantResources` is what enables the `grants.*` ops
  (`embed/embed.go`) — so the broker is off in the product today.
- `GrantResource` is a value struct (name, auth/token/revoke URLs,
  client id + secret, scopes, redirect, exchange config), and
  [`grants.md`](../../02-DESIGN/soulstream-identity/grants.md)'s
  configuration surface states the catalog is "declared configuration,
  no per-user rows (D26's spirit: linking is the user's own act on
  their own prefix)". That sentence is the topic's central tension and
  Bar 4 exists to resolve it in the open.
- The linking ceremony is authorization-code + PKCE. It needs a
  browser, and the shell is the only surface in the ecosystem that has
  one.
- The agent-side seam is workloads' `wrap` and its `mcp_args`; nothing
  there fetches a token today.

## 2026-08-19 — seam reconnaissance: Bar 2 kills the obvious answer

Before any rig, a code trace over the two seams the answer has to live
in. It cost an hour and eliminated the implementation everybody would
have written first, which is what the method is for.

**The obvious answer, stated so it can be refuted.** A tool call needs a
token; `wrap` already renders the harness's MCP client configuration
**per run** — `writeMCPConfig`, mode 0600, into the run dir, with the
comment "it may carry a credential" already on it
(`wrap/harness.go`) [measured, code trace]. So: fetch `grants.access`
at wake, write the token into that per-run config beside the soulstream
door, let the harness call the remote itself. Short-lived, per-wake,
never persisted between runs.

**Bar 2 refuses it, and Bar 2 is right.** The harness runs with
`cmd.Dir = spec.RunDir` and `MCP_CONFIG` pointing into that same
directory [measured, code trace] — the agent reads that file, by
construction. A token written there is a credential in something the
agent can read, which is the *exact* wording of the bar, and the bar was
pre-registered as pass/fail with no acceptable location. A per-run file
is a better place than a checked-in `.mcp.json`; it is not zero, and the
bar does not grade on a curve.

Worth naming precisely, because the two are easy to conflate: the
agent's **own** NATS credential already lives on disk where the agent
can read it, and that is fine — it is the agent's identity, the thing it
is supposed to hold. Bar 2 is about the **remote system's** token, which
is somebody else's. `sanitizedEnv` already strips every `SOULSTREAM_*`
from the inherited environment so nothing of the host's leaks in
(`wrap/harness.go`) — the discipline this bar extends is one the repo
already keeps.

**What that leaves.** If the token cannot reach the agent, the agent
cannot call the remote directly — so the call has to go through
something that holds no token and fetches one per invocation on the
caller's own admission. That is a **door that forwards tool calls**, not
an injector. Which lands precisely on the chokepoint D37 named and did
not build: "the waker's wake decision and **the door's tool-call
forwarding**"
([`tenancy.md`](../../02-DESIGN/soulstream-identity/tenancy.md)). The
seam this topic shares with
[`guardrail-human-end`](../guardrail-human-end/README.md) is therefore
not a coincidence to note at graduation — it is where the answer lives.

**And the catalog question got concrete.** The broker holds
`resources map[string]Resource`, built once in `New(...)` and never
written again — no mutex, because nothing mutates it [measured, code
trace: `internal/grants/grants.go`]. So a runtime catalog is not a
configuration knob away; it is either a mutable, locked map plus a new
op (the route that strains D26's "no per-user rows", which Bar 4 must
argue rather than assume), or the catalog lives somewhere that is not
the broker and resources stay declared configuration.

**Consequences for the bars, recorded before the next step rather than
after it:** Bar 2 is now expected to pass *by construction* under the
forwarding-door shape and to fail under any injection shape, which
makes it a design discriminator rather than a measurement — it will
still be run as written, against a real agent, because a bar that is
argued and not measured is not a bar. Bar 1 is unchanged. Bar 3 is
unchanged and becomes the interesting one: a forwarding door has to
carry the *calling person* to the remote, not the agent, and that is
`AccessOnBehalf` with a bounded delegation, which already exists.

**Open, for the operator:** the shape this points at is a door with an
outbound half, and soulstream-mcp is parked (episode 0071) with stdio
MCP the choice of record for this iteration. Whether the forwarding door
is the stdio door growing an outbound side, or the parked remote door
waking up, is a scope decision this topic should put to the operator
before it spends a rig on either.

## 2026-08-19 — the door question answers itself

Followed the trace one step further rather than asking, and the scope
question above dissolved. Three facts, all read off the code:

1. **soulstream-core imports nothing of this ecosystem** — its `go.mod`
   has no `impire-io` requirement at all — and the cycle guard says
   neither core repo imports the other
   ([`soulstream-mcp/cycleguard_test.go`](../../../soulstream-mcp/cycleguard_test.go),
   the guard episode [0027](../../04-JOURNEY/0027-soulstream-dx-hardening-and-the-cycle-guard.md)
   set and 0107's Bar 5 re-proved) [measured, code trace]. Fetching
   `grants.access` needs the identity client, so **the forwarding half
   cannot live in `soulstream-core/mcpserver`**. It has to sit in an
   adapter position, where the two halves already meet.
2. **The agent's stdio door is already in an adapter position.** It is
   not a separate component at all: `soulstream wrap` points the harness
   at its own executable with the verb `mcp`
   (`cmd/soulstream/wrap.go`: `lane.MCPCommandLoc = exe`,
   `lane.MCPArgs = []string{"mcp"}`), and `cmdMCP` connects a realm
   client and runs `mcpserver.NewServer(client)` over stdio [measured,
   code trace]. The product binary already composes core *and* identity
   elsewhere, so the outbound half costs it no new dependency and
   breaks no guard.
3. **The tool surface is already extensible, for free.**
   `mcpserver.NewServer(c, opts...)` returns the SDK's own
   `*mcp.Server` rather than a wrapper, so a composer can
   `mcp.AddTool(s, …)` its own tools onto the same server [measured,
   code trace]. Forwarding tools register beside the record's tools
   with **no core change**.

So the answer to "which door" is neither of the two the last entry
posed: the forwarding half belongs to **whoever composes the door**, and
the product binary is already that for stdio. `soulstream-mcp` is the
same position for remote clients and inherits the same half when its
own demand arrives — which is what its founding article already calls
itself ("this module imports BOTH core repos"). Nothing needs to be
un-parked, and the 0071 focus is not strained: stdio stays the choice of
record.

The scope question is withdrawn rather than escalated. What replaces it
is narrower and belongs to Bar 4: the catalog still has to live
somewhere, and the broker's resource map is still built once and never
written.

## 2026-08-21 — the operator splits the question in two

The operator, reading the door finding: "does that mean we have an
adapter for each mcp server but might not run the mcp server ourselves?
Which would allow for remote MCP servers to be used while still allowing
stdio or 'remote' MCP servers to be ran as well through workloads?"

The question is sharper than the topic's own framing, and it splits
**adding a tool** into two things this journal had been treating as one:

- **Reaching one nobody here runs** — a hosted remote MCP server behind
  OAuth. No process to supervise; the whole problem is identity and
  credentials. This is what the topic was opened for, and it is grants
  lane 2/3 behind the forwarding door.
- **Running one** — a stdio server, or a remote one this deployment
  hosts. That is a workload, and **workloads already has the vocabulary
  for it**: `declaration.RoleTool` is documented as "a capability other
  workloads call", and `LifecycleService` ("long-lived; runs until
  stopped") is the only lifecycle `Validate` accepts, so every workload
  the room runs today is already a long-lived service [measured, code
  trace: `declaration/declaration.go`]. Running an MCP server is an
  existing declaration shape pointed at a different artifact, not new
  machinery.

**One correction to the operator's framing, and one invariant it
suggests.** Not an adapter per server: the agent talks to one door (the
stdio `soulstream mcp` its wrap config already points at), and that door
speaks MCP-client to N targets and re-exposes their tools. Per-server
adapters would put N entries back in the agent's own configuration, and
the catalog with them, which is the shape Bar 2 pushed us out of. And
the invariant, proposed here for the design that graduates: **the door
speaks only to endpoints and never spawns a process.** A stdio target is
run by workloads and exposed; the door's view stays uniform, supervision
stays where supervision belongs, and the door stays small enough to be
worth trusting.

**Consequence for Bar 4, arrived at from a direction the topic did not
anticipate.** The catalog's tension was declared-configuration versus
runtime data (D26's "no per-user rows"). If half the catalog is workload
declarations — which are *already* runtime data, already ops on the
record — then arguing the other half must remain static configuration is
much harder. The operator's question is an argument **for** the runtime
catalog. Bar 4 is unchanged as a criterion; what changed is that its
losing options now have to answer this.

**Open, and named rather than assumed:**

- **Tool-name collisions** when one door re-exposes several servers.
  Prefixing is user-visible, so it is a design decision, not an
  implementation detail.
- **Discovery is chicken-and-egg**: listing a target's tools needs a
  credential too, so either the door lists lazily per person or the
  catalog declares the tool surface.
- **The two kinds have different identity stories, and conflating them
  is the mistake this design could most easily make.** For a remote
  nobody here runs, the point is that the remote sees the *calling
  person* (grants + a bounded delegation). For one this deployment runs,
  who may call it is a Soulstream authorization decision — the guardrail
  chokepoint, and D34 lane 4 territory. Same door, different question
  behind it.
- **Doc/code drift noticed in passing**: `RoleTool`'s comment says "Not
  accepted in M1.1" while `Validate` accepts it. Not verified end to
  end; recorded so the next reader does not take either at face value.

**Scope held, deliberately.** The topic's pre-registered question is
about the tool nobody here runs — clause (b), "the remote sees the
calling person", does not apply the same way to a tool this deployment
hosts. So running-our-own stays out of the question and in this journal
as the finding that makes the endpoint invariant right. No bar is
amended.

**Next, and it needs the operator:** the bars are now measurable in a
rig. Bar 1 and Bar 3 can run against a stand-in authorization server the
way episode 0104's Bar 2 did (Dex, or this ecosystem's own idp, which
speaks OIDC and now RFC 8693). Bar 2 needs a real agent through
`soulstream wrap` against a remote MCP server that requires OAuth —
standable locally. What no rig can supply is episode 0104's still-open
residue, inherited here: **one real third-party provider** (a GitHub or
Google OAuth app), which is an account act only the operator can
perform.

## 2026-08-21 — Bar 3 measures PASS; Bar 1's baseline is measured

The operator re-ordered the plan before any rig was spent, and was
right to: Bar 4 is a statement bar (prototyping two catalog homes
overbuilds for a paragraph), 0107's build-order precedent puts the most
irreversible piece last, and the concentrated unmeasured risk was
Bar 3 — a wrapped agent's door connects as the *agent's* persona while
the remote must see the *calling person*, a composition (C4 consent →
D33 delegation → `AccessOnBehalf` → remote) never run outside the
embedgate rig. So: rig first with the catalog hardcoded, Bar 4 last.

**The rig** (session scratchpad, `bar3rig/`, consumer position:
published tags core v0.11.1 + identity v0.9.0, module path outside the
namespace, zero replaces): the embedgate ceremony (operator-mode server,
auth callout, memory resolver), a per-user stand-in AS with strict
rotation whose authorization codes name the remote user, a **real MCP
server as the remote** (go-sdk streamable HTTP behind Bearer auth,
logging the authenticated remote user per `tools/call`), and **the
forwarding door exactly as the journal's shape says**:
`mcpserver.NewServer(agent's realm client)` + one `mcp.AddTool` — the
record's tools and the forwarding tool on one server, zero core
changes — holding the subject's delegation and *no token*, fetching
authority per call via `grants.access` on-behalf. Agents are MCP
clients over in-memory transports, the way a harness holds its stdio
door.

**Bar 3 — PASS, 6/6 runs** [measured]:

- Two humans (daan-ext, avery-ext), each with an agent (scribe-daan,
  scribe-avery). Each human links their own grant (the code binding
  custody to their remote user), issues C4 consent on the record, and
  the mint consults the projection before signing (the S8 split held
  end to end).
- Each agent took the same tool call through its door. The remote's
  attribution log: `[remote-daan, remote-avery]` — each call the
  calling person's own remote subject, **zero cross-attributions**.
- A **stolen delegation refused**: avery's agent presenting daan's
  delegation from its own server-proven connection was refused.
- **Revocation**: daan unlinking the resource refused his agent's next
  call — which never reached the remote — while avery's kept serving
  (post-revoke remote log gained exactly one `remote-avery`).
- Every on-behalf access **audited both personas** (subject= lines for
  both humans; `grants.access` lines naming the acting agent).
- **Custody**: no access token the AS ever minted appeared in anything
  the agent saw (tools list + every tool result), positive control
  fired.
- Door round trip (consent-consulted access + remote MCP session +
  call): **1.8–3.4ms** across 6 runs.

**Bar 1 — FAIL on today's mechanism, by construction, with the
baseline priced** [measured, 6 runs]: no hot-add exists (the broker's
resource map is built once), so adding a resource is a plane restart.
Measured through the embed seam under a 5ms-cadence probe on the
pre-existing resource: **max gap between successful accesses
9.2–16.9ms, 1–2 failed accesses of ~200, and the added resource served
its first link ceremony 5.5–9.2ms after the restart began**. The number
that matters for Bar 4: an in-process plane restart is **~10ms of
surface outage** — the house runs identity as an in-process plane, so
a product that can cycle the plane pays roughly this, not a process
restart. Caveat recorded: this is the embed-seam restart on a warm
process; a full process restart (spawn + reconnect + bucket open) was
not measured and would be larger.

Two things the rig surfaced for the design, not for the bars:

- **Where the delegation lives is a design decision Bar 2 will force.**
  The delegation is not an outbound credential (it authorizes only its
  named actor, server-proven, at the identity plane) — it is the
  agent's own authority artifact, like its NATS credential. But it has
  to reach the door somehow, and "in the door's config" needs saying
  out loud rather than happening by default.
- **The door refuses in words**: a refused on-behalf access comes back
  to the agent as an `IsError` tool result naming the refusal, and the
  refused call never touches the remote. That behaviour fell out of
  the composition and should be kept deliberately.

## 2026-08-21 — Bar 2 measures PASS through the real wrap machinery

**The rig grew the agent half**: two binaries built from the rig module —
a **door** (the stand-in for what the product's `soulstream mcp` verb
would grow: environment-only lane, `mcpserver.NewServer` + the
forwarding tool, stdio transport, authority fetched per call) and a
**scripted harness** (everything a real assistant is to the wrap
machinery — reads the generated MCP config, launches the door over
stdio, takes one tool call, emits the JSONL terminal event — with no
LLM, because the bar measures the custody surface, not the reasoning).
The wrap side is the real thing at its published tag: `wrap.Wrapper`
(workloads v0.6.0) — mention wake, catch-up, run directory, generated
`mcp.json`, sanitized environment, terminal extraction, outcome op.

**Bar 2 — PASS, 5/5 runs** [measured]: the human's mention woke the
wrapped agent; the harness launched the door from the run's own
`mcp.json`; the door minted nothing and stored nothing, fetching
authority per call (`grants.access` on-behalf under the standing C4
consent); the remote attributed the call to `remote-daan`; the outcome
op on the record carries the remote's own answer — the whole loop
person → record → wake → harness → door → identity plane → remote →
record, in ~2s wall-clock including the wake machinery. Then the scan:
**every file the run left on disk** (the run tree: `mcp.json`,
`events.jsonl`, `stderr.txt`, plus the harness's own dumps of its
environment and argv, captured from inside the process) grepped for
**every access and refresh token the AS ever minted — zero hits**, with
the planted control found and the three scan surfaces asserted present.
The plane's audit named the subject through the whole wrapped run.

Two mechanical findings about wrap worth keeping (cost: two wrong runs):
`{{MCP_CONFIG}}` and `{{PROMPT}}` are argv template variables, not
environment — a harness is *handed* its config path in its own
command line, which is exactly where the real presets put it; and
`{{BODY}}` fills the Prompt, which then fills argv as `{{PROMPT}}` —
the two-stage fill is deliberate and undocumented-by-example until now.

**Where the bars now stand:** Bar 3 PASS (6/6), Bar 2 PASS (5/5),
Bar 1 FAIL-by-construction with its baseline priced (~10ms in-process
plane cycle), Bar 4 owed its statement. The one residue no rig closes:
the real third-party provider confirmation (a GitHub/Google OAuth
app) — the operator act inherited from episode 0104. The rig
(scratchpad `bar3rig/`, per how-we-work) is the draft of the standing
gates the build inherits at graduation.

## 2026-08-21 — Bar 4, the draft statement (for the operator to ratify)

**The catalog lives on the identity plane as a small admin op family
(`resources.add|remove|list`), its records in the sealed custody the
plane already has — the D36 secret store holding the client secret,
the public half beside it — loaded into the broker under a lock.**
Owner: soulstream-identity. Cost: one op family, a mutable resource
map, and nothing new at the chokepoint — the guardrail (D37) already
sits on the op path, and "who may add a tool" is exactly the kind of
op it exists to evaluate.

The losing options, and why they lost:

- **Static configuration (the status quo)** loses on the actor, not
  the price. Bar 1 measured its price honestly — an in-process plane
  cycle is ~10ms of surface outage, near-invisible under a 5ms probe —
  so the case against it was never performance. It is that adding a
  tool stays an operator-and-restart act (the shell's class (c)),
  which can never serve the demand this topic was opened by: a person
  adding a tool from the shell. A deployment class that refuses
  runtime mutation keeps static declaration as a merged baseline —
  the design should allow both lists, one store.
- **The record as catalog** loses on custody: a client secret cannot
  rest on an append-only log every persona in the realm can read, and
  splitting one entry across record (public half) and store (secret
  half) buys incoherence for no property. The record's half of the
  operator's split stays exactly where it is: a tool this deployment
  *runs* is a workload declaration, already runtime, already the
  record's.
- **A per-user catalog** was never a candidate and D26 stands un-bent:
  entries remain realm-level declarations, no per-user rows anywhere —
  linking remains the person's own act on their own prefix, which is
  the isolation property the transport already enforces.

Under this statement Bar 1's pass becomes the graduated design's
acceptance criterion (add through the op under the probe, zero
restarts), not a research re-run.
