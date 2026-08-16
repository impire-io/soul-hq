# 0003 — BYO NATS: the ceremony against a server soulstream does not run

**Status**: designed 2026-08-16 at the operator's direction (two
flavours, operator mode required, the operator applies the account half
themselves); resolves design
[`0001-soulnode-composition.md`](0001-soulnode-composition.md) §4's
**[O]** (the BYO ceremony subset) and closes the roadmap's
`ngs-capabilities` question by scope decision. Validation tags per the
[design README](README.md); requirement language per the same.

The capability, in one sentence: **a person who already runs a NATS
server — their own, or a Synadia Cloud BYON — can found a soulstream
realm on it, with soulstream never holding an operator or account
master key and never editing a server it does not own.**

## 1. Scope: two flavours, one refusal [D]

- **Self-hosted** — a server whose operator controls its configuration
  and speaks `nsc`. soulstream emits the exact commands and config
  fragments; the operator applies them.
- **Synadia Cloud BYON** — dedicated servers under Synadia's managed
  operator, account authoring exposed through the control-plane API.
  The wiring is measured
  ([episode 0038](../../04-JOURNEY/0038-soulstream-remote-mcp-node.md):
  callout enabled, wired, and fired via `control-plane-sdk-go`).

**Out, by name:**

- **NGS shared plans** — no access to auth-callout configuration
  [operator direction, 2026-08-16], and the admission model (0001 §5)
  is callout or nothing. The `ngs-capabilities` research topic the
  roadmap held open for this is closed unneeded: the flavour that
  required the answer is out of scope. *Reversal condition*: Synadia
  exposing external-authorization configuration on shared-plan account
  JWTs (observable: the setting appearing in their console or API)
  reopens NGS as a third flavour behind its own pass.
- **Conf-file auth servers** — the permission model rides on scoped
  signing keys carrying `jwt.UserScope` templates
  ([`nats-surface.md`](../soulstream-identity/nats-surface.md) D15/D25),
  an object that exists only in operator mode. A BYO target not in
  operator mode MUST be refused by name, the refusal stating the
  requirement and pointing at the kit's conversion fragments (§4) — an
  honest break, never a degraded lane. There is no second admission
  design for conf-auth and none is wanted.

Both flavours keep 0001 §5 admission byte-identical to the bundle:
operator mode, callout, sentinel + external credential. Constitution
(house, article II) already demands the embedded server run "exactly as
a hosted deployment does"; BYO is that sentence read in the other
direction.

## 2. The substrate contract [D]

What soulstream requires from any BYO server — the load-bearing
capability is **account-JWT authoring**, not config-file access:

1. Operator mode, nats-server ≥ 2.12, JetStream enabled.
2. An **AUTH account** whose JWT carries
   `authorization.{auth_users, allowed_accounts, xkey}` per
   [`auth-callout.md`](../soulstream-identity/auth-callout.md) D21 —
   `allowed_accounts` naming the realm account explicitly, never `*`.
3. A **realm account** with JetStream quota the realm can live in, a
   plain workload signing key, and a scoped signing key whose
   `jwt.UserScope` template is 0001 §4 step 4's, verbatim, values
   filled for this realm.
4. A way to distribute the sentinel. No config edit is required —
   the public bearer deny-all sentinel creds file (D19's fallback) is
   the default; `default_sentinel` is a nicety the self-hosted flavour
   MAY apply since it owns its config.

The substrate keeps: its operator key, its SYS account, every other
account on the server. soulstream MUST NOT author or push account JWTs
on a self-hosted substrate, MUST NOT write any server's configuration
file, and MUST NOT request Cloud-API scopes beyond the two accounts and
their keys. What it refuses to touch is everything it did not bring.

## 3. The split: local material, then the account half, then the wire [D]

0001 §4's steps regroup. Steps 1–2 (operator, SYS) belong to the
substrate in both flavours, permanently. Steps 3–4 (the two accounts)
become the **account half**, executed by the flavour's driver (§4, §5).
Steps 5–8 collapse into **local material** and the **wire half**,
both soulstream's own in every mode:

- **Local material** (before anything touches the substrate):
  soulstream generates the callout xkey, vault first key, service
  surface key (step 6), the AUTH issuer user keypair, and — in the
  self-hosted flavour — the three signing keys themselves (AUTH
  signing key, realm plain + scoped signing keys). A signing key is an
  nkey like any other; what makes it a signing key is the account JWT
  listing its public half, which is the account half's job.
- **The account half**: the two accounts come to exist carrying the
  public keys the local material declared (§4, §5 per flavour).
- **The wire half** (steps 7–8, unchanged in spirit): buckets, realm
  provisioning, founding acts over the public client — import the
  signing keys into the vault, mint the sentinel (locally, with the
  vault-held AUTH signing key; it MUST carry `issuer_account`, the
  D19 measured fact), create the first API token, the one secret
  `init` prints. Bypass-lane users (step 5) are minted locally with
  the vault-held keys — the AUTH issuer user's public key was declared
  to `auth_users` by the account half; the realm service, ops, and
  archivist users are signed by the plain signing key with their
  permissions in the JWT. No user is ever created on the substrate's
  side of the boundary.

**Custody, stated once** [D]: the vault ends holding the same keys as
the bundle — AUTH signing key, realm scoped + plain signing keys,
curve keys — in every mode. Only the direction key material travels
differs: self-hosted, soulstream generates and **no seed crosses the
boundary in either direction — public keys out, two account public
keys back**; Synadia Cloud, the platform generates and the programmatic
signing-key group's seed is returned exactly once, straight into the
vault [measured, 0038]. An operator or account master key never
travels at all, because neither side ever offers one.

## 4. The self-hosted flavour: the kit [D]

`soulstream init` against a self-hosted substrate emits **the kit** —
one generated document (stdout and `<state>/byo-kit.md`, regenerated
idempotently) with exact values, no placeholders the operator must
compute:

1. **`nsc` commands** — create/edit the AUTH account
   (external authorization: the issuer user's public key,
   `allowed_accounts` = the realm account by name, the xkey public
   half) and the realm account (JetStream limits, the two signing-key
   public keys, the scope template on the scoped key). Commands
   reference accounts by name so they run in one pass.
2. **Server config fragments** — for a server not yet in operator
   mode: `operator`, `system_account`, and the standard
   `nats-resolver` stanza, stated as the conversion the operator is
   choosing to perform, not a step soulstream performs. For a server
   already in operator mode: nothing required; `default_sentinel`
   offered as the one optional line.
3. **The push and the hand-back** — `nsc push` per their resolver,
   then the kit's final command prints the two account public keys;
   `init` consumes them (config: `byo.auth_account`,
   `byo.realm_account`) — the only thing that crosses back, and it is
   public.

The kit assumes `nsc` fluency [operator direction]: it shows commands,
it does not teach nkey custody. soulstream then **verifies
behaviourally and refuses by name**: connect with the self-minted
service creds (proves the plain signing key landed), whoami via
`$SYS.REQ.USER.INFO`, JetStream reachable on the realm account, the
issuer subscription live on the AUTH account, one callout smoke round
(sentinel + token admits scoped; garbage refuses). Each failed probe
names the kit item that was not applied, and `init` re-run remains a
verified no-op (0001 §4's idempotence rule, unchanged). Verification
never repairs: a nonconformant substrate gets a named refusal, not a
mutation — the provisioning posture soulstream-core already holds.

*Reversal condition*: if self-hosted founding fails in practice on the
`nsc` half (observable: recurring abandoned inits or support issues at
the kit step), the reopened question is a driver that authors the
account half itself against a handed operator *signing* key — a
different custody class, taken deliberately or not at all.

## 5. The Synadia Cloud flavour: the kit applied by an API [D]

Same split, no kit document: the account half is driven through the
control-plane API — the account, the programmatic signing-key groups
(seeds imported into the vault on their one return), the issuer user
under an on-demand group (the platform refuses users under
programmatic groups [measured, 0038]), callout configuration on the
AUTH account. The wiring is `cmd/byon-setup`'s, graduated from
soulstream-mcp into the ceremony as this flavour's driver; the Cloud
API token arrives by environment (`SOULSTREAM_SYNADIA_TOKEN`), is used
for the account half, and is not persisted in the state dir.
Verification and the wire half are §4's, identical.

## 6. Interface: the config block [D]

Per 0001 §2, BYO is configuration, not architecture: embedded server
`enabled: false`, every plane URL pointing at the substrate. This pass
adds one block:

```
byo: {
  flavour: "self-hosted" | "synadia-cloud",
  auth_account:  <public key>,   // self-hosted: handed back by the kit
  realm_account: <public key>,   // both flavours, once known
  synadia: { system: <id> }      // synadia-cloud only; token via env
}
```

No alias keys, no fallback spellings (pre-v1 clean-break rule). A
plane still MUST NOT behave differently because its URL is not
loopback; the bundle remains the deployment where every URL happens to
be `127.0.0.1`.

## 7. What this deliberately is not

- **Not account lifecycle.** Founding-time authoring of two accounts,
  once. Runtime account creation, suspension, and naming stay with
  `platform-tenancy-guardrails` A1–A8 — with one constraint recorded
  here for that topic: this pass commits to **both** custody arms of
  A8 (provider-held root behind an API, and operator-held root behind
  a human), so any A8 answer MUST serve both.
- **Not NGS shared, not conf-auth** (§1), not remote planes (0001 §2
  already carries those), not clustering (unmeasured, workloads 0003
  §7), not a second admission lane of any kind.
- **Not zero-touch for the operator.** The constitution's "no manual
  key step" (house, article V) governs soulstream's ceremony; the
  account half on a self-hosted substrate is the substrate operator's
  own domain — the job core's identity design already names as
  outside the protocol, like a DBA's. The kit makes that job exact;
  it does not pretend to absorb it. The bundle's easy path is
  untouched; BYO is the hard path kept possible, folded behind
  `byo.flavour`.

## 8. Acceptance criteria

1. **Self-hosted founding**: against a stock nats-server in operator
   mode with an empty realm, `soulstream init` emits the kit; the kit
   applied verbatim and the two keys handed back, `init` verifies and
   completes, and `soulstream up` (embedded server off) meets M1.1's
   semantics — status answers, sentinel + printed token admits the
   persona scoped to its own prefix, garbage and revoked tokens refuse
   with audited refusals, `init` re-run is a verified no-op.
2. **Refusals are named**: a conf-auth target is refused naming
   operator mode and the kit's conversion fragments; a partially
   applied kit fails verification naming the specific unapplied item;
   the wire never carries a specific reason (D20's generic-refusal
   rule holds).
3. **Synadia Cloud founding**: the same end state reached with the
   account half driven by the control-plane API on a Cloud BYON
   (0038's rig class), the signing-key seeds entering the vault on
   their single return, the issuer user under an on-demand group.
4. **Custody audit**: after founding, the state dir contains no
   operator seed, no account master seed, and no Cloud API token; on
   a self-hosted substrate, no file owned by the server was written
   by soulstream, and the substrate's resolver holds no JWT soulstream
   pushed.
