# 0002 — soulstream-shell: the module shape

**Status:** open — graduated from research 2026-08-14 (episode
[0078](../../04-JOURNEY/0078-shell-the-module-contract.md)); every
load-bearing claim was measured there unless tagged otherwise. This
document describes the architecture as built (shell v0.3.0, composed
in soulstream v0.9.0) and is the argument for any future module.

## §1 The four layers

- **`shell/` — the pure frame.** Generic OIDC sign-in and in-memory
  sessions (bearer custody, nothing persisted), page chrome (transport
  bar, icon rail skeleton), Datastar/SSE plumbing, design assets, and
  the module contract. It says nothing about the product it frames
  that composition did not hand it. Its import purity is
  **compiler-grade and standing**: `internal/purity` walks
  `go list -deps` over every `shell/` package and fails on any module,
  support-layer, or `github.com/impire-io/*` path; its positive
  control must keep firing. Allowed external cost: go-oidc, oauth2,
  go-jose.
- **`soulstream/` — the module-support layer.** Everything the shell
  refuses: bearer → sentinel + auth-callout admission acting as the
  signed-in person, realm/identity clients, the persona directory, the
  mention tray. Modules import support; the shell never does.
- **`modules/*` — the human surfaces.** Each registers through the
  contract and imports the shell's exported packages, the support
  layer, and components — never another module (standing test:
  no-module-imports-another).
- **`embed/` (and `cmd/`) — composition.** The one place shell,
  support, and modules meet; the product composes here (its plane
  hands deployment facts through `embed.Options`).

## §2 The contract

A module presents: **identity** (slug + plain-language name),
**activation predicate** (`Active()` — a comparison over facts the
deployment declared, never probing, never shell-side config),
**navigation contribution** (rail entries; badges own their patch
targets), and **route mounting**. Composition passes declared facts
(issuer, `AdminBase`, …) downward; absence of a fact deactivates the
module that keys on it — routes 404, navigation contributes nothing.

**Cross-linking** (`shell/link.go`): a module asks the shell for a way
into another module's screen by identity + screen kind + subject; the
shell puts the ask to the active modules; the owning module builds its
own link or declines; nothing else resolves. The asking module renders
the href or its honest fallback (plain text) — never a dead link.

## §3 The custody rules (inherited, re-proven)

Design [0001](0001-soulhelm-the-helm.md)'s doors stand over the module
shape: sessions in memory only, the shell signs and acts as no one
(every act rides the signed-in person's own admission and, for
class-(b) surfaces, their own bearer — delegated authority, never
borrowed), and the consumer-position e2e rides published tags with the
sole replaces being the repo under test and the local probe.

## §4 Acceptance (the bars as standing gates)

1. `internal/purity` green with its control firing.
2. Both e2e arms green: the full ceremony (present) and the
   external-AS arm (idp plane absent — no admin navigation, 404s,
   everything else working).
3. The custody scan clean with its fired positive control.
4. The cross-link both-arms assertions and the outside-namespace probe
   (`e2e/moduleprobe`, contract-only imports, zero shell changes).

## §5 Open [O]

- **[O1] Extraction.** The pure shell lifts into a standalone
  framework when a second real consumer exists; the purity test keeps
  that cheap. Until then, in-repo packages.
- **[O2] The idp's own console.** "People & sign-in" duplicates
  soulstream-idp's `/admin`; retiring that console to ceremony pages
  is the idp's scope decision, by demand.
- **[O3] Third-party modules.** Proven possible (the probe); a real
  outside module decides packaging conventions when it arrives.
