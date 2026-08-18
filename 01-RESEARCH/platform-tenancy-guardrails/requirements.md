# Requirements — tenancy, guardrails, and what they need

*The enumerated output of this topic. States **what must be true**, not how to
build it. Companion to [`README.md`](README.md), which carries the question,
the bars, and the decisions that are not bars.*

## Provenance rule

**This rule governs this document and everything derived from it.** The
direction was set by comparison with a separate proprietary platform. Nothing
here reproduces text, code, identifiers, table structures, or vocabulary from
that system or any other proprietary source; requirements are stated in plain
words so the design is derived independently.

Where a mechanism has public prior art — NATS import/export and account token
positions, policy expression languages, OIDC and WebAuthn, ordinary
secret-store and optimistic-concurrency patterns — derive from the public
source directly. Where a requirement below has no public prior art, design
forward from the requirement as written; **do not consult a proprietary
implementation for its answer.** A lightly-reworded proprietary table is
copied expression as surely as a copied file.

Anyone extending this work inherits the rule.

---

## 1. Settled — decisions taken, with their reasons

Taken in discussion before the topic opened, and recorded in
[episode 0064](../../04-JOURNEY/0064-ecosystem-the-platform-turn.md) with their
reasoning, so a later change argues against something real rather than
rediscovering it.

**S1 — A realm is an account; say so.** The protocol already defines a realm as
one NATS account and as the tenancy and trust boundary
([`core/01-protocol.md`](../../02-DESIGN/soulstream-core/core/01-protocol.md)). The
vocabulary should say *account*, because that is what it is and it is legible
to anyone who knows NATS. Naming, not architecture.

**S2 — A principal belongs to exactly one account.** No cross-account
membership and no cross-account identity. Cross-account *communication* stays
out of scope.

**S3 — Inference stays out.** An agent configures its model provider directly.
Recorded cost: model spend becomes the one axis the tenancy boundary does not
cover — no central usage accounting, catalogue, egress control, or quota. This
will resurface the first time someone asks what a tenant costs. Deferred, not
rejected.

**S4 — Agents and assistants are one concept, not two.** A conversational
assistant and an autonomous agent are both personas, both run as workloads, and
both use the existing declaration unchanged — it already carries *as which
persona* and *anchored to which topic*
([`0001-soulrealm-runtime.md`](../../02-DESIGN/soulstream-workloads/0001-soulrealm-runtime.md)).
The difference is only what the program does when it wakes: one wakes on work
and acts, the other wakes on being addressed and replies — and the
wake-on-mention path already exists
([`core/02-identity.md`](../../02-DESIGN/soulstream-core/core/02-identity.md)).
Therefore: no new persona type, no new declaration field, no new protocol.

**S5 — No persona classification field.** Do not reintroduce a kind, type, or
class. It was introduced, demoted to presentation-only metadata, and then
removed outright, on the ground that the protocol cannot verify what sort of
entity controls a key and so refuses to record the claim
([`extensions/registry.md`](../../02-DESIGN/soulstream-core/extensions/registry.md)).
Everything wanted from a classification — who operates this, what it does, what
it offers — is expressible without it. This has now been re-derived three times
(persona kind, agent-versus-assistant, the persona passport); if it is ever
overturned, overturn it deliberately and once.

**S6 — Delegated authority, never borrowed identity.** A persona acting for
another does not assume its identity; it presents its own identity plus a
grant. Attribution stays honest and consistent with the standing refusal of
on-behalf-of attribution; the grant is revocable without touching either
persona; it can be scoped and time-bounded; every use is attributable to both
parties. Identity inheritance offers none of these.

**S7 — The persona registry stays in soulstream.** Signature verification needs
the author's published signing key, so moving the registry would give that
verification path a dependency on the identity service — which the cycle guard
([episode 0027](../../04-JOURNEY/0027-soulstream-dx-hardening-and-the-cycle-guard.md))
exists to prevent. A deployment running no registry at all must remain a
working deployment.

**S8 — Grants split across both components.** Enforcement belongs to the key
custodian: a grant is only real if the party holding the key refuses to act
without it. The *record* of issuing and revoking belongs in the op-log, because
a grant its issuer cannot review or watch being exercised is not meaningfully
given. Neither component may import the other, so this follows the established
seam pattern — one side defines the interface, the other implements it,
consumers wire.

---

## 2. Requirements

`[EXISTS]` already satisfied · `[NEW]` must be built · `[OPEN]` decision
outstanding (see *Decisions that are not bars* in
[`README.md`](README.md)).

### A. Accounts and tenancy

The largest gap. The account *boundary* exists; account *lifecycle* does not
exist at all.

- **A1 `[NEW]`** An operator must be able to bring a new account into existence
  at runtime, without editing static configuration and without restarting
  anything.
- **A2 `[NEW]`** An account must not be reachable or discoverable before it is
  fully constituted. Partial creation must not leave a usable half-account:
  either the whole account works, or nothing referring to it resolves. *(The
  ordering that achieves this is a design question; the requirement is only
  that no intermediate state is usable.)*
- **A3 `[NEW]`** An operator must be able to suspend an account so that no
  principal in it can connect or act, without destroying its data, and to
  reverse that.
- **A4 `[NEW]`** Something must map a human-meaningful account name to the
  account's cryptographic identity. Callers name accounts by name; the
  substrate identifies them by key. This mapping is also what A2's ordering
  hangs on.
- **A5 `[NEW]`** Creating account N+1 must require no change to anything that
  already exists — no static manifest edit, no redeploy of a running
  component, no downtime for existing accounts.
- **A6 `[EXISTS]`** Cross-account exposure of a shared surface is already
  designed: a service exported with the account token at a declared position,
  with the substrate forcing each importing account's own key into that
  position ([`nats-surface.md`](../../02-DESIGN/soulstream-identity/nats-surface.md),
  D14). A1–A5 build on this; it does not need reinventing.
- **A7 `[OPEN]`** Where does account-creating authority live? The custodian
  holds keys and is the natural enforcement point, but account creation may
  deserve its own narrow surface rather than being folded in. Decide before
  building A1.
- **A8 `[OPEN]`** Must account creation work where a hosting provider custodies
  the root signing key and exposes only an API, or only where we hold it? These
  have materially different trust properties — the first keeps the deployment
  portable, the second does not. Shapes the whole seam.

**Rename work (S1), separable from the above:**

- **A9 `[NEW]`** Rename the concept throughout: library surface, CLI flags,
  environment variables, config files, state directory layout, documentation.
  Roughly 90 non-test Go files mention it (a mention count, not an assessment
  of how many must change). Mechanical except for A10.
- **A10 `[OPEN]`** The realm name is bound into the canonical signed form of
  every operation, so this is not purely cosmetic at the wire. Decide whether
  that field continues to hold a human-readable name — signatures stay stable,
  resolution defers to A4 — or becomes the account's cryptographic identifier,
  which makes the record self-contained but changes the canonical form and
  invalidates existing signatures. Recommend the former; decide explicitly.
- **A11 `[OPEN]`** `soulrealm`'s name derives from the old term and will read
  oddly afterwards. Rename, or accept the drift, deliberately.

### B. Guardrails

Nothing here exists. From-scratch design, derived from these requirements plus
public prior art only.

- **B1 `[NEW]`** Some operations must be checked before they proceed, on
  grounds the transport cannot express. The substrate answers *may this
  principal reach this subject*; it cannot answer *may this specific
  invocation, with these arguments, at this moment, proceed*. The guardrail
  answers only the second.
- **B2 `[NEW]`** A guardrail must never re-answer the first question.
  Re-deriving "is this principal in this account" in application code
  duplicates enforcement the substrate already performed, and is a defect
  rather than defence in depth — the constitution's position that the server is
  the verifier of record.
- **B3 `[NEW]`** An evaluation must be able to conclude in more than two ways.
  Beyond proceeding and refusing, it must be able to conclude that the action
  may proceed *only after a human decides*. Without that, every uncertain case
  is pre-resolved as a blanket allow or a blanket deny, and operators choose
  allow.
- **B4 `[NEW]`** When an evaluation defers to a human, the resulting approval
  must be bound to the specific invocation it approved, usable exactly once,
  and unable to authorize anything else. An approval that generalises is a
  privilege escalation with extra steps.
- **B5 `[NEW]`** Every evaluation must be observable, including the ones that
  permit. A guardrail whose allows are invisible cannot be audited, tuned, or
  trusted.
- **B6 `[NEW]`** Rules must be data, changeable without redeploying what they
  govern, and every enforcement point must converge on a change without
  restarting.
- **B7 `[NEW]`** The rule language must be sandboxed, terminating, and free of
  side effects. Evaluation sits on the request path: a rule must not be able to
  hang it and must not be able to reach anything.
- **B8 `[OPEN]`** Advisory (a library each service chooses to call) or
  unskippable (enforced where a service cannot bypass it)? Advisory is far
  simpler and adequate for first-party code; unskippable is what hosting
  someone else's code eventually requires. Decide which is being built, and
  whether the design must leave room for the other.
- **B9 `[OPEN]`** What is the input to an evaluation? At minimum the acting
  principal, the operation, its arguments, and the time. The principal is
  available cheaply and trustworthily because the server proves it
  ([`nats-surface.md`](../../02-DESIGN/soulstream-identity/nats-surface.md), D15).
  Confirm nothing else is needed before designing a richer input.
- **B10 `[OPEN]`** Where do guardrails sit relative to grants (§C)? A grant is
  a standing narrow authorization; an escalation approval is a one-shot
  authorization. Plausibly the same mechanism at two lifetimes — deciding now
  avoids building two overlapping systems.

### C. Grants

The mechanism S6 depends on. Cheap to start: the recording half is additive
vocabulary with no cryptography.

- **C1 `[NEW]`** A persona must be able to authorize another to do a specific,
  named thing on its behalf — without the second being able to act *as* the
  first.
- **C2 `[NEW]`** A grant must be scoped to a resource or class of action,
  time-bounded, and revocable independently of both personas.
- **C3 `[NEW]`** Exercising a grant must be attributable to both parties: the
  acting persona under its own name, and the granting persona as the source of
  the authority. Neither may be inferable only by reading the other.
- **C4 `[NEW]`** Issuing and revoking must be recorded where the issuer can
  review it and watch it being used. A grant whose exercise is invisible to the
  granter is not meaningfully consented to.
- **C5 `[NEW]`** Enforcement sits with the party holding the key or secret. A
  grant checked only by the party that benefits from it is a suggestion.
- **C6 `[NEW]`** Nothing pre-provisioned. A grant exists when issued and not
  before — consistent with D26's rule that a persona's own artifacts
  materialize on first use rather than through a provisioning act
  ([`agent.md`](../../02-DESIGN/soulstream-identity/agent.md)).
- **C7 `[OPEN]`** A grant store is a store, and a store was dissolved at
  [episode 0029](../../04-JOURNEY/0029-soulidentity-the-registry-dissolves.md)
  for being a second source of truth restating facts other artifacts carried.
  The grant store must be able to say why it is not that. The argument
  available: a grant is an *original* fact with no other home — nothing else
  records that this persona authorized that one. Write it down at decision
  time; it is the first objection anyone familiar with that removal will raise.
- **C8 `[OPEN]`** Must a grant be presentable to a third party as a bearer
  artifact, or is it always looked up by the enforcing party? The first enables
  offline verification and cross-service use; the second is far simpler and
  much harder to leak.

### D. Secret storage in the custodian

The custodian holds keys and answers signing and minting requests. Grants (§C)
and per-persona upstream credentials both need it to hold *secrets*, not only
keys.

- **D1 `[NEW]`** Store, retrieve, list, and delete a secret at a named path.
- **D2 `[NEW]`** Concurrent writers must not silently lose each other's writes;
  a caller must be able to condition a write on the version it read.
- **D3 `[NEW]`** Paths must be namespaced so per-persona, per-account, and
  per-service secrets cannot collide, and no caller can name a path outside its
  own reach. The scheme is a design decision; the requirement is that reach is
  **structural**, not checked after the fact.
- **D4 `[NEW]`** Every path resolves within the calling account's own tree. The
  same path name in two accounts names two different secrets, with no
  configuration required to make that true.
- **D5 `[NEW]`** Stored values must be encrypted at rest under a key the
  storage backend does not hold.
- **D6 `[NEW]`** Values must be sealed end to end between caller and service,
  so the transport never carries them in the clear.
- **D7 `[EXISTS]`** The sealing machinery is already in use for the existing
  surface — envelope, service key, and pinning discipline need extending to
  new operations, not reinventing.
- **D8 `[NEW]`** A caller must be able to have the service act *with* a secret
  without receiving it. Signing already works this way; the same property must
  hold for new material, or the storage half undoes the custody guarantee that
  makes the service worth having.
- **D9 `[OPEN]`** Does one service hold both key custody and general secrets,
  or do they sit side by side behind one contract? One is simpler and already
  holds the hardest material; two keeps a narrow, auditable key surface from
  growing a large general-purpose one. Decide before D1.

### E. Agents and assistants

Per S4 and S5, almost nothing is required — recorded so the absence is
deliberate rather than an oversight.

- **E1 `[EXISTS]`** Running a conversational persona needs no new machinery:
  the workload declaration, per-persona notify subjects, and mention
  notifications already express it.
- **E2 `[NEW]`** Name the pattern in the design and make issuing a narrow,
  short-lived, hard-scoped credential ergonomic. The minting primitive exists
  (a scoped, time-bounded credential for the caller's own key, with no
  exportable credential file — D28); the *pattern* is undocumented, so everyone
  will reinvent it differently.
- **E3 `[OPEN]`** **Decide before assistants ship — it is a wire change.** If
  an assistant publishes under the human's name, nothing afterwards
  distinguishes what the human did from what their assistant did: the record
  carries an author but not the acting credential. Options: (a) accept
  indistinguishability — simplest, honest to "the assistant is me", but
  forensically empty and awkward for guardrails and audit; (b) record the
  acting credential's public key as **provenance** — no competing identity
  claim, the author remains solely accountable, the same category as
  distinguishing who committed from who authored. (b) does not reopen S6,
  because it never asserts a second party. Recommend (b); it adds a field to
  the signed record, so it is cheap now and expensive later.
- **E4 `[OPEN]`** How does a human discover which personas answer when
  addressed, versus which run unattended and will silently ignore them? A real
  need with no home. Any answer must not become S5 by another route: a
  capability a persona advertises about itself, self-declared and unverified,
  with behaviour never branching on it — not a taxonomy of what personas are.
  The service-advertisement field in profiles is the nearest precedent.

### F. Correctness gap found along the way

- **F1 `[NEW]`** A persona's signing key materializes in the vault on first use
  (D26), and nothing publishes the public half to the registry — so a persona
  can begin signing records no one can verify, until someone manually publishes
  a profile. Something must close that loop. Under S7's no-mutual-import rule
  this is a consumer's wiring job, which means it needs an owner or it stays
  unowned. Small, self-contained, independent of everything else here.
  **Confirmed 2026-08-18** [measured: code trace, JOURNEY.md entry]: every
  reader keyring is registry-profiles + TOFU pins (`realmKeyring`,
  `keyringFor`); `keys.public` is consulted by no reader; no
  `PersonaSigner` consumer (product runner/archivist, the remote node's
  per-user pool) publishes a profile — the gap is the shipped default for
  every identity-plane-signed persona, masked in dogfood by manual profile
  publication. Owner: the `PersonaSigner` consumers, via one core
  `registry` ensure-act.

---

## 3. Explicitly out of scope

Recorded so they are not rediscovered as gaps:

- Model access and inference (S3), and with it usage accounting, model
  catalogue, egress control, and per-tenant quota.
- An identity provider of our own beyond what the fold already is. Deployments
  bring their own.
- Cross-account communication and federation.
- Presence and liveness. Wanted eventually (E4 is adjacent), but note it does
  not belong in the durable profile store: status churns, profiles do not, and
  a store that keeps history would record a revision every time a laptop opens.

## 4. Suggested order

By dependency, and by how much each teaches before it commits anything:

1. **F1** — small, isolated, fixes a real correctness gap, commits nothing.
2. **A9/A10** — the rename, while the surface is smallest. A10 decided first.
3. **C1–C4** (recording half only) — additive vocabulary, no cryptography, no
   enforcement contract yet. Reveals the shape of a grant in real use.
4. **D** — secret storage, which C5 and per-persona upstream credentials both
   need. D9 decided first.
5. **C5–C6** — enforcement, once D exists and C's shape is known from use.
6. **B** — guardrails, with B10 resolved against whatever C became.
7. **A1–A8** — account lifecycle. Last not because it matters least but because
   it is the most irreversible, and everything above teaches something about
   what an account must own.

**E3 is the exception** — a wire-format decision that grows
disproportionately more expensive after assistants ship. Decide early even
though the work lands late.
