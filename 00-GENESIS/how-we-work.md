# How we work

The process companion to [`constitution.md`](constitution.md): the pipeline,
the lifecycles, the duties, and how all of it is enforced.
[`../README.md`](../README.md) holds the one-screen map.

**The two-repo reality.** This repository (soul-hq) is the headquarters for
all five components: research, designs, the roadmap, and the journey live
here. Code, frozen spec-kit artifacts (`specs/NNN-*/`), and plain-words docs
live in each component's own repository (a sibling checkout:
`../soulstream`, `../soulrealm`, `../soulidentity`, `../soulnode`,
`../soulfold`). Every artifact in this repo names its component; ecosystem-
level work uses the component tag `ecosystem`.

## The pipeline

```
question ──/research-start──▶ 01-RESEARCH/<slug>/     (state: active)
                                   │
                     /research-graduate <slug>
                        │            │           │
                     design       artifact    abandoned
                        │            │           │
                        ▼            │           │
        02-DESIGN/<component>/ doc   │           │
                        │            ▼           ▼
   /speckit-specify (in the      04-JOURNEY episode (always; folder removed)
    component repo)
                        │
                        ▼
 <component>/specs/NNN-*/ + code ──landed──▶ /journey-log episode
                        │                        + ROADMAP.md updated
                        ▼
        design docs updated (behavioral changes propagate back)
```

Two hard boundaries:

- **Research never goes through spec-kit.** Spec-kit assumes you know what
  you're building; research exists to find out *whether* to build. Research
  uses the pre-registration method below, in `01-RESEARCH/`.
- **Implementation always follows a design.** For the spec-kit components
  (soulstream, soulrealm, soulnode) a design doc in `02-DESIGN/<component>/`
  is written to be the argument to `/speckit-specify`, run in the component
  repo; the generated plan's Constitution Check reads GENESIS through that
  repo's `.specify/memory/constitution.md` symlink (which points here). For
  the D-number components (soulidentity, soulfold) a roadmap milestone points
  at the design doc and D-numbers it realizes; a capability that isn't
  decided yet starts as research, not as code.

## Research (`01-RESEARCH/`)

One folder per open topic, any component, created with
`/research-start <slug>`. The folder's `README.md` (from
[`../01-RESEARCH/TEMPLATE.md`](../01-RESEARCH/TEMPLATE.md)) carries: Title,
Component, State (`active | graduated | abandoned`), Abstract, the Question,
and **pre-registered bars** — the pass/fail criteria written *before* any
experiment runs. The folder's `JOURNEY.md` records the investigation as it
happens.

- **Method:** hypothesis → cheap discriminating experiment → verdict, one
  variable at a time. What a "discriminating experiment" is varies by
  component — a demonstrated NATS behavior (soulstream), a spike running a
  real workload (soulrealm), a consumer-position rig wiring real component
  releases (soulnode), a protocol rig against a stock client (soulfold,
  soulidentity). Experiment scripts live in the session scratchpad;
  conclusions, documents, and principled code changes land in git.
- **Always committed and pushed** — even work that will be abandoned. The
  point is a permanent trail; abandoned research keeps its full history in
  git after the folder is gone.
- **Ending:** `/research-graduate <slug> --to design|artifact|abandoned`
  composes the topic's journey into the next-numbered `04-JOURNEY/` episode
  (verdict, evidence tags, reversal condition included), creates or updates
  the component's design doc when the outcome is a design, and removes the
  topic folder in every case. An abandoned topic is a *result*, recorded with
  the same care as a success.

## Design (`02-DESIGN/<component>/`)

The normative design, one folder per component; each folder keeps its own
conventions and its own README index:

- **soulstream/** — `core/` (the protocol; a realm running only this is a
  working soulstream) + `extensions/` (optional conventions).
- **soulrealm/**, **soulnode/** — numbered documents (`0001-…` onward).
- **soulidentity/**, **soulfold/** — named documents carrying global
  D-numbered decisions.

Documents are written functional-level — explicit enough that
`/speckit-specify` can turn one into a spec without guessing: the capability,
its seams, its configuration surface, its acceptance criteria. Every
behavioral change made during implementation propagates back into the design
docs it touches — the docs describe the system as it *is*. A new capability
that isn't yet decided starts as research, not as a design doc.

## Implementation (`03-IMPLEMENTATION/` + the component repos)

[`ROADMAP.md`](../03-IMPLEMENTATION/ROADMAP.md) is the live plan for the
whole ecosystem: a holistic "where we are" and one section per component —
what gets built, in what order, behind which research gate. No dates; gates,
not calendars. Features run in the component repo (spec-kit flow on a
numbered feature branch where that is the repo's way; `specs/NNN-*/`
artifacts freeze when the feature lands). **Landing a feature means, in the
same working session:** the component repo's quality gate green and the
feature merged, and the soul-hq commit carrying the roadmap update, the
journey episode, and design propagation. Frozen `specs/` bodies are
point-in-time artifacts and are not rewritten after the fact; their `Status`
field records shipping. The roadmap file itself is load-bearing: changes to
it are decisions and belong in the journey as episodes like any other.

## Journey (`04-JOURNEY/`)

The append-only log for the whole ecosystem: one numbered episode
(`NNNN-<component>-<slug>.md`) per landed feature, concluded research topic,
or load-bearing decision — one chronological sequence across all components,
written with `/journey-log` (or `/research-graduate`, which writes it for
research). The [`TEMPLATE.md`](../04-JOURNEY/TEMPLATE.md) requires: what
happened with the honest numbers, what was refuted or reversed,
evidence-class tags on load-bearing claims, and a **Reversal condition**
line. `README.md` carries the preamble, the per-component "Where things
stand" summaries, the episode index, and the pre-merge numbering map — all
refreshed with every episode. Episodes 0001–0049 were renumbered from the
five per-project journals when the hqs merged; the map in the README resolves
every old reference.

## The working agreement (anti-drift)

The four correctives are constitution articles (Part II there); this is how
they run day to day:

- **When to teach-back:** any decision that changes a protocol's shape, a
  boundary, a scope, a criterion, or a public claim. The assistant asks for
  the restatement; the decision is recorded only after it survives.
- **Tagging:** write `[measured]` / `[mechanism-argument]` / `[judgment]`
  inline where the claim is made — in conversation, in episodes, in design
  docs. If a debate is being closed by anything other than `[measured]`, stop
  and say so.
- **Reversal conditions:** phrased as observable evidence, not vibes. Written
  at decision time, never retrofitted.
- **Adversarial pass:** for calls that change a protocol's shape or a core
  boundary, the assistant argues the other side at full strength *before* the
  decision.

## Enforcement (how this stays true without willpower)

1. **The constitution symlink.** Each component repo's
   `.specify/memory/constitution.md` →
   `../../../soul-hq/00-GENESIS/constitution.md`, so every spec-kit plan is
   checked against GENESIS mechanically.
2. **The structural lint.** soul-hq's `internal/hqlint` rides `make test`
   here: layout, research-state legality, episode numbering/component tags
   and required fields, index completeness, and that relative links inside
   the hq resolve.
3. **The skills.** `/research-start`, `/research-graduate`, `/journey-log`
   make the transitions one command each — available in soul-hq for holistic
   sessions and in each component repo (retargeted to write here) so the
   journey duty happens at feature-landing time. They stage explicit paths,
   commit signed, and never push — pushing stays a human act.
4. **Orientation.** Every repo's `CLAUDE.md` and `AGENTS.md` point every
   session here first.

## Quality gates (the non-negotiables, in one place)

- Component repos: `make fmt && make test && make lint` (or `make check`
  where that is the repo's named gate) — all green, nothing skipped, before
  any "done" and before every commit.
- soul-hq: `make fmt && make test && make lint` — includes the hq structural
  lint (`internal/hqlint`).
- Keep pure logic separate from NATS I/O so it unit-tests with no server.
- Sign every commit. Never commit `.claude/settings.local.json`.
- NATS-native first (S1) and smallest viable implementation (S2) apply to
  every change, product or research.
