# 02-DESIGN — the normative design, per component

What each component *is*, specified functionally: **what must exist** and
**how each part behaves**. An implementer should be able to build a working
system from these documents without needing undocumented decisions. Each
component folder keeps its own conventions and its own README index:

| Folder | Component | Convention |
|---|---|---|
| [`soulstream/`](soulstream-core/README.md) | the record | `core/` (the protocol — a realm running only this is a working soulstream) + `extensions/` (optional conventions) |
| [`soulstream-workloads/`](soulstream-workloads/README.md) | the room | Numbered documents (`0001-…`), next free number onward |
| [`soulstream-identity/`](soulstream-identity/README.md) | the name | Named documents carrying global D-numbered decisions (D1, D2, …) |
| [`soulstream/`](soulstream/README.md) | the house | Numbered documents (`0001-…`), next free number onward |
| [`soulstream-idp/`](soulstream-idp/README.md) | the fold | Named documents carrying global D-numbered decisions |

**The spec-kit rule:** every document here is written explicit enough to be
the argument to `/speckit-specify` (run in the component's repo) — the
capability, its seams, its configuration surface, and its acceptance
criteria, with no guessing left to the spec writer. Graduating research
enters through `/research-graduate`; behavioral changes made during
implementation propagate back here — the docs describe each system as it
*is* (see [`../00-GENESIS/how-we-work.md`](../00-GENESIS/how-we-work.md)).
The reasons behind non-obvious calls live in
[`../00-GENESIS/rationale.md`](../00-GENESIS/rationale.md) (soulstream, the
root) and in the D-numbers and design texts themselves (the others).
