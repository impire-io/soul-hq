# 02-DESIGN — the normative design

What SoulIdentity *is*, functional-level: capabilities, seams, configuration
surfaces, acceptance criteria. Load-bearing decisions carry D-numbers with
their reasoning, so future changes argue against the real reasons. Behavioral
changes made during implementation propagate back here — these docs describe
the system as it is.

| Document | What it covers |
|---|---|
| [`agent.md`](agent.md) | The agent: vault, registry, oracle, mint — decisions D1–D13, D27–D29 and the milestone-1 shape |
| [`nats-surface.md`](nats-surface.md) | The NATS surface (M3): subject space, server-enforced principal, sealed envelope, KV vault, admin gate — decisions D14–D18 |
| [`auth-callout.md`](auth-callout.md) | Auth callout (M4): sentinel connection contract, issuer-as-mint, AUTH topology, the mapping shape — decisions D19–D22 |
| [`grants.md`](grants.md) | Outbound grants (the broker): the `grants.*` op family, the second custody domain, Article I's derived-credential line, delegated acting, the four lanes — decisions D30–D34 |
| [`tenancy.md`](tenancy.md) | Tenancy and guardrails: the `accounts.*` lifecycle with pluggable authority, the general secret store, the guardrail evaluator and its chokepoints, approvals as one-shot delegations — decisions D35–D38 |
| [`external-tools.md`](external-tools.md) | External tools: the two-layer catalog (record discovery face + plane custody), the `resources.*` op family, and the forwarding door's invariants (adapter position, endpoints only, no token held, the remote sees the calling person) — decisions D39–D41 |
| [`approvals.md`](approvals.md) | Approvals: the deferral becomes a durable ticket with its own TTL and witnessed expiry, the loop's public ends (`PresentApproval`, status/pending/list reads), notification as composition's duty, per-rule approver policy — decisions D42–D45, and D38's hash-sketch correction |

Future documents arrive by research graduation (see
[`../01-RESEARCH/README.md`](../../01-RESEARCH/README.md)) or design propagation
from landed work; the roadmap names the expected ones (auth callout with
claims-derived authorization, attestation issuance, sealing keys).
