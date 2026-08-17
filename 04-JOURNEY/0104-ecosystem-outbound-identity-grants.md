# Episode 0104 — Outbound identity: every remote call carries the calling user (2026-08-17)

The question, opened after an evening's design conversation and measured
overnight at the operator's direction: can a persona's soulstream
identity be exchanged, at call time and only by that persona (or its
provably delegated agent), for a remote-system credential — without any
long-lived secret leaving identity-plane custody? The inbound door
solved its half by holding nothing (the remote node, episodes
0038/0047); outbound was still the industry default the topic exists to
refuse: one credential in the agent's MCP config, every persona
collapsed into a single remote user. Three pre-registered bars, all
measured on real components; full verdict numbers in the graduation
commit's topic README.

**Bar 1 — soulstream-to-soulstream, PASS** [measured]: the real fold
(embed seam, virtual-passkey ceremonies) stood as the external AS
behind the real identity plane and the real remote node — the 0054/0055
compositions finally run *multi-user*. Two humans enrolled by passkey,
each admitted through the door as exactly its own persona (`u-hex`
oids), one topic carrying a turn from each with **zero
cross-attributions** on an independent reader; 10/10 writes across
2.5× a 5s callout TTL; a dead bearer refused in **12.6ms** against a
2×TTL bound — the node's freshest-wins pool re-probes immediately,
refusal is not TTL-paced. No claim-profile work was needed: the fold
already speaks the validator's vocabulary, which is 0054's
indistinguishability decision paying out.

**Bar 2 — grant custody, PASS on the registered stand-in protocol**
[measured]: a broker prototype custodies per-persona OAuth grants
against Dex (rotation on). The scripted PKCE linking dance completes;
3/3 refresh rotations CAS-persist the successor; an 8-way concurrent
double-refresh under `-race` loses nothing; revocation deletes custody
and the next access refuses. The clause that matters most was
server-enforced, not broker logic: **a persona asking for another's
grant dies as a publish permission violation** — the broker's delivery
log shows the victim's subject reached it exactly once (its own
request). The pre-registered residue stands open: one real-provider
confirmation (GitHub/Google link + rotation + revoke), a morning act.

**Bar 3 — delegated acting, PASS** [measured]: on-behalf-of exists only
as a minted, bounded delegation — subject, actor, resources, scopes,
TTL — minted only against standing, revocable consent. One allowed path
against four refusal classes, every refusal audited naming both
personas; and on the wire the actor is the server-proven subject-token
principal, so a *stolen, validly-signed* delegation refuses as an actor
mismatch. What the rig deliberately did not close: the minting home —
episode 0082's runner was cut two days later (0083→0085, the wrapper
mints nothing), so the design argues wrapper-mint-restoration against
**subject-signed delegations** verified from the D26 key directory
[judgment].

Refuted or corrected along the way: the topic's own premise that the
remote-node prototype needed recovering (it shipped — soulstream-mcp
v0.1.0, the product's public door); the assumption the key vault could
custody grants (its records are immutable by design — Create-only seam,
closed kind enum, derive-requires-public-key — so grants need their own
custody domain, a numbered decision, not a case-arm) [measured, code
map]; and a first Dex reading of "no rotation" that was the rig
refreshing inside the 1s reuse interval — provider retry semantics, not
a rotation failure [measured].

What it opened, cross-component and named: the fold's **token-lifetime
knob is now demanded** — outbound revocation propagates in access-token
`exp` + callout TTL, and 0103's watch item ("per-client token-lifetime
knobs stay deferred until a milestone demands them") has met its
milestone [measured]; the fold's **RFC 8693 token exchange** gap is
measured from both ends (the idp refuses the grant type; audience is
deployment-fixed), so a persona reaching several remotes from one
session needs the exchange grant grown behind an Article II argument —
Entra itself speaks OBO, so indistinguishability is arguable; the
identity plane needs the **grants surface and custody domain** (design
D30–D34); and the wrapper needs its **per-run overlay seam restored**
to carry a delegation into a run. The design graduates to
[`grants.md`](../02-DESIGN/soulstream-identity/grants.md).

Reversal condition: if any lane can only be made to work by placing a
long-lived secret (refresh token, client secret, seed) in agent,
workload, or MCP-client configuration — outside identity-plane custody
— the central-broker direction is wrong for that lane and reopens
(observable: a passing rig wiring that requires such a secret in client
config). If the MCP specification ships a native per-user
identity-propagation mechanism adopted by a mainstream host
(observable: a published spec revision plus one shipping host), the
linking-ceremony lane is re-argued against it before its build
continues.

Trail: design [`grants.md`](../02-DESIGN/soulstream-identity/grants.md)
(D30–D34); the concluded research topic (removed on graduation, full
history in git — opened `d1cc57f`, measured `a0c57ef`, this graduation
commit); episodes [0038](0038-soulstream-remote-mcp-node.md)/
[0047](0047-soulstream-remote-mcp-node-built.md) (the inbound door),
[0054](0054-soulfold-m4-the-fold-in-the-fleet.md)–
[0057](0057-soulnode-the-folded-realm.md) (the compositions Bar 1
extends), [0082](0082-ecosystem-agent-participation.md) (the acting
substrate), [0103](0103-ecosystem-the-session-outlives-its-token.md)
(the refresh grant and the watch item this episode calls due); rig
modules in the session scratchpad per how-we-work.
