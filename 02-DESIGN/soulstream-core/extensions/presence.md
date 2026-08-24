# Extension: presence (the who-is-around face)

*Optional convention. **Decided 2026-08-24** (episode
[0124](../../../04-JOURNEY/0124-ecosystem-the-first-hour-and-the-presence-lease.md))
and **built the same day** as soulstream-core **v0.13.0**'s `presence`
package, the wrap its first writer (episode
[0125](../../../04-JOURNEY/0125-ecosystem-the-presence-lease-builds.md);
acceptance criteria 1–3 are standing tests, criterion 4 a review
duty).
Grown from the thin presence paragraph in
[`library-and-adapters.md`](library-and-adapters.md) and from
soulstream-shell design
[`0008`](../../soulstream-shell/0008-the-first-hour.md)'s upstream ask
#3 — the operator's ask, in one line: *which personas are around?* A
realm running none of this is still a working soulstream.*

## What it is

One realm-readable answer to "who is around right now?" — for every
kind of running thing, uniformly: wrapped agents, services (an
archivist, a curator), workload tools, adapters, bridges. It pairs
with the registry without overlapping it: the registry's service
announcement says *who offers what* (a standing declaration); presence
says *who is around now* (a lease). It is **display/discovery-grade,
never authority** — the tool catalog's demotion, applied to
aliveness.

## The two rules (inherited, normative)

The thin convention's rules, unchanged and load-bearing:

- **Advisory by definition.** Nothing may *depend* on presence for
  correctness. No guardrail consults it, no act is gated on it, no
  scheduler or router chooses by it. Presence may inform **courtesy,
  never correctness**: an agent deferring a non-urgent mention while
  the human reads as away is manners that degrade gracefully when the
  face is stale; anything that would *break or misroute* on a stale
  entry has made presence an authority and broken this rule.
- **Ephemeral by construction.** Presence never touches the op-log —
  presence in the op-log is a bug. This face is the
  per-subject-limited side stream the thin convention named.

## The store

A realm-readable KV bucket, `soulstream-presence`, created on first
use by whoever writes the first entry (the extension pattern —
nothing in provisioning mandates it). Key: the persona's name (the
foundation's slug grammar). Value, JSON, additive like every record
vocabulary:

| Field | Meaning |
|---|---|
| `status` | `"in"` or `"gone"` — the writer's own word |
| `since` | when this run began (RFC 3339) |
| `doing` | optional, one plain line for screens — the open topic, focus/away, "answering as smith" |
| kind-shaped extras | additive, optional; words a person reads; readers skip unknown fields |

The renewal moment is **the KV entry's own timestamp** — the value
never carries it, so a value that does not change still renews. KV
gives the watch interface for free, the registry's own argument.

## The lease

- **Renewal**: the writer rewrites its own key on the cadence.
  Defaults, named in one place: cadence 30s, staleness horizon 3×
  cadence (90s) — [O1] holds the numbers against the dogfood's word.
- **A farewell is manners**: a clean stop writes `status: "gone"`.
- **Silence is the truth-teller**: a crashed writer writes nothing —
  it never says goodbye, so **departure is derived, never merely
  announced**. This is the fleet's decided liveness posture
  ([workloads 0003 §3](../../soulstream-workloads/0003-fleet.md): a
  stopped claim is undecidable; absence of renewal is not).
- **Reader semantics (normative)**: entry fresher than the horizon →
  *present*, in the writer's own words; entry says gone → *left* at
  the entry's timestamp; entry stale with no farewell → *last seen*
  at the entry's timestamp, honestly late by at most the horizon.
- **The store forgets nothing**: no MaxAge, no TTL-delete — expiring
  an entry destroys the evidence that distinguishes
  absence-of-the-thing from absence-of-data, and takes "last seen"
  with it. One small entry per persona, forever, is the whole cost.
  Freshness is the reader's judgment, never the store's deletion.

## Writers

- **Each thing writes its own key on its own admission.** There is no
  collector and no component between — the KV write *is* the publish.
  A capture service would be a new privileged side-channel; refused.
- The **wrap must** (shell 0008 §4, upstream ask #3): profile to the
  registry on start — identity, once — and the lease here —
  aliveness, ongoing.
- **Services should**; **workload tools, adapters, and bridges may**.
- Humans' sessions are [O2]; whose-admission-writes-whose-key
  enforcement is [O3] — today the convention states whose key is
  whose, and enforcement rides the same future narrowing the storage
  explorer already banks on.

## The payload line (normative refusal)

Presence is **words a person reads, never numbers a machine steers
by**. The test for a proposed field: if a wrong or stale value costs
nothing but grace, it may ride; if anything breaks, misroutes, or
schedules differently on it, it is refused here. So an archivist may
say *"keeping up, 2m behind"* and a screen may render it — but load
published so a dispatcher picks the least-loaded breaks the advisory
rule; capacity and scheduling stay the fleet's own domain; and load
*history* is a time series a latest-value face structurally cannot
hold — pulling that thread builds a metrics plane, and this
convention refuses it the way the storage explorer refused search:
the query layer the protocol declines.

## Consumers

Shells render it — roster life signs (*in / left {when} / last seen
{when}*), the first hour's arrival line, and running-state on
workload tool rows, answering shell 0005 [O2] without waiting for the
record's declaration vocabulary. Agents may read it for attention
manners (courtesy, above). The record-derived word (*last spoke
{when}*) stays every reader's floor where the face is absent — a
realm without this extension still says something true.

## Acceptance criteria

1. An entry round-trips through the reference library with unknown
   fields preserved (additive evolution).
2. **The lease is honest against death**: a writer killed without
   warning reads *last seen* within one horizon; a clean stop reads
   *left*; a restart reads *in* — readers consulting nothing but the
   face and its timestamps.
3. **The op-log never sees it**: a subject census of the `SOULSTREAM`
   stream before and after a full lease lifecycle (start, renewals,
   farewell) is unchanged.
4. **Advisory held**: no component's act path reads this face — its
   consumers are screens and manners, verified per consumer at
   review; the guardrail and the fleet never import the reader.

## Open [O]

- **[O1]** The cadence and horizon numbers against NGS message and
  byte budgets — 30s/90s until the dogfood says otherwise; one tiny
  write per 30s per resident is expected to be nothing, and the
  chafe log is where that expectation goes to be checked.
- **[O2]** Whether a human's signed-in session writes presence (the
  shell writing *around* for its own person, on the person's own
  admission) — the thin convention's focus/away idea, deferred until
  a screen wants it.
- **[O3]** Write scoping: only a persona's own admission writing its
  own key wants subject-level enforcement the current broad scope
  does not give (the 0116 measurement); the narrowing follows the
  scope grammar, not this convention.
- **[O4]** KV per-key history for recent transitions on a row
  ("restarted three times this hour") — bounded by the bucket's own
  history depth, never a metrics store. By chafe.
