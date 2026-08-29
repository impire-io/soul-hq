# Episode 0153 — A runtime of your own serves the realm: the openai arm in the house, and byon thinks (2026-08-29)

The operator asked for the provider key to be set up on byon and the
answer became a pivot: there is no Anthropic key — the provider is
**lemonade-server on beno4**, a local runtime speaking the OpenAI
dialect with GLM-4.7-Flash loaded. That is the demand episode
[0146](0146-inference-the-other-dialect.md) built the adapter for
("local runtimes become plane instances"), arrived one hop early: the
adapter existed, but the **house** wired only `standin` and
`anthropic` — no openai arm, no way to say where a runtime of one's
own lives.

The house grew the arm the same hour (`soulstream` `8cb48f9`, released
as **v0.14.0-rc.6**): `planes.inference.instances` gains `base_url`,
and the openai case honors the adapter's own contract — keyless with
an explicit base URL is legal (a local runtime authenticates nobody),
the public API takes a key from the plane's own custody. Verify
refuses the shapes the house cannot honour by name: neither key nor
base URL, and a base URL on an adapter that does not take one; the
keyless shape round-trips `config.json` with its URL intact [measured,
the refusal table + the positive arm]. No new custody surface exists:
`base_url` is an address, never a credential — byon now runs a
thinking plane with **no provider secret anywhere**, honestly, because
the runtime is the deployment's own.

The deployment, measured end to end on byon [all measured]: lemonade
answers a direct completion on beno4 in ~130ms of prediction; from
beno1 its surface answers 200 across the LAN; with `planes.inference`
declared (one keyless instance, `GLM-4.7-Flash-GGUF`, backups beside
as the pattern demands) the journal says `thinking (agents)
http://127.0.0.1:8600`, the door refuses keyless with 401 on both
dialects, the `$SRV.INFO.infer` scatter answers in **470µs** with the
instance's model and both subjects — and one real generation rode the
plane's own anycast: a NATS request on `SOULSTREAM.SVC.infer-chat`
answered `served` with the grammar's usage headers riding
(`Infer-Usage-Output-Tokens: 165`, `Infer-Model: GLM-4.7-Flash-GGUF`).
beno1's realm thinks on beno4's GPU through the record's own
substrate. The catalogue holds byon's first name — `glm-flash`,
pinned — so the Models sheet's empty state is retired on the realm
that recorded it, and a declaration may now carry
`inference:{model:"glm-flash"}`.

Two small operational truths of record: the tailnet's SSH policy
refuses this operator's key to beno1 while the LAN admits it (the
route is the LAN's), and lemonade's OpenAI surface lives on `:13305`
(`:9000` speaks something else; `:8080` on beno4 is another service
entirely) — found by probing, not assumed. What stands unchanged:
`claude` is still not installed on beno1, so a declared agent is
claimed but its engine unserved; the per-serve door key remains the
dispatcher's act, which is why the plane's proof rode the substrate
rather than the door.

Reversal condition: none — records a completed build, release, and
deployment with their measurements; the verify refusal table and the
keyless round-trip test are the arm's standing tripwires. The
keyless-instance call reverses if a runtime the deployment does NOT
own is ever pointed at by `base_url` — that reading makes the missing
key a custody hole, and the openai secret lane (already built) becomes
the demanded default.

Trail: `soulstream` `8cb48f9`, tag `v0.14.0-rc.6` (release verified:
four tarballs + checksums, the tap at `0.14.0-rc.6`, the beno1 install
checksum-verified and printing its version);
[`DOGFOOD.md`](../03-IMPLEMENTATION/DOGFOOD.md) carries the byon
record; the adapter [episode 0146](0146-inference-the-other-dialect.md);
the plane's founding
[design 0001](../02-DESIGN/soulstream-inference/0001-the-inference-plane.md);
the sheet this lights up
[episode 0151](0151-shell-the-models-surface.md).
