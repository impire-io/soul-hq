# Episode 0155 — The second canon builds: a values swap, not a rebuild (2026-08-29)

Design 0011 built the day it was decided, autonomously, on branch
`second-canon` in both repos — shell `3a1ca5d`, idp `32855fd` — with
the full gates green on the first complete run: fmt, build, every
unit test, the consumer-position e2e whole (35.8s, the offline render
gate now fetching Geist, the custody scans, the ceremony), lint clean
[measured]. The port's central finding: **the class vocabulary was
the load-bearing seam.** Because every screen speaks the component
layer's class names and the class names are canon-neutral, the swap
was a values change — the token source rewritten (paper/ink, one teal
accent, Geist variable binaries vendored with the OFL, borders over
bevels, no textures, 80/140/220/360ms ease-out), the chrome reshaped
once (`frame.go`: the ink topbar and icon spine became the 232px
labeled sidebar, wordmark at the top, the signed-in name at the
foot), the canon prose re-vendored — and **zero module render files
changed** [measured: the diff]. Design 0011 §8 had planned three
movements against a predicted 654-line re-class and ~250 moved test
lines; the build moved three test sites and nothing else, ledgered in
0011 §11.

The honest steps re-stepped rather than fell. The 560px "bar sheds a
strip" rung died with the bar; at 900px the sidebar now sheds its
labels to a slim 60px icon rail beside the conversations drawer, the
labels coming back over the content on the frame's own signal — and
the ladder test asserts the new ladder as a ladder. At 390px the
measurement holds: scrollWidth equals innerWidth, nothing scrolls
sideways [measured, in the browser]. The 0011 §4 re-homing reads on
the screens as designed: the mention card hardens to ink, Scribe
carries the teal dot with "operated by Daan" in words, whose-is-whose
stays on alignment. The fold moved the same day — `webstyle` a pure
values swap, every selector held, its gate green.

What remains is §9.4 by its own terms: the operator's eyes. The ten
screens, before and after, are published side by side (with the
390px honesty shot and the gate results) for that call; the branches
stay unmerged until it lands. Merge, tag, and the product pin move
ride the acceptance, not this episode.

Reversal condition: none of its own — records a completed build
against design 0011; the decision's reversal conditions live in
episode 0154, and §9.4's acceptance is the standing gate this build
waits behind.

Trail: soulstream-shell branch `second-canon` (`3a1ca5d` — tokens,
chrome, fonts, canon prose, three test sites); soulstream-idp branch
`second-canon` (`32855fd` — webstyle); design
[`0011`](../02-DESIGN/soulstream-shell/0011-the-second-canon.md) §11
(build amendments in place); the before/after review shown to the
operator.
