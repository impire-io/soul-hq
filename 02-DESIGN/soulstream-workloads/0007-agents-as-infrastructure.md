# 0007 — Agents as infrastructure: the standing dispatcher

**Graduated from** research topic `agents-as-infrastructure`
([episode 0141](../../04-JOURNEY/0141-ecosystem-agents-as-infrastructure.md));
all five bars measured. This document specifies the serve arm design
[`0004-wrap.md`](0004-wrap.md) §9 promised back when its reversal
condition fired: a realm operating declared agents as infrastructure —
nobody's laptop, centrally credentialed — over the same engine, with
design [`0003-fleet.md`](0003-fleet.md)'s placement answering which
node serves a wake. Evidence classes: `[V]` measured in the graduation
spikes; `[O]` open, decided at spec time or by the operator.

## 1. What must exist [V]

A **dispatcher**: a standing process (one per fleet node) that makes
submit-and-forget real. Its whole loop, every part a shipped mechanism:

- **Watch** the placement topic (live subscription with materialise as
  catch-up — the wrap engine's own pattern; the spike's poll is not the
  build shape).
- **Race** open placements through the 0003 claim path: `ClaimWork`,
  re-materialise, serve only if the read-back names this node owner —
  first claim in stream order wins, the rest void `[V — contested 2/2
  split, one live claim per item, every run]`.
- **Resume on start** every placement the log says this node already
  owns — no new op, no handshake, no local state: the record is the
  position `[V — a fresh instance resumed and deduped in one poll]`.
- **Serve** each owned placement by running the wrap engine
  (`wrap.DeclaredConfig` + `wrap.Wrapper`) for the declared agent on a
  connection bound to that agent's persona. The engine brings
  catch-up, exactly-once outcomes, and the 0006 budget at admission
  with it `[V — cycle halted at declared MaxHops, delegation clean]`.
- **Answer probes** for owned placements and **sweep** peers' silent
  claims (0003 §6 unchanged — projection nominates, probe vetoes,
  ordinary abandon decides) `[V — failover served in ~1.05s at a 1s
  bound; live owners never reclaimed; zero probe ops on the stream]`.
- **Resolve inference credentials at wake time** (§4).

Deliberately absent, unchanged from 0004: consumer state beside the
log, a coordinator, any new realm vocabulary. Submission IS
`fleet.Submit` — an ordinary work item carrying the declaration.

## 2. The serve seam [V — resolved (b), built]

`fleet.Node.TryPlace` hardwires `Runner.Launch` — correct for backend
workloads, wrong for engine-served agents. **Resolved at the spec pass
(specs/011, episode 0143) as (b): the dispatcher owns its claim path**
— claim + read-back + serve beside `fleet.Node`'s probe/sweep halves
(which need no Runner `[V]`); `fleet` is untouched, so a dispatcher
node and a runner node share one realm without either knowing about
the other. The self-selection line is `dispatcher.Servable`: an agent
with a wake set is engine-served, everything else stays the Runner
path's — pinned by a standing test. The build added one node-local
knob the design had not named: `RaceBackoff`, so a declaration a node
cannot serve does not become a claim/abandon spin on the record — it
delays a decision, never makes one.

## 3. The declaration grows `inference` [O]

The spike rode `args` (`--provider anthropic`); the build gives the
requirement its first-class name (0005 §5's discipline — names, never
grants, nothing in the declaration can widen anything):

```json
"inference": { "provider": "anthropic", "model": "claude-sonnet-5" }
```

- Names resolve node-side; a declaration carrying a credential-shaped
  value refuses by name (pre-v1 clean break, no compatibility shim).
- **Amendment 2026-08-28** — the successor clause fired the day of
  graduation: at the operator's direction the inference story became
  its own research topic, `inference-plane` (a fleet of stateless
  single-model instances behind capability subjects; the model an
  instance attribute resolved client-side, never a request
  parameter). If it graduates, this block's names resolve against the
  **plane's catalogue** (virtual model names / capability + tags)
  rather than directly to a provider key, and `provider` here becomes
  the fallback vocabulary. The dispatcher's spec pass should hold
  this block's schema until that topic answers.
- Absent block: the wrap default — the harness's own ambient
  authentication (a person's signed-in assistant). A dispatcher
  serving a declaration with no `inference` block and no ambient lane
  refuses the placement loudly at claim time, never half-serves.
- Whether `model` rides the harness invocation as a template variable
  (`{{MODEL}}`) is settled at spec time with the template's owner.

## 4. Inference credentials: custody and injection [V]

- **The lane is the D36 secret store** — an infrastructure inference
  key is deployment configuration. The dispatcher resolves
  `providers/<provider>` from **its own** secrets tree at wake time
  `[V — 1.98ms per resolve, one secrets.get per wake in the audit]`
  and injects it into the harness's environment via the shipped
  `wrap.Template.Env` seam `[V — a real subprocess read and matched
  it]`; env API keys are the non-interactive lane real harnesses
  already accept [mechanism-argument].
- **Custody is structural** `[V]`: no canonical scope carries a
  secrets op-tail — the agent scope reaches no identity-plane subject
  at all, and the persona scope's tails are sign/keys/grants/
  approvals/seal only. All three probes (agent scope own-prefix,
  persona scope own-prefix, cross-prefix) die at the server as
  permissions violations with zero service decisions. The build adds
  no policy here; it must simply not widen any template.
- The secret value may exist in exactly two places: the sealed store,
  and the harness process's environment for the run's duration. Never
  in the declaration, the record, a run artifact, or an outcome
  `[V — census clean]`.
- **The grants broker is the other lane, by name** [O]: an agent
  thinking on a *person's* provider account is outbound identity
  (D30–D34), not deployment config. Out of scope for the first build;
  the boundary is recorded so nobody reaches for the secret store to
  fake it.
- Secret naming (`providers/<name>`), rotation, and per-tenant trees
  are spec-time decisions [O].
- **Amendment 2026-08-28** — research `inference-plane` (opened at the
  operator's direction) tests moving custody one layer further from
  the agent: provider credentials living with **inference-plane
  instances**, the harness reaching the plane with realm credentials
  only, and this section's env-injection demoting to the fallback for
  provider-native harness features the plane cannot carry. The
  measured mechanics here (wake-time resolve, `Template.Env`,
  structural scope denials) stand either way — the question is what
  the injected credential unlocks, not how it travels.

## 5. Credentials for the served agent [V→O]

Per served placement the dispatcher needs an engine connection bound
to the agent's persona. The measured shape `[V]`: a D28
`mint.ephemeral` against the deployment's persona-scope role — the
engine ran a full declared-agent wake under a minted canonical
persona-scope credential on an operator-mode server. The capability
credential (spec 010's agent scope, for the harness's tool door) is
unchanged. Open [O]: which role key the founding installs for the
dispatcher's engine mints, and the credential's TTL/renewal cadence
across long-lived serves — the product's founding ceremony decides
(soulstream spec, not this repo).

**The build sharpened the [O] (episode 0143, finding 3): the tool
door's credential is per-agent while the engine template is
per-node.** `wrap.Template.MCPEnv` carries the persona and its
credential, so one dispatcher serving many personas from one base
config cannot give each agent's tool door its own authority.
`ConnectAgent` covers the engine's connection only. The shipped
`dispatcher serve` builds door lanes carrying URL+realm alone —
degraded but safe (the door can impersonate nobody; replies always
post on the agent's own client). Before a real harness is wired under
the dispatcher, the engine seam must become per-persona — either
`ConnectAgent` returning an engine config beside the client, or a
second `EngineFor(ctx, persona)` hook — decided with the product
wiring, where the spec-010 agent-scope mint is at hand.

## 6. Drain and crash are different ends [V]

- **Crash** (connections dropped): nothing posts — the half-finished
  wake's outcome attempt dies with the connection, and the successor
  (a restart, or a peer via reclaim) re-serves it exactly once on the
  deterministic outcome id `[V — 0 outcomes while dead, then exactly
  1]`.
- **Drain** (context cancelled, engines waited): an in-flight harness
  returns a failure and the engine posts the failure self-report — the
  agent's own testimony, by the engine's contract `[V — kind=
  self_reported observed]`.

The dispatcher's stop ceremony therefore chooses deliberately: drain
for operator-intended stops (the record hears the truth), abrupt exit
for supervision-restart paths (the successor speaks instead). A stop
that accidentally drains mid-deploy would post spurious failure
testimony — the build treats the choice as config, not chance.

Two sharpenings from the build (episode 0143): the drain-time
self-report is a **`Retries:1` property** — with retries configured, a
wake cancelled between attempts parks instead of self-reporting, and
the successor speaks; and **an engine that stops on its own leaves the
served set**, silencing this node's probe answers for that placement,
so a peer's sweep reclaims it into a fresh race while the local
backoff keeps the node's own retry off a hot loop — both paths
converge on the record [the build's reading, adopted as the design's].

## 7. What the shell gets [V, module design at build]

Bar 5 proved the whole declare→submit→served→answer loop drivable from
a minted persona-scope session admission (561ms end to end) through
public ops and published packages only — no shell-only upstream
surface, no `internal/` import. The **declare surface** (author a
declaration, submit it, watch it come alive) is therefore a pure
class-(a) shell module, designed at its build against `fleet.Submit`'s
shape and this document. The refusal arm stands: an agent-scope
credential cannot even publish a submission `[V]`.

## 8. Acceptance criteria (the graduation bars as standing tests)

1. Submit-and-forget: submitter gone, mention answered exactly once;
   dispatcher restart resumes from the log with no re-claim and no
   duplicate; hard-kill mid-run leaves zero outcomes, restart serves
   exactly once.
2. Two nodes: every contested placement one owner, one live claim;
   failover reclaims `claim,abandon,claim`, a wake in the window
   answered exactly once by the survivor; zero probe ops on-stream.
3. The declared budget halts the uncooperative cycle at its bound,
   op-lessly and loudly, through the dispatcher path; the legitimate
   delegation completes with zero refusals under defaults.
4. The provider secret resolves at wake time from the dispatcher's
   tree; all three scope probes server-denied; the value absent from
   record, declaration, and run artifacts.
5. The loop from a session admission end to end; the agent-scope
   submission refused by the server.

## 9. Open, named [O]

- ~~§2 the serve seam's shape~~ — resolved (b) and BUILT
  (`specs/011-dispatcher`, episode 0143).
- §3 the `inference` block's exact schema and the `{{MODEL}}`
  template variable — the plane graduated (episode 0142); the block
  closes against its catalogue's names at the product wiring.
- §4 secret naming/rotation/per-tenant trees; the grants-broker lane
  for person-owned provider accounts — its own demand gate.
- §5 the founding's role naming and engine-credential TTL/renewal —
  the product's spec.
- §7 the shell declare-surface module design — at its build.
- ~~Live subscription on the placement topic~~ — paid in the build:
  the dispatcher watches live with materialise-poll as catch-up only.
- **`artifact` in the declaration schema** (episode 0143, finding 2):
  required by `Validate`, meaningless for engine-served agents — the
  engine runs the node's harness template, never the declared
  executable. Candidate: optional when a wake set is present. A
  pre-v1 clean break for the declaration spec, taken deliberately,
  not silently.
- **No retirement path** (episode 0143, finding 6): nothing un-places
  a placement — the dispatcher never posts `work.done`, and
  `fleet.Release` serves the Runner path only. Retiring a declared
  agent (and what it means for its persona and topics) needs its own
  vocabulary decision.
- **The zombie cap** (design 0003 §3, still open, inherited): an
  engine that answers probes but is wedged suppresses reclaim
  indefinitely.
