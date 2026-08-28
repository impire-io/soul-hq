# Episode 0146 — The other dialect: OpenAI compatibility, both directions (2026-08-28)

At the operator's direction the plane learned the second wire dialect
the ecosystem actually speaks — built by a parallel agent, reviewed and
merged the same hour (`soulstream-inference` `fc7bdcf`, branch
`003-openai-compat`, ~1,900 insertions). The point is not one provider:
**Chat Completions is the dialect OpenAI, Ollama, vLLM, LM Studio and
llama.cpp all answer**, so the adapter makes every local runtime a
plane instance, and the door surface makes every OpenAI-speaking
harness a realm thinker.

- **`adapter/openai`** — the same discipline as the first adapter
  (SSE consumed unbuffered, partial output standing and metered,
  provider errors wrapped so the diagnosis survives, 400 telling the
  truth) with the dialect's own facts held honestly:
  keyless-with-explicit-BaseURL is legal because a local runtime
  authenticates nobody, keyless-default-URL is refused at construction
  because it can only fail, and `[DONE]` carries the whole terminal
  weight — this dialect has no other terminal event, so its absence
  refuses as truncation (the adapter duty, exercised).
- **The door speaks both shapes over one plane connection** — factored
  so the two surfaces share everything except dialect: one `think`
  middle (encode, route, collect), one `Authorize`, each surface
  reading the key from the header its own clients send (`X-Api-Key`
  for Messages, `Authorization: Bearer` for Chat Completions) —
  *which header carries the key is dialect, not policy*. `GET
  /v1/models` arrives behind an optional hook the catalogue wiring
  fills, with the recorded rule that it must advertise the same names
  `Route` resolves.
- **`internal/providercode`** — the shared-constants call: provider
  error codes (`provider_auth`, `provider_rate_limited`) live in one
  internal package so the string a caller matches on cannot drift
  between adapters, while each adapter's public constants alias it
  and `wire`'s vocabulary stays provider-agnostic.

Gates [measured, my verification]: full `make check` green on the
merge; adapter, door, and instance suites 3× `-race` (openai 2.3s,
door 3.7s, instance 5.2s); `-tags live` vet clean. The live arms
(`make test-live-openai` with a real key — or an Ollama at an explicit
BaseURL, which needs no key at all) are the operator's pending acts,
alongside the anthropic one.

Reversal condition: none — records a completed build; the dialect's
truncation semantics and keyless rules are pinned by standing tests.

Trail: `soulstream-inference` `fc7bdcf` (`eca1e23` the adapter,
`75ecb90` the door surface); design
[`0001-the-inference-plane.md`](../02-DESIGN/soulstream-inference/0001-the-inference-plane.md)
§4/§6; the first adapter [episode
0145](0145-inference-the-first-provider.md).
