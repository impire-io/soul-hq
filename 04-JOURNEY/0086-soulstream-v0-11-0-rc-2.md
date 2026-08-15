# Episode 0086 — v0.11.0-rc.2: wrap ships (2026-08-15)

The second release candidate, cut the evening of the day the whole arc ran
— research (0082), the waker (0083), the folds (0084), the reshape to wrap
(0085). The product pins the versions that carry it: **soulstream-core
v0.8.4** (`PostTurnIdempotent` at v0.8.3; the external-subcommand seam at
v0.8.4 — `soulstream wrap` reaches `soulstream-wrap` from PATH),
**soulstream-workloads v0.3.0** (the wrap engine and `soulstream-wrap`;
the daemon cut with its reversal in design 0004 §9), **soulstream-shell
v0.4.3** (the credential screen's four set-up folds, ending in "make it
answer mentions on its own"); idp v0.4.1 and archivist v0.3.0 ride
unchanged. Product gate green (`make fmt/test/lint`, node suite 7.2s);
tag `v0.11.0-rc.2` signed and pushed; the release pipeline (episode
0058's) builds the archives and marks the prerelease itself
`[measured: pins upgraded cleanly, gate green on first run]`.

What the candidate means in one sentence: a person on this RC can create
an agent in the shell, paste one block, and either talk *through* their
assistant (the MCP door) or let it answer mentions *for* them
(`soulstream wrap`) — from their own machine, with their own logins.

Reversal condition: none — records a cut; the direction decisions live in
episodes 0082/0085 and design 0004.

Trail: soulstream `69af93d` + tag `v0.11.0-rc.2`; component tags
workloads `v0.3.0`, shell `v0.4.3`, core `v0.8.4` (all signed, pushed
today); episodes [0082](0082-ecosystem-agent-participation.md)–[0085](0085-workloads-wrap-run-your-agent-where-you-are.md)
are the day's record.
