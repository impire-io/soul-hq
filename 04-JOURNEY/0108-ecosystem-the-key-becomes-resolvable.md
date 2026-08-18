# Episode 0108 — F1 closes: every signing persona becomes readable (2026-08-18)

The first build from episode 0107's design order, landed the same
evening: `registry.EnsureSigningKey` (soulstream-core **v0.9.0**) —
create-or-add, idempotent, display metadata preserved, a different
stored key refused as `ErrKeyConflict` so rotation stays the one
key-change door — and the three `PersonaSigner` consumers now call it
at signer construction: the remote node's per-user pool
(soulstream-mcp **v0.2.0**), and the product's archivist and runner
(on soulstream main). Loud on failure, never refusing admission or a
launch. The gap this closes was the shipped default [measured, 0107's
trace]: no reader consults `keys.public`, so every
identity-plane-signed persona rendered unknown-key until someone
published a profile by hand. Now the profile appears the moment the
signer does.

Reversal condition: none — records a completed build. (The
warn-and-continue failure mode reopens as a refusal if a deployment
measures silent unknown-key persistence despite the ensure call —
observable as the warn line recurring across restarts.)

Trail: design
[`extensions/tenancy.md`](../02-DESIGN/soulstream-core/extensions/tenancy.md)
(registry additions); core `bf69fd3` (v0.9.0), soulstream-mcp
`7822e3c` (v0.2.0), soulstream `a840497`; episode
[0107](0107-ecosystem-platform-tenancy-guardrails.md) (the confirmed
gap and the owner decision).
