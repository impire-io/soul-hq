# 04-JOURNEY — the narrative record

What was built, what was measured, what was believed and then refuted, and
what each episode taught — for the whole ecosystem, in one chronological
sequence. The design docs in `../02-DESIGN/` and the frozen specs in the
component repos say what each system *is*; these episodes say how we *got
here* — including the reversals, because a refuted assumption is as
load-bearing as the shipped code.

> **Keeping this log alive:** whenever a feature lands, a research
> investigation concludes, or a load-bearing decision is made, add a numbered
> episode with `/journey-log` (research topics get theirs via
> `/research-graduate`). Episodes are named `NNNN-<component>-<slug>.md` —
> one shared sequence; component tags are single-word (`core`,
> `workloads`, `identity`, `idp`, `shell`, `mcp`, `cli`, `soulstream`
> for the product, `ecosystem` for cross-cutting). Episodes ≤ 0069 keep
> the pre-rename tags — the [naming map](#the-naming-map-2026-08-13)
> resolves them. Follow
> [`TEMPLATE.md`](TEMPLATE.md) — including its required Reversal-condition
> line and evidence-class tags. Honesty rules apply here as everywhere:
> record what actually happened, including failures, reversals, and findings
> that contradicted expectations. This duty is anchored in
> [`../00-GENESIS/how-we-work.md`](../00-GENESIS/how-we-work.md); the
> numbering and index are enforced by `internal/hqlint`.

Episodes `0001`–`0049` were carried from the five per-project journals when
the hqs merged (2026-08-02, episode
[0050](0050-ecosystem-one-hq.md)) and renumbered into this shared sequence —
ordered by the original commit time of each episode, links retargeted,
prose untouched. The [numbering map](#pre-merge-numbering-map) below
resolves every pre-merge reference (frozen specs and old commit messages
cite the per-project numbers).

## Where things stand, per component

The per-component summaries below were carried whole at the hq merge and are refreshed with every new episode, as before.

### soulstream-core — the record (as of 2026-08-02; named soulstream until episode 0069)

**The node module is consumable** ([episode
0010](0048-soulstream-the-node-becomes-consumable.md)): its landing-day
`replace => ../` dropped the same day v0.7.0 made it unnecessary — the
module now pins the tag and downstream compositions can require it
(soulnode's Phase 2 front door is the first). Full node suite measured
green against the tag; co-development rides an untracked `go.work`
(soulrealm 0011's discipline, adopted).

The reference library has shipped the MVP and most of day-2 — **`v0.7.0`**
(2026-08-02) is current: foundation + op-log engine, CLI + MCP clients,
signing, rollup, scatter-gather discovery, the curator, work stages 1–2,
distribution, config-file identity, persona accountability
([episode 0001](0007-soulstream-genesis-and-the-reference-library.md), the founding
retrospective), the **memory convention**
([episode 0003](0009-soulstream-memory-convention-and-exhibits.md)): collective search as
graded testimony, portable self-authenticating exhibits, a public witness
surface — with the first archivist a separate repository,
[impire-io/soulstream-archivist](https://github.com/impire-io/soulstream-archivist),
**verified live against the NGS realm** (2026-07-26) — and **provisioning
byte limits** ([episode 0004](0010-soulstream-provisioning-byte-limits.md)): limit-
enforced accounts provision out of the box. Just landed: the **signer seam**
([episode 0006](0026-soulstream-the-signer-seam.md)) — signing delegated through
`identity.Signer` to an external custodian (SoulIdentity's M2 wiring
point), local keys the first implementation, a failing signer failing the
publish loudly — hardened for release the same day
([episode 0007](0027-soulstream-dx-hardening-and-the-cycle-guard.md)): typed-nil
signers refused at `Connect`, responder callbacks carrying the error
instead of a `-1` sentinel, and the cycle-guard dependency rule (neither
core repo imports the other; consumers wire the structural interface)
recorded on both sides — shipped together as **`v0.6.0`**. Then the **remote
MCP node** ([episode 0009](0047-soulstream-remote-mcp-node-built.md), **`v0.7.0`**): a
URL into the realm for clients that cannot install anything — credential-free
bearer passthrough onto per-principal callout-admitted connections, the tool
surface now a public embeddable `mcpserver` (SoulNode's fourth upstream ask,
plus `soulstream_whoami`), sign-in via an external OIDC authorization server
only (soulfold the intended default, the AS-facing contract proven to be the
interface), all five user stories measured on an in-process admission edge,
and a trust model that closes the prototype's forged-hint DoS. The **two-week
dogfood run started 2026-07-27** (protocol:
[`../03-IMPLEMENTATION/DOGFOOD.md`](../03-IMPLEMENTATION/DOGFOOD.md))
— daan, smith, and scribe on the NGS realm, the archivist keeping; its chafe
log feeds the eg-walker and sealed-topics gates. The central architectural
bet — leaderless coordination, no coordinator and no consensus — stands,
un-refuted.

**The project's working structure lives in `hq/`**
([episode 0002](0008-soulstream-adopting-the-hq-way.md)): GENESIS holds the vision, the
constitution (v1.1.0, wired into every spec-kit plan via the
`.specify/memory/constitution.md` symlink, now carrying the anti-drift Working
Agreement), and how-we-work; research runs a `/research-start` →
`/research-graduate` lifecycle; this journal is numbered episodes with the
structure enforced by `internal/hqlint` on the standard gate. The first
research topic has concluded: **sealed-topics**
([episode 0005](0011-soulstream-sealed-topics.md)) ran its four pre-registered bars in
a day and graduated to design — the sealed design survives the shipped
substrate, with amendments now folded into
[`../02-DESIGN/extensions/sealed-topics.md`](../02-DESIGN/soulstream-core/extensions/sealed-topics.md)
(the `{"ct"}` payload wrapper, signature-covered epoch/nonce, signing-chain
endorsement of sealing keys, key-carrying sealed baselines); the build's
priority stays gated on the dogfood chafe log (to 2026-08-10). A second
research topic has concluded: **remote-mcp-node**
([episode 0008](0038-soulstream-remote-mcp-node.md)) graduated to design after Bars 1–3
and its reversal-condition measurement PASSED on a live Synadia Cloud BYON —
a credential-less MCP node that passes the caller's bearer through to auth
callout, admitting a no-install client as a real, signed realm member. Bar 4
found the one gap that shaped the build: Claude Desktop's hosted connector
authenticates by OAuth only (no static-header lane), which the node answers
with an external OIDC authorization server — that topic is now **built and
shipped** as feature 018 ([episode 0009](0047-soulstream-remote-mcp-node-built.md)). What
is *not* yet built is the rest of the forward plan in
[`../03-IMPLEMENTATION/ROADMAP.md`](../03-IMPLEMENTATION/ROADMAP.md):
eg-walker live co-editing, sealed topics, and a browser/WebSocket client.

### soulstream-workloads — the room (as of 2026-08-01; named soulrealm until episode 0069)

**Soulstream is pinned at v0.6.0 — the dev replace is gone** ([episode
0011](0037-soulrealm-pinned-to-the-record.md)): soulnode's composition research named
the filesystem `replace` a release blocker for any consumer wanting to pin
soulrealm; soulstream's tag proved current (main = v0.6.0 + docs only
[measured]) and the whole gate runs green against it. Co-development now
rides an untracked `go.work`; soulstream changes arrive by tag bump.

**The project was founded** ([episode 0001](0001-soulrealm-genesis.md)): soulrealm is the
runtime companion to soulstream — soulstream records, soulrealm runs. The hq is
bootstrapped from the sibling project's proven structure, with a constitution
whose non-negotiable article is the **substrate boundary** (soulrealm never
becomes a store of record; everything worth keeping flows back to topics as
ops).

**The substrate question is decided** ([episode
0002](0002-soulrealm-the-substrate-decision.md)): after a live NEX spike and a source
read, **soulrealm builds its own runtime — NEX as influence, not dependency —
with the soulstream op-log as the single control plane.** Measured: role
(`agent`/`tool`) is orthogonal to lifecycle (`service`/`function`/`job`); NEX
issues scoped per-workload identity for free and is embeddable via public
options (so a fork was never forced). The `[judgment]` call to rebuild rather
than embed turned on **not running a second control plane** (`$NEX.control.*`)
beside the op-log — recorded after teach-back, with the embed case argued at
full strength first. That opened design doc
[`0001-soulrealm-runtime.md`](../02-DESIGN/soulstream-workloads/0001-soulrealm-runtime.md): the
single-plane runtime, the role×lifecycle model, a realm-semantic per-workload
minter, lifecycle-as-ops, and pluggable backends, with an honest NEX influence
ledger and its `[O]` open sub-questions.

**The first slice is specced** ([`specs/001-launch-an-agent/spec.md`](../../soulstream-workloads/specs/001-launch-an-agent/spec.md)):
declare an `agent`/`service`, mint a persona-scoped credential, launch it
native, post a turn as its persona, lifecycle visible as ops — no second
control plane. Minimal spec-kit scaffolding is in place (`.specify/` with the
constitution symlinked). The signing story is resolved soulrealm-held.

**Scope is soulstream-only** ([episode 0003](0003-soulrealm-soulstream-only-scope.md)):
soulrealm depends on soulstream and nothing else — no Impire-platform services
(identity, tenancy, vault) for now. The minter stays a seam for a future
external authority, but none is designed in.

**M1.1 is implemented** ([episode 0004](0004-soulrealm-the-first-agent-runs.md)): the Go
module `github.com/impire-io/soulrealm` exists, and an agent launched by
soulrealm posts a turn attributed to its persona while its lifecycle shows up
as `work.open/claim/done` on the topic — proven end-to-end (SC-001, SC-002),
whole gate green. The plan's bet held: **no new soulstream vocabulary** —
soulrealm is the work.md "runner". Six packages (declaration, minter,
backend/native, runner, two cmds), pure logic split from I/O so most tests need
no server; the native backend proves it does not leak soulrealm's secrets into
a workload. All five success criteria met (SC-003 enforcement via an in-process
operator-mode server).

**M1.2 is done** ([episode 0005](0005-soulrealm-a-tool-answers.md)): soulrealm launches a
`tool` service and an agent discovers it by name and calls it (uppercase round
trip). Added the tool role, role-aware scopes, and the runner's launch/stop
(services don't self-exit). A measured lesson landed a boundary: tool
request-reply is transient, so it rides soulrealm's own `SOULREALM.SVC.*`
subjects, not the stored `SOULSTREAM.>` stream (which would ack and race the
reply). SC-001/002/003 proven end-to-end.

**The hq is now aligned with its own contract** ([episode
0006](0006-soulrealm-hq-alignment.md)): the "hq structural lint" the constitution and
skills had cited as the enforcement backbone is finally built —
`internal/hqlint`, a test-only Go package that rides `make test` and the commit
gate. Along the way: README/CLAUDE status corrected to Phase 1 (M1.1 + M1.2
done); specs 001/002 marked Shipped with 002's spec-kit short-circuit recorded
honestly; the full spec-kit flow vendored from pra (which also fixed a
plan/tasks template that had baked in pra's constitution principles); and
Article VI clarified (constitution 0.1.1).

**M1.3 is done — Phase 1 is complete** ([episode 0007](0020-soulrealm-a-second-wall.md)):
the byte-identical M1.1/M1.2 declarations run inside **microsandbox**
microVMs (an open amendment: the roadmap had said Docker/Firecracker), with
the identical op mapping on the topic and a measured isolation boundary — a
host path readable natively is denied in-guest. The `msb` CLI is supervised
as a child process (no CGO SDK); loopback NATS is rewritten to the guest's
host alias under a host-only network policy; backend selection is node-side
(`SOULREALM_BACKEND`), and the declaration still cannot name a backend. The
default suite stays hermetic (stub CLI); `make test-msb` boots real microVMs.
A refuted diagnosis is on record (":ro unsupported" was really msb failing on
symlinked mount sources), plus a named limitation (remote NATS needs the
`public` profile — Fleet-era).

**The Kubernetes gate is met** ([episode 0008](0024-soulrealm-kubernetes-backend.md)):
the `kubernetes-backend` research topic pre-registered four bars and four
spikes measured all four PASS — a prototype backend behind the *unchanged*
seam ran the byte-identical M1.1/M1.2 declarations as pods on kind (agent,
tool round trip, crash → abandon), and a scope probe inside a pod against
Synadia NGS was denied out-of-scope with its credential delivered as a
Secret. One expectation inverted: on non-loopback NATS, Kubernetes is
*ahead* of the microVM backend (ordinary pod egress vs msb's Fleet-era
`public` profile). Named honestly: pods are weaker isolation than microVMs —
the case is adoption, not isolation. Opened design
[`0002-kubernetes-backend.md`](../02-DESIGN/soulstream-workloads/0002-kubernetes-backend.md).

**M2.1 is done — Phase 2 is complete** ([episode
0009](0028-soulrealm-a-third-wall-lands.md)): `backend/k8s` landed through the full
spec-kit flow — one runner-supervised pod per workload, credential as a
Secret that never touches host disk, artifact as a per-run OCI image on the
CA-trusted base pushed digest-pinned to the operator's registry. Five e2e
scenarios green on kind + a local registry in ~26 s (`make test-k8s`);
default gate hermetic; `make test-msb` still green. Two recorded decisions
from planning: the artifact channel **reversed** from the draft's node HTTP
to an OCI-registry interface (maintainer decision — an open amendment to
design 0002's candidates), and client-go was chosen after teach-back, kept
entirely inside `backend/k8s`. **Next:** research gates for the later
horizons — Fleet, sandboxes (stage 5), tool ecosystem — and the `nats://`
artifact-addressing question at the artifact-registry milestone.

**The Fleet gate is met** ([episode 0010](0033-soulrealm-fleet.md)): the `fleet`
research topic pre-registered three bars and four spikes measured all three
PASS — placement **is** `work.claim` (exactly-one-launch 120/120 contested
rounds, replay-reconstructible, zero transient signaling), node death
reclaims via *projection nominates → probe vetoes → ordinary `work.abandon`
decides* (10/10 kills per variant within bound; the probe eliminates the
live-silent false positive at zero cost on true deaths), and a node without
the signing seed launches a scope-enforced workload (expiry floor measured
at 10 ms). Two open reversals on the record: Bar 2 amended pre-run when
work.md's timeout-by-projection surfaced, and spike 3's judgment against
scoped signing keys fell to spike 4's tag-template measurement — the minter
role dissolves into the identity plane (`soulidentity`), amending episode
0003's soulstream-only scope. Opened design
[`0003-fleet.md`](../02-DESIGN/soulstream-workloads/0003-fleet.md); roadmap Phase 3 (Fleet) is
unblocked. **Next:** the spec-kit pass for the first Fleet milestone; the
soulidentity tags-on-mint addition gates the preferred minting path.

### soulstream-identity — the name (as of 2026-08-02; named soulidentity until episode 0069)

**The default IdP is a sibling — soulfold, the refusal holds**
([episode 0019](0039-soulidentity-soulfold-the-default-idp.md)): the operator's
default-IdP question — should SoulIdentity, bearing the name, also be
the passkey-first OIDC provider deployments get out of the box — was
answered by holding the vision's refusal: identity truth stays in the
deployment's IAM, and the default IAM becomes **soulfold**
(`github.com/impire-io/soulfold`), a NATS-native (JetStream KV),
embeddable, passkey-first OIDC provider the callout issuer treats
identically to Entra — the D23 seam only, no side-channel, no shared
store, no precedence [mechanism-argument]. "Default" is distribution
wiring (`--oidc-issuer`), replaceable by any OIDC provider by config.
Pocket-id itself was ruled out as the bundle not for being Node (it is
Go since v1) but for being an application with an SQL store, not an
embeddable library. Named, not built: D23 multi-issuer dispatch for
deployments running soulfold beside a second issuer.

**The embed seam — D29, M2's second consumer-proven addition** ([episode
0018](0036-soulidentity-the-embed-seam.md); D29 in
[`../02-DESIGN/agent.md`](../02-DESIGN/soulstream-identity/agent.md), feature
`specs/002-embed-seam/`): soulnode's composition research measured the
wall — provisioning is public, the serve assembly was `cmdServe`-only,
and two consumers had already ridden the module-namespace dodge to reach
`internal/`. The public `embed` package now exposes the one seam:
`Run(ctx, Options)`, value-only options, custody unchanged (D13),
provisioning still wire-only; the daemon is its first consumer (one
assembly, two entrypoints). The proof is compiler-grade:
`e2e/embedgate/`'s module path sits outside the repo namespace, so an
`internal/` import cannot compile; the gate runs the M4 admission shape
through `embed.Run` [measured]. A drain lesson strengthened the contract:
`Run` flush-confirms its unsubscribes, so returned means silent. Existing
gates unchanged, uncached-green [measured].

**Role selection by declared name — D28, the first consumer-proven
M2 addition** ([episode 0017](0035-soulidentity-role-selection-by-name.md); D28 in
[`../02-DESIGN/agent.md`](../02-DESIGN/soulstream-identity/agent.md), answering
soulidentity#1): soulrealm's fleet needs one scoped signing key per role
on one realm account — the exact observable D5's amendment reversal
condition watched. The `mint.ephemeral` op issues an ephemeral scoped
user JWT against a **named role** for a **caller-supplied public key**
(no seed in either direction; the response is the JWT alone), with
**tags** in the user claims for scoped templates to resolve and a
required TTL (D22's bound). The nouns, corrected by the operator the
same day: a **team is the account, the tenant** (0013's "teams are
accounts"); the declared signing key is a **role** — wire field, client,
CLI, and internals renamed before any consumer wired in. Proven in the
M3-gate e2e [measured]: the by-name-minted JWT admits to the
operator-mode realm; with a second role imported the binding path refuses
as ambiguous while both roles stay reachable by name. Named, not built:
the token lane's own named-role answer (node enrollment), and per-role
tag policy.

**One noun: persona** ([episode 0016](0032-soulidentity-one-noun-persona.md); D27 in
[`../02-DESIGN/agent.md`](../02-DESIGN/soulstream-identity/agent.md), constitution 1.3.1):
persona == identity — the ecosystem's one noun for the represented
subject, adopted from soulstream's fixed terminology; *principal* is the
server-proven (account, user) a connection speaks as (D15's term), and
"identity" survives only in the product name. The persona is born at
first encounter (D26); the vault is where its durable artifacts live. A
vocabulary pass over vision, constitution, missions, and load-bearing
comments — no type, op, or JSON field changed [measured: gate green].

**The vault is the directory — ephemeral users, keys on first touch**
([episode 0015](0031-soulidentity-the-vault-is-the-directory.md); D26 in
[`../02-DESIGN/nats-surface.md`](../02-DESIGN/soulstream-identity/nats-surface.md)): episode
0014's persona-directory trust path was refuted by the operator the same
day — identity truth lives in the IAM, users are ephemeral, and no
per-user act may exist anywhere. A persona key is a capability artifact,
not identity (an access token carries no user key and cannot sign
records): the caller's own key **materializes inside the vault on first
touch**, owner-bound, and `keys.public` is the **open directory read** —
readers build verification keyrings from the identity plane; no profile
store. The sealed-communication key follows the same pattern when D9
lands.

**M2's first gate criterion is measured — the seam carries a real record**
([episode 0014](0030-soulidentity-the-cross-service-proof.md), re-proven on the D26
shape in 0015): a Soulstream record signed through the running
SoulIdentity service verifies in a real realm. The proof lives in `e2e/`
— a consumer-position module importing both repos (the cycle guard's
shape), riding `make test`, now with **zero per-user acts**: one team
key, one minted credential, the key materializing at signer construction,
the reader's keyring one `keys.public` answer; announce, baseline, and
turn read `SigVerified` — `unknown-key` without the keyring, the negative
control [measured]. What remains of M2 is the **node half** (one pooled
connection per user, no node-held creds), in soulstream's remote MCP node
feature — which publishes nothing per user; it builds keyrings from the
identity plane.

**The identity registry is dissolved — authorization lives in the ACLs
and the bindings** ([episode 0013](0029-soulidentity-the-registry-dissolves.md); D25 in
[`../02-DESIGN/nats-surface.md`](../02-DESIGN/soulstream-identity/nats-surface.md), amending
D2/D5/D6/D18/D22/D24): the operator's one-identity-one-persona question
unwound the ledger field by field. Persona keys carry their owner
(account, user) at import and `sign.record`/`keys.public` check the
caller against the binding; every mint resolves its signing key by the
target account's D24 team binding (ambiguity refuses); the management ops
are gated by the server's own permission enforcement on the op tail —
`requireAdmin`, the `admin` flag, `identities.*`, and `internal/registry`
are deleted; mint is an operator op (self-mint died with the row that
authorized it). The token store is the one registry standing. All three
e2e gates re-proven on the new shape [measured], including the new
op-tail proof (a represented user publishing a management op on its own
prefix: server-refused, zero service decisions) and the revocation bound
(re-admitted 5.25 s after connect at a 5 s TTL). The client now carries
the M2 seam surface: `PersonaSigner` satisfies soulstream's
`identity.Signer` structurally, exercised in the M3 rig. Next: **M2 —
consumers wire in** (the cross-repo gate proof, then the remote MCP
node).

**The Entra/OIDC lane is open — role == team, no mapping store**
([episode 0012](0025-soulidentity-entra-role-claim-lane.md); D23–D24 in
[`../02-DESIGN/auth-callout.md`](../02-DESIGN/soulstream-identity/auth-callout.md), built as
spec-driven feature `specs/001-entra-oidc-backend/` through the newly
enabled speckit flow): an external client presents an Entra access token
instead of an API token; the issuer validates it against one pinned
issuer/audience via JWKS discovery (fail-closed, key rollover without
restart) and authorizes by the `roles` claim — the role value IS the team
name, resolving directly against the vault's account signing keys, which
now carry their account binding at import. D22's sketched rule table was
refuted before it was built; no catalog, no per-user entries; admin and
personas never come from claims. Gate met on the stub rig [measured]:
sealed-leg admission with full attribution, the nine-row refusal matrix,
`sit_` lane untouched, and the accepted revocation asymmetry demonstrated
(token lifetime + one TTL; cached token re-admitted 5.2 s after connect
at a 5 s TTL). Real-tenant verification is a documented manual runbook.
Next in the execution order remains **M2 — consumers wire in**.

**The subject space gained its ecosystem namespace**
([episode 0011](0023-soulidentity-the-shared-subject-prefix.md), D14 amended at the
operator's direction): the root is `<prefix>.soulidentity` with the prefix
shared across all soulstream components (`--prefix` /
`SOULSTREAM_PREFIX`), empty by default. Environments coexist in one realm,
and the account token sits at declared position `P+2`, so a cross-account
export (`account_token_position`) extends D15's principal proof by
configuration alone. The M3 gate e2e now runs fully prefixed
(`prod.soulstream.soulidentity.>`) [measured]; the honest cost — prefix
mismatch is silent timeouts — is mitigated by the startup root log and the
shared environment variable.

**Milestone 4 — auth callout, the front door — is shipped**
([episode 0010](0022-soulidentity-m4-auth-callout-ships.md); design
[`../02-DESIGN/auth-callout.md`](../02-DESIGN/soulstream-identity/auth-callout.md), D19–D22,
researched in [episode 0008](0019-soulidentity-sentinel-credential-flow.md) and
[episode 0009](0021-soulidentity-claims-mapping-shape.md)): an external-identity client
holds a sentinel creds file (minted by the admin-gated `sentinel.mint` op,
public by design) and its API token; the issuer — one process, two
connections, the callout subscription in the dedicated AUTH account —
digest-validates against the token store (records name an identity, carry
no policy), authorizes via the registry row, and mints a TTL-bounded
scoped JWT for the server-assigned key with the vault's role keys. Gate
met [measured]: admission with server-enforced scope and full audit
attribution; bypass-lane connections produce zero callout decisions;
invalid and revoked tokens refused; the D21 xkey leg proven (sealed
requests and responses). Token management is four admin-gated surface ops.
Entra/OIDC arrives later as validator configuration on the same D22
interface. Open: `ngs-capabilities` (blocked on operator access to the
Synadia account; gates only the NGS deployment class), and next in the
execution order, **M2 — consumers wire in**.

**Milestone 3 — the NATS-native rebuild — is shipped**
([episode 0007](0018-soulidentity-m3-the-nats-native-rebuild.md); design
[`../02-DESIGN/nats-surface.md`](../02-DESIGN/soulstream-identity/nats-surface.md), D14–D18):
the service answers over sealed NATS request/reply on the caller's own
subject prefix — the principal proven by the server's publish-permission
enforcement (D15, via the scoped-key template
`soulidentity.{{account-subject()}}.{{name()}}.>`) — and the vault seals
into NATS KV with both xkeys deployment-supplied. Act-as (D6) is enforced
and audited against real principals; management is admin-gated in the
registry (D18, the socket trust model's successor). All four gate criteria
measured in the e2e proof: unauthorized act-as refused and logged, wire and
store ciphertext-only (positive-control-verified), cross-prefix requests
refused by the server itself. The socket agent, `NATSOption`, file
keystore, and `sign/nonce` are deleted. Next in the execution order: M4
(auth callout, the front door), then M2 (consumers wire in). Design and
review that preceded the build:
[episode 0005](0016-soulidentity-the-nats-surface-design.md),
[episode 0006](0017-soulidentity-design-review-amendments.md).

**The first-key story is decided — M3's research gate is open**
([episode 0004](0015-soulidentity-first-key-story.md), D13): the unwrapping xkey for the
KV backend's envelope encryption is a local root secret on the service host
— decided as a `0600` file, then amended at design review to
deployment-supplied environment configuration
([episode 0006](0017-soulidentity-design-review-amendments.md)) — named honestly as
plaintext in the same trust class as the creds file, with the envelope's
real gain being that broker disks,
replicas, and backups never hold plaintext seeds. All three pre-registered
bars passed [measured]: the sealed round-trip survives broker+service
restart unattended, the store holds ciphertext only
(positive-control-verified), and the from-nothing bootstrap is two operator
acts plus one automatic service act.

**The mission was re-centered, twice in one day**
([episode 0002](0013-soulidentity-the-identity-plane-re-centering.md) then
[episode 0003](0014-soulidentity-nats-only-and-the-connection-ladder.md); constitution
1.1.0 → 1.2.0): SoulIdentity is the identity plane of the Soulstream
ecosystem — the representation of identity for humans and agents, delivered
as a **NATS-only** service with xkey-sealed E2E request/reply (D11, D12).
There is no socket: a presented creds file bypasses SoulIdentity entirely
(self-custody, server-verified natively), everything else arrives through
auth callout, authorized by the declared registry or by validated claims in
the presented token (D2). NATS KV with xkey envelope encryption is the
vault's initial backend. The milestone-1 socket surface, `NATSOption` seam,
and file keystore are transitional until the NATS-native rebuild (M3).

**Milestone 1 — the walking skeleton — is shipped**
([episode 0001](0012-soulidentity-genesis-and-the-walking-skeleton.md)): vault, declared
identities, the agent on a Unix socket, scoped minting, and the `client`
package with `NATSOption`, proven end to end against an operator-mode NATS
server — mint through the agent, nonce signed in the vault, scope enforced by
the server, no seed ever in the client process [measured]. The design is
twelve numbered decisions in [`../02-DESIGN/agent.md`](../02-DESIGN/soulstream-identity/agent.md);
the plan is [`../03-IMPLEMENTATION/ROADMAP.md`](../03-IMPLEMENTATION/ROADMAP.md)
(execution order M3 → M4 → M2: the NATS-native rebuild, then auth callout as
the front door, then consumers wire in over the NATS surface). The first
release tag arrived 2026-08-02 — `v0.1.0`, for soulfold's pin; open
questions before their milestones (NGS callout
capabilities, the sentinel-credential flow, the first-key story, the
claims-mapping shape, service round-trip latency) are named on the roadmap.

### soulstream — the product, the house (as of 2026-08-13; named soulnode until episode 0069)

**The helm plane — the cockpit joins the bundle** ([episode
0068](0068-soulnode-the-helm-plane.md); `specs/007-the-helm-plane/`):
`planes.helm` runs soulhelm v0.1.0 through its public embed seam — on
by default beside the fold, absent-block-means-disabled on old state
dirs, `helm console` the fourth URL at `up`. Decision of record:
`SessionIssuer()` — enabling the helm switches the identity plane's
OIDC admission lane on in local mode with the bundled fold as AS (an
explicit external AS wins; ephemeral listeners resolve to no issuer
and the helm skips loudly). Named, not built: a dedicated scoped helm
ceremony user (the plane hands the ops lane today).

**The front of house — URLs, the console, no collision**
([episode 0012](0062-soulnode-the-front-of-house.md); `v0.3.0`): `up`
now logs the MCP door, the fold sign-in, and the admin console URLs;
the fold is on by default (`init && up` lands a person at a passkey
prompt with an admin console); the door and fold can't share a
listener (refused by name; public mode needs two fronted routes); and
the bare-IP RP-id footgun is caught at load (issuer defaults to
localhost). Two bugs the live browser run surfaced were fixed
upstream and pinned: the invite link dead-ended (soulfold grew a
standalone `/enroll`, v0.3.1) and a phantom duplicate user (CreateUser
made index-first, v0.3.2). Verified the whole human path in a real
Chromium: click the enrol link → create a passkey → land in the
console.

**v0.1.0 released — the house gets a shipping label**
([episode 0009](0058-soulnode-the-release-pipeline.md)): CI (first
runner run green, folded-realm gate included) and the tag-triggered
goreleaser release, archivist pattern; `v0.1.0` published with
linux/darwin binaries, an artifact round-tripped, checksum-verified,
answering `soulnode version` → `0.1.0`. No windows on purpose (the
ceremony's owner-only modes). The operator made soulrealm public —
the whole consumed stack now fetches with zero credentials, and the
"private-module credential" blocker is gone. **Next:** day-2 items,
soulfold M3 upstream, or Phase 3's tsnet gate.

**The folded realm — one binary, real people**
([episode 0008](0057-soulnode-the-folded-realm.md);
`specs/006-the-fold-plane/`): `planes.fold` runs the bundled
passkey-first OIDC provider (soulfold by tag, public embed seam) on
the node's own JetStream, and public door mode defaults its AS at it.
Measured end to end with zero external services: DCR → passkey
sign-in (first touch enrolls) → bearer MCP session at the door →
`whoami` names the fold identity; founding token coexists; old state
dirs load unchanged. Startup order recorded as load-bearing (fold
before identity plane). Two upstream tags rode along (soulfold v0.1.1,
v0.1.2). **Next:** release/distribution polish, the fold's M3
upstream, or Phase 3's tsnet gate.

**The public door opens — Phase 2 completes**
([episode 0007](0055-soulnode-the-public-door.md);
`specs/005-public-door/`): `planes.door` grew
`public_url`/`auth_issuer`/`auth_audience` additively and the node now
serves the whole hosted-client OAuth story — a scripted client knowing
only the door URL walks challenge → resource metadata → DCR →
code+PKCE → token, and the bearer forms an MCP session whose realm
identity is the token's subject (lane=oidc, role=realm in the audit).
Founding token coexists; undeclared roles and garbage refuse; the
listener stays loopback with HTTPS as deployment fronting. Measured
against the upstream AS contract's stand-in — AS-agnostic, soulfold
the intended default (its own admission proof is episode 0054). Bonus
root-cause fix: listen port 0 no longer collides with nats-server's
default-port reading. **Next:** the bundled-fold wiring (soulfold M5),
day-2, or Phase 3's tsnet gate when fronting measures insufficient.

**The door is open — Phase 2's local mode is done** ([episode
0006](0049-soulnode-the-door-opens.md); `specs/004-the-front-door/`): upstream's
remote MCP node (soulstream 018, v0.7.0; made consumable in its journey
0010) runs as the door plane — `planes.door`, loopback HTTP, static
bearers through the realm's own callout, custodying nothing. Measured: a
real MCP client with the founding token forms a session, lists the tool
surface, and `whoami` names the realm-admitted owner; garbage bearers
refuse; the state dir is untouched by the door. The vision sentence
executes: `init && up`, paste the token, work. Public mode waits on
soulfold upstream. The pseudo-version exception closed 2026-08-02: all
four upstreams tagged (soulrealm v0.1.0, archivist v0.2.0,
soulstream/node v0.7.0, soulidentity v0.1.0) and soulnode pins them by
tag. **Next:** public door mode when soulfold lands, or Phase 3's tsnet
gate when Phase 2 fronting is measured insufficient.

**Phase 1 is complete — an agent runs** ([episode
0005](0045-soulnode-an-agent-runs.md); `specs/003-an-agent-runs/`): `soulnode
workload start` is the invocation-scoped runtime plane — upstream's own
`agent-echo`, declared unchanged, launches with a minted TTL-bounded
credential, posts as its persona, and its lifecycle lands as a completed
work item owned by `runner`; nothing credential-shaped lingers after end
of life. The ceremony carries the two-keys split (plain workload minting
key beside the scoped admission key; inventory 20). The composition
caught a consumer-proven upstream bug on its first enforcing run —
agents lacked `$JS.API.INFO` for the realm client's availability probe —
fixed upstream first (soulrealm `3fee11f`), pin bumped. Three
pseudo-version pins await upstream tags. **Next:** Phase 2 — the front
door, gated on soulstream's `018-remote-mcp-node` (in flight upstream).

**M1.2 is done — the realm remembers** ([episode
0004](0044-soulnode-the-realm-remembers.md); `specs/002-realm-joins/`): `init`
founds the record substrate; the memory plane runs in-process (public
`keeper`/`archive` on a realm client signing through the identity plane —
persona key vault-held, nothing on disk). Measured: the owner's full path
through admission (post a turn → archivist keeps it, author `owner` →
memory answers with attribution and a citation) rides `make test` in
~5 s; restart continuity exactly-once; the disabled-plane arm clean.
`config.json` carries `realm` (default `home`) and the first plane block.
**Next:** M1.3 — an agent runs (the runtime plane; carries the known
workload-signing-key design question).

**M1.1 is done — first boot is real** ([episode
0003](0042-soulnode-first-boot-is-real.md); `specs/001-init-and-up/`): `soulnode
init && soulnode up` founds and runs a realm — the whole ceremony
persisted into one state directory (17 artifacts, owner-only modes), the
founding acts through public surfaces, the first token printed once, the
embedded operator-mode server + identity plane on ordinary loopback
connections. Measured: init in 0.15 s; the found→admit→refuse→revoke→
restart e2e rides `make test` in ~1 s; re-init is a verified no-op. One
refuted assumption on record (refuse-on-0755 became tighten-then-verify).
Dependency exception tracked: soulidentity pinned at a pseudo-version
until it tags. **Next:** M1.2 — the realm joins (provisioning + the
archivist plane).

**The composition gate is met — Phase 1 is unblocked** ([episode
0002](0040-soulnode-the-composition-gate.md)): the `single-binary-composition` topic
measured all three pre-registered bars PASS — embedded admission parity
(4 ms token-lane admission, server-asserted principal, refusals audited),
the embed surface per component (three of four upstream asks *delivered*
before graduation: soulidentity's public `embed.Run`, the archivist's
public `keeper`/`archive`, soulrealm pinned to tagged soulstream v0.6.0;
the fourth — soulstream's public MCP surface — held for the maintainer's
open `remote-mcp-node` topic, gating only Phase 2), and the first-boot
ceremony fully generated by code from an empty directory. The maintainer's
transport call superseded the proposed split: **everything through
loopback** — decomposition is configuration, never architecture. The
constitution ratified at **1.0.0** with Article III reworded accordingly.
Design [`0001-soulnode-composition.md`](../02-DESIGN/soulstream/0001-soulnode-composition.md)
is open. **Next:** the spec-kit pass for M1.1 (`soulnode init` + the
server and identity plane).

**The project was founded** ([episode 0001](0034-soulnode-genesis.md)): SoulNode is
the single-binary distribution of the Soulstream ecosystem — embedded NATS,
SoulIdentity, the archivist, soulrealm, and an MCP front door in one process
on a machine the user owns, with the first boot (`soulnode init`) treated as
the product. The hq is bootstrapped from soulrealm's proven structure. The
constitution's load-bearing articles are **composition, not invention**
(SoulNode wires public tagged surfaces; domain logic lands upstream) and
**same shape as any deployment** (operator-mode NATS + auth-callout locally,
no dev fork). Feasibility entered measured; so did the obstacles
(SoulIdentity's serve path was `internal/`, soulrealm's `replace`
directive) — both since dissolved by the upstream landings above.

### soulstream-idp — the fold (as of 2026-08-03; named soulfold until episode 0069)

**The admin console — a browser and your passkey**
([episode 0011](0061-soulfold-the-admin-console.md); `v0.3.0`): the
admin surface split (D25) — the JSON API moved to `/api/admin`, and a
server-rendered console landed at `/admin`, authenticated by a
passkey session gated on the `admin` group. From a browser: create
people, mint their enrolment invites (shown once), set groups,
disable accounts, register clients — CSRF-guarded, no server-box
access. Verified in a real Chromium session with a CDP virtual
authenticator, which caught a genuine constraint the Go test could
not: WebAuthn refuses a bare-IP issuer host (use `localhost` or the
fronted name, never `127.0.0.1`). Browser-session logic extracted to
`internal/websession`.

**M3 ships — the fold is complete**
([episode 0010](0060-soulfold-m3-the-lifecycle.md);
`specs/005-the-lifecycle/`, **v0.2.0**): users, groups→roles, invites,
client registration, and the `/admin` JSON surface on the graduated
lifecycle design — enrollment everywhere by single-use digest-stored
invite, first-touch deleted, the four research bars riding `make test`
permanently. The M3 gate measured: four acts to a signed-in admin;
membership changes in the next token; admin-surface authz. Soulnode
composed it the same hour: the founding invite prints once beside the
founding token, and the owner's bearer carries admin (fold surface) +
realm (callout admission) with the two authorities cleanly separate.
**Every milestone on the fold's roadmap is shipped.**

**The bootstrap story concludes — invitation is the only door**
([episode 0009](0059-soulfold-bootstrap-story.md)): M3's gating
research passed all four bars against the fold itself — a fresh fold
to a signed-in admin in four counted acts, invites exactly-once
(25/25 races) and digest-stored with positive controls, the opened
store's 29 artifacts all refusing as enrollment, and the pocket-id
audit fixing M3's scope (its open `/setup` bootstrap refused — the
operator-act invite is stronger trust at the same act count).
First-touch enrollment is deleted, not disabled. Third design doc:
[lifecycle](../02-DESIGN/soulstream-idp/lifecycle.md) (D20–D24). **The M3
build lands next.**

**M5 ships — the embed seam**
([episode 0008](0056-soulfold-m5-the-embed-seam.md);
`specs/004-embed-seam/`): the public `embed.Run(ctx, Options)`
assembly (D29 pattern) with the AS-contract half the bundled story
needs — RFC 7591 DCR and the fixed token audience, both opt-in. The
compiler-proof consumer gate (`e2e/embedgate`, module path outside the
namespace) runs discovery → DCR → passkey sign-in → audience/
roles-bearing tokens; the M4 rig's fold half now rides the seam;
`soulfold serve` is the seam's first consumer; `authtest` went public
for consumers' own gates. **The public-door path is complete on the
fold's side — M3 (lifecycle + bootstrap research) is all that
remains.**

**M4 ships — the fold joins the fleet**
([episode 0007](0054-soulfold-m4-the-fold-in-the-fleet.md);
`specs/003-fold-in-the-fleet/`, the `e2e/` rig module): a browser
user's passkey sign-in becomes a NATS admission — fold-issued access
token + public sentinel through soulidentity's auth callout at its
published tag v0.1.0, the role value resolving against the declared
binding, scope server-enforced, refusals audited; and the identical
rig passes with an Entra-shaped stub (constitution II demonstrated).
The fold's tokens speak Entra's claim vocabulary (oid /
preferred_username / roles) because the seam's verifier keys subjects
by oid. **Soulnode's public door is unblocked; M5 (embed) next.**

**M2 ships — the refusal becomes behavior**
([episode 0006](0053-soulfold-m2-passkeys.md); `specs/002-passkeys/`):
WebAuthn ceremonies from go-webauthn are the only way through — the
username-form login is gone, first touch registers a passkey, second
touch asserts it, all proven in `make test` by a virtual authenticator
doing real ES256 ceremonies through the full HTTP flow (stock go-oidc
RP, restart mid-flow, forged POSTs inert). The D14 origin matrix
refuses all four foreign shapes server-side; no credential secret in
the store, positive-control-verified. Pending human act: the
physical-authenticator runbook (feature quickstart). First-touch
enrollment is the loud interim until M3's bootstrap research. **Next:
M4 — the callout admission proof.**

**M1 ships — the OP skeleton is real**
([episode 0005](0052-soulfold-m1-the-op-skeleton.md);
`specs/001-op-skeleton/`): discovery, JWKS, and the code+PKCE flow on
the sealed store, all measured in `make test` against an embedded
nats-server — a stock go-oidc RP signs in end to end (tokens verifying
against published JWKS, subject the seeded user), a full restart
between the login POST and the token exchange is invisible, forged
POSTs (missing/wrong CSRF, foreign Origin) leave the store's revision
unmoved, a rotation runs under a never-restarted verifier at 61/61
verifications, and the custody scan stays clean with its positive
control. The page inventory is exactly {login, error}. **Next: M2
(passkeys), then M4 before M3** per the operator's public-door
priority.

**The envelope is decided — M1 is unblocked**
([episode 0004](0051-soulfold-kv-encryption-at-rest.md)): the last
M1-gating research concluded with all four pre-registered bars passing
[measured] — the store's proven mechanics survive xkey sealing
byte-for-byte (restart 7/7, matrix 24/24, CAS 8,000/8,000,
redemption 100/100), a stopped-store scan with positive controls
recovers nothing (and forced the D6 amendment digesting the username
index key), filestore encryption was measured leaving the NATS surface
plaintext (the decisive asymmetry), and the cost landed at +44 B / ~57 µs
per record and +1.19 ms on a real end-to-end sign-in. The store design
grew D16–D19; every research gate M1 names is now concluded to design —
the build is next.

**The sign-in surface is decided**
([episode 0003](0046-soulfold-session-and-ui-shape.md)): the second M1-gating
research concluded with all four bars passing [measured] — the whole
flow is two server-rendered pages and zero JavaScript (15/15 rig
checks with a stock RP), a process restart mid-flow is invisible, the
CSRF posture rejects every forged shape with zero state change, and
the WebAuthn origin/RP-ID matrix (10/10) fixed the naming rules:
renaming the public host invalidates enrolled passkeys. Second design
doc: [session-and-ui](../02-DESIGN/soulstream-idp/session-and-ui.md) (D9–D15). M1's
two named research topics are done; one new topic was opened at the
operator's direction before the build starts — KV entry protection at
rest (xkeys), since the record envelope sits inside the store-shape
one-way door.

**The store is decided** ([episode 0002](0043-soulfold-kv-schema-and-key-lifecycle.md)):
M1's gating research concluded with all four pre-registered bars passing
[measured] — restart round-trip 6/6 byte-identical, additive decode
25/25 (and the cross-version RMW trap measured into design rule D3),
CAS with zero lost updates and exactly-once code redemption 100/100,
and a full JWKS rotation with 0 failures in 466 verifications under a
never-restarted go-oidc verifier. The fold's first design doc landed:
[store-and-key-lifecycle](../02-DESIGN/soulstream-idp/store-and-key-lifecycle.md)
(D1–D8, including RS256 for the seam). M1's remaining research is the
session and UI shape; the build follows it.

**Genesis — the fold is founded from a refusal**
([episode 0001](0041-soulfold-genesis-the-fold.md)): soulidentity's default-IdP
question (its journey 0019) resolved with the identity plane's refusal
holding — the ecosystem's default IAM is this sibling project: a
NATS-native (JetStream KV), embeddable, passkey-first OIDC provider that
consumers reach exclusively through standard OIDC, indistinguishable
from Entra by design. Build-vs-adopt was examined (pocket-id is an
application, not an embeddable library; no embeddable Go passkey-IdP
library exists [judgment]); the constitution (1.0.0) fixes the founding
constraints — passkeys, not passwords; indistinguishable by design —
and the roadmap sequences M1 (the OP skeleton) behind its KV-schema and
key-lifecycle research. No product code exists yet.

### soulstream-shell — the shell (as of 2026-08-13; named soulhelm until episode 0069)

**Founded, shipped, and composed in one day** ([episode
0067](0067-soulhelm-founding-and-first-light.md); research graduated
the same morning, [episode 0066](0066-ecosystem-soulsystem-cockpit.md)):
**v0.1.0** is the observe surface (board, topics with earned signature
verdicts, storage, plane health) rendered in the cassette-light token
source over Datastar SSE, with fold sessions (RFC 7591 DCR +
code+PKCE, memory-only, each session admitting as itself through the
OIDC callout lane) and the first act — `work.open` signed as the
signed-in principal. The surface is closed until sign-in (an open
scope amendment: sessions moved into the founding release rather than
ship an unauthenticated realm viewer). The consumer-position e2e
(module path outside the namespace, upstreams at published tags) walks
the whole human ceremony in ~4 s inside `make test`. Soulnode composed
it the same day as `planes.helm` ([episode
0068](0068-soulnode-the-helm-plane.md)).

**The direction changed the same day** — the operator reframed the
component as **the shell**: a pure modular frame containing zero module
logic, agnostic from Soulstream by contract, every human surface — the
observe core, the fold's administration, the agent designer/manager to
come — a module plugging in through one exported contract, with the
shell providing composition (registration, activation, navigation,
sessions, cross-linking). Research topic:
[`shell-module-contract`](../01-RESEARCH/shell-module-contract/README.md),
four pre-registered bars. And the ecosystem naming re-centering renames
the component **soulstream-shell** ([episode
0069](0069-ecosystem-one-name-soulstream.md); bare *soulshell* was
decided and superseded the same day), executed in the sweep ([episode
0070](0070-ecosystem-the-rename-sweep.md), v0.2.0).

**The focus (2026-08-14,** [episode
0071](0071-ecosystem-the-focus.md)**):** participation enters scope —
the mission is the **usable cockpit**: view topics, collaborate
directly (post, reply, comment, open topics), mention notifications,
riding the backend-held session admission already measured. The
shell-module-contract bars ride that build, with the collaboration
module as Bar 1's second module. **The composer landed the same
morning** ([episode 0072](0072-shell-the-composer.md)) — post and
anchored reply as the session's own principal, two pre-existing
rendering bugs fixed — and **the chat shape the same day** ([episode
0073](0073-shell-the-chat-shape.md)): conversations rail, own-message
bubbles, threaded replies, docked composer, plus the `make screens`
helper whose first real screenshots caught the retired *soulsystem*
wordmark. **The screen review landed the same morning** ([episode
0074](0074-shell-the-spine-and-the-details.md)): the icon rail with
Home ("Your realm at a glance"), the capped centered column, and the
details panel — People, Status, Waiting on — all verified on live
screenshots. **Next:** mention notifications; the "You" display nit;
the module re-homing spike (Bar 1).

## Episode index

| # | Component | Episode |
|---|---|---|
| 0001 | soulrealm | [Genesis: soulrealm gets an HQ (2026-07-22)](0001-soulrealm-genesis.md) |
| 0002 | soulrealm | [The substrate decision: a from-scratch, NEX-influenced runtime (2026-07-22)](0002-soulrealm-the-substrate-decision.md) |
| 0003 | soulrealm | [Soulstream-only scope: the platform waits (2026-07-22)](0003-soulrealm-soulstream-only-scope.md) |
| 0004 | soulrealm | [The first agent runs (2026-07-22)](0004-soulrealm-the-first-agent-runs.md) |
| 0005 | soulrealm | [A tool answers (2026-07-22)](0005-soulrealm-a-tool-answers.md) |
| 0006 | soulrealm | [HQ alignment: the lint gets built (2026-07-24)](0006-soulrealm-hq-alignment.md) |
| 0007 | soulstream | [Genesis to v0.3: the protocol, and the library that proves it (2026-07-11 → 2026-07-24)](0007-soulstream-genesis-and-the-reference-library.md) |
| 0008 | soulstream | [Adopting the hq way: the process gets a constitution (2026-07-24)](0008-soulstream-adopting-the-hq-way.md) |
| 0009 | soulstream | [The memory convention: the realm learns to be asked (2026-07-25)](0009-soulstream-memory-convention-and-exhibits.md) |
| 0010 | soulstream | [Provisioning byte limits: the strict landlord gets a one-command realm (2026-07-27)](0010-soulstream-provisioning-byte-limits.md) |
| 0011 | soulstream | [Sealed topics survive the substrate: four bars, one encoding amendment (2026-07-27)](0011-soulstream-sealed-topics.md) |
| 0012 | soulidentity | [Genesis: the design thread and the walking skeleton (2026-07-28)](0012-soulidentity-genesis-and-the-walking-skeleton.md) |
| 0013 | soulidentity | [The identity-plane re-centering (2026-07-28)](0013-soulidentity-the-identity-plane-re-centering.md) |
| 0014 | soulidentity | [NATS-only and the connection ladder (2026-07-28)](0014-soulidentity-nats-only-and-the-connection-ladder.md) |
| 0015 | soulidentity | [The first-key story: a local file, named honestly (2026-07-28)](0015-soulidentity-first-key-story.md) |
| 0016 | soulidentity | [The NATS-surface design: the principal is the subject (2026-07-28)](0016-soulidentity-the-nats-surface-design.md) |
| 0017 | soulidentity | [Design review: seeds from the environment, no v1, D15 taught back (2026-07-28)](0017-soulidentity-design-review-amendments.md) |
| 0018 | soulidentity | [M3: the NATS-native rebuild ships (2026-07-28)](0018-soulidentity-m3-the-nats-native-rebuild.md) |
| 0019 | soulidentity | [The sentinel-credential flow: URL + token is enough (2026-07-28)](0019-soulidentity-sentinel-credential-flow.md) |
| 0020 | soulrealm | [A second wall: the microsandbox backend (2026-07-28)](0020-soulrealm-a-second-wall.md) |
| 0021 | soulidentity | [The claims-mapping shape: one pipeline, policy never in the credential store (2026-07-28)](0021-soulidentity-claims-mapping-shape.md) |
| 0022 | soulidentity | [M4: auth callout ships, the front door opens (2026-07-28)](0022-soulidentity-m4-auth-callout-ships.md) |
| 0023 | soulidentity | [The shared subject prefix: one namespace for the ecosystem (2026-07-28)](0023-soulidentity-the-shared-subject-prefix.md) |
| 0024 | soulrealm | [A third wall on rented ground: Kubernetes as a backend (2026-07-28 → 2026-07-29)](0024-soulrealm-kubernetes-backend.md) |
| 0025 | soulidentity | [The Entra lane: role == team, no mapping store (2026-07-29)](0025-soulidentity-entra-role-claim-lane.md) |
| 0026 | soulstream | [The signer seam: signing learns to be delegated (2026-07-29)](0026-soulstream-the-signer-seam.md) |
| 0027 | soulstream | [DX hardening: the seam's two sharp edges, and the cycle guard (2026-07-29)](0027-soulstream-dx-hardening-and-the-cycle-guard.md) |
| 0028 | soulrealm | [A third wall lands: the Kubernetes backend ships (2026-07-29)](0028-soulrealm-a-third-wall-lands.md) |
| 0029 | soulidentity | [The registry dissolves: authorization in the ACLs and the bindings (2026-07-29)](0029-soulidentity-the-registry-dissolves.md) |
| 0030 | soulidentity | [The cross-service proof: the seam carries a real record (2026-07-29)](0030-soulidentity-the-cross-service-proof.md) |
| 0031 | soulidentity | [The vault is the directory: ephemeral users, keys on first touch (2026-07-29)](0031-soulidentity-the-vault-is-the-directory.md) |
| 0032 | soulidentity | [One noun: persona (2026-07-29)](0032-soulidentity-one-noun-persona.md) |
| 0033 | soulrealm | [Fleet: the log nominates, evidence vetoes, the log decides (2026-07-29 → 2026-07-31)](0033-soulrealm-fleet.md) |
| 0034 | soulnode | [Genesis: SoulNode gets an HQ (2026-07-31)](0034-soulnode-genesis.md) |
| 0035 | soulidentity | [Role selection by declared name: the ephemeral mint op (2026-07-31)](0035-soulidentity-role-selection-by-name.md) |
| 0036 | soulidentity | [The embed seam: the serve assembly becomes public (2026-08-01)](0036-soulidentity-the-embed-seam.md) |
| 0037 | soulrealm | [Pinned to the record: the soulstream replace drops (2026-08-01)](0037-soulrealm-pinned-to-the-record.md) |
| 0038 | soulstream | [The remote MCP node: a URL into the realm, proven on the BYON (2026-07-30 → 08-01)](0038-soulstream-remote-mcp-node.md) |
| 0039 | soulidentity | [Soulfold: the default IdP is a sibling, the refusal holds (2026-08-02)](0039-soulidentity-soulfold-the-default-idp.md) |
| 0040 | soulnode | [The composition gate: three bars PASS, the ecosystem opens its seams (2026-07-31 → 2026-08-02)](0040-soulnode-the-composition-gate.md) |
| 0041 | soulfold | [Genesis: the fold is founded from a refusal (2026-08-02)](0041-soulfold-genesis-the-fold.md) |
| 0042 | soulnode | [First boot is real: init and up land (2026-08-02)](0042-soulnode-first-boot-is-real.md) |
| 0043 | soulfold | [The store is decided: four bars, four passes (2026-08-02)](0043-soulfold-kv-schema-and-key-lifecycle.md) |
| 0044 | soulnode | [The realm remembers: provisioning and the memory plane land (2026-08-02)](0044-soulnode-the-realm-remembers.md) |
| 0045 | soulnode | [An agent runs: Phase 1 is complete (2026-08-02)](0045-soulnode-an-agent-runs.md) |
| 0046 | soulfold | [Two pages, zero scripts, and a name that becomes a door (2026-08-02)](0046-soulfold-session-and-ui-shape.md) |
| 0047 | soulstream | [The remote MCP node, built: the door that holds nothing (2026-08-02)](0047-soulstream-remote-mcp-node-built.md) |
| 0048 | soulstream | [The node becomes consumable: the replace drops (2026-08-02)](0048-soulstream-the-node-becomes-consumable.md) |
| 0049 | soulnode | [The door opens: the MCP front door lands (2026-08-02)](0049-soulnode-the-door-opens.md) |
| 0050 | ecosystem | [One hq: the five headquarters merge](0050-ecosystem-one-hq.md) |
| 0051 | soulfold | [The envelope is decided: xkeys seal the store (2026-08-02)](0051-soulfold-kv-encryption-at-rest.md) |
| 0052 | soulfold | [M1 ships: the OP skeleton on the sealed store (2026-08-02)](0052-soulfold-m1-the-op-skeleton.md) |
| 0053 | soulfold | [M2 ships: passkeys, the refusal becomes behavior (2026-08-02)](0053-soulfold-m2-passkeys.md) |
| 0054 | soulfold | [M4 ships: the fold joins the fleet (2026-08-02)](0054-soulfold-m4-the-fold-in-the-fleet.md) |
| 0055 | soulnode | [The public door opens (2026-08-03)](0055-soulnode-the-public-door.md) |
| 0056 | soulfold | [M5 ships: the fold learns to live in your house (2026-08-03)](0056-soulfold-m5-the-embed-seam.md) |
| 0057 | soulnode | [The folded realm: one binary, real people (2026-08-03)](0057-soulnode-the-folded-realm.md) |
| 0058 | soulnode | [v0.1.0: the house gets a shipping label (2026-08-03)](0058-soulnode-the-release-pipeline.md) |
| 0059 | soulfold | [The bootstrap story: invitation is the only door (2026-08-03)](0059-soulfold-bootstrap-story.md) |
| 0060 | soulfold | [M3 ships: the fold is complete (2026-08-03)](0060-soulfold-m3-the-lifecycle.md) |
| 0061 | soulfold | [The admin console: a browser and your passkey (2026-08-03)](0061-soulfold-the-admin-console.md) |
| 0062 | soulnode | [The front of house: URLs, the console, no collision (2026-08-03)](0062-soulnode-the-front-of-house.md) |
| 0063 | soulfold | [The console gets a face (2026-08-03)](0063-soulfold-the-console-gets-a-face.md) |
| 0064 | ecosystem | [The platform turn: tenancy, guardrails, and eight decisions (2026-08-04)](0064-ecosystem-the-platform-turn.md) |
| 0065 | ecosystem | [The ecosystem goes fair-code (2026-08-08)](0065-ecosystem-fair-code.md) |
| 0066 | ecosystem | [The helm: the cockpit earns its design (2026-08-13)](0066-ecosystem-soulsystem-cockpit.md) |
| 0067 | soulhelm | [Founding and first light: the helm is real (2026-08-13)](0067-soulhelm-founding-and-first-light.md) |
| 0068 | soulnode | [The helm plane: the cockpit joins the bundle (2026-08-13)](0068-soulnode-the-helm-plane.md) |
| 0069 | ecosystem | [One name: Soulstream (2026-08-13)](0069-ecosystem-one-name-soulstream.md) |
| 0070 | ecosystem | [The rename sweep: eight repos, one evening (2026-08-13)](0070-ecosystem-the-rename-sweep.md) |
| 0071 | ecosystem | [The focus: a product, not a platform (2026-08-14)](0071-ecosystem-the-focus.md) |
| 0072 | shell | [The composer: the shell stops being a window (2026-08-14)](0072-shell-the-composer.md) |
| 0073 | shell | [The chat shape: a rail, a conversation, a docked composer (2026-08-14)](0073-shell-the-chat-shape.md) |
| 0074 | shell | [The spine and the details: the operator's screen review lands (2026-08-14)](0074-shell-the-spine-and-the-details.md) |

## The naming map (2026-08-13)

The ecosystem naming re-centering (episodes
[0069](0069-ecosystem-one-name-soulstream.md) /
[0070](0070-ecosystem-the-rename-sweep.md)) renamed every project under
the one brand. Episodes 0001–0069 keep their filenames and component
tags — the record is append-only; this map resolves them. Future
episodes use the short tag.

| Old name (tag in episodes ≤ 0069) | Repository today | Episode tag now |
|---|---|---|
| soulstream (the record library) | soulstream-core | core |
| soulrealm | soulstream-workloads | workloads |
| soulidentity | soulstream-identity | identity |
| soulnode | soulstream (the product) | soulstream |
| soulfold | soulstream-idp | idp |
| soulhelm | soulstream-shell | shell |
| soulstream/node (nested module) | soulstream-mcp | mcp |
| soulstream-archivist | soulstream-archivist (unchanged) | — |
| — | soulstream-cli (to be founded) | cli |

## Pre-merge numbering map

Old (per-project) episode → new shared number. Frozen specs, old commit messages, and archived documents cite the old numbers.

| Component | Old # | New # | Episode |
|---|---|---|---|
| soulrealm | 0001 | 0001 | [0001-soulrealm-genesis.md](0001-soulrealm-genesis.md) |
| soulrealm | 0002 | 0002 | [0002-soulrealm-the-substrate-decision.md](0002-soulrealm-the-substrate-decision.md) |
| soulrealm | 0003 | 0003 | [0003-soulrealm-soulstream-only-scope.md](0003-soulrealm-soulstream-only-scope.md) |
| soulrealm | 0004 | 0004 | [0004-soulrealm-the-first-agent-runs.md](0004-soulrealm-the-first-agent-runs.md) |
| soulrealm | 0005 | 0005 | [0005-soulrealm-a-tool-answers.md](0005-soulrealm-a-tool-answers.md) |
| soulrealm | 0006 | 0006 | [0006-soulrealm-hq-alignment.md](0006-soulrealm-hq-alignment.md) |
| soulstream | 0001 | 0007 | [0007-soulstream-genesis-and-the-reference-library.md](0007-soulstream-genesis-and-the-reference-library.md) |
| soulstream | 0002 | 0008 | [0008-soulstream-adopting-the-hq-way.md](0008-soulstream-adopting-the-hq-way.md) |
| soulstream | 0003 | 0009 | [0009-soulstream-memory-convention-and-exhibits.md](0009-soulstream-memory-convention-and-exhibits.md) |
| soulstream | 0004 | 0010 | [0010-soulstream-provisioning-byte-limits.md](0010-soulstream-provisioning-byte-limits.md) |
| soulstream | 0005 | 0011 | [0011-soulstream-sealed-topics.md](0011-soulstream-sealed-topics.md) |
| soulidentity | 0001 | 0012 | [0012-soulidentity-genesis-and-the-walking-skeleton.md](0012-soulidentity-genesis-and-the-walking-skeleton.md) |
| soulidentity | 0002 | 0013 | [0013-soulidentity-the-identity-plane-re-centering.md](0013-soulidentity-the-identity-plane-re-centering.md) |
| soulidentity | 0003 | 0014 | [0014-soulidentity-nats-only-and-the-connection-ladder.md](0014-soulidentity-nats-only-and-the-connection-ladder.md) |
| soulidentity | 0004 | 0015 | [0015-soulidentity-first-key-story.md](0015-soulidentity-first-key-story.md) |
| soulidentity | 0005 | 0016 | [0016-soulidentity-the-nats-surface-design.md](0016-soulidentity-the-nats-surface-design.md) |
| soulidentity | 0006 | 0017 | [0017-soulidentity-design-review-amendments.md](0017-soulidentity-design-review-amendments.md) |
| soulidentity | 0007 | 0018 | [0018-soulidentity-m3-the-nats-native-rebuild.md](0018-soulidentity-m3-the-nats-native-rebuild.md) |
| soulidentity | 0008 | 0019 | [0019-soulidentity-sentinel-credential-flow.md](0019-soulidentity-sentinel-credential-flow.md) |
| soulrealm | 0007 | 0020 | [0020-soulrealm-a-second-wall.md](0020-soulrealm-a-second-wall.md) |
| soulidentity | 0009 | 0021 | [0021-soulidentity-claims-mapping-shape.md](0021-soulidentity-claims-mapping-shape.md) |
| soulidentity | 0010 | 0022 | [0022-soulidentity-m4-auth-callout-ships.md](0022-soulidentity-m4-auth-callout-ships.md) |
| soulidentity | 0011 | 0023 | [0023-soulidentity-the-shared-subject-prefix.md](0023-soulidentity-the-shared-subject-prefix.md) |
| soulrealm | 0008 | 0024 | [0024-soulrealm-kubernetes-backend.md](0024-soulrealm-kubernetes-backend.md) |
| soulidentity | 0012 | 0025 | [0025-soulidentity-entra-role-claim-lane.md](0025-soulidentity-entra-role-claim-lane.md) |
| soulstream | 0006 | 0026 | [0026-soulstream-the-signer-seam.md](0026-soulstream-the-signer-seam.md) |
| soulstream | 0007 | 0027 | [0027-soulstream-dx-hardening-and-the-cycle-guard.md](0027-soulstream-dx-hardening-and-the-cycle-guard.md) |
| soulrealm | 0009 | 0028 | [0028-soulrealm-a-third-wall-lands.md](0028-soulrealm-a-third-wall-lands.md) |
| soulidentity | 0013 | 0029 | [0029-soulidentity-the-registry-dissolves.md](0029-soulidentity-the-registry-dissolves.md) |
| soulidentity | 0014 | 0030 | [0030-soulidentity-the-cross-service-proof.md](0030-soulidentity-the-cross-service-proof.md) |
| soulidentity | 0015 | 0031 | [0031-soulidentity-the-vault-is-the-directory.md](0031-soulidentity-the-vault-is-the-directory.md) |
| soulidentity | 0016 | 0032 | [0032-soulidentity-one-noun-persona.md](0032-soulidentity-one-noun-persona.md) |
| soulrealm | 0010 | 0033 | [0033-soulrealm-fleet.md](0033-soulrealm-fleet.md) |
| soulnode | 0001 | 0034 | [0034-soulnode-genesis.md](0034-soulnode-genesis.md) |
| soulidentity | 0017 | 0035 | [0035-soulidentity-role-selection-by-name.md](0035-soulidentity-role-selection-by-name.md) |
| soulidentity | 0018 | 0036 | [0036-soulidentity-the-embed-seam.md](0036-soulidentity-the-embed-seam.md) |
| soulrealm | 0011 | 0037 | [0037-soulrealm-pinned-to-the-record.md](0037-soulrealm-pinned-to-the-record.md) |
| soulstream | 0008 | 0038 | [0038-soulstream-remote-mcp-node.md](0038-soulstream-remote-mcp-node.md) |
| soulidentity | 0019 | 0039 | [0039-soulidentity-soulfold-the-default-idp.md](0039-soulidentity-soulfold-the-default-idp.md) |
| soulnode | 0002 | 0040 | [0040-soulnode-the-composition-gate.md](0040-soulnode-the-composition-gate.md) |
| soulfold | 0001 | 0041 | [0041-soulfold-genesis-the-fold.md](0041-soulfold-genesis-the-fold.md) |
| soulnode | 0003 | 0042 | [0042-soulnode-first-boot-is-real.md](0042-soulnode-first-boot-is-real.md) |
| soulfold | 0002 | 0043 | [0043-soulfold-kv-schema-and-key-lifecycle.md](0043-soulfold-kv-schema-and-key-lifecycle.md) |
| soulnode | 0004 | 0044 | [0044-soulnode-the-realm-remembers.md](0044-soulnode-the-realm-remembers.md) |
| soulnode | 0005 | 0045 | [0045-soulnode-an-agent-runs.md](0045-soulnode-an-agent-runs.md) |
| soulfold | 0003 | 0046 | [0046-soulfold-session-and-ui-shape.md](0046-soulfold-session-and-ui-shape.md) |
| soulstream | 0009 | 0047 | [0047-soulstream-remote-mcp-node-built.md](0047-soulstream-remote-mcp-node-built.md) |
| soulstream | 0010 | 0048 | [0048-soulstream-the-node-becomes-consumable.md](0048-soulstream-the-node-becomes-consumable.md) |
| soulnode | 0006 | 0049 | [0049-soulnode-the-door-opens.md](0049-soulnode-the-door-opens.md) |
