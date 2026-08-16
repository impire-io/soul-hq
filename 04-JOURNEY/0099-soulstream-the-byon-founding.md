# Episode 0099 — First contact: a realm founded on a live Synadia Cloud BYON (2026-08-16)

**The question:** spec 010's one unmeasured half — does the
synadia-cloud flavour found a real realm on a real BYON? It does now
[measured]: realm `byon` stands on the Impire DEV system's
agent-connected BYON (nats-server 2.12.7, the machine from episode
0038), founded by the released binary plus the fix stack first contact
forced. The callout **ADMITTED the founding persona on Synadia's own
infrastructure** and refused a garbage token with the audited generic
refusal; `up` served every plane against the substrate; `init` re-run
reported the verified no-op (13 artifacts — smaller than self-hosted's
15 by exactly the two artifacts this flavour never holds); the custody
audit passed (no master material, no PAT anywhere on disk); and the
after-state diff against the recorded before-state showed the system's
two pre-existing callout configs **byte-identical**, one new config,
two new accounts — additive, as design 0003 §2 demands.

**First contact broke things, and every break became a fix with a
replaying test, landed on main the same hour** [all measured live]:

1. **The default state dir is the client's home** — the product
   collided with the soulstream client tooling's config dir (context,
   persona keys) and misread it as a damaged ceremony. Now a named
   refusal pointing at `--state`, writing nothing (`7c1f030`).
2. **The awaiting state must resume** — the synadia flavour saves
   before the platform returns any seed, and Load demanded seeds that
   could not exist yet; a founding interrupted mid-account-half could
   never re-load its own state (`35bfa51`).
3. **A once-returned seed must hit disk the moment it arrives** — a
   mid-run failure orphaned the scoped group, its seed dead with the
   process; recovery took a hand-driven disable-then-delete (the
   platform refuses deleting an active group). The driver now hands
   each seed to `OnSeed` — persisted before the next API call — and
   the incident is a test (`7e68e6b`).
4. **The channel to an agent-connected BYON is lossy** — beno1's
   private-link idle watchdog cycles the tunnel while requests are in
   flight (journalctl: "No tunnel requests received in 9m53s" while
   closing fourteen live tunnel inboxes), losing ~50% of mutations to
   500 "nats: timeout". Every driver mutation now retries bounded
   through 5xx, list-first so a create that landed despite its timeout
   is found, never doubled (`7e68e6b`). The watchdog itself is
   Synadia's bug to fix — evidence handed over.
5. **The platform's callout surface is anti-idempotent** — re-enabling
   callout and re-adding a target or user draw a persistent 500 "an
   unexpected error occurred", so a resumed founding died on steps
   that had already succeeded. Enable and wiring are now read-first,
   and the stub replays the 500s so add-first can never come back
   (`dd7c53a`, `823a97b`).

Two platform facts banked besides: fresh accounts arrive with an
auto-created "Default" on-demand group, and a system carries multiple
callout configs side by side (three coexist on DEV) — the 0095-era
worry that enabling ours could rewire the old wiring is dead.

**The caveat became a measurement:** `sealed_requests=false` — the
platform set no callout xkey, so the callout runs unsealed on this
deployment class, printed at founding exactly as design 0003 §6's
as-built note promised. Accepted for the dev-class BYON.

Process note, honestly: two commits this arc (`823a97b` and one in the
0098 arc) went out with a red gate behind piped exit codes before
being fixed forward — the gate is only a gate when the exit code is
read; chained pipes stop at the commit boundary now.

Reversal condition: the unsealed callout is accepted for the dev-class
BYON — a deployment that needs sealed callout requests on Synadia
Cloud (observable: a production founding refusing to proceed unsealed,
or the platform exposing xkey seed export) reopens the sealing seam.
The retry budget holds while one founding completes within a few
resumes; foundings routinely needing more (observable: the run record
of the next live founding) widen the budget or add an outer loop.

Trail: soulstream `specs/010-byo-nats/quickstart.md` (the run record);
commits `7c1f030`, `35bfa51`, `7e68e6b`, `dd7c53a`, `823a97b`,
`e5a0c7e`, `9b3aaa0`; the realm at `~/.soulstream-byon` on the DEV
BYON; design [`0003-byo-nats.md`](../02-DESIGN/soulstream/0003-byo-nats.md);
episodes [0038](0038-soulstream-remote-mcp-node.md) (the machine and
the first wiring), [0095](0095-soulstream-byo-nats-designed.md)/[0096](0096-soulstream-byo-nats-ships.md)/[0097](0097-soulstream-v0-13-0-rc-1.md)
(the arc). The six fixes await v0.13.0-rc.2 — still gated on the
`HOMEBREW_TAP_TOKEN` secret (episode 0098's pending act).
