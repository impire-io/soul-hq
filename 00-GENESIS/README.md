# 00-GENESIS — why the ecosystem exists and how it decides

This folder is the fixed point every decision — in any component — is held
against. It changes rarely, deliberately, and always with a journey episode
recording why.

| File | Role |
|---|---|
| [`vision.md`](vision.md) | What the Soulstream ecosystem is — the record, the room, the name, the fold, the house — who it's for, where it's pointed, and what it refuses to become |
| [`constitution.md`](constitution.md) | The testable articles: shared principles S1–S5, the anti-drift working agreement, and per-component articles keeping their original numbering. Canonical copy — every component repo's spec-kit Constitution Check reads it through its `.specify/memory/constitution.md` symlink |
| [`how-we-work.md`](how-we-work.md) | The process: the cross-repo pipeline, research lifecycle, quality gates, documentation duties, and the working agreement in daily terms |
| [`rationale.md`](rationale.md) | How the record (soulstream, the ecosystem's root) got here — the reasons behind every non-obvious call. Not normative |

The five founding genesis sets (per-project visions and constitutions) are
frozen in [`../99-ARCHIVE/genesis/`](../99-ARCHIVE/genesis/); the reasons
behind non-obvious component calls also live with the decisions themselves —
D-numbers and numbered designs in [`../02-DESIGN/`](../02-DESIGN/README.md).

## The decision test

When a choice comes up — a new direction, a shortcut, a scope change — run it
through, in order:

1. **Vision**: does it serve what [`vision.md`](vision.md) says the ecosystem
   is for? If it serves something else (a bigger platform, a convenient
   coordinator, a special door for one client, a component hoarding truth
   that belongs to the record), say so out loud.
2. **Constitution**: does it violate an article — shared or the component's
   own? Articles don't bend for product work. The load-bearing questions are
   usually: does this stay NATS-native (S1), is it the smallest thing that
   satisfies the spec (S2), and does it cross the component's named
   non-negotiable (soulrealm I, soulnode I, soulidentity I, soulfold I/II)?
   If an article genuinely must change, that's an amendment with a version
   bump and a journey episode, never a quiet exception.
3. **Working agreement**: if the decision is load-bearing, it does not get
   recorded until it survives teach-back, carries its evidence class, names
   its reversal condition, and (for changes to a protocol's shape or a core
   boundary) has had the other side argued at full strength. See
   [`how-we-work.md`](how-we-work.md).

If the test doesn't produce a clear answer, the decision waits for the human.
