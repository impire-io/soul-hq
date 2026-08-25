# 0008 — soulstream-shell: the first hour

**Status:** decided 2026-08-24 (episode
[0124](../../04-JOURNEY/0124-ecosystem-the-first-hour-and-the-presence-lease.md);
drafted 2026-08-23), **built 2026-08-25** as soulstream-shell
v0.11.0-rc.3 (episode
[0127](../../04-JOURNEY/0127-shell-the-first-hour-builds.md); [O1]'s
fresh-eyes install still open). The operator's ask, in one line: *most
people are lost once they installed soulstream, not knowing what to do
next.* Episode
[0116](../../04-JOURNEY/0116-ecosystem-what-shipped-without-a-human-end.md)
named features without a human end and the arc that followed closed
them (episodes
[0122](../../04-JOURNEY/0122-ecosystem-the-shell-arc-lands.md)/[0123](../../04-JOURNEY/0123-shell-the-sheet-shape.md));
this document names the same disease one level up: **the product
without a first hour**. Every screen assumes the person already knows
why they are there, and the first thing every new person sees is an
empty realm.

## §1 The problem

A person founds a realm, signs in, and stands in a house with nothing
in it: no conversations, no agents, no tools, nobody else. Each act
that would fill it is built and honest alone — the agents screen hands
one paste-able block (the easy-easy rule, kept), the tools screen adds
both halves in one act, People & sign-in mints invites, the start card
opens a conversation — but nothing sequences them, and nothing says
which comes first or why any of them matters. What guidance exists is
disconnected asides: the Home agents card's "Set one up" hint when the
roster is empty, "No conversations yet. Start one above." [measured,
code trace: `modules/overview/render.go`]. And an act that succeeds
ends on the acting screen with no evidence the added thing *arrived*:
the agent paste block ends at a terminal on another machine, and the
screen that minted it shows a roster count, not a life sign.

## §2 Guidance is a reading, never a store (normative)

The first-steps surface derives **entirely from state the shell
already reads**: the board for conversations, the agents roster
through the support layer, the tool catalog, the people roster, the
session's own grants list. It is display-grade in exactly the
catalog's sense
([`tool-catalog.md`](../soulstream-core/extensions/tool-catalog.md) —
the A10 demotion): a map of what the realm holds and what act would
fill it, never an authority on anything, recomputed at every render.

Three refusals follow, each already a house rule:

- **No onboarding store.** No "completed" flag, no dismissal state, no
  helm-owned store — 0001 §4's constraint, unchanged. The card cannot
  drift from the realm because it is not a record of anything: restart
  the shell mid-journey and the same card derives again.
- **No per-person tour.** D26 refuses per-user rows; guidance is about
  the *realm's furnishing*, not the person's education. A second
  person signing into a furnished realm sees a furnished realm — the
  board is the tour. What is legitimately per-session (whether this
  person may take an admin act, whether they have connected a tool)
  rides reads the session already owns per render (0005 §3, [O1]).
- **No wizard.** No modal sequence, no gating, no step that must be
  taken before another screen works. Every step is a link to a
  standing screen; the person who ignores the card loses nothing.

## §3 The first-steps card on Home

While the realm is young, Home leads with one card above the plane
readouts: the steps that turn an empty realm into a working one, each
derived, each a link, each vanishing as the realm fills.

| Step | Shown while | Leads to |
|---|---|---|
| **Set up your assistant** — it gets a name of its own and answers when you mention it | the agents roster is empty (deployments with the agents surface on — the declared fact the Home card already reads) | `/agents` |
| **Start a conversation** — mention your assistant by name and it answers | the board holds no topic | the start card / `/conversations` |
| **Connect a tool** — services your assistant can use on your behalf | the catalog is empty (admin sessions: add; once entries exist, sessions whose own standing shows nothing connected: connect) | `/tools` |
| **Invite someone** — they sign in with a passkey of their own | the roster holds one person (sessions the People screen offers the act to — display-plus-gate, 0005 §3's precedent; the deeper authority stays at the plane) | People & sign-in |

- **Order is the product judgment**: something answers you first
  (steps 1–2 are the product clicking), hands second, company third.
  §6 [O1] names the test this ordering must survive.
- A step whose condition the realm already satisfies renders **done**
  while any step remains — progress is visible and costs no store,
  because done is derived like everything else. When every step a
  session can see is satisfied, the card is absent, forever, with
  nothing to dismiss.
- **Register**: the person's words throughout (0007 §3's calibration —
  smart and not technical). "Assistant", "tool", "invite"; component
  bynames and machine-room vocabulary never appear; the banned-word
  gates hold over this card like any surface.

## §4 An act ends when the thing arrives (normative)

The screen where an add-act happened states the added thing's arrival
**from the realm's own evidence, live** — the approvals tick's
mechanics (0007 §5, the storage tail's shape), pointed at the thing
just added. A screen never claims what the realm cannot say.

Applied where the evidence stands today:

- **Tools** already meet this: the add writes both halves and says
  which failed; connect ends with the person's standing updated on
  `/tools` (0005 §2). The catalog stays a map, not a warranty — drift
  still fails at use, in words (D39).
- **People** already hold the half they can: the invite card is
  shown-once; the roster shows the person once they first sign in.
  Whether the roster ticks live so the inviter *watches* them arrive
  is [O4] — optional manners, same mechanics.
- **Agents** are the named pain and the gap. The wrap announces
  nothing on start [measured, code trace: `cmd/soulstream/wrap.go` —
  no profile op], so the realm's first evidence of a wrapped agent is
  **its first answer**. Two moves, honest to that:
  1. The shown-once card gains its next step in words: *start a
     conversation and mention {name} — its answer is how you know it
     is running.* The guided path makes the first evidence arrive
     fast, in the conversation, where it is worth something.
  2. The roster row carries the derived life sign: *hasn't spoken
     yet* / *last spoke {when}*, read from the record like every
     lifecycle word (0007 §5) — never a liveness the realm did not
     witness. This is the floor; where the presence face below
     stands, the row prefers its fresher word.
  **Upstream ask #3** (continuing 0001's numbering, grown 2026-08-24
  from "announce on start" by the operator's ask — *which personas
  are around*): the wrap becomes the first mandatory writer of the
  presence convention
  ([`presence.md`](../soulstream-core/extensions/presence.md), drafted
  from the thin paragraph the same day). On start it publishes the
  agent's profile (identity,
  once, on the registry) and takes up a **presence lease** on a
  per-subject-limited side face — renewed on a cadence, a farewell
  written on clean stop, and *staleness read as gone*: a crashed
  wrap never says goodbye, so **departure is derived from silence,
  never merely announced** — the fleet's own liveness posture
  ([workloads 0003 §3](../soulstream-workloads/0003-fleet.md)).
  The convention's two rules hold unbent: advisory (no guardrail, no
  act may depend on it) and off the op-log ("presence in the op-log
  is a bug"). With the lease standing, the roster says *{name} is
  in / left {when} / last seen {when}* while the person still holds
  the paste block, and the first-steps card's arrival line reads it
  live. This design does not gate on it; the two moves above stand
  without it. **Built 2026-08-24** (episode
  [0125](../../04-JOURNEY/0125-ecosystem-the-presence-lease-builds.md):
  core v0.13.0's `presence` package, the wrap wired on soulstream
  branch `011-presence-lease`) — the arrival line has a face to read.

## §5 An empty state points somewhere (normative)

Every screen whose emptiness a furnishing act would fill says, in the
person's words, what the thing is and offers the one act — the Home
agents card's shape ("none yet — Set one up — your assistant gets a
name of its own…"), made the rule. For a session the act is not
offered to, the empty state still says what would fill the screen and
whose act that is. Screens where empty is the good state (approvals'
"Nothing is waiting for a decision.") explain and offer nothing —
the rule is scoped to furnishing, not to quiet.

## §6 What this does not change

No act endpoint moves or changes shape; the card and the empty states
are renderings of reads that already ride their decided lanes (0005
§3: catalog on the shared read lane, standing on the person's own
admission, rosters on the surface's standing). No new store, no new
authority, no new privileged side-channel — the purity gates and the
custody scans hold as they stand. The record's vocabulary is
untouched; only more of its translation reaches screens. The
one-URL tool add (MCP auth discovery, dynamic client registration) is
an identity-plane research question that would make step 3 shorter —
it is not this document's, and this document does not wait for it.

## §7 Acceptance criteria

1. **Fresh realm, first sign-in**: Home leads with the first-steps
   card; taking each act flips its mark on the next render; restarting
   the shell mid-journey re-derives the identical card — the no-store
   proof, standing beside the custody scan.
2. **Furnished realm**: the card is absent; a person signing into a
   realm with agents, topics, tools, and company never sees it.
3. **The agent path end-to-end**: add → the shown-once card names the
   next step → block pasted on another machine → the agent mentioned
   in a topic → its answer arrives and the roster row's word flips
   from *hasn't spoken yet* — read from the record, live, no reload.
4. **The empty-state sweep**: every screen with a furnishing act
   offers that act when empty; approvals keeps its quiet; banned-word
   gates green; zero horizontal overflow at 1000/390 px.

## §8 Open [O]

- **[O1]** The step ordering is asserted, not measured. The test: one
  fresh-eyes install by someone who is not the operator, chafe log
  open from minute one — the dogfood discipline pointed at the first
  hour. If the person reaches for a different first act, the order
  follows the person.
- **[O2]** Upstream ask #3's substance is a core-extension decision,
  not this document's:
  [`presence.md`](../soulstream-core/extensions/presence.md) (decided
  with this doc, episode 0124) — **one face for everything
  that runs**, the lease and its reader semantics, and the payload
  line (words a person reads, never numbers a machine steers by). A
  standing face also answers 0005 [O2] (running-state on workload
  tool rows) without waiting for the record's declaration vocabulary.
  This design reads whatever that convention decides and stands
  whether it lands ahead of this build, with it, or after.
- **[O3]** Whether step 3 should appear for non-admin sessions in a
  realm whose catalog is empty ("ask your admin to add one") or stay
  silent until there is something to connect. Drafted: silent — a
  step nobody on screen can take is a wall, not a door.
- **[O4]** The People roster's live tick, so an inviter watches the
  invited arrive. Same mechanics as approvals; by chafe.
