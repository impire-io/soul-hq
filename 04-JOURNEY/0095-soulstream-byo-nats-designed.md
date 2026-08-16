# Episode 0095 — BYO NATS designed: founding on a server we don't run (2026-08-16)

**The question:** composition 0001 §4 carried the product's oldest
deferred [O] — against a user-supplied server, which ceremony steps
apply and what soulstream refuses to touch. The operator set the scope
in discussion and the pass landed as design
[`0003-byo-nats.md`](../02-DESIGN/soulstream/0003-byo-nats.md):
exactly two flavours — a **self-hosted operator-mode server** whose
operator speaks `nsc`, and **Synadia Cloud BYON** driven through the
control-plane API — with the operator applying the account half
themselves. soulstream emits the kit (exact `nsc` commands and server
config fragments, no placeholders) and never authors or pushes an
account JWT on a server it does not own [judgment — operator
direction, recorded with its reversal condition]. NGS shared plans
are out by name: no access to callout configuration, and admission is
callout or nothing (0001 §5) — which closes the roadmap's
`ngs-capabilities` research topic *unopened*; the flavour that needed
the answer left the scope. Conf-auth servers refuse by name: the
permission model rides on scoped signing keys carrying `jwt.UserScope`
templates (nats-surface D15/D25), an object that exists only in
operator mode [mechanism-argument].

**The design's central fact is a custody symmetry.** A signing key is
an ordinary nkey — what makes it a signing key is the account JWT
listing its public half, and that listing is the account half's job.
So in the self-hosted flavour soulstream generates its three signing
keys locally and **no seed crosses the boundary in either direction**:
public keys out in the kit, two account public keys handed back at the
end [mechanism-argument]. Synadia Cloud is the mirror image: the
platform generates, and the programmatic signing-key group's seed
returns exactly once, straight into the vault — measured in episode
0038's BYON run, together with the platform refusing users under
programmatic groups (the issuer user needs an on-demand group)
[measured, 0038]. In both flavours the vault ends holding the same
keys as the bundle; an operator or account master key never travels
at all. No user is ever created on the substrate's side — bypass
users and the sentinel are minted locally with vault-held keys, the
issuer user's public key having been declared to `auth_users` by the
account half.

**What was reversed:** the [O]'s sketched default — "BYO mode requires
the accounts to exist and only runs steps 5–8" — survived as
mechanism and died as UX: the accounts must still exist before the
wire half, but they exist because the kit generated their exact shape,
not because a documentation page asked the operator to derive it. The
assistant's proposal that soulstream author the account half itself
against a handed operator signing key was rejected by the operator —
it survives only inside the reversal condition, a different custody
class to be taken deliberately or not at all.

**What it taught:** the load-bearing capability a BYO substrate must
offer is **account-JWT authoring, not config-file access** — the only
config-file item left in the whole design is optional
(`default_sentinel`; D19's public sentinel creds file already covers
the deployments that cannot set it). Steps 1–2 (operator, SYS) are
the substrate's permanently; 3–4 became the per-flavour account half;
5–8 regrouped into local material and the wire half, both soulstream's
own in every mode. Constitution article V holds without strain: the
account half on a self-hosted substrate is the substrate operator's
own domain — the "DBA job" core's identity design names as outside
the protocol — made exact by the kit, never absorbed. **Opened:** a
constraint recorded for `platform-tenancy-guardrails` — this pass
commits to both custody arms of its A8 (provider-held root behind an
API, operator-held root behind a human), so any A8 answer must serve
both.

Reversal condition: recurring abandoned inits or support issues at the
kit's `nsc` step (observable: issues recorded against founding on
self-hosted substrates) reopen the account-authoring driver; Synadia
exposing external-authorization configuration on shared-plan account
JWTs (observable: the setting appearing in their console or API)
reopens NGS as a third flavour behind its own pass.

Trail: design [`0003-byo-nats.md`](../02-DESIGN/soulstream/0003-byo-nats.md)
(new; indexed in the [design README](../02-DESIGN/soulstream/README.md));
[`0001-soulnode-composition.md`](../02-DESIGN/soulstream/0001-soulnode-composition.md)
§4's [O] resolved to a pointer;
[`auth-callout.md`](../02-DESIGN/soulstream-identity/auth-callout.md)'s
NGS tail closed (D11's reversal condition's NGS half, without
triggering); [`ROADMAP.md`](../03-IMPLEMENTATION/ROADMAP.md)
(`ngs-capabilities` struck, Later-horizons BYO entry updated); grounded
in [episode 0038](0038-soulstream-remote-mcp-node.md)'s Synadia Cloud
measurements; this change-set's commits.
