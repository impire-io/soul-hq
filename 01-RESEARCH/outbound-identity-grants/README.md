# How does a persona's identity cross outward — so every remote MCP or API call carries the calling user, not a shared credential?

**Component:** ecosystem
**State:** active
**Started:** 2026-08-17

## Abstract

Inbound identity is solved: the remote MCP node holds no credentials and
the auth callout decides who each connection is. Outbound is not — an
agent or shell calling a remote MCP server or third-party API today would
do it the way the ecosystem's peers do, with one credential in client
config, which collapses every persona into a single remote user. This
topic investigates the outbound inversion: no credential ever lives in
agent, workload, or MCP-client config; all outbound grants live behind a
broker surface on the identity plane where the principal is
server-asserted, so the only token any caller can obtain is its own. A
decisive answer unlocks the remote node's open decision 018 (the external
OIDC AS is soulstream-idp), a grants surface on soulstream-identity, and
a provable on-behalf-of story for agents.

## The question

Can a persona's soulstream identity be exchanged, at call time and only
by that persona (or its provably delegated agent), for a remote-system
credential — across three lanes: soulstream-to-soulstream (idp as the
external AS), third-party OAuth (broker-custodied grants), and
runner-delegated agent acting — without any long-lived secret leaving
identity-plane custody?

## Pre-registered bars

- **Bar 1 — soulstream-to-soulstream, per-user, both ends ours.**
  *Protocol:* a rig wires a NATS server (local topology mirroring the
  BYON AUTH shape; impire-dev BYON authorized as the confirming
  environment) with soulstream-identity's callout OIDC lane pointed at a
  local soulstream-idp as the issuer (its JWKS, its claims — adding an
  issuer-claim-profile config to the validator if idp's claim names
  differ from the Entra lane's, recorded as a finding, not a bar
  amendment). Two personas with distinct idp accounts mint access tokens
  and connect through the remote-node bearer path. *Pass:* persona A is
  admitted as A and persona B as B (server-asserted via
  `$SYS.REQ.USER.INFO`, 2/2); a write by A is attributed to A and never
  to B (0 cross-attributions); an expired or revoked token is refused at
  the next reconnect within the callout TTL bound.
- **Bar 2 — grant custody against a stand-in third party.**
  *Protocol:* a broker prototype on the identity plane's NATS surface
  custodies per-persona OAuth grants against a local stand-in AS (Dex or
  Ory Hydra in docker, refresh-token rotation ON). The linking dance
  (auth-code + PKCE) completes for persona A through broker ops;
  `grants.access` returns a short-lived access token. *Pass:* A's
  connection obtains A's token (verified live against the stand-in's
  userinfo/introspection); B's connection is refused A's grant **by
  subject-space permission**, not broker logic (the request never
  reaches the broker); ≥3 consecutive provider-side refresh rotations
  survive with the CAS-stored successor each time and a deliberate
  concurrent double-refresh loses no grant; `grants.revoke` deletes the
  record and the next `grants.access` refuses. *Named residue:* the
  stand-in substitutes for a real provider (Daan asleep; consent needs
  the account owner) — a real-provider confirmation (GitHub or Google,
  one link + one rotation + one revoke) is a pre-registered morning
  step, and Bar 2 is not fully closed until it runs.
- **Bar 3 — on-behalf-of exists only as a minted, bounded delegation.**
  *Protocol:* the runner (episode 0082's trust anchor — it mints for the
  agent, never the agent for itself) mints a per-run delegation grant:
  subject = the asking persona, actor = the agent persona, bounded
  resources, scopes, and TTL. *Pass:* the agent's
  `grants.access(resource, on_behalf_of=subject)` with a valid
  delegation returns the subject's token bounded to the delegated
  scopes; the same call without a delegation (or with an expired one) is
  refused and produces an audit record naming both personas (2/2
  refusals, 2/2 audit records); revoking the standing consent refuses
  the next delegation mint while past attribution stays intact.

## Reversal condition

If any lane in the rig can only be made to work by placing a long-lived
secret (refresh token, client secret, seed) in agent, workload, or
MCP-client configuration — outside identity-plane custody — the
central-broker direction is wrong for that lane and reopens (observable:
the rig's config file containing such a secret as the only passing
wiring). Independently: if the MCP specification ships a native per-user
identity-propagation mechanism adopted by a mainstream host (observable:
a published spec revision plus one shipping host), lane 2's
linking-ceremony design is re-argued against it before implementation
continues.

## Verdict

- **Bar 1 — PASS** [measured, local rig per the registered protocol]:
  the real fold (embed seam, virtual-passkey ceremonies) as the
  external AS behind the real identity plane and the real remote node;
  ada admitted as `u-cd00f20fc32ba2a3`, grace as `u-aed28616fbe46d78`
  (2/2, each exactly itself), one topic with a turn from each and **0
  cross-attributions** on an independent reader; 10/10 writes across
  2.5× a 5s callout TTL on a living bearer; a dead bearer refused in
  **12.6ms** (bound 2×TTL); no claim-profile change needed. Honest
  finding, not a failure: refresh-token revocation does not refuse a
  signature-valid access token — revocation propagates in
  access-token `exp` (1h, package constant) + callout TTL, so the
  fold's deferred token-lifetime knob is now demanded [measured].
- **Bar 2 — PASS on the registered stand-in protocol; the
  pre-registered real-provider morning step remains open** [measured]:
  the scripted PKCE linking dance completes against Dex; 3/3
  consecutive refresh rotations with the CAS-stored successor; an
  8-way concurrent double-refresh under `-race` loses nothing (8/8
  served, stored token still redeemable, both reuse-window regimes);
  revocation deletes custody and the next access refuses;
  **the principal clause is server-enforced** — bob's request for
  alice's grant died as a publish permission violation and the broker
  received alice's subject exactly once. Named residue: one real
  provider (GitHub or Google) link + rotation + revoke, a morning act.
- **Bar 3 — PASS** [measured, -race]: delegations mint only against
  standing consent covering every requested resource and scope; 1
  allowed vs 4 refusal classes (absent, expired, wrong caller,
  out-of-bounds resource), every refusal audited naming both personas
  (5/5); consent revocation refuses the next mint with the audit
  intact; and on the wire the actor is the server-proven
  subject-token principal — a *stolen, validly-signed* delegation
  refused as an actor mismatch. The delegation's minting *home*
  (wrapper mint-lane restoration vs subject-signed via the D26 key
  directory) is a design argument, deliberately not closed by the rig
  [judgment, argued in the design doc].
