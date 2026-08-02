# soul-hq — project instructions

The headquarters of the Soulstream ecosystem — soulstream (the record),
soulrealm (the room), soulidentity (the name), soulnode (the house),
soulfold (the fold). This repo holds research, designs, the roadmap, and the
journey for all five; code and frozen specs live in the sibling component
repos (`../soulstream`, `../soulrealm`, `../soulidentity`, `../soulnode`,
`../soulfold`).

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
  `04-JOURNEY/` (components: soulstream, soulrealm, soulidentity, soulnode,
  soulfold, ecosystem — one shared sequence). Use `/journey-log`; research
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
