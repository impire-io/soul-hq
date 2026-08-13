# soulsystem-cockpit — investigation journey

Topic opened 2026-08-13. Appended as the investigation happens.

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
