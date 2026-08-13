# Episode 0068 — The helm plane: the cockpit joins the bundle (2026-08-13)

Soulnode composed soulhelm the day it was born: `planes.helm`
(`specs/007-the-helm-plane/`, merged `e9d85ad`) runs the human cockpit
through soulhelm's public embed seam at its published tag v0.1.0 —
composition, not invention, twice over: the node hands the helm its
ops read lane, the public sentinel, and the resolved sign-in issuer;
the helm founds and owns nothing. `up` now logs a fourth URL:
`helm console`. On by default beside the fold; **absent block means
disabled** on state dirs founded before the plane existed (the fold's
upgrade rule, reused) [measured: `TestHelmWiring`].

**The decision of record — `SessionIssuer()`**: enabling the helm
switches the identity plane's OIDC admission lane on *in local mode*,
pointed at the bundled fold — before this, the lane existed only under
public door mode. An explicit `planes.door.auth_issuer` (external AS)
wins; the door's all-three-or-none public-mode contract is untouched;
an ephemeral fold listener (`:0`) resolves to no issuer and the helm
skips its surface with an audit warning, never silently
`[mechanism-argument]`. This was found the honest way: the helm's own
e2e hit `Authorization Violation` on a default local boot, because
fold bearers had nowhere to admit.

A first-draft hazard was caught before it shipped: wiring the issuer
default at config Load would have persisted two of the three
public-door fields and tripped Verify's all-three-or-none refusal on
the next save — the resolution moved to a method on the ceremony
state, mutating nothing [mechanism-argument].

Gate [measured]: the whole soulnode suite green with the plane in —
`TestHelmPlane` (the surface serves, closed until sign-in, `/live`
refuses 401), `TestHelmDisabled` (no URL), `TestHelmWiring` (defaults,
absent-block, collision and issuer refusals by name) — zero `replace`
directives, soulhelm pinned by tag.

Named, not built: a dedicated scoped `helm` ceremony user (today the
plane hands the ops lane — operator standing, more than observe
needs); the configure surfaces in the helm UI; the grant lane for
standalone helms (rides `platform-tenancy-guardrails`).

Reversal condition: if the helm plane's default-on posture measurably
burdens foundings that never use a browser (observable as operators
disabling it as a matter of course), it becomes opt-in like the fold
was at its own genesis.

Trail: soulnode `specs/007-the-helm-plane/`, commits `d033175` +
merge `e9d85ad`; soulhelm episode
[0067](0067-soulhelm-founding-and-first-light.md); design
[`0001-soulhelm-the-helm.md`](../02-DESIGN/soulstream-shell/0001-soulhelm-the-helm.md).
