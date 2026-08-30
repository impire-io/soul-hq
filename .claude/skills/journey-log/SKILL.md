---
name: "journey-log"
description: "Append a numbered journey episode for a completed feature, concluded investigation, or load-bearing decision; refresh the index, the component's Where-things-stand, and the roadmap."
argument-hint: "What happened (feature landed / decision made), or empty to log the work just completed in this session"
compatibility: "Requires the soul-hq structure (04-JOURNEY/TEMPLATE.md)"
metadata:
  author: "soul-hq"
user-invocable: true
disable-model-invocation: false
---

## User Input

```text
$ARGUMENTS
```

If `$ARGUMENTS` is empty, the subject is the work completed in the current
session; reconstruct it from the conversation and `git log` (including the
component repo's log when the work landed there).

## Steps

1. **Scope check.** An episode is warranted for: a landed feature, a concluded
   investigation (research topics get theirs via `/research-graduate` instead —
   do not duplicate), or a load-bearing decision (spec change, criterion
   amendment, refuted hypothesis, direction call). If none applies, say so and
   stop. Identify the component tag (`core`, `workloads`, `identity`,
   `idp`, `shell`, `mcp`, `cli`, `inference`, `soulstream` for the
   product, or `ecosystem` for cross-cutting work).

2. **Write the episode.** Next free number `NNNN` in `04-JOURNEY/`, file
   `NNNN-<component>-<short-kebab-slug>.md`, following `04-JOURNEY/TEMPLATE.md`
   exactly: what happened with the honest numbers (spreads, not bare means),
   what was refuted or reversed, what it taught or opened, evidence-class tags
   (`[measured]` / `[mechanism-argument]` / `[judgment]`) on load-bearing
   claims, the trail (docs, sibling-repo specs, commits), and the required
   **Reversal condition:** line. For direction decisions the working agreement
   applies first: teach-back before recording, adversarial pass for calls that
   change a protocol's shape or a core boundary
   (`00-GENESIS/constitution.md`, Part II).

3. **Refresh the surroundings, same change-set:** add the episode to the index
   in `04-JOURNEY/README.md`; refresh the component's "Where things stand"
   section there; update the affected item in the component's section of
   `03-IMPLEMENTATION/ROADMAP.md`; propagate behavioral changes into the
   `02-DESIGN/<component>/` docs they touch.

4. **Gate, commit (never push).** Full quality gate green
   (`make fmt && make test && make lint` — includes the `internal/hqlint`
   structural lint under `make test`). Stage only the touched paths by explicit
   pathspec (never `git add .`/`-A`); signed commit (`git commit -S`):
   `docs(journey): NNNN — <title>` with the standard co-author trailer
   (`Co-Authored-By: Claude Fable 5 <noreply@anthropic.com>`). The episode
   commit lands here in the same working session as the feature commit in the
   component repo. Remind the human to push both repos.
