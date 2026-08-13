# Episode 0070 — The rename sweep: eight repos, one evening (2026-08-13)

Episode [0069](0069-ecosystem-one-name-soulstream.md)'s decision was
executed the same evening: every repository, module path, binary, and
the sanctioned wire vocabulary now carries the one brand. Eight repos
moved in dependency order, each gate green before its push, every tag
release green in CI `[measured]`:

| Repository | Tag | What moved |
|---|---|---|
| soulstream-core (ex-soulstream) | v0.8.0 | module path; the nested `node/` module extracted out; plugin URLs follow |
| soulstream-identity (ex-soulidentity) | v0.2.0 | module path, binary; **wire segment `soulidentity` → `identity`** |
| soulstream-workloads (ex-soulrealm) | v0.2.0 | module path, binary; **`SOULREALM.*`/`SOULREALM_*` → `SOULSTREAM.*`/`SOULSTREAM_*`** |
| soulstream-archivist | v0.3.0 | re-pinned to core (name was already the pattern) |
| soulstream-idp (ex-soulfold) | v0.4.0 | module path, binary, store prefix `soulstream_idp_`; passkey prompt says "Soulstream" |
| soulstream-mcp (founded) | v0.1.0 | extracted from `soulstream/node v0.7.0`; own CI/release; adapter-position cycle guard kept, core-side halves now structural |
| soulstream-shell (ex-soulhelm) | v0.2.0 | module path, `shellserver` package; e2e interim-pinned, then re-proven (below) |
| **soulstream** (ex-soulnode) | **v0.8.0** | the product takes the bare name; all seven pins at new paths; `planes.helm` → `planes.shell`; binary `soulstream` |

The traps the sweep hit, recorded for the next one:

- **The checksum DB burns versions into a path, not a repo.** The
  product reuses `github.com/impire-io/soulstream`, whose v0.1.0–v0.7.0
  are immutably recorded from the record library — so the product's
  first tag is **v0.8.0**, skipping v0.4–v0.7 of its own lineage
  `[mechanism-argument]`.
- **The circular pin between shell and product** (the shell's e2e boots
  the product; the product pins the shell) resolved in the recorded
  order: the shell tagged v0.2.0 with its e2e pinned wholesale to the
  pre-rename published composition (stated in its go.mod), the product
  tagged v0.8.0 pinning shell v0.2.0, and the shell's ceremony re-proof
  landed as the next commit — the whole human ceremony passes against
  the renamed constellation in ~3.9 s `[measured]`.
- **The wire went with the repos** under 0069's pre-v1 waiver: identity
  subjects at `<prefix>.identity.>`, workloads env/subjects at
  `SOULSTREAM_*`/`SOULSTREAM.*` (the transient SVC subjects stay
  outside the record stream's `SOULSTREAM.TOPICS.>` capture), the idp
  store prefix renamed. **Old state dirs will not load unchanged** —
  named plainly: `planes.helm` blocks and identity-plane scoped-key
  templates in existing ceremonies carry the old tokens. Re-founding a
  realm is the migration story, as decided.
- One pre-rename `soulstream` (then `soulnode`) daemon was running on
  the dev machine and was stopped gracefully to free its port; its
  state dir is untouched on disk — and is now of the old shape.

What remains open, honestly: the local stdio MCP adapter binary in
soulstream-core is still named `soulstream-mcp` — the same filename the
new remote server releases — a collision deferred to the
**soulstream-cli founding** (which will likely absorb the adapter);
frozen `specs/` and episodes ≤ 0069 keep old names by design (the
naming map in [`README.md`](README.md) resolves them); the
`platform-tenancy-guardrails` topic still speaks the old vocabulary and
gets its terminology pass when next touched; the msb/k8s environment
gates in soulstream-workloads (`test-msb`, `test-k8s`) were not run in
this sweep — hermetic default gates only.

In the hq: design folders renamed to the repo names, links retargeted
across every area including the append-only episodes (the hq-merge
precedent — prose untouched, links follow), living docs re-worded,
hqlint's component list updated, and episode filename tags go **short**
(`core`, `workloads`, `identity`, `idp`, `shell`, `mcp`, `cli`,
`soulstream`, `ecosystem`) because the episode grammar is single-word —
the naming map carries the vocabulary.

Reversal condition: inherited from episode 0069 (the product/client
naming and bare-name reversal readings); the sweep itself records
completed mechanical work — the reversal of any single rename is a new
sweep, not an undo.

Trail: soulstream-core `71b162c`, soulstream-identity `6ce4124`,
soulstream-workloads `c9d6670`, soulstream-archivist `5a3f16a`,
soulstream-idp `382d4de`, soulstream-mcp `7e89fe8` (founding),
soulstream-shell `7975c4e` + `57de0b0` (re-proof), soulstream
`24caf23`; the hq sweep is this change-set.
