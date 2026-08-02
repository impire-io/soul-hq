# Episode 0050 — One hq: the five headquarters merge (2026-08-02)

The operator's direction: holistic work across the soul projects kept
requiring five separate visits — soulstream, soulrealm, soulidentity,
soulnode, and soulfold each ran their own `hq/` (same rules, five trees).
This repository, **soul-hq**, is the merge: one headquarters for the
ecosystem, following the same rules the five followed individually.

What happened, with the honest numbers [measured]:

- The five `hq/` lineages were extracted with `git filter-repo` and merged
  here with **full history**: 125 source commits carrying 141 files, every
  pre-merge path reachable with `git log --follow`. This matters because the
  research lifecycle deletes topic folders on graduation — git history *is*
  the trail, so a snapshot copy would have lost the record.
- **04-JOURNEY**: all 49 episodes renumbered into one chronological sequence
  (ordered by each episode's original commit time; same-commit ties keep
  their original relative order), filenames now `NNNN-<component>-<slug>.md`.
  The pre-merge numbering map in [`README.md`](README.md) resolves every old
  reference; episode prose was left untouched, so in-text mentions of old
  numbers ("episode 0005") read through the map. Link targets were rewritten
  mechanically (123 links across episodes and designs).
- **00-GENESIS**: one vision, one constitution (v1.0.0) — shared articles
  S1–S5 factored from the five ancestors, the Working Agreement carried
  verbatim (it was already identical), and per-component articles keeping
  their original numbering so existing citations still resolve. The five
  founding sets are frozen in
  [`../99-ARCHIVE/genesis/`](../99-ARCHIVE/genesis/) (ancestors: soulstream
  1.1.0, soulrealm 0.1.1, soulidentity 1.3.1, soulnode 1.0.0, soulfold
  1.0.0).
- **02-DESIGN**: per-component folders, each keeping its own convention
  (core/extensions, `NNNN-` numbers, D-numbers). **03-IMPLEMENTATION**: one
  ROADMAP with a cross-component state table; the five roadmaps carried
  whole, every gate state intact — soulfold M1 still gated on the open
  `kv-encryption-at-rest` topic (the only live research topic, moved in
  as-is; concluded later the same day —
  [episode 0051](0051-soulfold-kv-encryption-at-rest.md)), soulstream
  sealed-topics still gated on the dogfood chafe log (to 2026-08-10),
  soulnode's public door still gated on soulfold upstream.
- **Enforcement moved with the structure**: a unified `internal/hqlint`
  rides `make test` here (layout, component tags, contiguous numbering,
  reversal-condition lines, research states, link resolution); the
  lifecycle skills (`/research-start`, `/research-graduate`, `/journey-log`)
  live here and, retargeted, in each component repo. The per-project `hq/`
  trees, their `internal/hqlint` packages, and their skill copies are
  removed from the component repos in the same working session; each repo's
  `.specify/memory/constitution.md` now symlinks to the merged constitution.

What was superseded rather than lost: the five per-project genesis sets,
area READMEs, templates, and roadmap files — their content lives on in the
merged documents, their texts in `99-ARCHIVE/` or git history [measured: the
migration's inventory check accounted for every pre-merge file as moved,
archived, or carried into a merged document].

What it opened: cross-component work — a research topic, design change, or
roadmap re-sequencing that touches several components — is now one commit in
one repository, with the `ecosystem` component tag for exactly that class of
work.

Reversal condition: if the merged hq measurably obstructs per-component work
— observable as component landings repeatedly blocked on soul-hq conflicts,
or a component's roadmap section and its repo's actual state drifting apart
across multiple landings because the hq is "too far away" to update in the
same session — the merge un-splits along the same filter-repo path it came
in by, with the numbering map run in reverse.

Trail: this repository's first six commits (the genesis commit, five
`chore: absorb <component>/hq` merges, and the reshape series); the
numbering map in [`README.md`](README.md); the removal commits in the five
component repos.
