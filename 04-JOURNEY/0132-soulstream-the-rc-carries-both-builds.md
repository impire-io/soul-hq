# Episode 0132 — The rc carries both builds: v0.14.0-rc.1 across the stack (2026-08-25)

The day's two builds (episodes
[0130](0130-ecosystem-the-agent-declaration-builds.md)/[0131](0131-core-sealed-topics-build.md))
became installable the same day, at the operator's direction: a
coordinated pre-release line so people test through Homebrew.
**soulstream-core `v0.14.0-rc.1`** (the system stream + sealed topics;
goreleaser pre-release with archives). **soulstream-workloads
`v0.8.0-rc.1`** (the agent declaration), its core pin bumped from the
stale `v0.11.1` to the rc and verified standalone — `GOWORK=off` build
and full suite green against the pinned tag, the check module consumers
actually feel `[measured]`. **soulstream `v0.14.0-rc.1`**, the house:
pins to both rcs, and the wrap verb grows the `--declaration`
passthrough (the same wiring as `soulstream-wrap`'s, design
soulstream/0002's composition rule — a thin main over the pinned
library), so the brew-installed binary drives declared agents' four
wake kinds. `soulstream init`/`up` now provisions `SOULSTREAM_SYSTEM`
for free through core's `Provision`.

The release rode the shipped rails: goreleaser's `prerelease: auto`
marked all three; the tap formula moved to `0.14.0-rc.1` on the
release's own commit (`brew: soulstream v0.14.0-rc.1`) — RCs publish to
the tap deliberately, pre-v1 the candidates ARE the releases. Verified
from the outside `[measured]`: the darwin_arm64 release tarball's
binary prints `0.14.0-rc.1` and its usage carries `--declaration`.

One real defect surfaced and died: the product release's Test job
failed its first run on `TestWrapLifeAgainstFoundedRealm` — the spec-011
rig was the one sibling still binding the fixed plane defaults
(8080/8378/8500), racing the node package for 8378 under package
parallelism `[measured: CI run #81 attempt 1, and the same one-off
locally]`. Fixed on main (`47dc354`) with the foldplane rig's pattern —
reserve the fold's port, clear the issuer so the derived default
matches, everything else `:0` — and the rig dropped from 10s (port
grinding) to under 1s. The tag's release run passed on re-run; the fix
rides the next tag.

Not in this rc: the byon deployment stays on the v0.13.0-rc line
(rc.11) — the soak continues; adopting the rc there is the operator's
own act. The sealed-topics ceremonies live in core's CLI (`key
sealing …`), not the product binary — a product surface for them is a
composition question by demand.

Reversal condition: none — records a completed release with its
measured verifications; the rc line itself is judged by the soak, as
every candidate is.

Trail: tags `soulstream-core v0.14.0-rc.1`, `soulstream-workloads
v0.8.0-rc.1` (pin bump `7a428be`), `soulstream v0.14.0-rc.1`
(`be41db6`, flake fix `47dc354`); the tap
`impire-io/homebrew-tap@Formula/soulstream.rb` at `0.14.0-rc.1`; builds
[episode 0130](0130-ecosystem-the-agent-declaration-builds.md) and
[episode 0131](0131-core-sealed-topics-build.md).
