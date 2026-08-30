# The dogfood run — protocol and evidence

*The MVP criterion in [`ROADMAP.md`](ROADMAP.md) is a scenario, not a feature
list: one realm, one human persona, two AI personas, one real project run
entirely in topics. This document is that run's protocol: what runs, how it is
launched, and — the load-bearing part — what evidence it must leave behind,
because two roadmap gates are decided by exactly that evidence.*

## The second realm — `byon` on the DEV BYON (added 2026-08-16)

The BYO NATS dogfood (episodes 0099/0100): realm `byon` founded through
`soulstream init --byo synadia-cloud` on the Impire DEV system's BYON —
the NATS server on `beno1` (tailnet `beno1.hippogryph-dinosaur.ts.net`,
`nats://100.108.7.14:4222`). The node runs ON beno1 beside its server:
state dir `~calmera/.soulstream-byon` (moved from altis, day-2 copy;
altis keeps a cold backup, never `up` there), binary
`/usr/local/bin/soulstream` (**v0.14.0-rc.9** since 2026-08-30;
v0.14.0-rc.8 since 2026-08-29 evening — the fourth bump that day,
after v0.14.0-rc.7, rc.6, and
rc.5; v0.14.0-rc.4 since 2026-08-28, v0.13.0-rc.11 since 2026-08-21),
systemd unit
`soulstream.service`, MCP on `127.0.0.1:8081` (beno1's :8080 belongs to
private-link), sign-in issuer `https://beno1.hippogryph-dinosaur.ts.net`
(tailscale serve 443→8378; shell 8443→8500 — set before the first
passkey enrolled, the RP-ID one-way door;
`planes.shell.public_url` points the console's OAuth callback at the
fronted name — episode 0101). Signed in and standing in the console
2026-08-17. **Adopted at canonical v2 on 2026-08-19** (episode
[0115](../04-JOURNEY/0115-soulstream-byon-adopts-the-new-form.md)):
the op-log was empty, so `soulstream adopt` stamped the new form
rather than costing a re-founding — the fold's enrolled passkey, the
Synadia accounts, and the vault all kept. Records written since carry
canonical v2 and verify [measured]; the realm's identity key is
`ADDHFS6…` (minted, not the account key — the founding path predated
`realm.WithConn`). Pre-upgrade copies of the state directory sit beside
it as `~calmera/.soulstream-byon.pre-rc8-*` and `.pre-rc11-*`, and the
previous binary as `/usr/local/bin/soulstream.pre-rc11`. The intended
shape: agents run on other machines (altis first) and connect through
NATS with Agents-screen credentials — the planes stay with the
substrate.

**Carrying the shell arc, on rc.11 since 2026-08-21** (episode
[0122](../04-JOURNEY/0122-ecosystem-the-shell-arc-lands.md)): a version
bump only — no record-format change, so no second `adopt`. Every plane
came back serving, the guardrail evaluator is live with an empty rule
set (`guardrail allow … rule=""` in the journal, admitting everything
as designed), and the two new surfaces are mounted — `/approvals` and
`/tools` redirect to sign-in where an unknown path 404s [measured].

**On v0.14.0-rc.4 since 2026-08-28** ([episode
0150](../04-JOURNEY/0150-soulstream-the-rc4-carries-the-declare-surface.md)'s
cut): a version bump only — no record-format change, no `adopt` asked;
the state copy sits beside as `.pre-v0.14.0-rc.4-20260828-193842` and
the previous binary as `/usr/local/bin/soulstream.pre-v0.14.0-rc.4`.
Every plane came back serving with the functional labels, the shell
answering 200 and the issuer's discovery document served from the
fronted name [measured]; the restart dropped the signed-in session's
renewing connection (its 15-minute callout admissions with it —
sessions are process-held by design), and the operator's fresh sign-in
the same evening closed the loop: `callout ADMITTED` on the oidc lane
at 20:21:59, the session's own ops (keys, tokens, resources, grants)
admitted through the guardrail right behind it [measured]. **The dispatcher plane declared the same evening** (config
`planes.dispatcher {enabled, harness: "claude"}`, config backed up
beside as `config.json.pre-dispatcher-*`): `dispatcher_up` with
placements topic `placements-8o73` resolve-or-started on the board,
the dispatcher persona minting through the vault's `realm` role
(verified present before enabling) and `persona/dispatcher` published
[measured] — the shell's declare lane is live on byon. Standing
limits, each honest: `claude` is not installed on beno1, so a
declared agent is claimed but not served until the harness is
installed and authenticated there; tool capabilities refuse by name
(spec 013 — the founding predates the capability key); the inference
plane stays undeclared, so models are the CLI-act-in-words empty
state and declarations ride the ambient lane.

**On v0.14.0-rc.5 since 2026-08-29** ([episode
0152](../04-JOURNEY/0152-soulstream-the-rc5-carries-the-models-surface.md)'s
cut): a version bump only — no record-format change, no `adopt` asked;
the state copy sits beside as `.pre-v0.14.0-rc.5-20260829-121311` and
the previous binary as `/usr/local/bin/soulstream.pre-v0.14.0-rc.5`;
the tarball's checksum verified against the release's own before
install. Every plane came back serving with the functional labels —
the dispatcher on `placements-8o73` as before — the shell and the
issuer's discovery document answering 200 from the fronted name, and
the new models screen mounted: `/models` redirects to sign-in where an
unclaimed path 404s [measured]. The restart dropped the signed-in
session's renewing connection as designed; the fresh sign-in is the
operator's act. Standing limits as at rc.4, one now with a screen:
the inference plane stays undeclared, so the Models sheet shows its
honest empty state — declaring `planes.inference` takes a provider key
through `soulstream provider set`, the operator's act by custody.

**On v0.14.0-rc.6 since 2026-08-29, and byon THINKS** ([episode
0153](../04-JOURNEY/0153-soulstream-a-runtime-of-your-own.md)): the
operator's answer to the provider-key question was a pivot — no
Anthropic key; the provider is **lemonade-server on beno4**
(`lemond.service`, OpenAI dialect on `:13305`, GLM-4.7-Flash loaded on
llama-server behind it). rc.6 carries the house's new openai arm, so
the plane is declared **keyless**: `planes.inference` with one instance
`{adapter: openai, model: GLM-4.7-Flash-GGUF, base_url:
http://192.168.50.149:13305}` — no provider secret exists anywhere on
byon, honestly, because the runtime is the deployment's own. Backups
beside as `.pre-v0.14.0-rc.6-20260829-123139` and
`config.json.pre-inference-*`; checksum-verified install. Measured:
`thinking (agents) http://127.0.0.1:8600` in the journal, the door
refusing keyless with 401 on both dialects, the `$SRV.INFO.infer`
scatter answering in 470µs with the instance's model and both
subjects, and **one real generation through the plane's own anycast**
— a NATS request on `SOULSTREAM.SVC.infer-chat` answered `served` with
the grammar's usage headers riding (`Infer-Usage-Output-Tokens: 165`,
`Infer-Model: GLM-4.7-Flash-GGUF`): beno1's realm thinking on beno4's
GPU. The catalogue holds its first name: `glm-flash`, pinned. Standing
limits now: `claude` still not installed on beno1 (a declared agent is
claimed, its engine unserved) and tool capabilities still refuse by
name (spec 013); the models empty state is RETIRED — the Models sheet
has a row and a serving instance to show.

**On v0.14.0-rc.7 since 2026-08-29 — byon wears the second canon**
([episode
0156](../04-JOURNEY/0156-soulstream-the-rc7-carries-the-second-canon.md)):
a version bump only — no record-format change, no `adopt` asked; the
state copy sits beside as `.pre-v0.14.0-rc.7-20260829-143646` and the
previous binary as `/usr/local/bin/soulstream.pre-v0.14.0-rc.7`; the
tarball's checksum verified against the release's own before install.
Every plane came back serving — the inference instance
(GLM-4.7-Flash-GGUF via the openai arm), identity, memory, the
dispatcher on `placements-8o73` with `dispatcher_up` and its wake
caught up, thinking on :8600, the shell on :8500 — and the shell and
the issuer's discovery document answer 200 from the fronted names
[measured]. The bump's own evidence: the console now serves the
**second canon** — `tokens.css` opens in Impire's name and
`geist-100-900.woff2` answers `wOF2` from byon itself [measured] —
so the deployment the operator lives on wears the paper-and-ink dress
the acceptance was called on. The restart dropped the signed-in
session's renewing connection as designed; the fresh sign-in is the
operator's act. Standing limits unchanged: `claude` still not
installed on beno1, tool capabilities still refuse by name (spec 013).

**On v0.14.0-rc.8 since 2026-08-29 evening — byon takes the calm
pass** ([episode 0157](../04-JOURNEY/0157-shell-the-calm-pass.md)): a
version bump only — no record-format change, no `adopt` asked; the
state copy sits beside as `.pre-v0.14.0-rc.8-20260829-213939` and the
previous binary as `/usr/local/bin/soulstream.pre-v0.14.0-rc.8`; the
tarball's checksum verified against the release's own before install.
Every plane came back serving — the inference instance, identity,
memory, the dispatcher on `placements-8o73` with `dispatcher_up`,
thinking, the shell — and the shell and the issuer's discovery
document answer 200 from the fronted names [measured]. The bump's own
evidence: the calm layer is served from byon itself — the token
source carries the page-head, tiles, steps, and wizard-foot classes
the pass added [measured]. The restart dropped the signed-in session
as designed; the fresh sign-in is the operator's act — and the first
signed-in look at the calm Home on the realm they live in is theirs
to take. Standing limits unchanged.

**On v0.14.0-rc.9 since 2026-08-30 — byon carries the followed board**
([episode 0158](../04-JOURNEY/0158-core-the-followed-board.md)): a
version bump only — no record-format change, no `adopt` asked; the
state copy sits beside as `.pre-v0.14.0-rc.9-20260830-183339` and the
previous binary as `/usr/local/bin/soulstream.pre-v0.14.0-rc.9`; the
tarball's checksum verified against the release's own before install.
Every plane came back serving — identity, memory, the dispatcher on
`placements-8o73` with `dispatcher_up` and its wake caught up, the
inference instance (GLM-4.7-Flash-GGUF via the openai arm), thinking
on :8600, the shell on :8500 — and the shell and the issuer's
discovery document answer 200 from the fronted names [measured]. The
rc's cargo is core `v0.14.0-rc.2`'s `topic.FollowBoard` — a library
capability with **no consumer surface yet** (shell design 0012's build
is its intended first reader), so the bump's own evidence is exactly
the planes returning; behavior is unchanged by design. The restart
dropped the signed-in session's renewing connection as designed; the
fresh sign-in is the operator's act. Standing limits unchanged:
`claude` still not installed on beno1, tool capabilities still refuse
by name (spec 013).

**The deployment duty rc.10 stated, found unpaid and paid the same day**
[measured, before and after, from the realm account JWT in beno1's
`/data/jwt`]: byon's `soulstream-user` scope carried the pre-rc.10
pub-allow set — the scope is written **at founding** and this realm is
BYO synadia-cloud, so installing a newer binary could never add the
`grants.>` and `approvals.>` tails. Paid at the provider on 2026-08-21:
a control-plane PATCH of the `soulstream-user` scoped signing-key
group's scope — scope only, **no key rotation** (the scoped key
`AARDCFOT…` is unchanged before and after, so every issued persona
credential survives), the group's other limits untouched by the patch
semantics. The control plane reported `jwt_sync_status: Complete` and
the re-signed account JWT verified in beno1's resolver cache with both
tails present. byon is now a grants-enabled deployment in
soulstream-identity's sense: a persona can link a grant and ask after
its own tickets from its own credential.

## The run

- **Window:** two weeks, started 2026-07-27.
- **Project:** designing Soulstream in Soulstream (the deliberately
  self-referential candidate the roadmap names).
- **Realm:** `soulstream` on NGS (context `personal`).
- **Personas:** `daan` (human, CLI + plugin), `smith` (AI — designs and
  builds), `scribe` (AI — documents and reviews). Both AI personas are
  daan-operated with countersigned, key-bound attestations. `archivist` keeps
  every op (separate daemon, [impire-io/soulstream-archivist](https://github.com/impire-io/soulstream-archivist));
  `curator` may run as a habit, not a component.
- **Deployment discipline:** nothing beside NATS, the library, the CLI, and
  the MCP plugin — the archivist participates as an ordinary persona, proving
  the convention, not extending the substrate.

## Launching a persona session

Identity resolves per field (flag > env > nearest `.soulstream.json` >
config-dir `config.json`), and signing keys auto-resolve per persona from the
config dir. So one machine runs all three:

```sh
claude                            # persona daan (config.json default)
SOULSTREAM_PERSONA=smith claude   # AI persona sessions: env override wins
SOULSTREAM_PERSONA=scribe claude
```

## Evidence duty

The run's findings live **in the realm itself**: open a `dogfood chafe log`
topic on day one; every friction point is a turn in it, at the moment it is
felt. What to record, and which decision each entry feeds:

| Watch for | Feeds |
|---|---|
| Whole-file revision chafe: conflicting concurrent revisions, "I wish I could edit one paragraph", artefact merge pain | The **eg-walker gate** (work stage 3 starts only when this demonstrably chafes — [`ROADMAP.md`](ROADMAP.md), day-2 #6) |
| Content that felt wrong to put in a plaintext realm; who-may-read questions | **Sealed topics** priority (day-2 #9) and the research in [`../01-RESEARCH/`](../01-RESEARCH/README.md) |
| Memory queries that failed or surprised (archivist substring search, grading verdicts, coverage honesty) | Archivist search design; memory-convention refinements |
| Operational friction: NGS connection/stream/byte limits, MCP connection accumulation, rollup cadence, notification noise | Backlog ordering (provisioning limits, MCP connection thrift, curator digests) |
| Vocabulary gaps: things said *about* the work that had no op type | Whether remaining vocabulary is actually remaining |

## Definition of done

The roadmap's own words: two weeks; topics announced, threaded conversations
with mentions and attachments, work items claimed and done, topics closed —
and at the end, a journey episode ([`../04-JOURNEY/`](../04-JOURNEY/README.md))
that turns the chafe log into decisions: eg-walker gate verdict, sealed-topics
priority verdict, and the next roadmap order.
