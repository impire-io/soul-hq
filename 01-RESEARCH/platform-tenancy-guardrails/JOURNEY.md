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
