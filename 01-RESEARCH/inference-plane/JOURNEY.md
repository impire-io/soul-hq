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
