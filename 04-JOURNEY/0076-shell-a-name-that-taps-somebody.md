# Episode 0076 — Type @, pick a person: a name that taps somebody (2026-08-14)

The gap episode [0075](0075-shell-mentions-and-a-word-retired.md) named
was closed from both ends the same day, at the operator's direction
("meaningful tags and some sort of autocomplete").

**Core grew its first consumer-proven arm since the focus** — the
topic package's turn and comment paths accept *explicit* mentions
beside the body (union with the parsed slugs, deduped, invalid
dropped; additive API, zero wire change — the mentions field always
existed). Tagged **v0.8.1** `[measured: 101 test lines, gate green]`.

**The shell got the picker.** Typing `@` asks the server who is in
the conversation (the same participants the details panel derives, off
the same record) and morphs a list into its own target inside the
composer; typing narrows by name, word, or handle; Enter/Tab picks,
Escape declines. **What is written is never touched**: the body
reaching the record says `@Daan` because that is what was typed; who
it taps rides beside it through core's new arm. Resolution is decided
server-side at post against the record — a deleted token untaps its
person, a hand-typed name counts only when exactly one person answers
to it, ambiguity stays as typed. Rendered back, a name is marked only
if the record says a slip actually reached an inbox — text that merely
looks like an address stays plain.

**C4's reversal condition was tested and survived** `[measured]`: the
morph model held — an open picker with a half-written message under it
survived four stream ticks unchanged. The honest cost is two pieces of
page-local state no server can know unasked (the caret; that Escape
declined a list), neither of which is state about the record, which is
re-read on every keystroke. The Datastar rendering pick stands, with
its cost now named instead of feared.

Reversal condition: none — records a completed build; C4's own
reversal reading stays live in design 0001 §5.

Trail: soulstream-core `71f818e` (v0.8.1), soulstream-shell `1b1c63c`
+ `7c386d3` (the pin); screenshots `shell-v5-meaningful-tag.png` /
`shell-v5-autocomplete.png` shown to the operator.
