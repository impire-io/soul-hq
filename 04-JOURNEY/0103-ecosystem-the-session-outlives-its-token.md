# Episode 0103 — The session outlives its token: the refresh grant crosses the seam (2026-08-17)

**The bug, reported from use:** a signed-in shell went bad "after a
while" — every act erroring, the screen still saying signed in. The
chain, read across three components [mechanism-argument, confirmed by
the gates below]: the idp mints one-hour access tokens and —
deliberately, "not in M1's scope" — no refresh token; the shell copied
that bearer once at sign-in into an in-memory session that never
expires; the identity plane's callout bounds every NATS admission at
15 minutes (its revocation-propagation TTL), so the server bumps the
per-session connection on schedule and each reconnect re-presented the
same aging bearer. Within the hour every bump re-admits silently; at
the first bump past it the callout refuses, nats.go aborts reconnection
on repeated auth errors, and the person's lane is dead for good — while
the admin surface 401s beside it and the cookie keeps saying welcome.

**The fix, in the layers that own it.** soulstream-idp grows the
refresh grant: `offline_access` mints a rotating refresh token beside
the access token — digest-keyed record (D12, the token verbatim
nowhere server-side), single-use by the same CAS flip codes redeem
through, 30 days idle bounds any one token while rotation renews the
line, revocation deletes. The M1 gate grew the scenario [measured]:
a sign-in without `offline_access` still carries nothing; with it the
grant refreshes without a ceremony, the new tokens verify against JWKS
for the same subject, the rotated successor comes back, and the spent
token is refused like a replayed code.

soulstream-shell stops copying: the session custodies the grant — the
access token of the moment plus the source that renews it — and
everything acting as the person asks for the current bearer: the NATS
admission through a token handler on every (re)connect, the admin lane
per call. A session that can no longer produce a living credential
ends honestly on the next request — closed, drained exactly once,
forgotten, the sign-in card instead of a screen of errors
[measured, table-tested]. An issuer that grants no refresh token
leaves the honest half: the session ends with its one token instead of
rotting behind it. The shell's e2e (idp replaced by the local sibling)
ran both halves together green [measured].

**What it taught.** A session is exactly as alive as its credential —
holding one is custody of the *grant*, not of a string; every copy of
a bearer is a place for it to die unnoticed. And the 15-minute callout
bumps turned out to be load-bearing diagnostics: the lane died at the
first bump past the hour, which made "after a while" a readable
schedule instead of a mystery.

Reversal condition: none — records a completed fix. Watch items, not
conditions: per-client token-lifetime knobs stay deferred until a
milestone demands them; the no-refresh fallback (session ends with its
one token) is acceptable until a deployment class says otherwise.

Trail: soulstream-idp `b6d3b2d` (the grant, the gate scenario);
soulstream-shell `3c9f0ed` (the renewing session, the honest ending) —
both on main, tags and the product pin pending; the seam read:
soulstream-identity's callout TTL (`embed/embed.go`, 15m default), the
idp's hour (`internal/provider/provider.go`).
