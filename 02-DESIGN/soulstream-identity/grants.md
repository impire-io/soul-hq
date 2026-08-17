# soulstream-identity — outbound grants, the broker (D30–D34)

*The design graduated from research topic `outbound-identity-grants`
([episode 0104](../../04-JOURNEY/0104-ecosystem-outbound-identity-grants.md), Bars 1–3
measured on the real fold, plane, and node). The problem it answers: an
agent or shell calling a remote MCP server or third-party API today
would hold one credential in client config, collapsing every persona
into a single remote user. The inversion this design fixes in place:
**no outbound credential ever lives in agent, workload, or MCP-client
configuration** — outbound grants live behind the identity plane, where
the principal is server-asserted, so the only token any caller can
obtain is its own. Decisions continue soulstream-identity's numbering:
D30–D34.*

## D30 — The grants surface is an op family on the principal-scoped subject space

Outbound credentials are served by the identity plane as a `grants.*`
op family at `identity.<account>.<user>.grants.<op>` — the D14 subject
space, the D15 server-asserted principal, the D16 sealed envelope, the
D25 ACL-gated op tail, unchanged. The ops:

| Op suffix | Body → reply | Notes |
|---|---|---|
| `grants.link.start` | `{resource}` → `{authorize_url, link_id}` | begins the authorization-code + PKCE ceremony at the resource's AS |
| `grants.link.complete` | `{link_id, code_or_callback}` → `{}` | redeems the code; custody begins |
| `grants.access` | `{resource, on_behalf_of?, delegation?}` → `{access_token, expires_at}` | the workhorse; D32/D33 govern what it may return |
| `grants.list` | `{}` → `{grants: [{resource, linked_at, scopes}]}` | own grants only, by construction |
| `grants.revoke` | `{resource}` → `{}` | deletes custody; best-effort RFC 7009 upstream |

A represented user's scope template grows exactly one line —
`…{{account-subject()}}.{{name()}}.grants.>` — and the isolation
property is the transport's, not the broker's: a persona publishing to
another's grants subject is refused by the server before the service
hears it [measured: the request died as a publish permission violation
and the broker's delivery log showed the victim's subject reached it
exactly once]. Multi-token op suffixes need no parser change (the
dispatch `Join`). The client package mirrors the wire types in the same
commit, per the repo's standing rule.

**Reversal condition** (D15's, inherited and restated): a deployment
class whose permission templates cannot scope the grants op tail forces
a service-side gate keyed on declared configuration, as a new
D-decision.

## D31 — Grants are a second custody domain, not the key vault

The key vault's records are immutable by design — Create-only store
seam, closed kind enum, `derive()` requiring a public half — and a
rotating refresh token violates all three [measured, code map]. Grants
therefore get their **own sealed store**: a separate KV bucket (default
`SOULIDENTITY_GRANTS`), records sealed `xkv1` to the same first key
(D13/D17 custody unchanged), keyed `grant/<persona>/<resource>`, with
**CAS update semantics** (JetStream `Update(rev)`) — the one property
the key vault refuses and rotation demands. The callout token store is
not the home either: it writes unsealed digests, never secrets.

The rotation discipline, measured under `-race`: redeem at the
provider, CAS-write the successor, then return the access token; a CAS
loser hands out its still-valid access token and writes nothing — the
winner's successor is the line (8/8 concurrent accesses served, the
stored token still redeemable, both reuse-window regimes). Named
honestly: the crash window between redeem and CAS-write is where a
rotating line can still die; providers' reuse intervals exist for it,
and the design orders the writes to shrink it, not to zero. Access
tokens are cached in memory only, keyed (persona, resource, scopes),
never at rest.

**Reversal condition**: a provider whose rotation semantics cannot
survive the redeem-then-write order (observable: a measured lost line
with the discipline followed) forces a write-ahead intent record into
the grant store, as an amendment here.

## D32 — Article I holds: the broker returns derived credentials, never the custodied secret

A broker op that returns tokens collides head-on with
custody-without-possession unless the line is drawn explicitly, so it
is drawn here: **the refresh token never crosses the wire, in either
direction, ever** — there is no grants analogue of `ExportSeed`, and
unlinking is deletion, not export. What `grants.access` returns is the
provider's **short-lived access token**: a derived, expiring artifact,
the same custody class as a minted ephemeral JWT (D28) — custody
produces artifacts; it never surrenders the material that makes them.
The sealed envelope (D16) keeps even that artifact off the broker's
wire in plaintext. Every `grants.access` is one audit entry naming the
server-proven principal, the resource, and the decision.

## D33 — On-behalf-of is a minted, bounded delegation; the actor is the server-proven principal

An agent acting for a person is never ambient config. `grants.access`
with `on_behalf_of` requires a **delegation**: `{subject, actor,
resources, scopes, issued_at, expires_at}` plus a signature, honored
only when it verifies, is live, names the *server-proven caller* as
actor and the named subject as subject, and covers the requested
resource. Measured: the one allowed path against four refusal classes
(absent, expired, wrong caller presenting a stolen delegation,
out-of-bounds resource), every refusal audited **naming both
personas**; a stolen, validly-signed delegation refuses as an actor
mismatch because the actor claim is checked against the subject-token
principal, never against anything the caller says. Standing behind
every mint is a **consent record** — subject allows actor these
resources and scopes — revocable; revocation refuses the next mint and
rewrites no history.

The minting home, argued and decided as direction [judgment]:
**subject-signed** — the asking persona's own key signs the delegation,
and the broker verifies it from the `keys.public` directory (D26), so
no new trust root exists anywhere. The alternative (a privileged
runner co-signing, 0082's shape) died with the waker (the wrapper
mints nothing); restoring a mint lane for delegation would rebuild a
trust anchor D26 already provides. Carriage: the wrapper's per-run
overlay seam (the cut `MCPOverlay`) returns to hand the delegation
into the run's MCP env — a soulstream-workloads change, named in its
roadmap.

**Reversal condition**: a deployment where the subject cannot sign at
mention time (observable: a delegation needed for a subject whose
persona key cannot materialize — e.g. a non-signing client class,
recorded as an issue) reopens the co-signing mint lane as a new
D-decision.

## D34 — Four lanes, named honestly by what the remote supports

1. **Soulstream-to-soulstream.** The fold is the external AS for the
   remote realm's callout; admission is per-user with zero
   cross-attribution [measured, Bar 1 — the 0054/0055 compositions run
   multi-user]. No stored grants, no broker in the path. Two demands
   this lane places upstream, both measured: the fold's
   **token-lifetime knob** (revocation propagates in access-token
   `exp` + callout TTL; 1h is a package constant today — 0103's watch
   item has met its milestone) and, for several remotes from one
   session, **RFC 8693 token exchange** (the fold refuses the grant
   type; audience is deployment-fixed at startup). The exchange grant
   is an Article II argument the fold's roadmap owns — Entra itself
   speaks OBO, so indistinguishability is arguable.
2. **Third-party OAuth** (most SaaS APIs and remote MCP servers). The
   linking ceremony + D31 custody + D32 access [measured on the Dex
   stand-in; the real-provider confirmation is the named residue].
   The provider-facing HTTP machinery is genuinely new surface — no
   outbound HTTP exists in core or workloads — which is exactly why it
   lives behind the identity plane once, not in every agent.
3. **Token-exchange-capable enterprise IdP.** Where the deployment's
   IdP speaks RFC 8693, the broker gains an exchange backend and
   stores nothing — same surface, no custody. Preferred over lane 2
   whenever available.
4. **Single-credential APIs.** A remote that cannot tell users apart
   gets no pretense: the grant belongs to a realm/service persona, who
   may act through it is a soulstream authorization decision, and
   attribution lives in the audit. The design refuses to dress this up
   as user-scoped.

## Configuration surface

- The grants bucket name; the resource catalog — each resource:
  a name (legal slug), the AS issuer or API base, client registration
  (public + PKCE preferred), scopes. Declared configuration, no
  per-user rows (D26's spirit: linking is the user's own act on their
  own prefix).
- The scope-template line for the grants op tail (deployment duty,
  stated in docs beside the two existing shapes).
- Lane 3: per-issuer exchange configuration when present.

## Acceptance criteria (the build's gate, from the measured bars)

1. A persona obtains its own grant's access token; another persona's
   request for it is refused by the *server* (delivery-log proof), and
   the broker makes zero identity decisions [Bar 2 shape].
2. ≥3 provider-side rotations CAS-persist; a concurrent double-refresh
   under `-race` loses nothing; revocation deletes custody, refuses
   the next access, and best-effort revokes upstream [Bar 2 shape].
3. An on-behalf call with a valid delegation serves bounded access;
   absent/expired/stolen/out-of-bounds refuse; every on-behalf
   decision audits both personas; consent revocation refuses the next
   mint [Bar 3 shape].
4. The refresh token appears nowhere on any wire and nowhere at rest
   unsealed (positive-control grep, the D13 idiom) [D32].
5. One real third-party provider completes link → rotation → revoke
   [the Bar 2 residue, closing].
