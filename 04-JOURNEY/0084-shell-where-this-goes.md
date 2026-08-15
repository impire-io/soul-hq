# Episode 0084 — Where this goes (2026-08-15)

The operator's ask, minutes after the waker landed: "include instructions
in the shell's agents module on how people can configure their
claude-code/pi.dev/…". The shown-once credential screen (episode 0079's
copy-paste block) now carries a **"Where this goes"** section — three
folds, each in the receiving program's own terms `[measured, unit +
whole gate]`:

- **Claude Code**: the block above is already the exact `.mcp.json` —
  save it in the folder the agent works from, or merge the entry into an
  existing one; headless runs want `--mcp-config`.
- **Codex**: the same five values rendered as a **filled TOML table** for
  `~/.codex/config.toml` — a second copy-paste block on the same
  shown-once reveal, repeating the exact variable names the stdio door
  resolves (one spelled wrong is an agent that does not start — the
  screen's own standing rule, kept).
- **Anything else that speaks MCP** (pi.dev, opencode, …): the shape in
  plain words — a declared stdio server named `soulstream`, the command,
  the five variables. No per-assistant machinery for programs whose
  config formats we would only be guessing at; the generic fold covers
  them honestly.

Plus the one line every machine needs first: the `soulstream-mcp` install
(go install or release download). Plain-language rule held: folds are
labeled by the program people run, not by ecosystem bynames. The secret
appears in the codex fold as it does in the block — same reveal, same
shown-once lifetime, nothing new kept anywhere. The pinned agents e2e
(0079's ceremony) passed untouched — the added folds sit outside the
`data-credential` contract `[measured]`.

Reversal condition: none — records a completed build. If a named
assistant's real config format diverges from what a fold claims
(observable: an operator following a fold verbatim and the agent not
starting, reported as an issue), that fold is corrected or demoted to the
generic one.

Trail: soulstream-shell `967f150` (feat(agents): the credential screen
says where it goes); design 0002 §2's agents module carries the surface;
episode [0079](0079-shell-agents-join-the-stream.md) is the block this
explains, episode [0083](0083-workloads-the-waker-lands.md) the waker
those agents can now be woken by.
