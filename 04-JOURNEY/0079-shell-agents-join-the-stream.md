# Episode 0079 — Agents join the stream (2026-08-14)

The operator's ask closed the founding loop: "create agent credentials
so we can manage the agents on the system… configure the stdio mcp
server to use those creds… we will know which agent it is and it can
start collaborating." Landed across three repos, the full path
measured end to end `[measured]`:

- **Core grew the revocable lane** (v0.8.2): `realm.Config` dials
  direct with URL + sentinel creds + token — an agent's machine has no
  saved context, and until now the stdio door could name an identity
  but never be handed one. Neither half admits alone; the callout
  exchanges the pair for a scoped principal; taking the token away
  refuses the next connection.
- **The shell grew the Agents module** (v0.4.0): "The machine voices
  somebody here answers for." Add an agent (handle + shown-as) and the
  support layer mints its identity-plane token over the node-standing
  ops lane (the founding design's sanctioned class-(b) arm — the
  shell still holds nothing), while the creating person's own session
  signs the `operated_by` attestation: "You vouch for what you add:
  your name goes on it, signed with your own key, and stays there."
  The credential appears once, as the exact copy-paste block the stdio
  adapter accepts. Manage: list, revoke ("take the credential away"),
  re-mint — and the operator claim survives credential changes,
  because who vouched for a voice is a thing that happened.
- **The product declares the fact** (v0.10.0): the Agents module
  activates where this node issues agent credentials; absent, no rail
  entry, 404s.

The e2e proves the collaboration, not just the ceremony: an agent
created from the browser connects with exactly the emitted credential,
posts into the conversation, and renders as a **teal voice with its
operator named** through the real path — no rig seeding; an @-mention
from the composer reaches its inbox; browser revocation refuses the
next connection within the token lane's bound; and the replacement
credential is proven to re-admit, not only offered `[measured]`.

Named nit: the Agents table shows the operator's raw persona id where
the display name belongs (the O3 mapping family). Tags: core v0.8.2 →
shell v0.4.0 → soulstream **v0.10.0**, all gates and releases green.

Reversal condition: none — records a completed build; the revocation
bound is the identity plane's recorded token-lane semantic and moves
only with it.

Trail: soulstream-core `7806c14`, soulstream-shell `a6b88dc` +
`eeae564`, soulstream `df53d9e`; screenshot `shell-v9-agents.png`
shown to the operator; design 0002 §2's activation pattern carries a
second real instance.
