# 0009 — soulstream-shell: the declare surface

**Status:** designed 2026-08-28, at the operator's direction — the last
human chapter of the agents-as-infrastructure arc (episodes
[0140](../../04-JOURNEY/0140-ecosystem-the-focus-agents-as-infrastructure.md)–[0147](../../04-JOURNEY/0147-soulstream-the-thinking-house.md)).
Bar 5 of the graduated research measured the whole
declare→submit→served→answer loop drivable from a session admission in
561ms through public surfaces only; spec 014 built the serving side.
What is missing is episode 0116's disease one more time: a capability
with no human end — declaring an agent today is handwriting JSON and a
CLI verb.

## §1 The surface, in one sentence

The agents sheet grows a second lane — **Declared agents** beside the
personal paste block — where a person authors a declaration in a form,
submits it on their own session admission, and watches it come alive on
the acting screen; a placements list shows every declared agent's state
as read from the record.

## §2 The form (authoring)

The form's output IS the declaration: the same JSON `soulstream agent
submit` takes, shown in a folded "as JSON" view so there is never a
second schema and the CLI path stays learnable from the screen. Fields,
mapping 1:1 onto `declaration.Declaration`:

- **Name** (persona) and **home conversation** (topic picker over the
  board; may start a new one — the conversation-lifecycle acts exist).
- **Wakes**: mention on by default; topic/schedule/subject folded under
  "more ways to wake it" with the kind-specific fields the schema
  takes. Delivery classes surface as the words design 0005 fixed
  (a subject wake says it can miss).
- **Instructions** (optional): a stage-1 artefact reference — topic +
  artefact name pickers.
- **Model** (optional): a picker over the catalogue's *names* (§4).
  Absent means the ambient lane, said in words.
- **Tools** (optional): names from the tool catalog, riding as the
  declaration's capabilities.
- **Budget**: the defaults visible and editable, never hidden — the
  colony gate is a fact a person may read.

Validation is upstream's, surfaced inline: the form round-trips
through `declaration.Parse`/`Validate` and shows the refusal's own
words (the credential-shaped-model refusal included). No shell-side
schema duplicate.

## §3 Submit and arrival

Submit is `fleet.Submit` on the session's own admission — the measured
Bar 5 loop; the shell performs no act it could not perform through
public packages. Arrival follows 0008's principle: the act ends when
the realm's own evidence shows the thing arrived — the acting screen
watches the placement live and renders `open → claimed by <node>` as it
happens. **A realm with no dispatcher plane shows the open placement
with the truth in words** ("declared; nothing serves agents here yet —
the deployment enables the dispatcher plane"): honest waiting, never a
spinner.

## §4 The lists (reading)

- **Declared agents**: the placements topic materialised — persona,
  wake kinds, model name, state (open/claimed/owner), the declaration
  unfolded on demand. Reconstructed from the log at every render; no
  shell store (0001 §4, standing).
- **Models**: the catalogue's names, read straight from the realm KV
  as names only — the picker's source and a read-only list.
  `model set` stays the CLI's this slice; writing the catalogue from
  the shell is a named [O], not a refusal.
- **No retirement is offered.** Nothing in the ecosystem un-places a
  placement (design 0007 §9's open); the surface must not invent a
  stop it cannot perform. The state shows; the act waits for the
  vocabulary.

## §5 Boundaries, each already a house rule

- **Pure consumer.** One new dependency, named openly:
  `soulstream-workloads`' published `declaration` and `fleet` packages
  (parse/validate/submit — the exact wire format, one definition). The
  catalogue is read as KV names to avoid a second new module
  dependency. No `internal/` import anywhere (the standing purity
  gates extend, not bend).
- **No secrets.** Provider keys never pass through the shell: D36
  custody is caller-own with no cross-tree access (spec 014, wall 3),
  so the surface hands the person the paste-able
  `soulstream provider set` command instead — the easy-easy pattern
  the agents sheet already uses for wrap.
- **Plain words on screen.** "Agents", "Declared agents", "Models" —
  never the bynames.
- **No onboarding store, no dismissal state** — first-steps guidance,
  if the card grows a step for this, derives per render (0008 §2).

## §6 Acceptance criteria

1. The form emits a declaration byte-compatible with the CLI's input;
   upstream refusals (invalid wake, credential-shaped model, missing
   fields) surface inline in their own words.
2. Submit lands the placement through the session admission; the
   submitting session holds nothing afterward (close the tab, the
   agent still arrives).
3. Arrival: the acting screen renders open → claimed live from the
   record; the no-dispatcher realm shows the honest words instead.
4. The declared-agents list reconstructs from the log alone across a
   shell restart.
5. The models picker lists exactly the catalogue's names; an empty
   catalogue offers the CLI act in words.
6. Purity gates unchanged and green; no module imports another; the
   one new dependency is the two named workloads packages.
7. Empty states offer their act (0008's rule): a realm with no
   declared agents explains what declaring is and points at the form.

## §7 Amendments in place (2026-08-29 — the calm pass, episode 0157)

- **The form is the kit's three-step wizard** (shell v0.11.0-rc.7):
  Name it → Wake it → Instruct it, steps as visibility on a
  page-local signal — the form stays one and the §2 sections live
  inside their steps. The **limits stay unfolded** inside step three:
  a fold was tried and this design's own test refused it, by its own
  words — a bound nobody sees is a bound nobody knows they are
  running under. The wizard does the calming; the numbers stay on the
  screen.
- **A home that does not exist yet is started at declare time**
  (`f6e1cbb`): the home picker's first choice is "a new one, named
  after it" — on a fresh board the only choice — and the act starts
  that conversation (named after the agent) before validating and
  submitting, the same start-on-first-use the placements topic gets.
  The one-codec claim holds whole: `declarationFrom` stays pure, an
  empty topic is still invalid as a document, and what lands on the
  record always carries a real path. The JSON view says plainly that
  the home arrives at declare time rather than letting the validator
  refuse a gap the act closes; the answer names the home it made; the
  e2e reads the conversation named after the agent off the board.
