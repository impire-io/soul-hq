# soul-hq — project instructions

The headquarters of the Soulstream ecosystem — soulstream-core (the
record), soulstream-workloads (the room), soulstream-identity (the
name), soulstream (the product, the house), soulstream-idp (the fold),
soulstream-shell (the shell), soulstream-mcp (the remote door),
soulstream-archivist (the keeper), soulstream-inference (the mind).
This repo holds research, designs,
the roadmap, and the journey for all of them; code and frozen specs
live in the sibling component repos (`../soulstream-core`,
`../soulstream-workloads`, `../soulstream-identity`, `../soulstream`,
`../soulstream-idp`, `../soulstream-shell`, `../soulstream-mcp`,
`../soulstream-archivist`, `../soulstream-inference`).

**Read [`AGENTS.md`](AGENTS.md) first** (reading order + the
non-negotiables), then hold decisions against
[`00-GENESIS/`](00-GENESIS/README.md). Where things stand:
[`04-JOURNEY/README.md`](04-JOURNEY/README.md). The plan:
[`03-IMPLEMENTATION/ROADMAP.md`](03-IMPLEMENTATION/ROADMAP.md).

Conventions:

- Quality gate before every commit: `make fmt && make test && make lint` —
  all green, none skipped; `make test` includes the hq structural lint
  (`internal/hqlint`).
- The journey duty: every landed feature, concluded research topic, or
  load-bearing decision gets an episode `NNNN-<component>-<slug>.md` in
  `04-JOURNEY/` (component tags, single-word by the episode grammar:
  `core`, `workloads`, `identity`, `idp`, `shell`, `mcp`, `cli`,
  `inference`, `soulstream` for the product, `ecosystem` for
  cross-cutting — one
  shared sequence; episodes ≤ 0069 keep the pre-rename tags, resolved
  by the naming map in `04-JOURNEY/README.md`). Use `/journey-log`; research
  via `/research-graduate`. Refresh the index, the component's "Where things
  stand", and the roadmap in the same change.
- Episodes 0001–0049 carry the five per-project journals (merged
  2026-08-02); the numbering map in `04-JOURNEY/README.md` resolves old
  per-project references found in frozen specs and old commit messages.
- Links to component-repo content use sibling paths
  (`../<component>/specs/...`); hqlint skips them, so verify them by eye
  when you write them.
- Sign every commit. Never push from a skill — pushing stays a human act.
  Never commit `.claude/settings.local.json`.
