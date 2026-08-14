# Episode 0081 — The last administrator stays (2026-08-14)

The operator found they could saw off the branch they sit on: People &
sign-in let the founding owner — the only administrator — be disabled
or demoted with their own bearer, after which nobody could administer
the deployment again short of store surgery.

The rule landed where it is authoritative (soulstream-idp `2eaa7e9`,
v0.4.1): **the admin group may never become empty of enabled people.**
Both lethal paths refuse — disabling, and taking the admin group away
— for any client, any bearer, including the person themselves; there
is no user-delete op on any surface, so those two are the whole lethal
surface. The refusal speaks the same plain words on the machine API
(409) and the console flash: "the last administrator stays — add
another administrator first."

**Enforced atomically, not counted-and-hoped**: the store has no
cross-key transaction (D4), so the enabled administrators moved under
one key (`idx.admins`) and every lethal act rides its CAS. Measured:
25 rounds × 8 racing disables of the last two administrators — 100
accepted, 100 refused, exactly one administrator standing every round
— beside a control proving the pre-guard shape empties the group
`[measured]`.

The shell reflects it honestly (v0.4.2): the lethal actions are not
offered on whoever is currently the last enabled administrator (a
quiet mono note instead), and a refusal that still arrives renders in
the idp's own words. With a second administrator enrolled, the first
can be disabled again — the rule protects the *last*, not any one
person. **Named, not built**: whether the founder deserves protection
beyond last-admin-standing (with two admins, one can still disable the
founder, reversibly) — the operator's evaluation decides.

Tagged as the operator asked: idp v0.4.1 + shell v0.4.2 composed into
**soulstream v0.11.0-rc.1** — a pre-release, goreleaser marking it so
— the evaluation candidate carrying the whole two-day arc.

Reversal condition: none — records a completed build; the invariant is
a floor, and only a stronger founder rule would amend it.

Trail: soulstream-idp `2eaa7e9`, soulstream-shell `d6ff5e8`,
soulstream `20c8fc2` (v0.11.0-rc.1); lifecycle design D20–D24 gains
the invariant at its next re-vendor.
