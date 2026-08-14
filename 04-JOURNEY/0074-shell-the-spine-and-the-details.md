# Episode 0074 — The spine and the details: the operator's screen review lands (2026-08-14)

The first real screenshots earned a review, and the review earned a
slice — episode [0073](0073-shell-the-chat-shape.md)'s queue built the
same morning:

- **The icon rail** — a slim dark spine at the far left: Home and
  Conversations up top, System status and Sign out at the bottom,
  expanding on a chevron to icons-with-labels (page-local state, no
  server round-trip). Home answers the operator's "no easy way back":
  `/home` is *Your realm at a glance* — Storage, People & sign-in,
  Talking, the conversations list — with the custody promise in the
  copy: "the shell keeps none of it."
- **The measure** — the conversation column caps at a readable width
  and centers, composer included; bubbles keep their sides within it.
- **The details panel** — People (participants derived server-side
  from the record, with message counts), Status in plain words
  ("Going on — people are talking here."), and **Waiting on** — the
  work vocabulary projected honestly: an open item reads "waiting for
  someone to pick up …", a claimed one "«name» is working on …", done
  items counted quietly. Hidden below 1180 px; live via the stream.
- A vendored favicon ends the console 404.

Verified on live screenshots before push `[measured]`: rail collapsed
and expanded, the capped column, the panel showing both waiting
states, Home reachable from the conversation. Whole gate green,
~6.7 s uncached at review `[measured]`. Named nit for the next slice:
the signed-in person renders as their raw persona id in People and
the top bar — it should say "You" / the display name (O3's fallback
showing through).

Reversal condition: none — records a completed build.

Trail: soulstream-shell `22154b2`; screenshots
`shell-v3-conversation.png` / `shell-v3-rail-expanded.png` /
`shell-v3-home.png` shown to the operator; the rail's markup kept
contribution-ready for the module contract
([`shell-module-contract`](../01-RESEARCH/shell-module-contract/README.md)).
