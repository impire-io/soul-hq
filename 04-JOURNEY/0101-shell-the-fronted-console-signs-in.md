# Episode 0101 — The fronted console signs in: PublicURL crosses the seam (2026-08-16 → 08-17)

**The bug, found by the first fronted deployment:** the byon realm's
console went behind `tailscale serve`
(`https://beno1.hippogryph-dinosaur.ts.net:8443` → loopback 8500), the
passkey ceremony succeeded — the fold issued a code — and the browser
was then sent to `http://127.0.0.1:8500/callback`: the visitor's own
machine. The shell registered its *bound* address as the OAuth
redirect at DCR time, because nothing in its options could say what
origin browsers actually reach it on. Correct exactly when the browser
is on the same host (the bundle); a dead end behind any front
[measured live].

**The fix, in the layers that own it:** soulstream-shell **v0.7.0**
grew `Options.PublicURL` — the DCR registration and the oauth2 config
build the redirect from it (`redirectBase`, unit-tested; bound address
stays the default) — and the product wires **`planes.shell.public_url`**
through the ceremony, the config round-trip (garbage refused by name),
the shell plane, and the `up` output. Deployed on beno1, the live
`/login` now carries
`redirect_uri=https://beno1…:8443/callback` [measured on the
deployment], and the operator's passkey sign-in landed in the console
— the first human standing in a fronted shell on a BYO substrate
[measured, 2026-08-17].

**The honest middle: rc.3 shipped a non-fix.** The config, ceremony,
docs, and pin all landed — but the one line handing `public_url` to
the plane was edited with a whitespace-sensitive string replace that
silently no-oped, no test covered the pass-through, and a green gate
released it. The deployment's own `/login` redirect exposed it in
minutes. rc.4 carries the line (via an edit that fails loudly on
mismatch) and a plane-level assertion on the LIVE `/login` redirect in
`TestHelmPlane` — the composition gate now observes the fact itself,
not just the config that implies it. Same lesson as this arc's piped
exit codes: a gate only guards what it actually observes; this session
banked that lesson twice and both spots now fail loudly.

Reversal condition: none — records a completed fix, measured on the
deployment class that found it. Watch item, not a condition: the
fronted console's session cookie carries no `Secure` attribute yet;
it rides HTTPS today because the front terminates TLS — worth closing
when public-mode hardening gets its pass.

Trail: soulstream-shell `1c39067` + tag `v0.7.0`; soulstream `9f344f3`
(rc.3, the wiring gap), `ccf0dfd` (rc.4, the line + the test); tags
[`v0.13.0-rc.3`](https://github.com/impire-io/soulstream/releases/tag/v0.13.0-rc.3)/[`v0.13.0-rc.4`](https://github.com/impire-io/soulstream/releases/tag/v0.13.0-rc.4)
(both pipeline-fed to the tap); the byon deployment (episode
[0099](0099-soulstream-the-byon-founding.md),
[`DOGFOOD.md`](../03-IMPLEMENTATION/DOGFOOD.md)).
