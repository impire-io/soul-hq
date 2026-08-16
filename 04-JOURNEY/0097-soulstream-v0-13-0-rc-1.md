# Episode 0097 — v0.13.0-rc.1: bring your own server (2026-08-16)

The fourth pre-release, cut the same day its headline landed. The
tag-triggered pipeline (episode 0058's) built and published on the
first run: linux/darwin × amd64/arm64 with checksums, marked
prerelease by the `-rc` suffix. The darwin_arm64 artifact
round-tripped — checksum verified, `soulstream version` stamps
`0.13.0-rc.1`, and the usage carries both BYO founding paths
[measured].

**Pins unchanged** from v0.12.0-rc.1 (core v0.8.4, workloads v0.4.0,
shell v0.6.0, idp v0.5.0, archivist v0.3.0) — this candidate is one
arc over the last: **BYO NATS** (episodes
[0095](0095-soulstream-byo-nats-designed.md)/[0096](0096-soulstream-byo-nats-ships.md),
design 0003, specs/010). A person on this RC can found their realm on
a server soulstream does not run: self-hosted through the kit (nsc
commands and config fragments with real values, behavioural
verification, refusals by name, a callout smoke round), or Synadia
Cloud BYON through the control-plane API in one command. Operator mode
is the requirement in both; conf-auth servers and NGS shared plans
refuse by name. No operator or account master key ever travels.

Known-open on this candidate, stated in the spec: the live Synadia
Cloud founding is a manual runbook not yet run against a real BYON,
and on that platform the callout may run unsealed when the xkey is
platform-custodied (printed at founding).

Reversal condition: none — records a completed cut; the RC soaks
toward v0.13.0.

Trail: tag
[`v0.13.0-rc.1`](https://github.com/impire-io/soulstream/releases/tag/v0.13.0-rc.1);
soulstream commits `3b396d7`/`d3282a2` (the feature), `d2676b1` (docs
name the candidate); episodes 0095–0096 for the arc it ships.
