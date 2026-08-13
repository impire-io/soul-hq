# Episode 0067 — Founding and first light: the helm is real (2026-08-13)

Hours after its research graduated (episode
[0066](0066-ecosystem-soulsystem-cockpit.md)), the sixth component was
founded and shipped: **soulhelm v0.1.0**
([impire-io/soulhelm](https://github.com/impire-io/soulhelm)), the
soulsystem's human cockpit, built by porting the topic's proven rigs
into real packages. The whole build rode the design's own acceptance
gates and landed green the same day [measured].

What v0.1.0 is: the **observe surface** — board, topics with signature
verdicts *earned* from the identity plane's directory, storage
readings from the stream, plane health probes — rendered server-side
in the cassette-light token source (fonts and icons embedded, zero
external fetches) and pushed as Datastar SSE morphs; **fold sessions**
— the helm registers itself with the deployment's AS via RFC 7591 DCR,
runs code+PKCE, and each session opens its *own* NATS admission
(sentinel + the person's fold-issued bearer through the OIDC callout
lane), memory-only; and the **first act** — `work.open` through the
session's connection, attributed and signed as the signed-in
principal. The helm signs as no one and custodies nothing. Public
`embed.Run` seam (D29 pattern) plus a standalone `soulhelm serve`.

One scope amendment, made openly: the roadmap had split M1 (observe)
from M2's sessions, but shipping an *unauthenticated* realm viewer
would have violated the design's own custody story — so sessions moved
into the founding release, and the surface is closed until sign-in
`[judgment]`. The configure surfaces (classes (b)/(c)) remain the next
milestone.

The standing e2e gate is the research made permanent [measured]: a
consumer-position module (`soulhelm.invalid/e2e`, upstreams at
published tags, soulnode booted in-process) walks passkey enrolment
(soulfold's public `authtest` doing real ES256 ceremonies), sign-in,
the live view, an act that lands in the realm authored by the fold
principal, sign-out, the Bar-2 custody scan with a fired positive
control, and the Bar-4 offline-render check — ~4 s inside `make test`.
One correctness fix earned during the build: the DCR redirect URI must
carry the *bound* listener address, so ephemeral ports work.

Reversal condition: the graduated one, kept live — if the helm cannot
remain a pure consumer (a required `internal/` import, a privileged
surface existing only for it, or a helm-owned store of record), the
component dissolves back into the existing surfaces.

Trail: design
[`0001-soulhelm-the-helm.md`](../02-DESIGN/soulstream-shell/0001-soulhelm-the-helm.md);
soulhelm `bb02c24` (founding commit), tag `v0.1.0`; composed by
soulnode the same day (episode
[0068](0068-soulnode-the-helm-plane.md)).
