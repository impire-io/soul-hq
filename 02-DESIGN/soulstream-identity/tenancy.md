# soulstream-identity — tenancy, secrets, and the guardrail evaluator (D35–D38)

*Graduated from research topic `platform-tenancy-guardrails` (episode
0107; Bars 1–3 and 5 measured, eleven decisions taken by the operator
2026-08-18). The record-side halves — the grant-record vocabulary, the
profile field, the ensure-signing-key act, and the two canonical-form
amendments — live in
[`../soulstream-core/extensions/tenancy.md`](../soulstream-core/extensions/tenancy.md);
neither module imports the other (Bar 5 held through the grant build).
Decisions continue the global numbering: D35–D38.*

## D35 — Account lifecycle is an op family on the identity plane, with pluggable authority

Tenancy's missing half is lifecycle: the account *boundary* exists (a
realm is one NATS account), account *birth, suspension, and resolution
at runtime* do not. They become a narrow, named op family on the
management surface — `accounts.create`, `accounts.suspend`,
`accounts.resume`, `accounts.resolve` — gated exactly as every
management op is: by the deployment's permission template on the op
tail (D25), reachable only by operator-class credentials. Account
creation is a key-custody act, and the custodian already holds the
keys; a separate account service would be two services over one secret
(operator decision, A7).

**The authority backend is a seam** (operator decision, A8): the
`LocalOperator` backend signs account JWTs with the operator signing
key custodied in the vault; the `ProviderAPI` backend drives a hosting
provider's control plane (Synadia Cloud BYON — the account-by-API half
proven at the byon founding). Both arms must pass the same acceptance
criteria.

**The ordering rule, measured**: the account artifact is built
**complete** — limits, imports, everything — and only then pushed as
one act. Measured on the local arm [6/6 runs]: store → first full
round trip in 543µs–774µs, zero restarts, zero edits to existing
accounts, a pre-existing account's continuous probe uninterrupted; and
across ~2,400–2,800 probes per run through the creation window, zero
partial successes — no usable half-account ever exists (Bars 1 and 2).

`accounts.resolve` is the name→key mapping (A4) — **display-layer, not
security-layer**, because the canonical record carries the account key
itself (A10, the core amendment): verification never depends on
resolution. Names bind first-seen; a name reuse refuses.

Suspension: no principal in a suspended account can connect or act;
the account's data is untouched; resume reverses it. The mechanism
(resolver update, revocation) is the build's choice; the requirement
is the observable.

**Acceptance criteria**: (1) Bars 1 and 2 re-measured through the real
op family on both authority arms — the provider arm is the topic's one
named residue; (2) suspend refuses the next connection within one
callout TTL and resume restores; (3) `accounts.*` unreachable for
represented users (transport-refused, delivery-log proof).

## D36 — The custodian stores secrets, one service, domains side by side

The grants build answered D9's question by composition: a second
sealed custody domain inside the one service works cleanly (D31 —
own bucket, same first key, same surface discipline). The general
secret store is the third domain, and the pattern is now fixed:

- **Ops on the principal-scoped surface**: `secrets.put|get|list|delete`
  at `…<account>.<user>.secrets.>` — reach is structural (the transport
  op tail), so per-persona trees cannot collide and no caller can name
  a path outside its own reach (D3/D4: the same path in two accounts
  is two secrets by construction).
- **CAS revisions** on every write (D2 — the concurrent-writer rule the
  grants store measured under `-race`).
- **Sealed** `xkv1` to the deployment's first key at rest (D5/D7), the
  D16 envelope on the wire (D6).
- **Act-with, don't hand out** (D8): the preferred consumers are
  service-side acts that use a secret without returning it — the
  grants broker is the existing example; raw `secrets.get` exists for
  the owner's own prefix only, and is the exception, not the pattern.

**Acceptance criteria**: the D31 test battery generalized — CAS loser
loses nothing, at-rest positive-control grep, cross-persona reach
refused by the server with the delivery-log proof.

## D37 — The guardrail evaluator, and where it cannot be skipped

A guardrail answers the question the transport cannot: *may this
specific invocation, with these arguments, now, proceed* (B1) — and
never re-answers what the server already proved (B2).

- **Language**: CEL — sandboxed, non-Turing-complete, side-effect-free
  (B7 by construction, verified not assumed). Invalid or type-broken
  rules refuse at compile.
- **Input** (B9): `{server-proven principal, action, args, time}` plus
  counters the evaluator itself keeps for rate rules — never
  caller-supplied state. Nothing richer until a real rule demands it.
- **Three outcomes** (B3): allow, deny, **defer-to-human** — resolved
  per D38.
- **Rules are data** (B6): loadable and hot-reloadable without
  restarting the enforcement point.
- **Every evaluation is observable, including allows** (B5): one audit
  line per decision.
- **The evaluator discipline, measured and mandatory**: a tight cost
  limit sized to op-path rules (~10k), an interrupt check every ~100
  steps, and a **context deadline (25ms) as the hard stop** —
  belt-and-braces, never one mechanism. The scare that fixed this: a
  cost limit alone terminated a 100k²-element input bomb only after
  622ms; with the full discipline every hostile case died in under
  1ms. Allow-path cost at 100 compiled rules: p99 206–220µs against a
  2ms budget [measured].
- **Placement** (operator decision, B8): **unskippable at the
  capability chokepoints** — the custodian's op path first (refuse
  there and the action never had authority), the waker's wake decision
  and the door's tool-call forwarding following — and an **advisory
  library** everywhere else, same evaluator both places. A mediation
  proxy over plain publishes is refused: that is the second control
  plane episode 0002 already declined. Hosting untrusted code means
  its *capabilities* are guardrailed; its plain publishes remain
  transport-permission-bound, stated honestly.

**Acceptance criteria**: the Bar 3 matrix re-measured at the real op
path (hostile refusals, the latency budget), plus a hot rule change
converging without restart.

## D38 — An escalation approval is a one-shot delegation

One mechanism at two lifetimes (operator decision, B10): when an
evaluation defers to a human, the approval artifact is a **D33
delegation** — subject: the approving human; actor: the requesting
principal; bounds: the single invocation, named by
`hash(principal, action, args, time)`; TTL: minutes; verified from the
D26 directory like every delegation, honored only from its named
actor. Usable exactly once; it can authorize nothing else (B4). The
guardrail's defer emits the request; the human's yes mints the
delegation; the retry presents it. No second approval system exists.

**Reversal condition** (named at decision time): if binding one
invocation cannot be expressed without leaking its arguments into the
artifact (observable: an approval artifact carrying sensitive args),
this reopens as a new D-decision.

## Build order (from the topic's dependency argument)

F1's ensure-act (core doc) first — smallest, closes a live correctness
gap. Then the core grant-record vocabulary (C4) with Bar 4 as its gate,
then D36 (secrets), then D37/D38 (the evaluator at the custodian
chokepoint), then D35 (accounts — most irreversible, and everything
above teaches what an account must own). The A9/A10 rename rides its
own core cycle. Everything behind the 0071 focus gate: built when the
product demands it, designed now so the demand has an answer.
