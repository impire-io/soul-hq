# inference-plane — investigation journal (started 2026-08-28)

## 2026-08-28 — Bar 1 measured: PASS — and inference capabilities ARE tools

**Hypothesis:** the fleet shape (stateless single-model instances,
capability subjects, queue-group anycast, discovery-resolved unicast)
composes with our substrate's existing scope machinery without widening
anything.

**Rig:** the standing operator-mode rig from the 0141 spikes (real
operator/SYS/APP realm, identity plane through `embed.Run`, both
canonical scoped signers). Two stand-in instances built on the NATS
micro framework, each carrying `model`/`tags`/`formats` as instance
metadata. **Working namespace: the capability is a tool** —
`SOULSTREAM.SVC.infer-chat` — chosen to test the boldest composition
first.

**Measured, first run + 3 consecutive `-race` runs [measured]:**

- **anycast**: 40 requests in 4.5ms, split alpha=24/beta=16 — the queue
  group is the router, both instances serve, nothing else needed;
- **resolve**: micro's own `$SRV.INFO` scatter returned both instances
  with model/tags metadata; the client filtered `model:beta` and got
  the exact instance id — discovery IS the resolve surface, no
  bespoke endpoint;
- **unicast**: 10/10 requests to the pinned suffix answered by exactly
  the pinned instance (a queue group of one);
- **truth for free**: an unserved capability (`infer-embed`) answered
  no-responders — the routing layer reporting the fleet honestly;
- **the scope arm, the load-bearing find**: an agent-scope mint tagged
  `tool:infer-chat` reached the plane's anycast through the SHIPPED
  agent template — **zero scope widening; inference capabilities are
  tools under the existing `{{tag(tool)}}` machinery** — and a
  tag-less mint died at the server. The reversal condition's first
  reading (scope enforcement vs fleet shape) did not fire; it
  inverted: the fleet shape lands on the scope machinery we already
  have.

**Findings for the design, recorded not asserted:** under today's
exact-token tag template an agent can neither unicast an instance
suffix nor reach `$SRV` — **resolve-and-pin is an infrastructure act**;
the agent gets anycast, and pinning (the virtual-model catalogue's
job) happens node-side. That division is probably right — the
catalogue was going to live infrastructure-side anyway (Bar 4) — but
it is now a measured boundary, not a preference.

## 2026-08-28 — Bar 3 measured: PASS — the real harness thinks through the door

**The probe first, because it decides the reversal condition:** the
real production harness pointed at a local base URL with an env API key
announced, in its own words, that the env key *takes precedence over
the interactive login*, sent a health check (`HEAD /api/hello`) and a
genuine Messages-API request carrying the realm key in `X-Api-Key`, and
consumed our response. The "no harness accepts a realm door" reversal
reading is dead on the strongest harness there is [measured].

**Rig:** a minimal realm door — Messages-API-compatible HTTP front
(health check, key auth, request translated to a plane anycast, the
answer returned as SSE stream or one-shot JSON per the request's own
`stream` flag) — over one stand-in instance whose provider credential
exists only instance-side. The whole authentication world handed to
the harness: the door URL and a realm-issued key, injected exactly as
the dispatcher would (the shipped `Template.Env` seam).

**Measured, first run + 3 consecutive `-race` runs, live harness each
time [measured]:**

- the real `claude -p` completed a full round trip through door +
  plane in **2.21s**: one door call, one plane round trip, the plane's
  marker text verbatim in the harness's printed answer — the thinking
  demonstrably came through the realm;
- **refusal arm**: a keyless request answered 401 at the door with
  zero plane deliveries;
- **custody**: no provider material in the harness's environment
  (constructed, checked), none in its output, realm key ≠ provider
  key — the provider credential never left the instance.

**What the build takes from this:** the door is small — a health
endpoint, key auth, one translation, one SSE shape — and the harness
needs nothing bespoke: base-URL + key is the lane every API-compatible
harness already speaks. Per-run keys minted at wake time (instead of
the spike's static string) are the obvious hardening, and the door is
where metering headers will surface.

## 2026-08-28 — Bars 4+5 measured: PASS, same afternoon

**Bar 4 — names, not routes** [measured, first run + 3× `-race`]: a
virtual model name resolved node-side through a catalogue descriptor
(`{model_pin, default_params}`) into resolve-and-pin or anycast.
Re-pointing the name — alpha to beta, and the effort default with it —
moved the traffic with **zero caller change**: same name, same request
body, same capability subject. Un-pinning the name fell back to
anycast with both instances serving. The model appeared in no subject,
no header, and never in the caller's hands; effort rode the request
`params`. Left for the design: where the catalogue LIVES (a realm KV
entry the node watches is the natural shape) — the spike proved the
resolve-time mechanism, deliberately not the storage.

**Bar 5 — the record stays the only record** [measured, first run +
3× `-race`]: three conversation rounds where each request's context
was assembled fresh from the topic's materialisation (the stand-in
instance reports how many turns it saw: 1, then 3, then 5 — the
context demonstrably traveled complete). The serving instance was
killed and replaced mid-conversation; round 3 completed unchanged
with the full context — nothing the plane held mattered because it
held nothing. The census: every stream on the server is the realm's
own provisioning; the plane created no stream, no KV, no bucket. The
third reversal reading (statelessness breaking under real context)
did not fire at spike scale; the design must still carry the wire
discipline honestly for long conversations (chunking/reference
escape hatches — a design section, not a blocker).

## 2026-08-28 — Bar 2 measured: PASS (parallel spike, independently verified)

Run as a parallel spike in its own consumer-position module
(`infplane-grammar`, nine tests), then verified independently — the
suite re-run and the client read line by line from this side before
recording [measured, my run: 9/9 green, `-race` ×3 green].

**The one-client-loop claim holds honestly.** All four arms — streamed
(5 frames + a mid-stream progress frame + terminator, 1.24ms),
one-shot (single reply, 429µs, provably no terminator on the wire),
oversized one-shot (120 bytes arriving as 2 frames though the request
asked `stream=false`, 343µs), and mid-stream error (3 frames kept,
stream abandoned, partial output standing marked incomplete) — went
through one `collect` function whose only branch is the grammar's own
one-shot tell: *first message, content-bearing, unnumbered*. That is
the grammar deciding, not the caller special-casing [measured].

**Invariants pinned as protocol errors, not conventions:** a content
frame carrying status refuses; an unnumbered content frame mid-stream
refuses; a sequence jump refuses; the terminator's sequence is a gap
check against frames actually seen (inbox delivery is at-most-once —
the check caught a deliberately dropped frame in the suite).

**Two mechanism findings beyond the brief:**

- **No-responders never enters the loop.** The server answers an
  unserved subject with its own empty-payload 503 control frame — the
  same idiom this grammar uses — and the client library intercepts it
  at the subscription, surfacing a typed error in ~1.4ms. "No
  instance" arrives as itself, immediately, no timeout burned, and
  the grammar reserves no status code for it [measured].
- **The payload budget charges headers, not subjects.** At
  `max_payload=1024`, the largest deliverable content frame body was
  998 bytes — 26 bytes are the serialized headers, charged against
  the payload limit; the subject is not [measured]. The design's
  wire-size section inherits this as fact.

**All five bars now PASS.** Ready to graduate; the design's home
(component/repo) is the operator decision pending.

### Bar 2 addendum — the spike's full report, facts the design needs

- **The one-path evidence, sharpened:** the request's `stream` flag
  never reaches the reply loop (grep-verified: two occurrences, both
  request-side); arms 2 and 3 pass **literally identical caller
  arguments** and differ only in the instance's result size — the
  caller cannot tell one-shot from streamed and does not need to
  [measured]. Wire counts asserted from a third connection.
- **`max_payload` recovery is client-side and survivable:** a
  too-large publish returns `ErrMaxPayload` before the wire and the
  connection stays usable — the streaming fallback is implementable
  precisely because the limit does not kill the connection
  [measured]. Overhead equals the serialized header block byte for
  byte (asserted as an equality, confirmed in the client library's
  own size computation); subject and reply are not charged.
- **The per-frame threshold consequence:** an instance's streaming
  threshold is `max_payload` minus its own serialized headers, and
  the budget shrinks by one byte when the sequence number gains a
  digit — a hardcoded threshold fails on the last frames of a long
  stream. Compute per frame from the actual header block.
- **Grammar decisions the design must take deliberately** (found by
  building, pre-decided nowhere):
  1. where a terminal error's code lives — the spike kept the result
     struct clean and returned a typed error beside the partial
     output; the design picks the contract;
  2. **strictly consecutive, not merely monotonic** sequences — the
     failure mode being designed against is silent truncation (short
     text marked complete, the worst outcome this grammar can
     produce); the spike checks consecutiveness at each frame AND the
     terminator's count;
  3. **every empty frame must carry a recognized discriminator** — an
     empty frame with neither status nor progress is unroutable and
     refuses as a protocol error.
- **Fan-out is ordered per subscription, not simultaneous across
  subscriptions** — an observer connection receives its copy after
  the requester has already returned. Bit the spike once as a test
  flake; every future rig watching the wire from a second connection
  waits for counts then requires a quiet window.
