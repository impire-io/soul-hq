# 0006 — soulstream-shell: the approvals module

**Status:** designed with its build — decided 2026-08-21 (episodes
[0119](../../04-JOURNEY/0119-ecosystem-guardrail-human-end.md)/[0121](../../04-JOURNEY/0121-ecosystem-the-approvals-loop-closes.md)
carry the design and build this screen is the human end of). This is
[`approvals.md`](../soulstream-identity/approvals.md)'s [O3]: where a
human sees "X wants to do Y" and their tap closes D38's loop.

## §1 The gap

The loop runs end to end by client call (episode 0121): a deferral
opens a ticket, a human's yes mints and presents, the retry serves.
But "a human" is still a terminal session. The topic that produced all
of this opened on exactly this absence — no inbox, no screen, nobody
told.

## §2 The decided surface

One screen, `/approvals`, its key on the spine drawn only for sessions
that could act — the admin role, or a persona named in some standing
rule's approvers clause — with the pending count as the key's mark:
the spine is where a person learns something waits for them.

- **The list**: every ticket awaiting a decision — who asked (the
  principal, by display name where the directory has one), what they
  asked to do (the action), under which rule, and how long the window
  has left. **Never the arguments**: the ticket's own privacy line
  (D42), said on the screen so the absence reads as designed rather
  than broken — "the request is named by its fingerprint; the rule is
  your context."
- **Approve / Deny**, one tap each: the session **mints** the answer
  as the person's own signed act (their persona key, their signature —
  an approval is a D33 delegation and this module never pretends
  otherwise), and the surface **delivers** it on the originator's tail
  through the node-standing lane. The plane's refusals come back in
  its words — an approver outside the rule's clause reads the by-name
  refusal, not a shell invention.
- **Resolved honestly**: an act that raced the clock ("expired"), a
  second answer ("denied"/"approved" already) — each in the plane's
  own words, and the list re-reads so the screen shows what is now
  true.
- **Empty state**: "Nothing is waiting for a decision."

## §3 Mint and delivery, kept apart (normative)

The one subtlety this screen must not blur (D44): the *yes is the
person's* — minted on their session, signed by their key, verified by
the plane against the directory and the rule's clause — and the
*delivery is the surface's* — `approvals.present` must arrive on the
originator's own tail, and the originator is an agent or service that
is not in this browser, so the surface carries the signed artifact
there on the node-standing lane. The shell signs nothing as anyone:
carrying a sealed envelope is not authority, and the plane's actor
binding means a delivered artifact converts only the invocation its
signer named. Delegated authority, never borrowed identity — with
delivery named as the third thing it is.

## §4 Activation (normative)

A declared fact, never a probe (design 0002): the deployment says its
identity plane runs the guardrail (`GuardrailOn` through the embed
options; soulstream declares true since v0.13.0-rc.9's enable). Absent,
the module is nowhere — no key, no routes, 404.

## §5 Acceptance criteria

1. The standing gate closes the loop through the screen against a real
   deployment: a deferred op's ticket appears on the screen → the
   admin-or-approver session approves → the originator's retry serves
   → the ticket reads spent. The deny arm likewise, ending denied.
2. The key is drawn for the admin and for a named approver, not for a
   plain session; the routes refuse a plain session's posted act in
   the plane's words.
3. No ticket's arguments appear anywhere — there is nothing to leak by
   construction, asserted over the served pages anyway.
4. Plain register; zero horizontal overflow at 1000/390 px.

## §6 Open [O]

- **[O1]** The pending count on the spine reads the plane per render
  (the conversations tally's own precedent). A busy deployment may
  want the support layer debouncing it. By chafe.
- **[O2]** Argument visibility for approving humans stays the plane's
  line ([`approvals.md`](../soulstream-identity/approvals.md) [O3]):
  if rule names prove insufficient context in practice, the answer is
  richer *rule* vocabulary, not arguments on tickets.
- **[O3]** Notifying an approver who is not looking at the shell (a
  mention in a conversation? mail?) is composition's next reach —
  deliberately not this module's.
