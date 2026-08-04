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

*Empty until graduation.*
