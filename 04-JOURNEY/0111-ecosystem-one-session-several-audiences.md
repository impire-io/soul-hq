# Episode 0111 — One session, several audiences: the exchange grant lands on both ends (2026-08-19)

The second of episode 0104's due demands closed, and lane 3 with it.
The fold (**v0.8.0**, spec `006-token-exchange`) serves RFC 8693
through the certified library's `TokenExchangeStorage` seam — audience
re-scoping only: same subject, roles carried, scopes narrow never
widen, and the exchanged token's expiry bounded by the subject token's
remaining life, so the v0.7.0 token-lifetime knob's revocation bound
survives exchange. Actor tokens refuse by name (delegation lives with
the identity plane); only fold-issued subject tokens exchange.
Discovery says only what is true: the exchange methods live on an
opt-in storage wrapper handed over exactly when `--exchange-audiences`
declares targets — support upstream is interface-derived, so honesty
had to be structural. Measured in the e2e [measured]: the full passkey
ceremony → exchange with swapped `aud`, same `sub`, carried roles,
bounded `exp`; four refusal rows (undeclared audience, actor token,
foreign subject token, widened scope); discovery on both arms.

Two contract facts found by measurement, recorded: the library's
exchange path authenticates the client by **Basic only** (the form's
`client_id` is not read), so a public client presents its id with an
empty secret — registration is its authentication, and a presented
secret still always refuses; and public-client exchange required
amending the fold's blanket `AuthorizeClientIDSecret` refusal to accept
exactly that shape, nothing more.

The identity plane's broker (**v0.7.0**) gained the **lane 3 backend**
(D34): a resource declaring `exchange_token_url` + `exchange_audience`
is served by exchanging the caller's OWN bearer — presented, never
retained — against the wire contract the fold's e2e just measured from
the RP side. No linking ceremony, nothing at rest (asserted by the
store staying empty), on-behalf refused by name: there is no custody to
redeem for a subject. Lane 3 is now what D34 preferred it to be —
same surface, no custody.

One process finding, recorded because it is the honest kind: `v0.8.0`
was tagged, released, then **retagged** one commit later after a lint
fix (a local variable shadowing a builtin). The published binaries
therefore came from the pre-fix commit while the tag pointed past it,
and the second release run failed on the already-existing release. The
source difference was a rename only, and the module pin resolved
correctly either way — but a tag whose artifacts disagree with it is
not a tag. Fixed by cutting **v0.8.1** from the settled commit so the
binaries match their tag [measured: release green, 5 assets]. The
lesson for the night's own practice: verify a gate's exit status
before tagging, never a piped tail — three of tonight's runs hid a
non-zero exit behind `| tail`.

Reversal condition: none — records completed builds against a
published RFC. (Spec 006's own condition stands: a consumer proving a
need the actor-token refusal cannot meet reopens that scope, argued
against D33 first.)

Trail: idp `4963042`+`4ba1786` (v0.8.0, spec
`specs/006-token-exchange/`), identity `39b75d6` (v0.7.0); episodes
[0104](0104-ecosystem-outbound-identity-grants.md) (the demand measured
from both ends), [0106](0106-idp-the-token-lifetime-knob.md) (the knob
whose bound survives exchange).
