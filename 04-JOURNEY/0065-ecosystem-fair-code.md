# Episode 0065 — The ecosystem goes fair-code (2026-08-08)

A licensing decision, taken deliberately: every impire-io project moves
from MIT to the **Sustainable Use License** (n8n's fair-code license,
version 1.0, derived from the Elastic License 2.0 — n8n explicitly
invites adoption). The terms in one sentence: use, modify, and
self-host freely for internal or non-commercial purposes; offering the
software to others as a paid product or service requires an agreement.

- **Why.** The ecosystem's bet is that the substrate is the product and
  people run it themselves — that stays free, unchanged. What MIT also
  allowed was a third party packaging the same substrate as a hosted
  commercial service without any relationship to the project. Fair-code
  keeps the self-hosting door wide open while making that path a
  conversation.
- **What changed.** `LICENSE` replaced (SUL v1.0 verbatim, with a
  provenance preamble) and the README license section rewritten in
  soulstream, soulrealm, soulidentity, soulfold, soulnode,
  soulstream-archivist, poseres (pra), imps, and soul-hq itself. The
  public explanation lives at
  [impire.io/license](https://impire.io/license/) (shipping with the
  website rework in the same drive).
- **What did not change.** Every version already published under MIT
  stays MIT — a granted license is irrevocable; the new terms apply
  from the next release of each component. Sole authorship was
  verified per repo (`git shortlog -sne`: one author, two email
  identities) before relicensing, so no third-party consent was
  needed.
- **Known cost, accepted.** SUL is not OSI-approved, so pkg.go.dev
  stops rendering docs for the Go modules and the module proxy may
  decline to cache them (`go get` falls back to `direct` and keeps
  working). If protocol-library adoption friction ever bites,
  carving the client libraries back out to MIT is a one-file revert
  per repo.

Reversal condition: if fair-code measurably suppresses the adoption
the ecosystem exists for — clients not being built, realms not being
stood up — the libraries return to MIT first, the products only if
that isn't enough.
