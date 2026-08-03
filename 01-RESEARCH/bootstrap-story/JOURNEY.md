# JOURNEY — bootstrap-story (started 2026-08-03)

## 2026-08-03 — the rig is the fold, all four bars, no amendments

Prototyped the invite mechanism directly against the fold on its
lifecycle branch (the store, ceremonies, and HTTP surface are real —
the strongest measurement position): digest-stored single-use invites
as the only enrollment right; Begin refuses credential-less users
without one; Finish consumes the invite by CAS before the credential
binds. Bars measured in `internal/lifecycle/bars_test.go` (2, 3) and
`internal/serve/bootstrap_test.go` (1, plus the M3 gate observables);
Bar 4 is the pocket-id desk audit (docs.pocket-id.org — notably: its
first admin arrives via an open /setup page, which the audit REFUSES
for the fold). All four passed on their first full runs; no bar was
amended. Numbers in the verdict.
