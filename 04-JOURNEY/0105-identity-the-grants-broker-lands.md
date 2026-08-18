# Episode 0105 — The grants broker lands: outbound identity in custody (2026-08-18)

Episode 0104's design (D30–D34) became running, released code in one
overnight-plus-morning arc: the operator's 2026-08-17 directive built
slice 1 on a branch, and the morning review completed the slice and
merged it — soulstream-identity **v0.3.0**. The `grants.*` op family
serves on the principal-scoped surface; per-persona OAuth custody lives
in its own sealed CAS bucket (the key vault untouched, D31); the only
thing any caller ever receives is the derived short-lived access token
(D32); on-behalf-of is a subject-signed, bounded delegation honored only
from its actor, both personas audited on every decision (D33). Merged
fast-forward, two commits, +2,453 lines; gate green uncached across all
three modules (8 packages ok, 0 lint issues, no skips) [measured].

The review verdict on the overnight slice: **sound** — and three
additions rode the review pass. The delegation gained a not-before check
(a future-dated `issued_at` now refuses as not yet valid; D33 amended);
the no-key-subject refusal path got direct tests at both layers (the
op-level test had been refusing via subject mismatch while its comment
claimed the no-key path); and the provider now sends
`Accept: application/json` — **found writing the runbook, not running
it**: GitHub's token endpoint answers form-encoded without that header
[mechanism-argument from GitHub's documented contract; the fix is
exercised only against JSON stand-ins until SC-005 runs], so the
real-provider walk would have failed on its first step.

The slice's own gate now carries the transport clause in consumer
position [measured]: `e2e/embedgate` `TestGrantsGate` runs the full M4
admission ceremony with the scope template grown by exactly one line
(the grants op tail), links a persona against a *strict* rotating
stand-in AS (a stale refresh token refuses, so the second access passing
proves the rotated successor was custodied), then proves isolation the
server's way — the imposter's publish to the victim's grants subject
dies as a permissions violation on the imposter's own connection, the
request never answered, and the delivery log shows the victim's subject
served exactly twice: her own two calls. Revocation deletes custody and
the next access refuses. 2.2s inside `make test`.

Two build decisions were kept at review and propagated into the design:
access contention is bounded by time (5s), not round count (a fixed
retry count was measured starving a loser at three attempts), and a
contender whose redemption fails polls briefly for the record's revision
to move before declaring the line dead — the loser can observe the
winner's redeem→CAS-write gap from outside and must not mistake it for
a lost line. Named residue, accepted: a revoke racing a rotation
best-effort-revokes the pre-rotation token upstream; custody deletion is
the decision either way. The CLI grew the ceremony
(`grant link|access|ls|revoke`), and the README states the deployment's
scope-template duty beside the two existing shapes (D25).

What stays open, honestly: **SC-005, the real-provider walk** — the
runbook is written into the feature's quickstart (GitHub preferred, its
rotating refresh tokens exercise D31 live; Google as the stable-token
alternative with the offline ask baked into the catalog's AuthURL), but
registering an app and consenting in a browser is a human act, and the
Bar 2 residue from 0104 stays open until the operator runs it. Also
still upstream of consumers: the wrapper's per-run overlay seam
(workloads' roadmap) to carry a delegation into a run, and the fold's
token-lifetime knob and RFC 8693 exchange (its roadmap, called due by
0104).

Reversal condition: if the real-provider walk fails in a way the
stand-in could not show (observable: a failed SC-005 run recorded in the
feature's quickstart checklist), the provider surface reopens before any
consumer wires in. D31's own reversal stands: a provider whose rotation
semantics cannot survive the redeem-then-write order (observable: a
measured lost line with the discipline followed) forces a write-ahead
intent record into the grant store.

Trail: design [`grants.md`](../02-DESIGN/soulstream-identity/grants.md)
(D30–D34, amended at review); episode
[0104](0104-ecosystem-outbound-identity-grants.md) (the graduated
research this builds); soulstream-identity
`specs/003-grants-broker/` (spec, plan, tasks with the review record,
quickstart with the SC-005 runbook); commits `4c19201` (slice 1,
overnight) and `dd38817` (the review pass), merged to main and tagged
`v0.3.0`.
