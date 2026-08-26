# Rigs — platform-account-topology

Discriminating experiments for the bars. Each is a self-contained Go
module; `cd <bar> && go run .`. Preserved in-topic so the trail is
reproducible; removed on graduation, retained in git history per
how-we-work.

- `bar1/` — cross-account export preserves the D15 principal proof
  (account_token_position enforcement, with a negative control).
- `bar2/` — one-act tenant birth and admission: the AUTH allowed_accounts
  coupling and the scoped-mint-on-plain-key admission. Fully standalone —
  reproduces the identity code's JWT shapes verbatim (cited in the source),
  no soulstream imports.
- `bar3/` — isolation through shared services: real provisioned SOULSTREAM
  streams in two tenants, a shared two-connection service, adversarial probes.
  Uses a local `replace` for `soulstream-core` (imports the real
  `realm.ProvisionOn`); adjust the path in `go.mod` to run elsewhere.
- `bar4/` — the multi-tenant human: token lane (live), OIDC ambiguity
  (roleFor reproduced + fuzzed), persona-name scope (collision reproduced).
