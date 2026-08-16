# Episode 0093 — Pre-v1 renames are clean breaks (2026-08-16)

A next-day reversal, at the operator's direction: "I don't want
backwards compatibility in pre-v1." Episode 0092's rename had shipped
with read-forever alias keys, a creds-filename fallback, a state-subdir
fallback, and accepted legacy flag spellings — compatibility machinery
for a product whose whole audience predates v1. All of it is gone
(soulstream `dd32cbd`, net −41 lines): one schema, one code path.

What replaced the fallbacks is not silence but a **named refusal**: the
byname-era keys are still *detected* (two `json.RawMessage` fields that
are never read), so a realm founded under them is refused with the
hand-migration spelled out — rename the keys (`door→mcp`,
`fold→signin`), `mv users/fold.creds users/signin.creds`, `mv fold/
signin/` — or re-init. Refusing by name is not compatibility; silently
dropping a sign-in plane because its key spelled the old way would
have been the worst of both worlds [judgment, and the operator's
standing honesty rule]. Pinned both ways [measured]: the ceremony test
asserts the refusal names the migration, and a live tamper of a fresh
realm's config produced the exact sentence. `SOULNODE_STATE` became
`SOULSTREAM_STATE` in the same pass, no alias — same class of
leftover.

Design 0001 §2 now states the rule generally: **no compatibility shims
before v1** — renames are clean breaks, honest refusals over silent
misreads. Episode 0092's reversal-condition (deprecation notes on the
dual keys) is moot with the keys gone.

Reversal condition: v1 itself — the first tagged v1.x makes every
config key, flag, and on-disk name a compatibility surface, and this
rule inverts.

Trail: design
[`0001-soulnode-composition.md` §2](../02-DESIGN/soulstream/0001-soulnode-composition.md);
soulstream `dd32cbd`; the fallbacks it removes landed in
`specs/009-one-console-one-vocabulary` (episode 0092).
