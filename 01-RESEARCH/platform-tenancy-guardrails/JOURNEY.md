# Journey — platform tenancy and guardrails

*The investigation as it happens. Bars are in [`README.md`](README.md);
requirements in [`requirements.md`](requirements.md).*

---

## 2026-08-04 — the topic opens, with its direction already set

Unusually for a research topic, this one opens with eight load-bearing
decisions **already taken** (S1–S8 in
[`requirements.md`](requirements.md), recorded as
[episode 0064](../../04-JOURNEY/0064-ecosystem-the-platform-turn.md)). That is
worth flagging rather than hiding: the normal shape is question first,
decisions at graduation. Here a comparative analysis against a separate
proprietary platform — built by the same author, for the same problem space —
settled the direction before the topic existed, and the topic exists to
enumerate what the direction *requires* and to settle what remains open.

The honest consequence: **this topic's output is mostly not measured.** Five
bars are genuinely measurable and are pre-registered. Ten decisions are not,
and are listed as such in `README.md` under *Decisions that are not bars*
rather than dressed up as experiments. When the topic graduates, its verdict
must keep that separation — `[measured]` for the bars, `[judgment]` or
`[mechanism-argument]` for the rest. A verdict that blurs them would be the
failure mode this ecosystem's working agreement exists to prevent.

**What the comparison established** (recorded in the episode, not repeated
here): the two systems are the same architecture built for opposite halves of
the problem. The other has tenancy, guardrails, a model plane, an audit
derivation, and a Kubernetes deployment story; this one has the collaboration
substrate, the workload runtime, and the single-binary deployment the other
has an *unresolved research effort* about. Two capabilities are worth having
from that comparison — tenancy and guardrails — and both must be built here
from requirements, never ported.

**The provenance rule** was written before any requirement, and is at the head
of `requirements.md`. It is stricter than "no code": design-document prose,
tables, identifiers, and vocabulary are copyrightable expression too. The rule
survives this topic and binds anyone extending it.

**Three findings changed the scoping** before a single requirement was written:

1. **A realm is already one NATS account** — stated normatively in
   [`core/01-protocol.md`](../../02-DESIGN/soulstream-core/core/01-protocol.md).
   The expected "make realms into accounts" work turned out to be a rename
   (S1). The real gap is lifecycle, which does not exist at all
   `[mechanism-argument]`.
2. **The realm name is bound into the canonical signed record.** So the rename
   is not purely cosmetic at the wire, which is what makes A10 a one-way door
   rather than a naming preference `[mechanism-argument]`.
3. **The persona registry already covers the "passport" requirement** —
   `operated_by` with a countersigned operator attestation, chains terminating
   at a principal, key distribution, service advertisement
   ([`extensions/registry.md`](../../02-DESIGN/soulstream-core/extensions/registry.md)).
   What was thought to be a build turned out to be two narrow gaps: status and
   richer capabilities `[mechanism-argument]`.

**One correctness gap surfaced and is not yet confirmed** (F1): a persona's
signing key materializes in the vault on first use, and nothing publishes the
public half to the registry — so a persona could sign records nobody can
verify. Inferred from the design docs, **not traced through running code**.
Confirm before acting; if confirmed it is the cheapest item in the topic and
should go first.

**Open at the moment of writing:** every `[OPEN]` marker in
`requirements.md`, and the latency budget for Bar 3, which cannot be fixed
until the enforcement point is chosen (B8). That budget goes here, with its
number, before Bar 3 runs — not into the bar retroactively.

---

## 2026-08-18 — F1 confirmed by code trace, and it is the default, not a corner

The trace the opening entry demanded ran against the code as shipped
[measured: code trace]. Both halves of the gap are real:

- **Signing side.** Three real consumers construct `PersonaSigner` —
  the product's runner (`soulstream/node/workload.go:65`) and archivist
  (`soulstream/node/node.go:454`), and the remote node's per-user pool
  (`soulstream-mcp/pool.go:455`, every door persona). The persona key
  materializes in the vault on first touch
  (`soulstream-identity/internal/service/service.go` `ownedPersonaKey`
  → `GeneratePersonaKey`), owner-bound — and **no profile publication
  exists on any of these paths** (product, mcp, shell, and archivist
  repos all swept: the shell only reads the directory, the rig
  publishes profiles only in test setup, the archivist references the
  registry nowhere).
- **Reading side.** Every reader builds its keyring from **registry
  profiles + TOFU pins** and nothing else: core's `realmKeyring`
  (`soulstream-core/internal/cli/connect.go:75`, backing CLI, stream,
  memory, discover, inbox) and the node's `keyringFor`
  (`soulstream-mcp/pool.go:483`). The identity plane's `keys.public`
  directory (D26) is consulted by **no reader anywhere** — it is real
  and proven in identity's own e2e, but no core-side consumer wires it.

The consequence is stronger than the opening entry guessed: this is not
a persona that *could* sign unverifiably — it is the **shipped default
for every identity-plane-signed persona**. The product's own runner and
archivist, and every person admitted through the remote door, sign ops
that every reader in the realm renders `unknown-key`, unless someone
manually published a profile for them (which the dogfood personas did,
masking the gap).

The fix stays what S7/S8 dictate — a **consumer wiring job**: the party
constructing a `PersonaSigner` owns publishing (or refreshing) the
persona's registry profile with the signing public key beside it. A
reader-side fallback to `keys.public` is refused by the cycle guard
(core cannot import the identity client). Owner and shape land with
this topic's graduation outputs; the natural home is one core helper
(`registry` gains an ensure-signing-key act) called by the three
consumers above at signer construction.

---

## 2026-08-18 — the grants broker lands elsewhere, and half of §C/§D lands with it

The `outbound-identity-grants` topic (episode 0104) graduated overnight
and its build merged this morning (episode 0105, soulstream-identity
v0.3.0). That work was scoped to *outbound* credentials — but it built
the very mechanisms §C and §D of this topic enumerate, for the
outbound-resource class. Reconciled here so this topic measures what
remains rather than re-planning what exists.

**Bar 5 — PASS** [measured 2026-08-18]: after the full grant build
(identity v0.3.0, delegations wired against soulstream's signing
shape), both core modules' dependency graphs contain zero references to
the other — `go.mod` and the complete `go mod graph`, both directions,
all zero. The consumer-position e2e module imports both, which is the
cycle guard's sanctioned shape (episode 0027). The grant work S8
worried about is now real code, and the guard held.

**§D — the custody pattern is proven; the general surface is not
built.** D31's second custody domain measures every mechanism §D
names, for its domain: CAS conditional writes (D2) under `-race`;
paths namespaced by construction (`grant/<persona>/<resource>`) with
reach structural via the transport op tail (D3/D4's property); sealed
at rest under the deployment's first key, positive-control-verified
(D5); the E2E-sealed surface (D6/D7); and act-with-without-receiving
(D8) — the caller gets derived access tokens, the refresh token never
crosses [all measured, identity's suite]. What §D still lacks is the
*general* secret store: D1's arbitrary caller-named paths do not exist
— the broker custodies one record shape it defines itself. **D9 gains a
data point, not a closure**: the broker put a second custody domain
*inside* the one service (own bucket, same first key, same surface
discipline) and it composes cleanly — evidence for one-service, but the
question as posed (general secrets beside key custody) is still open.

**§C — the delegation half exists; the standing-grant half does not.**
For the outbound-resource action class: C1 (authorize without
acting-as), C3 (dual attribution, each readable independently — every
on-behalf decision audits both personas), C5 (enforcement with the
custodian), C6 (nothing pre-provisioned) are running, tested code
[measured]. C8 is **answered for persona-to-persona delegation**
[judgment, D33]: the grant *is* a presentable bearer artifact —
subject-signed, verified by the enforcer from the D26 directory, no new
trust root — honored only from its named actor, which removes the
leak-amplification that made C8's bearer arm scary. What remains: C2's
independent revocation (today a delegation ends by expiry; revoking the
subject's grant custody disturbs the subject's own access — D33's
standing **consent record**, revocable and refusing the next mint, is
designed but deliberately unbuilt), and C4 entirely — the op-log
recording half (S8's soulstream side): grant issuance, revocation, and
exercise are visible today only in the identity plane's audit log, not
in any topic the granter can watch. C7's argument sharpens: the
refresh-token store is unambiguously an original fact (a secret with no
other home); the *consent record* is the artifact that must still answer
the dissolved-registry objection when C4 is designed.

**Bar 4 — partially measurable, not yet a PASS.** The real machinery
measures the granted action performed and the dual attribution [both
measured]. The bar's third clause — revocation stops the action while
disturbing *neither* persona's own standing — is exactly the unbuilt
consent record: today's only revocation lever (custody deletion) kills
the subject's own access too. Bar 4 waits on C2/C4's remaining half,
and the delegation matrix is its measurement rig when that lands.

---

## 2026-08-18 — pre-registration before the Bar 1/2/3 spikes

Written before either rig runs, per the topic's own discipline.

**Bar 3's enforcement point and budget, fixed now.** Point: the
custodian's op path (the B8 memo's recommendation — the chokepoint a
caller cannot bypass because the capability lives there). Budget:
**added p99 ≤ 2ms** on a representative allow path with a rule set of
~100 compiled rules, measured as evaluator-in-the-loop minus
evaluator-absent on the same op shape. Hostile-rule pass criteria as
the bar states: refused at compile/load, or terminated by the
evaluator's own cost bound — the caller never waits unboundedly. The
candidate rule language is CEL (public prior art: sandboxed,
non-Turing-complete, per-evaluation cost limits) — B7's three
properties by construction, to be verified, not assumed. If B8's
eventual recording moves the enforcement point, the bar re-runs there;
this run stands as the evaluator's own measurement either way.

**Bar 1/2's rig shape.** An embedded operator-mode server (the
embedgate ceremony idiom, memory account resolver — the same resolver
class the product's founding uses); a service account exporting a
request/reply surface with `account_token_position` forcing each
importer's own account key into the subject (D14/A6's mechanics, a
stand-in responder behind them); account A pre-existing with a
continuous round-trip probe for the full duration. At T: account B's
JWT is constructed **complete** — limits, the import, everything —
and only then stored into the resolver, the A2 ordering bet being that
no intermediate state ever exists on the wire. Probes: before the
store, B's user connect must fail closed (unknown account) and any
partial success fails Bar 2; after, B's user connects and completes a
round trip on the shared surface. Bar 1 passes only with zero
restarts, zero edits to A's configuration, and A's probe uninterrupted
throughout. Stated honestly up front: the shared surface is a stand-in
responder riding the real export/import JWTs — the identity plane
rides the same mechanics by configuration (episode 0011's D14
amendment) [mechanism-argument], and wiring the full plane into this
rig is the graduation-time upgrade if the stand-in leaves doubt. The
**provider arm** (A8: Synadia Cloud BYON control-plane API) is not
runnable from this desk without the operator's portal token; it is
named as the remaining Bar 1 arm, not silently dropped.

---

## 2026-08-18 — Bars 1, 2, and 3 measured: three passes and one honest scare

Both rigs built and run this session (scratchpad modules `tenancy-rig`
and `guardrail-rig`, per how-we-work). All numbers are spreads over
repeated runs, not means.

**Bar 1 — PASS (local arm)** [measured, 6/6 runs]: account B, its JWT
built complete and stored as the one act, went from store to **first
full round trip in 543µs–774µs** across runs — zero server restarts,
zero edits to account A, and A's continuous probe (230–238 round trips
per run at 5ms cadence) recorded **zero failures** with a max
inter-success gap of 6.01ms. The server resolves a runtime-stored
account on demand; nothing existing is touched. The provider arm (A8)
remains named and unrun.

**Bar 2 — PASS** [measured, 6/6 runs]: 2,418–2,826 pre-creation probes
per run — connect attempts with valid-but-unknown-account credentials,
each also attempting the shared surface — **all failed closed, zero
partial successes**; post-creation, the first probe was already a full
success (connect + round trip) in every run, with zero
connected-but-unreachable states observed. The A2 ordering bet held:
build the account artifact complete, then store — the store is the
atom, and no intermediate state ever exists on the wire.

**Bar 3 — PASS, and the bar earned its keep** [measured, 3 runs]: the
hostile set against CEL: unparseable and type-broken rules **refused at
compile**; a nested-comprehension cost bomb **terminated in
689µs–982µs**; a 100k²-element input bomb **terminated in 484µs–913µs**;
the catastrophic-backtracking regex probe ran linear in 14µs (RE2 by
construction). Allow path at 100 compiled rules per op over 10k ops:
**p50 58–69µs, p99 206–220µs, max ≤ 492µs** — an order of magnitude
under the pre-registered 2ms budget. **The scare, recorded as a design
output:** the first run used a cost limit alone (500k) and the input
bomb, though terminated, took **622ms** to die — cost accounting is a
step bound, not a wall-clock bound. The discipline that passes, and
that the eventual build must carry: a tight cost limit sized to op-path
rules (10k), an interrupt check every ~100 steps, and a **context
deadline (25ms) as the hard stop** — B7 is realizable, but only as
belt-and-braces, never as a single mechanism.
