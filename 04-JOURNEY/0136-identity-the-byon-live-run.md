# Episode 0136 — The BYON live run: two defects caught, then the provider arm measures sound (2026-08-27)

The one honest bound episode 0135 named — the provider arm
compile-proven but unmeasured, the control-plane token the operator's
alone — was paid today. Three runs of `TestBar1ProviderArm` against
the real Synadia Cloud BYON system: the first two each caught a defect
no local rig or compile proof could see; the third measured the arm
sound end to end.

**Run one: 400 Bad Request on the D47 coupling.** The amend wrote
`allowed_accounts` alone into the AUTH account's authorization patch,
and the JWT law refuses that shape: external authorization cannot have
accounts without users (nats-io/jwt). Two defects behind the one
status code: the rig's throwaway AUTH stand-in was born bare — no
`auth_users`, so no allowed-accounts write on it could ever be valid —
and the production path itself never carried `auth_users`/`xkey`
forward, which under replace-the-object patch semantics would have
stripped the real AUTH's callout users. Fixed fail-closed (identity
`a0545c8`): the amend reads the whole authorization object and writes
it back whole — correct whether the control plane merges the patch per
field or replaces the object, so the semantics question never needs
answering — a userless AUTH is refused by name as
not-a-callout-account, and the stand-in is seeded one throwaway auth
user so it has the shape a real AUTH has. This is exactly why 0135
tagged the arm [mechanism-argument] and not [measured]: the mechanism
argument was wrong twice.

**Run two: the coupling held live, and two more edges surfaced.** AUTH
learned both tenants — but the round-trip clause could not run: a BYON
system's control-plane view exposes no client-reachable URL
(`direct_connection_opts` unset; the control plane reaches the server
through its agent), so the URL is the operator's fact, passed
explicitly. And the failure path leaked: `t.Fatalf` skipped the
probe's stop, the goroutine outlived the test body, and a straggling
iteration re-created account A after cleanup had freed its name — a
leaked probe account in the real system. Fixed (identity `31279c6`):
the probe's stop is a `t.Cleanup` registered after the deletion
cleanup (LIFO — it runs first) that joins the goroutine before any
name is freed; `TestCleanupProbeAccounts` swept the leak.

**Run three: PASS [measured].** Account births 4.36s and 3.29s through
the control plane, against Bar 2's 5s admission bound (the local arm
births in 395–776µs — the provider's API and resolver push set the
pace, and the bound still holds); the scoped round trip alive in the
newborn account — D47's inert defect is gone on BYON, the minted user
subscribing and publishing inside the persona scope; the AUTH
stand-in's `allowed_accounts` carrying both tenants when read back
from the control plane; suspend refusing and resume restoring through
the provider API; the pre-existing account probed 12 times through the
window, uninterrupted; zero probe accounts left in the system.

What it taught: a live run is not a formality after a compile proof —
both defects lived in the seam between our request shape and the
provider's validation, the one place no local stand-in can testify.
And the partial-patch defect was the dangerous kind: on a
merge-semantics control plane it fails loud; on a replace-semantics
one it would have silently disconnected every callout user of the
real AUTH. With this, every residue of the 0133 topology verdict is
closed — both arms of D47 are measured, local and BYON.

Reversal condition: if the whole-object write proves to race another
writer of the AUTH account's authorization (observable: an auth user
or xkey added between the read and the write disappearing after an
`accounts.create`), the amend gains the provider's concurrency guard
or narrows to a field patch with the merge semantics confirmed in
writing — never a silent lost update.

Trail: identity `a0545c8` (whole-object authorization write + faithful
stand-in), `31279c6` (the probe joins before cleanup); run log
2026-08-27 (three runs: 400 → URL unset + the leak → PASS); design
[`platform-topology.md`](../02-DESIGN/soulstream-identity/platform-topology.md)
(as-built note updated); episodes
[0133](0133-ecosystem-platform-account-topology.md)/[0134](0134-identity-tenants-born-admissible.md)/[0135](0135-ecosystem-the-residues-close.md).
