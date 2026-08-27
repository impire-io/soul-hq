# Episode 0135 — The residues close: tenancy reaches the hand, both arms carry D47 (2026-08-27)

Episode 0134 left two named residues; the operator asked whether to
address them, and both closed the same day — one measured to its end,
one carried to the edge of what this machine can measure.

**The tenancy surface exists, and its acceptance criterion is met
through the real op family** (identity `df8e4a3`). The `accounts.*`
family had no public caller since it shipped; the client now mirrors
it (`AccountCreate/Resolve/Accounts/Suspend/Resume`), and the new e2e
drives D47's criterion entirely through public surfaces (`embed.Run` +
the sealed client): a tenant born over `accounts.create`, a token
issued for it, a sentinel+token client admitted **usable** — subscribe
and round trip inside the persona scope in **11.2ms**
create-to-admission, a foreign publish drawing the permissions
violation, suspend refusing the next admission, resume restoring it
[measured]. One observation recorded: during suspension the issuer
still mints; the server (Conn=0) is the refusal — the verifier of
record, working as designed.

**The provider arm carries D47** (identity `f6a1a33`): the
programmatic signing-key group is created WITH the persona template as
its scope (`SigningKeyGroupCreateRequest.Scope` — the same inert-user
defect existed on BYON), and creation teaches the AUTH account each
tenant by read-union-write on its `jwt_settings` authorization
(idempotent, fail-closed between acts). The byon_live measurement was
upgraded to match — the round trip mints the real scoped shape, and a
throwaway AUTH stand-in proves the coupling against the real control
plane without touching the system's actual AUTH. **Honest bound**: no
control-plane token is persisted on this machine (custody working as
designed), so the live run remains the operator's one command:
`SOULSTREAM_CP_TOKEN=… SOULSTREAM_CP_SYSTEM=… go test -tags byon_live
-run TestBar1ProviderArm ./internal/accounts/ -v` — compile-proven
under the tag, unmeasured until then [mechanism-argument].

**The house grew tenants** (soulstream `8b458dd`, spec
`012-tenants-in-the-house`, merged `87e456a`) — the product's half the
0133 verdict named, taken as a full cycle at the operator's direction.
Three composition decisions, each an amendment to design 0001
propagated in this change:
- **The resolver persists** (§3): the embedded server moved from
  `MemAccResolver` to a dir resolver under `<state>/resolver`, seeded
  **create-if-absent** from the founding JWTs — never overwrite,
  because the runtime amends stored JWTs (AUTH learns each tenant,
  D47) and re-seeding founding shapes would silently unlearn every
  tenant at restart. Found live: the dir resolver refuses bare
  `TrustedKeys`; the operator claims it requires are synthesized in
  memory from the ceremony's operator seed.
- **In-memory synthesis over new artifacts**: the SYS user (the
  `SystemConn`'s credential) and the operator claims are built per
  start from seeds the ceremony already persists — nothing new on
  disk, nothing to migrate; realms founded before tenancy gain the
  capability on their next `up` (the operator key enters the vault as
  an ensure at start, F1's posture).
- **The hand**: `soulstream account create|list|show|suspend|resume`
  drives the sealed ops over the operator's creds. BYO stays honestly
  off — no operator material on that side (design 0003), and the
  service's own refusal is what the command shows.

The 012 gate, measured: `account create` → usable token-lane admission
**8.8ms** through the running house; the tenant **still admits after
the node stops and starts** on the same state directory (the
persistence clause — resolver dir intact, AUTH's amendment kept, the
vault binding minting); suspend/resume through the same surface; the
M1.1 gate unchanged and green on the dir resolver [all measured].

What remains, named: the BYON live run (the operator's token), and the
per-tenant realm provisioning + shared-service builds (D46/D48/D49)
behind the focus gate — a created tenant is admissible and isolated
today, but its `SOULSTREAM` stream arrives when the record's
provisioning reaches multi-tenant composition.

Reversal condition: if the create-if-absent resolver seed proves to
mask a needed founding-JWT update (observable: a ceremony change that
must reach an existing realm's resolver and cannot), the seeding rule
gains an explicit re-seed act as a new decision — never a silent
overwrite.

Trail: identity `df8e4a3` (client surface + op-family e2e), `f6a1a33`
(provider-arm parity), `95a8d9e` (kind vocabulary); soulstream
`8b458dd`/`87e456a` (spec `012-tenants-in-the-house`); designs
[`platform-topology.md`](../02-DESIGN/soulstream-identity/platform-topology.md)
(as-built notes) and
[`0001-soulnode-composition.md`](../02-DESIGN/soulstream/0001-soulnode-composition.md)
(§3 resolver amendment); episodes
[0133](0133-ecosystem-platform-account-topology.md)/[0134](0134-identity-tenants-born-admissible.md).
