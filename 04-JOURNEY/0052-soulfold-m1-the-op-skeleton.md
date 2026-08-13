# Episode 0052 — M1 ships: the OP skeleton on the sealed store (2026-08-02)

The fold's first product code landed the same day its last research
gate concluded: discovery, JWKS, and the authorization-code + PKCE flow
served by `zitadel/oidc` v3, storage on the four D1 buckets under the
D16 envelope, the D7/D8 key lifecycle, the two-page D9 surface with the
D13 walls, and a seeded user and client standing in for the ceremonies
(`specs/001-op-skeleton/` in the repo, feature branch `001-op-skeleton`).

The gate, all in `make test` against a real embedded nats-server, no
mocks [measured]:

- **A stock go-oidc RP completes sign-in end to end** — discovery,
  authorize, login POST, callback, token exchange — holding an ID token
  and a JWT access token (D15) that verify against the published JWKS,
  subject the seeded user. The page inventory across every flow in the
  gate is exactly {login, error} (D9).
- **The process is disposable, the store is not** (D11): a full
  fold-and-server restart injected between the login POST and the token
  exchange is invisible — the exchange completes and verifies; a
  browser session minted before the restart completes a fresh sign-in
  after it with zero pages rendered; nothing is re-seeded.
- **Forged state-changing POSTs change nothing** (D13): missing CSRF,
  wrong CSRF, and valid-CSRF-with-foreign-Origin are all refused with
  the auth-request record's revision unmoved — and the legitimate
  submission still completes afterward, because refusals never consume
  the one-shot token.
- **Keys roll over under a live verifier** (D7's I1/I2 in code, store
  acceptance #4): a never-restarted go-oidc verifier saw 61
  verifications with 0 failures across create-pending → activate →
  retire; retirement refused while the straggler's
  `last_signed_expiry` was in the future; after it passed, the retired
  key left JWKS and a fresh keyset refused the straggler while
  verifying the active key.
- **The store defends itself** (store acceptance #1–#3, #5): restart
  byte-identical ciphertext; the additive matrix through the envelope
  with the D3 RMW trap still demonstrable; exactly-once code
  redemption 50/50 races; the custody scan clean over the stopped
  store dir and the API-level dump with an unsealed positive control
  proving the scan; `expires_at` authoritative on reads (D5); a nil
  sealer refused outright — the fold has no plaintext mode (D16); a
  group- or world-readable seal seed refused at load (D17).

One deliberate reading recorded: the D9 page inventory counts
*rendered pages* — zitadel's `/authorize/callback` attaches an HTML
content-type to its 302 bodies, and a redirect is not a page
[judgment]. `make check` green: fmt, tidy, build, test, lint at zero
issues.

What it opened: M2 (passkeys — the login form is the stub it
replaces), M4 (the callout admission proof now has a real issuer to
prove against), M5 (the serve assembly is shaped for the embed lift:
value-only Options, `Open`/`Run` split). The quickstart records the
operational honesty: the issuer host becomes a one-way door at first
enrollment (D14), and the seal seed is backed up separately or the
store is honestly lost (D17).

Reversal condition: none — records a completed build against
already-graduated designs; the designs' own reversal conditions
(episodes 0043, 0046, 0051) stay armed.

Trail: `specs/001-op-skeleton/` (spec, plan, tasks, quickstart);
designs [store-and-key-lifecycle](../02-DESIGN/soulstream-idp/store-and-key-lifecycle.md)
(D1–D8, D16–D19) and [session-and-ui](../02-DESIGN/soulstream-idp/session-and-ui.md)
(D9–D15); the `001-op-skeleton` branch merge in the soulfold repo.
