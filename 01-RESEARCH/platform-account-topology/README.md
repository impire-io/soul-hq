# Can the platform's services live in one dedicated account, serving many tenant accounts?

**Component:** ecosystem
**State:** active
**Started:** 2026-08-26

## Abstract

A realm is normatively one NATS account (protocol §"realm"; 0064 S1),
and account lifecycle already runs at runtime on both authority arms
(D35, episodes 0110/0114). What does not exist is the topology the
operator now asks for: a dedicated platform-services account beside
per-tenant accounts — today every platform service (identity surface,
ops, archivist, runner, signin) connects with creds signed by the one
realm account itself, and the built `accounts.*` family is dark in the
shipped product (no `SystemConn`). This topic investigates whether the
platform account is sound: it must make real the cross-account seam
D15 designed and deferred (`account_token_position` exports), answer
the AUTH `allowed_accounts` coupling, re-pin persona-name scope for a
many-account deployment, and argue honestly against the D14 amendment,
which chose subject prefixes over "account gymnastics" for environment
separation. A decisive PASS unlocks multi-tenant hosting as
configuration; a FAIL keeps services per-tenant-account and closes the
question with the reasons recorded. The 0071 focus gate's demand
condition has arguably fired — this is the product asking.

## The question

Can the Soulstream platform services run in one dedicated platform
account, serving many per-tenant accounts, without weakening the
server-enforced principal proof (D15), tenant isolation, or one-act
tenant birth (D35)?

## Pre-registered bars

- **Bar 1 — the cross-account surface preserves the principal
  proof.** Protocol: a rig on a dir-resolver operator-mode server; the
  platform account exports the identity surface with
  `account_token_position = P+2`; two tenant accounts (A, B) import
  it. Pass: a caller in A reaches its own
  `<prefix>.identity.<A-key>.<user>.>` ops through the import with the
  server having forced A's account key into the token, and every
  attempt from A to name B's key — or to reach the surface at all with
  a forged account token — is refused **by the server**, shown in the
  delivery log, with zero service-side authorization decisions added.
- **Bar 2 — tenant birth is one act, end to end, through the
  product.** Protocol: the composed node with `SystemConn` wired;
  `accounts.create` on the local authority arm, then realm
  provisioning, then first admission through the token lane (which
  forces the AUTH `allowed_accounts` coupling to a measured answer).
  Pass, over ≥ 6 runs reported as spreads: account store → admitted
  round trip ≤ 100ms; create-to-first-token-admission ≤ 5s including
  realm provisioning; zero restarts; zero edits to existing tenant
  accounts; an existing tenant's continuous probe uninterrupted
  (max gap reported); across all pre-creation probes, zero usable
  half-tenants (0107 Bars 1–2 restated at the new scope).
- **Bar 3 — isolation holds through the shared services.** Protocol:
  one platform archivist and one runtime serving tenants A and B;
  adversarial probes from a principal in A. Pass: no subject, stream,
  KV bucket, or object in B is readable or writable from A —
  server-refused with delivery-log proof — including via any surface
  the shared services expose; and every service-performed act lands in
  exactly one tenant's stream, attributed to the acting persona and
  tenant (E3 intact across the boundary).
- **Bar 4 — the multi-tenant human is expressible and loud.**
  Protocol: one external identity admitted into both A and B. Pass:
  the token lane admits the same human into two tenants with two
  tokens, cross-tenant signing refused; the OIDC lane's behavior for a
  subject holding roles in two tenants is deterministic and named
  (an account-selection mechanism, or the D24 ambiguity refusal kept
  deliberately — either way measured, never order-dependent); and the
  chosen persona-name scope (account-scoped vs deployment-unique) is
  demonstrated: the same persona name in two tenants either signs in
  both, or the second refuses loudly naming the first owner — no
  silent shadowing (D26's first-owner-wins re-measured at this scope).

## Reversal condition

If the platform account forces what 0064 named as the ejection
evidence — a core wire change beyond additive vocabulary, a mutual
import across the cycle guard, or a privileged identity tier — or if
the export path cannot preserve the D15 proof (observable: a rig run
in which a tenant-A caller is admitted to a subject carrying tenant
B's account token, or in which reaching the exported surface requires
wildcard publish grants), then the platform-services account does not
belong in this architecture: services remain per-tenant-account, the
D14 amendment's prefix answer stands, and the topic ends `abandoned`
with the run recorded.

## Verdict

*All four bars measured 2026-08-26 (rigs preserved in `rigs/`, full
detail in [`JOURNEY.md`](JOURNEY.md)). The topic's own question —
does the platform-services-account topology fit inside the core
invariants — is answered **YES**: the reversal condition never fired
(no wire change beyond additive vocabulary, no mutual import across the
cycle guard, no privileged identity tier; the cross-account export path
preserves the D15 proof). The topology is **sound and belongs in the
architecture**. It is not free: the bars surfaced two required fixes,
one required design discipline, and two operator decisions, all
localised to the identity plane — none touching the wire or a core
invariant.*

**Bar 1 — cross-account export preserves the principal proof: PASS**
[measured, 3 runs; mechanism-argument from nats-server's own
`TestJWTAccountProtectedImport`]. `account_token_position` enforces at
import-definition time: a tenant can only import its own account key at
the token position, so it cannot construct a route to another tenant's
surface. Server-refused (no-responder), zero service-side decisions, a
negative control confirming the position is the load-bearing mechanism.
The D15 proof extends across the account boundary by export
configuration alone — the topology's central seam is real.

**Bar 2 — one-act tenant birth and admission: CONDITIONAL PASS**
[measured, 5 runs]. Birth is one fast act (395–776µs; 0110's 1.69ms
full-engine number corroborated) and the combined-fix admission works
end to end (callout round trip 2–3.5ms, well under the 5s bar). BUT the
as-built authority cannot birth a *usable* tenant — two measured,
required fixes: **(1)** `accounts.create` never adds the new tenant to
AUTH `allowed_accounts`, so the callout refuses admission
(`Authorization Violation`); **(2)** it installs the tenant signing key
*plain* while every mint issues `SetScoped(true)` users, so the minted
user inherits 0-subscription/0-payload sentinel limits — **admitted but
inert**. Both fixes measured sufficient (install the signing key as a
scoped signer with the persona template; add the tenant to
`allowed_accounts`). The product wires none of this today (`SystemConn`
absent, no client/CLI surface for `accounts.*`).

**Bar 3 — isolation holds through shared services: PASS** [measured, 3
runs, on real provisioned `SOULSTREAM` streams]. The only shared-service
model that composes with account isolation is per-tenant connections;
under it, a tenant-A principal cannot read (P1) or write (P2) tenant B's
account data, cannot address the service's B-connection (P3), and a
client-supplied tenant claim is ignored — the connection decides, never
the payload (P5, D15 for shared services); every service act lands in
exactly one tenant's stream (P4). Two disciplines the design must
record: per-tenant account-scoped connections, and tenant-from-connection
never-from-payload.

**Bar 4 — the multi-tenant human: PASS, with a finding and a decision**
[measured, 3 runs; C2/C3 reproduce the identity algorithms verbatim].
Token lane (C1): the same human holds two tokens into two tenants,
admitted to both and server-isolated — 0064-S2 is a non-issue for the
token lane. OIDC lane (C2): D24's ambiguity refusal is deterministic and
order-independent (a length count), so an OIDC human assigned into two
tenants **cannot connect** — the **operator decision**: add an explicit
tenant-selection input to the OIDC lane, or keep the deliberate refusal.
Persona-name scope (C3): the vault keys persona keys as `persona/<user>`
with no account component, so a shared platform vault **silently shadows**
same-named personas across tenants (fails the bar's no-silent-shadowing
clause) — a **required fix**: make persona-key names account-scoped
(per-tenant vault buckets, or a `persona/<account>/<user>` D26
clean-break rename; the control proves the latter).

### What graduation-to-design must carry

Recommended outcome: **`--to design`** — the topology is viable and the
design should capture the seam and the fixes. The design cycle must
resolve, with the operator, the decisions the bars reserved:

1. **Persona-name scope** (Bar 4 C3, required): per-tenant vault vs.
   account-scoped persona-key rename. Natural companion to the S1
   realm→account rename.
2. **OIDC multi-tenant humans** (Bar 4 C2, decision): tenant-selection
   mechanism vs. keep the two-tenant refusal.
3. **Where the AUTH `allowed_accounts` coupling is performed** (Bar 2,
   required): folded into `accounts.create`, or a stated separate act.
4. **Tenant signing key as a scoped signer** (Bar 2, required): the
   authority must install the persona template, not a plain key.
5. **The export seam** (Bar 1): `account_token_position = P+2`, adopting
   D14/D15's already-designed extension path as the platform surface's
   deployment shape.
6. **Shared-service disciplines** (Bar 3): per-tenant connections;
   tenant derived from the server-proven connection, never a payload.

Items 1–2 are genuine operator direction calls (the decision test:
options remain, so they wait for the human); 3–6 are mechanical and
measured-sufficient. *Resolved 2026-08-27 by the operator:* **(1)
per-tenant vault buckets** — persona keys are not used outside their
tenant, so each tenant's persona custody lives on its own JetStream,
reached through the service's per-tenant connection; the platform vault
keeps platform custody only. **(2) keep the deliberate OIDC refusal** —
constitution III; the reversal condition (a real blocked multi-tenant
OIDC human) is named in the design. Graduated `--to design`:
[`../../02-DESIGN/soulstream-identity/platform-topology.md`](../../02-DESIGN/soulstream-identity/platform-topology.md)
(D46–D49), episode 0133.
