# Episode 0060 — M3 ships: the fold is complete (2026-08-03)

The last milestone on the fold's roadmap landed the same morning its
research graduated: users, groups, invites, client registration, and
the `/admin` surface — the lifecycle design (D20–D24) realized, with
the research prototype as its foundation and the four bars riding
`make test` as permanent gate tests (`specs/005-the-lifecycle/`,
branch `005-the-lifecycle`, **v0.2.0 tagged**). Soulnode composed it
the same hour (`008-founding-invite`).

What the milestone made true [measured]:

- **Invitation is the only door.** Enrollment everywhere — the CLI
  fold, the embedded fold, every rig — now requires a live,
  single-use, digest-stored invite; first-touch enrollment is deleted
  from the codebase, and every pre-M3 gate (M1, M2, M4, M5, both
  consumer-position rigs) runs green on the new mechanism.
- **The M3 gate**: from nothing to a signed-in admin in four counted
  acts, the admin-role claim JWKS-verified; an admin-surface group
  change surfaces in the target's next token (engineering → platform,
  asserted exactly); non-admin bearers 403, bare requests 401.
- **The admin surface** is a JSON API under `/admin` — no pages, D9's
  inventory intact — authenticated by the fold's own admin-role
  tokens against its own published keys. Its one bearer-carrying
  response is the invite mint, shown once.
- **In the bundled shape**, the fold plane seeds the founding persona
  into `admin`+`realm` and delivers the founding invite through the
  embed seam's new `InviteSink`; `soulnode init`/`up` print the enroll
  URL once beside the founding token. The identity plane is
  indifferent on purpose: `admin` is inert at the callout (no declared
  role key names it) while `realm` keeps admitting — the fold's admin
  surface and the realm's admission stay two separate authorities.
- The pending M2 runbook was updated in place so its steps stay
  runnable under the invite rule (the human act itself is still
  pending).

With M3, **every milestone on the fold's roadmap is shipped** — M1
through M5, three design docs (D1–D24), v0.2.0 — and the soulnode
story composes it end to end: one binary, a founding invite, a passkey
in a browser, an MCP session in a realm.

Reversal condition: none — records a completed build; the lifecycle
design's own reversal conditions (episode 0059) stand.

Trail: `specs/005-the-lifecycle/` (spec, quickstart with the four-act
walk and the admin-API tour); design
[lifecycle](../02-DESIGN/soulstream-idp/lifecycle.md); soulfold `v0.2.0`;
the `005-the-lifecycle` and `008-founding-invite` merges.
