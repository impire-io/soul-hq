# Episode 0051 — The envelope is decided: xkeys seal the store (2026-08-02)

M1's last gating research asked whether the fold's KV records should be
sealed with the ecosystem's own key machinery — xkey (X25519 curve key)
sealing at the application layer, the operator's stated hypothesis —
and what key-custody story makes that more than obfuscation. The
question gated M1 because the record envelope (plaintext JSON vs sealed
blob) sits inside the store-shape one-way door. Four bars were
pre-registered against the alternatives (JetStream filestore
encryption; field-level sealing) before the rig ran.

All four passed, against an embedded nats-server v2.14.4, sealed and
unsealed arms, positive controls on every scan [measured]:

- **The envelope preserves the proven store properties.** Restart
  round-trip 7/7 byte-identical ciphertext, plaintext decoding to the
  originals; the additive matrix 24/24 cells identical sealed vs
  unsealed plus the D3 RMW trap demonstrated unchanged through the
  envelope; CAS landing at exactly 8,000/8,000 with 49,243 rejections
  all retried (~3,165 accepted contended writes/s sealed vs ~7,116
  unsealed — half the throughput, orders beyond an IdP's write rate);
  code redemption exactly-once in 100/100 races, both arms.
- **Custody is real.** Rig stopped: zero marker hits across the store
  directory and a full API-level dump — raw and base64, keys and
  values — with the seal seed mechanically outside the store; the
  unsealed control found 5 disk + 6 API hits, so the scan proves
  something. One amendment forced here: D6's `idx.username.<username>`
  key would have leaked the username in the KV *key*, so the index key
  is now digested (`idx.username.<hex(sha256(lower(username)))[:32]>`)
  — D12's rule extended from bearer secrets to user-supplied names.
- **The alternative was measured, not dismissed.** Filestore
  encryption (ChaCha): restart green, disk scan clean — and an
  ordinary NATS-API reader still saw 6 plaintext hits. Server-side
  encryption defends the disk; only the app-layer envelope also
  defends the NATS surface that D1 explicitly shares with a parent
  deployment. That asymmetry, demonstrated live, is the decision.
- **The cost is known.** Seal/open p50 ≈ 57 µs (~17,400/s), +44 bytes
  per record flat; added p50 per KV record operation +45.5 µs put /
  +57.2 µs get (bar: 1 ms); end-to-end sign-in — a stock go-oidc RP
  through authorization-code + PKCE against a real zitadel/oidc
  provider on the sealed store, tokens verified against published JWKS
  — median 3.23 ms → 4.41 ms, **+1.19 ms** (bar: 10 ms).

Nothing was refuted; the reversal condition (custody degenerating in
every deployment shape) was never approached — the shared-JetStream
shape keeps the seed unreachable from the store's own surfaces, and the
single-binary shape keeps it out of the store-dir artifact set, with
the full-machine-backup caveat named honestly in the design
[judgment].

What it opened: the store design grows the envelope (D16–D19 in
[store-and-key-lifecycle](../02-DESIGN/soulfold/store-and-key-lifecycle.md)) and
amends D6's username index. **M1 is now unblocked** — every research
gate the roadmap names for it has concluded to design. The rig's
provider + storage shape is the M1 skeleton's reference.

Reversal condition: a supported deployment shape whose seed custody
demonstrably collapses to "readable by every principal that can read
the ciphertext" (the Bar-2 recovery succeeding once that shape's own
config/env joins the artifact set), or measured envelope overhead
breaching the pre-registered bars under real fold load — either
reopens this decision toward filestore-encryption-plus-minimization,
as the topic pre-registered.

Trail: verdict and topic journey in git history at
`01-RESEARCH/kv-encryption-at-rest/` (folder removed by this
graduation); design
[store-and-key-lifecycle](../02-DESIGN/soulfold/store-and-key-lifecycle.md)
(D6 amended, D16–D19 added); rig in the session scratchpad
(`xkeyrig/` — envelope, matrix, CAS, custody scans, filestore arm,
flow rig), stack pinned in the verdict.
