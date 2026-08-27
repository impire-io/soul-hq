# Episode 0138 — The sealed record gains its custodian: D9 builds (2026-08-27)

The oldest standing "later" on identity's roadmap — sealing keys, D9,
"waits on Soulstream sealed topics build" — was unblocked by episode
0131 and is built. The design pass came first
([`sealing-keys.md`](../02-DESIGN/soulstream-identity/sealing-keys.md),
D50–D53): design propagation, not research, because core's `Unwrapper`
seam fixed and measured the whole mechanism (NaCl anonymous sealed box,
32-byte X25519, base64 public halves) and what remained were decisions —
the vault kind with D26's reserved first-touch clause realized (D50),
one `seal.unwrap` op releasing artifacts and never material (D51, the
D32 parallel), one directory door (`keys.public` grows the `sealing/`
grammar; the persona template grows exactly one tail — D52), and
consumer-side ensure-publication as the F1 posture applied to sealing
(D53). The supposed D26-versus-registry tension the exploration flagged
dissolved on inspection: F1 already made consumer-side ensure-publication
of endorsed public halves the shipped default, and
`registry.EnsureSealingKey` takes exactly a `Signer` and a public string
[mechanism-argument, anchored in episode 0108].

The build (identity `specs/005-sealing-custody`, merged `72bd164`,
unreleased) is the `PersonaSigner` story told again for sealing, mirror
by mirror: `KindPersonaSealingKey` with the owner binding, in-vault
`box.OpenAnonymous`, `GenerateSealingKey` closing the cross-instance
first-touch race with an `ErrExists` re-Get, `ownedSealingKey`'s
no-probe refusals (foreign keys refuse identically whether they exist or
not — pinned by unit), and the client's `PersonaUnwrapper` satisfying
core's seam structurally with fail-fast foreign-owner construction.
The D9 gate rides `make test` in consumer position [measured]: two
custodial members run the full D53 ceremony, a sealed topic carries
three messages across two epochs, and the member materialises full
plaintext through a counting custodian at **exactly 2 unwraps — one per
epoch, never per message** (D9's no-oracle line, now a number); the
sealed mention body opens through the same one op (the arbitrary-length
shape, proving the vault's derivation byte-compatible with core's wrap
targets); the no-unwrapper view stays structure-only; the directory
serves public halves only, and `ExportSeed` remains reachable solely
through user-key credential export — no path to a sealing seed exists
[measured; code-trace for the export claim].

What it opened, named in the design's [O] ledger: rotation (with the
honest problem stated — old epoch wraps were sealed to the old key, so
rotation is never replace-and-forget), batch unwrap (reopens if
materialise latency measures dominated by round trips; today's number
is one per epoch), product wiring (no consumer constructs sealed
topics yet — grep-verified; surfaces wire the custodian F1-style when
sealed topics reach the house), and the scope re-render note for
pre-existing accounts (templates without the new tail don't gain the op
silently). The e2e modules pin core `v0.14.0-rc.1`; the pin re-lands at
the core release.

Reversal condition: if oracle-shaped traffic appears in the audit log
(observable: `seal.unwrap` rates tracking message volume rather than
epoch changes on a deployment), D51's one-op stance is re-examined
before any batch convenience is added — the gate tightens, it never
widens.

Trail: design
[`sealing-keys.md`](../02-DESIGN/soulstream-identity/sealing-keys.md)
(D50–D53; agent.md D9 and nats-surface.md D26 carry the realized
pointers); identity `specs/005-sealing-custody` (`ee2ab18`, merged
`72bd164`); core's handover in `specs/021-sealed-topics` (episode
[0131](0131-core-sealed-topics-build.md)); the F1 precedent (episode
[0108](0108-ecosystem-the-key-becomes-resolvable.md)).
