# 0001 — soulstream-shell: the helm

**Status:** open — graduated from research 2026-08-13 (episode
[0066](../../04-JOURNEY/0066-ecosystem-soulsystem-cockpit.md)); every
load-bearing claim below was measured there unless tagged otherwise.
This document is written to be the argument to `/speckit-specify` in
the soulstream-shell repo (founded 2026-08-13, v0.1.0 — §10).

## §1 What soulstream-shell is

The soulsystem's human cockpit: one browser surface where a person
**observes** the whole running system — topics, turns, work items,
memory, plane health — and **configures** it, beside the MCP door that
serves machines. It is a *cockpit*, not a client: posting turns
(participation) is a named successor, out of scope here.

## §2 Placement — a sibling component (C1/C2)

- Its own repo (`github.com/impire-io/soulstream-shell`), own constitution, own
  gates — the archivist/soulstream-idp consumer-position pattern.
- A public **embed seam** `embed.Run(ctx, Options)` (the D29 pattern),
  which soulstream composes as `planes.helm` by tag — the fourth URL
  logged at `up` — plus a standalone `soulstream-shell serve` for realms
  running the components without soulstream.
- **Pure consumer, structurally guaranteed**: the e2e gate is a module
  whose path sits outside every component namespace (the
  `e2e/embedgate` shape), pinning all upstreams by tag, zero `replace`.
  Measured: the entire observe surface needed zero upstream additions.
- Two candidates were refused by constitutions, not preference: the
  fold's console (a Soulstream-aware IdP collapses the two planes) and
  soulstream-native code (composition, not invention)
  [mechanism-argument].

## §3 The observe surface

All reads ride a backend-held `realm.Client` (read-only config — no
persona) against the realm, plus the soulstream-identity `client` for the
public key directory:

- **Board** — every topic with lifecycle (`topic.Board`).
- **Topic view** — `Materialise`/`Follow`: turns, comments, work items,
  attachments, with signature verdicts *earned* (keyring built from
  `keys.public`; no keyring means `unknown-key` shown honestly).
- **Plane health** — door liveness (HTTP on its published URL), fold
  liveness, storage budgets, workload lifecycle read from the work
  vocabulary.
- **Memory** — *participation, not observation* [measured]: the query
  is a posted op, so the memory panel exists only inside a signed-in
  session and asks as that principal.

## §4 The configure surface (C6, Bar 3)

Every mutation maps to one of exactly three classes — the enumeration
is normative; a mutation that fits none is a design change, not a
feature:

| Class | Mechanism | Examples |
|---|---|---|
| (a) op on the record | the signed-in principal's own realm client | work open/claim/complete/abandon; topic lifecycle transitions; workload start/stop via the runner vocabulary |
| (b) admin op on an existing plane surface | soulstream-identity token ops; soulstream-idp `/api/admin` | mint/revoke tokens; people, invites, groups; OAuth clients |
| (c) config-file change with stated restart semantics | `config.json` `planes.*`; apply-on-restart | enable/disable a plane; door public mode (all-three-or-none) |

Founding-only fields (`listen`, `realm`, listener addresses) are not
helm mutations — soulstream refuses them post-founding by design. No
helm mutation may require a new privileged side-channel or a helm-owned
store [measured: Bar 3 held it].

**The lane constraint** [measured]: class (b) identity-plane ops are
refused to the owner's admission-lane connection (op-tail enforcement,
D25). As a soulstream plane the helm may use the node's operator
standing; standalone, it needs the state dir's ops credentials — until
**upstream ask #2** (an owner-reachable token-management lane or a
delegated grant, converging with `platform-tenancy-guardrails`' grant
model) resolves. The helm never works around this with a side-channel.

## §5 Rendering — Datastar over SSE (C4, decided on the bench)

- Backend is source of truth: the helm renders HTML server-side and
  pushes `datastar-patch-elements` morphs over SSE; one vendored
  ~34 KB bundle; page-local JS only. Measured live, including across
  an in-place node restart.
- The **NATS-WebSocket lane (C5)** remains the decided browser
  transport for the *participant client* horizon — **upstream ask
  #1**: soulstream's embedded server grows a WebSocket listener (one
  options field). The helm does not gate on it.
- Build lessons of record: the Datastar release bundle (npm `+esm`
  does not self-initialize); `data-init` / `data-on:click`; one-shot
  SSE responses race on shared targets — mutation results get their
  own target elements.
- Reversal: an interaction needing client-held state or sub-roundtrip
  latency SSE morphing cannot express reopens this pick with that
  interaction as evidence.

## §6 Identity, sessions, custody (Bar 2)

- **Sign-in is the fold** (or any OIDC AS the deployment configures):
  the helm registers itself via RFC 7591 DCR, runs code+PKCE, and
  verifies the ID token. No pre-provisioned client.
- **Delegated authority, never borrowed identity (S6)**: each session
  opens its *own* NATS admission — sentinel + the user's fold-issued
  bearer through the OIDC callout lane — and every act rides that
  connection, signed by the principal's own persona key (materializing
  on first touch). The helm signs as no one.
- **Custody**: sessions in memory only; nothing credential-shaped in
  helm storage — the standing e2e gate repeats Bar 2's scan with a
  positive control that must fire.
- Attribution rides soulstream-idp's persona-shaped ids (`u-…`); the helm
  owns the id→display-name mapping on screen [measured].

## §7 The design-system contract (C3, C8, Bar 4)

- Canon: **cassette futurism in a light key** — the Soulsystem Design
  System project is the source of truth; the helm consumes it as a
  single token/component CSS source, fonts and icons vendored, zero
  CDN (the offline render gate is standing).
- Component identity by **label** (mono strips, wordmarks), never a
  second color system; amber/teal stay strictly the human/machine
  channel semantic.
- **Copy (C8)**: surfaces say what things do — Storage, Connections,
  People & sign-in, Agents, Active topics, Latest activity. Component
  bynames never appear in product UI.

## §8 Acceptance criteria (the research bars as standing gates)

1. The consumer-position module compiles against pinned tags only —
   zero `replace`, `internal/` imports impossible by module path.
2. After a full sign-in→act→sign-out session, the helm's storage scans
   clean of credential shapes, with a fired positive control.
3. Every shipped mutation appears in the §4 table with its class; one
   e2e per class stays green.
4. The helm's screens and one fold page render from the shared token
   source with zero external fetches and fonts verifiably loaded.

## §9 Open questions [O]

- **[O1]** Upstream ask #1 timing: the WebSocket listener lands in
  soulstream ahead of the participant-client topic, or with it.
- **[O2]** Upstream ask #2: the grant-shaped answer to standalone
  class-(b) authority — tracked in `platform-tenancy-guardrails`.
- **[O3]** ~~The id→display mapping source~~ — first answer as built
  (§10): the ID token's `name` claim, id fallback; richer mapping open.
- **[O4]** The participant client — its own research topic, not helm
  scope.

## §10 As built (v0.1.0, 2026-08-13 — episode 0067)

Propagated from the founding build:

- **Sessions shipped with the founding release** (open amendment to the
  M1/M2 split): the surface is closed until sign-in — an
  unauthenticated realm viewer would contradict §6.
- **The soulstream composition** (episode 0068): the plane hands the ops
  lane as the helm's read lane (§4's node-standing arm, as designed;
  the dedicated scoped user is a named hardening), the public sentinel,
  and `SessionIssuer()` — enabling the helm switches the identity
  plane's OIDC lane on in local mode with the bundled fold as AS.
- **O3, first answer**: display names come from the ID token's `name`
  claim, falling back to the persona-shaped id; a richer directory
  mapping stays open.
- The DCR redirect URI carries the *bound* listener address (ephemeral
  ports work).

