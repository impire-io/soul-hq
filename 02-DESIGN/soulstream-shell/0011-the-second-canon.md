# 0011 — soulstream-shell: the second canon

**Status:** decided, built, and **ACCEPTED 2026-08-29** — the
operator called the canon held on the before/after screens (§9.4
paid; episodes
[0154](../../04-JOURNEY/0154-shell-the-second-canon.md)/[0155](../../04-JOURNEY/0155-shell-the-second-canon-builds.md)/[0156](../../04-JOURNEY/0156-soulstream-the-rc7-carries-the-second-canon.md)).
Merged and shipped: shell **v0.11.0-rc.6**, idp **v0.8.2**, the house
pinning both as **v0.14.0-rc.7**. Build amendments are marked in §11. The
operator authored a new design system and a full console kit in a
design project (**"New Impire Design System"**, the design-sync lane —
the same lane the first canon rode, episode
[0077](../../04-JOURNEY/0077-shell-the-canon-held.md)); this document
decides what the shell adopts from it and what it explicitly does
not. Grounded in a structural read of the shell repo at v0.11.0-rc.5
(2026-08-29): one token source (`shell/assets/tokens.css`, 869
lines), one chrome file (`shell/frame.go`, 202 lines), 654
markup-bearing lines across 20 module render files, ~250 test lines
asserting on markup.

## §1 The decision, in one sentence

The **Impire design system replaces Soulsystem as the shell's visual
canon** — tokens, type, depth, motion, and the chrome — while every
screen keeps its name, its route, its module, and its grammar: a
re-skin adopted **visual-for-visual**, with the kit's new platform
vocabulary deliberately left at the door (§3).

## §2 The canon that arrives

The project is the source of truth and the shell vendors it, exactly
as the first canon worked ("re-vendor on canon change — the project
wins", episode 0077): the token layer into `tokens.css` layer 1, the
canon prose into `docs/design-canon.md`, fonts and icons into
`shell/assets/`. What changes, token-for-token:

| | Soulsystem (first canon) | Impire (second canon) |
|---|---|---|
| World | cassette futurism in a light key — warm plastic, bevels, grain | paper and ink — calm surfaces, hairline borders, nothing decorative |
| Surfaces | warm shell tones `#FBF8F1`→`#A2937A`, one CRT glass | paper `#FAFAF7`, card white, sunken `#F2F2EE`; ink `#0B0E13`→`#F5F6F8` |
| Accents | **two at equal weight** — amber (human) / teal (machine) | **one** — teal `#1FA88E`, used sparingly; red/green/blue stay state-only signals |
| Type | Archivo + JetBrains Mono | Geist + Geist Mono (single families, weight is the only play) |
| Depth | bevels, etch lines, press travel | 1px borders first; two small shadows; no bevels |
| Radii | 2/3/5/8/12 | 2/4/6/10, pills for tags only — unchanged in spirit |
| Motion | 70/120/180/280ms `--ease-mech` | 80/140/220/360ms ease-out — same discipline, softer curve |
| Texture | grain, ribbing, scanlines | none |

Continuities worth naming so the port keeps them deliberately: the
mono **label strip stays the component's identity** (the kit's eyebrow
is 12px uppercase at the same `0.14em` tracking the strips carry
today), and the **dark glass for raw output survives** — the kit sets
its payload panes ink-on-ink with teal-tinted text, so the storage
explorer's byte-for-byte panels keep their glass, without the
scanlines. The kit's dark theme (`[data-theme="dark"]`) is **not**
adopted this arc — the shell stays in its light key; dark is a named
horizon (§10).

## §3 The scope line (normative)

**Only the skin moves.** Holding, by name:

- Every nav label and route: Home, Conversations, Tools, Approvals,
  People & sign-in, Agents, Models, System status, Storage, Sign out
  — and the conditional absence of the three declared-fact modules.
- The sheet grammar
  ([0007](0007-the-sheet-shape.md)) whole: h1 → lede → act key →
  table → stow folds → result line → slide-over. The kit's console
  screens follow the same rhythm (its PageHead / table / disclosure /
  slide-over are the same parts under other names), which is why this
  is a re-skin and not a redesign.
- The copy register (0001 §7 C8, 0007 §3): plain functional names, no
  component bynames, machine-room vocabulary off the screens.
- The module contract ([0002](0002-the-module-shape.md)), the lanes
  (0001 §4/§6), every act endpoint, every test's meaning.

**Not adopted**, by name: the kit's console vocabulary and the
architecture underneath it — Bridges, gateway, vault, credential-scope
pills (`obo-downstream` and kin), teams-as-NATS-accounts, host
bridges. Those come from the design project's foundation documents,
which describe a platform soulstream does not have; renaming Tools to
Bridges without that platform would put words on screen the realm
cannot honor. The foundation documents ride the roadmap as a named
horizon (future research topics), and the day that architecture is
researched and built, the *names* revisit rides it — not this
document.

## §4 One accent — re-homing the channel, not reversing it

The first canon's load-bearing semantic was **two accent channels at
equal weight**: amber the voice that answers for itself, teal the
voice somebody else answers for (episode 0077 — accountability, not
species, inheriting feature 014's rationale). The second canon has
one accent, so **color stops being the vehicle; the reading
survives**. Whose a message is keeps riding alignment; who answers
for a voice rides what already says it in words — the byline, the
operator named beside an agent in People, the countersigned
`operated_by` claim the record itself keeps — plus the small live/idle
dot the kit puts in bylines. This does not fire 0077's reversal
condition: the accountability reading stands; only its paint is
retired. If living with one accent loses something real — the operator
can no longer tell at a glance which voices somebody else answers for
— that observation reopens this section with the screens as evidence.

## §5 The chrome

`shell/frame.go` is the whole chrome and it changes shape once: the
ink topbar and 248px icon spine become the kit's **topbar-less,
232px labeled sidebar** — wordmark at the top, the nav items with
icon *and* label, the signed-in persona and Sign out at the foot. The
unread and pending marks keep their spine seats. The gate keeps its
shape (one card, one action) in the new dress. The content column
moves from 1080px to the kit's 1180px; 0007 §6's details-drawer step
is a viewport breakpoint and holds as written.

## §6 What the port must hold (the standing gates)

- **The offline render gate.** Geist and Geist Mono are vendored
  woff2 — the kit's Google-Fonts `@import` does not come along. Zero
  external fetches stays a measured property. (The first canon's
  phantom `jetbrainsmono-500.woff2` `@font-face` — declared, never
  vendored, a silent 404 today — dies with the font swap.)
- **The fits gates.** No inline `style=` in served markup — the kit's
  JSX inline styles are translated into the shell's class layer, which
  is the right pressure anyway; no `width`/`min-width` over 360px in
  `tokens.css`; every `<table>` keeps its scroll container.
- **The banned-word gates.** New class names pass them like the old
  ones did; the kit's component names arrive as the shell's own
  functional class vocabulary, not as bynames.
- **The purity gates and the module contract** — untouched by
  construction; this arc adds no import.
- **Icons hold.** The 21 vendored Lucide icons match the kit's icon
  convention (24×24 line, currentColor); the set carries over.

## §7 Reach

All ten screens move in one arc — including **Models** and **System
status**, which have no kit screen and get the canon by the sheet
grammar they already follow — and the **fold moves with the shell**:
soulstream-idp's `internal/webstyle` (bound to the first canon by
[`session-and-ui.md`](../soulstream-idp/session-and-ui.md) §D30)
takes the same tokens, so sign-in does not wear the old clothes at
the product's most visible moment. The idp's port is its own change
in its own repo, gated the same way, landing after the shell's token
layer exists to vendor from.

## §8 The landing shape

One arc on a branch, three movements, because the tests assert on
markup and will not be green between movements: (1) the token and
component layer — `tokens.css` rewritten, canon prose re-vendored,
fonts swapped; (2) the chrome — `frame.go`; (3) the screens — the 20
render files re-classed module by module, each module's markup
assertions moving with it. `make screens` before and after, the
screenshots shown to the operator — the first canon's drift lesson
(episode 0077: fast slices styling surfaces plausibly instead of
speaking the canon) is the standing risk, and eyes on every screen is
its counter.

## §9 Acceptance criteria

1. Full gate green (`make check`), including fits, purity,
   banned-word, and the offline render gate with Geist verifiably
   loaded — no external fetch, no 404'd font.
2. Every screen renders from the new token source; no Soulsystem
   token, bevel variable, or texture reference survives in
   `tokens.css` or the served markup.
3. Names, routes, copy, and act endpoints byte-unchanged where the
   scope line says they hold (the diff shows classes, chrome, and
   canon — not vocabulary).
4. The operator reviews the ten screens side by side (before/after)
   and calls the canon held.

## §10 Open questions [O]

- **[O1] The mark.** The kit's sidebar leads with the Impire mark;
  adopting a mark is a naming act, not a canon act — the shell keeps
  its current wordmark set in Geist until the operator decides the
  branding question in its own frame.
- **[O2] Dark.** The kit defines an opt-in dark theme; the shell
  ships light-only this arc. Adopting dark is a small follow-up once
  the token layer is the kit's.
- **[O3] The vocabulary horizon.** When the foundation architecture
  (bridges, gateway, vault, teams) graduates from research, the
  console vocabulary revisit rides those designs — tracked on the
  roadmap, explicitly out of this document's scope.

## §11 As built (2026-08-29, branch `second-canon`)

- **§8's three movements collapsed into one commit.** The class
  vocabulary held whole, so the port was a values swap: **zero module
  render files changed** — the predicted 654-line re-class never
  happened, and of the ~250 markup-asserting test lines only three
  sites moved (the frame's own tests, the collapse-ladder table, the
  mention-mark test). The blast radius was the token source, the
  chrome, and the canon prose.
- **The ladder re-steps.** The 560px "bar sheds a strip" step dies
  with the bar; the 900px step now also sheds the sidebar's labels to
  a slim icon rail (`--spine-slim`, 60px), with the labels back over
  the content on the frame's own `$rail` signal. Measured at 390px:
  scrollWidth = innerWidth, nothing sideways.
- **Geist vendored as the variable binaries** (100–900, 58KB each,
  OFL alongside) from the `geist` npm package — one file per family,
  the archivo-shape kept; the e2e offline gate re-pointed and green.
- **The fold moved the same day** — `internal/webstyle` a pure values
  swap, every selector held, idp gate green.
- **The one filled key per view held for free**: `.btn` is teal,
  `.btn.ghost` (the product's own secondary everywhere) is the white
  key with the ink edge — no call sites changed.
