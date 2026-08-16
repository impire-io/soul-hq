# Episode 0094 — v0.12.0-rc.1: the clean-break candidate (2026-08-16)

The third pre-release, and a new candidate series rather than rc.3 of
the old one — the shape changed underneath it, on purpose [judgment:
an RC series should not absorb clean breaks]. The tag-triggered
pipeline (episode 0058's) built and published on the first run:
linux/darwin × amd64/arm64 with checksums; the darwin_arm64 artifact
round-tripped — checksum verified, `soulstream version` stamps
`0.12.0-rc.1`, the usage carries the new verbs [measured].

**Pins**: core v0.8.4, workloads v0.4.0, shell v0.6.0, idp v0.5.0,
archivist v0.3.0. What a person on this candidate gets over
v0.11.0-rc.2 (episodes 0089–0093, two days of arc):

- **One binary, one paste** (0089): native `soulstream wrap` and
  `soulstream mcp`; the Agents screen leads with the portable paste
  block; `go install` is gone from every instruction.
- **One canon** (0090): the sign-in and admin pages wear the same
  cassette-light system as the console.
- **One console** (0091): full administration in the shell, drawn only
  for administrators; the idp's HTML console unmounts in the bundle
  (`/admin` → 404).
- **One vocabulary, no fallbacks** (0092/0093): `planes.signin`,
  `planes.mcp`, `--signin-listen`, `--mcp-listen`, `SOULSTREAM_STATE`
  — and a realm founded under the byname-era names is refused with the
  hand-migration spelled out, never silently misread.

The break is the release's headline, stated in the tag message: a
realm founded on an earlier candidate needs the three-line
hand-migration (rename the config keys `door→mcp`/`fold→signin`,
`mv users/fold.creds users/signin.creds`, `mv fold/ signin/`) or a
re-init.

Reversal condition: none — records a completed cut; the RC soaks
toward v0.12.0, superseding the v0.11.0 series.

Trail: tag
[`v0.12.0-rc.1`](https://github.com/impire-io/soulstream/releases/tag/v0.12.0-rc.1);
commits `49baf70` (docs name the candidate), site `0e56918`; episodes
0089–0093 for the arc it ships.
