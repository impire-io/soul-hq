# Episode 0098 — brew install soulstream: the tap opens (2026-08-16)

**What happened:** soulstream became one Homebrew command —
`brew install impire-io/tap/soulstream`. Two halves: the tap repo
(`impire-io/homebrew-tap`, new, public) carries a hand-published
formula for `v0.13.0-rc.1` so the path works today, and the release
pipeline (episode 0058's, extended: `brews` in `.goreleaser.yaml`,
`HOMEBREW_TAP_TOKEN` passed through the workflow) owns the formula
from the next tag on — RCs included, deliberately: pre-v1 the
candidates ARE the releases, and the tap is the easy-easy install (no
Go toolchain, nothing unpacked, nothing on PATH by hand).

**Measured:** on a clean PATH (`which soulstream` → not found),
`brew install impire-io/tap/soulstream` installed 0.13.0-rc.1 from the
published release artifact (the formula pins the release checksums),
`soulstream version` stamps `0.13.0-rc.1`, and `brew test soulstream`
passes [measured, darwin_arm64]. `goreleaser check` validates the
extended config.

**Two decisions, named:**

- **Formula over cask.** goreleaser deprecates `brews` in favour of
  `homebrew_casks`, but casks are macOS-only and the BYO/self-hosted
  audience runs Linuxbrew — the formula serves both platforms today.
  Deprecation accepted with eyes open. *Reversal condition*:
  goreleaser v3 removing `brews` (observable: the release run failing
  on the key) forces the move — cask for macOS plus a documented
  Linux path, or a self-templated formula step.
- **RCs publish to the tap** (`skip_upload: false`, not `auto`): a tap
  that only serves stable versions would serve nothing before v1.

**The one open dependency, said loudly:** the cross-repo formula push
needs `HOMEBREW_TAP_TOKEN` (fine-grained PAT, contents read/write on
the tap repo only) as a soulstream Actions secret. Until it is set,
the next tag's release run FAILS at the brew step — the secret is an
operator act pending at this episode's close.

Trail: [`impire-io/homebrew-tap`](https://github.com/impire-io/homebrew-tap)
(first commit: the v0.13.0-rc.1 formula); soulstream `09a8902` (the
goreleaser `brews` block, workflow env, README + getting-started lead
with brew); episode [0097](0097-soulstream-v0-13-0-rc-1.md) (the
candidate the formula pins); episode 0058 (the pipeline this extends).
