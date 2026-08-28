# Episode 0147 — The thinking house: submit-and-forget meets a person's hands (2026-08-28)

Spec 014 closes the arc episode 0140 opened two days ago — built by a
parallel agent against designs 0007 and inference-0001, reviewed and
verified before the merge (`soulstream` `79354b9`, eight commits,
2,590 insertions, `specs/014-the-thinking-house`). The whole storyline
the operator asked for on Wednesday is now one walk: **declare an
agent, name its model, submit, walk away** — the house serves it,
the harness thinks through the realm's own door, and the model behind
the name is whatever the catalogue says today.

Two opt-in planes, each filling seams and inventing nothing:

- **`planes.dispatcher`** runs upstream's serve loop and fills exactly
  its two seams. `ConnectAgent` mints each served agent's engine
  credential against the founding's **existing realm role** — design
  0007 §5's open question answered with a key every realm this binary
  ever founded already holds; the TTL is the revocation bound and
  there is deliberately no renewal loop: an expiry ends the engine,
  the placement returns to the race, the next serve mints fresh.
  `EngineFor` gives each agent its own MCP lane, and — when the
  declaration names a virtual model — the door URL and a per-serve
  key through the shipped `Template.Env` seam.
- **`planes.inference`** runs the catalogue (a realm KV bucket, read
  fresh per resolution — watching is the named [O]), the door, and
  the configured instances, with the two custody lines **never
  meeting**: provider credentials resolve from the plane principal's
  own D36 tree at start and exist nowhere the door can reach; door
  keys are issued per serve by the dispatcher plane and exist nowhere
  an instance can read. Issuing for a persona revokes its predecessor;
  stopping the plane revokes everything. Per-WAKE keys are recorded as
  the [O] they are — the mint would need a seam inside the engine's
  admission path that workloads does not offer — rather than faked.
- **The hands**: `soulstream agent submit <declaration.json>`,
  `soulstream model set|ls`, `soulstream provider set`.

The gate (`TestM15ThinkingHouse`, 708 lines, verified independently —
1.66s single run, suites 3× `-race` in 50.5s) walks the composition:
found, up with both planes, catalogue pointed at the stand-in, a
declaration submitted with `inference:{model:...}` and the submitter
gone, a mention answered exactly once through door and plane, the
custody scan clean, a declaration WITHOUT an inference block serving
exactly as before (the ambient lane untouched), and a stop-and-start
of the whole house resuming the serve from the log [measured]. One
defect found and fixed on the way: planes now release what they hold
when a start fails partway (`c803ebb`).

What remains of the storyline, named: the shell's declare surface
(design 0007 §7 — the one human-facing chapter left), per-wake door
keys and catalogue watching (this spec's [O]s), the live provider arms
(`make test-live-anthropic` / `test-live-openai`, the operator's
acts), and the next rc to carry the whole arc to installers.

Reversal condition: none — records a completed build against two
graduated designs; the composition gate stands as the arc's tripwire.

Trail: `soulstream` `79354b9` (spec `67fd2ef`, ceremony `6d79604`,
planes `9c0bae1`/`619f7ff`, CLI `3e740bb`, gate `6daf150`, the
failed-start fix `c803ebb`, models hook `fbe0e82`); designs
[`0007`](../02-DESIGN/soulstream-workloads/0007-agents-as-infrastructure.md)
and
[`inference 0001`](../02-DESIGN/soulstream-inference/0001-the-inference-plane.md);
the arc [episodes 0140](0140-ecosystem-the-focus-agents-as-infrastructure.md)–[0146](0146-inference-the-other-dialect.md).
