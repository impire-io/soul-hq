# 0002 — Wrap in the house: one binary, one paste

**Status**: designed 2026-08-15 at the operator's direction (no Go
toolchain, no PATH assembly on an agent's machine); realizes the product
half of workloads design
[`0004-wrap.md`](../soulstream-workloads/0004-wrap.md) §8 and closes its
§10 "shell surfacing" item. Validation tags per the
[design README](README.md); requirement language per the same.

The capability, in one sentence: **a person who has downloaded the
`soulstream` release binary and signed in to their assistant can run
their agent by pasting one block from the Agents screen — nothing else
installed, nothing edited.**

## 1. The two verbs [D]

The product CLI (`cmd/soulstream`) grows two verbs. Both are thin mains
over libraries the product already pins; neither adds wire vocabulary.

### `soulstream wrap`

The personal wrapper of workloads design 0004, out of the product
binary: `wrap.Preset`/`wrap.LoadTemplate` + `wrap.Wrapper` from
`soulstream-workloads/wrap`. Flags MUST mirror `soulstream-wrap`:
`--harness claude|codex`, `--template <file>`, `--scratch`,
`--run-timeout`, `--retries`, `--inbox-limit`. The lane resolves from
the five `SOULSTREAM_*` environment names (URL, CREDS, TOKEN, REALM,
PERSONA) — the same block the Agents screen mints; flags win over env
where both exist. Refusals are loud and name the missing piece, in
`soulstream-wrap`'s own wording (a revoked credential is refused at
connect, before anything is posted).

The harness's tool door is the product's own executable: the lane sets
`MCPCommandLoc = os.Executable()` and `MCPArgs = ["mcp"]` (0004 §5's
`mcp_args`), so the per-run MCP config launches `soulstream mcp` and
**no second binary exists anywhere in the path**.

### `soulstream mcp`

The stdio tool door: `realm.Connect` from the same five env names
(flags `--url --creds --token --realm --persona` win where given), then
core's public `mcpserver.NewServer(client)` served over the MCP stdio
transport. No context files, no keystore, no config-dir resolution —
the agent lane is environment-only, because a config file can never
carry a credential. An agent with no signing key speaks unsigned; that
is the lane's honest state today (measured in the shell's e2e) and this
verb MUST NOT pretend otherwise.

What this deliberately is not: core's `soulstream-mcp` (contexts,
key files, per-project identity) stays the component world's door; the
core CLI's external-subcommand seam (0004 §8) is untouched.

## 2. Composition [D]

Per design 0001, the product is composition: both verbs consume public,
tagged surfaces — `soulstream-workloads/wrap`, `soulstream-core/realm`,
`soulstream-core/mcpserver` — already in `go.mod`. No goreleaser
change; the one release binary simply answers more verbs. Usage text
names both.

## 3. The lane contract [D]

One block, five names, minted once by the Agents screen (episode 0079's
credential): `SOULSTREAM_URL`, `SOULSTREAM_CREDS` (a *path* — see §4
for how the screen makes it true on any machine), `SOULSTREAM_TOKEN`,
`SOULSTREAM_REALM`, `SOULSTREAM_PERSONA`. Spelled exactly this way in
every surface that prints or reads them; one wrong name is an agent
that does not start.

## 4. The credential-screen contract [D]

The shell's Agents screen is the product's console, so what it prints
is a product contract:

- The shown-once credential card leads with **the paste block** — the
  primary path, before any fold. The block, pasted unedited into a
  terminal on any machine:
  1. creates `$HOME/.soulstream/` and writes the sentinel creds file
     to `$HOME/.soulstream/<handle>.creds` (the sentinel is public by
     construction — a bearer that denies itself everything — so
     printing it is free; *writing it locally* is what makes the block
     machine-independent, where the deployment's own path is not),
  2. runs `soulstream wrap --harness claude` under the five values
     (`--harness codex` offered as a comment line; `--template` in
     prose beside it).
- **Portability rule**: the block MUST run unchanged under POSIX shells
  *and* fish — no heredoc, no `export`; a multiline single-quoted
  `printf` write and an `env VAR=… soulstream wrap …` prefix are the
  proven-portable shapes.
- The block carries the one secret (`sit_` token); it remains
  shown-once, replaced by the next act on the screen, and the screen
  keeps saying so.
- The demoted folds keep the hard paths possible: Claude Code
  interactive (`.mcp.json`), codex TOML, anything-else MCP — all now
  spelling the door as `command: soulstream, args: ["mcp"]`, the one
  binary everywhere.
- `docs/getting-started.md` steps 6–7 teach the same: download, paste.
  The strings `go install` MUST NOT appear in the getting-started.

## 5. Acceptance criteria

1. On a machine with only the release `soulstream` binary and a
   signed-in assistant, the paste block from the Agents screen —
   unedited — answers a mention posted while the wrapper was off:
   exactly one reply turn authored by the agent (0004 §11.1 through
   the product verb).
2. `soulstream mcp` with the five env values answers an MCP initialize
   over stdio; with the credential revoked it refuses loudly and
   serves nothing.
3. `soulstream wrap` with a missing persona/realm/connection refuses
   with the message naming what is missing; an unknown `--harness`
   names the two presets and the template escape.
4. The credential screen at 1440px and 390px: the paste block above
   the folds, no horizontal document scroll, section rhythm per the
   shell's design canon.
5. `getting-started.md` contains no `go install`.
