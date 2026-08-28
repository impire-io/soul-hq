# Episode 0145 — The first provider: M2's anthropic adapter (2026-08-28)

The mind's second slice landed hours after its first — built by a
parallel agent against design 0001 §4, reviewed and verified before the
merge (`soulstream-inference` `1bc8f55`, branch `002-anthropic-adapter`,
946 insertions): the **anthropic adapter**, the first real provider
behind the `Adapter` seam. One transcript becomes one Messages API
request, the provider's SSE is consumed as it arrives, and every text
delta is emitted the moment it exists — **nothing buffered**, because
output already emitted stands with the caller even when the stream
fails afterwards; the reply grammar's contract, not an optimisation.

The custody line held in code: the credential is construction-time
config — nothing in the package reads an environment variable or a
store; the deployment resolves the key from the instance's own custody
tree and hands it to `New`, which refuses a half-configured provider
outright. Provider-specific error codes (`provider_auth`,
`provider_rate_limited`) live adapter-side, because `wire`'s vocabulary
names conditions the plane itself can reach and stays
provider-agnostic — the instance carries a stream error's code
verbatim, so an adapter's own word travels the wire without `wire`
having to know it. One deliberately right non-default: the HTTP client
carries no timeout of its own — a client-level timeout bounds the whole
body read and would cut a long stream mid-flight; the instance's
`RunTimeout` bounds a generation through the context instead.

Gates [measured, my verification runs]: the hermetic SSE-stub suite
(streaming order, stop mapping, auth/rate-limit refusals with zero
emits, mid-stream disconnect with the emitted chunks standing, params
mapping and system concatenation asserted in the captured requests)
plus the instance-level composition, `-race` ×3; `make check` green on
the merge. Named for the operator's own hands: **the live provider arm
has not run** — `make test-live-anthropic` with a real key (it fails
rather than skips without one; a green run means the provider
answered). Effort's provider mapping (thinking budgets) is a named
[O], dropped silently for now by documented choice.

Reversal condition: none — records a completed build; the live arm's
first run is the standing residue, recorded here so it cannot be
mistaken for measured.

Trail: `soulstream-inference` `1bc8f55` (`114ad43` the adapter,
`0f4c613` the live-target fix); design
[`0001-the-inference-plane.md`](../02-DESIGN/soulstream-inference/0001-the-inference-plane.md)
§4; M1 [episode 0144](0144-inference-the-mind-serves.md).
