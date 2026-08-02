# JOURNEY — kv-encryption-at-rest (started 2026-08-02)

## 2026-08-02 — the rig, all four bars, one amendment

Built `xkeyrig/` in the session scratchpad (stack: nats-server v2.14.4,
nats.go v1.52.0, nkeys v0.4.16, zitadel/oidc v3.48.1, go-oidc v3.20.0):
the envelope (`nkeys.CreateCurveKeys`, self-addressed `Seal`/`Open`, raw
sealed bytes as the KV value), the four D1 buckets, the D6 shapes at
schema 1 and 2, and re-runs of the store design's three mechanic rigs
with the seal layer underneath — plus the filestore-cipher arm, the
custody scans with positive controls, and a real zitadel/oidc provider
with a stock go-oidc RP for the flow cost.

**The amendment the custody bar forced:** D6's `idx.username.<username>`
key leaks the username *in the KV key* — the store-dir scan would find
it before any value is read. Bar 2's pass required digesting the index
key (`idx.username.<hex(sha256(lower(username)))[:32]>`), the same rule
D12 applies to bearer secrets. Exact-match lookup is all the fold ever
needs there, so nothing is lost. Recorded as a D6 amendment in the
graduation, raw numbers in the verdict.

All four bars passed on the first full-suite run after the amendment;
verdict numbers in the topic README. No other bar was amended.
