# Episode 0064 — The platform turn: tenancy, guardrails, and eight decisions (2026-08-03 → 2026-08-04)

A comparative analysis against a separate proprietary platform — built by the
same author, for the same problem space — asked whether the ecosystem should
absorb it. The answer was narrower and more useful than the question: the two
are the same architecture built for opposite halves of the problem. That one
has tenancy, guardrails, a model plane, an audit derivation, and a Kubernetes
deployment story; this one has the collaboration substrate, the workload
runtime, and the single-binary deployment that one has an *unresolved research
effort* about `[judgment]`. Exactly two of its capabilities are worth having —
**tenancy** and **guardrails** — and both are now enumerated as requirements in
research topic
[`platform-tenancy-guardrails`](0107-ecosystem-platform-tenancy-guardrails.md).

**Nothing is ported.** The other system is proprietary; this one is MIT. Only
ideas cross, and the provenance rule written at the head of the topic's
requirements is stricter than "no code" — design-document prose, tables,
identifiers, and vocabulary are copyrightable expression too. Every requirement
is therefore stated as *what must be true*, with the prior answer deliberately
withheld, so the design is derived independently `[judgment]`.

**Eight decisions were taken**, recorded with their reasoning so a later change
argues against something real:

- **S1** — a realm *is* an account; the vocabulary should say so. Discovered
  mid-analysis: the protocol already defines it that way, so the expected
  architectural work is a rename `[mechanism-argument]`.
- **S2** — a principal belongs to exactly one account `[judgment]`.
- **S3** — inference stays out; an agent configures its provider directly. The
  recorded cost: model spend becomes the one axis the tenancy boundary does not
  cover `[judgment]`.
- **S4** — agents and assistants are one concept. A conversational assistant
  and an autonomous agent are both personas running as workloads under the
  existing declaration; the difference is only what the program does when it
  wakes, and the wake-on-mention path already exists. No new type, no new
  field, no new protocol `[mechanism-argument]`.
- **S5** — no persona classification field. Re-derived and re-refused for the
  third time (persona kind, agent-versus-assistant, the persona passport),
  each time on the ground stated in 014: the protocol cannot verify what sort
  of entity controls a key, so it refuses to record the claim
  `[mechanism-argument]`.
- **S6** — delegated authority, never borrowed identity. An agent acting for a
  human presents its own identity plus a grant, never the human's identity.
  Attribution stays honest, the grant is independently revocable, every use is
  attributable to both parties `[judgment]`.
- **S7** — the persona registry stays in soulstream. Decisive reason: signature
  verification reads the author's published signing key, so moving it would
  give that path a dependency on soulidentity — exactly what the cycle guard
  ([0027](0027-soulstream-dx-hardening-and-the-cycle-guard.md)) prevents
  `[mechanism-argument]`.
- **S8** — grants split: enforcement with the key custodian, the record of
  issue and revoke in the op-log, wired by consumers along the signer-seam
  pattern rather than by an import `[mechanism-argument]`.

**What was refuted.** Three expectations died on contact with the design docs.
The "make realms into accounts" work is a rename, not an architecture change —
the protocol says so normatively. The persona "passport" was thought to need
building; the registry already carries `operated_by` with countersigned
attestation, chains terminating at a principal, key distribution, and service
advertisement, leaving only status and richer capabilities. And the intuition
that soulidentity manages personas was simply stale — the identity ledger was
dissolved at [0029](0029-soulidentity-the-registry-dissolves.md), and personas
are born at first encounter with nothing pre-provisioned
`[mechanism-argument]`.

**What it opened.** Two one-way doors now have deadlines rather than dates.
Whether the signed record holds an account's name or its key (A10) blocks the
rename, because the realm name is bound into the canonical form. Whether the
record notes the *acting credential* (E3) must be decided before assistants
ship: today nothing distinguishes what a human did from what their assistant
did under the same name. Recording the acting key as provenance makes no
competing identity claim and so does not reopen S6 — but it adds a field to the
signed record, which is cheap now and expensive later `[judgment]`.

One unconfirmed correctness gap also surfaced: a persona's signing key
materializes in the vault on first use and nothing publishes the public half to
the registry, so a persona could sign records nobody can verify. Inferred from
the design docs, **not traced through running code** — confirm before acting.

Reversal condition: if tenancy or guardrails cannot be added without changing a
core invariant — observable as a required wire-format change beyond additive
vocabulary, a required import between the two core modules, or a required
privileged tier in the identity model — then the platform ambition does not
belong in the core, the topic ends `abandoned`, and the capability lives in a
layer above instead.

Trail: [`01-RESEARCH/platform-tenancy-guardrails/`](0107-ecosystem-platform-tenancy-guardrails.md)
(README, requirements, JOURNEY); no code changed.
