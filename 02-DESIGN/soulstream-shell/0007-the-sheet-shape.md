# 0007 — soulstream-shell: the sheet shape

**Status:** designed with its build — decided 2026-08-23. The
cognitive-load pass over every sheet screen: what a list page leads
with, where creation lives, what a destructive key must ask, and the
words rows are allowed to say. Grew out of reviewing the screens the
shell arc landed (episode
[0122](../../04-JOURNEY/0122-ecosystem-the-shell-arc-lands.md)) as a
whole rather than one at a time.

## §1 The problem

Each sheet screen was honest alone and heavy together. Three list
pages (Tools, People & sign-in, Agents) led with their add-forms —
most visits are about what exists, not about making more. The tools
form showed twelve fields at once, roughly half of them dead for
whichever kind was picked. Destructive keys fired on one tap on three
screens while archiving a conversation asked first. Record vocabulary
(`proposed`, `active`, `dormant`) and machine strings (`work.open ok ·
{op} · (signed=true)`, the full principal) leaked onto human rows.
And below 1180px the details column disappeared and took the
close/archive acts with it.

## §2 The list-screen rhythm (normative)

Every sheet list screen reads, top to bottom: **h1 → lede → the act
key → the list → folded side-matter → the result line → the
slide-over**. The table is what the screen is for and it leads;
creation is a deliberate act with a surface of its own.

- **The slide-over** (`shell.SlideOver`, `.slideover` in the shell's
  CSS layer): the sheet's own side panel, from the right, at every
  width — the same mechanics as the conversations rail's drawer
  (transform + delayed visibility, the one modal scrim) and the same
  page-local signal, `$panel`, because no screen has both. Served with
  the page, never morphed, so a half-written form survives whatever an
  act patches. Its form owns a result line of its own
  (`{module}-add-note`): what goes wrong answers beside the fields it
  is about; what goes right closes the panel (`datastar-patch-signals`
  via `shell.PatchSignals`) and answers on the screen it slid over.
- **The stow fold** (`.stow`): the `<details>` fold under a name a
  sheet may say — the banned-word gates keep component bynames off
  product screens, class names included, so `archfold` stays on the
  conversations surface and the sheets fold with this one. Used for:
  the admin's "Apps that sign people in" section (a day-one task),
  approvals' "How this works" prose, the tools form's "Provider
  sign-in" block, and the storage message panel's Headers and Signed
  bytes — the payload and the verdict lead, the wire forms rest one
  click away, all still served.

## §3 Forms are sections, and show what the choice reads

A form with a kind-select is sections: what every kind takes first
(name, kind, description), then **one labeled section per kind and
none for the other** — the select drives a page-local signal
(`data-bind:kind`), each section stands behind `data-show`, and a
label strip names it the way the details panel names its own sections.
The tools form's remote branch is "Connected service" (its address)
and "Provider sign-in" (the seven provider fields with a plain
sentence about where they come from); the runs-here branch is "Runs
here" (address, runs-as). A field both kinds take but mean differently
— the address — lives in each section with its own words, `disabled`
while hidden (`data-attr:disabled`) so only the living one submits.
Required inputs carry `required`; optional ones say "— optional" in
their placeholder.

**The register, calibrated (operator's rule, 2026-08-23):** the
audience is smart and not technical. Machine-room vocabulary never
reaches a screen — "identity plane" becomes "the sign-in service",
"where its MCP server answers" becomes "from the service's
documentation" — and an explanation that merely restates jargon
explains nothing. The machine room's own names survive only where the
person will meet them elsewhere: the provider's console says "scopes",
so the form does too.

## §4 A destructive key asks first

Every destructive act stands behind a question, in the archive
confirm's own shape: the row's key `@get`s an ask route
(`/people/disable-ask`, `/people/client-remove-ask`,
`/agents/revoke-ask`, `/tools/remove-ask`) that patches the question
— what changes, what stays, `Yes, …` / `Keep it` — into the screen's
result line; asked about nobody, the route clears it, which is what
"Keep it" does. The act endpoints are unchanged and remain the
authority; the question is the surface's own manners. Buttons say one
short word (**Disable**, **Enable**, **Revoke**, **Remove**) with the
whole sentence in the hover — plain language does not mean long
sentences on keys.

## §5 Rows speak the person's words

- Lifecycle words on rows are the person's (`stateWords`): new, going
  on, quiet, closed, archived — the details panel already said these
  things in sentences; the record's vocabulary stays on the record. An
  unknown word arrives as itself: newer records outrank the list.
- A work mark's sentence agrees with its stamped strip: opened / took
  up / finished, never "opened" under a `claimed` stamp.
- Raw identifiers ride hovers, never cells: the approvals principal,
  the full op id behind the status screen's plain answer ("Work item
  opened by … — signed and on the record"), the day behind a chat
  message's clock, the moment behind an agent's Added date. The
  approvals window says "4m10s left" with the exact deadline in the
  hover — and the screen re-reads itself every five seconds
  (`/approvals/live`, the storage tail's shape): countdowns run,
  arrivals appear, expiries go, and the mark on the spine counts down
  to nothing and back. The tally is always an element carrying its id,
  so the tick can patch it either way; the result line stays the
  acts' own, never the stream's.

## §6 The details column steps aside, not away

Below 1180px the conversation's details column becomes a drawer over
the thread on its own signal (`$info`, declared on the thread; the
Details key lives in the conversation's head and exists only at those
widths), with a shut key of its own and no scrim — one surface keeps
one scrim, and the conversation stays readable beside it. The
lifecycle dock rides over the drawer, so close/archive and their
answers stay reachable at every width. The class binding
(`data-class:open="$info"`) is on every render of the panel, because a
morph replaces attributes — lose it once and the drawer never opens
again.

## §7 What this does not change

No act endpoint moved or changed shape; the asks are additive GET
routes. The record's vocabulary is untouched — only its on-screen
translation. The credential card, the invite card, and the shown-once
discipline are as they were. The purity gates, the banned-word gates,
and the collapse ladder all still hold, re-pointed where the ladder's
1180 step now reads "drawer" instead of "gone".
