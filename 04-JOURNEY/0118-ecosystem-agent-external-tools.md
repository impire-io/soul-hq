# Episode 0118 — External tools: the agent reaches out, the person is who the remote sees (2026-08-19 → 2026-08-21)

The question, opened by episode 0116's evaluation ("don't we need the
ability to add external tools, like MCP servers?"): how does a person
add an external tool to their soulstream and let their agent use it,
such that no outbound credential ever exists in anything the agent can
read, the remote sees the calling person, and the catalog never becomes
a per-user store. Four pre-registered bars; graduated to design with
three measured and the fourth ratified.

**The shape found before any rig, by code trace**: the obvious
implementation — write a per-wake token into the run's MCP config —
dies on Bar 2's own wording (the harness reads that file by
construction), which leaves a **door that forwards and holds nothing**;
the door question then dissolved ("which door?" is the wrong question —
the forwarding half belongs to whoever composes a door, since core
imports nothing of the ecosystem and `mcpserver.NewServer` returns the
SDK's own server, so forwarding tools register beside the record's with
zero core changes). The operator's own question split the topic
cleanly: reaching a tool nobody here runs is identity work; running one
is a workload (`role: tool` — the vocabulary already existed), and the
door speaks **endpoints only, never spawning**.

**The bars** (rig in the session scratchpad, consumer position, all
published tags, zero replaces; spreads over 5–6 runs):

- **Bar 3 — the remote sees the caller: PASS 6/6** [measured]. Two
  humans' agents, one real MCP remote behind OAuth: every call
  attributed to that person's own remote subject, zero
  cross-attributions; a stolen delegation refused as an actor
  mismatch; one person's revocation refused their agent's next call —
  which never reached the remote — while the other's kept serving;
  every on-behalf access audited both personas. The full C4-consent →
  D33-delegation → on-behalf composition ran outside its home rig for
  the first time. Door round trip 1.8–3.4ms.
- **Bar 2 — no credential where the agent reads: PASS 5/5**
  [measured], through the **real wrap machinery** at its published
  tag: mention wake → scripted harness (everything a real assistant is
  to wrap except the LLM) → door over stdio → per-call on-behalf →
  remote → outcome op carrying the remote's answer, ~2s end to end.
  The scan over every byte the run left on disk — including the
  harness's own dumps of its environment and argv — found zero of the
  provider's ever-minted access/refresh tokens; the planted control
  fired.
- **Bar 1 — usable without restart: FAIL by construction, baseline
  priced** [measured]. No mechanism exists; adding a resource is a
  plane restart costing 9.2–16.9ms of max probe gap, the added
  resource serving 5.5–9.2ms in. The pre-registered fail lane was not
  taken — the number became the bar the design's build must beat.
- **Bar 4 — where the catalog lives: ratified** [judgment, operator
  decision, twice amended by the operator's own questions]: **one
  record-borne discovery catalog for every tool** (the A10 pattern —
  display-layer resolution — applied to tools), custody split behind
  it: client secrets in the D36 store behind a new `resources.*` op
  family, workload declarations unchanged. Static config lost on the
  actor (a config file cannot be a button), not on its measured ~10ms
  price; the record lost custody of secrets; the first draft's
  identity-plane-only catalog lost the unified view; the literal
  persona registry lost on key-anchoring but contributed the
  two-layer pattern.

Refuted or corrected along the way: the token-in-config implementation
(by the topic's own Bar 2, before it was built); the topic's first
framing of "which door" (an adapter-position property, not a component
choice); and the first Bar 4 draft (the operator's unified-view
objection).

What it opened: the build items in the graduated designs —
[`soulstream-identity/external-tools.md`](../02-DESIGN/soulstream-identity/external-tools.md)
(D39 the two-layer catalog, D40 `resources.*` with Bar 1's pass as its
acceptance criterion, D41 the door invariants) and
[`soulstream-core/extensions/tool-catalog.md`](../02-DESIGN/soulstream-core/extensions/tool-catalog.md)
(the discovery convention) — all behind the 0071 focus gate, designed
now so demand has an answer. The rigs are the drafts of the standing
gates those builds inherit. The real-provider residue inherited from
episode 0104 closed mid-topic [operator-attested].

Reversal condition: D39's carries the direction — if the catalog's two
halves drift in practice faster than link-time failure surfaces it
(observable: recurring "tool isn't serving" against entries believed
live), the write becomes one transactional surface as a new
D-decision. For the door: if a remote class emerges whose protocol
cannot ride per-call token fetch (observable: a required tool whose
sessions outlive any fetchable token), the no-custody invariant
reopens by name rather than being quietly weakened.

Trail: designs
[`external-tools.md`](../02-DESIGN/soulstream-identity/external-tools.md)
and
[`tool-catalog.md`](../02-DESIGN/soulstream-core/extensions/tool-catalog.md);
the concluded topic (removed on graduation, full history in git —
opened `e4619ab`, the Bar 2 refutation `b1365e4`, the door dissolution
`200c883`, the operator's split `831c371`, Bars 3+1 `9874d2d`, Bar 2 +
the Bar 4 draft `4df45ab`, the amendment `843f081`, ratification
`31586cd`, verdicts `4435093`); episodes
[0116](0116-ecosystem-what-shipped-without-a-human-end.md) (the ask),
[0104](0104-ecosystem-outbound-identity-grants.md)/[0105](0105-identity-the-grants-broker-lands.md)
(the broker underneath),
[0109](0109-ecosystem-consent-enters-the-record.md) (the consent
vocabulary the rig exercised), [0083](0083-workloads-the-waker-lands.md)/[0085](0085-workloads-wrap-run-your-agent-where-you-are.md)
(the wrap machinery Bar 2 rode).
