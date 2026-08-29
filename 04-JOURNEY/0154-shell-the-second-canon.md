# Episode 0154 — The second canon: the skin moves, the words hold (2026-08-29)

The operator authored a new design system and a full console kit in a
design project ("New Impire Design System", the design-sync lane the
first canon rode — episode [0077](0077-shell-the-canon-held.md)) and
asked for the shell to be revisited against it. The structural read
of the shell at v0.11.0-rc.5 sized the ground [measured, read from
the repo]: one token source (`tokens.css`, 869 lines), one chrome
file (`frame.go`, 202 lines), 654 markup-bearing lines across 20
render files, ~250 test lines asserting on markup, the fits/purity/
banned-word/offline gates all standing. Held side by side, the kit
maps onto the shell almost one-to-one — its Overview, Conversations,
Approvals, People, Workloads, and Audit screens are Home,
Conversations, Approvals, People & sign-in, Agents, and Storage in
new dress, down to both audit surfaces refusing text search for the
same reason and both folding advanced fields behind disclosures
[judgment, from the side-by-side read]. The decision landed as design
[`0011`](../02-DESIGN/soulstream-shell/0011-the-second-canon.md):
**adopt the canon visual-for-visual** — paper/ink, Geist, one teal
accent, borders over bevels, the icon spine becoming a labeled
sidebar — while every name, route, act endpoint, and the sheet
grammar (design 0007) hold.

Two refusals shaped the scope, both the operator's calls. The kit's
*platform vocabulary* — Bridges, gateway, vault, credential-scope
pills, teams-as-accounts — is not adopted: its foundation documents
describe an architecture soulstream does not have, and renaming Tools
to Bridges ahead of it would put words on screen the realm cannot
honor [mechanism-argument, the honesty rule applied to a name]. Those
documents stand on the roadmap as a named horizon (future research
topics), and the vocabulary revisit rides their graduation. And the
first canon's two-channel accent semantic is **re-homed, not
reversed**: 0077's reading (accountability, not species) survives —
whose a message is stays on alignment, who answers for a voice stays
in words and the byline's dot — but color stops being the vehicle,
because the second canon has one accent [judgment]. The re-skin
reaches all ten screens (Models and System status get the canon by
the grammar they already follow) and then the idp fold
(`session-and-ui.md` §D30's companion), so sign-in does not wear the
old clothes.

What it opened: milestone M-second-canon — three movements on one
branch (tokens → chrome → screens), gated on the full check plus the
operator calling the canon held on before/after screens — and three
named [O]s: the Impire mark (a naming act, parked), the kit's dark
theme (a follow-up once the tokens are the kit's), the vocabulary
horizon. One incidental defect found by the read and folded into the
arc: the first canon's `@font-face` for `jetbrainsmono-500.woff2`
points at a file never vendored — a silent 404 on every
medium-weight-mono request — which dies with the font swap
[measured].

Reversal condition: the one-accent call reopens if the operator,
living on the new screens, can no longer tell at a glance which
voices somebody else answers for — that reading reopens 0011 §4 with
the screens as evidence. The vocabulary refusal is not
reversal-shaped: the foundation architecture graduating from research
is its named path in.

Trail: design
[`0011-the-second-canon.md`](../02-DESIGN/soulstream-shell/0011-the-second-canon.md)
(new), design 0001 §7 amended in place, the roadmap's ecosystem note
("The second canon and the Impire horizon") and shell section; the
design project's `ui_kits/product/*` (the console kit),
`colors_and_type.css` (the tokens), `docs/foundation/*` +
`docs/bridges/*` (the horizon); episode 0077 (the first canon, whose
vendoring lane and drift lesson this arc inherits).
