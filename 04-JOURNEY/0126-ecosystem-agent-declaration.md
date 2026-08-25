# Episode 0126 — Agent declaration: the record declares, the room runs (2026-08-23 → 2026-08-25)

The question, pre-registered with four bars before any experiment ran:
does "declare an agent from instructions — who I am, what I can do,
what I need to act on, and deploy, all from the shell" decompose
entirely into existing vocabulary, with no new machinery beside the
log, no core wire change, and no human/machine branching? A scratch
rig (session scratchpad, embedded servers, real component libraries)
answered all four bars **PASS [measured]**, race-checked:

- **Wake is vocabulary** — one declaration, four wake sources
  (mention / topic op / schedule tick / external subject); every wake
  left exactly one outcome op under a deterministic UUIDv5 id (wrap's
  own shape, generalized); a restart mid-backlog answered nothing
  twice, verified at both the stream and the post-attempt level
  (5.43s; `-race -count=3` 18.6s). Delivery class is per source kind
  and must be declared honestly: the core-subject wake fired while the
  dispatcher was down was *lost* — at-most-once, measured — while
  stream-backed kinds replay exactly.
- **Instructions are an artefact of the record** — the registration
  references a stage-1 artefact, never a host path; the tip is
  materialised per wake (digest-checked, structurally uncached), so a
  revision reached the *running* agent with no redeploy, and a
  dispatcher death lost scratch only, never history (0.93s).
- **Capability is the identity plane, unchanged** — the identity plane
  ran in-process (D29 embed seam); D28 `mint.ephemeral` minted against
  the declared role name with the declaration's tools as tags; a
  `SOULSTREAM.SVC.{{tag(tool)}}` scoped template — the ecosystem's
  first `{{tag()}}` template in Go code, the fleet design's open item
  — enforced at the transport: granted tool answered, ungranted
  refused with zero responder deliveries and zero authorization code
  in the rig (1.64s).
- **The shell verb is composition** — the full declare flow drove
  through shipped surfaces only (attestation token → self-published
  "attested" profile; instructions and the registration itself as
  artefacts; placement as `work.open`; the dispatcher booting from the
  registration read back out of the record; `open:owner → claim:sprite
  → done:sprite`). The only missing piece is the verb that sequences
  the steps (0.41s; suite `-race -count=2` 18.9s).

Refuted or reversed: the owner's teach-back carried one claim that did
not survive as a finding — heartbeats-with-TTL — and it dissolved into
the same-day presence thread on main, which decided both points better
(no TTL: freshness is the reader's judgment and "last seen" evidence
survives; advisory, never authority). This topic's passing suggestion
that a heartbeat KV could serve as the fleet's transient evidence was
withdrawn for the same reason [judgment]. The adversarial pass refuted
"wake vocabulary belongs to the inhabitant" (wake determines credential
scope and placement, which the inhabitant must not self-declare
[mechanism-argument]) — while two of its arguments *stand as design
obligations*: instruction revision is privilege escalation, so the
soul topic is a guarded surface; and declared agents put colonies one
op away, so mention-wake agents ship before topic-wake colonies.

What it taught and opened: growth confined to schema and provisioning,
never machinery — a `wake` section and a record-form artifact scheme
in the workloads declaration (the roadmap's named hole), one additive
`SOULSTREAM_SYSTEM` stream (owner's call: streams consolidate under
one system home; KV faces stay bucket-per-face) with the shadow-record
drift guard written now, and the enforcement-read gap `[O]`: minted
agent credentials cannot re-derive record position under operator-mode
scopes (only `$JS.API.INFO` is granted). The reconciler's shape is
answered by house grammar — a persona with ordinary credentials, never
a privileged tier — with one hand-started process per node as the
bootstrap fixed point. Named successors stand: agent-wakes-agent loop
safety before any colony, runtime join/leave (watched, unfired), tag
policy (the identity plane's own item). Design
[`0005-agent-declaration.md`](../02-DESIGN/soulstream-workloads/0005-agent-declaration.md)
carries the functional spec; declared agents inherit profile-on-start
and the presence lease by being the same kind of persona the wrap
already is.

Reversal condition: a bar's mechanism failing under the shipped
implementation — the declaration needing a coordination store beside
the op-log (observable: a mutation that never appears as an op or
realm KV write), or wake exactly-once breaking under a real backend
(observable: an outcome id counted twice on the stream) — reopens the
topic; a code path that must branch on human vs machine to pass review
abandons the direction rather than patching it.

Trail: research topic `01-RESEARCH/agent-declaration/` (pre-registered
bars, per-bar journal with raw numbers, teach-back and adversarial
pass; folder removed at graduation, full history in git); branch
`research/agent-declaration` (b968572 → 40d6ee4, merged baca4c3);
design [`0005-agent-declaration.md`](../02-DESIGN/soulstream-workloads/0005-agent-declaration.md);
context [`extensions/presence.md`](../02-DESIGN/soulstream-core/extensions/presence.md),
[`0004-wrap.md`](../02-DESIGN/soulstream-workloads/0004-wrap.md) §10,
[`extensions/work.md`](../02-DESIGN/soulstream-core/extensions/work.md) stage 4.
