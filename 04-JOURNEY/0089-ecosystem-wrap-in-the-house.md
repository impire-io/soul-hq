# Episode 0089 — Wrap in the house: one binary, one paste (2026-08-15)

The operator looked at the agents screen and called it what it was: too
dense, and the wrap — the only way an agent gets the full experience —
buried as the last of four folds behind a manual five-variable export
and two `go install` lines. The direction that reshaped the day: people
must not need a Go toolchain; the `soulstream` command carries what's
needed. Easy is easy, hard stays possible.

What the audit found first was sharper than a layout complaint: the
product binary the releases page hands out had **no `wrap` verb at
all** — design 0004 §8 gave `soulstream wrap` to the *core CLI's*
external-subcommand seam, and the product never picked it up, so the
getting-started's own step 7 could not work with the binary its step 1
downloads [measured — `unknown command "wrap"` from the shipped
usage]. The fix ran the pipeline whole, four repos in one arc:

- **Design** (this repo): 0004 §5 grew `mcp_args` (the per-run MCP
  config's `args` array, so a *subcommand* can be the tool door), §8
  recorded the product-native occupancy beside the seam — the
  embedding alternative won the adversarial pass against
  self-provisioning downloads (a supply-chain surface with no custody
  answer) and multi-binary archives (goreleaser builds only this
  module's mains) [judgment, argued]; new product design
  [`0002-wrap-in-the-house.md`](../02-DESIGN/soulstream/0002-wrap-in-the-house.md)
  pinned the verbs, the five-name lane, and the paste-block contract
  (portable: POSIX *and* fish — no heredoc, no `export`).
- **soulstream-workloads** (`specs/007-mcp-args`, v0.4.0):
  `Template.MCPArgs`/`Lane.MCPArgs` through preset and
  `writeMCPConfig`, additive — an argless template writes the exact
  shape it always wrote [measured, gate green].
- **soulstream** (`specs/008-wrap-in-the-house`): `soulstream wrap`
  (the wrap library, door pointed at `os.Executable() + ["mcp"]`) and
  `soulstream mcp` (`realm.Connect` + core's public `mcpserver` over
  stdio, environment-only — a config file can never carry a
  credential). `getting-started.md` teaches download-and-paste;
  `go install` is gone from it [measured: grep exits 1].
- **soulstream-shell** (v0.5.0): the credential card leads with the
  paste block — writes the creds file itself, so the same block works
  on any machine, closing the latent server-local-path wrinkle — with
  a Copy key, and folds the hard paths (`.mcp.json`, codex TOML,
  plain-words MCP, the creds file) as quiet cards below, all spelling
  the door `command: soulstream, args: ["mcp"]`. The screen took the
  canon's rhythm: sections breathe, the stray "—" note renders empty,
  an unnamed operator's handle stands alone, the roster's empty state
  points at the form, and the overview grew an Agents card with the
  way to set one up. Gate green including the consumer-position e2e
  (14.5 s); no horizontal scroll at 390 px [measured].

**Proven live, the whole story** [measured]: release-shaped binary →
`init` → `up` → passkey enrolled in a real browser → agent "clerk"
created on the screen → the paste block copied *unedited* and run
under fish with only the product binary on PATH → a mention posted
while the wrapper was **off** was caught up and answered by a real
`claude -p` run in 6.697 s — exactly one reply under the deterministic
wake id, self-wake guard holding, the teal machine-channel card on the
shell with its honest `unsigned` verdict. `soulstream mcp` answered an
MCP initialize over stdio; the revoked credential was refused loudly
(`Authorization Violation`, exit 1) with nothing posted. One
real-world wrinkle recorded: the operator's fish config re-orders
PATH, so a stale installed `soulstream` shadowed the fresh one until
`--no-config` — the paste block itself was never edited.

Standing exception, tracked openly: the product's `go.sum` cannot
carry the workloads v0.4.0 / shell v0.5.0 pins until those tags are
pushed (`go mod tidy` ignores workspace replaces by design); a local
`go.work` carried the gates. Push order: workloads (main + v0.4.0),
shell (main + v0.5.0), then `go mod tidy` in the product and push it —
pushing stays the human's act.

Reversal condition: a fix landing in `soulstream-wrap` that the
product cannot pick up within its release cadence — observed as
divergent behavior for the same template — moves the product back to
shipping the component binaries in its release archive instead of
embedding the verbs (recorded in design 0004 §8).

Trail: designs
[`soulstream/0002-wrap-in-the-house.md`](../02-DESIGN/soulstream/0002-wrap-in-the-house.md),
[`soulstream-workloads/0004-wrap.md`](../02-DESIGN/soulstream-workloads/0004-wrap.md)
(§5/§8/§10 amended); specs `../soulstream-workloads/specs/007-mcp-args/`,
`../soulstream/specs/008-wrap-in-the-house/`; commits — soul-hq
`eebfd87`; workloads `a6b92e4`/`e46ee53`/merge `a7de142`, tag v0.4.0;
product `8f6a114`/`b001c5c`/merge `95e436b`/`af25a99`; shell
`2a78e07`, tag v0.5.0.
