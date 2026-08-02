# Episode 0053 — M2 ships: passkeys, the refusal becomes behavior (2026-08-02)

The founding refusal — passkeys, not passwords (constitution I) —
stopped being policy and became what the code does: WebAuthn
registration and login ceremonies from `go-webauthn` replaced the M1
seeded-user stub (`specs/002-passkeys/`, branch `002-passkeys`). The
login page's form POST is gone; the only way through the fold is a
ceremony. The page carries the one script D9 always reserved the right
to carry — `navigator.credentials` is unreachable without it — and
nothing else.

The gate, measured in `make test` [measured]:

- **Register-then-login at library level**, driven by a virtual
  authenticator doing real ceremonies — ES256 keys, honest rpIdHash,
  flags, counters, ASN.1 signatures, CBOR "none" attestation. First
  touch (zero credentials) is a registration that appends exactly one
  credential record; second touch is an assertion that advances the
  stored sign count; a consumed ceremony record refuses replay.
- **The e2e gate now runs the passkey flow end to end**: the stock
  go-oidc RP's sign-in authenticates only by ceremony (begin →
  authenticator → finish → callback → code → tokens verifying against
  JWKS); the mid-flow restart lands between ceremony and exchange and
  stays invisible; the browser session still skips the page; the
  retired username-only POST signs nobody in.
- **The name is a wall** (D14, session acceptance #4): ceremony
  responses from a scheme flip, a port change, a subdomain, and a
  foreign host are all refused server-side — four shapes, zero
  passes.
- **Nothing stored can impersonate** (constitution I): the
  credential's private scalar appears nowhere in the opened records;
  the same scan finds the public key, so it proves something.
- Forged begin/finish POSTs (missing/wrong CSRF, foreign Origin)
  leave the auth-request revision unmoved; the one-shot token spans
  both ceremony POSTs and is cleared only in the CAS write that
  completes the flow (D13).

Two things said honestly rather than quietly: **first-touch
enrollment** (the account's first authentication binds the passkey) is
the M2 stand-in for M3's researched bootstrap story — trust on first
authentication, loud in the quickstart, not a durable policy
[judgment]; and the **physical-authenticator runbook** is documented
(quickstart.md) but running it is a human act — the virtual
authenticator proves the ceremonies, not a finger on Touch ID. That
run is pending.

What it opened: M4 is next (the operator's public-door priority,
episode 0052's sequencing note): the fold now issues real tokens for
real passkey users, which is exactly what the callout admission proof
wants to verify.

Reversal condition: none — records a completed build against
already-graduated designs; D14's naming door and the M3
bootstrap-story research stand as written.

Trail: `specs/002-passkeys/` (spec, plan, tasks, quickstart with the
runbook); design
[session-and-ui](../02-DESIGN/soulfold/session-and-ui.md) (D9, D13,
D14) and [store-and-key-lifecycle](../02-DESIGN/soulfold/store-and-key-lifecycle.md)
(D2/D6 additive credentials); the `002-passkeys` merge in the soulfold
repo.
