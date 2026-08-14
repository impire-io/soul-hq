# Episode 0072 — The composer: the shell stops being a window (2026-08-14)

The first slice of episode [0071](0071-ecosystem-the-focus.md)'s usable
cockpit landed the same morning: a signed-in human writes a message
into a conversation from the topic view, and it lands on the record
through their **own** admitted connection — authored by their principal,
signed with their key, the shell still signing as no one. Reply came
nearly free: a per-message control moves the composer's anchor, the
anchor is resolved against the record rather than taken on the
browser's word, and the answer lands as a comment on the message it
names. The whole gate — the consumer-position ceremony now including
compose, post, and anchored reply — runs green in ~6.9 s, re-verified
uncached at review `[measured]`.

The build honored the recorded rendering lesson (one-shot act
responses and the live stream never share a patch target): the
composer owns three targets — box, anchor, note — so a half-written
message survives every morph.

**Two pre-existing bugs surfaced and fixed**, both load-bearing:

- **SSE frames truncated at the first newline.** Patch frames were
  written as a single `data:` line, and a raw newline ends an SSE
  field — every dash frame reached the browser cut off at the first
  vendored icon's line break. Each line now gets its own `data:` line;
  the gate grew a frame bar whose positive control fails against the
  old writer `[measured]`.
- **Icons rendered unsized** outside buttons: `Icon()` stripped the
  vendored `width`/`height` while only `.btn svg` had a CSS restore.
  The vendored size stays on.

First episode under the naming map's short tag vocabulary (`shell`).

Reversal condition: none — records a completed build; the participation
direction's reversal lives in episode 0071.

Trail: soulstream-shell `1a8784d` (built by a delegated session against
the episode-0071 brief, reviewed and gate-re-run before push); design
0001 §1 as amended; the
[`shell-module-contract`](../01-RESEARCH/shell-module-contract/README.md)
bars await the module re-homing spike.
