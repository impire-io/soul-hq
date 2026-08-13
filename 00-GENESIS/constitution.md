<!--
Sync Impact Report
==================
Version change: (five ancestors) → 1.0.0 (MAJOR — the five per-project
constitutions merge into one ecosystem constitution)
Ancestry (final versions, frozen in ../99-ARCHIVE/genesis/):
  - soulstream   1.1.0  (ratified 2026-07-12, last amended 2026-07-24)
  - soulstream-workloads    0.1.1  (drafted 2026-07-22, never formally ratified)
  - soulstream-identity 1.3.1  (ratified 2026-07-28, last amended 2026-07-29)
  - soulstream     1.0.0  (ratified 2026-08-02)
  - soulstream-idp     1.0.0  (ratified 2026-08-02)
Structure:
  - Shared articles S1–S5 factor out the principles all five ancestors carried
    (NATS-native, smallest-viable, docs-first, research gates, all-green gate).
  - Component articles keep their ORIGINAL numbering (e.g. "soulstream-workloads
    constitution I" still names the substrate boundary), so every existing
    citation in design docs, journey episodes, and frozen specs resolves.
  - The Working Agreement is carried verbatim in substance; it was already
    identical across the five.
Relocation:
  - Canonical copy lives at soul-hq/00-GENESIS/constitution.md. Each component
    repo's .specify/memory/constitution.md is a relative symlink to it
    (../../../soul-hq/00-GENESIS/constitution.md), so every spec-kit plan's
    Constitution Check reads these articles.
Follow-up TODOs: none
-->

# The Soul Constitution

One constitution for the Soulstream ecosystem — **soulstream** (the record),
**soulstream-workloads** (the room), **soulstream-identity** (the name), **soulstream-idp** (the
fold), and **soulstream** (the house). The canonical copy lives at
`soul-hq/00-GENESIS/constitution.md`; each component repo's
`.specify/memory/constitution.md` is a symlink to it, so every spec-kit
plan's Constitution Check reads these articles. Decisions are held against
this file and [`vision.md`](vision.md) — see the decision test in
[`README.md`](README.md).

Shared articles (S1–S5) bind every component. Component articles bind their
component and keep the numbering they had in that component's own
constitution, so existing citations ("constitution II", "constitution VI")
keep resolving. On conflict, a component article is the more specific rule
and wins for its component.

## Part I — Shared Principles

### S1. NATS-Native First

NATS is the platform, not a dependency. A working deployment is a NATS server
with JetStream, credentials, and the protocol — nothing else.

- Every capability MUST be implemented with built-in NATS and JetStream
  primitives (streams, consumers, KV, Object Store, subject hierarchies,
  headers) before any custom mechanism is considered. Designs MUST evaluate
  current NATS server features — including atomic batch publishing, batched
  direct get, message scheduling, per-message TTLs, and optimistic concurrency
  via `Nats-Expected-Last-Subject-Sequence` — before proposing custom code.
- Infrastructure that duplicates a NATS capability — databases, coordinators,
  API tiers, external queues — is prohibited. If NATS can express it, NATS
  MUST express it.
- Coordination is deterministic rules, idempotent operations, and optimistic
  concurrency. Elections and consensus rounds are banned.
- Where a component's front door is legitimately not NATS (soulstream-idp's HTTP
  door — WebAuthn is origin-bound), the store and the deployment story stay
  NATS-native and the exception is named in that component's articles.

**Rationale**: the ecosystem's premise is that the "what is needed" list stays
short. Every custom component beside NATS is a component that can fail, drift,
or require operation independently of the stream that is the system of record.

### S2. Smallest Viable Implementation

- Every feature MUST be the smallest implementation that satisfies its
  specification. Anything not required by an acceptance scenario or a concrete
  consumer is cut or deferred.
- Speculative generality is prohibited: no configuration options, abstraction
  layers, or plugin points added "for later". A seam exists only where the
  design names its concrete future occupants.
- Growth is expressed as new vocabulary over the existing log, new backends on
  existing seams, or new modes on existing surfaces — never new machinery
  beside them.
- Scope creep is a review blocker, not a style concern. Reviewers MUST reject
  additions that exceed the spec, however well-built.

**Rationale**: the original idea was once buried under its own elaborations.
Keeping each change minimal is how each component stays answerable to "what is
needed" — and how the security-adjacent components stay small enough to audit.

### S3. Documentation Is a First-Class Citizen

- Every concept is explained simply — plain words, one concept per page, an
  everyday analogy before any technical detail. A concept that cannot be
  explained simply is a signal the concept itself is too complicated.
- No feature is complete until its concepts are documented. Docs ship in the
  same change as the behavior they describe; stale documentation is a bug with
  the same severity as a failing test.
- Plain words beat invented terms (the new-term test: if the plain word works,
  use it). Invented vocabulary is a budget, spent only on concepts a plain
  word cannot carry.
- Load-bearing design decisions are recorded where the design lives — numbered
  decisions (D-numbers) or numbered design documents — with their reasoning,
  so future changes argue against the real reasons.

### S4. Research Gates Before Build Spends

Every build milestone names the research gate it depends on (a decided
substrate, a proven contract, a measured constraint, a public surface that
exists), and no machinery is built ahead of its gate. Speculation is research,
recorded in `01-RESEARCH/`; it is not code. Exit criteria are written before
the work and amended only openly, with the raw findings recorded.

### S5. All-Green Quality Gate

Done means the full gate is green with nothing skipped: each component repo's
standard gate (`make fmt && make test && make lint`, or `make check` where
that is the repo's named gate), and soul-hq's own gate (which includes the hq
structural lint, `internal/hqlint`). Tests that need no NATS server run
without one; anything touching NATS uses an in-process server or fake
transport so suites have no external dependency. Sign every commit. Never
commit `.claude/settings.local.json`. Hook or gate failures are blocking —
fixed before anything else continues.

## Part II — The Working Agreement (Anti-Drift)

Adopted 2026-07-24 with soulstream's hq adoption and carried by every
component since; identical across all five ancestors. It guards a specific
failure mode: a fluent counterpart steering the maintainer on a load-bearing
call he cannot independently check in the moment, without either party
intending it. Applies to every load-bearing decision — a protocol or boundary
change, a scope call, a criterion, or a public claim.

1. **Teach-back as a gate.** No load-bearing direction decision is recorded
   until the maintainer can restate the argument for it in his own words. If
   he can't, the decision isn't ready — the deficit is in the explanation, not
   the listener.
2. **Claims carry their evidence class.** Every load-bearing claim is tagged
   **[measured]** (a reading in the repo: a test, a demonstrated behavior, a
   benchmark), **[mechanism-argument]** (a reasoned case from how the system
   works, attackable by reasoning), or **[judgment]** (taste or preference).
   Only measured closes a debate.
3. **Decisions record the reversal condition.** Every direction decision gets
   a "what would change our minds" line written *when the decision is made*
   (the journey episode template requires it), phrased as an observable
   reading, so a future reversal is a clean, anticipated turn instead of
   drift.
4. **Adversarial pass on direction changes.** For decisions that change a
   protocol's shape or a core boundary, the other side is argued at full
   strength before the decision — the maintainer never sees only the most
   convincing case.

## Part III — Component Articles

### Soulstream — the record

Soulstream's three original articles are the ancestors of the shared
principles and live there now:

- **I. NATS-Native First** → shared article S1.
- **II. Smallest Viable Implementation** → shared article S2. Soulstream's
  specific test stands: if a design addition does not survive the "what is
  needed for a working soulstream" list staying short, it goes in the design's
  `extensions/` or it goes nowhere.
- **III. Documentation Is a First-Class Citizen (ELI5)** → shared article S3.
  Soulstream's specific duty stands: the `docs/` folder explains every concept
  simply enough for a five-year-old, and every user story's task list includes
  its `docs/` task.

### soulstream-workloads — the room

- **I. The Substrate Boundary (NON-NEGOTIABLE).** soulstream-workloads is a runtime,
  never a store of record. The authoritative home of any artefact — its bytes,
  its history, its current state — is the soulstream topic. soulstream-workloads
  launches, supervises, observes, and retires workloads; everything worth
  keeping flows back into the topic as ops. A workload that dies loses scratch
  state, never history. No feature may make soulstream-workloads the place a piece of
  durable truth lives. This article does not relax for convenience.
- **II. One Identity, No Privileged Tier.** Every workload runs as a persona
  with scoped NATS credentials — the same kind of identity a human persona
  holds. No bot API, no elevated standing for machine personas; credentials
  are always scoped to what a workload needs and no more. Behaviour may never
  branch on whether a persona is human or machine.
- **III. Contracts Orthogonal to Backends.** A workload contract is defined
  independently of the isolation backend that runs it (native process,
  Docker/OCI, microVM, Kubernetes). A contract must never leak which backend
  is in use, and a backend must be swappable per node without touching a
  single workload declaration. Where the two axes meet — resource limits,
  credential injection, lifecycle signals — the seam is explicit and
  documented. If a contract can only be satisfied by one backend, the design
  is wrong.
- **IV. Research Gates Before Build Spends** → shared article S4.
- **V. Execution Is Observable and Attributable.** A workload's lifecycle —
  start, progress, result, exit — is visible in the realm and attributable to
  a persona, or the work is not done. Execution is a stream of operations
  anyone in the topic can follow and replay. If a runtime backend cannot
  surface lifecycle as ops, that gap is a named limitation, written down,
  never a silent hole.
- **VI. All-Green Quality Gate** → shared article S5.

### soulstream-identity — the name

- **I. Custody Without Possession.** Secrets live in the vault and answer
  requests; they are never handed out. No API surface may return a seed,
  private key, or any material sufficient to sign without the vault. Every
  exception is a named custody escape (today: credential export), always
  explicit in the request, always loud in the audit log, never a side effect.
  Every signing operation is attributable and logged. In-process access to key
  material stays inside the vault package; the process boundary is the custody
  boundary.
- **II. The Server Is the Verifier of Record.** The NATS server enforces;
  soulstream-identity decides only what is genuinely its own. Transport permissions
  live NATS-side — scoped signing keys or callout-issued JWTs — enforced by
  the server on every connection, including the op tail of the subject (D25).
  soulstream-identity's own policy surface is exactly the declared bindings, token
  records, and validated claims; who exists is the IAM's truth. Validations
  the server will repeat are diagnostics, never gates.
- **III. Smallest Viable Implementation** → shared article S2.
- **IV. Documentation Is a First-Class Citizen** → shared article S3.

### soulstream — the house

- **I. Composition, Not Invention (NON-NEGOTIABLE).** soulstream contains no
  domain logic. Identity behavior lives in soulstream-identity, runtime behavior in
  soulstream-workloads, record behavior in soulstream; soulstream wires their public
  surfaces together and adds only what composition itself requires. Components
  are consumed as tagged releases through public packages — never `internal/`
  paths, never `replace` directives on main — and a soulstream release names the
  component versions it bundles. If a feature cannot be built without new
  domain behavior, that behavior lands upstream first. This article does not
  relax for convenience.
- **II. Same Shape as Any Deployment.** The embedded NATS server runs operator
  mode with auth-callout admission, exactly as a hosted deployment does. No
  dev-only auth lane, no local-only bypass, no admission shortcut the rest of
  the ecosystem lacks. Divergence between the two shapes is a bug, never a
  feature.
- **III. One Process, Planes by Configuration.** Enabled planes run in one
  process, each on an ordinary loopback NATS connection; repointing or
  disabling a plane is configuration, never a different build. Workloads
  always run outside the process, through soulstream-workloads's isolation backends; a
  workload failure never takes the node down, and a node-plane failure is
  surfaced and named, never silent.
- **IV. Research Gates Before Build Spends** → shared article S4.
- **V. First Boot Is the Product.** `soulstream init` performs the entire
  ceremony with zero manual key steps, and `soulstream up` reaches a connectable
  realm on a fresh machine in minutes. The distance from download to working
  realm is a measured, guarded number. First-boot regressions are release
  blockers; a manual step added to the ceremony is a constitution violation,
  not a documentation task.
- **VI. All-Green Quality Gate** → shared article S5.

### soulstream-idp — the fold

- **I. Passkeys, Not Passwords.** The fold authenticates users with WebAuthn
  ceremonies and nothing else. No password lane exists, ever — not as a
  fallback, not for bootstrap, not behind a flag. Nothing the fold stores may
  be sufficient to impersonate a user; OAuth client secrets are stored as
  digests only. Recovery is re-enrollment through a deliberate, logged act —
  never a knowledge factor and never an unattended email loop.
- **II. Indistinguishable by Design.** Consumers reach the fold only through
  the OIDC spec surfaces — discovery, JWKS, the authorization endpoint, the
  token endpoint. No Soulstream-only claim, header, endpoint, or side-channel
  exists; soulstream-identity's callout issuer MUST be unable to tell the fold from
  Entra (soulstream-identity's D23 seam is the contract). Group names surface as
  roles-claim values that *name* declared roles; they carry no permissions,
  and the NATS server remains the verifier of record. Signing keys rotate
  through JWKS the way the spec says.
- **III. Smallest Viable Implementation** → shared article S2. soulstream-idp's
  specific rule stands: protocol comes from certified, maintained libraries —
  the OP core and the WebAuthn ceremonies are consumed, never hand-rolled.
- **IV. Documentation Is a First-Class Citizen** → shared article S3.

## Technology Constraints

Shared:

- **Language**: Go, across the ecosystem. Official, maintained NATS client
  libraries only (`nats.go`, `nkeys`, `jwt/v2`, embedded `nats-server` for
  tests); deprecated clients (e.g. `nats.ws`) MUST NOT be used. Connect via
  `github.com/synadia-io/orbit.go/natscontext`; modern `nats.go/jetstream`
  API.
- **NATS server**: target a modern release (2.12+ as of ratification). Any
  feature relying on a specific server capability MUST state its minimum
  server version in its plan.
- **Persistence**: JetStream only (streams, KV, Object Store). No external
  databases.
- **Dependencies**: identity-adjacent components are judged by their audit
  surface — keep the dependency tree small enough to read.

Component-specific:

- **soulstream-identity**: secrets at rest ride NATS KV with xkey envelope
  encryption — only ciphertext is ever stored; NATS is the only transport
  (xkey-sealed request/reply; there is no socket and no TCP listener). The
  pre-connection moment belongs to the connection ladder: creds-file bypass or
  auth callout.
- **soulstream-idp**: the OP core from `zitadel/oidc`, the ceremonies from
  `go-webauthn/webauthn`; hand-rolled protocol or cryptography is a review
  blocker. The front door is HTTP(S), named honestly; JetStream KV is the
  store; the serve assembly is embeddable (the ecosystem's embed pattern,
  soulstream-identity D29).

## Development Workflow & Quality Gates

- The working structure lives in this repository (soul-hq) as described in
  [`how-we-work.md`](how-we-work.md): research (`01-RESEARCH/`) → design
  (`02-DESIGN/<component>/`) → implementation (the component repos' spec-kit
  flow, sequenced by `03-IMPLEMENTATION/ROADMAP.md`) → journey
  (`04-JOURNEY/`, one numbered episode per landed feature, concluded research
  topic, or load-bearing decision).
- Research never goes through spec-kit; implementation in a component repo
  always does (or follows the design's D-numbers where that is the repo's
  named flow). Every plan MUST pass the Constitution Check gate; violations
  are either removed or justified in Complexity Tracking.
- **Landing work means, in the same working session:** the component repo's
  gate green and its feature merged, and the soul-hq commit carrying the
  journey episode, the roadmap update, and design propagation. Neither half
  is optional.
- The hq structural lint (`internal/hqlint` in soul-hq) enforces this
  repository's invariants under `make test`.
- Commits are signed, in every repository.

## Governance

- This constitution supersedes all other practices for every component. On
  conflict with a README, CLAUDE.md, or any template, the constitution wins.
  On conflict between a shared and a component article, the component article
  wins for its component.
- **Amendments**: made by editing this file, and MUST include an updated Sync
  Impact Report, a version bump, and a journey episode recording the why and
  the reversal condition.
- **Versioning policy** (semantic): MAJOR — removing or redefining an
  article; MINOR — a new article or section, or materially expanded guidance;
  PATCH — clarifications, wording, and non-semantic refinements.
- **Compliance review**: every spec-kit plan's Constitution Check enforces the
  shared articles and the component's own; every review verifies the change is
  NATS-native, minimal, and documented. Complexity MUST be justified or
  removed.

**Version**: 1.0.0 | **Ratified**: 2026-08-02 | **Last Amended**: 2026-08-02

*Amendment history:*
- *1.0.0 (2026-08-02)* — the five per-project constitutions merge into one
  ecosystem constitution with shared articles S1–S5 and per-component
  articles keeping their original numbering. Ancestors, frozen with their
  full texts in [`../99-ARCHIVE/genesis/`](../99-ARCHIVE/genesis/):
  soulstream 1.1.0, soulstream-workloads 0.1.1 (draft), soulstream-identity 1.3.1, soulstream
  1.0.0, soulstream-idp 1.0.0. Recorded in the consolidation episode in
  `../04-JOURNEY/`.
