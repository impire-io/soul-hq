# What module contract makes soulshell a pure shell every human surface plugs into?

**Component:** soulhelm
**State:** active
**Started:** 2026-08-13

*(The component renames **soulhelm → soulshell — the shell** — decided
2026-08-13, execution pending; the Component field above carries the tag
that is legal until the rename sweep lands. See the naming decision below.)*

## Abstract

The component shipped v0.1.0 the same day it was designed — the observe
surface, fold sessions, one act. First light, not yet the place humans
live. The operator's direction (2026-08-13): it becomes **soulshell**, the
surface humans use most and architecturally a **pure shell** — a modular
frame that contains **zero module logic** and is **agnostic from
soulsystem by contract**: the shell's packages import nothing
soulsystem-specific, ever, checked mechanically. Every human surface —
the observe/cockpit core, the fold's administration, the agent
designer/manager to come — is a module built *beside* the shell that
plugs in through one exported contract; the shell's own job is
composition — registration, activation, navigation, sessions, and the
cross-module facilities (cross-linking) that make separate modules feel
like one product. A decisive answer lets M2 (design 0001 §4's configure
surfaces) be built module-shaped instead of as one-off screens. Scope is
held: participation stays the named successor topic ([O4]), later
arriving as a module through this same contract; extraction of the shell
into a standalone generic framework is a *later, cheap* act gated on a
second real consumer — the agnosticism bar is what keeps it cheap.

## The question

**What exported module contract lets soulshell host every human surface
as an externally-built module — registered through one seam, activated by
what the deployment runs, cross-linked through the shell — while the
shell itself imports nothing soulsystem-specific and custodies nothing?**

Sub-questions that prove to need their own investigation become successor
topics rather than growing this one; the candidates are named under
*Decisions that are not bars* below.

## Pre-registered bars

Written before any experiment. The spikes run in the session scratchpad
against the component's existing rigs (the consumer-position e2e, the
browser rig, the custody scan); conclusions and principled code land in
git.

- **Bar 1 — a pure, agnostic shell, proven by two real modules.** The
  existing observe/configure core re-homed as a module, and a new
  fold-administration module (people, invites, groups, OAuth clients over
  the fold's `/api/admin`), both registering through one exported
  contract (identity, activation predicate, navigation contribution,
  route mounting). **Pass:** both run through the same contract, and the
  shell's packages import **no module and no soulsystem component** —
  checked mechanically from the import graph (`go list -deps` over the
  shell packages), not by eye; the existing consumer-position e2e stays
  green on the composed build.
- **Bar 2 — activation follows the deployment.** The consumer-position
  rig runs two arms: fold plane present, fold plane absent. **Pass:**
  with the fold absent, the fold module contributes no navigation and its
  routes answer 404; with the fold present, its screens appear and act —
  and the module learns which arm it is in from public surfaces only
  (plane discovery, config the deployment already declares), zero
  shell-specific configuration.
- **Bar 3 — the standing doors survive the shell.** Design 0001 §8's
  gates run over the module-shaped build, at the composed level.
  **Pass:** gate 1 (pure consumer: pinned tags only, zero `replace`,
  `internal/` imports impossible by module path) and gate 2 (the custody
  scan clean after a full sign-in→act→sign-out session, positive control
  fired) both green; sessions stay in memory; the NATS-admission
  machinery (bearer → sentinel + callout → per-principal connection)
  lives in the soulsystem module-support layer, never in the shell; no
  module needed a privileged surface or a shell-owned store to earn its
  place.
- **Bar 4 — the shell composes.** Two halves, both required. *Cross-
  linking:* module A renders a link into module B's screen through the
  shell's facility without importing B; with B active the link resolves,
  with B absent it degrades honestly (hidden or named-unavailable) —
  both arms measured in the rig. *Outside pluggability:* a probe module
  whose module path sits outside the component's namespace (the
  `e2e/embedgate` shape) registers, contributes navigation, and renders
  one screen through the exported contract alone — compiling against
  exported API only, with zero shell changes.

## Decisions that are not bars

Design decisions this topic must take or route, recorded honestly rather
than dressed as measurements:

- **The name.** Decided 2026-08-13: **soulhelm → soulshell — the
  shell**. The trail, honestly: *helm* fell for its collision with Helm
  charts; bare *cockpit* was refused for colliding with cockpit-project
  (Red Hat's browser server console — the same product category);
  *soulshell* names the architecture itself and stays in the
  constellation's byname family. Execution is the next work item: repo
  rename + module path + fresh tag; soulnode's `planes.helm` →
  `planes.shell` with the pin bump; the hq sweep (design folder,
  roadmap, component lists, hqlint) and a journey episode — landed
  together so the docs keep describing what is.
- **Where module code lives.** Decided by the operator: **never in the
  shell** — the shell contains no module logic at all. The remaining
  open arm — a component's module in that component's own repo (the fold
  growing a shell-module package, a dependency-direction question
  against its constitution) versus sibling module repos versus the
  composition layer — is answered by the experiments.
- **The fold's standalone `/admin` console** (its episode 0061): once
  the fold module exists, whether soulfold's own console retires to the
  ceremony pages (sign-in, enrol, error — the OIDC surface itself) is a
  soulfold scope decision, recorded at graduation.
- **The agent designer/manager** — declaring, launching, and observing
  workloads over soulrealm's surface — is the third-module horizon. It
  motivates the contract; it is not built here.
- **Participation** (posting turns, talking to agents) stays [O4], the
  named successor, decided shell-first; it later lands as a module
  through this contract and still gates on upstream ask #1.
- **Extraction.** A standalone generic framework is founded only when a
  second real consumer outside soulsystem exists; Bar 1's agnosticism
  check is what makes that extraction cheap when its day comes.

## Reversal condition

Written now: **if two real modules cannot share the shell without
module-specific shell code** — observable as the contract degenerating
into a switch over module names, or a module unable to earn its place
without a privileged surface or a shell-owned store of record — then the
shell is the wrong shape at this size: the component returns to growing
screens directly, and pluggability waits for the third real consumer to
make the argument again. Narrower: **if the shell cannot stay
soulsystem-agnostic** — observable as a real module need the contract
cannot express without importing a soulsystem shape — then the
agnosticism constraint falls (recorded with that concrete need as the
evidence), not necessarily the module idea. Design 0001's own reversal
condition stays live underneath: a required `internal/` import, a
privileged surface existing only for this component, or a component-owned
store of record dissolves it, shell or not.

## Verdict

*Empty until graduation.*
