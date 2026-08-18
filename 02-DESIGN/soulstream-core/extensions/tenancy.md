# Tenancy hooks — the record's half (grants vocabulary, profile additions, canonical-form amendments)

*Graduated from research topic `platform-tenancy-guardrails` (episode
0107). This is the soulstream-core half of the S8 split: the record of
authority lives in the op-log where its issuer can review and watch it;
enforcement lives with the key custodian
([`../../soulstream-identity/tenancy.md`](../../soulstream-identity/tenancy.md),
D35–D38). Neither module imports the other — measured intact after the
grant build (Bar 5).*

## The grant record vocabulary (C4)

Additive ops, no cryptography beyond ordinary signing:

- **`grant.issue`** — the granting persona authorizes another to do a
  specific, named thing: `{grantee, scope (resource or action class),
  expires_at}`, signed by the granter like any op. A grant exists when
  issued and not before (C6 — nothing pre-provisioned).
- **`grant.revoke`** — `{grant_id}` (the issue op's id), signed by the
  granter; rewrites no history.
- **Exercise attribution** (C3): an op performed under a grant names
  both parties — the acting persona as author, the granting persona as
  the source of authority — each readable without inferring the other.

**The projection duty (C7), stated as a design requirement**: any
enforcement-side state derived from these ops (the custodian's consent
view, D38's approval checks) MUST rebuild exactly from replaying them —
the op-log record is the source of truth; the custodian holds a
projection plus original secrets, never a second authority. This is
the pre-answered objection to episode 0029's registry dissolution.

**Presentation split (C8)**: short-lived per-call delegations are
*presented* (the running D33 artifact); standing consent is *looked
up* by the enforcer from its projection — never a bearer artifact.

**Gate criterion**: Bar 4 as pre-registered in the topic — a granted
action performed with dual attribution, and revocation that stops the
action while disturbing **neither** persona's own standing. The first
two clauses are already measured on the delegation machinery; the
third is what this vocabulary adds, and the delegation matrix is its
measurement rig.

## Registry additions

- **The ensure-signing-key act (F1)**: a persona's signing key
  materializes in the identity-plane vault on first touch, and today
  *nothing* publishes the public half — unknown-key is the shipped
  default for every identity-plane-signed persona [measured, code
  trace: no reader consults `keys.public`; readers build keyrings from
  registry profiles + TOFU pins only]. The `registry` package grows an
  ensure act — publish-or-refresh the persona's profile carrying its
  signing public key — and **every `PersonaSigner` consumer calls it
  at signer construction** (the product's runner and archivist, the
  remote node's per-user pool). Readers stay unchanged. Small, first
  in the build order.
- **The responds-when-addressed field (E4)**: profiles gain one
  self-declared, unverified field saying whether the persona answers
  when addressed. Presentation-only; **no code path may branch on it**
  — if any ever does, the field is cut rather than trusted (that is S5
  returning by another route). Lands as a `registry.md` amendment with
  its build.

## Canonical-form amendments (decided 2026-08-18 — land with their builds, as clean breaks)

Two operator-taken one-way doors amend the canonical signed form.
They are recorded here until their builds fold them into
[`../core/01-protocol.md`](../core/01-protocol.md) and the signing
spec; both are pre-v1 clean breaks — an old-shape realm is detected
and refused by name with the re-founding migration spelled out.

- **A10 — the account key replaces the realm name.** The canonical
  form's realm field becomes the account's public key: the signature
  scopes to the true trust root (a name-scoped signature could be
  replayed into a realm reusing the name), and core aligns with the
  identity plane, where an account already *is* its key. Human names
  become display-layer resolution (`accounts.resolve`, D35) — display
  can be wrong without breaking verification.
- **E3 — a required acting-credential field.** Every signed record
  names the server-proven principal of the connection that published
  it — whose hand held the pen, distinct from the author. Two evidence
  grades, the same split the signing story already has: on the
  custodial lane, `sign.record` **refuses** a canonical whose acting
  field does not name the server-proven caller (custodian-verified
  provenance); on the self-custody lane it is testimony-grade
  self-claim. No second identity is asserted; the author remains
  solely accountable (S6 intact).

**Reversal conditions** (from the decision record): A10 — a
deployment class needing record verification where no account-key
directory can travel with the exhibit reopens the display-layer
question, not the wire. E3 — a signing lane that legitimately cannot
know its acting principal beyond the self-custody case reopens the
required-ness as a new decision.
