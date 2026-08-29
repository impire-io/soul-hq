# 0010 — soulstream-shell: the models surface

**Status:** drafted 2026-08-28, at the operator's direction — the
thinking house's operator window. Closes the [O] design
[0009](0009-the-declare-surface.md) §4 named ("writing the catalogue
from the shell is a named [O], not a refusal") and answers the line
the byon soak recorded the same week
([`DOGFOOD.md`](../../03-IMPLEMENTATION/DOGFOOD.md)): the inference
plane undeclared, models a CLI-act-in-words empty state. Grounded in
the measured facts of spec 014 ([episode
0147](../../04-JOURNEY/0147-soulstream-the-thinking-house.md)) and the
plane's founding design
([inference 0001](../soulstream-inference/0001-the-inference-plane.md));
evidence classes as there.

## §1 The surface, in one sentence

The spine grows a **Models** sheet — the catalogue's names as a table
everyone reads, authoring and re-pointing as a slide-over form whose
output is the identical entry `soulstream model set` writes (one
codec, no second schema), removal behind a one-word ask that counts
the declared agents naming the name, and a **Serving now** reading
drawn live from the plane's own discovery — with provider credentials
staying exactly where wall 3 put them: outside the shell, behind a
paste-able command.

## §2 The sheet (reading)

- **Model names**: one row per catalogue entry — name, capability,
  and where it points in the person's words (the pinned model, the
  tag words, or "any instance that serves it"). Each entry unfolds
  "as JSON" in a stow, the whole stored truth — including fields the
  form does not write. Reconstructed per render; no store (0001 §4,
  standing).
- **The picker's source is unchanged.** Design 0009's models picker
  keeps reading names only; this sheet reads entries whole. That
  amends the stated names-only policy at its comment
  (`soulstream/declare.go`) precisely, not broadly: reading a route
  the surface cannot use is reading somebody else's configuration —
  a surface that *manages* the catalogue is the party whose
  configuration it is.
- **Serving now**: the plane's instances from one discovery scatter
  per render (`client.Resolve`, the house's own 150ms window) —
  model, capability, tags, formats; raw instance ids ride hovers
  (design 0007). A serving instance IS the provider evidence: the
  plane refuses to start on an unresolvable secret, so this reading
  derives provider standing without touching any custody.
- **Words by the declared fact.** A new config-derived fact,
  `InferenceOn` (§5), tells the empty states apart honestly: fact
  absent — "nothing serves models here yet; the deployment turns on
  model serving" with the CLI act in words, no spinner; fact present
  with nothing discovered — honest waiting. An empty catalogue offers
  the form and the paste-able CLI line (0008's rule).
- **Live** at the agents cadence: one stream re-rendering the names
  and serving elements wholesale; the result line belongs to the acts
  and is never written by the stream.

## §3 The form (authoring)

- **Fields**: Name (the record's name grammar, refused upstream in
  its own words); Capability (default `chat`, free text — a later
  capability is a new name by the plane's own design); **Points at**
  — any instance (anycast), a pinned model (picker over the distinct
  models Serving now discovered, plus free text: a name may honestly
  point at a model nobody serves yet, and no-responders tells the
  truth), or tags folded under "more ways to match" (exact key:value,
  all must match). The form would be the ecosystem's first hand for
  tags — the CLI writes only capability and pin today; the matching
  mechanism is shipped and measured.
- **`default_params` is not offered.** Stored upstream but applied
  nowhere on the door path (`Descriptor.Apply` has no caller) — an
  editor for a knob that does nothing would lie. The stow shows what
  exists; editing waits on upstream ask #2 (§6).
- **Set IS re-point**: the same act pre-filled from the row. The
  screen says in words what spec 014 measured — the catalogue is read
  fresh per resolution, so a re-point moves the very next request; no
  restart, no redeploy.
- **Remove** stands behind the ask (key: Remove), and the sentence
  counts: when the placements topic is declared, "N declared agents
  name this — their next serve refuses until a name points again"
  (the dispatcher refuses an unpointed name whole); when the shell
  cannot count, it says so rather than guessing.
- **Validation is upstream's, one codec**: the entry the form writes
  is byte-identical to the CLI's for the same inputs, via ask #1. No
  shell-side schema duplicate exists anywhere.

## §4 Lanes and custody (normative)

- **Reads ride the shared read lane** — names, entries, and discovery
  alike (the catalog class of design 0005 §3). `$SRV.INFO` is in no
  persona scope by the plane's own design (0001 §2: resolving and
  pinning are infrastructure's acts), and the shared lane is the
  infrastructure standing this surface already holds.
- **Writes are the person's own act on their own admission.** The
  canonical persona scope carries `$KV.>` — spec 014's load-bearing
  asymmetry (persona wide, agent narrow) — so a session's put rides
  the transport's own permission, mirroring `fleet.Submit`: the
  surface acts as nobody. A deployment whose founding predates the
  canonical tails owes the scope-patch duty the dogfood run has paid
  once already — a named wall, not a surprise.
- **The admin gate is drawn, and named for what it is.** The
  transport grants every persona the same write; the shell draws the
  acts for admin sessions and refuses a non-admin's posted act, and
  the screen must not imply this narrows what the realm actually
  permits (the storage module's honesty rule, applied to a write). A
  real chokepoint is upstream's [O] (§7).
- **No secret through the shell** — wall 3 re-carried: provider keys
  arrive by the paste-able `soulstream provider set` block, the value
  riding its environment variable as a placeholder the person fills;
  the shell reads no custody tree and serves no key material.
- **No door-key surface.** Per-serve keys have no read surface
  upstream, deliberately; the sheet shows nothing it would have to
  hold.
- **No lifecycle act.** Instances are the deployment's configuration;
  the surface must not invent a stop it cannot perform (0009's
  no-retirement rule, restated for processes).

## §5 Boundaries, each already a house rule

- **One new dependency, named openly**: soulstream-inference's
  published client surface (`client`, `wire`) — the purity allowlist
  grows a bar with its reached-ness control, the exact pattern the
  workloads bar set. No `internal/` import anywhere.
- **One new declared fact**: `InferenceOn`, from the deployment's
  config, absent-means-absent — ordering-safe because the helm plane
  starts before the inference plane, so the fact is the config's, the
  pattern `declaredPlacements` already rides.
- **Plain words on screen**: "Models", "Serving now" — catalogue,
  plane, anycast, pin stay in the machine room; the standing
  banned-word gates extend, not bend.
- **No store**: every list a reading per render; nothing dismissible,
  nothing cached.

## §6 Upstream asks

1. **(inference)** The catalogue's entry encoding becomes the plane's
   published contract: the JSON shape
   `{capability, model_pin?, tags?, default_params?}`, the bucket
   name, and get/set/list over a handed KV handle — with the house's
   `node/catalogue.go` rewired to consume it and the shell's spelled
   bucket constant retired (the drift episode
   [0148](../../04-JOURNEY/0148-shell-the-declare-surface.md) flagged,
   closed at its source). Provisioning stays the product's: the
   plane's census keeps creating nothing.
2. **(soulstream, by demand)** A `default_params` consumer on the
   door's route path — before any surface offers to edit them.

## §7 Acceptance criteria

1. The e2e walks the sheet on a real realm at published tags: add a
   name from the form → the row and the declare picker show it from
   the same reading, the stow's JSON byte-equal to the stored entry;
   re-point → the next read shows the moved pin; remove behind the
   ask → gone everywhere; a bad name or missing capability refuses in
   upstream's own words.
2. Codec equality asserted, not assumed: the sheet's entry and
   `soulstream model set`'s entry for the same inputs are
   byte-identical.
3. Serving now shows a configured stand-in instance's model and
   capability from one discovery scatter; with the fact absent the
   sheet says the honest words and offers the CLI act — no spinner,
   no empty table implying a fault.
4. A non-admin session is offered no act and its posted act refuses;
   no served copy implies the gate is the realm's own permission.
5. The provider needle: no key value appears anywhere the shell
   serves or keeps — the paste block carries a placeholder; control
   fired.
6. Purity gates green and grown: the inference allowlist bar with its
   control; frame purity, no module importing another, banned words,
   the 360px floor, every table inside its wrapper — all standing.
7. Names and serving re-render wholesale on one stream at the agents
   cadence; the result line is never written by live.

## §8 Open, named [O]

- **The write authority**: today the persona scope's own `$KV.>` is
  the whole permission — a catalogue-write chokepoint (guardrail-class
  authority, or per-bucket narrowing at founding) is upstream's
  decision when a realm wants it narrower than everyone.
- **`default_params` editing** — gated on ask #2.
- **Metering** (usage per key, per persona) — gated on inference 0001
  §6's own [O]: where metering lands is undecided, and no surface
  precedes its home.
- **Catalogue watching** (0001 §9) — per-render reads stay until
  resolution cost measures real.
- **Serving detail past `Resolve`'s decode** (endpoints, queue
  groups, meaningful versions — the instance version is a hardcoded
  constant upstream) and **door-key visibility** (which serves hold
  keys — no read surface exists) — by demand.
- **Instance lifecycle from the screen** — waits for a vocabulary, as
  retirement does in 0009.
