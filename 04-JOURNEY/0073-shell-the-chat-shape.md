# Episode 0073 — The chat shape: a rail, a conversation, a docked composer (2026-08-14)

The operator redirected the cockpit's form with a reference screenshot
— "more of a traditional chat interface" — and the second slice of
episode [0071](0071-ecosystem-the-focus.md)'s usable cockpit landed it
the same morning: a left rail of conversations (the board projection,
live), the conversation as a chat column — the signed-in person's own
messages right-aligned in amber bubbles, everyone else's left-aligned
cards with name and time, anchored replies visually attached under the
message they answer, work items rendered as quiet system lines in the
flow — and the composer docked at the bottom ("Write a message…",
Send, the reply state above it as a pill with Cancel). Attribution is
decided server-side from the record and the session, never the
browser's word `[measured: asserted in the gate]`. The whole
consumer-position ceremony stays green, ~6.8 s uncached at review
`[measured]`.

**The screens helper was born with it** — `make screens` boots a full
realm, performs the passkey ceremony at protocol level with the
virtual authenticator, seeds a two-person conversation, and hands a
signed-in session cookie to a real browser. The first real screenshots
it produced immediately earned their keep: the wordmark and titles
still said *soulsystem* — the name episode
[0069](0069-ecosystem-one-name-soulstream.md) retired — fixed within
the hour. A missing favicon is the named nit.

What waits on the operator's eye before more surface grows: the
right-hand details panel (work items, attachments, people) and mention
notifications — both queued behind the screenshots now in hand.

Reversal condition: none — records a completed build; the direction's
reversal readings live in episode 0071.

Trail: soulstream-shell `5690f41` (the shape, built by a delegated
session against the operator's reference), `a3946e1` (the retired
name); screenshots `shell-chat-conversation.png` /
`shell-chat-reply-state.png` shown to the operator; design 0001's
rendering rules honored (stream and act responses never share a patch
target).
