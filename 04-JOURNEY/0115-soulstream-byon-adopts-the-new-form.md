# Episode 0115 — byon adopts the new form: the migration that refused to be a re-founding (2026-08-19)

The operator asked for the byon realm to be re-founded on rc.7 — the
migration episode 0112's refusal names. Looking at the realm first
changed the answer: **its op-log held zero records.** Nothing had ever
been signed there. The break exists because v1 signatures can never be
re-signed; where none exist, re-founding protects nothing and costs
real things — the fold's enrolled passkey users, the realm's name, and
a fresh set of Synadia accounts (the existing accounts' signing-key
seeds were returned exactly once and are custodied in the old state
directory, so a same-name founding is blocked at the provider anyway).

So the refusal got precise instead of merciful. **`soulstream adopt`**
(rc.8) reads the realm's own op-log and decides: an empty log adopts —
stamping the canonical version and leaving every other field founding
wrote exactly as it was, including keys this build does not recognise —
a populated one refuses **by count** with re-founding spelled out, and
`--force` takes the cost deliberately while saying what it costs. The
ceremony's `AdoptV2` is the only writer of that field outside founding
and decides nothing itself; the weighing lives in the command.

**The deployment, on beno1** [measured]: state directory backed up,
service stopped, rc.8 installed, `soulstream adopt` answering *"the
op-log was empty, so nothing was mixed"*, service restarted — every
plane serving (identity, memory, MCP, sign-in, shell console). Then the
proof, run on the host against the live realm: a signed turn posted and
read back — **wire version 2, the acting credential present, the
signature verified** against a keyring built from the directory. What
adoption preserved, checked rather than assumed: the fold's 8 user
records (the operator's enrolled passkey among them), its 6 clients and
2 keys, the identity vault's 4 entries, both Synadia accounts.

Two things the live run surfaced:

**F1 closed itself in production.** The personas bucket had been empty
for three days; the moment rc.8 started, the archivist's profile
appeared — `EnsureSigningKey` (episode 0108) publishing at signer
construction, on a realm that had been quietly unverifiable.

**The realm identity is a minted key, not the account key.** A10
prefers the deployment's real account key, but the product's founding
path calls `ProvisionOn` without a connection, so the connectionless
branch minted one. Recorded rather than papered over: byon's identity
is `ADDHFS6…` and stays so (first-provision-wins), both forms are
opaque keys that bind signatures identically, and the gap is now closed
in code — `realm.WithConn` (core v0.11.2) plus both founding paths
passing their ops connection, unreleased on main so it rides the next
rc.

Reversal condition: if a realm ever adopts and *then* proves to hold v1
records the count missed (observable: a `legacy-shape` verdict on a
realm adopted as empty), the count is reading the wrong thing — the
precondition must move from stream-message-count to a scan for signed
records, and adopt must re-verify before trusting itself.

Trail: product `9c05120` (v0.13.0-rc.8, `cmd/soulstream/adopt.go`,
`ceremony/adopt_test.go`), `0730a61` (the WithConn wiring, unreleased);
core `v0.11.2`; episodes
[0112](0112-ecosystem-the-canonical-form-breaks-clean.md) (the break
and its refusal), [0108](0108-ecosystem-the-key-becomes-resolvable.md)
(the F1 act that closed itself here),
[0099](0099-soulstream-the-byon-founding.md) (the realm this migrates).
