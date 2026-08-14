# Episode 0075 — Mentions land, and a word retires (2026-08-14)

Slice 4 of episode [0071](0071-ecosystem-the-focus.md)'s usable
cockpit: **being tapped on the shoulder reaches the surface.** The
record's mention convention was always there — the library parses
`@name` from a posted body and drops a slip in each named persona's
inbox — and the shell was the only one not listening. Each signed-in
session now follows its own inbox over its own admitted connection
(the shell's read lane never reads anybody's mail): a count on the
spine's Conversations key, a mark on the conversation row, and the
mentioning message highlighted where it was said — that last one read
off the record, not the inbox, so it holds forever. Reading clears the
marks; the tray is session-memory only — nothing about who has read
what reaches the record. A slip pointing outside the board is not a
mark (anyone may drop one; a mark that opens onto nothing could never
be cleared), and a replayed inbox never resurrects a message already
read `[measured: all asserted in the gate, ~9.6 s uncached]`.

**The honest gap, named for upstream:** the mention grammar is a
lowercase slug keyed by persona id — `@u-590479…` taps a fold-issued
person, `@Daan` taps nobody. The shell does not pretend otherwise: it
never rewrites what somebody typed, and the People list carries each
person's handle behind their name. A display-name mention grammar
needs the directory mapping O3 left open — now with a concrete
consumer need against soulstream-core's mention convention.

Also landed: **the signed-in person is "Daan", not `u-…`** — display
name in the top bar and the People list (YOU pill on the name, raw
handle demoted to detail), the rig enrolling people with real names —
and **"realm" retired from every screen** at the operator's direction
("strange and ancient"): *Your soulstream at a glance*, the top-bar
strip now just the instance name, and the gate mechanically asserts no
served copy says the word again.

Reversal condition: none — records a completed build; the mention
read-state deliberately resets on sign-in (session memory only) until
real use shows it chafing — that observation would reopen it as a
design question, not a silent store.

Trail: soulstream-shell `a213c06`; screenshots
`shell-v4-home-mentions.png` / `shell-v4-conversation.png` shown to
the operator (badge lit, badge cleared by reading, highlight standing).
