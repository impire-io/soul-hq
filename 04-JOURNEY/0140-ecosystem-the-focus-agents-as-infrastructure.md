# Episode 0140 — The focus: agents as infrastructure (2026-08-27)

The operator redirected the effort the evening the declaration story
became installable (episode
[0139](0139-soulstream-the-rc2-carries-the-tenants-and-the-capabilities.md)):
the next things built are **a standing dispatcher allowing
submit-and-forget** and **a shell surface for declaring agents**,
because that is what real users report missing `[judgment — the
operator's read of demand, the evidence class this decision can have]`.
The declaration itself is whole — wake, instructions, capabilities,
budget, one JSON contract, v0.14.0-rc.2 — but an agent still runs only
where a person runs `soulstream wrap`: it lives and dies with a
terminal. What people want is to hand the realm a declaration and walk
away.

Nothing here is a reversal of the record — it is the record's own
clause firing. Design
[`0004-wrap.md`](../02-DESIGN/soulstream-workloads/0004-wrap.md) §9 cut
the central daemon *with* its reversal condition: "a realm operating
agents as infrastructure — nobody's laptop, centrally credentialed —
brings the serve arm back as a fleet-era feature over the same engine,
with design 0003's placement answering which node runs a wake." That
observable has now been reported from outside `[mechanism-argument:
the condition was written for exactly this signal]`. Design
[`0006-loop-safety.md`](../02-DESIGN/soulstream-workloads/0006-loop-safety.md)
§6 pre-paid the safety half: the wake budget was built to sit at any
future dispatcher's admission seam.

What the decision opened, in order:

- Research topic **`agents-as-infrastructure`** (ecosystem), opened the
  same evening with five pre-registered bars: submit-and-forget
  exactly-once across dispatcher restart, placement by the 0003 claim
  path across two nodes, the 0006 budget measured at the dispatcher's
  admission (not assumed from §6), provider/model credentials custodied
  in the identity plane and unreachable from the agent's scope, and the
  whole declare→submit→served→answer loop drivable from the shell's
  pure-consumer position.
- The **inference providers and models** question, named as part of
  this storyline at the operator's direction: wrap dissolved it by
  wrapping the person's own signed-in assistant; a dispatcher-served
  agent has no signed-in person, so what serves its thinking must be
  declared and custodied. Bar 4 carries it; if it outgrows one bar it
  becomes a named successor topic.
- The **shell declare surface** is the storyline's human end, designed
  against the dispatcher's submit op once the research fixes that op's
  shape — Bar 5 exists so the shell's one-way door (pure consumer) is
  proven reachable before that design is written.

What steps back, named honestly: the sealed-topics product wiring —
next in the queue this same afternoon — moves behind the dispatcher
arc; it stays a recorded [O] in
[`sealing-keys.md`](../02-DESIGN/soulstream-identity/sealing-keys.md)
and loses nothing by waiting. The 0071 focus (the usable cockpit,
stdio MCP as the agents' door) is not repealed — this is its next
chapter, not its replacement: the cockpit gains the declare surface,
and stdio MCP remains how a wrapped or dispatched agent participates.

Reversal condition: the research topic's own two readings govern — if
every viable harness demands interactive per-person authentication no
custodied credential satisfies, or if submit-and-forget cannot hold
exactly-once without a store of record beside the log, the serve arm
does not return in this shape and the finding is recorded where the
demand can see it. The focus itself reverses only on the operator
reading the same demand differently.

Trail: research pre-registration `01-RESEARCH/agents-as-infrastructure/`
(`7f895c1`; graduated and removed the next morning — [episode
0141](0141-ecosystem-agents-as-infrastructure.md), git history keeps
the folder); design 0004 §9 (the fired clause), design 0006 §6 (the
waiting seam); the demand-side context in episode
[0116](0116-ecosystem-what-shipped-without-a-human-end.md)'s method —
evaluation finds what is built but unreachable by a person.
