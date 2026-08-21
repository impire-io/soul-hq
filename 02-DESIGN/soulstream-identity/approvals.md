# soulstream-identity — approvals: the ticket, and the loop's human end (D42–D45)

*Graduated from research topic `guardrail-human-end` ([episode
0119](../../04-JOURNEY/0119-ecosystem-guardrail-human-end.md)). The
problem: D37/D38 shipped whole and headless — the defer emits and the
approval mints, and the loop cannot close, because presenting is
unreachable (no client method, sealed envelope internal) and no human
ever learns anyone is waiting. This document designs the missing ends.
It amends D38 in place where noted and continues the global numbering:
D42–D45. The measured ground: the emit half works today (rule +
deterministic, argument-sensitive invocation id in the refusal), the
mint half is public (an approval is a D33 delegation naming
`invocation:<id>`), and a correctly minted approval leaves the retry
still deferring [measured, 6/6].*

**D38 correction, textual**: the invocation is named by
`hash(principal, action, argument bytes)` — **time is not in the
hash**, and must not be: a retry at a later moment could never match a
time-salted name. The implementation always did this; the design's
sketch said otherwise and is corrected here [measured].

## D42 — The ticket: a deferral is state, with a clock and witnesses

A defer today is a stateless refusal. It becomes a **ticket** the plane
custodies:

- **States**: `pending → approved → spent`, with `pending → denied`
  (the human's no) and `pending → expired` (the clock's no). Expiry is
  a **first-class, recorded outcome — never a silent death** (operator
  direction, 2026-08-21).
- **Two clocks, deliberately distinct**: the **ticket TTL** is the
  human's window (long, configurable per deployment; hours-scale); the
  **approval TTL** is the retry's window (minutes, one-shot, already
  built — D38 unchanged), and its clock starts at the yes, so human
  latency never eats the retry's window.
- **Async by construction, an invariant**: no call ever waits for a
  human. The deferred op returns immediately with a **structured
  refusal** — ticket id, rule, ticket expiry — machine-readable, not
  prose; **the retry is the execution**.
- **Durable**: tickets live in a sealed store (the plane's existing
  custody machinery), surviving plane restarts — at human-scale TTLs,
  in-memory loss would expire every pending ask unwitnessed, which the
  notified-expiry requirement forbids [judgment, from the operator's
  direction]. Approvals themselves stay in-memory and fail closed
  (D38 unchanged): they are minutes-scale and re-mintable while the
  ticket stands.

**Reversal condition** (the topic's, carried and softened): if real
tickets recurringly expire before any human sees them despite visible,
notified expiry (observable: the expired/answered ratio staying high on
a real deployment), the answer is pre-authorization policy — rules that
say yes in advance — with the one-shot ticket as the exception, not a
longer clock.

## D43 — The loop's public ends: present, and read

Two additions to the published surface, none to the mechanism:

| Surface | Shape | Notes |
|---|---|---|
| `client.PresentApproval` | `(invocationID, Delegation)` → error | the missing link: one public method over the existing `approvals.present` op |
| `approvals.status` | `{invocation_id}` → `{state, expires_at}` | on the principal's own op tail — the originator asks about its own tickets |
| `approvals.pending` | `{}` → tickets awaiting decision | the approver-side read, management-gated like `guardrail.load` until D45 refines who sees what |
| `guardrail.list` | `{}` → the standing rule set | the read half `guardrail.load` never had; a surface shows rules without keeping a copy |

The permission-template addition is the same shape `grants.>` was: the
represented user's tail grows `…{{name()}}.approvals.>` as deployment
duty.

## D44 — Notification is composition's, and it is correctness

The plane can neither write the record (the cycle guard — zero core
edges [measured]) nor push where anyone may listen (a persona's sub
permissions are `_INBOX.>` and `SOULSTREAM.>`) — so the plane **owns
the truth and sends nothing**:

- **The originating adapter carries the news**: the door tells its
  agent (a deferred tool call answers promptly as pending-with-ticket;
  a later status change reaches the agent by the door's watch or its
  next attempt), the shell tells its human (the approvals screen reads
  `approvals.pending` as the approver, `approvals.status` as the
  originator).
- Both legs of notification are named in the operator's direction:
  **"a human is now required"** to the originator at defer time
  (satisfied structurally by the D42 refusal), and **"your ticket
  expired"** at expiry (satisfied by status reads + the adapters'
  watch; a durable inbox notification is the shell's to post as
  composition, never the plane's).
- **Promptness is an acceptance criterion, not politeness**: the
  approval TTL starts at the yes, so "approved — retry now" must reach
  the originator inside that window or the one-shot is unusable. The
  gate measures defer → yes → retry-served end to end with the
  notification path in the loop.

## D45 — Who may approve is policy, per rule

Today any directory-resolvable persona's signature approves. The rule
gains an optional `approvers` clause — persona names, or the sign-in
surface's group names as its tokens already carry them — checked at
presentation beside the existing actor/binding/expiry checks. Absent,
the current behaviour stands, stated: any directory-resolvable persona.
A well-formed approval from a persona outside the clause **refuses by
name** — the demonstration Bar 5 could not run until presenting existed
becomes this decision's acceptance criterion.

## Acceptance criteria (the bars as the build's gate)

1. The full loop closes in consumer position: defer (structured
   refusal) → pending visible to the approver → yes (mint + present in
   one act) → retry served; the same approval refuses a second use;
   an unanswered ticket expires at its TTL with the expiry readable
   and the adapters' notification observed [Bar 1's shape].
2. The artifact's wire and at-rest arms of the leak scan run at last:
   planted argument absent everywhere, positive control fired [Bar 3's
   completion].
3. `guardrail.list`/`approvals.pending` serve a surface that holds no
   copy between requests; another principal's status is refused by the
   server, not the surface [Bar 4's shape].
4. The D45 refusal: an approval from outside a rule's approvers clause
   refuses by name [Bar 5's shape].
5. The cycle guard re-measured at zero after the build [Bar 2's
   standing half].

## Open [O]

- **[O1] The approver-side scope** of `approvals.pending` once D45
  lands: per-rule approvers seeing only their rules' tickets is the
  natural refinement; decided when D45 builds.
- **[O2] Rate counters** for defer floods stay where 0110 left them:
  by demonstrated need.
