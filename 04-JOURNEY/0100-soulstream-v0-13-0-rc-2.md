# Episode 0100 — v0.13.0-rc.2: first contact hardened, and the tap feeds itself (2026-08-16)

The fifth pre-release, cut hours after rc.1: the six fixes the first
live BYON founding forced (episode
[0099](0099-soulstream-the-byon-founding.md)), pins unchanged. The
pipeline built and published on the first run — and, for the first
time, **pushed the Homebrew formula itself**: the `HOMEBREW_TAP_TOKEN`
secret landed, goreleaser committed `brew: soulstream v0.13.0-rc.2` to
the tap, and `brew upgrade soulstream` on a machine that installed
rc.1 from the hand-published formula moved to rc.2 stamping
`0.13.0-rc.2` [measured]. Episode 0098's pending operator act closes;
the tap is now fully pipeline-fed.

What a person on this candidate gets over rc.1: a BYO founding that
survives the real world — the client-config-dir collision refused by
name, `--synadia-system` by name or id, the awaiting state resuming,
once-returned seeds persisted the moment they arrive, bounded retries
through a lossy agent channel, and read-first callout operations
against the platform's anti-idempotent 500s. Every fix carries a test
that replays its live incident.

Reversal condition: none — records a completed cut; the RC soaks
toward v0.13.0, superseding rc.1.

Trail: tag
[`v0.13.0-rc.2`](https://github.com/impire-io/soulstream/releases/tag/v0.13.0-rc.2);
soulstream commits `7c1f030`…`9b3aaa0` (the fix arc), `7f462f5` (docs
name the candidate); the tap's `5f0a0b3` (goreleaser's first formula
push); episodes 0098 (the tap opens) and 0099 (the arc it ships).
