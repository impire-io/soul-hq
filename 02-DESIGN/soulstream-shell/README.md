# 02-DESIGN / soulstream-shell — the shell

The normative design for **soulstream-shell**, Soulstream's human
surface — the browser entry beside the MCP door. Numbered documents, soulstream-workloads/
soulstream convention.

| Doc | What it decides |
|---|---|
| [`0001-soulhelm-the-helm.md`](0001-soulhelm-the-helm.md) | The founding design: placement, surfaces, rendering architecture, identity and custody, the design-system contract, acceptance criteria. |
| [`0002-the-module-shape.md`](0002-the-module-shape.md) | The module shape: the pure shell, the support layer, the contract (activation, navigation, routes, links), purity as a standing gate — graduated from `shell-module-contract` (episode [0078](../../04-JOURNEY/0078-shell-the-module-contract.md)). |
| [`0003-conversation-lifecycle.md`](0003-conversation-lifecycle.md) | Conversation lifecycle from the shell: starting (rail fold + Home card), the close-then-archive ladder with its two-step confirm, the archived fold, truthful copy under partial failure — class (a) of 0001 §4, participation scope of episode [0071](../../04-JOURNEY/0071-ecosystem-the-focus.md). |
| [`0005-the-tools-module.md`](0005-the-tools-module.md) | The tools module: the catalog with each person's own standing, connect/disconnect as the person's own ceremony through the module's callback, the admin's both-halves add — the external-tools design's human end. Designed and built 2026-08-21, episode [0122](../../04-JOURNEY/0122-ecosystem-the-shell-arc-lands.md). |
| [`0006-the-approvals-module.md`](0006-the-approvals-module.md) | The approvals module: the guardrail's tickets on a screen, one tap minting the person's signed answer and delivering it to the originator's tail — mint and delivery kept apart, activation by the deployment's declared GuardrailOn fact. Designed and built 2026-08-21, episode [0122](../../04-JOURNEY/0122-ecosystem-the-shell-arc-lands.md). |
| [`0004-the-storage-explorer.md`](0004-the-storage-explorer.md) | The storage explorer: reading the stores themselves — the list over a subject pattern, one message whole, the live tail — on the signed-in person's own admission rather than the shared read lane (amending 0001 §3 for this surface), with search and any persistent index refused as the query layer the protocol declines. Designed and built 2026-08-19, episodes [0116](../../04-JOURNEY/0116-ecosystem-what-shipped-without-a-human-end.md)/[0117](../../04-JOURNEY/0117-shell-the-store-shows-what-it-holds.md). |
| [`0007-the-sheet-shape.md`](0007-the-sheet-shape.md) | The sheet shape: the list-screen rhythm (tables lead, add-forms in the slide-over on the frame's panel signal), the `.stow` fold, kind-branched forms, the ask a destructive key stands behind, the person's words on rows, and the details drawer below 1180px. Designed and built 2026-08-23, episode [0123](../../04-JOURNEY/0123-shell-the-sheet-shape.md). |
| [`0008-the-first-hour.md`](0008-the-first-hour.md) | The first hour: guidance as a reading, never a store — the derived first-steps card on Home, the arrival principle (an act ends when the realm's own evidence shows the thing arrived), empty states that offer their act, and upstream ask #3 (the wrap's presence lease — announce on start, renew on a cadence, staleness read as gone). Decided 2026-08-24 (episode [0124](../../04-JOURNEY/0124-ecosystem-the-first-hour-and-the-presence-lease.md)), built 2026-08-25 as v0.11.0-rc.3 (episode [0127](../../04-JOURNEY/0127-shell-the-first-hour-builds.md)). |
| [`0009-the-declare-surface.md`](0009-the-declare-surface.md) | The declare surface: the agents sheet's second lane — author a declaration in a form whose output IS the CLI's JSON (no second schema), submit on the session's own admission (the measured Bar 5 loop), arrival as open→claimed read live from the record with the no-dispatcher realm answered in honest words; the placements and models lists as readings; no retirement offered (the vocabulary does not exist), no secrets through the shell (D36 is caller-own), one new dependency named openly (workloads' declaration+fleet). Designed 2026-08-28 (the agents-as-infrastructure arc's last human chapter). |
| [`0010-the-models-surface.md`](0010-the-models-surface.md) | The models surface: a Models sheet — the catalogue's names as a table everyone reads, authoring/re-pointing as a slide-over form byte-identical to the CLI's entry (one codec via upstream ask #1), removal behind an ask that counts the declared agents naming the name, Serving now from the plane's own discovery; writes as the person's own act on their own admission (`$KV.>` is the persona scope's), the admin gate named as a courtesy line, no secret and no door-key surface, no lifecycle act. Closes 0009 §4's catalogue-writing [O]. Drafted 2026-08-28, built 2026-08-29 (episode [0151](../../04-JOURNEY/0151-shell-the-models-surface.md)). |

Born from research topic `soulsystem-cockpit` (graduated 2026-08-13,
episode
[0066](../../04-JOURNEY/0066-ecosystem-soulsystem-cockpit.md); the
topic folder lives in git history). The repo was founded the same day
([impire-io/soulstream-shell](https://github.com/impire-io/soulstream-shell), v0.1.0,
episode [0067](../../04-JOURNEY/0067-soulhelm-founding-and-first-light.md))
and `soulstream-shell` is a full component of the hq: legal journey tag,
constellation entry, this design folder.
