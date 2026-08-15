# Episode 0092 — The names say what they do (2026-08-15)

The bynames "the door" and "the fold" left every surface a user reads —
the operator's direction, taken all the way: slogans and identifiers
included. The inventory that drove the sweep counted ~93 prose
occurrences and ~20 identifier-level ones across the product, the idp,
the shell, the core docs, and the website [measured, the sweep agent's
file-by-file inventory].

**soulstream** (`specs/009-one-console-one-vocabulary`): the config keys
became `planes.signin` and `planes.mcp` with the byname-era spellings
read forever; fresh founds write `users/signin.creds` and the `signin`
NATS user, while founded realms keep their artifacts by fallback — a
founded user's name is baked into its account JWTs and cannot be renamed
without re-ceremony [mechanism-argument]. Flags gained
`--signin-listen`/`--mcp-listen` with the old spellings accepted; `up`
prints functional labels; the migration fixture (legacy keys + legacy
creds name) rides `make test` [measured]. Verified live: a fresh found
writes `planes: [memory, mcp, signin, shell]` and `signin.creds`, `up`
prints the new labels, `/admin` answers 404 [measured].

**The website** (impire-io.github.io): the story, the pages, the map
legend, and the mirrors — the flagship line is now **"One protocol. No
bot lane."** The get-started page mirrors the product's real output.
One honest wrinkle for the record: the site commit swept in the
operator's in-flight story-detail components (`MapDetails.astro`,
`SoulStory.astro` changes) because the working tree was staged with
`-A` against the hq's own explicit-pathspec rule — the build is green
and the components are additive, but publishing them was not this
sweep's call to make. Their newest deltas remain local to the
operator's tree.

**soulstream-core docs**: component-referent "door" language went
functional (the MCP adapter, the remote endpoint, the front desk that
checks who came in — titles included); the workshop analogy keeps its
literal furniture — a fridge door is a fridge door [judgment on the
line between byname and object].

**soulstream-idp / soulstream-shell / soulstream-identity / soul-hq**:
READMEs, page copy, flag help, and the hq roster label swept in the
same pass (the idp's README also stopped claiming genesis five shipped
milestones later).

Reversal condition: none for the copy — records a completed sweep. For
the identifier rename, the read-forever aliases are the reversal
insurance: if the dual-key read is ever measured to confuse (a support
question naming both spellings), the legacy keys gain a deprecation
note in the docs, never a removal.

Trail: designs
[`0001-soulnode-composition.md` §2](../02-DESIGN/soulstream/0001-soulnode-composition.md);
`../soulstream/specs/009-one-console-one-vocabulary/`; commits —
product `c61072d`/merge `a561828`, site `8912d63`, core `6b4f831`,
identity `c7d832e`, idp `7ce4383`, shell `73cc95b`.
