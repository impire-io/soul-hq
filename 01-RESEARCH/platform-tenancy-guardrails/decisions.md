# Decision memos — drafted for graduation, DECIDED 2026-08-18

*Drafted 2026-08-18, then **taken by the operator the same day** in a
remote teach-back review, question by question. Two outcomes differ
from the drafts below and supersede them: **A10 → the account KEY in
the canonical form** (the operator invoked the pre-v1 clean-break rule,
discounting the draft's compat argument; the key scopes signatures to
the true trust root and aligns core with the identity plane's
account=key vocabulary; migration is the standard refuse-by-name with
the re-founding step) and **E3 → the acting-credential field is
REQUIRED** (custodian-verified on the sign.record lane, testimony-grade
on the self-custody lane — the same evidence-class split the signing
story already has). All other memos were adopted as drafted. The
verdict in README.md is the record; the design docs carry the outcomes.*

---

## A7 — Where account-creating authority lives

**Recommendation: a new narrow op family on the identity plane
(`accounts.create/suspend/resume/resolve`), operator-gated by the D25
transport ACL, with the authority backend pluggable.** [judgment]

Creating an account is a key-custody act — locally it means the
operator signing key signs an account JWT and pushes it to the
resolver, and that key belongs in the vault under the same Article I
custody as everything else. A separate account service would need the
same vault access: two services over one secret is a seam without a
boundary. But the surface must be *named and narrow* — `accounts.*` op
tails the ACL can gate independently — never folded into an existing
op. The backend seam exists because of A8: the same surface fronts
either the local operator key or a provider's control-plane API.

Would change it: a deployment class where the account authority cannot
share a process with the custodian (observable: a provider whose
account API demands credentials the vault must not hold).

## A8 — Provider-custodied root keys: supported, or not

**Recommendation: both arms, behind A7's backend seam — and A8
dissolves into Bar 1's run matrix rather than leaving as a successor
topic.** [judgment; each arm's feasibility already evidenced]

The evidence exists on both arms already: the BYO work drove Synadia
Cloud BYON's account half through the control-plane API — accounts
created by API against a root we never hold (episodes 0095/0096,
founded live in 0099) — and the self-hosted product signs accounts
with a local operator key at founding. What A8 actually still asks is
*runtime* creation on the provider arm (founding-time is proven), and
that is precisely Bar 1 run under both custody arrangements, which the
README already demands. No separate topic needed unless Bar 1's
provider arm fails for a reason the API contract can't express.

Would change it: Bar 1 failing on the provider arm structurally
(observable: the control-plane API cannot create an account without a
restart or an edit to existing accounts' configuration).

## A10 — The signed record: account name, or account key (ONE-WAY DOOR — operator decides)

**Recommendation: the human-readable name stays in the canonical form;
resolution is A4's job.** [judgment]

For: (1) changing the canonical form invalidates every existing
signature — byon and the dogfood realm are live history; (2) A4's
name→key mapping must exist regardless (callers name accounts by name
at connection time), so the record gains no capability by carrying the
key; (3) the record stays human-legible and portable; (4) S1/A9's
rename then never touches the wire. Cost, named honestly: the record is
not self-contained — attribution across realms leans on the mapping,
and a reused account name aliases history. That is a resolution-layer
concern (A4 can bind name→key with first-seen semantics, the TOFU
pattern the registry already uses for personas).

Would change it: a real cross-realm verification need where the mapping
cannot travel with the record (observable: an exhibit whose account
cannot be resolved by any reachable directory).

## A11 — The soulrealm rename

**Closed by events.** The naming re-centering (episodes 0069/0070,
2026-08-13) renamed the project to `soulstream-workloads` in the
eight-repo sweep — after this topic opened. The *term* "realm" in the
wider vocabulary is A9's rename work and stands; the project-name
question this marker asked is answered. [measured: the sweep landed]

## B8 — Guardrail: advisory or unskippable

**Recommendation: unskippable at the capability chokepoints, advisory
as a library elsewhere — one evaluator, embedded at both.** [judgment]

The transport cannot express B1's check, and a proxy tier in front of
arbitrary publishes is a second control plane — the exact shape episode
0002 refused. But this architecture already has chokepoints a caller
*cannot* bypass, because that is where the capability lives (C5's own
logic): the custodian's op path (sign, mint, grants — refuse there and
the action never had authority), the waker's wake decision, the door's
tool-call forwarding. Enforce unskippably at those; first-party
services embed the same evaluator advisorily. Hosting someone else's
code then means: their workload's *capabilities* are guardrailed even
though their arbitrary publishes are not — which is the honest
statement of what a transport-permission system plus capability
chokepoints can enforce.

Consequence for Bar 3: the enforcement point is the custodian's op
path, and the latency budget must be fixed in JOURNEY.md before the
run (proposed there: added p99 ≤ 2ms on a representative allow path).

Would change it: a hosted-code requirement that a non-capability action
be unskippably checked (observable: a concrete abuse through plain
publishes the chokepoints cannot see).

## B9 — The evaluation input

**Recommendation: `{principal, action, args, time}` — the server-proven
principal, nothing caller-claimed — plus evaluator-held counters for
rate rules; no richer input until a real rule demands it.** [judgment]

The principal is free and trustworthy (D15). Frequency ("at this
moment" includes "how often lately") needs state, and that state must
be the evaluator's own, never caller-supplied. B4's one-shot approvals
need an invocation id, which is derivable from (principal, action,
args, time) by hashing — no new input. Confirm against the Bar 3 spike
before freezing.

## B10 — Grants and escalation approvals: one mechanism or two

**Recommendation: one mechanism at two lifetimes — an approval is a
one-shot grant, and D33's delegation is already its artifact.**
[judgment]

The subject-signed, bounded, presented-per-call delegation (running
code since v0.3.0) has exactly B4's required properties: bound to named
scope, expiring, verified by the enforcer from the directory, honored
only from its named actor. An escalation approval is that artifact with
the human as subject, the workload as actor, the bounds narrowed to one
invocation id, and a TTL of minutes. Building a second approval
mechanism would duplicate a verified one. The guardrail's third outcome
(B3) then *mints a delegation request*; the human's answer is a
delegation; the retry presents it.

Would change it: B4's binding proving too fine for the delegation shape
(observable: an invocation that cannot be named by resources/scopes
without leaking its arguments into the artifact).

## C7 — Why the grant store is not the store dissolved at 0029

**The argument, written for decision time:** [mechanism-argument]

The dissolved registry restated facts other artifacts already carried —
derivable, hence a second source of truth. The grants domain holds two
things, neither derivable: (1) **original secrets** (the refresh token
has no other home — custody is the point), and (2) once C4 lands, an
**enforcement projection of the op-log's grant record** — and a
projection is not a source of truth, it is the same class as work.md's
claim projections: the op-log record (grant issued, revoked, exercised
— reviewable by the granter, S8's soulstream half) is authoritative,
and the custodian's view rebuilds from it. The 0029 objection dissolves
exactly when the store can answer "replay the ops and you get my
state"; the design duty C4 carries is to keep that true.

## C8 — Grant presentable, or looked up

**Answered for the persona-to-persona case (recorded in JOURNEY.md
2026-08-18):** the delegation is a presentable subject-signed artifact
verified from the D26 directory, and actor-binding removes the bearer
leak-amplification. **For standing grants (C2/C4's consent record):
looked up by the enforcer** — the consent record lives with the
custodian as the op-log projection (C7), so presentation adds nothing
but a leak surface. One mechanism presents (short-lived, per-call); the
other is consulted (standing state). [judgment]

## E3 — Record the acting credential as provenance (ONE-WAY DOOR — operator decides)

**Recommendation: option (b) — one provenance field in the signed
record — with a sharper enforcement shape than the requirement
sketched: the custodian stamps it.** [mechanism-argument for the
shape; the door itself is the operator's]

Because signing already flows through the identity plane, `sign.record`
knows the server-proven principal of the connection presenting the
canonical bytes. If the canonical form carries an acting-credential
field, the service can **refuse to sign when the field does not name
the caller** — making the field provenance-grade (custodian-verified at
creation) rather than a self-claim. No second identity is asserted; the
author remains solely accountable (S6 intact); guardrails and audit
gain the who-held-the-pen fact E3 wants. Cheap now, expensive after
assistants ship — the requirement's own urgency note stands.

Would change it: a signing lane that legitimately cannot know the
acting principal (observable: a creds-file bypass signer where the
persona key never touches the identity plane — the self-custody lane —
which would make the field optional-when-absent, stated in the spec).

## E4 — Discovering who answers

**Recommendation: a self-declared capability field in the registry
profile (the service-advertisement precedent), presentation-only,
behavior never branching on it.** [judgment]

`responds_when_addressed: true|false|absent` (vocabulary to be bikeshed
at design time) in the profile a persona publishes about itself —
unverified by construction, honest about being unverified, and S5-safe
because nothing enforces or branches on it. The F1 fix gives this a
natural home: once every signing persona has an ensured profile, the
field has somewhere to live for exactly the personas a human would
address.

Would change it: the field being trusted by tooling (observable: any
code path branching on it — that is S5 returning, and the field should
be cut instead).
