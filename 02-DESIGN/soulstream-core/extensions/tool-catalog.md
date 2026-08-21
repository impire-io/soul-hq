# Extension: the tool catalog (discovery face)

*Optional convention. Graduated from research topic
`agent-external-tools` ([episode
0118](../../../04-JOURNEY/0118-ecosystem-agent-external-tools.md)); the
custody half and the forwarding door live in
[`soulstream-identity/external-tools.md`](../../soulstream-identity/external-tools.md)
(D39–D41). A realm running none of this is still a working soulstream.*

## What it is

One realm-readable answer to "which tools does this soulstream have?" —
uniform across the tools this deployment runs and the remote systems
nobody here runs. It is **display/discovery-grade, never authority**:
the same demotion A10 gave names. Authority lives where it always did —
admission at the transport, outbound custody behind the identity
plane, consent on the record's own grant vocabulary.

## The store

A realm-readable KV bucket, `soulstream-tools`, created on first use by
whoever writes the first entry (this is an extension: nothing in
provisioning mandates it). Key: the tool's name (the foundation's slug
grammar). Value, JSON, additive like every record vocabulary:

| Field | Meaning |
|---|---|
| `name` | the tool's catalog name (equals the key) |
| `kind` | `"remote"` — reached through the identity plane's resource of the same name — or `"workload"` — run by this deployment |
| `persona` | workload kind only: the tool workload's persona, resolvable in the persona registry like any participant |
| `endpoint` | where a door reaches the tool — the remote MCP server's URL for a remote entry, the workload's serving address for a workload one |
| `description` | one plain-language line for screens and agents |

A `remote` entry deliberately carries **no** OAuth endpoints, client
ids, or secrets — the ceremony's half lives with the custody half on
the identity plane, keyed by the same name, so the record never holds
a secret and never partially describes one. The service `endpoint` is
not ceremony: where a door reaches a tool is the catalog's to say for
both kinds — the build's first correction (episode
[0120](../../../04-JOURNEY/0120-ecosystem-the-tools-arc-builds.md)),
made when the door found the remote's own URL living nowhere. A
`remote` entry still refuses a `persona`: a remote tool holds no realm
identity.

`kind` here classifies **catalog entries**, not personas — the persona
taxonomy the protocol removed stays removed. A workload tool's persona
is a persona like any other; what answers for it is the registry's
operator claim, not this field.

## Duties and honesty

- **The writer writes both halves.** Adding a `remote` tool means the
  identity-plane op *and* this entry, by the same hand (the shell's
  module, an operator script). This convention never promises the
  halves agree.
- **Drift fails at use, in words**: an entry whose plane resource is
  absent yields "this tool isn't serving" at link time; an entry whose
  workload is stopped yields the endpoint's own refusal. A reader
  treats the catalog as a map, not a warranty.
- **Consumers**: doors resolve their target list here; shells render
  it; agents discover it the way they discover everything else — by
  reading the realm. The identity plane is **never** a consumer: it
  imports nothing of the record, and this convention keeps it that way.

## Acceptance criteria

1. An entry of each kind round-trips through the reference library
   with unknown fields preserved (additive evolution).
2. A door composing only this convention and the identity plane's
   published surface resolves its targets with no other configuration.
3. The drift case reads honestly: a `remote` entry with no plane
   resource produces a named link-time refusal, not an empty ceremony.
