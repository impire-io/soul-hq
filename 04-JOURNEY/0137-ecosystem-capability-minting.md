# Episode 0137 — Capability minting: the declaration's names become the credential (2026-08-27)

The named follow-on of the agent-declaration build (episode 0130) is
built, across three repos in one arc. A declaration's `capabilities
{role, tools[]}` — schema-only since workloads `specs/009` — now
reaches the minter seam and comes out the other side as a credential
the server clamps: workloads `specs/010` widens `minter.Scope` with the
selectors, renders the mint-tag vocabulary in one refusing surface
(`MintTags` — a value that could alter subject grammar never leaves),
and narrows the local minter's `SOULSTREAM.SVC.>` wildcard to exactly
the declared tools; identity `specs/004` exports the canonical agent
scope beside the persona scope (`AgentScopePubAllow/SubAllow` — the
workloads agent derivation with the dynamic parts as tag functions,
notify riding `{{name()}}` so reachability cannot drift from
attribution); the product's `specs/013` founds the **agent capability
key** — a scoped signer on the realm account rendering that template —
and routes capability-bearing declarations through the scoped lane.
Measured end to end [measured]: identity's e2e turned research 0126's
rig into a standing gate and closed the arc's two first-verify risks —
**multi-value tag expansion** (two `tool:` tags, both tools answer
through one credential; a third refused with zero responder
deliveries) and the **zero-matching-tag line drop** (a tool-less mint
still admits and reaches nothing); the product's `TestM14` ran
upstream's scope-probe both ways — granted its own subject, the run
completes `done`; granted a *different* tool, the probe's own publish
dies at the server ("Permissions Violation for Publish to
SOULSTREAM.SVC.probe-ping") and the run ends in the runner's abandon —
the narrowing bites through the full authority chain, with zero
authorization code in any runtime (design 0005 §10 #3).

One plan was refuted by the wiring, and the reversal is the episode's
load-bearing find. Design 0005 §5's letter — "the runtime calls D28
`mint.ephemeral` with the declared role name" — assumed the agent role
key could enter the identity vault beside the realm's persona role.
It cannot, **by decided design**: the binding-resolved lanes
(`RoleForAccount` — durable mint, token creation, callout issuance)
refuse a multi-role account as ambiguous (D5 as amended; identity's M3
gate proof 6 measures exactly this refusal), so the import would have
broken every token-lane sign-in on the deployment
[mechanism-argument, anchored in the M3 measurement]. The resolution
is an open amendment recorded in workloads spec 010: the scoped lane
ships **local-first** — `ScopedSigningKeyMinter` produces the
identical D28 claim shape (permission-less scoped user, tags, TTL)
signed by the state-held agent seed; enforcement never leaves the
server. The D28 op lane remains the fleet-era path for seedless
nodes, gated on identity's already-named "token lane's named-role
answer" — the collision is the reason that item exists.

What it opened, named: the Synadia BYON arm (the driver's fourth
signing-key group with the agent template) waits on a live run — the
0136 discipline — so a BYON realm today refuses capability
declarations by the same named refusal as a pre-capability realm; the
tag-policy watch now has its first consumer (fires when submission
opens beyond operator-grade trust); descendant tag scoping stays
untested; and the D47 persona-scope adoption landed as a rider (the
product ceremony now renders `client.PersonaScope*` instead of its
local copy — byte-identical, one source). Standing exception, the
0089 precedent: the product rides sibling trees via untracked
`go.work` until workloads/identity mains push; then both pins bump and
`go mod tidy` closes it.

Reversal condition: if a deployment needs capability minting where the
agent seed cannot live beside the runtime (observable: a fleet node
without the state directory needing to launch capability-bearing
declarations), the D28 op lane builds behind the token lane's
named-role answer — the local lane never grows a seed-shipping
workaround.

Trail: workloads `specs/010-capability-minting` (`9fa70a9`/`8d4475f`,
merged `14e95a0`; the scoped-lane amendment `2465dba`); identity
`specs/004-agent-scope` (`26a6e87`, merged `e032687`); soulstream
`specs/013-capability-minting` (`60b5727`, merged `baf5a7e`); designs
[`0005-agent-declaration.md`](../02-DESIGN/soulstream-workloads/0005-agent-declaration.md)
(§5 as-built note) and
[`0003-fleet.md`](../02-DESIGN/soulstream-workloads/0003-fleet.md)
(§5's stale tags-[O] resolved); episodes
[0126](0126-ecosystem-agent-declaration.md)/[0130](0130-ecosystem-the-agent-declaration-builds.md).
