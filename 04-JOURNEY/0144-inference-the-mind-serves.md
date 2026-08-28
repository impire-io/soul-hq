# Episode 0144 — The mind serves: M1 of the inference plane (2026-08-28)

The ninth component's first slice landed the day the component was
founded — built in parallel with the dispatcher (episode 0143) and
published at
[impire-io/soulstream-inference](https://github.com/impire-io/soulstream-inference)
(founding commit `8f4f4a6`, 24 files). The realm can now serve its
agents' thinking: instances, the grammar, the client, the catalogue
descriptor, and the door — design 0001 realized package for package,
with the graduation bars standing as the repo's own tests.

What shipped, and where the design's words became law:

- **`wire`** — the grammar with its invariants enforced at
  construction AND refused at reading (a terminator cannot carry
  content because the builder has no payload parameter to give it);
  strictly consecutive sequences with the terminator's count as second
  check; the per-frame payload budget computed per frame (the
  digit-shrink fact from the graduation, now `MaxFrameBody`); the
  transcript codec, one source for client and instance.
- **`client`** — the one `Collect` loop; `Resolve` over micro's own
  discovery; the catalogue `Descriptor` routing names to anycast or
  resolve-and-pin with default params merging under the request's.
- **`instance`** — the `Adapter` seam (one adapter, one model, the
  pairing is the instance); anycast + unicast endpoints; the caller's
  wish overruled by an exact fit check (`wire.Fits`, the serialized
  header block charged byte-for-byte); credentials deliberately
  absent — nothing in the package reads a secret.
- **`adapter/standin`** — the deterministic model every hermetic gate
  runs against, reporting context size and effort so tests prove the
  request traveled whole.
- **`door`** — health, key auth, one translation, SSE + one-shot;
  custodies nothing; a keyless request dies with zero plane
  deliveries.

The gates [measured]: the five §8 acceptance criteria green hermetic
and 3× `-race` — including the agent-scope arm on a real operator-mode
server (the scope template carried as a test fixture with the product
e2e named the drift court, workloads' dual-write pattern) and the
statelessness arm (context 1→3→5 across a mid-conversation instance
replacement, census zero streams). The live arm
(`make test-live`, behind a build tag so nothing skips): **the real
harness thought through the M1 door in 2.40s** — and caught the one
bug the research spike missed: the harness sends `system` as content
blocks, not a string; the door reads both now. Two small learnings on
the way: a scoped-signer user JWT must be explicitly marked scoped
(the identity plane's mint does it internally; a direct fixture must
too), and CI on the founding push is the repo's birth certificate.

Not in M1, named: the `anthropic` adapter (M2 — the first real
provider behind the same seam, live test env-gated), the catalogue's
durable home and per-wake door keys (M3, the product wiring), the
release pipeline (goreleaser + tag, at the first consumer's demand —
the archivist precedent).

Reversal condition: none — records a completed build against the
graduated design; design 0001's principles carry their own reversal
observables, and the census/statelessness tests are their standing
tripwires.

Trail: [impire-io/soulstream-inference](https://github.com/impire-io/soulstream-inference)
`8f4f4a6`; design
[`0001-the-inference-plane.md`](../02-DESIGN/soulstream-inference/0001-the-inference-plane.md);
graduation [episode 0142](0142-ecosystem-the-inference-plane.md); the
parallel build of the same day [episode
0143](0143-workloads-the-dispatcher-builds.md).
