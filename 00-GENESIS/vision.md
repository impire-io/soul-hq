# The Soul Vision

## What the ecosystem is

The Soulstream ecosystem — **the soulsystem** — is **a stream on which humans
and AI collaborate as peers, on infrastructure you own**: a protocol with the
components that make it real, never a platform. Every persona, human or AI,
holds the same kind of credential, publishes the same operation record, and is
addressed the same way. There is no bot API and no human API. There is one
protocol.

The name is a solar system, near enough: a record at the centre that
everything writes to, the name / the room / the door in orbit around it, and
the house that gathers all four into one binary you can run
([impire.io/soulsystem](https://impire.io/soulsystem)). Five components, one
sentence each:

| Component | Role | What it is |
|---|---|---|
| **soulstream** | *the record* | The protocol and its reference library: topics as shared workbenches, operations, baselines, personas — collaboration as an op-log over NATS. |
| **soulrealm** | *the room* | The runtime: launches, supervises, observes, and retires a realm's agents and tools as workloads; everything worth keeping flows back to the record. |
| **soulidentity** | *the name* | The identity plane: the home of the persona — vault-held keys, signing oracle, credential minting; signs and mints instead of handing out keys. |
| **soulfold** | *the fold* | The default IAM: a self-hosted, passkey-first OIDC provider — who exists and who belongs — standing exactly where Entra or any OIDC provider may stand instead. |
| **soulnode** | *the house* | The single-binary distribution: the whole stack on a machine you own — `soulnode init && soulnode up`, point a client at the printed URL. |

Soulstream is the record, soulrealm is the room, soulidentity is the name,
soulfold is the fold, soulnode is the house.

## The founding bet

The ecosystem's bet is soulstream's bet, extended outward: **the "what is
needed" list stays short, and everything else is vocabulary over the log,
backends on named seams, or optional extensions.** A working deployment is a
NATS server with JetStream, credentials, and the protocol — no API tier, no
database, no coordinator. Each component holds its own version of the bet:

- **soulstream** — a working soulstream is a NATS server + JetStream, a
  stream, an identity per persona, the protocol, and baselines. Nothing else.
- **soulrealm** — a running agent is a persona, not a service tier.
- **soulidentity** — custody without possession: like an ssh-agent, it signs
  instead of handing out keys; the NATS server stays the verifier of record.
- **soulfold** — an IdP the ecosystem cannot special-case: standard OIDC,
  passkeys only, replaceable by configuration.
- **soulnode** — owning your realm costs one binary and one command, at the
  same shape a hosted deployment runs. Nothing less is worth self-hosting.

## Who it is for

People and agents who want to collaborate as peers, without a second-class
door for the AI — and who want the result on hardware they control.
Concretely: independent professionals and small teams standing up a realm
(account-enforced tenancy over their own NATS), humans on the CLI or a
browser, agents through MCP, all first-class personas on one stream. The
substrate is the product; the bet is that useful clients grow around a stable
protocol rather than the protocol growing around one client.

## Where it is pointed

Each component's sequencing lives in
[`../03-IMPLEMENTATION/ROADMAP.md`](../03-IMPLEMENTATION/ROADMAP.md); the
horizons that shape design today:

- **More vocabulary, never more machinery** (soulstream): versioned artefacts
  and work items shipped; live co-editing (eg-walker) and sealed topics
  gated on real use; a second library language when the wire must prove it is
  the contract.
- **One workload contract, many walls** (soulrealm): native process, microVM,
  and Kubernetes backends proven; fleet without orchestration — placement as
  `work.claim` on the op-log, no coordinator.
- **Consumers wire in** (soulidentity): the signer seam and callout lanes are
  live; attestation issuance and sealing keys arrive when their consumers do.
- **The fold closes the loop** (soulfold): the OP skeleton, passkey
  ceremonies, lifecycle, the consumer-position proof against soulidentity's
  callout, then the embed seam for the distribution.
- **The house ships** (soulnode): first boot is the product; the front door
  on your own network; day 2 without an operator — the state directory *is*
  the realm.

## What we refuse to become

The refusals are the ecosystem's spine; each component's constitution article
makes its own testable (see
[`constitution.md`](constitution.md)):

- **A platform instead of a protocol.** The core answers "what is needed" and
  nothing more; the rest is extensions.
- **A system with a coordinator or a consensus round.** Deterministic rules,
  idempotent ops, optimistic concurrency — no steward, no election, no lock
  service.
- **A second door for the AI.** No bot API, no attribution laundering.
  Delegation is scoped credentials or a separately named persona.
- **A store of record outside the record.** Soulrealm never becomes the home
  of durable truth; sandboxes and runtimes hold views, never the artefact.
- **A KMS, an identity provider, or a parallel permission system.**
  Soulidentity represents subjects; it never authenticates them (identity
  truth lives in the deployment's IAM — the fold or any OIDC provider) and
  never second-guesses what the NATS server enforces.
- **A password store or a privileged peer.** Soulfold is passkeys-only and
  must remain indistinguishable from any external OIDC provider.
- **A dev-mode fork or a hosting platform.** Soulnode runs the real shape on
  your machine; every component remains independently deployable without it.
- **A wrapper around NATS.** Lean on NATS, don't wrap it: subject permissions
  are nearly the whole security model.

## How ambition stays honest

Invented vocabulary is a budget, spent only on concepts a plain word can't
carry (*persona*, *realm*, *baseline* earn their place). Every ambition sits
behind a named, pre-registered research gate, and no build outruns its gate.
Direction decisions record what would change our minds when they are made, so
a future reversal is a clean, anticipated turn instead of drift. Claims carry
evidence classes, and only `[measured]` closes a debate. The full discipline
lives in [`constitution.md`](constitution.md) and
[`how-we-work.md`](how-we-work.md); the per-component founding visions are
frozen with their genesis sets in
[`../99-ARCHIVE/genesis/`](../99-ARCHIVE/genesis/).
