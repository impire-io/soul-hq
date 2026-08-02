# How does the first admin's first passkey come to exist — and how does anyone after them?

**Component:** soulfold
**State:** active
**Started:** 2026-08-03

## Abstract

M3 (the lifecycle) is the fold's last milestone, and its gate names a
from-nothing bootstrap: the first admin enrolled and signed in "in a
counted, documented number of acts". Today the fold's interim is
first-touch enrollment — the account's first authentication binds the
passkey, loud in every quickstart — which is honest for a loopback dev
realm and indefensible on a hostile network. This topic decides the
enrollment trust story: how the first admin's passkey is born, how
invites carry enrollment rights to everyone after, what an invite
token *is* (custody, lifetime, single-use), and what the fold's
equivalent of soulidentity's first-key ceremony looks like. The
pocket-id surface audit (which admin capabilities a real deployment
needs, held against constitution III) rides alongside as the scope
input for M3's admin surface.

## The question

What enrollment-trust mechanism replaces first-touch enrollment —
for the first admin (from nothing) and for every later user (from an
invite) — such that possession of the fold's store never suffices to
enroll, and the whole story fits the ecosystem's ceremony discipline?

## Pre-registered bars

- **Bar 1 — the from-nothing ceremony is counted and closed.**
  Protocol: a rig founds a fresh fold and walks the first admin from
  zero to signed-in-with-a-passkey; every act is enumerated (commands
  run, secrets shown, their custody). Pass: the count is fixed and
  documented; no act leaves a reusable secret behind (whatever
  bootstrap credential exists is single-use or expires, demonstrated
  by replaying it and being refused); the ceremony ends with
  first-touch enrollment structurally disabled for that fold
  [measured].
- **Bar 2 — invites are bearer-shaped and D12-honest.** Protocol: the
  invite mechanism in the rig — mint an invite for a username, enroll
  through it, then attack: replay the consumed invite, present an
  expired one, present a forged one, and scan the store for the
  invite's verbatim secret. Pass: exactly one enrollment per invite;
  expired and forged refuse with zero state change; the store holds
  only a digest of the invite token (D12's rule extended), verified
  with a positive control [measured].
- **Bar 3 — the store alone cannot enroll.** Protocol: Bar 2's rig
  stopped; the attacker holds the complete store directory and the
  seal seed (the envelope open is conceded — this bar is about
  enrollment trust, not confidentiality). Pass: no artifact in the
  store suffices to complete an enrollment ceremony against a running
  fold (invite digests don't invert; no credential secret exists to
  lift — constitution I's line extended to enrollment) [measured].
- **Bar 4 — the M3 surface is the smallest that serves a real
  deployment.** Protocol: the pocket-id capability audit — enumerate
  its admin surface, mark each capability needed/deferred/refused for
  the fold's actual deployment shapes (single-binary bundled,
  standalone), with the reason. Pass: the audit table exists with
  every row justified against constitution III, and M3's spec scope is
  exactly the "needed" rows [judgment, documented].

## Reversal condition

If every mechanism that closes first-touch enrollment demonstrably
requires either an out-of-band secret channel the single-binary
distribution cannot provide (observable: Bar 1's ceremony cannot be
completed with only the terminal that ran `soulnode init` plus a
browser), or a standing secret whose custody story collapses to "in
the store beside the records" (Bar 3 fails for every candidate), then
first-touch enrollment survives as a *stated deployment mode* for
loopback/trusted networks — renamed honestly, never silently — and
M3's gate is amended openly to say so.

## Verdict

<Empty until graduation. Filled by /research-graduate: PASS/FAIL per bar with the
honest numbers, each load-bearing claim tagged [measured] / [mechanism-argument]
/ [judgment].>
