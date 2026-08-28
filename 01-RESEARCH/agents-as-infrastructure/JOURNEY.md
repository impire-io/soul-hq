# agents-as-infrastructure — investigation journal (started 2026-08-27)

## 2026-08-28 — Bar 1 measured: PASS, both arms, first day

**Hypothesis:** submit-and-forget needs no new machinery — `fleet.Submit`
already puts a declaration on the record as an ordinary work item, the
wrap engine already serves declared wakes, and the log alone should let a
restarted dispatcher resume. The discriminating experiment is the
composition nobody has run: a dispatcher that claims placements and runs
the wrap engine per claimed agent, with the submitter gone and the
dispatcher dying mid-story.

**Rig:** consumer-position Go module in the session scratchpad
(`aai-bar1`), pinning `soulstream-core v0.14.0-rc.1` +
`soulstream-workloads v0.8.0-rc.2` — public surfaces only. Embedded
JetStream server, hermetic realm. The dispatcher prototype is ~90 lines;
its whole loop:

- poll the placement topic's materialisation;
- **resume arm**: item `claimed` with `owner == self` → serve it (no new
  op, no handshake — the log is the position);
- **race arm**: item `open` → `ClaimWork` → re-materialise → serve only
  if the read-back names self owner (fleet.TryPlace's discipline, minus
  its `Runner.Launch` coupling);
- serve = connect a client as the declared persona, `wrap.DeclaredConfig`
  over it, `wrap.Wrapper.Run` (mock invoker).

**Arm 1 — submitter gone, dispatcher restarts between triggers.**
Submit → submitter's connection closed before the dispatcher ever sees
the item. Measured over 3 runs [measured]:

- submit → claimed+served: 159.7–162.1ms (dominated by the spike's 150ms
  poll interval — the mechanism itself is ~10ms of that);
- mention answered by the dispatcher-served agent: 54.8–55.1ms;
- dispatcher stopped whole (engines + clients), second mention posted
  into the silence (outcome count provably unchanged while down), fresh
  dispatcher instance with the same claim persona and zero local state:
  `resume_from_log` found its claim, the engine's catch-up read
  `inbox=2`, logged `wake_already_answered` for the served trigger, and
  answered the missed one in 162.9–214.5ms from restart;
- exactly 2 outcomes, exactly 2 harness invocations, no duplicate op
  ids; replay from a fresh auditor client reconstructs the whole story —
  one submission, exactly **one live claim** (resume never re-claimed),
  two attributed outcomes — from ops alone.

**Arm 2 — hard kill between trigger and outcome.** First attempt
produced a finding, not a pass: a *graceful* stop is not a crash — the
cancelled invoker returned `OK:false` and the engine dutifully posted
the failure **self-report** (`kind=self_reported`) before dying. That is
the engine's contract (failures are the agent's own testimony), and it
draws a line the dispatcher design must keep: **drain and crash are
different ends** — a drained dispatcher reports, a killed one leaves the
record clean for its successor. Re-run as a true crash (connections
dropped mid-invocation, invoker never returning): zero outcomes while
dead [measured], restart re-served the wake — invocation count 2
(at-least-once), outcome count 1 (exactly-once), outcome landing 208ms
after restart; the dead process's half-finished post visibly dies on the
closed connection and lands nothing [measured]. The deterministic
outcome id (`WakeOpID`) is the whole mechanism.

**Verdict on the bar:** PASS — 3 consecutive `-race` runs of both arms
green. No store beside the log was needed; the reversal condition's
second reading (submit-and-forget demands a dispatcher-owned store) did
not fire.

**What it opened:**
- The resume arm is the piece that exists nowhere shipped: fleet's
  `TryPlace` hardwires `Runner.Launch` (backend workloads), so a real
  dispatcher needs either a fleet seam that takes a serve function or
  its own claim path — a design question for graduation, not a blocker.
- The spike polls; a real dispatcher should watch the placement topic's
  subject live with poll as catch-up (the wrap engine's own pattern).
- Liveness (Bar 2): this dispatcher never answers probes; the crash arm
  here leaves the item claimed by a ghost until the successor instance
  resumes it. With TWO nodes, the fleet's probe/sweep decides instead —
  that is exactly Bar 2's territory, next.
- Credential custody is faked (hermetic realm, no auth): the dispatcher
  connects "as" the declared persona freely. Bar 4 owns the honest
  version.

## 2026-08-28 — Bar 2 measured: PASS, same day

**Hypothesis:** the 0003 claim path needs nothing new to decide which
dispatcher node serves an agent — and the reclaim discipline (projection
nominates, probe vetoes, ordinary abandon decides) carries the *serving*
handoff too, with the wake-dedup mechanism guarding the takeover the way
it guards a restart.

**Rig:** the Bar 1 spike's dispatcher grown the fleet half, all shipped
surfaces: it answers `fleet.ProbeSubject(self)` for the items it serves,
and runs `fleet.Node.Sweep` (the shipped reclaim — Sweep needs no
Runner) every 300ms with a 1s reclaim bound. Two nodes, four agents
each with its own desk topic, all submissions contested (both nodes up
before the first submit), submitter gone.

**Measured, 3 consecutive `-race` runs + the first [measured]:**

- contested placement: split node-a=2 / node-b=2 every run, every item
  exactly one live claim, zero double-launches;
- every agent answered its mention wherever it landed (4/4, one
  invocation each);
- live owners never reclaimed: sweeps ran continuously from startup
  (every claim older than the 1s bound within a few polls), zero
  abandons on any live item across all runs;
- failover: node-a hard-killed owning 2 agents, a mention posted into
  the window — answered by the survivor 1.051–1.064s after the kill
  (reclaim bound 1s + sweep cadence; the bound is the knob), exactly
  one outcome (no double serve), every reclaimed item's timeline
  exactly `claim,abandon,claim`, node-b's own items undisturbed;
- the survivor's takeover re-serves the reclaimed agents and
  `wake_already_answered` drops their already-served triggers — **the
  restart dedup and the failover dedup are the same mechanism**, which
  is the finding that makes the dispatcher's serve arm boring (good);
- census: zero probe traffic on the stream [measured, the fleet gate's
  clause re-proven through the dispatcher].

**Verdict on the bar:** PASS. No new vocabulary, no coordinator; the
one composition the shipped fleet lacked was "serve via the wrap engine
instead of Runner.Launch" — which is a seam question, not a mechanism
question.

**Still open for the build design (not bars):** a node that dies and
RESTARTS between the death and a peer's reclaim resumes its own claims
legitimately (it answers probes again) — correct by construction,
worth a standing test at build time. The spike's poll should become a
live subscription with poll as catch-up.
