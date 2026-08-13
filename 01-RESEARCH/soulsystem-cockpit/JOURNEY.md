# soulsystem-cockpit — investigation journey

Topic opened 2026-08-13. Appended as the investigation happens.

## 2026-08-13 — C3 decided: cassette light is canon, identity by label

The operator decided both halves of C3, ahead of Bar 4 as that bar
requires.

**Canon: the cassette-futurism light system** (the Soulsystem Design
System project) becomes the ecosystem's visual language `[judgment]`.
The deciding argument: the amber/teal rule — human channel and machine
channel as two accents at deliberately equal weight, neither allowed to
outrank the other — is the vision's "humans and AI as peers" turned
into a color system; the violet/rose dark theme is semantically mute
beside it. Supporting the call: it is the only *complete* system of the
two (tokens, seventeen components, guidelines, voice canon, two UI
kits; `webstyle` is one CSS string), and the migration cost is at its
historic minimum — episode
[0063](../../04-JOURNEY/0063-soulfold-the-console-gets-a-face.md) was a
presentation-only change with reversal condition *none*, and the
impire.io website rework is already in flight (episode
[0065](../../04-JOURNEY/0065-ecosystem-fair-code.md)), so the site
adopts the canon rather than being repainted twice.

Costs accepted openly: this is a **committed light-only look** — "one
dark surface: CRT glass" is a core rule, and there is no dark mode; a
future dark key would be a deliberate amendment to that rule, never a
drift. The fold's four shipped pages get restyled (soulfold repo work,
sequenced later — Bar 4 needs only one fold page rendered from the
shared source as proof); the violet/rose language is retired from the
soulsystem. Middle paths were examined and rejected: re-keying the
design project dark breaks its own central rule (the CRT surface reads
as special only because the world is light), and a dual canon makes
every future surface pick a side forever.

**Per-component identity rides labels, not color**
`[mechanism-argument]`: amber/teal are *channel* colors (who is
speaking), the rose-edge idea was a *component* color (where you are) —
two color systems on one screen fight, and the equal-weight rule cannot
survive a third accent competing beside it. Component identity uses the
register the system already owns: mono label strips, wordmarks, and
badges ("the fold", "the house"). The rose edge retires with the dark
language.

## 2026-08-13 — Bar 3 measured: PASS — three classes from the browser; Bar 1 complete on both halves

The helm prototype (scratchpad `bar1/rig/proto/`, same consumer-position
module, pinned tags): a Datastar page over a live in-process realm with
the three pre-registered mutation classes wired as buttons. Everything
below happened in a real Chromium against the running system [measured].

**Bar 1's browser half, closed on the decided path**: the dashboard
renders live realm state in the browser — NATS and door URLs, topic
count, turns, work items — through the backend-held realm client and
`datastar-patch-elements` morphs, staying live *across an in-place node
restart*. With spike A's compile guarantee and read coverage, **Bar 1
is met in full**; the C5 WebSocket lane remains the participant-client
horizon and upstream ask #1 stands on its own.

**Bar 3, class by class:**

- **(a) op on the record — PASS.** `work.open` from the browser button;
  the item lands signed on the topic and the live dashboard shows it
  within a second (`work "spike work item" · open`).
- **(b) plane admin op — PASS, with the spike's sharpest finding.**
  The founding owner's admission-lane connection is **refused** for
  `tokens.create` — an explicit server-side `Permissions Violation` on
  the op tail `soulidentity.<account>.owner.tokens.create` in the
  audit log: D25's enforcement working exactly as designed, observed
  from the consumer position. The node's ops lane then mints the token
  (`sit_…` returned once, issue audited `user=ops`,
  `target_user=helm-guest`). The mutation maps to an *existing*
  surface — no new side-channel — but **the lane question is real**:
  a helm running as a soulnode plane inherits the node's operator
  standing; a standalone helm would need state-dir creds or a
  delegated grant. That is upstream ask #2 (an owner-reachable
  token-management lane), and it lands squarely in
  `platform-tenancy-guardrails`' grant work — C7's parallelism is
  already paying.
- **(c) config file — PASS.** `planes.memory.enabled=false` written to
  `config.json`, the node stopped and restarted in-place (the stated
  semantics), clients auto-re-admitted (callout `ADMITTED` in the
  audit), and the memory probe reads **0 answers where the same query
  read 5 before** — the plane is verifiably gone, and the dashboard
  never stopped rendering.

**Bar 3 verdict: PASS** — table enumerated before the run (previous
entry), one demonstration per class green, and no mutation in the
table required a new privileged side-channel or a helm-owned store.

**Bar 2 status: not run.** Its ceremony needs the helm to authenticate
the browser user against the bundled fold (OIDC + passkey enrolment
through the printed invite, virtual-authenticator-driven), act as that
principal, sign out, then the storage scan with positive control. One
preview finding is already on the record: the spike's own stdout log
captured the fold invite (`sfi_…`) — precisely the class of leak the
Bar 2 scan exists to catch, and a ready-made positive control.

**Datastar lessons banked for the build**: the npm `+esm` build does
not self-initialize (use the release bundle); the v1 attributes are
`data-init` and `data-on:click`; concurrent one-shot SSE responses
race on a shared target element — last write wins — so mutation
results need per-action target ids or the `/live` channel.

## 2026-08-13 — Bar 3's mutation table, written before the spike runs

Every mutation helm v1 offers, mapped to its class as Bar 3 requires —
enumerated here first, demonstrated after. Participant acts (posting
turns, comments) are outside v1 by C6 and deliberately absent.

| Helm v1 mutation | Class | Existing surface |
|---|---|---|
| Open / claim / complete / abandon a work item | (a) op on the record | `topic` work vocabulary via a realm client, signed as the acting principal |
| Close / transition a topic | (a) op on the record | `life.transition` op, same client |
| Mint / revoke / list access tokens | (b) admin op on a plane surface | soulidentity token ops (`CreateToken` / `RevokeToken` / `Tokens`) |
| People: create, invite, groups, disable | (b) admin op on a plane surface | soulfold `/api/admin` (passkey-session gated) |
| Register an OAuth client | (b) admin op on a plane surface | soulfold `/api/admin` |
| Enable / disable a plane (memory, door, fold) | (c) config file | `config.json` `planes.*`; applies on node restart — the stated semantics |
| Door public mode (`public_url`+`auth_issuer`+`auth_audience`) | (c) config file | `config.json`, all-three-or-none, restart |
| Start / stop a workload | (a) op on the record | work vocabulary via the runner (soulrealm); *named, not demonstrated in this spike* |

Founding-only fields (`listen`, `realm`, listener addresses) are not
helm mutations — soulnode refuses them post-founding by design. The
spike demonstrates one mutation per class: `work.open` (a),
`CreateToken` (b), `planes.memory.enabled=false` (c). No mutation in
the table needs a new privileged side-channel or a helm-owned store —
that claim is what the spike now has to hold up.

## 2026-08-13 — Bar 1 completed on the Go half; spike B measured; the rendering discrimination decided

**Bar 1's remaining reads landed on the spike-A rig** [measured]: a
claimed work item materializes with `status=claimed · owner=owner`; the
memory plane answers a live query (5 archivist answers, each with a
citation); the door answers HTTP on its published URL (400 to a bare
GET — the MCP endpoint refusing a non-MCP request, which is the
liveness signal). One finding of record: **asking memory requires a
persona** — `MemoryQuery` refuses a persona-less client because the
query is itself a posted op. The helm's memory panel is participation,
not observation; it must ride the signed-in principal's client
[measured, mechanism in the memory convention].

**Spike B — the WebSocket bench** (scratchpad `bar1/wsbench/`): an
embedded nats-server v2.12 with its WebSocket listener enabled (the
knob is one field in the embedded server's options — upstream ask #1
is plumbing, not invention), one publisher emitting soulstream-shaped
ops every 700 ms, both candidate rendering paths beside it, everything
vendored and loopback-only:

- **Web component + `nats.js` (browser speaks NATS):** connected over
  `ws://`, subscribed, rows streaming live in a real Chromium
  [measured]. Cost observed: a four-bundle vendored chain
  (nats-core + nkeys + nuid + tweetnacl, ~130 KB) with import-path
  surgery, and the realm credential must live in browser JS (the
  MCP-client trust class — acceptable under S6, but real).
- **Datastar (backend-held subscription, SSE patches):** the same feed
  live via `data-init="@get('/feed')"` and
  `datastar-patch-elements` morphs [measured]. One 34 KB
  self-initializing bundle, zero import surgery, zero credentials in
  the browser; the backend is spike A's already-proven typed read
  path. One protocol lesson recorded: the npm `+esm` build does not
  self-initialize and the load trigger is `data-init` (v1), not
  `data-on-load`.

**The discrimination — decided** (delegated; `[judgment]` on measured
evidence): **the helm renders through Datastar** — backend as source
of truth, SSE patches, spike A's consumer-position read path as the
data layer, credentials never in the browser (Bar 2's custody property
falls out for free), and the house no-pipeline pattern kept (one
vendored bundle). **The NATS-WebSocket lane stays decided (C5) for
what it is**: the realm's browser transport for the *participant
client* horizon and any client that genuinely needs raw subjects —
upstream ask #1 stands, but the helm itself no longer gates on it.
Reversal condition: if a helm interaction needs client-held state or
sub-roundtrip latency that SSE morphing cannot express, this pick
reopens with that interaction as the evidence.

## 2026-08-13 — Bar 1, spike A measured: the consumer position compiles and reads live

The Go-shell half of Bar 1, run the same afternoon (rig in the session
scratchpad, `bar1/rig/`):

- **The module**: `soulhelm.invalid/rig` — outside every component's
  namespace, so an `internal/` import cannot compile — requiring
  soulnode **v0.3.1**, soulstream **v0.7.0**, soulidentity **v0.1.0**
  by tag (soulfold v0.3.3, soulrealm v0.1.0, the archivist v0.2.0
  arriving transitively, all tagged). **Zero `replace` directives**
  [measured]; `go vet` clean.
- **One run, ~1.5 s wall**: boot a whole realm in-process from
  soulnode's public `ceremony`/`node` surfaces (ephemeral ports) →
  found it (`sit_` token) → connect as the founding owner on the
  admission lane → `PersonaSigner` through the identity plane (the key
  materializing on first touch) → `StartTopic` + `PostTurn` → then a
  second, read-only realm client: `Board` returns the topic,
  `Materialise` returns 1 contribution on an active lifecycle, and the
  keyring built from `keys.public` yields **`sig=verified`** — the
  verdict earned, not defaulted [measured].
- **What this proves of Bar 1**: the compile guarantee and the
  live-read path, entirely on public tagged surfaces — the Go half
  needed **zero** upstream asks.
- **What remains**: plane-health and workload-lifecycle reads on the
  same rig; and spike B — the browser half (`nats.js` over WebSocket,
  Datastar-vs-web-components discrimination) — gated on upstream ask
  #1 (soulnode's WebSocket listener), with a stock operator-mode
  nats-server as the honest interim bench.

## 2026-08-13 — C4–C8 decided: browser-live over the realm's wire, plain words on screen

Four decisions from the operator in one sitting, plus a copy rule that
became C8.

**C8 — plain words on human surfaces** (operator, reviewing the Bar 4
screens): the component bynames — "the door", "the record", "the room" —
are "too geeky" as UI copy. Human-facing labels say what things do:
Storage, Connections, People & sign-in, Agents, Active topics, Latest
activity. Product names (soulstream, soulfold) belong in mono detail
rows; the bynames stay in internal docs and the journey. The Bar 4 rig's
copy was reworked the same hour and re-rendered (copy-only change — the
measurement stands). This rule is an upstream ask on the Design
project's voice guidance.

**C4 — a browser-live hypermedia UI**, not a compiled SPA (operator,
overriding the drafted server-rendered recommendation): the realm made
reachable over its own wire, with **Datastar or web components** as the
rendering layer — the
[datastar-skills](https://github.com/cbeauhilton/datastar-skills) repo
(skills: `datastar`, `nats-jetstream`) is the named resource. Recorded
honestly: the two candidates place the NATS connection differently —
Datastar's own tao is "backend is source of truth" over SSE (a
helm-held realm connection pushing DOM patches), while web components
pair with `nats.js` speaking NATS directly in the browser. The pick
between them is the rig's first discriminating question, not a
decision taken today. Bar 4's server-rendered finding stands as the
measured fallback either way.

**C5 — NATS over WebSocket** (operator): the realm reachable from the
browser over NATS's native WebSocket transport, `nats.js` only (the
standing rule: `nats.ws` is deprecated). **Named upstream ask #1**, per
Bar 1's rule: soulnode's embedded server exposes loopback TCP only
today — no WebSocket listener exists; the knob (and its founding-time
ceremony treatment) is soulnode's to grow. Custody note for Bar 2: the
browser holding the *user's own* token is the MCP-client trust class —
delegated to self, not helm custody; the bar's scan still applies to
the helm's own storage.

**C6 — observe + configure** stands as pre-registered; the participant
client remains a named successor topic. **C7 — fully parallel** with
`platform-tenancy-guardrails`: all configure surfaces may be built now,
and the rework risk where tenancy's decisions land differently is
accepted openly rather than hedged.

**Bar 1 shape note, recorded before the bar runs:** C4/C5 split the
prototype into a Go shell and a browser half. The compiler-enforced
consumer-position guarantee applies to the Go shell (module path
outside every component namespace, pinned tags, no `replace`); the
browser half speaks only public wire surfaces (NATS subjects, JetStream
API, the realm's published vocabulary). If measuring forces more than
this clarification, the amendment lands here with its reasons, openly.

## 2026-08-13 — Bar 4 measured: PASS — one source, two surfaces, zero external fetches

The bar ran the same day C3 was decided, on a rig in the session
scratchpad (`bar4/`: the shared source, two pages, vendored assets,
screenshots — scripts stay out of git per how-we-work).

**The rig.** One `shared.css`: the design system's seven token files
carried verbatim, with exactly one change — the Google Fonts CDN
`@import` replaced by vendored `@font-face` (Archivo variable woff2,
90,096 B, `wght` 400–800 *and* the `wdth` 62–125 axis; JetBrains Mono
400/500/700) — plus one shared component layer (transport bar, label
strips, LED pips, molded keys, fields, pills, segmented meters, tables,
the CRT screen) written to the system's own specs. Ten Lucide SVGs
vendored once (unpkg; one via jsDelivr after an unpkg 1102 error) and
inlined at build. Two pages consume it: the fold's sign-in re-rendered
with the *same markup structure* as the shipped `internal/ui` login
template, and the helm's first screen (four plane cards, a topics
table with channel-colored author pills, the tape-tail CRT).

**The measurement**, in a real Chromium against a loopback-only server:

- Requests: helm 5, fold sign-in 4 — every one to `127.0.0.1:8471`;
  `performance.getEntriesByType("resource")` filtered for external
  origins returns **empty on both** [measured]. The only console entry
  is the browser's automatic `/favicon.ico` probe against the serving
  host (404, no external fetch).
- Fonts real, not fallback: `document.fonts.status` = `loaded`;
  Archivo 400–800 and JetBrains Mono 400/700 loaded, the wordmark
  visibly condensed at `'wdth' 88` — the variable axis survived
  vendoring. JBM 500 declared-unused on these two pages, honestly
  noted [measured].
- No per-surface fork: both pages link the same `shared.css`;
  page-local `<style>` is layout-only and every value in it is a
  token.
- C3 exercised on screen: identity by label — "THE HELM" and "THE
  DOOR" strips; amber/teal held strictly to the channel semantic
  (author pills: daan amber, scribe/smith teal); no rose anywhere.

**A finding that feeds C4, not decided here:** the whole proof needed
no JS framework and no asset pipeline — the shared source is
consumable as a single CSS string plus static markup, exactly the
shape `webstyle` ships today (constitution III compatible). The design
system's JSX component layer was *not* consumed directly: the token
layer is shared verbatim, but the component layer was re-expressed as
CSS classes from the system's written specs. Server-rendered
consumption is proven cheap [measured]; direct JSX consumption remains
unproven and is C4's remaining question.

## 2026-08-13 — C1 and C2 decided: a sibling component, soulhelm — the helm

The operator decided the cockpit's home and its name.

**Home: a sibling component repo** — the sixth component, following the
consumer-position pattern the ecosystem has proven twice (the archivist,
soulfold): its own repo and constitution, a public embed seam
(`embed.Run(ctx, Options)`, the D29 pattern) that soulnode composes as a
plane by tag — the fourth URL logged at `up`, beside the door, the fold
sign-in, and the admin console — plus a standalone binary for
deployments that run the components without soulnode. The other two
candidates fell to constitutions, not taste:

- **Growing the fold's console** is blocked by the fold's own one-way
  door — "any Soulstream-only claim, endpoint, or side-channel is a
  constitution-II amendment"; a Soulstream-aware IdP collapses the
  ecosystem's two planes into one `[mechanism-argument]`.
- **Cockpit code native to soulnode** breaks "composition, not
  invention" (domain logic lands upstream — the article that has held
  through six features) and couples the human surface to the house: a
  realm running the components without soulnode would have no cockpit
  at all `[mechanism-argument]`.

The sibling shape also collapses a happy coincidence: Bar 1's
consumer-position rig (module path outside every component's namespace)
*is* the sibling repo's natural e2e gate — the bar's proof becomes the
component's standing test `[judgment]` on the final shape.

**Name: `soulhelm` — the helm** `[judgment]`: the place where the
vessel is steered and watched, which is what a cockpit is; it keeps the
one-noun family (the record, the room, the name, the house, the fold,
the helm). `soulbridge` was rejected for a real collision — soulstream's
*Later* section already names Slack/email protocol bridges — and
`soulconsole` because "console" already names the fold's admin surface.

Follow-through, named not done: the `soulhelm` component tag enters the
hq (hqlint's component list, the README constellation, the journey
conventions) when this topic graduates and the repo is founded — not
before. The topic stays `ecosystem`-tagged.

## 2026-08-13 — scoping: what already exists

The topic opens with more prior art than expected, found before any
experiment ran:

- **The design half is real.** The **Soulsystem Design System** exists as a
  Claude Design project (updated 2026-08-13, same day this topic opened):
  seven token files, seventeen components with props contracts and usage
  prompts, eighteen guideline specimens, a voice-and-tone canon, and two UI
  kits — a marketing site and a **Stream Console** (sign-in, transport-bar
  chrome, a three-column live stream with human/machine turns and a CRT
  tape readout, a reel archive, settings). The cockpit's first screens have
  in effect been sketched. Its brand rule of record: amber = human channel,
  teal = machine channel, **at deliberately equal weight** — the vision's
  "humans and AI as peers" as a color system.
- **The canon conflict is real too.** The design system is cassette
  futurism in a *light* key (warm shell tones, one dark CRT surface);
  soulfold ships a *dark* language (`internal/webstyle`, violet accent,
  rose edge, episode
  [0063](../../04-JOURNEY/0063-soulfold-the-console-gets-a-face.md))
  descended from impire.io. The design project's own readme records it was
  authored from a one-line brief with no codebase supplied — the two have
  never seen each other. C3 exists because of this.
- **Two declared substitutions in the design system** matter to Bar 4:
  fonts (Archivo / JetBrains Mono standing in) and icons (Lucide loaded
  from a CDN). Self-containment means vendoring both.
- **The entry asymmetry** motivating the topic: the MCP door gives machines
  the full tool surface; humans get the fold's identity console only.
  Soulnode's front of house (episode
  [0062](../../04-JOURNEY/0062-soulnode-the-front-of-house.md)) already
  logs door, sign-in, and console URLs — the cockpit would be the fourth
  URL on that line.
- **Sequencing note (C7):** the configure surface will render exactly what
  [`platform-tenancy-guardrails`](../platform-tenancy-guardrails/README.md)
  is deciding (accounts, grants, guardrails). Observe-first work is
  independent of it; the configure half likely trails it.
