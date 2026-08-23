# Episode 0123 — The sheet shape: tables lead, forms slide over, keys ask first (2026-08-23)

The operator's read of the landed screens: the forms overwhelm ("the
advanced fields should fold away — a slide-out"), the module pages
carry more cognitive load than their jobs need, and "Take sign-in
away" is too long to put on a button. A survey of every form and every
screen found the shapes behind the feeling [measured]: three list
pages led with their add-forms; the tools form showed twelve fields
with roughly half dead for whichever kind was picked (the handler
branches on Kind, the form did not); destructive keys fired on one tap
on three screens while archiving asked first; record vocabulary
(`proposed`, `active`) and machine strings (`work.open ok · {op} ·
(signed=true)`, the full principal) sat on human rows; and below
1180px the details column vanished and took close/archive with it.

One pass answered all of it (design
[`0007-the-sheet-shape.md`](../02-DESIGN/soulstream-shell/0007-the-sheet-shape.md)):
a shared slide-over on the frame's own panel signal — the rail
drawer's mechanics, from the right, at every width — so Tools, People
& sign-in and Agents lead with their tables and creation is a
deliberate act with a surface and a result line of its own (success
closes the panel over `datastar-patch-signals`); the tools form
branches on Kind with the seven provider fields folded further; every
destructive act stands behind an additive GET ask in the archive
confirm's shape (the act endpoints unchanged), behind one-word keys —
Disable, Enable, Revoke, Remove — with the whole sentence in the
hover; rows speak the person's words (new / going on / quiet / closed
/ archived; a work mark's sentence agrees with its stamped strip); raw
ids ride hovers; the storage message panel leads with payload and
verdict, Headers and Signed bytes folded; and the details column
becomes a drawer on its own signal instead of disappearing. 24 files,
+803/−174, on the shell's `ux-cognitive-load` branch atop v0.11.0-rc.2;
the full gate green — unit suites, the realm-booting e2e (21.8s), lint
clean [measured].

The standing gates earned their keep and were re-pointed, not
weakened: the banned-word gate refused the `.archfold` class name on
the people screen — the byname ban reads markup, class names included
— which is why the sheets' fold is named `.stow` [measured]; the
collapse ladder's 1180 step now asserts the drawer where it asserted
`display:none`; the details panel's button audit learned the drawer's
shut key by name; and the e2e rail assertion tripped on exactly the
copy this pass exists to change (`active` → `going on`) [measured].

What it taught: the confirm-ask generalizes cleanly from the archive
pattern precisely because acts never move — asks are additive GETs,
so every direct-POST e2e stayed green [mechanism-argument]. And plain
language cuts both ways: sentences belong in ledes and confirms, one
word belongs on a key — the sentence rides the hover [judgment,
the operator's own words].

Reversal condition: dogfood chafe entries naming the slide-over —
creation undiscovered behind its key, or the confirm tap as friction
on an act somebody performs routinely — reopen the panel-by-default
and ask-every-time choices; the components stay either way.

Trail: design
[`0007-the-sheet-shape.md`](../02-DESIGN/soulstream-shell/0007-the-sheet-shape.md);
the survey and pass in this session; commit `41a0042` on
`ux-cognitive-load` in
[impire-io/soulstream-shell](https://github.com/impire-io/soulstream-shell).
