# Episode 0106 — The token-lifetime knob: revocation propagation gets a dial (2026-08-18)

The first of episode 0104's two due demands on the fold, closed the day
after it was called: the access-token lifetime — a package constant
(1h) since M1, deliberately deferred "until a milestone demands it"
(0103's watch item) — is now a deployment knob. Outbound revocation
propagates in access-token `exp` + callout TTL [measured, 0104 Bar 1],
so the deployment that cares about the bound tunes the lifetime: one
value through the whole seam (`provider.New`, `serve.Options`,
`embed.Options.AccessTokenLifetime`, `--access-token-lifetime`), zero
meaning the default hour. Shipped as **v0.7.0**.

Measured in the e2e on both arms [measured]: a fold serving with a
5-minute lifetime issued an access token whose `exp − iat` is exactly
5 minutes through the full ceremony (authorize → passkey → code →
exchange), and the unset knob still issues the hour — the default is a
default, not a new behavior. Gate green across all three modules, no
skips.

Deliberately not built, named: per-client lifetimes (the constant's
original comment survives — no milestone has demanded them) and the
product-side plumbing (`planes.signin` does not yet expose the knob;
that lands with the next product cycle that wants it). The second 0104
demand — the RFC 8693 exchange grant — remains the fold's next due
item, spec first.

Reversal condition: none — records a completed build/measurement. (The
knob's default reverses trivially by configuration; no direction was
taken.)

Trail: soulstream-idp commit `dcba05a`, tag `v0.7.0`; episode
[0104](0104-ecosystem-outbound-identity-grants.md) (the demand and its
measurement), [0103](0103-ecosystem-the-session-outlives-its-token.md)
(the watch item that armed it); e2e `TestAccessTokenLifetimeKnob` in
`e2e/issuers_test.go`.
