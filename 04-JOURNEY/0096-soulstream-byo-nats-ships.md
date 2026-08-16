# Episode 0096 — BYO NATS ships: the kit, the probes, the driver (2026-08-16)

**What happened:** design 0003 went from decided to built in the same
day — soulstream `specs/010-byo-nats`, merged to main (`3b396d7`,
merge `d3282a2`). `soulstream init --byo self-hosted --url …` emits
**the kit** (exact `nsc` commands and config fragments, public keys
only, byte-identical on every re-run) and exits; the operator applies
it and hands back two account public keys; `init` re-run mints the
bypass users (signing-key-signed, `IssuerAccount` set — no master key
exists on our side to sign with), **probes the substrate behaviourally**
(anonymous connect must refuse — a conf-auth server is named and
refused before anything runs; the ops user proves the realm account and
its plain signing key; whoami; JetStream; the issuer user proves the
AUTH wiring), boots the planes against the substrate URL, runs the wire
half, and closes with one **callout smoke round** — sentinel + token
must admit scoped, garbage must refuse, and a failed smoke takes the
sentinel marker back out. `--byo synadia-cloud` drives the same account
half through the control-plane API (byon-setup graduated out of
soulstream-mcp: idempotent by name, the plain workload group it never
created added, `log.Fatal` replaced by named refusals), one command, no
kit.

**Measured.** The whole self-hosted flavour end to end
[measured, `node/byo_test.go`]: a rig plays the operator — a stock
nats-server stood up from a *config file* in operator mode, its two
account JWTs authored from nothing but the publics the kit renders —
and the full founding passes M1.1 semantics through it (persona
server-asserted, scoped to its own prefix, on a server soulstream never
configured; garbage refused with the audited refusal; restart on the
same state; 0.33 s). The custody audit passes by file name and by
content: no operator/SYS/master artifact in the state dir, the rig's
operator and master seeds appear in no state-dir file, and the kit
document carries no seed [measured]. The kit's `nsc` sequence ran
verbatim against real nsc — including the `{{account-subject()}}`
scope templates and the `describe --field sub` hand-back — all `[ OK ]`
[measured, quickstart]. The Synadia driver is stub-proven (sequence,
idempotence-by-name, the lost-seed refusal: a programmatic seed is
returned exactly once, so a group without its custodied seed refuses by
name) [measured against the stub]; **the live Cloud founding is the
quickstart's manual runbook, deliberately not in `make test`** — the
Entra-lane precedent. Gate: fmt, full suite, lint — green, nothing
skipped.

**What the build taught, beyond the design:** (1) the config block
needed `url` — the design sketch had nowhere for the substrate to live;
(2) the issuer user's seed is phase-1 *persisted* material
(`keys/issuer-user.nk`), because its public key rides in the kit before
its creds can exist; (3) `n.url` really was the single seam — the whole
BYO branch in `node.Start` is "skip the server, point n.url at the
substrate", and no plane changed; (4) one honest gap stands: on Synadia
Cloud the platform may custody the callout xkey, and then the callout
runs **unsealed** on our side — printed at founding, recorded in design
0003 §6's as-built note, re-measured whenever the runbook next runs.
Deltas propagated back into design 0003.

Reversal condition: inherits design 0003's two (recurring abandoned
inits at the kit's `nsc` step reopen the account-authoring driver;
Synadia exposing callout config on shared plans reopens NGS as a third
flavour). New, from the build: if the unsealed-callout state on Synadia
Cloud proves unacceptable in the runbook's first live pass (observable:
the sealed-requests requirement failing the founding there), the
synadia flavour gains a blocking requirement on platform xkey export —
or loses the flavour.

Trail: soulstream `specs/010-byo-nats/` (spec, plan, tasks, quickstart
with the nsc validation and the live runbook); commits `3b396d7` +
merge `d3282a2`; design
[`0003-byo-nats.md`](../02-DESIGN/soulstream/0003-byo-nats.md) (§6
as-built note); [episode 0095](0095-soulstream-byo-nats-designed.md)
(the design pass); [episode 0038](0038-soulstream-remote-mcp-node.md)
(the Synadia facts the driver rides on);
[`ROADMAP.md`](../03-IMPLEMENTATION/ROADMAP.md) (BYO NATS moved from
later-horizon to landed).
