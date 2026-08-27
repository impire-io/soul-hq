# Episode 0139 — The rc.2 carries the tenants and the capabilities: v0.14.0-rc.2, identity v0.12.0, workloads v0.8.0-rc.2 (2026-08-27)

The 08-27 arc became installable the same evening it landed, on the
0132 rails. **soulstream-identity `v0.12.0`** — the first tag since
v0.11.0, carrying the whole run: D47 tenants-born-admissible
(`447ec6b`), the `accounts.*` client surface (`df8e4a3`), the provider
arm's D47 parity (`f6a1a33`) with both BYON live-run fixes
(`a0545c8`/`31279c6`), `KindNATSOperatorKey` (`95a8d9e`), the agent
scope (`e032687`, specs/004), and the sealing custodian (`72bd164`,
specs/005). No binary pipeline fires — as with every identity release,
the signed module tag is the release. **soulstream-workloads
`v0.8.0-rc.2`** — capability minting (`14e95a0`) and the local-first
scoped lane (`2465dba`), the gate run `GOWORK=off` against the pinned
core `v0.14.0-rc.1` so the tag proves what module consumers actually
feel [measured: full suite, the 108s integration pass included].
**soulstream `v0.14.0-rc.2`** — the pins move from the day's
pseudo-versions to the named tags (`8f72f66`, taking the idp `v0.8.1`
patch that had sat unadopted), carrying spec 012 (tenants in the
house) and spec 013 (capability minting in the house) plus rc.1's
wraplife port fix. Core stayed untagged on purpose: its main has not
moved since rc.1. Every gate ran uncached-green before its tag
[measured].

Verified from the outside, the 0132 standard: the release marked
prerelease with all four platform tarballs, the tap formula at
`0.14.0-rc.2` on the release's own commit, and the darwin_arm64
tarball's binary printing `0.14.0-rc.2` and answering `soulstream
account` [measured]. The 0137 standing exception (workloads/identity
pins bump after their mains push) closed inside this cut — `aa354b4`
rode the mains, this release names them.

One boundary held itself: the assistant's session could push the
module tags but the product's release-triggering tag push was
refused by the permission layer — the push that publishes to people
stayed the operator's own hand, which is the house rule arriving by
mechanism rather than discipline.

With this rc the declaration story is installable end to end: `brew
install impire-io/tap/soulstream`, one JSON declaration (wake,
instructions, capabilities, budget), `soulstream wrap --harness claude
--declaration agent.json` — the names in the declaration becoming the
credential's exact reach. What the rc does *not* change: an agent
still runs where a person runs wrap. That gap became the next focus
the same evening (episode
[0140](0140-ecosystem-the-focus-agents-as-infrastructure.md)).

Reversal condition: none — records a completed release with its
measured verifications; the rc line is judged by the soak, as every
candidate is.

Trail: tags `soulstream-identity v0.12.0`, `soulstream-workloads
v0.8.0-rc.2`, `soulstream v0.14.0-rc.2` (pin bump `8f72f66`); the tap
`impire-io/homebrew-tap@Formula/soulstream.rb` at `0.14.0-rc.2`;
builds episodes [0134](0134-identity-tenants-born-admissible.md)–[0138](0138-identity-the-sealed-record-gains-its-custodian.md);
the previous cut [episode 0132](0132-soulstream-the-rc-carries-both-builds.md).
