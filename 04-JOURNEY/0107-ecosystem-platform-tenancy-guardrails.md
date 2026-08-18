# Episode 0107 — Tenancy and guardrails: the platform question answered (2026-08-04 → 2026-08-18)

The topic opened two weeks ago with an unusual confession: its
direction was already set (eight decisions, S1–S8, from the
comparative analysis of episode 0064), so its honest job was to
enumerate what the direction *requires*, measure what could be
measured, and take the rest as named judgment — under a provenance
rule stricter than "no code" that survives it. It closes today with
its question decisively answered: **tenancy and guardrails fit inside
the core invariants** — the reversal condition (a forced wire change
beyond additive vocabulary, a forced mutual import, a forced
privileged identity tier) never fired. The two wire changes it does
take are deliberate pre-v1 clean breaks, not forced ones.

**The bars, measured** (rigs per how-we-work; spreads, not means):

- **Bar 1 — an account is born at runtime: PASS on the local arm**
  [measured, 6/6]: a complete-built account JWT stored as one act went
  store → first full round trip in **543µs–774µs**, zero restarts,
  zero edits to existing accounts, the pre-existing account's
  continuous probe uninterrupted (max gap 6.01ms at 5ms cadence). The
  **provider arm** (Synadia BYON, per A8) is the one named residue —
  an operator act, its shape written.
- **Bar 2 — no usable half-account: PASS** [measured, 6/6]:
  2,418–2,826 pre-creation probes per run all failed closed, **zero
  partial successes** ever observed. Complete-artifact-then-one-store
  is the ordering that makes A2 true.
- **Bar 3 — a guardrail cannot hang the request path: PASS**
  [measured, 3 runs]: hostile CEL refused at compile or terminated
  sub-millisecond; allow-path **p99 206–220µs at 100 rules** against
  the pre-registered 2ms budget. The bar earned its keep: a cost
  limit alone let an input bomb take **622ms** to die — cost
  accounting is a step bound, not a wall-clock bound — so the
  evaluator discipline (tight cost limit + interrupt check + context
  deadline) is now a design mandate, not an option.
- **Bar 4 — a grant is real, scoped, revocable: not measured**,
  carried honestly as the C4 build's gate criterion; its first two
  clauses are measured on the delegation machinery, the third is
  exactly the consent record that build adds.
- **Bar 5 — the cycle guard survives the grant work: PASS**
  [measured]: zero cross-references in both modules' complete
  dependency graphs after identity v0.3.0 landed.
- **F1 — confirmed, and worse than suspected** [measured, code trace]:
  no reader consults `keys.public` and no `PersonaSigner` consumer
  publishes a profile — unknown-key is the *shipped default* for every
  identity-plane-signed persona, masked in dogfood by manual profile
  publication. The fix (a core registry ensure-act called at signer
  construction) is first in the build order.

**The decisions — all eleven taken by the operator 2026-08-18**, in a
remote teach-back, question by question. Two overrode the topic's own
drafts, both by invoking the pre-v1 clean-break rule against
compat-weighted recommendations: **A10 — the account key enters the
canonical signed form** (signatures scope to the true trust root, not
a reusable label; names demote to display-layer resolution)
[judgment], and **E3 — the acting-credential field is required** in
every signed record, custodian-verified on the `sign.record` lane and
testimony-grade on self-custody — the same evidence-class split the
signing story already carries [judgment; the stamping mechanism is
mechanism-argument]. The rest as drafted: `accounts.*` on the identity
plane with pluggable authority (A7/A8, dissolving A8 into Bar 1's run
matrix instead of a successor topic); A11 closed by the 0069/0070
rename; guardrails **unskippable at capability chokepoints, advisory
elsewhere** — the mediation proxy refused as episode 0002's second
control plane (B8); minimal evaluation input (B9); **approvals are
one-shot delegations** — D33's running artifact, no second mechanism
(B10/B4); the grant store answers the 0029 objection as
original-secrets-plus-op-log-projection with the replay property a
stated duty (C7); present the short-lived, look up the standing (C8);
a self-declared responds-when-addressed profile field that no code may
branch on (E4).

Refuted or corrected along the way: the topic's compat-weighted A10
draft (the operator's rule outweighed it); the expectation that §C/§D
were all-new (the grants broker built half of them mid-topic, and the
reconciliation measured what remained rather than re-planning what
existed); and the assumption that cost limits bound wall-clock time.

What it opened, named: the build items in their dependency order —
F1's ensure-act, the C4 grant vocabulary (Bar 4 its gate), the general
secret store (D36), the evaluator at the custodian chokepoint
(D37/D38), the `accounts.*` family last as the most irreversible
(D35), and the A9/A10/E3 core amendments as their own clean-break
cycle — all behind the 0071 focus gate, designed now so demand has an
answer. The Bar 1 provider arm stays an operator act.

Reversal condition: the designs carry their own (D38's
argument-leaking approvals; A10's exhibit-resolution case; E3's
unknowable-actor lane; B8's chokepoint-invisible abuse). For the topic
itself: if the C4 build cannot make Bar 4's third clause true — a
revocation that stops delegated action while disturbing neither
persona [observable: the delegation matrix failing that row] — the
one-mechanism consolidation (B10) reopens before anything ships on it.

Trail: designs
[`soulstream-identity/tenancy.md`](../02-DESIGN/soulstream-identity/tenancy.md)
(D35–D38) and
[`soulstream-core/extensions/tenancy.md`](../02-DESIGN/soulstream-core/extensions/tenancy.md);
the concluded topic (removed on graduation, full history in git —
opened 2026-08-04, F1 trace `0e1aec8`, reconciliation `30f2571`, memos
`0da7b25`, pre-registration `a12e278`, spikes `c45896f`, verdict
`070aafb`); episodes
[0064](0064-ecosystem-the-platform-turn.md) (the direction),
[0104](0104-ecosystem-outbound-identity-grants.md)/[0105](0105-identity-the-grants-broker-lands.md)
(the grants half it reconciled against),
[0027](0027-soulstream-dx-hardening-and-the-cycle-guard.md) (the guard
Bar 5 re-proved); rig modules in the session scratchpad per
how-we-work.
