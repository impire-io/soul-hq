# AGENTS — orientation and reading order

This repository is the headquarters of the Soulstream ecosystem:
soulstream-core (the record), soulstream-workloads (the room),
soulstream-identity (the name), soulstream (the product, the house),
soulstream-idp (the fold), soulstream-shell (the shell),
soulstream-mcp (the remote door), soulstream-archivist (the keeper),
soulstream-inference (the mind).
Code lives in the sibling component repos; research, designs, the
roadmap, and the journey live here.

## Reading order

1. [`README.md`](README.md) — the one-screen map of the areas and the
   pipeline.
2. [`00-GENESIS/vision.md`](00-GENESIS/vision.md) and
   [`00-GENESIS/constitution.md`](00-GENESIS/constitution.md) — what the
   ecosystem is and the articles no work may violate (shared S1–S5 +
   per-component articles keeping their original numbering).
3. [`00-GENESIS/how-we-work.md`](00-GENESIS/how-we-work.md) — the cross-repo
   pipeline, the research lifecycle, and the working agreement in daily
   terms.
4. [`04-JOURNEY/README.md`](04-JOURNEY/README.md) — where things stand, per
   component, and the honest log.
5. [`03-IMPLEMENTATION/ROADMAP.md`](03-IMPLEMENTATION/ROADMAP.md) — the live
   plan: gates, not calendars.

## The non-negotiables, in brief

- **The decision test** ([`00-GENESIS/README.md`](00-GENESIS/README.md)):
  vision, then constitution, then the working agreement. If it doesn't
  produce a clear answer, the decision waits for the human.
- **The journey duty:** every landed feature, concluded research topic, or
  load-bearing decision gets a numbered episode
  (`NNNN-<component>-<slug>.md`) in [`04-JOURNEY/`](04-JOURNEY/README.md) —
  `/journey-log` does this (research topics get theirs via
  `/research-graduate`). Landing work in a component repo and writing its
  episode here happen in the same working session.
- **Research never goes through spec-kit; implementation always follows a
  design.** Open questions run `/research-start` →
  `/research-graduate` with pre-registered bars.
- **Quality gate:** `make fmt && make test && make lint` here — all green,
  nothing skipped; `make test` includes the structural lint
  (`internal/hqlint`). The component repos have their own gates.
- **Claims carry evidence classes** (`[measured]` / `[mechanism-argument]` /
  `[judgment]`); only measured closes a debate. Direction decisions record
  their reversal condition when they are made.
- Sign every commit. Never push from a skill — pushing stays a human act.
  Never commit `.claude/settings.local.json`.
