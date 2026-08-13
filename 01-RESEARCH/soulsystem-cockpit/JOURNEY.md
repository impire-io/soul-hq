# soulsystem-cockpit — investigation journey

Topic opened 2026-08-13. Appended as the investigation happens.

## 2026-08-13 — C3 decided: cassette light is canon, identity by label

The operator decided both halves of C3, ahead of Bar 4 as that bar
requires.

**Canon: the cassette-futurism light system** (the Soulsystem Design
System project) becomes the ecosystem's visual language `[judgment]`.
The deciding argument: the amber/teal rule — human channel and machine
channel as two accents at deliberately equal weight, neither allowed to
outrank the other — is the vision's "humans and AI as peers" turned
into a color system; the violet/rose dark theme is semantically mute
beside it. Supporting the call: it is the only *complete* system of the
two (tokens, seventeen components, guidelines, voice canon, two UI
kits; `webstyle` is one CSS string), and the migration cost is at its
historic minimum — episode
[0063](../../04-JOURNEY/0063-soulfold-the-console-gets-a-face.md) was a
presentation-only change with reversal condition *none*, and the
impire.io website rework is already in flight (episode
[0065](../../04-JOURNEY/0065-ecosystem-fair-code.md)), so the site
adopts the canon rather than being repainted twice.

Costs accepted openly: this is a **committed light-only look** — "one
dark surface: CRT glass" is a core rule, and there is no dark mode; a
future dark key would be a deliberate amendment to that rule, never a
drift. The fold's four shipped pages get restyled (soulfold repo work,
sequenced later — Bar 4 needs only one fold page rendered from the
shared source as proof); the violet/rose language is retired from the
soulsystem. Middle paths were examined and rejected: re-keying the
design project dark breaks its own central rule (the CRT surface reads
as special only because the world is light), and a dual canon makes
every future surface pick a side forever.

**Per-component identity rides labels, not color**
`[mechanism-argument]`: amber/teal are *channel* colors (who is
speaking), the rose-edge idea was a *component* color (where you are) —
two color systems on one screen fight, and the equal-weight rule cannot
survive a third accent competing beside it. Component identity uses the
register the system already owns: mono label strips, wordmarks, and
badges ("the fold", "the house"). The rose edge retires with the dark
language.

## 2026-08-13 — Bar 4 measured: PASS — one source, two surfaces, zero external fetches

The bar ran the same day C3 was decided, on a rig in the session
scratchpad (`bar4/`: the shared source, two pages, vendored assets,
screenshots — scripts stay out of git per how-we-work).

**The rig.** One `shared.css`: the design system's seven token files
carried verbatim, with exactly one change — the Google Fonts CDN
`@import` replaced by vendored `@font-face` (Archivo variable woff2,
90,096 B, `wght` 400–800 *and* the `wdth` 62–125 axis; JetBrains Mono
400/500/700) — plus one shared component layer (transport bar, label
strips, LED pips, molded keys, fields, pills, segmented meters, tables,
the CRT screen) written to the system's own specs. Ten Lucide SVGs
vendored once (unpkg; one via jsDelivr after an unpkg 1102 error) and
inlined at build. Two pages consume it: the fold's sign-in re-rendered
with the *same markup structure* as the shipped `internal/ui` login
template, and the helm's first screen (four plane cards, a topics
table with channel-colored author pills, the tape-tail CRT).

**The measurement**, in a real Chromium against a loopback-only server:

- Requests: helm 5, fold sign-in 4 — every one to `127.0.0.1:8471`;
  `performance.getEntriesByType("resource")` filtered for external
  origins returns **empty on both** [measured]. The only console entry
  is the browser's automatic `/favicon.ico` probe against the serving
  host (404, no external fetch).
- Fonts real, not fallback: `document.fonts.status` = `loaded`;
  Archivo 400–800 and JetBrains Mono 400/700 loaded, the wordmark
  visibly condensed at `'wdth' 88` — the variable axis survived
  vendoring. JBM 500 declared-unused on these two pages, honestly
  noted [measured].
- No per-surface fork: both pages link the same `shared.css`;
  page-local `<style>` is layout-only and every value in it is a
  token.
- C3 exercised on screen: identity by label — "THE HELM" and "THE
  DOOR" strips; amber/teal held strictly to the channel semantic
  (author pills: daan amber, scribe/smith teal); no rose anywhere.

**A finding that feeds C4, not decided here:** the whole proof needed
no JS framework and no asset pipeline — the shared source is
consumable as a single CSS string plus static markup, exactly the
shape `webstyle` ships today (constitution III compatible). The design
system's JSX component layer was *not* consumed directly: the token
layer is shared verbatim, but the component layer was re-expressed as
CSS classes from the system's written specs. Server-rendered
consumption is proven cheap [measured]; direct JSX consumption remains
unproven and is C4's remaining question.

## 2026-08-13 — C1 and C2 decided: a sibling component, soulhelm — the helm

The operator decided the cockpit's home and its name.

**Home: a sibling component repo** — the sixth component, following the
consumer-position pattern the ecosystem has proven twice (the archivist,
soulfold): its own repo and constitution, a public embed seam
(`embed.Run(ctx, Options)`, the D29 pattern) that soulnode composes as a
plane by tag — the fourth URL logged at `up`, beside the door, the fold
sign-in, and the admin console — plus a standalone binary for
deployments that run the components without soulnode. The other two
candidates fell to constitutions, not taste:

- **Growing the fold's console** is blocked by the fold's own one-way
  door — "any Soulstream-only claim, endpoint, or side-channel is a
  constitution-II amendment"; a Soulstream-aware IdP collapses the
  ecosystem's two planes into one `[mechanism-argument]`.
- **Cockpit code native to soulnode** breaks "composition, not
  invention" (domain logic lands upstream — the article that has held
  through six features) and couples the human surface to the house: a
  realm running the components without soulnode would have no cockpit
  at all `[mechanism-argument]`.

The sibling shape also collapses a happy coincidence: Bar 1's
consumer-position rig (module path outside every component's namespace)
*is* the sibling repo's natural e2e gate — the bar's proof becomes the
component's standing test `[judgment]` on the final shape.

**Name: `soulhelm` — the helm** `[judgment]`: the place where the
vessel is steered and watched, which is what a cockpit is; it keeps the
one-noun family (the record, the room, the name, the house, the fold,
the helm). `soulbridge` was rejected for a real collision — soulstream's
*Later* section already names Slack/email protocol bridges — and
`soulconsole` because "console" already names the fold's admin surface.

Follow-through, named not done: the `soulhelm` component tag enters the
hq (hqlint's component list, the README constellation, the journey
conventions) when this topic graduates and the repo is founded — not
before. The topic stays `ecosystem`-tagged.

## 2026-08-13 — scoping: what already exists

The topic opens with more prior art than expected, found before any
experiment ran:

- **The design half is real.** The **Soulsystem Design System** exists as a
  Claude Design project (updated 2026-08-13, same day this topic opened):
  seven token files, seventeen components with props contracts and usage
  prompts, eighteen guideline specimens, a voice-and-tone canon, and two UI
  kits — a marketing site and a **Stream Console** (sign-in, transport-bar
  chrome, a three-column live stream with human/machine turns and a CRT
  tape readout, a reel archive, settings). The cockpit's first screens have
  in effect been sketched. Its brand rule of record: amber = human channel,
  teal = machine channel, **at deliberately equal weight** — the vision's
  "humans and AI as peers" as a color system.
- **The canon conflict is real too.** The design system is cassette
  futurism in a *light* key (warm shell tones, one dark CRT surface);
  soulfold ships a *dark* language (`internal/webstyle`, violet accent,
  rose edge, episode
  [0063](../../04-JOURNEY/0063-soulfold-the-console-gets-a-face.md))
  descended from impire.io. The design project's own readme records it was
  authored from a one-line brief with no codebase supplied — the two have
  never seen each other. C3 exists because of this.
- **Two declared substitutions in the design system** matter to Bar 4:
  fonts (Archivo / JetBrains Mono standing in) and icons (Lucide loaded
  from a CDN). Self-containment means vendoring both.
- **The entry asymmetry** motivating the topic: the MCP door gives machines
  the full tool surface; humans get the fold's identity console only.
  Soulnode's front of house (episode
  [0062](../../04-JOURNEY/0062-soulnode-the-front-of-house.md)) already
  logs door, sign-in, and console URLs — the cockpit would be the fourth
  URL on that line.
- **Sequencing note (C7):** the configure surface will render exactly what
  [`platform-tenancy-guardrails`](../platform-tenancy-guardrails/README.md)
  is deciding (accounts, grants, guardrails). Observe-first work is
  independent of it; the configure half likely trails it.
