# Episode 0090 — The fold wears the ecosystem's canon (2026-08-15)

The operator noticed what episode 0089's live proof had walked straight
through: the passkey pages — sign-in, enrolment, the admin console —
looked like a different product. They were: the fold's `webstyle`
carried the design system's *earlier* iteration (dark world, violet
accent with a rose edge), while the source of truth became the
Soulsystem light canon the shell vendored on 2026-08-14. Nothing in
the fold's design pinned the dark look — D9 pins the page inventory
and the page-local-JS rule, not the clothes [mechanism-argument: the
design docs were checked, no D-number names the theme].

The fix is design D30 (recorded in
[`session-and-ui.md`](../02-DESIGN/soulstream-idp/session-and-ui.md)):
the fold's one shared style block now implements the canon — shell
tones over the grain texture, warm charcoal ink, molded amber keys
with engraved lettering and 2px press travel, hairline-and-etch
edges, mono label strips, pills for tags and never for buttons, and
the canon's one dark surface (CRT glass with scanlines) reserved for
the admin console's shown-once token reveal. The inline-string rule
(idp article III, no asset pipeline) holds, so the type faces are the
canon's fallback stacks — upgraded where installed, honest system
faces where not. Same two pages, same classes, one file rewritten:
`internal/webstyle/webstyle.go` (+112/−72).

Verified by eye against a running realm [measured: screenshots at
1440px]: the sign-in card, and the admin console — people table,
invite lane, OAuth clients — now read as the same instrument as the
shell console beside them. The idp gate ran green including its e2e
and embed-gate suites; the product pins v0.4.2 and its own gate ran
green against the published tag.

Reversal condition: if the fold ever needs the canon's full type
faces (a measured legibility or brand need), the inline-string rule
yields to one embedded-font route before D30 reopens (recorded in the
design entry).

Trail: design
[`session-and-ui.md` D30](../02-DESIGN/soulstream-idp/session-and-ui.md);
commits — soulstream-idp `e0f3fb7`, tag v0.4.2; soulstream `99381ce`
(pin bump).
