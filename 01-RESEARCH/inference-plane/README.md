# The inference plane — can the realm serve its agents' thinking as a fleet of stateless, single-model instances, with the record staying the only record?

**Component:** ecosystem
**State:** active
**Started:** 2026-08-28

## Abstract

Episode [0141](../../04-JOURNEY/0141-ecosystem-agents-as-infrastructure.md)
closed the dispatcher question with the simplest inference lane that
could work: the declaration names a provider, the dispatcher resolves
one API key from custody and injects it into the harness's environment.
The successor clause armed there has now fired at the operator's
direction (2026-08-28): the inference story deserves a **plane**, not a
config knob. The adopted shape: inference is a realm service — a fleet
of **stateless, single-model instances**, each wrapping exactly one
model behind one adapter and subscribing only to the **capability
subjects** it can serve; the subject says *what* you want, never which
model; *which model* is an instance attribute a caller **resolves** for
and may pin by unicast, and *how hard to try* is a request parameter,
never a route. Routing belongs to the substrate (queue-group anycast)
and the client (resolve → filter → pin) — the plane contains no router.
Provider credentials live with the instances, so a harness thinks
**credential-free**: it reaches the plane with realm credentials only,
and metering rides headers the middleware can read without touching a
body. One deliberate divergence from anything this shape has done
elsewhere: Soulstream already *has* the conversation store — the
record. The plane must hold **no conversation state**: context travels
complete in each request, assembled from the record, and nothing an
agent says lives anywhere but topics (the substrate boundary,
constitution-grade). A decisive answer gives the dispatcher build its
real inference arm, gives declared agents model/effort vocabulary that
is names-not-grants all the way down, and gives the realm one metered,
custodied door to every provider.

## The question

What is the shape of the realm's inference plane — subjects, reply
grammar, instance resolution, credential custody, and the harness's
path into it — such that instances stay stateless per request (the
property queue-group balancing depends on), the model is never a
request parameter, no provider credential is reachable from any agent,
and the record remains the only store of conversation?

## Pre-registered bars

Written before any experiment runs; consumer-position rigs on
operator-mode servers under real scope enforcement, the graduation
spikes' standard.

- **Bar 1 — the fleet shape holds on our substrate.** Two stand-in
  instances wrapping different "models" serve one capability subject:
  anycast (the queue group) picks exactly one instance per request and
  both serve over a run of requests; the micro discovery surface
  returns each instance's metadata (model, formats, tags) and
  resolve → filter → unicast reaches exactly the pinned instance
  (a queue group of one); a capability nobody serves answers
  no-responders — the routing layer telling the truth for free.
  Protocol: an operator-mode rig with the instances' subjects
  scope-enforced; counts and metadata read from the wire, not the rig.
- **Bar 2 — the reply grammar is one client loop.** Streaming is
  content frames (monotonic sequence) terminated by an
  **empty-payload sentinel** whose headers carry status, last
  sequence, stop reason, and usage — meaning in headers, emptiness
  selecting the control plane, the substrate's own house style. A
  one-shot is a single reply with usage headers; a one-shot result
  too large for the wire streams instead, the first frame telling the
  caller which case it got; an error is terminal and partial output
  stands. Protocol: one client code path measured against all four
  arms (stream, one-shot, oversized one-shot, mid-stream error), plus
  the invariant checks (a terminator never carries content; a content
  frame never carries status).
- **Bar 3 — the harness thinks credential-free.** A dispatcher-served
  agent's harness reaches the plane with **realm credentials only**:
  the provider credential lives in the serving instance's own custody
  tree (the D36 lane the 0141 graduation measured) and the 0118-standard
  custody scan over everything the agent reads — declaration, record,
  environment, run artifacts — finds no provider material; the
  env-injection lane from design 0007 §4 demotes to the fallback for
  provider-native harness features the plane cannot carry. Protocol:
  a real subprocess completes an inference round trip through the
  plane while the scan and the scope probes (server-denied, the
  structural-custody standard) run against it.
- **Bar 4 — names, not routes.** The declaration's inference
  requirement is a **name** (a virtual model name or a capability +
  tag predicate), resolved node- or client-side into anycast or
  resolve-and-pin; effort/thinking ride request parameters; the
  concrete model appears in no subject and no request header.
  Re-pointing the name — the catalogue changes which instances
  satisfy it — moves the traffic with **zero declaration change**
  [the discriminating observable]. Protocol: one declared agent, one
  name, two catalogue states, the serving instance read from the
  response's informative headers.
- **Bar 5 — the record stays the only record.** After a multi-turn
  exchange through the plane, a census of every plane-side subject,
  stream, and bucket finds **no stored turn anywhere but the topic**:
  context traveled complete in each request, assembled from the
  record's materialisation, and the plane held nothing. Protocol: the
  census plus a restart of every instance mid-conversation — the next
  turn completes unchanged, proving nothing pinned and nothing
  cached mattered.

## Reversal condition

Observable readings that reverse the adopted shape, written now:

- **Scope enforcement and the fleet shape conflict.** Granting an
  agent reachability to a capability subject cannot be expressed
  without widening its scope beyond declared names (observable: Bar
  1's rig needs a wildcard broader than the capability family, or the
  agent scope must grow subjects unrelated to its declaration) — then
  the plane's subject grammar must be redesigned before anything
  builds on it.
- **No harness accepts a realm door.** Every viable harness insists on
  speaking to its provider directly with its own credential
  (observable: Bar 3 finds no lane — base-URL redirection, proxy
  environment, or native config — through which a real harness
  completes a round trip against the plane) — then the harness arm
  stays design 0007 §4's env-injection permanently and the plane
  serves only native realm callers, a much smaller thing than adopted.
- **Statelessness breaks under real context sizes.** Complete-context
  requests for realistic conversations exceed the wire discipline so
  routinely that per-instance state or chunking-with-stickiness
  becomes the norm rather than the escape hatch (observable: Bar 5's
  protocol cannot complete without instance affinity) — then the
  record-assembles-context division needs rethinking before the plane
  ships.

Media and realtime sessions are named out of scope for this topic — a
later horizon with its own gate; nothing here may foreclose it, and
nothing here builds for it.

## Verdict

Graduated to design 2026-08-28 — all five bars PASS, measured the day
the topic opened, in consumer-position spikes (operator-mode rigs where
the bar demanded scope enforcement; the live harness where the bar
demanded a real one), each suite 3× under `-race`.

- **Bar 1 — the fleet shape holds: PASS** [measured]. Anycast 40
  requests in 4.5ms split 24/16; micro's `$SRV.INFO` the resolve
  surface with model/tags as instance metadata; resolve→filter→pin
  unicast 10/10; an unserved capability answered no-responders. The
  load-bearing find: an agent-scope mint tagged `tool:infer-chat`
  reached the plane through the SHIPPED template — inference
  capabilities are tools, zero scope widening; resolve-and-pin is
  infrastructure's act (measured boundary).
- **Bar 2 — one client loop: PASS** [measured]. All four arms through
  one collect function whose only branch is the grammar's own
  one-shot tell; strictly-consecutive sequences + terminator count
  check (the enemy is silent truncation); `ErrMaxPayload` client-side
  and survivable; headers charged against the payload, subjects not
  (overhead byte-exact); the server's 503 control frame intercepted
  client-side as a typed error in ~1.4ms.
- **Bar 3 — the harness thinks credential-free: PASS** [measured].
  Real `claude -p` announced the env key takes precedence over its
  login, completed door+plane round trips in 2.21s with the plane's
  marker verbatim in its answer; keyless requests 401 with zero plane
  deliveries; no provider material anywhere the harness touched.
- **Bar 4 — names, not routes: PASS** [measured]. Re-pointing a
  virtual name moved traffic alpha→beta (effort default riding along)
  with zero caller change; un-pinned falls back to anycast; the model
  never in a subject, header, or the caller's hands.
- **Bar 5 — the record stays the only record: PASS** [measured].
  Context assembled fresh from the topic each round (1→3→5 turns
  seen); the instance killed and replaced mid-conversation with round
  3 unchanged; census clean — the plane created no stream, KV, or
  bucket.

None of the three reversal readings fired. Outcome: the founding
design of the new component,
[`02-DESIGN/soulstream-inference/0001-the-inference-plane.md`](../../02-DESIGN/soulstream-inference/0001-the-inference-plane.md).
