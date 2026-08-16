# 0001 — The soulstream composition

**Status**: graduated from `single-binary-composition` (journey
[episode 0002](../../04-JOURNEY/0040-soulnode-the-composition-gate.md)), 2026-08-02.
Governs Phase 1 of the [roadmap](../../03-IMPLEMENTATION/ROADMAP.md); the
front door (Phase 2) is seamed here but designed later.

Validation tags per the [design README](README.md): **[V]** validated by
the composition rig or a delivered upstream surface, **[D]** specified but
not yet run in soulstream itself, **[O]** interface fixed, internals open.

## 1. The shape [V]

One binary, one process, five planes — each a Go library consumed through
its public, tagged surface (constitution I), each talking to NATS over an
**ordinary client connection, loopback by default** (constitution III as
ratified). There is no in-process transport anywhere: the bundle is the
deployment where every URL happens to be `127.0.0.1`, which is exactly
what makes decomposition configuration rather than architecture.

| Plane | Surface consumed | Runs as |
|---|---|---|
| Server | `nats-server/v2/server` (embedded, operator mode, JetStream) | goroutine owning the listener |
| Identity | `soulstream-identity/embed.Run` + `soulstream-identity/client` | goroutine on its two conns |
| Memory | `soulstream-archivist/keeper` + `archive`, `soulstream/topic.RespondMemory` | goroutine on a realm client |
| Runtime | `soulstream-workloads/{runner,backend/native,minter,declaration}` + `soulstream/{realm,topic}` | goroutine supervising child processes |
| Front door | soulstream MCP surface (Phase 2 — see §8) | goroutine owning the HTTP listener |

Workloads never run in the process: the runtime plane launches them
through soulstream-workloads's backends (native first), and a workload death is an
op on the topic, never a node crash.

## 2. Configuration: planes are declared, connections are URLs [D]

Every plane carries the same connection block:

```
{ enabled: bool, url: string (default nats://127.0.0.1:<port>),
  creds: path (default: minted into the state dir) }
```

- **All defaults** → the bundle: embedded server enabled, every plane on
  loopback with state-dir creds. (As-built from 002: `config.json`
  carries `listen`, `realm` — fixed at founding — and
  `planes.memory.enabled`; each block gains `url`/`creds` when a plane
  can actually be pointed elsewhere.)
- **BYO NATS** → embedded server `enabled: false`, every URL points at
  the user's server. The ceremony's account half runs against that server
  (see §4's [O]).
- **Remote plane** → that plane `enabled: false` here, enabled on the
  other machine with a non-loopback URL. Same binary, same config schema,
  both sides.

A plane MUST NOT behave differently because its URL is or is not
loopback. The embedded server binds `127.0.0.1` only, port configurable
(default the NATS conventional port; refusing to start on a bind conflict
with a message naming the config key).

**Planes are named by function, and pre-v1 renames are clean breaks
[D — operator direction, 2026-08-15/16].** The config keys are
`planes.signin` (the bundled OpenID sign-in service, soulstream-idp)
and `planes.mcp` (the MCP endpoint for assistants); the flags are
`--signin-listen` / `--mcp-listen`; the state dir env var is
`SOULSTREAM_STATE`; founds mint the sign-in plane's NATS artifacts as
`users/signin.creds` (user `signin`) with plane state under
`<state>/signin/`. There are **no compatibility shims before v1**: no
alias keys, no alias flags, no path fallbacks — one schema, one code
path. A realm founded under the byname-era spellings is **refused by
name**, with the hand-migration spelled out in the refusal (rename the
config keys `door→mcp`, `fold→signin`; `mv users/fold.creds
users/signin.creds`; `mv fold/ signin/`) — an honest break, never a
silent misread. The sign-in plane, when embedded here, serves its
admin **API** and not its HTML console (soulstream-idp design D31):
the product's administration surface is the shell.

## 3. The embedded server [V]

Provisioned entirely in code — `server.Options` + `MemAccResolver`, no
config file, no `nsc` [measured, the rig]: `TrustedKeys` = the operator,
`SystemAccount` = SYS, resolver preloaded from the persisted account JWTs,
JetStream on `<state>/jetstream`. Boot order: server ready → identity
plane → realm provisioning → memory + runtime planes → front door.

## 4. The ceremony and the state directory [V for generation, D for persistence]

`soulstream init` generates, in dependency order, the inventory the research
enumerated and the rig proved executable from an empty directory
[measured]:

1. **Operator nkey** — the trust root (`TrustedKeys`).
2. **SYS account** nkey + JWT (grants `$SYS.REQ.USER.INFO` — the
   server-asserted whoami everything downstream leans on).
3. **AUTH account** nkey + JWT: `EnableExternalAuthorization(<issuer
   user>)`, `Authorization.AllowedAccounts` = the realm account,
   `Authorization.XKey` = the callout curve key; plus its **signing key**
   (vault name `auth/issuer` — signs admitted users and the sentinel).
4. **Realm account** nkey + JWT: JetStream unlimited locally, plus a
   **plain workload signing key** (`SigningKeys.Add` — the runtime
   minter's key; per-workload permissions ride in the JWTs it signs, a
   scoped key rejects them; as-built in 003) and the
   **scoped signing key** whose `jwt.UserScope` template *is* the admitted
   persona's permission set — `identity.{{account-subject()}}.{{name()}}.sign.record`
   / `.keys.public`, `identity.status` / `.xkey`, `SOULSTREAM.>`,
   `$JS.API.>`, `$KV.>`, `$O.>`, `$SYS.REQ.USER.INFO` (pub);
   `_INBOX.>`, `SOULSTREAM.>` (sub).
5. **Bypass-lane users**: AUTH issuer user; realm service, ops, and
   archivist users (account-key signed). These never leave the state
   dir. (The archivist's entry is transport only — its *persona* signing
   key is vault-held, materialized on first touch; as-built in 002.)
6. **Curve keys**: callout xkey (public in the AUTH JWT, seed to the
   issuer), vault first key, service surface key.
7. **Buckets** `SOULIDENTITY_VAULT` + `SOULIDENTITY_TOKENS`; **realm
   provisioning** (`realm.NewClient` + `Provision` — stream, notify,
   personas, objects).
8. **Founding acts through the public `client` over the node's own
   loopback connection** [measured — no in-process admin API exists
   upstream, and none is wanted]: import the realm scoped signing key and
   the AUTH signing key into the vault, mint the **sentinel** (public by
   design), create the **first API token** — the one secret `init` prints.

**Persistence and idempotence** [D — the delta beyond the rig]: seeds,
account JWTs, sentinel, and config persist under the state dir (`0700`
dirs, `0600` files); `init` on a non-empty state dir verifies and reports
— it MUST NOT regenerate a trust root that JetStream state already
depends on. Moving the realm to a new machine is copying the state dir
(vision: day 2). Custody follows soulstream-identity's D13: raw seeds on the
node's own disk are the accepted trust class, stated without euphemism.

The BYO-NATS ceremony subset — formerly this section's [O] — is
resolved by design [`0003-byo-nats.md`](0003-byo-nats.md)
(2026-08-16): steps 1–2 stay the substrate's permanently, steps 3–4
become the per-flavour account half (self-hosted kit or Synadia Cloud
API), steps 5–8 regroup into local material and the wire half, both
soulstream's own; no operator or account master key ever travels.

## 5. Admission [V]

Sentinel creds + `sit_` token; the callout issuer mints a TTL-bounded
scoped user; the principal is server-asserted (the expanded `sign.record`
grant names persona and account — no client-claimed identity anywhere).
Refusals (garbage, revoked) land in the audit log. All measured through
the public surfaces end to end. Revocation propagation bound = token TTL,
inherited from soulstream-identity's D22.

## 6. Plane wiring [V for identity/memory, D for runtime]

- **Identity**: `embed.Run(ctx, Options)` on two loopback conns (service
  account, AUTH account); options straight from the state dir. `Run`
  returned ⇒ surface silent (upstream contract, its journey 0018).
- **Memory**: `archive.Open(<state>/archive)` → `keeper.Run` +
  `keeper.Witness` + `topic.RespondMemory` on a realm client
  (`realm.NewClient` over loopback, archivist persona, signer from the
  identity plane's `client.PersonaSigner`).
- **Runtime** [V as-built in 003]: soulstream-workloads public packages, native
  backend, workloads as declarations — **invocation-scoped**: `soulstream
  workload start` supervises one declared workload as persona `runner`
  (transport creds from the ceremony, signer vault-held), minting with
  the ceremony's *plain* workload signing key (a scoped key rejects the
  minter's carried permissions — the two-keys split). The long-running
  node supervisor (claim-race placement, sweeper) is soulstream-workloads's own
  unbuilt Fleet milestone — soulstream consumes it when it lands upstream
  and MUST NOT invent one here (constitution I).

## 7. Shutdown and failure [D]

One signal context fans out to every plane; planes drain their
subscriptions, the runtime stops workloads through backend handles, the
server shuts down last. A plane failure is surfaced and named (log +
non-zero exit if fatal at boot; runtime restarts of planes are Phase 1's
[O] — default: fail loud, no silent restarts).

## 8. The front door [V as-built in 004]

MCP over streamable HTTP, static-bearer admission through the callout,
per-user pooled loopback connections, corpse eviction on dead pools —
landed upstream as soulstream 018 (v0.7.0, `soulstream/node`) and wired
as the door plane: `planes.door {enabled, listen}` (loopback,
`127.0.0.1:8080` default), the node holding the listener, upstream's
handler mounted, the existing sentinel, the audit logger; local mode
only. The door custodies nothing and consumes the same admission lane
as any client. Public mode (`public_url`, `auth_issuer` — the OAuth
story) joins the block additively when soulstream-idp exists upstream; HTTPS
today is a fronting concern (`tailscale serve`), with Phase 3's tsnet
behind its own gate.

## 9. Acceptance criteria (Phase 1, made precise per feature in specs/)

- **M1.1 — the server and the identity plane**: on an empty state dir,
  `soulstream init && soulstream up` reaches: sealed surface answers
  `status`; sentinel + printed token admits with the persona scoped to
  its own prefix; garbage and revoked tokens refuse with audited
  refusals; `init` re-run is a verified no-op; every artifact of §4 is on
  disk with the stated modes. Zero manual steps, zero external binaries.
- **M1.2 — the realm joins**: realm provisioned on the embedded server;
  the archivist keeps ops and answers the memory convention, attributed
  to its persona with a vault-held key.
- **M1.3 — an agent runs**: a declared agent workload launches through
  the runtime plane (native backend), posts a turn attributed to its
  persona, and its lifecycle appears as work ops — the soulstream-workloads M1.1
  proof, re-run inside soulstream.

## Requirement language

MUST / MUST NOT / MAY per the [design README](README.md); values marked
*default* ship unless configuration overrides them.
