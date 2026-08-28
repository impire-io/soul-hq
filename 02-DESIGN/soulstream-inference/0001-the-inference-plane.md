# 0001 — The inference plane: the realm serves its agents' thinking

**Graduated from** research topic `inference-plane`
([episode 0142](../../04-JOURNEY/0142-ecosystem-the-inference-plane.md));
all five bars measured the day the topic opened. This is the founding
design of **soulstream-inference (the mind)** — the component that turns
"which model, whose key, how metered" from per-harness configuration
into a realm service. Evidence classes: `[V]` measured in the
graduation spikes; `[O]` open, decided at spec time or by the operator.

## 1. Principles [V-grounded, load-bearing]

1. **Stateless per request.** Full context travels in every request;
   any instance can serve any request; nothing is sticky. This is the
   single property that preserves queue-group load balancing — and it
   is what lets an instance die mid-conversation with zero effect
   `[V — Bar 5's replacement mid-exchange]`.
2. **One instance, one model.** An instance wraps exactly one model
   behind one adapter and subscribes only to the subjects it can
   serve. Subscription IS the capability declaration; adding a model
   to the realm is starting a process, not editing a map.
3. **The model is an instance attribute, never a request parameter.**
   Callers code against capabilities; the model is discovered about an
   instance (metadata) and pinned by addressing it — it appears in no
   subject and no header `[V — Bar 4]`. *How hard to try* (effort,
   thinking budget) rides request parameters, never a route.
4. **Routing belongs to the substrate and the client.** Anycast is the
   queue group; unicast is the instance-suffixed subject; resolution
   is the micro discovery surface (`$SRV.INFO`) `[V — Bar 1]`. The
   plane contains no router; a smarter router is another client.
5. **An inference capability is a realm tool.** The plane's subjects
   live in the tool space (`SOULSTREAM.SVC.infer-<capability>`), so an
   agent's reachability is the shipped capability-minting machinery —
   the declaration names the tool, the mint narrows to it, zero scope
   change anywhere `[V — Bar 1's agent-scope arm]`.
6. **The record stays the only record.** The plane holds no
   conversation state — no stream, no KV, no bucket `[V — Bar 5's
   census]`. Context is assembled from the topic's materialisation by
   the caller; what deserves keeping flows back to the record as
   ordinary ops. This is the substrate boundary applied to thinking.
7. **Harnesses think credential-free.** Provider credentials live in
   the serving instance's own custody (its D36 secret tree); a harness
   receives only a realm key and a door URL `[V — Bar 3, real
   harness]`. Metering rides headers middleware can read without
   touching a body.

## 2. Subjects [V]

```
SOULSTREAM.SVC.infer-<capability>              anycast — the queue group decides
SOULSTREAM.SVC.infer-<capability>.<instance>   unicast — a queue group of one
```

First capability: `chat` (text/multimodal in, text out). Later
capabilities (`embed`, `transcribe`, `speak`, …) are new tool names
added when an instance serves them — an unserved capability answers
no-responders, the routing layer telling the truth for free `[V]`.

- The tool space keeps inference off the record's streams (the M1.2
  lesson: transient RPC never rides `SOULSTREAM.>`) and inside the
  agent scope's `{{tag(tool)}}` template.
- **Agents anycast; infrastructure pins** `[V — measured boundary]`:
  the exact-token tag template does not cover instance suffixes or
  `$SRV`, so resolve-and-pin belongs to the door, the dispatcher, and
  operator tools. This is by design — pinning is the catalogue's job
  (§5), and the catalogue is infrastructure.
- A field is promoted into the subject only when a party other than
  the handler must act on it without reading the body (broker
  authorization, a dedicated pool, selective subscription, stream
  segregation). Capability qualifies; model, effort, and format do
  not — they are metadata, parameters, and headers respectively.

## 3. The reply grammar [V]

Core NATS request/reply; the caller subscribes its own inbox (a plain
request would truncate a multi-frame reply `[V]`).

- **Streaming**: content frames — non-empty payloads with a
  **strictly consecutive** `Infer-Seq` from 1 — terminated by exactly
  one **empty-payload sentinel** carrying `Infer-Status: done|error`,
  the last content sequence (the gap check), `Infer-Stop-Reason`, and
  usage headers. Meaning lives in headers; emptiness selects the
  control plane — the substrate's own idiom.
- **One-shot**: a single non-empty reply with usage headers and no
  sentinel — correlation is the completion signal. A one-shot result
  too large for the wire **streams regardless**; the first frame's
  sequence tells the caller which case it got.
- **One client loop.** All four shapes (stream, one-shot, oversized
  one-shot, mid-stream error) are consumed by one loop whose only
  branch is the grammar's own tell: first message, content-bearing,
  unnumbered → one-shot `[V — the request's stream flag never
  reaches the loop]`.
- **Errors are terminal — and metered.** One error sentinel (status +
  error code + the usage seen before the break), then silence; partial
  output already delivered stands, marked incomplete `[V]`. Standing
  output rides the error terminator's usage header, because unmetered
  partials would undercount exactly the runs that need accounting most
  (the M2 finding, built). In the Go client the code travels as a
  typed error beside the partial result (the decided contract).
- **Invariants refuse as protocol errors**, not conventions `[V]`: a
  content frame never carries status; every empty frame carries a
  recognized discriminator (`Infer-Status` or `Infer-Progress`); a
  sequence jump refuses at the frame; the enemy is silent truncation.
- **No-responders needs no vocabulary**: the server's own 503 control
  frame surfaces client-side as a typed error in ~1.4ms without
  entering the loop `[V]` — "no instance" is distinguishable from
  "slow instance" for free.
- **Wire-size discipline** `[V]`: `max_payload` charges body + the
  serialized header block (byte-exact) and not the subject; the limit
  is client-side and survivable, which is what makes the streaming
  fallback implementable; a frame's threshold is computed from its
  actual headers (the budget shrinks a byte when the sequence gains a
  digit). Chunked *input* for oversized requests is a named later
  gate, not first-slice scope (§9).

## 4. Instances and adapters

An instance = one adapter + one model + one custody tree.

- Built on the micro framework: service name `infer`, instance
  metadata carrying `model`, `tags`, `formats` — the resolve surface
  is micro's own discovery, no bespoke endpoint `[V]`.
- The **adapter** is the seam: `Adapter` turns one request (context +
  params) into provider calls and frames. First adapters: the
  **stand-in** (deterministic, for every hermetic gate) and
  **anthropic** (the first real provider; its live test env-gated).
- The instance's provider credential comes from **its own D36 secret
  tree** (`providers/<name>` convention) — reachable by no canonical
  scope, structurally `[V — the 0141 custody probes]`.
- Effort/thinking/params map to provider request properties inside
  the adapter; the plane never interprets them.
- **Adapter duties the compiler cannot enforce** (M2's findings, now
  in the seam's contract): a provider stream ending without its own
  terminal marker is refused as an error, never returned as a clean
  result — silent truncation would otherwise enter through the one
  door the grammar cannot watch; the partial result beside an error
  meters the output that stands; the credential is construction-time
  only. Named [O]s from the first real adapter: **effort naming** —
  the provider's home is an effort *level* vocabulary (thinking
  budgets are gone on current models), so the decision is which plane
  word means which provider level, not a translation; **participant
  attribution** — the plane's transcript is strictly richer than the
  provider request shape, and writing the speaker into prompt text
  would cross the attribution-is-data line, so participants stop at
  the provider boundary until decided otherwise; and a recorded
  hazard: current models refuse a non-default `temperature`, which the
  adapter forwards anyway because the parameter belongs to the
  caller's model choice.

## 5. The catalogue: names, not routes [V mechanism, O home]

A **virtual model name** resolves infrastructure-side to a route
descriptor `{capability, model_pin | tag_predicate, default_params}`:
anycast when unpinned, resolve-then-unicast when pinned. Re-pointing
the name moves traffic — and default params — with zero caller change
`[V — Bar 4]`. The concrete model stays an instance attribute; a name
nobody serves answers no-responders, truthfully.

The catalogue's **home** is an [O]: a realm KV bucket the resolvers
watch is the natural shape (the product provisions it; the shell reads
it; the dispatcher and door resolve through it) — decided at the
product-wiring spec, not here. The resolver is a client library
(`client/resolver`), never a worker.

## 6. The door: harnesses think through the realm [V]

An HTTP front speaking the Messages API shape harnesses already speak:
health endpoint, key auth, one translation to a plane request, replies
returned as SSE (streaming) or JSON (one-shot). Measured with the real
harness: base-URL + env-key redirection is the lane, and the key takes
precedence over the harness's own login in its own words `[V — Bar 3,
2.21s round trips]`.

- The door custodies nothing and holds no provider credential; a
  keyless or wrong-key request dies at the door with zero plane
  deliveries `[V]`.
- **Per-run keys**: the dispatcher mints a short-lived door key per
  wake and injects `{BASE_URL, KEY}` through the shipped
  `Template.Env` seam — design 0007 §4's mechanics with the door as
  the target [O: the key mint rides the identity plane's existing
  machinery; its exact lane is the door spec's first decision].
- The door is the **metering surface**: usage headers aggregate here
  per key/persona [O: where metering lands — the record? — decided
  with the product wiring].

## 7. What the plane refuses to contain

- **No orchestration.** One request, one generation; the plane never
  triggers itself. Multi-agent coordination lives in workloads
  (the dispatcher, the wake machinery) consuming the record.
- **No router.** §1.4. Fallback chains, cost-aware selection: clients.
- **No conversation store.** §1.6. Incremental/server-held context is
  refused as a feature: the record materialises, the caller assembles.
- **No realtime media.** Live audio/video sessions are a later
  horizon with their own research gate; nothing here forecloses them,
  nothing here builds for them.

## 8. Acceptance criteria (the graduation bars as standing tests)

1. Two instances, one capability: anycast serves from both; resolve
   returns metadata; pin reaches exactly the pinned instance; an
   unserved capability answers no-responders; an agent-scope mint
   with the tool tag reaches anycast, a tag-less one is server-denied.
2. The four reply shapes through one client loop; invariants refuse
   as protocol errors; the gap check bites; error is terminal with
   partial output standing.
3. A real harness completes a round trip through door + plane on a
   realm key alone; keyless dies at the door; no provider material
   anywhere the harness touches.
4. Re-pointing a virtual name moves traffic and defaults with zero
   caller change; un-pinned anycasts.
5. Census: the plane creates no stream, KV, or bucket; an instance
   replaced mid-conversation changes nothing; context arrives
   complete from the record each request.

## 9. Open, named [O]

- The catalogue's home (realm KV, watched) — product-wiring spec (§5).
- The door-key mint lane and TTL; where metering lands — door spec (§6).
- Chunked input for oversized requests (unicast, ack-paced) — gated on
  a real conversation hitting the ceiling, with the record-reference
  escape hatch (`soulstream://` artifacts) considered first.
- Additional capabilities (`embed`, `transcribe`, `speak`) and
  provider adapters — each a process to start, by demand.
- Realtime media — its own research gate, later.
- ~~The dispatcher's `inference` block schema~~ — landed in the
  declaration (workloads `67a75e2`): a virtual name, agent-only,
  credentials refused by construction; the product wiring closes it
  against §5's catalogue.
- A wedged provider inside `RunTimeout`: the instance emits no
  progress frames today, so a stalled generation is silent until the
  timeout. A per-chunk idle deadline in adapters, or progress frames
  from the instance's stream path, closes it — by demand, with the
  first long-job capability.
