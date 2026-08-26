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

<Empty until graduation. Filled by /research-graduate: PASS/FAIL per bar with the
honest numbers, each load-bearing claim tagged [measured] / [mechanism-argument]
/ [judgment].>
