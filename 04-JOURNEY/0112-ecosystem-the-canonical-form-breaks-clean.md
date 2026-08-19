# Episode 0112 — The canonical form breaks clean: the realm's key signs, the record names the hand (2026-08-19)

The two wire-format one-way doors the operator took at episode 0107's
review, executed across the whole ecosystem in one pass —
soulstream-core **v0.11.0/v0.11.1**, every consumer re-pinned and
re-tagged, the product at **v0.13.0-rc.7**.

**A10 — the realm's key binds every signature.** The canonical form's
realm value is now the realm's cryptographic identity, not its reusable
display name: born at first provision from the connection's real
account key where the deployment has one and a minted key otherwise
(`realm.Identity`, `EnsureIdentity`, `Client.RealmKey`), first
provision winning forever, the seed deliberately dropped because
nothing ever signs with it — it is a name in key form. A name-scoped
signature could be replayed into any realm reusing the name; a
key-scoped one cannot. Names demote to display, and resolution
(`accounts.resolve`, D35) can be wrong without breaking verification.
Portable exhibits carry the key too, so offline verification is
alias-proof end to end — the archivist's keeper was stamping the
display name and its verdicts read `failed` until it followed.

**E3 — every record names whose hand held the pen.** `Soulstream-Acting`
is required on v2 records: the publishing identity, distinct from the
author, `Config.Acting` for the assistant case. Two evidence grades, as
decided: on the custodial lane the identity plane **refuses to sign** a
canonical whose acting field names anyone but the server-proven caller
(identity v0.8.0), which is what makes the field provenance rather than
a self-claim; on the self-custody lane it is honest testimony.

**Reads never hard-fail.** v1 records still parse, and their signatures
verify as the named status **`legacy-shape`** — never conflated with
`failed`, because they were signed over a form that no longer exists.
Signing on a realm with no identity refuses with the migration in the
error text. The product refuses a pre-v2 state directory by name at
load, with re-founding spelled out, and stamps the canonical version it
was founded under (the byname-era refusal's idiom, reused).

What the sweep cost, honestly: the whole ecosystem's fixtures gained
the acting field, and three real defects surfaced — the miss-caching
`RealmKey` (a client connected before provisioning never learned the
identity), `ProvisionOn` reaching the streams but not the identity (the
product's own founding path, caught by the shell's gate), and the
archivist's name-stamped exhibits. Every other consumer — workloads,
mcp, shell — compiled and passed **unchanged**, which is the seam
discipline paying out.

Reversal condition: the designs' conditions stand (A10's
exhibit-resolution case; E3's unknowable-actor lane). New from the
build: if `legacy-shape` proves to be read as failure by any surface
that matters (observable: a reader treating it as tamper evidence),
the status gains explicit rendering rather than being folded into
`unknown-key`.

Trail: designs
[`extensions/tenancy.md`](../02-DESIGN/soulstream-core/extensions/tenancy.md)
(A10/E3 as recorded at graduation); core `b1f8371` (v0.11.0),
`13c9357` (v0.11.1); identity `f1ed68d` (v0.8.0, the custodial stamp),
archivist `cbfda68`, and the pin-only releases across workloads, mcp,
and shell; product `a98e7a5` (v0.13.0-rc.7); episode
[0107](0107-ecosystem-platform-tenancy-guardrails.md) (where the doors
were taken).
