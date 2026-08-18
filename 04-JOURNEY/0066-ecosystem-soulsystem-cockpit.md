# Episode 0066 — The helm: the cockpit earns its design (2026-08-13)

The question — *where and as what does the human cockpit live, such
that it observes and configures the whole soulsystem while remaining a
pure consumer of public tagged surfaces?* — opened, ran its four
pre-registered bars, took eight decisions, and graduated to design in a
single day. The soulsystem gets its sixth component: **soulhelm — the
helm**, the human entry beside the MCP door.

**The design half already existed.** The topic opened on a discovery:
the Soulsystem Design System (a Claude Design project) carries tokens,
seventeen components, voice rules, and a Stream Console kit — cassette
futurism in a light key, with amber and teal as the human and machine
channels *at deliberately equal weight*: the vision's "humans and AI as
peers" turned into a color rule. It conflicted with the dark
violet/rose language soulfold shipped in episode
[0063](0063-soulfold-the-console-gets-a-face.md); the operator decided
the canon (C3): **cassette light, ecosystem-wide** `[judgment]`, with
per-component identity riding labels — mono strips and wordmarks —
never a second color system beside the channel pair
`[mechanism-argument]`. Costs accepted openly: a committed light-only
look, the fold's pages restyle later, violet/rose retires.

**The placement fell out of constitutions, not taste** (C1/C2). Growing
the fold's console would make the IdP Soulstream-aware — the exact
one-way door soulfold's constitution forbids; cockpit code native to
soulnode would be the first breach of "composition, not invention"
`[mechanism-argument]`. So the helm is a sibling component on the
twice-proven consumer-position pattern (the archivist, soulfold): its
own repo and constitution, a public `embed.Run` seam soulnode composes
as a plane by tag, a standalone binary for realms without soulnode.
`soulbridge` lost the name to soulstream's planned protocol bridges;
`soulconsole` to the fold's own console.

**The four bars, all PASS, all `[measured]`:**

- **Pure consumer, compiler-proven.** A module outside every
  component's namespace pins all six upstream artifacts by tag with
  zero `replace` directives, boots a whole realm in-process in ~1.5 s,
  and reads everything back: board, a turn at `sig=verified` with the
  keyring earned from `keys.public`, a claimed work item, five
  archivist answers with citations, the door answering HTTP.
- **Custodies nothing.** The whole human ceremony in a real browser —
  passkey enrolment on the bundled fold, sign-in via RFC 7591 DCR +
  code+PKCE, the session's own NATS admission (`lane=oidc role=realm`
  in the audit), a mutation signed by the *principal's* materialized
  persona key, sign-out — and the helm's storage scans clean with a
  positive control that fires. Sessions live in memory only.
- **Configuration without a second control plane.** The mutation table
  (eight rows, three classes) was committed before the spike ran; one
  mutation per class then ran from browser buttons: `work.open` on the
  record; `tokens.create` through the identity plane; and
  `planes.memory.enabled=false` with an in-place restart — clients
  re-admitted, memory answers falling 5 → 0. No mutation needed a new
  side-channel or a helm-owned store.
- **One design system, two consumers.** A single token source (the
  design system's seven files verbatim, fonts and icons vendored, the
  variable `wdth` axis surviving) renders the helm's first screen and
  the fold's sign-in page with every request on the serving host and
  the fonts genuinely loaded.

**What was refuted or reversed.** The drafted server-rendered
recommendation (C4) was overridden by the operator toward a
browser-live hypermedia UI — and the bench then split the difference
honestly: both candidates ran live, and **Datastar won the helm's
rendering** (backend as source of truth over SSE, one 34 KB bundle,
credentials never in the browser) while the NATS-WebSocket lane (C5)
stays decided as the *participant-client* transport `[judgment]` on
measured evidence. The founding owner's lane was refused for token
management — an explicit server `Permissions Violation` on the op
tail, D25's enforcement seen from the consumer position, turning a
convenience assumption into upstream ask #2. And one self-refutation
worth keeping: the first custody positive control was non-hex and
could not have fired — a control that cannot fire proves nothing.

**What it taught / what it opened.** Asking memory is participation,
not observation — the query is a posted op, so the helm's memory panel
rides the signed-in principal `[measured]`. Attribution rides
soulfold's persona-shaped ids end-to-end, with display names a helm UI
concern `[measured]`. The operator's copy rule became C8 and design-
system canon: human surfaces say what things do — Storage,
Connections, People & sign-in, Agents — and the component bynames stay
in internal docs. Two upstream asks stand: **#1** a WebSocket listener
in soulnode's embedded server (one options field); **#2** an
owner-reachable token-management lane, which converges with
[`platform-tenancy-guardrails`](0107-ecosystem-platform-tenancy-guardrails.md)'s
grant work — C7's decision to run fully parallel is already paying.
The participant client (humans posting turns from the browser) is a
named successor topic, not helm scope.

Reversal condition: if the helm cannot remain a pure consumer —
observable as a required `internal/` import, a required privileged
surface existing only for the helm, or a required helm-owned store of
record — the component dissolves back into the existing surfaces. If a
helm interaction needs client-held state or sub-roundtrip latency that
SSE morphing cannot express, the Datastar rendering pick reopens with
that interaction as the evidence.

Trail: design
[`0001-soulhelm-the-helm.md`](../02-DESIGN/soulstream-shell/0001-soulhelm-the-helm.md)
(the component's founding design doc); the rigs (consumer-position
module, WebSocket bench, helm prototype) ran in the session scratchpad
per how-we-work, their measurements recorded in the topic journey;
commits `4bb91ec` → `b174b94` plus this graduation (topic folder
removed, history in git).
