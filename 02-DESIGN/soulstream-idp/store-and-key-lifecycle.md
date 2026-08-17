# The store and the key lifecycle

**Graduated from research:** kv-schema-and-key-lifecycle, 2026-08-02 —
[episode 0002](../../04-JOURNEY/0043-soulfold-kv-schema-and-key-lifecycle.md);
extended by kv-encryption-at-rest (the envelope, D16–D19; D6 amended),
2026-08-02 —
[episode 0051](../../04-JOURNEY/0051-soulfold-kv-encryption-at-rest.md).
**Realized by:** M1 (the OP skeleton) on the
[roadmap](../../03-IMPLEMENTATION/ROADMAP.md).

What soulstream-idp keeps, where, and how it changes: the JetStream KV layout
that is the fold's only store, and the signing-key lifecycle that lets
JWKS roll over while consumers keep verifying, uninterrupted. Every
mechanism below passed a pre-registered bar in the graduating research
[measured]; the acceptance criteria at the end are those bars restated
as M1 gate tests. The storage contract this design serves is
`zitadel/oidc` v3's `op.Storage` (constitution III: the certified
library defines the protocol's demands; the fold supplies storage and
lifecycle).

## Decisions

### D1 — One KV bucket per record kind

Four buckets — `users`, `clients`, `keys`, `sessions` — each carrying a
configurable name prefix, default `soulstream_idp_`.

Reasoning: retention is configured per bucket, and the kinds genuinely
differ — sessions age out and want per-key TTL garbage collection
(`LimitMarkerTTL` set, see D5); users, clients, and keys are durable and
keep a short history for operator forensics. The prefix exists because
the fold may share a JetStream domain with its parent deployment; its
buckets must not squat on generic names. Survived Bar 1: all four
buckets found by lookup after restart, 6/6 records byte-identical
[measured].

### D2 — Records are JSON with a `schema` field; evolution is additive only

Every record is a JSON object whose first concern is `schema` (integer,
starting at 1). Evolution may only add optional fields. Renaming,
removing, or retyping a field is a breaking change and therefore a
stated migration under the roadmap's store-shape one-way door — never a
silent re-read.

Reasoning: JSON keeps records inspectable with stock NATS tooling and
needs no codegen (constitution III); Go's decoder ignores unknown
fields, which is exactly the additive property. Survived Bar 2: the
full v1↔v2 matrix over all four record kinds, 25/25 [measured].

### D3 — One writer version per deployment

A deployment runs a single writer schema-version at a time. A rolling
upgrade ships readers first, writers after; an old-version writer must
never read-modify-write a newer record.

Reasoning: measured, not speculative — a v1 reader that RMWs a v2
record silently drops the v2-only fields (demonstrated on the user
record: email and credentials vanish) [measured]. Additive *reads* are
safe; cross-version *writes* are not. For the M1 single binary this is
free; it becomes an operational rule the moment two fold versions can
touch one store.

### D4 — Birth by `Create`, every transition by `Update(revision)`

New records are written with KV `Create` (create-only; a duplicate is
an error, never an overwrite). Every mutation is a compare-and-swap
`Update` against the revision read, retried with a fresh read on
rejection. The schema never requires atomicity across keys: every
record must be independently valid at every reachable state.

Contended paths, named: auth-code redemption (D6 — a single-use CAS
flip; the loser's rejection *is* the single-use guarantee), signing-key
state transitions and the active-pointer flip (D7), and index-key
maintenance (D6).

Reasoning: KV has no cross-key transaction, so the design must not
want one. Survived Bar 3: 8 writers × 1,000 CAS cycles landed at
exactly 8,000 with 36,961 rejections all observably retried (~6,400
accepted contended writes/s — orders beyond an IdP's write rate);
code redemption won exactly once in 100/100 races of 8 [measured].

### D5 — `expires_at` is authoritative; TTL is garbage collection

Every expiring record carries `expires_at`, checked on every read; a
record past it is treated as absent regardless of its presence in KV.
KV TTLs (bucket-level, or per-key via `KeyTTL` on create in the
`sessions` bucket) only reclaim storage, on a schedule with slack
beyond `expires_at`.

Reasoning: TTL granularity, marker semantics, and server clock behavior
must never decide token or code validity — the security boundary lives
in the record, the janitor lives in the bucket [mechanism-argument].

### D6 — Buckets, keys, and record shapes (schema 1)

| Bucket | Key | Record |
|---|---|---|
| `users` | `<user-id>` | user |
| `users` | `idx.username.<hex(sha256(lower(username)))[:32]>` | index → user id (**amended by D16's research**: a plaintext username in a KV key would hand a store scan what the envelope withholds — D12's digest rule extends to user-supplied names; exact-match lookup is all the fold needs) |
| `clients` | `<client-id>` | client |
| `keys` | `key.<kid>` | signing key |
| `keys` | `active` | pointer → kid |
| `sessions` | `<request-id>` | session (auth request / token record) |
| `sessions` | `code.<hex(sha256(code))[:32]>` | single-use code index → request id ([session-and-ui](session-and-ui.md) D12: bearer secrets never verbatim) |
| `sessions` | `bs_<session-id>` | browser session ([session-and-ui](session-and-ui.md) D11) |

Record shapes (fields additive-only per D2; timestamps RFC 3339 UTC):

- **user** — `schema`, `id`, `username`, `display_name?`, `status`
  (`active | disabled`), `created_at`, `updated_at`. M2 adds
  `credentials[]` (credential id, **public key**, sign count — public
  material and digests only, constitution I) and profile fields,
  additively.
- **client** — `schema`, `client_id`, `name`, `redirect_uris[]`,
  `public` (M1 clients are public-with-PKCE), `created_at`. A future
  confidential client stores a secret **digest**, never the secret.
- **signing key** — `schema`, `kid`, `alg`, `state` (D7),
  `private_pkcs8`, `public_jwk`, `created_at`, `not_before`,
  `last_signed_expiry?`. These are the fold's own keys — the one place
  the store holds private material; it is private *to the fold*, not a
  user credential, and makes the store exactly as sensitive as any
  IdP's key store [judgment].
- **index** (`idx.*`, `code.*`, `active`) — `schema`, the target id,
  and for `code.*`: `consumed` (the CAS flip) and `expires_at`.
- **session** — `schema`, `id`, `kind`
  (`auth_request | access_token | refresh_token`), `client_id`,
  `user_id?`, `scopes[]`, `redirect_uri?`, `state?`, `nonce?`,
  `pkce_challenge?`, `pkce_method?`, `response_type?`, `auth_time?`,
  `done?`, `consumed?`, `csrf?` (one-shot,
  [session-and-ui](session-and-ui.md) D13), `created_at`,
  `expires_at`. A `refresh_token` record lives under the token's
  digest key (`rt.<digest>`, D12 — the bearer verbatim nowhere
  server-side) and carries the CAS `consumed` flip itself: rotation
  redeems each token exactly once, the code guarantee at the token
  scale (built 2026-08-17, episode 0103 — `offline_access` mints one,
  redemption rotates, revocation deletes).
- **browser session** (`bs_*`) — `schema`, `id`, `subject`,
  `created_at`, `expires_at` ([session-and-ui](session-and-ui.md)
  D11).

`op.Storage` mapping, so implementation doesn't guess:
`AuthRequestByID` → `sessions/<request-id>`; `SaveAuthCode` →
`Create sessions/code.<code>`; `AuthRequestByCode` → redeem
`code.<code>` (CAS flip of `consumed`, loser fails — single use), then
load the request; `CreateAccessToken` → `Create sessions/<token-id>`
with `kind: access_token`; `SigningKey` → `keys/active` → `key.<kid>`;
`KeySet` → all `key.*` filtered by state (D7).

### D7 — Key lifecycle: `pending → active → retiring → retired`, two invariants

One signing key is active at a time, selected by the CAS-updated
`active` pointer. JWKS publishes every key in state `pending`,
`active`, or `retiring`; `retired` keys stay in the store (audit,
`retired` is terminal) but leave the published set. Two invariants
carry the no-restart guarantee:

- **I1 — publish before sign:** a key enters JWKS (`pending`) at least
  one verifier cache-lifetime before it may sign.
- **I2 — unpublish after expiry:** a key may move `retiring → retired`
  only after `last_signed_expiry` — the latest `exp` it ever signed —
  has passed.

Rotation is then: create `pending` → (lead time) → flip pointer, old
key → `retiring`, stamp its `last_signed_expiry` → (until that instant
passes) → `retired`. Every transition is a D4 CAS.

Reasoning: I1 covers verifiers that cache JWKS on a TTL
[mechanism-argument]; refetch-on-unknown-kid verifiers like go-oidc
absorb the switch instantly [measured]. Survived Bar 4: one
never-restarted go-oidc verifier, 466 verifications, 0 failures across
a full rotation, with fresh-keyset controls proving the straggler
verified from published JWKS while retiring and was rejected after
retirement [measured].

### D8 — Tokens sign RS256

The fold's signing keys are RSA (2048-bit minimum), `alg: RS256`.

Reasoning: the consumer seam's verifier of record — soulstream-identity's auth
callout (its D23) — is pinned issuer + JWKS discovery + RS256, and
Entra publishes RS256; indistinguishable-by-design (constitution II)
decides the default [mechanism-argument]. The lifecycle mechanics are
algorithm-independent — the graduating rig ran ES256 for key-generation
speed and nothing in D7 noticed [measured]. Reversal: if the seam's
verifier of record grows ES256 support and a deployment need for it
appears, this becomes a per-deployment algorithm choice; the record
shape already carries `alg` per key.

### D16 — Records are sealed app-layer with the deployment's xkey

Every record value in the four buckets is stored as the raw sealed
output of an NATS curve-key (xkey, X25519) `Seal`, self-addressed to
the deployment's seal key; `Open` on every read. The plaintext inside
remains the D2/D6 JSON — the envelope changes custody, not shape. KV
*keys* stay structural (ids, digests, `active`); no user-supplied
plaintext may appear in a key (D6 as amended).

Reasoning: the graduating research re-ran the store's three mechanic
rigs through the envelope — restart byte-identical, additive matrix
24/24 identical to the unsealed baseline, CAS exact at 8,000/8,000
with exactly-once redemption 100/100 [measured]. The decisive threat
is the NATS surface itself: D1 lets the fold share a JetStream domain
with its parent deployment, and an ordinary API-level reader saw full
plaintext under server-side filestore encryption but nothing under the
envelope [measured]. Cost, priced on the real flow: +44 bytes and
~57 µs per record operation, +1.19 ms on an end-to-end sign-in
[measured]. Sealing halves contended CAS throughput (~3,165/s vs
~7,116/s accepted writes) — orders beyond an IdP's write rate
[measured].

### D17 — The seal seed's custody story

- **Birth**: generated at first start (single binary: by the fold;
  embedded: by or under the parent's ceremony), written `0600`.
- **Home, per deployment shape**: always outside the JetStream store
  directory. Single binary: a file beside the config, never under the
  store dir; embedded (soulstream): the parent's state area, sibling to
  — never inside — its `jetstream/` dir; shared JetStream: the seed
  stays with the fold process, and the JetStream operator holds only
  ciphertext — this is the shape where custody bites hardest.
- **Re-keying**: a stated re-seal migration — walk every record, open
  with the old key, seal with the new, CAS-write; flip the config;
  destroy the old seed. Never a silent re-read (the store-shape door).
- **Loss**: total, honest data loss of record contents. Deployments
  back up the seed *separately* from the store; a full-machine backup
  that captures both is plaintext-equivalent and the docs must say so
  [judgment].

### D18 — Tokens and codes stay out of the store's trust

Unchanged by the envelope and restated so nobody relies on it: bearer
secrets never appear verbatim server-side (D12) and nothing the fold
stores may be sufficient to impersonate a user (constitution I). The
envelope defends *confidentiality* of records; it is not the
impersonation boundary and must never be argued as one
[mechanism-argument].

### D19 — Filestore encryption is defense-in-depth, not a substitute

JetStream filestore encryption (ChaCha) protects the disk artifact
alone: restart green, disk scan clean, API reader sees plaintext — all
demonstrated [measured]. Deployments may enable it *in addition to*
the envelope (single binary: worth offering; shared JetStream: the
parent's call), but it satisfies no bar the envelope satisfies. Threat
table, from the graduating rig:

| Threat | Filestore cipher | App-layer envelope (D16) |
|---|---|---|
| Stolen disk / store-dir backup | covered | covered |
| NATS-API-level reader (shared domain, ops tooling) | **not covered** [measured] | covered [measured] |
| Root on the live host | not covered (server key on host) | not covered (seed on host) — stated honestly |
| Fold process compromise | not covered | not covered |

## Acceptance criteria (the M1 gate inherits these)

1. Restart round-trip: the full working set survives a server restart
   byte-identical, buckets found by lookup, no re-seeding.
2. Additive matrix: the schema-N ↔ schema-N+1 decode matrix over every
   record shape runs green in `make test`.
3. Exactly-once redemption: racing redeemers of one auth code produce
   exactly one winner, every round.
4. Rollover: a stock, never-restarted OIDC verifier sees zero
   verification failures across a full key rotation, old-key tokens
   verify until expiry, and the retired key is absent from published
   JWKS.
5. The envelope (D16–D17): every stored value opens only with the
   deployment seal key; a marker scan over the stopped store dir and a
   full API-level dump finds no record plaintext (positive-control
   scan proven); the seed lives outside the store dir and its loss/
   re-keying story is in the deployment docs.
