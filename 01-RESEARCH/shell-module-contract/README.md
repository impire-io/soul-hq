# What module contract makes soulstream-shell a pure shell every human surface plugs into?

**Component:** shell
**State:** active
**Started:** 2026-08-13

*(The rename executed 2026-08-13: the repo is
[impire-io/soulstream-shell](https://github.com/impire-io/soulstream-shell),
v0.2.0 — episode 0070. The Component tag uses the short episode
vocabulary: `shell`.)*

## Abstract

The component shipped v0.1.0 the same day it was designed — the observe
surface, fold sessions, one act. First light, not yet the place humans
live. The operator's direction (2026-08-13): it becomes
**soulstream-shell**, the surface humans use most and architecturally a
**pure shell** — a modular
frame that contains **zero module logic** and is **agnostic from
Soulstream by contract**: the shell's packages import nothing
Soulstream-specific, ever, checked mechanically. Every human surface —
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

**What exported module contract lets soulstream-shell host every human
surface as an externally-built module — registered through one seam, activated by
what the deployment runs, cross-linked through the shell — while the
shell itself imports nothing Soulstream-specific and custodies nothing?**

Sub-questions that prove to need their own investigation become successor
topics rather than growing this one; the candidates are named under
*Decisions that are not bars* below.

## Pre-registered bars

Written before any experiment. The spikes run in the session scratchpad
against the component's existing rigs (the consumer-position e2e, the
browser rig, the custody scan); conclusions and principled code land in
git.

- **Bar 1 — a pure, agnostic shell, proven by two real modules.** The
  existing observe/configure core re-homed as a module, and the new
  **collaboration module** (post, reply, comment, open topics, mention
  notifications — the usable-cockpit value of episode 0071; amended
  pre-experiment from the fold-administration module, which waits as a
  later module), both registering through one exported
  contract (identity, activation predicate, navigation contribution,
  route mounting). **Pass:** both run through the same contract, and the
  shell's packages import **no module and no Soulstream component** —
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
  lives in the Soulstream module-support layer, never in the shell; no
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

- **The name.** Decided 2026-08-13, twice in one day: *helm* fell for
  its collision with Helm charts; bare *cockpit* was refused for
  colliding with cockpit-project (Red Hat's browser server console —
  the same product category); *soulshell* won the morning and was
  superseded the same day when the collision-avoidance argument
  generalized into the ecosystem-wide **soulstream-*** scheme — the
  component is **soulstream-shell** (episode
  [0069](../../04-JOURNEY/0069-ecosystem-one-name-soulstream.md)).
  Execution rides the ecosystem rename sweep: repo + module path +
  fresh tag; the product's plane key with the pin bump; the hq sweep
  (design folder, roadmap, component lists, hqlint) — landed so the
  docs keep describing what is.
- **Where module code lives.** Decided by the operator: **never in the
  shell** — the shell contains no module logic at all. The remaining
  open arm — a component's module in that component's own repo (the fold
  growing a shell-module package, a dependency-direction question
  against its constitution) versus sibling module repos versus the
  composition layer — is answered by the experiments.
- **The fold's standalone `/admin` console** (its episode 0061): once
  the fold module exists, whether soulstream-idp's own console retires to the
  ceremony pages (sign-in, enrol, error — the OIDC surface itself) is a
  soulstream-idp scope decision, recorded at graduation.
- **The agent designer/manager** — declaring, launching, and observing
  workloads over soulstream-workloads' surface — is the third-module horizon. It
  motivates the contract; it is not built here.
- **Participation** entered scope 2026-08-14 (episode 0071, amending
  the shell-first call recorded here the day before): collaborating
  from the cockpit is the value center and rides the backend-held
  session admission — no upstream ask. The *browser-native* client
  (upstream ask #1) stays a separate parked horizon.
- **Extraction.** A standalone generic framework is founded only when a
  second real consumer outside Soulstream exists; Bar 1's agnosticism
  check is what makes that extraction cheap when its day comes.

## Reversal condition

Written now: **if two real modules cannot share the shell without
module-specific shell code** — observable as the contract degenerating
into a switch over module names, or a module unable to earn its place
without a privileged surface or a shell-owned store of record — then the
shell is the wrong shape at this size: the component returns to growing
screens directly, and pluggability waits for the third real consumer to
make the argument again. Narrower: **if the shell cannot stay
Soulstream-agnostic** — observable as a real module need the contract
cannot express without importing a Soulstream shape — then the
agnosticism constraint falls (recorded with that concrete need as the
evidence), not necessarily the module idea. Design 0001's own reversal
condition stays live underneath: a required `internal/` import, a
privileged surface existing only for this component, or a component-owned
store of record dissolves it, shell or not.

## Verdict

All four bars **PASS**, measured 2026-08-14 on real builds:

- **Bar 1 — PASS** `[measured]`: shell/ (pure frame + exported
  contract), soulstream/ (support), modules/overview +
  modules/conversations through one contract; `internal/purity` walks
  `go list -deps`, fails on any module/support/component path, its
  control fired on 7 dependencies; served HTML byte-identical across
  all six screens; e2e ~9.5 s uncached. Shell's whole external cost:
  go-oidc, oauth2, go-jose.
- **Bar 2 — PASS** `[measured]`: modules/admin ("People & sign-in")
  activates on one declared fact (the product's optional `AdminBase`);
  present arm 9.19 s (people listed, a real invite issued), absent arm
  3.34 s (external AS, no rail entry, 404s, everything else working).
  Authority delegated (the person's own bearer); the idp needed zero
  changes.
- **Bar 3 — PASS** `[measured]`: on published tags only (product
  v0.9.0 → shell v0.3.0, core v0.8.1, idp v0.4.0, mcp v0.1.0), sole
  replaces the repo-under-test and the never-published probe; custody
  scan with fired control inside the full ceremony; 13.5 s uncached.
- **Bar 4 — PASS** `[measured]`: the link facility resolves through
  the registry of active modules (present navigates, absent degrades
  to plain text, no module imports another); `e2e/moduleprobe`
  (outside even the impire-io namespace) composes through the exported
  contract alone, zero shell changes; the outside-module import test's
  control fired on 45 packages.

The reversal condition's observable stayed quiet throughout: no
special-casing entered the shell `[measured: the purity graph]`.
Graduates to design.
