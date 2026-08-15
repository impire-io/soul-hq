# Episode 0091 — The shell is the console (2026-08-15)

The operator's question cut to the product's shape: "the idp admin needs
to be incorporated into the shell and only accessible to admins — that's
the whole reason we do the shell, no?" Exploration made the answer small:
the idp already published a complete bearer-guarded JSON admin API
(`/api/admin/*` — users, groups, status, invites, OAuth clients), and the
shell's People & sign-in module already consumed three calls of it. The
consolidation was parity, not architecture.

**soulstream-shell v0.6.0**: the People screen grew the rest of the
contract — naming a new person (existence, never admission: the invite
still enrolls), group editing in place, shut-out-and-back-in, and the
applications that sign people in, registered and removed from the same
screen. The spine key is drawn only for sessions whose own token carries
the admin role — a display fact read from the person's own `roles` claim;
the sign-in service's RS256-verified refusal stays the authority behind
every act [mechanism-argument: the API re-verifies every call]. The e2e
walks the whole lifecycle shell-natively and proves the other arm: a
non-admin sees no key, and the refusal that meets them at the path is the
service's own, in its words [measured, `make check` green].

**soulstream-idp v0.5.0**: the HTML console became a serve/embed option
(design D31) — on by default for standalone deployments, unmounted when
the soulstream product embeds the plane, where `/admin` answers like any
path nobody claimed [measured: pinned 404 in the idp suite and the
product's node suite]. The API/console split keeps standalone estates
whole: the console was always presentation over the same API, never a
second authority.

**soulstream** (rode specs/009): the sign-in plane passes the console-off
option; `up` prints no admin-console line; the getting-started's
break-glass sentence points at the shell's People screen.

Reversal condition: a measured lockout incident that only a fold-local
console could have fixed brings the bundled console back behind an
explicit break-glass flag (recorded in D31).

Trail: design
[`session-and-ui.md` D31](../02-DESIGN/soulstream-idp/session-and-ui.md);
commits — idp `7ce4383` (tag v0.5.0), shell `73cc95b` (tag v0.6.0),
product in `specs/009` (episode 0092); soul-hq `da233d3`.
