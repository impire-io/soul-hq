# 02-DESIGN / soulstream-shell — the shell

The normative design for **soulstream-shell**, Soulstream's human
surface — the browser entry beside the MCP door. Numbered documents, soulstream-workloads/
soulstream convention.

| Doc | What it decides |
|---|---|
| [`0001-soulhelm-the-helm.md`](0001-soulhelm-the-helm.md) | The founding design: placement, surfaces, rendering architecture, identity and custody, the design-system contract, acceptance criteria. |
| [`0002-the-module-shape.md`](0002-the-module-shape.md) | The module shape: the pure shell, the support layer, the contract (activation, navigation, routes, links), purity as a standing gate — graduated from `shell-module-contract` (episode [0078](../../04-JOURNEY/0078-shell-the-module-contract.md)). |
| [`0003-conversation-lifecycle.md`](0003-conversation-lifecycle.md) | Conversation lifecycle from the shell: starting (rail fold + Home card), the close-then-archive ladder with its two-step confirm, the archived fold, truthful copy under partial failure — class (a) of 0001 §4, participation scope of episode [0071](../../04-JOURNEY/0071-ecosystem-the-focus.md). |
| [`0004-the-storage-explorer.md`](0004-the-storage-explorer.md) | The storage explorer: reading the stores themselves — the list over a subject pattern, one message whole, the live tail — on the signed-in person's own admission rather than the shared read lane (amending 0001 §3 for this surface), with search and any persistent index refused as the query layer the protocol declines. Designed and built 2026-08-19, episodes [0116](../../04-JOURNEY/0116-ecosystem-what-shipped-without-a-human-end.md)/[0117](../../04-JOURNEY/0117-shell-the-store-shows-what-it-holds.md). |

Born from research topic `soulsystem-cockpit` (graduated 2026-08-13,
episode
[0066](../../04-JOURNEY/0066-ecosystem-soulsystem-cockpit.md); the
topic folder lives in git history). The repo was founded the same day
([impire-io/soulstream-shell](https://github.com/impire-io/soulstream-shell), v0.1.0,
episode [0067](../../04-JOURNEY/0067-soulhelm-founding-and-first-light.md))
and `soulstream-shell` is a full component of the hq: legal journey tag,
constellation entry, this design folder.
