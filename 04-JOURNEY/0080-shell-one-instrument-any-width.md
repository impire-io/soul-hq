# Episode 0080 — One instrument, any width (2026-08-14)

The operator split their screen and the shell clipped — the Agents
page cut mid-word, the form and table off the right edge. The frame
had been built at one width. The responsive pass (soulstream-shell
`2fa4ca8`, v0.4.1) taught it to give ground in honest steps: the
details panel hides below 1180 px (as before), the conversations list
collapses (~900 px) into an overlay the rail's Conversations key
opens — badge intact — the content column goes fluid, the transport
bar truncates gracefully, tables scroll inside their own containers,
Home's cards wrap, and the conversation column holds down to phone
width.

Measured in a live browser at review `[measured]`: at 1000 px (the
operator's split) and 390 px (a phone), all four screens — Agents,
conversation, Home, People & sign-in — report
`scrollWidth == innerWidth`, zero horizontal overflow. The gate grew
hermetic guards so it stays true: every served table sits in a scroll
container, no served markup carries a fixed width past 360 px, every
page carries its viewport meta.

Tagged through: shell v0.4.1 → soulstream **v0.10.1**.

Reversal condition: none — records a completed build.

Trail: screenshots `shell-v10-split-agents.png` (the complained-about
page at the complained-about width, whole) and
`shell-v10-phone-conversation.png`; the canon
(`docs/design-canon.md`) held throughout — instrument-panel density,
not mobile mush.
