# Episode 0058 — v0.1.0: the house gets a shipping label (2026-08-03)

The last structural gap between "the folded realm works in `make
test`" and "a person can download and run it" closed: SoulNode grew CI
and the tag-triggered release pipeline — the archivist pattern,
composed unchanged — and **v0.1.0 is released** with binaries
(`007-ci-and-release` in the soulnode repo).

What landed and what was measured:

- **ci.yml**: gofmt, build, test, lint on every push and PR. The first
  run went green on the GitHub runner — including the folded-realm
  gate, the public-door OAuth walk, and the agent-under-enforcement
  suite, all on an ubuntu box that had never seen the code [measured].
- **release.yml + .goreleaser.yaml**: a `v*` tag runs the tests and
  publishes archives for linux/darwin × amd64/arm64 with checksums,
  the version stamped into `internal/version`. Proven live: the
  `v0.1.0` push produced the release; the darwin/arm64 artifact was
  downloaded back, its checksum verified, and the binary answers
  `soulnode version` → `0.1.0` [measured].
- **No windows, on purpose**: the founding ceremony refuses
  filesystems that cannot hold owner-only modes (0700/0600 custody is
  load-bearing), so a windows binary would refuse at `init`. Honest
  scope, recorded in the goreleaser config; revisit only with a
  measured windows custody story.
- **The operator made soulrealm public** — the one private module in
  the consumed stack — so the pipeline holds zero credentials: CI
  fetches every impire-io module straight from its repo (`GOPRIVATE`
  skips proxy lag), and the long-standing "private-module credential"
  blocker dissolves rather than being provisioned. Soulstream's
  dormant node CI can wake on the same grounds, its repo's call.

The vision sentence now has its distribution half: download one
archive, `soulnode init && soulnode up`. What still stands between
v0.1.0 and handing it to a stranger is product, not plumbing: the
fold's M3 lifecycle behind the loudly-interim first-touch enrollment
(research since concluded — [episode 0059](0059-soulfold-bootstrap-story.md)),
the physical-authenticator runbook (a human act, pending), and the
named-not-planned day-2 items (upgrade in place, backup/restore).

Reversal condition: none — records a completed build; the
windows-scope decision reverses only on a measured custody story, as
written in the config.

Trail: `.github/workflows/ci.yml`, `.github/workflows/release.yml`,
`.goreleaser.yaml`, tag `v0.1.0` and its
[release](https://github.com/impire-io/soulnode/releases/tag/v0.1.0);
the `007-ci-and-release` merge.
