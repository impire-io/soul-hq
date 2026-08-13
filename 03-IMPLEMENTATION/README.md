# 03-IMPLEMENTATION — what gets built, in what order

| File | Role |
|---|---|
| [`ROADMAP.md`](ROADMAP.md) | The ecosystem's live plan: a cross-component state table, then one section per component with its milestones, exit criteria, research gates, and one-way doors. No dates — gates, not calendars. |
| [`DOGFOOD.md`](DOGFOOD.md) | The soulstream dogfood run's protocol: the two-week scenario, how each persona launches, and the evidence duty that feeds the eg-walker and sealed-topics gates. |

## Conventions

- **Roadmap ↔ journey ↔ specs mapping:** a roadmap item is built in its
  component's repository — through the spec-kit flow (`/speckit-specify` →
  plan → tasks → implement, artifacts frozen in `<component>/specs/NNN-*/`)
  for the spec-kit components, or against the design's D-numbers for
  soulstream-identity and soulstream-idp — and lands together with a numbered episode in
  [`../04-JOURNEY/`](../04-JOURNEY/README.md). Feature numbers come from git
  branches in the component repo; episode numbers from the shared journey
  sequence; release versions from that repo's git tags (`v*`).
- **Landing a feature means, in the same working session:** the component
  repo's quality gate green and the feature merged, plus the soul-hq commit
  updating the roadmap with the measured outcome, writing the journey episode
  (`/journey-log`), and propagating behavioral changes into the
  [`../02-DESIGN/`](../02-DESIGN/README.md) docs they touch.
- **Exit criteria are written before the work** and amended only openly. The
  roadmap file itself is load-bearing: changes to it are decisions and belong
  in the journey as episodes like any other.
- **Frozen specs.** A `specs/NNN-*/` body in a component repo is a
  point-in-time artifact and is not rewritten after the feature ships; only
  its `Status` field records the shipping version. Pre-merge spec bodies
  reference journey episodes by their old per-project numbers — the map in
  [`../04-JOURNEY/README.md`](../04-JOURNEY/README.md) resolves them.
