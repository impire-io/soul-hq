# Episode 0069 — One name: Soulstream (2026-08-13)

The constellation renames. What began as a one-component correction —
*helm* falling to its Helm-charts collision during the shell reframe
(research topic
[`shell-module-contract`](0078-shell-the-module-contract.md))
— widened the same day into the operator's ecosystem-wide question: why
not name every project by its function under the one brand? Decided,
teach-back survived: **the ecosystem is Soulstream, and every component
carries its name** `[judgment]`:

| Today | Becomes |
|---|---|
| soulstream (the record library) | **soulstream-core** |
| soulnode | **soulstream** — the product: one binary that spins up a full soulstream |
| soulrealm | **soulstream-workloads** — it was always meant to run the agents and tools |
| `soulstream/node` (the nested remote-MCP module) | **soulstream-mcp** — extracted to its own repository |
| soulidentity | **soulstream-identity** |
| soulfold | **soulstream-idp** |
| the shell (ex-helm) | **soulstream-shell** |
| soulstream-archivist | unchanged — it was the pattern all along |
| *(new)* | **soulstream-cli** — the client CLI, its own project |

The decisions riding the scheme. The flagship bare name moves from the
library to the product — `soulstream init && soulstream up` becomes
literal `[judgment]`. The operator's one-CLI vision refined in the same
exchange: server product and client CLI split into `soulstream` and
`soulstream-cli` — the nats-server/natscli shape — with soulstream-cli
founding as its own project after the sweep and absorbing
soulstream-core's existing CLI client `[judgment]`. **Backwards
compatibility is explicitly waived pre-v1**, so wire vocabulary
(`SOULREALM.SVC.*`, env vars, config keys) renames fold into each repo's
sweep rather than lingering half-renamed. And *soulsystem* retires: the
ecosystem name is Soulstream, resolving a two-name ambiguity the record
had carried openly — episode 0066's own slug says `soulsystem-cockpit`
`[mechanism-argument]`.

What was refuted or reversed. *soulshell*, decided earlier the same day,
is superseded before it ever executed — the collision-avoidance argument
that killed *helm* and bare *cockpit* (cockpit-project) generalized into
the prefix scheme: compound names behind a unique prefix cannot collide
the way single words do, and function-forward names extend the C8 copy
rule from screens to repositories `[mechanism-argument]`. The
soulstream-mcp extraction also dissolves an ambiguity the mapping itself
created: with soulnode claiming the bare name, "soulstream node" would
have meant two different things.

What it opened: the execution sweep, ordered. The record vacates the
bare name first (soulstream → soulstream-core), *then* soulnode claims
it; GitHub's redirect for the record's old name dies at that moment —
accepted at zero external adoption, and the cost only grows later: Go
module paths don't follow redirects, so every rename is a new import
path with fresh tags and a coordinated pin move `[mechanism-argument]`.
The sweep covers module paths, cross-pins, sibling checkout directories,
CI/release configs, and the hq (design folders, roadmap, hqlint's
component list — for *future* episodes; 0001–0068 keep their tags, the
record is append-only, and a naming map beside the pre-merge numbering
map resolves old references). Component sections in the roadmap keep
their old names until each repo's sweep lands, so the docs keep
describing what is.

Reversal condition: if the product/client split confuses the first
external users — observable as repeated misdirected issues or
"which one do I install" questions over a sustained window — the
one-CLI instinct returns and soulstream-cli folds into the `soulstream`
binary. If the bare-name swap misroutes contributions or imports in
practice, the product takes a suffixed name back and the record's claim
to the flagship name reopens.

Trail: research topic
[`shell-module-contract`](0078-shell-the-module-contract.md)
(the same-day naming trail helm → cockpit → soulshell →
soulstream-shell); roadmap note in
[`ROADMAP.md`](../03-IMPLEMENTATION/ROADMAP.md); commits `01a23c7`,
`76c5084`, and this change.
