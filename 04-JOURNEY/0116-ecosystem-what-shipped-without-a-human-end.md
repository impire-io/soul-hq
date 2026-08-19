# Episode 0116 — What shipped without a human end (2026-08-19)

The operator, living on the release candidate as the focus gate asks,
named three gaps in one message: the status screen measures storage but
cannot show a message; adding external tools like MCP servers has no
home; and boundaries/guardrails have no surface at all. The survey that
followed found the second and third were not missing work — they were
**finished work with no human end**.

**The record has no viewer.** The Storage card renders `%d ops · %s`
and a meter against the roof the store declares; nothing anywhere in
the shell shows an op [measured, code trace]. When a turn does not
appear or a verdict surprises somebody, the only answer today is
`nats stream view` at a terminal with operator credentials, outside the
surface that reported the problem. Everything a viewer needs is already
public on the pinned core — `realm.Client.JetStream()`, `record.Parse`,
`Record.Canonical` — so this is a shell design, not a research question
[measured, code trace]. Design
[`0004-the-storage-explorer.md`](../02-DESIGN/soulstream-shell/0004-the-storage-explorer.md)
takes it, with one decision worth the argument: the explorer reads on
the **signed-in person's own admission**, not the surface's shared read
lane, because `SOULSTREAM.>` includes one person's notifications and
will include sealed payloads. The honest half of that decision is that
it narrows nothing today — the product's persona scope grants every
admitted persona `SOULSTREAM.>` and `$JS.API.>`, so anyone signed in
can already read the whole op-log [measured, code trace:
`soulstream/ceremony/ceremony.go`, `scopePubAllow`] — so the screen says
what your admission can read and never implies scoping it does not
have. The payoff is in the future tense: narrow the scope, or land
sealed topics, and the explorer follows with no surface change.

**The outbound broker is off in the house.** soulstream-identity v0.9.0
ships the whole grants surface and the product declares no
`GrantResources`, which is what enables the `grants.*` ops — so the
lane that exists to refuse "one credential in the agent's config" is
not reachable from anything a person touches [measured, code trace].
The module that would fix it is the easy half; the undecided half is
where the tool catalog lives, given that
[`grants.md`](../02-DESIGN/soulstream-identity/grants.md) states the
catalog is declared configuration with no per-user rows (D26's spirit),
and how a running agent gets a token at call time without one landing
in an `.mcp.json`. Opened as research topic `agent-external-tools`,
four pre-registered bars.

**The guardrail is built and headless.** Episode
[0110](0110-identity-the-tenancy-set-builds.md)'s D37/D38 build is
complete and default-off; the op dispatch serves `guardrail.load` and
`approvals.present` and nothing else; the client package exposes
neither; there is no read op for the standing rules and no feed of
decisions [measured, code trace]. D38's loop — defer emits, the human's
yes mints, the retry presents — has no home for its middle clause, which
is why nothing is guarding anything today. Opened as research topic
`guardrail-human-end`, five pre-registered bars, one of them the
non-obvious one: the natural carrier for "tell a person" is the
record's deliberately-general `PERSONA.NOTIFY` subject, and using it
would make the identity plane depend on the record — exactly what the
cycle guard (episode [0027](0027-soulstream-dx-hardening-and-the-cycle-guard.md))
and 0107's Bar 5 keep at zero.

**The two topics stay separate**, though they meet: D37 named three
chokepoints and built one, and the third — the door's tool-call
forwarding — is precisely where "which external tools may this agent
call, with what arguments" gets decided. Coupling them into one topic
would let the slower half hold the faster hostage; the guardrail
surface is buildable against shipped code, while the catalog needs a
decision that may reverse a stated design property. The seam is named
in both topic journeys so graduation does not forget it [judgment].

What this taught, beyond the three items: the focus gate is working as
written. Episode [0071](0071-ecosystem-the-focus.md) froze capability
growth behind real demand and the roadmap's next gate is "the operator
tries the whole system — evaluation decides what changes." This episode
is that evaluation returning its first verdict, and all three items
arrived as *use*, not as a plan.

Reversal condition: for the storage explorer, design 0004 §2 carries
its own — an operator who cannot debug a real failure because their own
admission cannot read what broke (observable: a recurring failure class
whose evidence sits on subjects the operator's persona is refused)
brings back a deliberately named operator lane as a numbered decision,
never a quiet fallback to the shared read lane. For the two topics,
their pre-registered reversal conditions stand. For this episode's own
direction — that surfaces are what these three gaps need — the
observable that would change it is the guardrail topic's Bar 1 finding
real deferrals expiring at TTL before any human reaches a browser, at
which point the answer for that half is pre-authorizing policy rather
than a screen.

Trail: design
[`soulstream-shell/0004-the-storage-explorer.md`](../02-DESIGN/soulstream-shell/0004-the-storage-explorer.md)
(new) and its index entry; research topics
[`agent-external-tools`](../01-RESEARCH/agent-external-tools/README.md)
and [`guardrail-human-end`](../01-RESEARCH/guardrail-human-end/README.md)
(pre-registered, unstarted); roadmap sections for the shell and the
identity plane; episodes
[0104](0104-ecosystem-outbound-identity-grants.md)/[0105](0105-identity-the-grants-broker-lands.md)
(the broker this finds unreachable),
[0107](0107-ecosystem-platform-tenancy-guardrails.md)/[0110](0110-identity-the-tenancy-set-builds.md)
(the guardrail this finds headless),
[0071](0071-ecosystem-the-focus.md) (the gate that produced the ask).
