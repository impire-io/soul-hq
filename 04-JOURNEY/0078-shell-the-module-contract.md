# Episode 0078 — The module contract: four bars, one shell (2026-08-13 → 2026-08-14)

The question — *what exported module contract lets soulstream-shell
host every human surface as an externally-built module, activated by
what the deployment runs, cross-linked through the shell, while the
shell itself imports nothing Soulstream-specific and custodies
nothing?* — opened the morning the component was still called the
helm, survived two renames and the focus re-scope, and graduated with
all four pre-registered bars measured PASS on real builds.

**The bars, with the honest numbers:**

- **Bar 1 — a pure, agnostic shell, proven by two real modules**
  `[measured]`. One package became four things that do not know each
  other: `shell/` (generic OIDC sessions, chrome, SSE plumbing, and
  the contract — identity, activation predicate, navigation
  contribution, route mounting), `soulstream/` (the module-support
  layer: admission, clients, directory, tray), `modules/overview` and
  `modules/conversations`, with `embed/` the one meeting place.
  `internal/purity` fails the build on any impure shell import — its
  control fired on 7 dependencies. Pixels held: served HTML
  byte-identical across all six screens. The shell's whole external
  cost: go-oidc, oauth2, go-jose.
- **Bar 2 — activation follows the deployment** `[measured]`. The
  module that is not always there — "People & sign-in" over the idp
  plane — activates on one fact the deployment declares (optional
  `AdminBase` in the product's embed options); `Active()` is a single
  comparison. Present arm 9.19 s: people listed, a real invite issued,
  authority riding the signed-in person's own bearer (the idp needed
  zero changes — its admin surface already verified its own tokens).
  Absent arm 3.34 s: idp plane off, sessions on an external AS, no
  rail entry, 404s, everything else working.
- **Bar 3 — the standing doors survive** `[measured]`. On published
  tags only (product v0.9.0 → shell v0.3.0), the sole replaces being
  the repo-under-test and the never-published probe; custody scan
  green with its fired control inside the full ceremony; 13.5 s
  uncached.
- **Bar 4 — the shell composes** `[measured]`. The link facility: a
  module names who it wants to reach and what about, the shell asks
  the modules this deployment runs, the owner builds its own link or
  declines — every name in a conversation's People panel is a way into
  that person's sign-in, with no import between the modules, degrading
  to plain text when the module is absent. And `e2e/moduleprobe`,
  outside even the impire-io namespace, composed a rail entry and a
  screen through the exported contract alone — the toolchain itself
  its jailer.

**What was refuted or reversed along the way:** the second Bar 1
module was amended pre-experiment from fold-administration to the
collaboration surface when the focus landed (episode
[0071](0071-ecosystem-the-focus.md)) — and fold-administration
arrived anyway, as Bar 2's absence proof. C4's rendering pick met its
own reversal condition (the autocomplete interaction) mid-topic and
survived with its cost named (episode
[0076](0076-shell-a-name-that-taps-somebody.md)). The topic also
carried the naming trail: helm → cockpit → soulshell →
soulstream-shell in two days, each fall recorded.

**What it opened, named:** extraction of the pure shell into a
standalone framework stays gated on a second real consumer — cheap by
construction, the purity test is the guarantee. The idp's own `/admin`
console is now duplicated by the shell's People & sign-in module; its
retirement to ceremony pages is soulstream-idp's scope decision, by
demand. Module packaging beyond in-repo packages is proven possible
(the probe) but not needed yet.

Reversal condition: the topic's own conditions stay live as standing
gates — if a future module cannot earn its place without
module-specific shell code or a shell-held store, the purity test is
where it will show; if the shell cannot stay Soulstream-agnostic, the
concrete need that breaks it reopens the agnosticism constraint with
that need as evidence.

Trail: design
[`0002-the-module-shape.md`](../02-DESIGN/soulstream-shell/0002-the-module-shape.md)
(the graduated design); the topic journey (in git history at
`01-RESEARCH/shell-module-contract/`, folder removed at graduation);
soulstream-shell `4960ef7` → `1522586`, soulstream `ce73152` +
`9421778` (v0.9.0); episodes 0071–0077 carry the product slices the
bars rode.
