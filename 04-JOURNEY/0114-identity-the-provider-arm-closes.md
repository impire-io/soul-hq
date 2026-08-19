# Episode 0114 — The provider arm closes: an account born on someone else's operator key (2026-08-19)

Episode 0107 graduated with exactly one named residue — Bar 1's
**provider arm**, the runtime-account-birth measurement re-run where a
hosting provider custodies the operator key and exposes only an API.
The operator handed over the Synadia control-plane token this morning;
the arm is now measured, and the tenancy topic has no open residue at
all. soulstream-identity **v0.9.0**.

**The backend** (D35's `Authority` seam, second implementation): account
birth is the account plus its **programmatic signing-key group**, whose
seed the platform returns exactly once and which goes straight to the
vault bound to the new account — nothing in the backend retains it, and
the operator key is never ours to hold. Suspension is the connection
limit dropped to zero, the provider-side twin of the local arm's
re-landed JWT.

**Measured live on the Impire DEV system** [measured]: an account born
at runtime in **~51s**, a principal minted against its group admitted
**through the real cloud** and completing a publish/subscribe round
trip, suspend and resume both landing, and the pre-existing account
resolving uninterrupted across **22 probes** spanning the whole window
— zero restarts of anything, and nothing in the system touched but the
two throwaway `bar1-probe-*` accounts, which the rig deleted.

**The honest number, beside the local arm's 1.69ms: ~51 seconds** — the
provider arm is four orders of magnitude slower to bring an account into
existence. That is propagation through someone else's control plane and
JWT distribution, not our path, and it is exactly the kind of fact A8
existed to surface: the two custody arrangements have materially
different runtime properties, and a deployment that expects
sub-second tenancy must hold its own operator key.

Two defects the live run forced, both fixed rather than worked around:
the control plane answers **5xx on a just-created account's next call**
while it settles — the same lossy-channel class the product's founding
driver already retries (episode 0099), so create, group-create,
suspend, and delete now share one retry discipline; and a cleanup path
using a bare context **cannot reach the API at all**, because the SDK
carries base URL and token on the context, not in headers. A first
failed run left three accounts behind, which is why a
`TestCleanupProbeAccounts` rig now exists and why it ran before the
definitive measurement.

The measurement never rides the default gate: it is `//go:build
byon_live`, needs `SOULSTREAM_CP_TOKEN`/`SOULSTREAM_CP_SYSTEM`/
`SOULSTREAM_CP_NATS_URL`, and skips without them.

Reversal condition: none — records a completed measurement. (D35's
`Authority` seam stands or falls with the two arms it now has; a third
provider whose API cannot express suspension without destroying data
would reopen the suspend contract, not the seam.)

Trail: design
[`tenancy.md`](../02-DESIGN/soulstream-identity/tenancy.md) (D35, A7/A8);
identity `3bdf8c7` (v0.9.0, `internal/accounts/providerapi.go` +
`byon_live_test.go`); episodes
[0107](0107-ecosystem-platform-tenancy-guardrails.md) (where the
residue was named), [0110](0110-identity-the-tenancy-set-builds.md)
(the local arm at 1.69ms), [0099](0099-soulstream-the-byon-founding.md)
(the lossy-channel finding this build inherited).
