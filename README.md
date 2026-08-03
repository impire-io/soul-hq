# soul-hq — the soulsystem's headquarters

Everything about *how the ecosystem is run* lives here, for all five
components at once. The five are **the soulsystem** — a record at the
centre with the rest in orbit, and a single binary that gathers them
all (the metaphor is a solar system; see
[impire.io/soulsystem](https://impire.io/soulsystem)). The code lives in
the component repositories (sibling checkouts:
[`../soulstream`](../soulstream) the record,
[`../soulrealm`](../soulrealm) the room,
[`../soulidentity`](../soulidentity) the name,
[`../soulnode`](../soulnode) the house, [`../soulfold`](../soulfold) the
fold), along with their frozen per-feature spec-kit artifacts
(`specs/NNN-*/`) and plain-words `docs/`. Everything else — why the
ecosystem exists, what we're investigating, what we've designed, what we're
building, and what happened — lives in one of the areas below.

| Area | What it holds | When you touch it |
|---|---|---|
| [`00-GENESIS/`](00-GENESIS/README.md) | One vision, one constitution, one working agreement, the founding rationale | When deciding *whether* / *how* to do something |
| [`01-RESEARCH/`](01-RESEARCH/README.md) | Active research topics (one folder each, any component) | While investigating an open question |
| [`02-DESIGN/`](02-DESIGN/README.md) | The normative design, per component (soulstream's `core/`+`extensions/`, numbered designs, D-numbered docs) | When research graduates, or a build changes behavior |
| [`03-IMPLEMENTATION/`](03-IMPLEMENTATION/README.md) | The ecosystem roadmap: what to build, in what order, behind which gate | When planning or landing a feature |
| [`04-JOURNEY/`](04-JOURNEY/README.md) | Numbered episodes, one shared sequence: the honest log of what happened | Whenever a feature lands, research concludes, or a load-bearing decision is made |

`99-ARCHIVE/` holds superseded material — the five pre-merge genesis sets and
soulstream's pre-restructure design set — kept for provenance with its
content frozen; it is history, not the live design.

## The pipeline

```
01-RESEARCH ──graduates──▶ 02-DESIGN/<component> ──/speckit-specify──▶ <component repo> specs/NNN + code
     │                         ▲                                            │
     │ (abandoned)             │ (behavioral changes                        │
     │                         │  propagate back)                           │
     ▼                         │                                            ▼
04-JOURNEY ◀──────── every ending writes an episode ◀──────── 03-IMPLEMENTATION
                                                                (ROADMAP updated)
```

- Research topics live in `01-RESEARCH/<slug>/`, name their component, and
  end in exactly one of three states: **graduated to design**, **graduated to
  artifact**, or **abandoned**. Every ending produces a numbered episode in
  `04-JOURNEY/`; the topic folder is then removed (git history keeps the full
  trail).
- Designs in `02-DESIGN/<component>/` are written functional-level and
  explicit enough to hand to `/speckit-specify` in the component repo.
  Implementation always goes through the component's flow (spec-kit, or the
  design's D-numbers); research never does.
- The lifecycle transitions are mechanized as skills: `/research-start`,
  `/research-graduate`, `/journey-log` — available here and, retargeted, in
  every component repo. The structure itself is enforced by
  `internal/hqlint`, which rides the standard quality gate (`make test`).
- This repository carries the full pre-merge history of every file: the five
  per-project `hq/` lineages were merged here with `git filter-repo` on
  2026-08-02 (episode
  [0050](04-JOURNEY/0050-ecosystem-one-hq.md)). Old per-project episode
  numbers resolve through the map in
  [`04-JOURNEY/README.md`](04-JOURNEY/README.md).

**If in doubt** — about whether to build something, how to decide, or whether
a shortcut is acceptable — the answer is in
[`00-GENESIS/`](00-GENESIS/README.md). Hold the decision against `vision.md`
and `constitution.md`; if it still isn't clear, that's a teach-back
conversation, not a judgment call to make alone.
