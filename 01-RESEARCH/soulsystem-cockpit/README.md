# Can the soulsystem grow a cockpit — one human surface to observe and configure the whole system — as a pure consumer?

**Component:** ecosystem
**State:** active
**Started:** 2026-08-13

## Abstract

The system's two entries are asymmetric. Machines get the MCP door — a URL
into the realm with the full tool surface. Humans get a passkey sign-in and
the fold's admin console, which manages *identity*, not the system: no view
of topics, personas, workloads, planes, or memory, and no configuration
surface beyond the fold's own lifecycle. The proposal is a **cockpit**: the
human entry into the soulsystem alongside the MCP server, where the running
system is observed and configured. The design half already exists — the
**Soulsystem Design System** (a Claude Design project, updated 2026-08-13)
carries tokens, seventeen components, and a *Stream Console* UI kit whose
equal-weight amber/teal channel rule is the vision's "humans and AI as
equals" made into a color system — but it diverges from the dark
violet/rose language soulfold ships today (`internal/webstyle`, episode
[0063](../../04-JOURNEY/0063-soulfold-the-console-gets-a-face.md)). A
decisive answer unlocks a design doc and the build of the ecosystem's human
front; a negative one confines the ambition to the existing surfaces before
any code is written.

## The question

**Where and as what does the human cockpit live, such that it observes and
configures the whole soulsystem while remaining a pure consumer of public
tagged surfaces — custodying nothing, creating no second control plane, and
wearing one design system shared with the fold?**

Sub-questions that prove to need their own investigation become successor
topics rather than growing this one (the one-question rule); the candidates
are named under *Decisions that are not bars* below — the participant
client (humans posting turns from the browser) is already one of them.

## Pre-registered bars

Written before any experiment. These are the claims this topic can settle by
**measurement**; the design decisions it must also take are listed
separately below, honestly, rather than dressed as bars.

- **Bar 1 — pure consumer, compiler-proven.** A cockpit prototype in
  consumer position — a module whose path sits outside every component's
  namespace, the `e2e/embedgate` shape, so an `internal/` import cannot
  compile — renders live realm state (topics, personas, workload lifecycle,
  plane health) against **pinned tagged upstream releases only**. **Pass:**
  the rig compiles and runs with zero `replace` directives; every
  capability it needs that is not public is recorded as a named upstream
  ask in `JOURNEY.md`, never dodged. A single namespace dodge or `replace`
  fails this bar.
- **Bar 2 — the cockpit custodies nothing.** After one complete human
  session against a live soulnode (sign in via the fold, observe, perform
  one configuration change, sign out), a scan of the cockpit's own storage
  — disk, KV, anything it can write — finds nothing credential-shaped and
  no store of record, with a positive control proving the scan can find a
  planted secret. The cockpit acts through admission as the signed-in
  principal, and the change performed is attributed to that principal —
  delegated authority, never borrowed identity (S6, episode
  [0064](../../04-JOURNEY/0064-ecosystem-the-platform-turn.md)). **Pass:**
  clean scan, positive control caught, attribution readable in the record
  or audit.
- **Bar 3 — configuration without a second control plane.** Before the
  spike runs, every mutation cockpit v1 offers is enumerated in a table in
  `JOURNEY.md`, each mapped to an existing class of surface: (a) an op on
  the record, (b) an admin op on an existing plane surface, (c) a
  config-file change with its restart/reload semantics stated. The spike
  then demonstrates **one mutation of each class** end-to-end from the
  browser. **Pass:** the table is complete, all three demonstrations green,
  and no mutation required a new privileged side-channel, a new subject
  namespace with elevated permissions, or a cockpit-owned store. A single
  unmappable mutation fails this bar.
- **Bar 4 — one design system, two consumers.** From a single
  token/component source derived from the Soulsystem Design System, render
  the cockpit's first screen **and** one existing fold page. Both must be
  fully self-contained: fonts and icons vendored, zero CDN or third-party
  fetches — measured by loading both surfaces with network access denied
  beyond the serving host. **Pass:** both render from the same source with
  no per-surface fork of the tokens; the canon decision (C3 below) is
  recorded with its reasoning before this bar runs, since it decides what
  "the same source" contains.

## Decisions that are not bars

Load-bearing calls this topic must take that no experiment can settle —
they resolve by design argument and will carry `[mechanism-argument]` or
`[judgment]`, never `[measured]`. Recorded here so the topic is honest
about which of its outputs are evidence and which are choices.

| Marker | Decision | Character |
|---|---|---|
| C1 | ~~Where the cockpit lives~~ — **decided 2026-08-13** (see `JOURNEY.md`): a sibling component repo with a public embed seam, composed by soulnode as a plane by tag (the soulfold/archivist pattern); the fold-console and soulnode-native candidates fell to their own constitutions | structural |
| C2 | ~~Its name~~ — **decided 2026-08-13** (see `JOURNEY.md`): **`soulhelm` — the helm**; `soulbridge` rejected (collides with soulstream's planned protocol bridges), `soulconsole` rejected (names the fold's surface) | cosmetic |
| C3 | ~~The design-system canon~~ — **decided 2026-08-13** (see `JOURNEY.md`): cassette-futurism light is canon; per-component identity rides labels (mono strips, wordmarks), never a second color system; the violet/rose dark language retires | direction |
| C4 | ~~Serving architecture~~ — **decided 2026-08-13** (operator, see `JOURNEY.md`): a browser-live hypermedia UI — Datastar or web components — over the realm's own wire, not a compiled SPA; the Datastar-vs-web-components pick is the rig's first discriminating question (they place the NATS connection differently) | structural |
| C5 | ~~Read transport~~ — **decided 2026-08-13** (operator): the realm reachable over **NATS WebSocket** (`nats.js`, never the deprecated `nats.ws`); named upstream ask: soulnode's embedded server has no WebSocket listener today | structural |
| C6 | ~~The v1 scope boundary~~ — **decided 2026-08-13**: observe + configure; the participant client (posting turns) is a named successor topic | direction |
| C7 | ~~Sequencing~~ — **decided 2026-08-13**: fully parallel with [`platform-tenancy-guardrails`](../platform-tenancy-guardrails/README.md); the rework risk on tenancy-shaped configure surfaces is accepted openly | sequencing |
| C8 | Human-facing copy: plain functional labels (Storage, Connections, People & sign-in, Agents), never the component bynames — **decided 2026-08-13** (operator: the bynames are "too geeky" on screens; they stay in internal docs) | content |

## Reversal condition

Written now: **if the cockpit cannot be built as a pure consumer** —
observable as a required `internal/` import from any component, a required
new privileged surface upstream that exists only for the cockpit, or a
required cockpit-owned store of record — then the cockpit does not belong
in the soulsystem as a component, this topic ends `abandoned`, and the
human path stays with the existing surfaces (the fold's console, the MCP
door, the CLI).

A second, narrower reading: if Bar 3's table shows that cockpit v1's real
mutations fall (almost) entirely into class (c) — config-file edits with
restarts — then the "cockpit" is a settings page wearing a grand name, and
the scope collapses to match the evidence rather than the ambition.

## Verdict

Graduated to design, 2026-08-13 — all four bars PASS, all eight
decisions taken, same day the topic opened.

- **Bar 1 — pure consumer, compiler-proven: PASS** [measured]. Module
  `soulhelm.invalid/rig` (outside every component namespace) pins
  soulnode v0.3.1, soulstream v0.7.0, soulidentity v0.1.0 by tag —
  soulfold v0.3.3, soulrealm v0.1.0, the archivist v0.2.0 arriving
  transitively — **zero `replace` directives**, `go vet` clean. One
  run, ~1.5 s wall: boot a whole realm in-process, found it, post a
  signed topic, read back board (1 topic), a turn at `sig=verified`
  (keyring earned from `keys.public`), a claimed work item, 5
  archivist answers each with a citation, the door answering HTTP.
  Browser half on the decided path: the Datastar dashboard rendered
  live realm state in a real Chromium, across an in-place node
  restart. Zero namespace dodges; two upstream asks recorded, neither
  needed for the Go half.
- **Bar 2 — custodies nothing: PASS** [measured]. Full ceremony:
  passkey enrolment on the bundled fold (virtual authenticator, real
  WebAuthn), helm sign-in via RFC 7591 DCR + code+PKCE, the session's
  own NATS admission (audit: `callout ADMITTED lane=oidc
  subject=u-90e41123cd67bc45 role=realm display=owner`), `work.open`
  as that principal with `signed=true` (the persona signer
  materialized for the fold principal — S6 held), sign-out. Scan of
  the helm's storage after the session: **zero** credential-shaped
  hits; the positive control fired (after the first, non-hex plant
  was corrected — a control that cannot fire proves nothing); session
  state memory-only.
- **Bar 3 — configuration without a second control plane: PASS**
  [measured]. The mutation table (8 rows, 3 classes) was committed
  before the spike ran. Demonstrated from browser buttons: (a)
  `work.open` on the record, live on the dashboard within a second;
  (b) `tokens.create` through the identity plane — the owner lane
  refused with an explicit server `Permissions Violation` on the op
  tail [measured], minted via the node ops lane, `sit_` returned
  once, issue audited; (c) `planes.memory.enabled=false` + in-place
  restart, clients re-admitted, memory answers **5 → 0**. No mutation
  required a new privileged side-channel or a helm-owned store. The
  narrower reversal reading did not fire: mutations spread across all
  three classes, not config-file edits alone.
- **Bar 4 — one design system, two consumers: PASS** [measured]. One
  `shared.css` (the seven token files verbatim, CDN font import
  replaced by vendored `@font-face`, the `wdth` axis surviving)
  renders the helm's first screen and the fold's sign-in page: helm 5
  requests, fold 4, every one to the serving host, external resource
  list empty, fonts genuinely loaded (`document.fonts`), no
  per-surface token fork. C3 was recorded before the bar ran, as the
  bar required.

Decisions C1–C8 are in the table above with their reasoning in
`JOURNEY.md`. Standing upstream asks: #1 a WebSocket listener in
soulnode's embedded server (one options field — plumbing, not
invention); #2 an owner-reachable token-management lane (soulidentity
/ the tenancy topic's grant work). Findings of record: asking memory
is participation, not observation [measured]; attribution rides
soulfold's persona-shaped ids end-to-end [measured]; the Datastar v1
protocol lessons (release bundle, `data-init`, `data-on:click`,
one-shot SSE races) are banked for the build.
