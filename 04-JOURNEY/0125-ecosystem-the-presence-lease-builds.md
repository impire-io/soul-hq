# Episode 0125 — The presence lease builds: the lamp lights (2026-08-24)

Episode 0124's decisions built the same day, both halves. **soulstream-core
v0.13.0** ships the `presence` package in the tool catalog's exact
shape — additive entry with an extras map, bucket on first write,
absence never an error — plus the ground the catalog never needed: a
`State` carries the KV entry's own timestamp as the renewal moment,
`Read(now)` derives *present / left / last-seen* as the reader's
judgment, and `Hold` is the whole lease in one call (write, renew on
the cadence, farewell on a fresh short context once ctx ends). One
mechanism choice worth its comment: renewal is a plain put, not
compare-and-set — the key's owner is its only legitimate writer, and
after a crash-and-restart a stale revision must never make the fresh
writer lose [mechanism-argument]. Measured in the package's own rig:
the lease-honesty triple (renew advances the timestamp with the words
untouched; cancel reads *left* forever; silence reads *last seen* past
the horizon), unknown fields and unfamiliar statuses surviving the
round trip, and the op-log census unchanged across a full lifecycle —
acceptance criteria 1–3 of the design, each a standing test
[measured]. A docs page ships with it (the lamp in the window),
closing the Article III gap the toolcatalog had left open.

**soulstream branch `011-presence-lease`** wires the wrap — spec-kit
flow by the operator's choice, the first spec folder since 010. Before
answering mentions the wrap ensures a directory floor and lights its
lamp; after the run loop returns it waits for the farewell before the
connection closes. Two constraints shaped the wiring: the wrap lane
holds **no signer**, so the profile floor is name + created-at with no
signing key and no pretence [measured, code trace]; and
`registry.Publish` **replaces** display metadata on an existing entry,
so the floor is lookup-first — publishing only into absence, never
speaking over the richer profile an agent's own harness published.
Everything is advisory: a lease failure is a log line, never a refusal
to answer — courtesy, never correctness, held in the wiring. The live
rig runs the whole story through real admission (founded node,
sentinel + token — the persona scope exactly as minted), which
confirms by measurement what 0124 had from a code trace: the scope's
`$KV.>`/`$JS.API.>` tails let a wrap create and write the presence
bucket with no ceremony change [measured, 0.3s in `make test`].

What remains is exactly what 0124 named: the shell 0008 build (the
first-steps card and the arrival line now have a face to read), the
live run on a standing deployment (the quickstart's pending human
act — install a wrap at the next tag, watch the lamp before and after
Ctrl-C), and the cadence numbers held against the dogfood's word.

Reversal condition: none — records a completed build/measurement; the
direction decisions and their reversal readings are episode 0124's.

Trail: soulstream-core `presence/` + `docs/presence.md`, commit
`5beea83`, tag **v0.13.0**; soulstream
`specs/011-presence-lease/` + `cmd/soulstream/wraplife.go` +
`wraplife_test.go` + the `cmdWrap` wiring, branch `011-presence-lease`
commit `72e2325` (pushed, awaiting merge); designs
[`extensions/presence.md`](../02-DESIGN/soulstream-core/extensions/presence.md)
and
[`0008-the-first-hour.md`](../02-DESIGN/soulstream-shell/0008-the-first-hour.md)
§4.
