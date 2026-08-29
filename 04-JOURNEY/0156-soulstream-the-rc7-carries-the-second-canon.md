# Episode 0156 — The rc.7 carries the second canon: held, merged, pinned (2026-08-29)

The operator looked at the ten screens side by side and called the
canon held — design 0011 §9.4, the one gate the build waited behind —
and the publish step followed in one motion. The `second-canon`
branches fast-forwarded onto both mains; the shell tagged
**v0.11.0-rc.6** and the fold **v0.8.2** (signed tags carrying the
decision's words); the house moved its pins onto both and cut
**v0.14.0-rc.7** (`f4f7c3f`), the full house gate green on the new
pins [measured]. The pins were gated before any tag existed on the
remote: go's fetches for the two modules were routed to the sibling
clones through a scoped git config (an `insteadOf` rewrite in a
throwaway `GIT_CONFIG_GLOBAL`, `GOPRIVATE` for the pair) — no
`replace`, no `go.work`, nothing in any repo; the same tags the
operator pushes are the ones the gate measured, so the recorded
`go.sum` hashes hold when the world fetches them from GitHub
[mechanism-argument: identical tag, identical module hash].

With rc.7, every human-facing surface of the product — the console's
ten screens and the fold's sign-in, enrolment, and admin pages —
wears the Impire canon, and the words on all of them are byte-for-byte
what rc.6 said. The three named [O]s stand: the mark (a naming act),
the dark theme (a small follow-up on the kit's own tokens), and the
vocabulary horizon (riding the foundation architecture's research,
not this arc).

Reversal condition: none — records an acceptance and its publish
step; the decision's reversal conditions live in episode 0154 (the
one-accent reading reopens if the operator can no longer tell at a
glance which voices somebody else answers for).

Trail: design
[`0011`](../02-DESIGN/soulstream-shell/0011-the-second-canon.md)
(§9.4 paid, status updated in place); soulstream-shell `v0.11.0-rc.6`
(= `3a1ca5d`), soulstream-idp `v0.8.2` (= `32855fd`), soulstream
`f4f7c3f` tagged `v0.14.0-rc.7`; the byon binary bump onto rc.7 is
the operator's dogfood act, tracked in
[`DOGFOOD.md`](../03-IMPLEMENTATION/DOGFOOD.md).
