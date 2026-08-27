# Sealing keys — the custodian for sealed topics (D50–D53)

*Realizes D9 (unwrap once, no decrypt oracle) against soulstream-core's
shipped `Unwrapper` seam (core `specs/021-sealed-topics`, episode 0131,
which hands over: "Remote sealing-key custodians — the `Unwrapper` seam
ships; the custodian implementation is soulstream-identity's D9"). This
document is design propagation, not research: the mechanism is fixed and
measured on the core side (NaCl anonymous sealed box to a 32-byte X25519
key; public half = std base64 of the 32 raw bytes; `Unwrap` never
`(nil, nil)`; implementations own their deadlines and concurrency
safety), and the remaining choices are decisions, taken here.*

## Context

Core's sealed topics call an `identity.Unwrapper` in exactly two places:
`collectSealState` (per-epoch wrapped **epoch keys**, 32 bytes out,
failures fold as warnings) and `Notification.Unseal` (sealed
`mention.notify` bodies, arbitrary length). Wiring is per-handle
(`Handle.UseSealing`); wrap targets come from the realm registry's
endorsed sealing keys (`registry.EnsureSealingKey`, which requires the
signing key published first). The seam is structurally satisfied — no
import in either direction; a consumer wires the two (the `Signer`
cycle-guard discipline, unchanged).

## D50 — The sealing key is a vault kind, owner-bound, first-touch materialized

A persona's sealing key is a vault record of kind
`persona-sealing-key`, named `sealing/<persona>`, carrying the same
`(account, user)` owner binding as its signing key (D6/D25). It
materializes **on first touch** — `keys.public` on the caller's own
sealing name, or the unwrap op reaching it — exactly the D26 pattern
that decision reserved for this moment ("the persona's X25519 sealing
key follows this same materialization pattern — decided now as the
pattern, built then"). Seed = std base64 of 32 raw X25519 bytes; public
half = std base64 of `curve25519.X25519(seed, basepoint)` —
byte-compatible with core's `SealingKeyFromSeed`, proven in the
consumer-position e2e. There is **no export sibling**: D7's one custody
escape stays creds-only. Cost carried openly, same as D26:
first-owner-wins on the name; a foreign-owned name refuses identically
whether it exists or not (no vault probing).

## D51 — One unwrap op; the released secret is an artifact, not identity material

`seal.unwrap` opens an anonymous sealed box addressed to the caller's
OWN sealing key, inside the vault, and releases the plaintext to the
session. Article I holds: the epoch key is "already a shared group
secret, not an identity" (D9) — the parallel of D32's grants ruling
that custody produces artifacts and never surrenders the material that
makes them; the X25519 seed never crosses any wire. One op serves both
core call shapes (32-byte epoch keys and arbitrary-length notify
bodies) because the seam itself makes no distinction and an opaque
wrapped blob cannot be classified by anyone but the box — so no length
guardrail pretends otherwise. The oracle risk D9 names is structurally
limited: sealed boxes exist only where core wraps (epoch keys, notify
bodies); message ciphertext is ChaCha20-Poly1305 under AAD and cannot
feed `box.OpenAnonymous`. Per-message decryption never gets an op.
Every unwrap is audited with the principal and the wrapped length —
never the bytes.

## D52 — One directory door

`keys.public` accepts the `sealing/` name grammar: the caller's own
sealing name materializes (D50); any persona's sealing public half is
an open directory read — the vault that custodies the keys IS the
directory (D26, unchanged). No `seal.public` sibling exists. The
canonical persona scope template (D47's one exported source,
`client.PersonaScopePubAllow`) grows exactly one new tail:
`seal.unwrap`. Deployments founded from the old template do not gain
the op until their account JWT re-renders — an ops note, not a silent
upgrade.

## D53 — Consumer-side ensure-publication (F1 applied to sealing)

Whoever constructs a `PersonaUnwrapper` pairs it with
`registry.EnsureSealingKey(ctx, rc, signer, unwrapper.PublicKey())`,
endorsed by the same custodial `PersonaSigner` — the F1 posture
(episode 0108: consumers ensure-publish endorsed public halves at
construction; the client library cannot call core, the cycle guard).
Core enforces the ordering: signing key published first. The client's
`PersonaUnwrapper` mirrors `PersonaSigner` line for line — constructor
resolves the public half once (the first touch) and fails fast on a
foreign owner; `Unwrap` is one sealed round trip, never `(nil, nil)`,
deadline owned by the client's per-request timeout.

## Deferred, named [O]

- **Rotation.** Vault records are immutable (`ErrExists`; import under
  a new name); core's `RotateSealingKey` has zero consumers. The real
  problem a rotation build must answer: old epoch wraps in topic
  history were sealed to the old key — rotation is never
  replace-and-forget; the sketch is versioned names
  (`sealing/<persona>/2`, highest = current for `PublicKey()`), unwrap
  routing current-then-history. Waits for a consumer that rotates.
- **Batch unwrap.** `collectSealState` costs one sealed round trip per
  epoch per materialise; epochs move on membership change only. The
  gate records the measured count; a batch op reopens if materialise
  latency measures dominated by unwrap round trips.
- **Product wiring.** No consumer constructs `SealingKey`/`UseSealing`
  outside core's own tests today (grep-verified across the ecosystem);
  the custodian ships with its consumer-position proof in identity's
  `e2e/`, and product surfaces wire it F1-style when sealed topics
  reach the house.
- **Scope re-render.** Pre-existing accounts carry templates without
  the `seal.unwrap` tail (D52's ops note).

## Where things stand

Designed 2026-08-27; build follows in identity `specs/005-sealing-custody`.
