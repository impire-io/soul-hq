# Episode 0134 — Tenants are born admissible: D47 lands (2026-08-27)

The defect episode 0133 measured closed the same day, both halves in
one commit (identity `447ec6b`, unreleased): `accounts.create` now
births a tenant whose minted users actually work. The signing key
lands as a **scoped signer carrying the canonical persona template** —
before this, the plain key left every `SetScoped` mint admitted but
inert (0 subscriptions, 0 payload), so 0110's "the mint path serves
the new tenant the moment the op returns" was false through admission
— and creation **amends the AUTH account's `allowed_accounts`** with
the new tenant (lookup, add, re-land complete; idempotent; tenant
first, AUTH second, so the between-acts window fails closed rather
than half-open). The upgraded live test admits the mint-shaped user
and proves it usable: subscribe + round trip inside the scope in
**2.77ms** store-to-usable-admission [measured], a publish outside the
scope drawing the server's permissions violation (the template is
applied, not merely present), and the tenant read back from AUTH's
stored JWT [measured]. Suspension/resume unchanged and re-proven
through the scoped shape [measured].

The one-source discipline D47 demanded is now code: the persona
template is exported once from `client`
(`PersonaScopePubAllow`/`PersonaScopeSubAllow`, prefix-aware), the
tenancy authority renders it, and the founding ceremony can adopt the
same export when the product next touches `ceremony.go` — until then
the two render identical lists by construction of the export
[mechanism-argument]. `AuthAccount` empty skips the coupling honestly:
a deployment with no callout half has no admission list to maintain.

Named residues, unchanged: the ProviderAPI arm's D47 parity (the BYON
account shape is the provider's — A8's standing operator-act residue),
and the product's halves (`SystemConn` unwired, no `accounts.*` client
surface) ride whichever product cycle adopts multi-tenancy.

Reversal condition: if the AuthAccount-empty skip proves to hide
misconfiguration — a deployment running the callout whose tenancy
authority was not told the AUTH account, tenants born unreachable
(observable: recorded as an issue) — the coupling becomes required
configuration validated at assembly, as a new D-decision.

Trail: design
[`platform-topology.md`](../02-DESIGN/soulstream-identity/platform-topology.md)
(D47, as-built notes propagated in this change); identity commit
`447ec6b` (fix + live-test proof); episode
[0133](0133-ecosystem-platform-account-topology.md) (the measurement
that found both halves).
