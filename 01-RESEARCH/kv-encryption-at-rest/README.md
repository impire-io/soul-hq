# Should KV records be sealed with xkeys — and where would the key live?

**Component:** soulfold
**State:** active
**Started:** 2026-08-02

## Abstract

The operator's direction: the fold's KV entries should be encrypted,
using the NATS-native key machinery. The precise hypothesis to test is
**xkey (X25519 curve key) sealing at the application layer** — nkeys
proper are Ed25519 *signing* keys and cannot encrypt; the `nkeys`
library's curve keys (`Seal`/`Open`) are the encryption primitive the
ecosystem offers. This gates M1: the record envelope (plaintext JSON
vs sealed blob) sits inside the store-shape one-way door, so deciding
after M1 lands means a stated migration. The topic must answer against
the alternatives (nats-server filestore encryption at rest;
field-level sealing of only the sensitive records) and must produce a
key-custody story that makes encryption more than obfuscation — a seed
stored beside the ciphertext defends nothing. It knowingly revisits
part of store design D2's reasoning (records inspectable with stock
NATS tooling); the design doc absorbs whatever verdict the bars
produce.

## The question

Does xkey-sealing of KV records preserve everything the store design
already proved (restart, additive evolution, CAS), against which
threat does it actually defend compared to server-side filestore
encryption, and what key-custody story — birth, location, rotation,
loss — makes it worth the inspectability it costs?

## Pre-registered bars

- **Bar 1 — the envelope preserves the proven store properties.**
  Protocol: re-run the store design's three mechanic rigs with a
  seal/unseal layer in place (xkey `Seal` on write, `Open` on read;
  plaintext inside remains the D2/D6 JSON): restart round-trip
  (ciphertext byte-identical, plaintext decodes), the schema-N ↔ N+1
  additive matrix through the envelope, and the CAS race on sealed
  records. Pass: all three green with numbers matching the unsealed
  baselines (zero lost updates; full matrix) [measured].
- **Bar 2 — custody is real, not decorative.** Protocol: with the rig
  stopped, take the complete store directory (and a full stream dump)
  and attempt recovery of any record plaintext using only those
  artifacts. Pass: no plaintext recoverable from store contents alone;
  the seal seed demonstrably lives outside the store; and the custody
  story is written — where the seed is born, where each deployment
  shape keeps it (single binary, embedded, shared JetStream), how
  re-keying works (a stated re-seal migration), and what seed loss
  means (honestly: total data loss) [measured for the recovery
  attempt; the story itself is design matter].
- **Bar 3 — the alternative is measured, not dismissed.** Protocol:
  the same working set on a nats-server with JetStream filestore
  encryption enabled (`cipher`); restart round-trip green; a disk scan
  for known plaintext markers comes back empty; then state which
  threat each option covers — server-side encryption defends the disk
  but not a NATS-API-level reader; app-layer sealing defends both but
  costs inspectability and custody complexity. Pass: both cells
  demonstrated, the threat table written [measured].
- **Bar 4 — the cost is known.** Protocol: seal/unseal micro-benchmark
  over the M1 record shapes (ops/s, added bytes per record) plus the
  flow rig's end-to-end sign-in with the envelope on. Pass: overhead
  quantified; added p50 per record operation under 1 ms and the
  sign-in flow's added wall time under 10 ms, or the numbers recorded
  as the reason to reverse [measured].

## Reversal condition

The custody story degenerates: if for every supported deployment shape
the seal seed necessarily ends up readable by the same principal that
can already read the ciphertext (observable: the rig's recovery
attempt succeeds once the deployment's own config/env is included in
the artifact set, in all shapes), then app-layer sealing adds a
migration burden and an outage class (seed loss) without covering a
real threat — and the direction reverses to server-side filestore
encryption plus constitution-I data minimization, recorded as such.

## Verdict

**Graduated to design, 2026-08-02. Four bars, four passes.** The rig
(session scratchpad `xkeyrig/`, stack: nats-server v2.14.4, nats.go
v1.52.0, nkeys v0.4.16, zitadel/oidc v3.48.1, go-oidc v3.20.0) ran
every arm against an embedded JetStream server, sealed and unsealed,
with positive controls on every scan.

- **Bar 1 — PASS [measured].** Restart round-trip: 7/7 records
  byte-identical ciphertext across a full server stop, plaintext
  decodes to the originals, buckets found by lookup. Additive matrix:
  24/24 writer×reader cells over six record shapes decode identically
  sealed vs unsealed, plus the D3 RMW demonstration through the
  envelope (the trap is unchanged — a design rule, not a mechanism the
  envelope could fix). CAS: 8 writers × 1,000 cycles landed at exactly
  8,000 sealed (49,243 rejections all retried, ~3,165 accepted
  contended writes/s vs ~7,116 unsealed — half the throughput, still
  orders beyond an IdP's write rate); code redemption won exactly once
  in 100/100 races of 8, both arms.
- **Bar 2 — PASS [measured].** With the rig stopped: zero marker hits
  (raw and base64) across the complete store directory, zero in an
  API-level dump of every key and value — including the username,
  because the amendment this bar forced digests the `idx.username.*`
  key (a plaintext username in a KV *key* would hand the scan exactly
  what the envelope withholds). The seal seed lives in a sibling
  custody dir, mechanically asserted outside the store. Positive
  control: 5 disk + 6 API hits unsealed — the scan is proven, not
  assumed. The custody story is design matter (D17): the seed is born
  at first start, lives outside the store dir per deployment shape,
  re-keying is a stated re-seal migration, seed loss is total data
  loss.
- **Bar 3 — PASS [measured].** JetStream filestore encryption
  (ChaCha): restart round-trip green, disk scan zero hits — and the
  decisive asymmetry demonstrated, not asserted: an ordinary NATS-API
  reader saw 6 plaintext hits under filestore encryption. Server-side
  encryption defends the disk; only the app-layer envelope also
  defends the NATS surface — the surface D1 explicitly shares with a
  parent deployment. Threat table in the design (D19).
- **Bar 4 — PASS [measured].** Seal p50 ≈ 57.4 µs, open p50 ≈ 57.2 µs
  (~17,400/s), +44 bytes per record flat across all six shapes. Added
  p50 per KV record operation: +45.5 µs (put), +57.2 µs (get) — bar
  was 1 ms. End-to-end sign-in (stock go-oidc RP, authorization-code +
  PKCE, JWT access token, id_token verified against published JWKS):
  median 3.23 ms unsealed → 4.41 ms sealed, **+1.19 ms** — bar was
  10 ms.

The reversal condition never approached: the custody story does not
degenerate — in the shared-JetStream shape the seed is unreachable
from the store's own surfaces, and even single-binary deployments keep
the seed out of the store-dir artifact set (a full-machine backup that
grabs both is named honestly in D17 as plaintext-equivalent
[judgment]). Direction confirmed: **app-layer xkey sealing ships as
the fold's store envelope** (D16–D19 in
[store-and-key-lifecycle](../../02-DESIGN/soulfold/store-and-key-lifecycle.md));
filestore encryption stays available as defense-in-depth, not a
substitute.
