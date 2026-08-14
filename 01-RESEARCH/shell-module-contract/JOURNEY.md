# Journey — shell-module-contract

*The investigation as it happens. Bars are in [`README.md`](README.md).*

**2026-08-13 — opened, then reframed before any experiment ran.** The
topic opened as `helm-skeleton` at the operator's direction (the helm
becomes the surface humans use most; every surface a module that earns
its place). The same session, two corrections sharpened it, both the
operator's: the shell contains **no module logic at all** — nothing
about the fold or any component goes into it; it facilitates modules
working together (cross-linking) and brings things together — and the
component renames: *helm* fell for the Helm-charts collision, bare
*cockpit* for cockpit-project, and **soulshell** won — agnostic from
Soulstream by contract, extraction into a generic framework deferred to
a second real consumer. Bars rewritten accordingly before any spike;
the folder renamed to `shell-module-contract`.

**Later the same day — the name widened past this topic.** The
collision-avoidance argument generalized into the ecosystem naming
re-centering ([episode
0069](../../04-JOURNEY/0069-ecosystem-one-name-soulstream.md)): every
project renames soulstream-*, and this component is
**soulstream-shell** — *soulshell* superseded before it ever executed.
The bars are untouched; only the name moved.

**2026-08-14 — the focus re-scope (episode 0071), before any spike.**
The operator's slim-down: participation enters the shell's scope (the
usable cockpit — view, collaborate, mention notifications — is the
value center), so Bar 1's second module is amended pre-experiment from
fold-administration to the **collaboration module**; fold admin waits
as a later module. The bars otherwise stand. Amendment recorded here
per the pre-registration rule.

**2026-08-14 — Bar 1 measured: PASS.** The spike (soulstream-shell
`4960ef7`) split the one package into four things that do not know
each other: `shell/` (the pure frame — generic OIDC sessions, chrome,
SSE plumbing, and the exported contract: identity, activation
predicate, navigation contribution, route mounting), `soulstream/`
(the module-support layer — bearer → sentinel + callout admission,
clients, directory, mention tray), `modules/overview` and
`modules/conversations` (the two real surfaces, each registering
through the one contract, importing neither the other nor the
composition), and `embed/` (the one place the pieces meet). The
agnosticism is compiler-grade and standing: `internal/purity` walks
`go list -deps` over every package under `shell/` and fails on any
module, support, or component path — its positive control runs the
same walk over the support layer and fired on 7 dependencies
[measured]. The shell's whole external cost: go-oidc, oauth2,
go-jose. Behavior and pixels unchanged [measured]: every e2e
assertion untouched and green (~9.5 s uncached), served HTML
byte-identical across all six screens (per-run ids masked), and a
fresh browser screenshot indistinguishable from the pre-spike canon
run. One honest structural note: the work-open act moved to the
screen that offers it — no module serves another's button; the
contract did NOT degenerate toward special-casing (the reversal
condition's observable stayed quiet). Bars 2–4 remain.

**2026-08-14 — Bar 2 measured: PASS.** The third module is the one
that is not always there: `modules/admin` — "People & sign-in" on
screen — list people, create an invite, take a sign-in away, through
the same exported contract with zero shell changes (the purity graph
is unchanged: seven non-stdlib packages). Activation is one declared
fact: the product's `embed.Options` grew optional `AdminBase` (absent
when the idp plane is off — soulstream `ce73152`), the support layer
carries the lane, and `Active()` is a single comparison — no probe,
no reachability guess, no shell configuration [measured]. Authority
is delegated, never borrowed: every call rides the signed-in person's
own bearer, and the idp needed NOTHING — `/api/admin` already
verified its own issued bearers with the admin role. Both arms
measured (shell `8366685`): present — rail entry, seeded people
listed, a real invite issued by the idp (TestShellGate, 9.19 s);
absent — the product boots with the idp plane off and sessions on a
standalone external AS, no rail entry, admin routes 404, sign-in and
conversations and mentions all working (TestExternalIdPGate, 3.34 s)
[measured]. Named for the stable-point evaluation: the founding
group value `realm` surfaces as data in GROUPS — a product-ceremony
naming follow-up, not shell copy. Bars 3–4 remain.
