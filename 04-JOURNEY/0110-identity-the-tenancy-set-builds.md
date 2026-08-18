# Episode 0110 — The tenancy set builds: secrets, the guardrail, approvals, accounts (2026-08-19)

Episode 0107's identity-plane design (D35–D38) became running code in
one overnight arc, three releases in dependency order:

**v0.4.0 — the general secret store (D36).** The sealed CAS pattern
extracted to `internal/sealedstore` the moment it had a second consumer
(the grants store refactored onto it, behavior unchanged, its suite
green throughout); `secrets.put|get|list|delete` on the principal-scoped
surface — caller-named paths structural under the persona's own tree,
D2's conditional writes (a stale revision loses loudly), sealed at rest
with the positive-control grep, always on beside the vault. Raw get is
the owner-only exception; act-with-not-receive stays the pattern.

**v0.5.0 — the guardrail at its chokepoint, approvals as delegations
(D37/D38).** Every sealed op is evaluated after unseal, before dispatch
— refused there, the action never had authority. CEL under the mandated
belt-and-braces (10k cost limit, interrupt every 100 steps, 25ms
deadline — the 622ms scare's lesson, now enforced in code); rules are
data, first match decides, an invalid rule refuses the whole load while
the running set stays whole, an erroring rule **fails closed**; every
evaluation observable. Defer names its invocation
(`hash(principal, action, argument bytes)`), and `approvals.present`
converts it with a subject-signed delegation — actor-bound (a stolen
approval refuses exactly as a stolen delegation does), usable exactly
once, dead in minutes, gone on restart, which fails closed. The
delegation primitives extracted to `internal/delegation`, shared with
the broker. `guardrail.load` is the hot reload, template-gated and
never gated by its own rules [all measured, both layers].

**v0.6.0 — tenancy at runtime (D35).** `accounts.create|suspend|resume|
resolve|list` behind the pluggable authority (A7/A8): the LocalOperator
arm signs COMPLETE account JWTs with the vault-held operator key (new
kind `nats-operator-key`) and lands them on the resolver through the
system-account connection as one act. Measured live on a dir-resolver
server [measured]: **store → admitted round trip in 1.69ms**, name
reuse refused first-seen, suspension refusing the next connection with
data intact, resume restoring admission. Creation completes the
composition: the new account's signing key enters the vault bound to
the new account — the D24 team binding — so the existing mint path
serves the new tenant the moment the op returns, and the seed appears
in no reply. Suspension's honest bound: it refuses *new* connections
immediately; in-flight connections drain at reconnect — the same
propagation class as D22's token revocation.

Named, not built: the ProviderAPI authority arm stays the seam (A8's
BYON half remains the operator-act residue); evaluator-held rate
counters wait for a rule that demonstrably needs them (B9's own rule);
per-rule approver policy for D38 (today any directory-resolvable
persona's signature approves — the mechanism landed, the policy layer
is a named follow-up).

Reversal condition: the designs' own conditions stand (D38's
argument-leaking approvals; D35's authority-cannot-share-a-process
observable). New from the build: if the guardrail's fail-closed rule
errors prove noisy in real deployments (observable: recurring
evaluation-error denials on well-formed ops), the error semantics
reopen as a D-decision rather than being quietly softened.

Trail: design
[`tenancy.md`](../02-DESIGN/soulstream-identity/tenancy.md) (D35–D38);
identity commits `0cc7cc1` (v0.4.0), `e7d6771` (v0.5.0), `ddeb9c0`
(v0.6.0); episode
[0107](0107-ecosystem-platform-tenancy-guardrails.md) (the design and
the measured discipline this build carries into code).
