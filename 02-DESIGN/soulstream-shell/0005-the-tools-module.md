# 0005 — soulstream-shell: the tools module

**Status:** designed with its build — decided 2026-08-21 (episodes
[0118](../../04-JOURNEY/0118-ecosystem-agent-external-tools.md)/[0120](../../04-JOURNEY/0120-ecosystem-the-tools-arc-builds.md)
carry the design and build this screen is the human end of). This is
[`external-tools.md`](../soulstream-identity/external-tools.md)'s [O4]:
the `resources.*` admin surface and the per-person linking ceremony —
the browser half the broker always needed.

## §1 The gap

Everything under this screen already runs (episode 0120): an operator
can `resources.add` a remote MCP server by client call, a person can
link their account by client call, and the agent's door forwards with
the person's own identity at the remote. What does not exist is the
screen: the demand episode 0116 opened with — *a person adds a tool
from the shell* — still ends at a terminal.

## §2 The decided surface

One screen, `/tools`, a key on the spine for every session — the
catalog is everyone's to read and linking is everyone's own act.

- **The list**: the catalog's entries (the record's discovery face)
  merged by name with the plane's resources — one row per tool: name,
  kind in plain words ("connected service" / "runs here"), description,
  and for remote tools **this person's own standing**: connected or
  not, read from their session's own grants list.
- **Connect / Disconnect** (per person, per remote tool): Connect
  starts the linking ceremony as the person and sends their browser to
  the provider's own sign-in; the provider returns them to the
  module's callback route, which completes the ceremony as the same
  session and lands them back on `/tools` with their standing updated.
  Disconnect revokes their own grant. Plain register: people connect
  accounts; "grant" and "link" stay in the machine room.
- **Add / remove a tool** (admin): one form, both kinds. A remote tool
  takes the catalog half (name, endpoint, description) and the
  ceremony half (auth/token URLs, client id + secret, scopes) — **the
  writer writes both halves in one act** (D39), plane first, catalog
  second, and says which half failed when one does. A workload tool
  takes name, persona, endpoint, description — catalog only. Remove
  reverses both halves; standing grants keep their custody, said on
  the screen (the plane's own semantic).

## §3 Custody and lanes (normative)

- **Reads**: the catalog rides the shared read lane (it is the realm's
  public shape, the board's own class); the plane's resource list and
  the per-person grants list ride, respectively, the support layer's
  management lane and the person's own admission.
- **Linking is the person's own act on their own prefix** — session
  identity client, never the shared lane (D30's isolation property is
  the transport's; this surface keeps it). The callback carries the
  ceremony's identity in OAuth `state` (= the broker's link id, its
  published contract) and completes against the session that started
  it: a callback arriving on someone else's session completes nothing.
- **Admin acts ride the node-standing lane** the deployment hands this
  surface — the agents module's precedent, restated: the credential
  ops are refused to a person's own admission by design, and no
  side-channel is grown. The shell gates the acts on the session's
  admin role — a display-plus-gate fact; the deeper authority is the
  guardrail standing at the plane's own op path (D37), where "who may
  add a tool" is evaluated regardless of what any surface believes.
- **The secret crosses once**: the add form's client secret goes to
  the plane and nowhere else — not into the catalog, not into any
  page the shell serves back, not into the session. The e2e's custody
  scan carries it as a needle.

## §4 Acceptance criteria

1. The standing gate walks the ceremony against a real deployment at
   published tags: admin adds a remote tool (stand-in AS) → both
   halves exist (plane resource + catalog entry) → a person connects
   (authorize redirect out, callback in, standing shows connected) →
   the person's grant serves an access → disconnect refuses the next
   one → admin removes the tool → both halves gone, the screen honest.
2. The secret needle appears nowhere the scan reads: no page served,
   no catalog entry, no shell state — control fired.
3. A non-admin session is offered no add/remove and its posted act
   refuses; connecting stays offered to everyone.
4. Plain register throughout; zero horizontal overflow at 1000/390 px.

## §5 Open [O]

- **[O1]** The linked-status read is per person per render; a realm
  with many tools may want the support layer caching it. By chafe.
- **[O2]** Workload rows are catalog-only today; when workloads'
  declaration vocabulary lands on the record, the row could offer
  running-state. A new decision then.
