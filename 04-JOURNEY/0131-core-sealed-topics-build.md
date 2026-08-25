# Episode 0131 — Sealed topics build: the locked binder (2026-08-25)

The build of the design-validated extension
([`extensions/sealed-topics.md`](../02-DESIGN/soulstream-core/extensions/sealed-topics.md);
research bars measured 2026-07-28, [episode
0011](0011-soulstream-sealed-topics.md)) — feature `021-sealed-topics`,
merged to core main `724f10d`, 37 files, +4372/−86, whole gate green
with workloads re-verified against the merged tree. Both research
amendments held as the law they were written to be: epoch and nonce
ride **inside** the signed `sealed.op` payload (a keyless reader
verifies an epoch rewrite flips the signature `[measured:
TestSealedOpSignatureCoversEpoch]`), and the AAD — newline-joined
`(realmKey, path, op-id, author, parents, epoch)` — refuses every
splice arm: cross-topic, cross-author, wrong parents, wrong epoch,
wrong op-id, wrong realm `[measured: TestSealAADSpliceMatrix]`, plus a
server-replay splice across topics `[measured]`.

What shipped, per package: `identity` grows the X25519 `SealingKey`
and the `Unwrapper` custody seam mirroring `Signer` — structurally
satisfied, no seed in the contract, so soulstream-identity's D9
custodian plugs in later without an import; `registry` grows
`sealing_key` + rotations (strict decode ships first — the Bar 2
ordering), endorsement through any validated Ed25519 chain key,
`EnsureSealingKey`/`RotateSealingKey`; `topic` grows the
`sealed.op`/`sealed.epoch` vocabulary, `StartSealedTopic`, the decrypt
pass before the unchanged pure fold (non-members materialise structure
with warnings, never errors `[measured]`), epoch bump and explicit
history handoff, the sealed rollup **re-carrying the wrapped-key
table** — the Bar 3 amendment proven live: a fresh member reads from
the post-purge baseline `[measured: TestSealedTopicFullLife]` — sealed
attachments (digest over ciphertext), and sealed mention bodies;
curator goes blind on purpose and collective search hands out nothing
readable `[measured]`; the keystore and CLI carry the `.x25519` sibling
seed with `key sealing init|show|publish|rotate`. Crypto as designed:
XChaCha20-Poly1305 for content; anonymous sealed boxes (`nacl/box`)
for the member wrap — sender authenticity already lives in the signed
op, so a sender-keyed primitive would be redundant surface
`[mechanism-argument]`.

One defect the build caught in the design's own blind spot: deriving
the topic id from the display name would have leaked the very name the
announcement encrypts — sealed topic ids now derive from the word
`sealed`, never the name. And one refusal sharpened: a mentioned
persona without a published sealing key is skipped with a warning,
never notified in cleartext. Both propagated into the design doc in
this change.

Reversal condition: the design's Bar 4 flip conditions stand unchanged
— membership scale or churn meaningfully beyond ~10, or member devices
no longer assumed trusted, moves MLS from upgrade path to prerequisite
and reopens the key-scheme decision. For the build itself: none beyond
those — records a completed, measured implementation.

Trail: `soulstream-core/specs/021-sealed-topics/` (all seven
artifacts) + `docs/sealed-topics.md`; branch `021-sealed-topics`
(seven signed commits, tip `1db797b`, merged `724f10d`, unreleased —
tagging stays a human act); design
[`extensions/sealed-topics.md`](../02-DESIGN/soulstream-core/extensions/sealed-topics.md);
research [episode 0011](0011-soulstream-sealed-topics.md).
