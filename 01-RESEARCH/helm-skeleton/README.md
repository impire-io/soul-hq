# What module contract turns the helm into the skeleton every human surface plugs into?

**Component:** soulhelm
**State:** active
**Started:** 2026-08-13

## Abstract

The helm shipped v0.1.0 the same day it was designed — the observe surface,
fold sessions, one act. First light, not yet the place humans live. The
operator's direction (2026-08-13): the helm becomes the surface humans use
most, and architecturally it becomes a **skeleton** — a pluggable frame in
which every human surface is a module that earns its place, activated by
what the deployment actually runs. The fold's administration is the first
module; the cockpit core itself is re-homed as one rather than being the
shell's privileged tenant; the agent designer/manager over soulrealm's
surface is the named horizon that makes the contract worth having. A
decisive answer lets M2 (design 0001 §4's configure surfaces) be built
module-shaped instead of as one-off screens; scope is held — participation
stays the named successor topic ([O4]), later arriving *as a module*
through this same contract.

## The question

**What exported module contract lets the helm host every human surface as a
module — registered through one seam, activated by what the deployment
runs, contributing navigation and screens — while the helm remains a pure
consumer that custodies nothing?**

Sub-questions that prove to need their own investigation become successor
topics rather than growing this one; the candidates are named under
*Decisions that are not bars* below.

## Pre-registered bars

Written before any experiment. The refactor spikes run in the session
scratchpad against the helm's existing rigs (the consumer-position e2e, the
browser rig, the custody scan); conclusions and principled code land in
git.

- **Bar 1 — two real modules, one contract, zero special-casing.** The
  existing observe/configure core re-homed as a module, and a new
  fold-administration module (people, invites, groups, OAuth clients over
  the fold's `/api/admin`), both registering through one exported contract
  (identity, activation predicate, navigation contribution, route
  mounting). **Pass:** both run through the same contract; no module
  identifier appears anywhere in the shell's packages (checked
  mechanically, not by eye); the existing consumer-position e2e stays
  green.
- **Bar 2 — activation follows the deployment.** The consumer-position rig
  runs two arms: fold plane present, fold plane absent. **Pass:** with the
  fold absent, the fold module contributes no navigation and its routes
  answer 404; with the fold present, its screens appear and act — and the
  module learns which arm it is in from public surfaces only (plane
  discovery, config the deployment already declares), zero helm-specific
  configuration.
- **Bar 3 — the standing doors survive the skeleton.** Design 0001 §8's
  gates run over the module-shaped build. **Pass:** gate 1 (pure consumer:
  pinned tags only, zero `replace`, `internal/` imports impossible by
  module path) and gate 2 (the custody scan clean after a full
  sign-in→act→sign-out session, positive control fired) both green; no
  module needed a privileged surface or a helm-owned store to earn its
  place.
- **Bar 4 — pluggability proven from outside.** A probe module whose module
  path sits outside the soulhelm namespace (the `e2e/embedgate` shape)
  plugs in through the exported contract alone: registers, contributes
  navigation, mounts one screen. **Pass:** it compiles against exported API
  only, its screen renders in the browser rig, and the shell needed zero
  changes to admit it.

## Decisions that are not bars

Design decisions this topic must take or route, recorded honestly rather
than dressed as measurements:

- **The fold's standalone `/admin` console** (its episode 0061): once the
  fold module exists, the operator's framing — *the fold is just another
  module; the helm is the skeleton in which modules earn their place* —
  puts the fold's human administration in the helm. Whether soulfold's own
  console retires to the ceremony pages (sign-in, enrol, error — the OIDC
  surface itself) is a soulfold scope decision, recorded at graduation.
- **The agent designer/manager** — declaring, launching, and observing
  workloads over soulrealm's surface — is the third-module horizon. It
  motivates the contract; it is not built here.
- **Participation** (posting turns, talking to agents) stays [O4], the
  named successor, decided shell-first today; it later lands as a module
  through this contract and still gates on upstream ask #1.
- **Module packaging** — modules as packages in the soulhelm repo versus
  sibling repos composed by the embedding binary (the soulnode plane
  pattern one level down) — is answered by the experiments, not
  pre-decided.

## Reversal condition

Written now: **if two real modules cannot share the shell without
module-specific shell code** — observable as the contract degenerating into
a switch over module names, or a module unable to earn its place without a
privileged surface or a helm-owned store of record — then the skeleton is
the wrong shape at this size: the helm returns to growing screens directly,
and pluggability waits for the third real consumer to make the argument
again. Design 0001's own reversal condition stays live underneath: a
required `internal/` import, a privileged surface existing only for the
helm, or a helm-owned store dissolves the component, skeleton or not.

## Verdict

*Empty until graduation.*
