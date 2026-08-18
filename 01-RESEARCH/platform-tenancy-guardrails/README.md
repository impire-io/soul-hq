# What must be true for the ecosystem to host tenants and code it does not own?

**Component:** ecosystem
**State:** active
**Started:** 2026-08-04

## Abstract

The ecosystem today runs *our* realms and *our* workloads well. Hosting other
people's tenants and other people's code needs two capabilities it does not
have at all: **tenancy** (accounts brought into existence, suspended, and
resolved at runtime) and **guardrails** (a check on an invocation that the
transport's subject permissions cannot express). This topic establishes what
must be true for both, and the smaller enablers they turn out to need — a
grant model, secret storage in the custodian, and one correctness gap found
along the way. A decisive answer unlocks the roadmap's next phase; without it,
every "can we host X" question is answered by guesswork.

The direction was set by a comparative analysis against a separate, proprietary
platform built by the same author for the same problem space. **Nothing from
that system is reproduced here** — see the provenance rule in
[`requirements.md`](requirements.md), which governs this topic and everything
derived from it.

## The question

**What must be true for the ecosystem to admit tenants and workloads it does
not own — and what does that require the components to grow?**

The requirements are enumerated in [`requirements.md`](requirements.md), with
each marked as already satisfied, to be built, or open. Sub-questions that
prove to need their own investigation become successor topics rather than
growing this one (the one-question rule in
[`../README.md`](../README.md)); the candidates are named under *Decisions that
are not bars* below.

## Pre-registered bars

Written before any experiment. These are the claims this topic can settle by
**measurement**; the design decisions it must also take are listed separately
below, honestly, rather than dressed as bars.

- **Bar 1 — an account is born at runtime.** Against a live deployment,
  create a new account and have a principal in it connect and reach the shared
  service surface. **Pass:** the new principal connects and completes a
  round trip; **zero** restarts of any running component; **zero** edits to
  any existing account's configuration; accounts that already existed serve
  continuously throughout (measured by an uninterrupted round-trip probe on a
  pre-existing account for the duration).
- **Bar 2 — no usable half-account.** Probe the account under creation
  continuously through the creation window (resolve the name, connect a
  principal, reach the shared surface). **Pass:** every probe before
  completion fails closed; **no** probe partially succeeds. A single
  partially-successful probe fails this bar.
- **Bar 3 — a guardrail cannot hang the request path.** Evaluate a
  deliberately hostile rule (non-terminating construct, pathological input
  size) at an enforcement point. **Pass:** the rule is refused at load or
  terminated by the evaluator, never the caller; and added p99 latency on a
  representative allow-path stays within a budget stated here before the run
  (to be fixed when the enforcement point is chosen — recorded in
  `JOURNEY.md`, not left implicit).
- **Bar 4 — a grant is real, scoped, and revocable.** A persona holding a
  grant performs the granted action; the record attributes **both** the acting
  persona and the granting persona, each readable without inferring it from
  the other. Revoking the grant stops the action. **Pass:** all three hold,
  and revocation disturbs neither persona's ability to act on its own behalf.
- **Bar 5 — the cycle guard survives the grant work.** After the grant model
  lands on both sides, neither core module's dependency graph contains the
  other. **Pass:** the module graph shows zero, as it did at
  [episode 0027](../../04-JOURNEY/0027-soulstream-dx-hardening-and-the-cycle-guard.md).

Bar 1 is measured under whichever root-key custody arrangements the answer to
**A8** admits — if provider-custodied roots are in scope, Bar 1 must pass under
both, and that is stated here rather than discovered later.

## Decisions that are not bars

These are the load-bearing calls this topic must take that **no experiment can
settle** — they resolve by design argument, and will carry
`[mechanism-argument]` or `[judgment]`, never `[measured]`. Recorded here so
the topic is honest about which of its outputs are evidence and which are
choices. Each is specified in [`requirements.md`](requirements.md) under the
marker given.

| Marker | Decision | Character |
|---|---|---|
| A10 | Signed record holds the account's name or its key | one-way door (wire format) |
| E3 | Whether the record notes the acting credential | one-way door (wire format) |
| D9 | One service for key custody and secrets, or two | structural |
| C7 | Why a grant store is not the store dissolved at [0029](../../04-JOURNEY/0029-soulidentity-the-registry-dissolves.md) | argument |
| C8 | Grants presentable to a third party, or looked up by the enforcer | structural |
| B8 | Guardrail advisory, or unskippable | structural |
| B10 | Grants and escalation approvals: one mechanism or two | structural |
| A7 | Where account-creating authority lives | structural |
| A8 | Whether provider-custodied root keys must be supported | direction |
| A11 | Whether soulrealm is renamed with the vocabulary | cosmetic |

Any of these that grows an investigation of its own — A8 most plausibly, since
it *is* partly measurable — leaves as a successor topic rather than expanding
this one.

## Reversal condition

Written now: **if tenancy or guardrails cannot be added without changing a core
invariant** — observable as a required change to the wire format beyond
additive vocabulary, a required import between the two core modules, or a
required privileged tier in the identity model — then the platform ambition
does not belong in the core, and this topic ends as `abandoned` with the
capability living in a layer above rather than inside
[the constitution's](../../00-GENESIS/constitution.md) boundaries.

A second, narrower reading: if the guardrail's third outcome (defer to a human)
goes unused in real deployments — observable as operators writing only allow
and deny rules over a sustained run — then B3 was speculation and the escalation
machinery should be cut rather than maintained.

## Verdict

Recorded 2026-08-18. The question is answered: tenancy and guardrails
can be added without changing a core invariant — the reversal condition
never fired (no mutual import, no privileged identity tier; the two
wire changes are clean breaks the operator's pre-v1 rule sanctions,
taken deliberately below).

**Bar 1 — PASS on the local arm** [measured, 6/6 runs]: an account
whose JWT is built complete and stored as one act went store → first
full round trip in **543µs–774µs**; zero restarts, zero edits to the
pre-existing account, whose continuous probe (230–238 round trips per
run, 5ms cadence) recorded zero failures, max inter-success gap 6.01ms.
**The provider arm (Synadia BYON, A8) is the one named residue** — the
operator act's shape is written; graduating with it named follows
0104's precedent.

**Bar 2 — PASS** [measured, 6/6 runs]: 2,418–2,826 pre-creation probes
per run all failed closed, **zero partial successes**; post-creation
the first probe was already a full success in every run. The A2
ordering (complete artifact, then the one store) leaves no observable
intermediate state.

**Bar 3 — PASS** [measured, 3 runs]: unparseable and type-broken rules
refused at compile; a nested-comprehension cost bomb terminated in
689µs–982µs and a 100k²-element input bomb in 484µs–913µs; the
backtracking regex probe ran linear (14µs, RE2). Allow path at 100
compiled CEL rules per op over 10k ops: p50 58–69µs, **p99 206–220µs**
against the pre-registered 2ms budget. The scare that became a design
output: a cost limit alone let the input bomb take **622ms** to die —
the evaluator discipline is cost limit + interrupt check + context
deadline, belt-and-braces, never one mechanism.

**Bar 4 — not measured; carried as the C4 build's gate criterion.**
Its first two clauses (granted action performed, dual attribution)
are measured on the delegation machinery [identity v0.3.0]; the third
— revocation disturbing neither persona — is exactly the unbuilt
standing consent record, and the delegation matrix is its rig.

**Bar 5 — PASS** [measured]: after the full grant build, both core
modules' complete dependency graphs contain zero references to each
other; the consumer-position e2e imports both — the sanctioned shape.

**F1 — confirmed** [measured, code trace]: no reader consults
`keys.public`, no `PersonaSigner` consumer publishes a profile —
unknown-key is the shipped default for every identity-plane-signed
persona. Owner: the signer consumers, via one core registry ensure-act.

**Decisions — all eleven taken by the operator, 2026-08-18 (remote
review, teach-back per question):** A10 **account key in the canonical
form** [judgment — overriding this topic's own draft recommendation;
the pre-v1 clean-break rule discounts the compat argument, and the key
scopes signatures to the true trust root]; E3 **required
acting-credential field, two evidence grades** [judgment; the
custodian-stamped half is mechanism-argument]; A7 `accounts.*` on the
identity plane, backend-pluggable; A8 both custody arms, dissolved
into Bar 1's matrix; A11 closed by events (the 0069/0070 rename); B8
unskippable at capability chokepoints, advisory elsewhere; B9 minimal
evaluation input + evaluator-held counters; B10 one mechanism —
approvals are one-shot delegations; C7 the projection argument; C8
present short-lived, look up standing; E4 self-declared profile field
[all judgment/mechanism-argument, none dressed as measured].

Outcome: **design** — `02-DESIGN/soulstream-identity/tenancy.md`
(D35–D38) and `02-DESIGN/soulstream-core/extensions/tenancy.md`.
