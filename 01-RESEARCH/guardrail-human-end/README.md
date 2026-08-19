# How does a deferred guardrail decision reach a human, get answered, and let the refused call proceed?

**Component:** ecosystem
**State:** active
**Started:** 2026-08-19

## Abstract

The guardrail is built and headless. soulstream-identity v0.5.0
(episode [0110](../../04-JOURNEY/0110-identity-the-tenancy-set-builds.md))
landed D37/D38 whole: CEL at the custodian chokepoint under the
mandated belt-and-braces, rules as data with `guardrail.load` as the
hot reload, three outcomes, every evaluation observable, and approvals
as one-shot subject-signed delegations — actor-bound, usable once, dead
in minutes. It is **off by default** (`-guardrail`), the product
enables it nowhere, and there is no way for a person to see a rule, see
a decision, or answer a defer [measured, code trace: the op dispatch
serves `guardrail.load` and `approvals.present` and nothing else; the
client package exposes neither].

So D38's loop is missing its human end. "The guardrail's defer emits
the request; the human's yes mints the delegation; the retry presents
it" describes a mechanism whose middle clause has no home. Until it
does, a deployment's only guardrail posture is allow-or-deny written by
an operator into a JSON file — which is why nothing is guarding
anything today.

## The question

How does a deferred decision reach a person, get answered by them, and
release the refused invocation — without inventing the second control
plane episode 0002 declined, without the shell keeping any rule or
decision state of its own, and without the approval artifact carrying
the arguments it binds?

## Pre-registered bars

- **Bar 1 — the round trip closes, end to end.** Protocol: the identity
  plane with the evaluator on and a rule whose effect is `defer`; a
  caller trips it; a person answers on a surface; the caller retries
  through `approvals.present`. **Pass:** the retry is served; the same
  approval refuses a second use; an unanswered defer's approval is dead
  at its TTL; a stolen approval refuses as an actor mismatch (the
  existing property, re-proven end to end). Wall-clock defer→served is
  recorded, not thresholded — this bar is about closure, not speed.
- **Bar 2 — the carrier does not break the cycle guard.** The obvious
  route for "tell a person something happened" is the record's
  `SOULSTREAM.PERSONA.NOTIFY.<persona-id>`, whose subject the protocol
  calls deliberately general. But the identity plane must not depend on
  the record: episode 0107's Bar 5 measured zero cross-references in
  both modules' dependency graphs, and 0027's cycle guard stands.
  Protocol: whatever carrier the candidate design picks, re-run the
  guard over both modules' complete graphs. **Pass:** zero
  cross-references. **Fail** rules the record out as the carrier and
  forces the request onto the identity plane's own surface — a numbered
  decision, taken openly.
- **Bar 3 — the artifact leaks nothing.** D38's own reversal condition,
  measured rather than argued. Protocol: defer an invocation whose
  arguments contain a planted secret; inspect the approval artifact on
  the wire and at rest. **Pass:** only `hash(principal, action,
  argument bytes)` appears, with the positive control proving the grep
  would have found the plant.
- **Bar 4 — the standing rules and the decision stream are readable
  without a second copy.** No read op exists today; the topic must add
  one or conclude that rules stay write-only. Protocol: a surface shows
  the running rule set and the recent decisions; a test asserts the
  surface holds no rule or decision state between requests; a second
  persona's request for the same read is refused by the server, not by
  the surface. **Pass:** all three.
- **Bar 5 — who may approve is answered, not inherited.** Today any
  directory-resolvable persona's signature approves; episode 0110 names
  per-rule approver policy as an open follow-up. **Pass:** the topic
  states the policy shape and demonstrates one refusal of a
  well-formed approval from a persona that should not have been able to
  give it.

## Reversal condition

The topic assumes the human end belongs on a screen. What would change
our minds, phrased as an observable: **if real deferrals expire at
their minutes-long TTL before any human sees them** — a person is not
reliably at a browser when an agent is working — then the answer is not
a surface at all. It is policy that pre-authorizes classes of
invocation in advance, and D38's one-shot delegation is the wrong
artifact for the common case rather than the right one nobody could
reach.

## Verdict

<Empty until graduation.>
