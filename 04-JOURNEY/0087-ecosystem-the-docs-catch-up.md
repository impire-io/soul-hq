# Episode 0087 — The docs catch up, and the site learns to onboard (2026-08-15)

The operator's observation, the morning after rc.2: the README and
getting-started were outdated, and the website said nothing about how to
actually begin. True on every count — the product README still opened
with the pre-rename "SoulNode" and "Phase 1 is complete"; the guide
claimed "there is no packaged release yet" the day after the archives
shipped, and knew nothing of passkeys, the shell, agents, or the
wrapper; the site's soulstream page pinned v0.7.0 and told the orbit
story in retired names (soulidentity/soulrealm/soulfold/soulnode), with
its "get the code" link pointing at a repo that has since become the
product.

What landed, three repos:

- **soulstream** (`13e4eaa`): README rewritten under the product's own
  name — current pre-release front and center, the five things a person
  on rc.2 can do (passkey sign-in, connect an assistant, mint an agent
  seat, wrap it, run declared workloads). The guide now runs the full
  arc: download the release → `init` (token **and** passkey invite,
  both shown once) → `up` (four URLs) → sign in → connect → the Agents
  screen → `soulstream wrap` → workloads → fronting; the honest
  not-yet list updated (public-mode fronting, BYO NATS,
  harness-as-workload).
- **impire-io.github.io** (`b628486`): a new
  **/soulstream/get-started/** page in the site's voice — five minutes
  from nothing to a wrapped assistant answering mentions — wired into
  the CTAs, the LLM mirrors, and the sitemap (9 urls, was 8). The
  protocol page sheds the stale pin (core v0.8.4) and the pre-rename
  orbit names, and gains the wrapper in "what's true today".
- Not swept, named honestly: the site's story pages
  (`src/content/story/*`) still narrate in the retired names — a full
  site rename pass is its own change, not smuggled into this one.

Reversal condition: none — records a docs alignment; the facts it
states are the release's (episode 0086).

Trail: soulstream `13e4eaa`; impire-io.github.io `b628486` (site builds
green: astro check + 9 mirrors); episodes
[0084](0084-shell-where-this-goes.md)–[0086](0086-soulstream-v0-11-0-rc-2.md)
are what the docs now describe.
