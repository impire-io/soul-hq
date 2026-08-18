# Episode 0109 — Consent enters the record, and Bar 4 finally measures (2026-08-18)

The C4 build (soulstream-core **v0.10.0**): `grant.issue` and
`grant.revoke` as additive vocabulary, the Grants projection folded
exactly like work items — granter-only revocation with everything the
state machine rejects kept visible as void — dual exercise attribution
(`TurnPayload.Authority` names the grant and its granter, so the acting
persona and the granting persona are each readable without inferring
one from the other), and standing grants **baked through rollup**: a
compaction that dropped consent would erase authority its issuer never
withdrew. Expiry deliberately stays a clock fact — the projection is a
pure function of the log, `ActiveAt` is where *now* enters — which the
suite pins with an issued-already-expired grant reading `active` in the
log-derived state and inactive at now [measured].

**Bar 4 — the tenancy topic's one unmeasured bar — now measures PASS**
[measured, consumer position]: on the full composition (core's consent
record + the identity plane's delegation broker + a minting surface
that consults the projection before minting — the S8 split exactly),
all three clauses held in one run of the embedgate gate: the granted
action served with both personas attributed on *both* surfaces (the
broker's on-behalf audit and the record's authority field); revocation
refused the next mint on the projection and the in-flight delegation
died at its TTL bound (the honest boundary D33 named); and neither
persona's own standing was disturbed — the subject's own access served,
the actor posted as itself. The represented-user scope template grew
the realm half for the composition, which is what the product's
ceremony grants a represented user anyway.

Reversal condition: episode 0107's stands — if the consent record
cannot make Bar 4's third clause true in a real deployment (observable:
the delegation matrix failing that row outside the rig), the B10
one-mechanism consolidation reopens. The TTL bound on in-flight
delegations is a named semantic, not a gap: tightening it to
immediate-kill would need delegation introspection per access, a new
D-decision.

Trail: design
[`extensions/tenancy.md`](../02-DESIGN/soulstream-core/extensions/tenancy.md)
(the vocabulary and the C7 replay duty); core `bb4e00e` (v0.10.0),
identity `dbf948a` (the Bar 4 gate,
`e2e/embedgate/bar4_test.go`); episodes
[0107](0107-ecosystem-platform-tenancy-guardrails.md) (Bar 4 carried as
this build's gate), [0105](0105-identity-the-grants-broker-lands.md)
(the enforcement half).
