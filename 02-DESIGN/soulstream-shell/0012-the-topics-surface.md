# 0012 — soulstream-shell: the topics surface

**Status:** drafted 2026-08-30 from research topic `the-topics-surface`
(bars 1–2 measured on the rig; bars 3–4 stand below as this design's
build gates). Operator direction (2026-08-30): the conversations list
is unmanageable once a person is part of many topics, and an agent's
home topic is not a conversation — hide it from the list, offer it on
the agent's detail. This document amends 0003 §3's list paragraph and
0009 §4's declared-agents reading; the lifecycle ladder, the truthful
copy rules, and every act stand unchanged.

## §1 The problem

The rail renders `topic.Board` raw: a flat list, name and a one-word
state, ordered lexicographically by path — arbitrary, since paths end
in random suffixes. It costs what it does not show: the shipped board
pays ≥3 sequential round trips per topic every tick — measured linear
at ~0.52ms/topic on loopback, ~105ms at 200 topics before any network
RTT [measured, topic journey 2026-08-30] — and it lists the machinery:
every declared agent's home (started at declare time since 0009 §7)
and the placements topic sit among the people's conversations
indistinguishably. The home topic is the agent's operational room —
the placement work item, non-record wake outcomes, failure
self-reports, the narrow credential's only writable subject — misfiled
as a conversation.

## §2 The mechanism — one followed board (upstream ask #4)

**Upstream ask #4 (core):** the record grows a followed board — one
ordered consumer over `SOULSTREAM.TOPICS.>` serving a live projection
of `BoardEntry` plus **last activity** (each op already carries its
timestamp). The rig measured the shape decisively: one pass replays a
whole 200×20 realm in 7.5ms cold (~1.6µs/msg) against ~105ms for one
shipped `Board` call, and the warm, incremental tick is ~free
[measured]. The declarations on the placements topic ride the same
stream, so the one pass also serves §3's partition.

The shell holds the projection in memory on the shared read lane,
rebuilt from the log on connect — the tray's shape, not a store:
nothing survives a restart because nothing needs to. The 1s tick
renders from the projection and performs **zero** JetStream reads —
strengthening 0003 §5/§8 ("the tick gains no per-topic reads") to its
limit. `placementsPath`'s second board call disappears into the same
projection.

The build binds to the ask; no interim half-mechanism ships. Until the
ask lands, the surface stays as it is — a stopgap that keeps the O(n)
board on the tick would ship the disease with a new label.

## §3 The list — conversations, ordered by life

Leaning into the record's idea of topics is mechanics, not vocabulary:
on screen people still read "conversation" (0001 §7 stands; the
operator's plain-language rule holds).

- **Order:** live conversations sort by last activity, newest first —
  the projection carries the timestamp; the reverse-lexicographic
  accident retires.
- **Partition:** machinery leaves the rail entirely. The rule, measured
  correct on all 206 rig entries including both edge cases [measured]:
  *machinery = the placements topic ∪ every declared home*. Home-ness
  is a **role read from the declarations at every render, never a
  property stored on the topic** — a re-pointed fleet (submission is
  additive; nothing un-places) leaves both homes machinery, honestly.
  A personal paste-block agent's home is not on the record and cannot
  be derived — it stays in the list, and the research topic's reversal
  condition watches whether that clutters in practice.
- **The folds stand:** the archived fold at the foot (0003 §3) is
  unchanged; closed conversations show with their state in plain words
  as today.
- **Deep-open is honest:** the conversation screen renders any path it
  is given, hidden ones included — hidden from the rail is not
  unroutable. An open machinery topic says whose room it is in plain
  words and points at the agent's detail; the rail still does not list
  it.

## §4 The agent detail — the room behind the agent

The agents module grows a detail surface (the `Link` contract's
`routeAgent` deepens from a roster highlight to a place): identity
(shown-as, handle, channel, operated-by, presence as the lease reads),
the declaration's facts as 0009 renders them (wakes with their
delivery classes, model, tools, placement state and owner), and **the
home topic's thread with a composer** — the room, where it belongs.

- **Talking wakes it.** The composer posts a turn on the home topic
  *mentioning the agent* — mention is the default-on wake; a bare turn
  would reach an agent only through a topic wake it may not have. One
  outcome per turn (the engine's idempotent wake, measured in episode
  0130); the answer lands on the same thread the person is watching.
- **A no-dispatcher realm answers in words** (0009 §3's rule extends):
  the composer stands, the thread shows the turn, and the surface says
  nothing serves agents here yet — honest waiting, never a spinner.
- **The thread is the conversation module's rendering reached at the
  detail, not a second implementation** — how the two modules share it
  without breaching the module contract (0002) is the build's one
  structural decision, made against the purity gates, not around them.
- **No mention goes dark.** A mention landing on a hidden home still
  counts in the spine tally within one tick, and the mark leads
  somewhere a click can follow — the agent's row and detail badge it.
  Hidden topics stay on the board, so the tray's standing check keeps
  their slips alive [mechanism-argument].

## §5 Boundaries, each already a house rule

- **No shell store** (0001 §4): the projection is memory, rebuilt from
  the log on connect — proven restart-identical on the rig [measured].
- **Every act on the session's own admission** (0001 §6): the detail's
  composer is the person's own signed post; the read lane never
  mutates. No new act class — posting a turn is class (a) as always.
- **One upstream dependency named openly:** ask #4. No new module
  dependency; the declaration reading stays `fleet.DeclarationOf` on
  the published packages.
- **Plain words on screen:** "Conversations", "Agents", "its room" —
  never board, topic, partition, projection.

## §6 Acceptance criteria (the research bars become the gates)

1. **Scale (Bar 2):** at 200 topics × 20 ops the tick performs zero
   JetStream reads; the projection's cold build completes under 1s at
   10k ops; the list orders by last activity and holds the order live.
2. **Partition (Bar 1):** e2e with ≥5 declared agents — one persona
   declared twice with two homes, one home the placements topic itself
   — zero misclassified rows across renders and across a shell
   restart, derived from the placements reading alone.
3. **The talk wakes (Bar 3):** a turn from the detail composer yields
   exactly one outcome, rendered on the detail without refresh; the
   no-dispatcher realm shows the honest words.
4. **Nothing dark (Bar 4):** a mention into a hidden home reaches the
   spine tally within one tick and a click lands on the message.
5. Deep-opening a hidden topic renders with the whose-room note; the
   rail never lists machinery; the archived fold behaves as 0003 §6.
6. Standing gates unchanged and green: purity, custody scan, plain
   register, the e2e whole.

## §7 Open questions [O]

- **[O1]** Sub-topics: the record carries `Parent`/`ParentKnown`; the
  rail stays flat. Grouping by parentage is its own decision when
  nesting appears in practice.
- **[O2]** Tags as the person's organizing tool (filter, pin): the
  announcement carries them; no surface reads them yet. Waits for
  demand, not for symmetry.
- **[O3]** Undeclared machinery (personal agents' homes): if realms
  accumulate enough of it that the list stays cluttered despite §3,
  record-side vocabulary returns as its own research topic — the
  registered reversal condition.
- **[O4]** A realm with thousands of topics may want a screen of its
  own beyond ordering and partition (0003 §7 [O2] carried forward).
