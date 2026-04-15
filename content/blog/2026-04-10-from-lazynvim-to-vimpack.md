---
title: "From lazy.nvim to vim.pack"
date: 2026-04-15
tags: ["neovim", "lua"]
categories: []
featured_image: "/blog/assets/neovim/nvim-v0.12-vimpack.png"
---

## Background

I've been using the excellent [lazy.nvim](https://github.com/folke/lazy.nvim)
package manager for more than three years now, and I've been super happy with
it. But with Neovim v0.12.0, `vim.pack` was shipped: a built-in (but still
experimental) plugin manager that manages plugins using Git, with no third-party
dependencies required, implemented by
[Evgeni Chasnovski](https://github.com/echasnovski) (see
[neovim/neovim#34009](https://github.com/neovim/neovim/pull/34009)), known for
his work on [mini.nvim](https://github.com/nvim-mini/mini.nvim/). This piqued my
interest, as I've found myself creating abstractions and isolations with
lazy.nvim that don't harmonize with my [grug brain](https://grugbrain.dev). So I
figured I wanted to see if I could simplify by moving onto `vim.pack`.

If you, like me, have a good deal of plugins installed, you might miss
out-of-the-box features from lazy.nvim such as lazy-loading, defining load order
of plugins while passing around opts from plugin to plugin, build commands,
per-project overrides, ease of management via a nice TUI, local plugin
development. This blog post aims to outline what I've done in my
[personal Neovim config](https://github.com/fredrikaverpil/dotfiles/blob/main/nvim-fredrik)
to solve all this with minimal helpers.

One caveat though; I realized I don't _want_ to pass opts between plugins. I can
just inline everything into each plugin's opts most of the time and it's fine
and simple. But I only came to this conclusion after first having solved that,
so if you have this as a requirement, I will cover it below.

Jump straight to [the summary](#summary) if you just want the TL;DR. Or strap in
for one deep rabbit hole!

### Quick intro to `vim.pack`

Evgeni has written
[A guide to `vim.pack`](https://echasnovski.com/blog/2026-03-13-a-guide-to-vim-pack) -
also comes with a nice
[YouTube video](https://www.youtube.com/watch?v=J1r0vrqOMJo). But in short:

In Neovim, plugins are stored in a dedicated directory — `site/pack/core/opt`
under Neovim's data path — and a lockfile (`nvim-pack-lock.json`) in your config
directory tracks exact revisions. The API is intentionally small:
`vim.pack.add()`, `vim.pack.update()`, `vim.pack.del()`, and `vim.pack.get()`.
There are no user commands and no TUI from the get-go.

A minimal setup might look like this:

```lua
-- init.lua
vim.pack.add({
  { src = "https://github.com/lewis6991/gitsigns.nvim" },
  { src = "https://github.com/hat0uma/csvview.nvim", version = "main" },
  { src = "https://github.com/Saghen/blink.cmp", version = vim.version.range("1.*") },
})
```

The full reference is available in `:h vim.pack` or in the
[online docs](https://neovim.io/doc/user/pack.html#vim.pack).

On first start, Neovim clones missing plugins and makes them available
immediately after the `add()` call returns.

> [!TIP-] Cheat sheet
>
> When running `vim.pack.update()`, Neovim opens a confirmation buffer with
> built-in LSP features for managing individual plugins:
>
> - `]]` and `[[` to navigate between plugin sections
> - `gO` (`:h vim.lsp.buf.document_symbol()`) to list all plugins
> - `K` (`:h vim.lsp.buf.hover()`) to show details about a pending change or
>   newer tag
> - `gra` (`:h vim.lsp.buf.code_action()`) to update, skip, or delete individual
>   plugins
>
> Confirm all updates with `:w`, or discard with `:q`.
>
> **Syncing across machines:** put `nvim-pack-lock.json` under version control.
> On a secondary machine, pull the lockfile, `:restart` Neovim (new plugins are
> installed automatically), then run
> `vim.pack.update(nil, { target = 'lockfile' })` to align all plugins to the
> lockfile's revisions. Plugins removed from the lockfile can be cleaned up with
> `vim.pack.del()`.
>
> **Reverting an update:** revert the lockfile (e.g.
> `git checkout HEAD -- nvim-pack-lock.json`), `:restart`, and run
> `vim.pack.update({ 'plugin' }, { offline = true, target = 'lockfile' })` to
> roll back to the previous revision.

## Lazy-loading with `vim.pack`

In essence, what we are talking about here is we want to defer e.g. expensive
`require()` calls or heavy work carried out by the plugin, for later, when
Neovim has "fully started up". This is the ever-lasting pursuit of making your
Neovim setup not showing any tangible delay to get to the welcome screen, which
puts your vscode-using colleagues in awe. 😉

### Neovim's stance on lazy-loading of plugins

Neovim's own plugin development guide (`:h lua-plugin`,
[online](<https://neovim.io/doc/user/lua-plugin/#_defer-require()-calls>)) is
very clear on this (and has implementation examples):

> [!QUOTE]
>
> Plugins should arrange their "lazy" behavior once, instead of expecting every
> user to micromanage it.

In other words, a well-written plugin already defers its heavy work internally.
That said, not all plugins are born equal and `vim.pack` offers no easy-to-use
lazy-loading machinery of its own. So if you want to control startup timing with
`vim.pack`, you have to wire it up yourself.

### Built in lazy capabilities in `vim.pack`...?

`vim.pack.add()` accepts a `load` option that controls _whether and how_ plugin
scripts are sourced after installation:

| `load` value | What happens                                                                                                                                    |
| ------------ | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| `true`       | Runs `:packadd` — adds the plugin to the runtime path **and** sources its `plugin/` and `ftdetect/` scripts immediately.                        |
| `false`      | Runs `:packadd!` — adds the plugin to the runtime path but does **not** source `plugin/` or `ftdetect/` scripts _right now_ (see caveat below). |
| `function`   | `vim.pack` does nothing; your function receives `{ spec, path }` and is fully responsible for loading the plugin however you see fit.           |

The default depends on _when_ `add()` is called. Internally, `vim.pack` checks
`v:vim_did_init` — a variable that is `0` during `init.lua` and exrc sourcing
(steps 7b–7c of `:h initialization`,
[online](https://neovim.io/doc/user/starting.html#initialization)) and flips to
`1` at step 10, _before_ `plugin/` files are sourced at step 11. So the default
is `false` only inside `init.lua` and `.nvim.lua` (exrc), and `true` everywhere
else — including `plugin/` files, autocmd callbacks, and interactive use. 😵‍💫

> [!CAUTION] Gotcha!
>
> `load = false` does **not** always mean "never load". During startup,
> `:packadd!` defers sourcing to the load-plugins phase (step 11 of
> `:h initialization`). Since the plugin is now on the runtime path, Neovim's
> normal startup walks it and sources `plugin/` files anyway. So `load = false`
> in `init.lua` only avoids _eagerly_ sourcing — the scripts still run moments
> later at step 11.
>
> But here's the _real_ gotcha: after startup (e.g. in an autocmd callback),
> `load = false` genuinely prevents sourcing, because the load-plugins phase has
> already passed and won't run again. 🤯

Remember, `vim.pack` is a library, not a framework (like lazy.nvim is), so we
shouldn't expect the same kind of UX. But still, I don't think this is easy for
most Neovim users to fully grok.

## Finding the right pattern

It was somewhere around this point I was wondering if it was really worth my
time to try and port my config from lazy.nvim into `vim.pack` yet, or maybe wait
until a future release which would bring a more user-friendly API. However, I'm
blessed with the deadly combo of curiosity and stubbornness, so I pressed on
with the perhaps naive viewpoint that plugins _should_ follow best practices. If
not, we could open issues or PRs to fix it.

It took some time back and forth of fiddling and trying things out, but I landed
on this:

> [!HINT] Fredrik's `vim.pack` pattern
>
> - All plugins should be assumed to use `load = true`, unless specifically
>   instructed otherwise. This is what the plugin author intended.
> - All plugins reside under the Neovim config's `plugin/` folder.
> - For lazy-loading, wrap `vim.pack.add` along with _all_ other configuration
>   in a `VimEnter` autocmd.
> - Write to a `_G.Config` "registry" before `VimEnter` which can store global
>   states, plugin opts etc. Once plugins have loaded after `VimEnter`, they can
>   read from the shared states in the registry, thus enabling cross-plugin opts
>   sharing.
> - For per-project overrides, collect them "`exrc`-style", but execute them
>   after all plugins have loaded.

This reduces much of the `vim.pack` loading complexity down to a simple mental
model.

### Lazy-loading

By deferring most of all plugin sourcing onto the `VimEnter` event, the startup
of Neovim becomes very snappy. A small
[`lazyload.lua`](https://github.com/fredrikaverpil/dotfiles/blob/main/nvim-fredrik/lua/lazyload.lua)
module handles this with two queues — `VimEnter` and an override queue — each
draining in order:

> [!EXAMPLE-] `lua/lazyload.lua`
>
> ```lua
> local M = {}
>
> local vim_enter_queue = {}
> local override_queue = {}
>
> local function drain(queue)
>   for _, entry in ipairs(queue) do
>     if not entry.sync then
>       vim.schedule(entry.fn)
>     end
>   end
>   for _, entry in ipairs(queue) do
>     if entry.sync then
>       entry.fn()
>     end
>   end
> end
>
> local function drain_override()
>   if not override_queue then return end
>   for _, entry in ipairs(override_queue) do
>     vim.schedule(function()
>       local ok, err = pcall(entry.fn)
>       if not ok then
>         vim.notify((".nvim.lua override error:\n%s"):format(err), vim.log.levels.ERROR)
>       end
>     end)
>   end
>   override_queue = nil
> end
>
> vim.api.nvim_create_autocmd("VimEnter", {
>   once = true,
>   callback = function()
>     drain(vim_enter_queue)
>     vim_enter_queue = nil
>     drain_override()
>   end,
> })
>
> function M.on_vim_enter(fn, opts)
>   local sync = opts and opts.sync or false
>   if vim_enter_queue then
>     table.insert(vim_enter_queue, { fn = fn, sync = sync })
>   elseif sync then
>     fn()
>   else
>     vim.schedule(fn)
>   end
> end
>
> function M.on_override(fn)
>   if override_queue then
>     table.insert(override_queue, { fn = fn })
>   else
>     vim.schedule(fn)
>   end
> end
>
> return M
> ```

`drain()` does two passes: first it schedules async callbacks via
`vim.schedule()`, then it runs synchronous ones. Callbacks are async by default
— pass `{ sync = true }` for the rare plugin that must be fully set up before
the UI draws (like a statusline).

The execution order is: `VimEnter` sync → `VimEnter` async → overrides. Since
`vim.schedule` is FIFO, overrides always run last — which is what makes
per-project overrides (covered later) work.

Then you defer the plugin from loading and setting up like so:

```lua
-- plugin/<name>.lua
require("lazyload").on_vim_enter(function()
  -- build command on plugin install/update
  vim.api.nvim_create_autocmd("PackChanged", { ... })

  -- add plugin
  vim.pack.add(...)

  -- configure plugin
  require("plugin").setup({ ... })

  -- keymaps
  vim.keymap.set( ... )
end)
```

[`VimEnter`](https://neovim.io/doc/user/autocmd.html#VimEnter) fires after all
`plugin/` files have been sourced. See `:h autocmd.txt`
[online](https://neovim.io/doc/user/autocmd.html) for all available events.

Callbacks registered after `VimEnter` has already fired execute immediately (or
via `vim.schedule` for async), so this pattern is safe to use from any `plugin/`
file regardless of load order.

> [!NOTE]
>
> Things like colorschemes and a startup dashboard would _not_ be lazy-loaded
> via `on_vim_enter`.

### Passing opts from one plugin to another

In lazy.nvim, any plugin spec can declare a dependency and pass opts into it.
lazy.nvim deep-merges all contributions automatically. For example,
[mason](https://github.com/mason-org/mason.nvim) can add itself to
[lualine](https://github.com/nvim-lualine/lualine.nvim)'s extensions list while
also passing opts to itself:

```lua
-- lazy.nvim style
return {
  "mason-org/mason.nvim",
  dependencies = {
    {
      "nvim-lualine/lualine.nvim",
      opts = {
        extensions = { "mason" },
      },
    },
  },
  opts = { PATH = "append" },
}
```

> [!WARNING]
>
> Systematic cross-plugin opts merging is arguably an anti-pattern in
> traditional Neovim idioms. From what I know, it originated with
> [LazyVim](https://github.com/LazyVim/LazyVim) and was designed with a distro
> in mind, where plugins could be cherry-picked by the user. Neovim's own
> `vim.lsp.config()` uses a similar layered merge, but scoped to a single
> subsystem — not one plugin reaching into another's config.
>
> That said, the decentralized approach is appealing to me, compared to tangling
> plugin configurations up in each other. But it comes at a cost; abstractions
> and complexity. And since I'm writing a personal Neovim config (not a distro),
> I here prefer simplicity and explicitness over complexity and implicitness.
> So, I've actually opted for _not_ doing this. I will, however, continue
> explaining how my setup supports it together with using `vim.pack`.

So, with `vim.pack`, there is no dependency graph and no automatic opts merging.
Every `plugin/` file is self-contained and loads in alphabetical order. How can
we achieve this behavior without lazy.nvim?

I solved this with a small shared global config "registry" that acts as a
central coordination point. Plugin files **register** data into it immediately
on load of `plugin/**/*.lua`, and the same plugin files can also **read** from
it in deferred callbacks, after all files have had their chance to contribute.

Let's look at the registry:

```lua
-- init.lua
local merge = require("merge")

_G.Config = {
  lsp = {},
  mason = {},
  conform = {},
  lint = {},
  lualine = {},
  -- ... other fields
}

function _G.Config.add(spec)
  merge(_G.Config, spec)
end
```

Each plugin gets its own namespace ("lsp", "mason", "conform" ...). Plugins with
a `setup(opts)` function store their opts under `.opts`, while other data (like
`mason.ensure_installed` or `lsp.servers`) lives alongside it.

The `merge()` function is a custom deep merge that **appends lists** (instead of
replacing them like `vim.tbl_deep_extend` does) and **recurses into dicts**:

```lua
-- lua/merge.lua
local function merge(base, override)
  for k, v in pairs(override) do
    if v == vim.NIL then
      base[k] = nil
    elseif type(v) == "table" then
      local bv = base[k]
      if type(bv) ~= "table" then
        base[k] = v
      elseif vim.islist(v) then
        for _, item in ipairs(v) do
          if type(item) == "table" or not vim.list_contains(bv, item) then
            table.insert(bv, item)
          end
        end
      else
        merge(bv, v)
      end
    else
      base[k] = v
    end
  end
  return base
end

return merge
```

This is the key difference from `vim.tbl_deep_extend("force", ...)`: when two
plugins both contribute to the same list (say, both add mason tools), the items
are appended and deduplicated instead of the second one silently replacing the
first. And `vim.NIL` lets you explicitly remove a key — useful in project-local
overrides where you want to disable something a lang file declared:

```lua
Config.add({
  lint = { linters_by_ft = { go = vim.NIL } },
})
```

Now any plugin file can contribute data to any other plugin — regardless of
alphabetical load order — because all contributions happen at load time
(immediate), and all consumption is deferred (`VimEnter`).

Mason contributes its lualine extension and reads its own opts from the
registry:

```lua
-- plugin/mason.lua

Config.add({
  lualine = { opts = { extensions = { "mason" } } },
})

require("lazyload").on_vim_enter(function()
  vim.pack.add({
    { src = "https://github.com/mason-org/mason.nvim" },
  })

  local merge = require("merge")

  local opts = { PATH = "append" }
  merge(opts, Config.mason.opts or {})
  require("mason").setup(opts)
end)
```

Lualine reads everything that was contributed:

```lua
-- plugin/lualine.lua

require("lazyload").on_vim_enter(function()
  vim.pack.add({
    { src = "https://github.com/nvim-lualine/lualine.nvim" },
  })

  local merge = require("merge")

  local opts = { extensions = { "man", "quickfix" } }
  merge(opts, Config.lualine.opts or {})
  require("lualine").setup(opts)
end)
```

### Build command

Some plugins need a build step after install or update, like compiling a binary,
downloading assets, or updating parsers. In lazy.nvim this is the `build` key.
With `vim.pack`, the equivalent is the `PackChanged` autocmd event.

`PackChanged` fires after a plugin's state has changed. The event data includes:

- `ev.data.kind` — `"install"`, `"update"`, or `"delete"`
- `ev.data.spec` — the plugin's full spec (including `.name`)
- `ev.data.path` — full path to the plugin directory

Here's how nvim-treesitter's `:TSUpdate` maps to a `PackChanged` hook:

```lua
-- plugin/nvim_treesiter.lua
require("lazyload").on_vim_enter(function()
  vim.api.nvim_create_autocmd("PackChanged", {
    callback = function(ev)
      if ev.data.spec.name == "nvim-treesitter" then
        vim.cmd("TSUpdate")
      end
    end,
  })

  vim.pack.add({
    { src = "https://github.com/nvim-treesitter/nvim-treesitter", branch = "main" },
  })
end)
```

### Local plugin development

When developing Neovim plugins, it's been really nice to be able to define
`dev = true` in lazy.nvim, which makes the plugin load from local disk instead
of the remote git repo. We can achieve this quite easily too.

```lua
-- lua/dev.lua
local M = {}

function M.prefer_local(local_path, remote_src)
  local expanded = vim.fs.normalize(vim.fn.expand(local_path))
  if vim.uv.fs_stat(expanded) then
    return expanded
  end
  return remote_src
end

return M
```

Then use it as the `src` in `vim.pack.add`:

```lua
local dev = require("dev")

vim.pack.add({
  { src = dev.prefer_local("~/code/public/neotest-golang", "https://github.com/fredrikaverpil/neotest-golang") },
})
```

### Per-project overrides

lazy.nvim has `.lazy.lua` for per-project config. Neovim has built-in exrc
(`:h exrc`, [online](https://neovim.io/doc/user/options.html#'exrc')) — drop a
`.nvim.lua` in your project and Neovim sources it on startup, guarded by a trust
prompt (`:h :trust`).

Here we can leverage the same capabilities as we have for a plugin; register
data into the `_G.Config` and lazyload plugins. But we can also perform
per-project overrides, by using the `on_override` autocmd wrapper. Anything
wrapped in this will be deferred to loading after all plugins have loaded, and
thus provide a "final say"; perfect for per-project settings.

Here's how I set custom markdown formatting and gopls overrides for a certain
project, by wrapping it all into `on_override`:

```lua
require("lazyload").on_override(function()
    -- Override markdown formatter
    require("conform").formatters_by_ft.markdown = { "mdformat" }
    require("conform").formatters.mdformat = {
        prepend_args = { "--number", "--wrap", "80" },
    }

    -- Override gopls settings
    vim.lsp.config.gopls.settings = {
        gopls = {
            analyses = {
                ST1000 = false,
                ST1020 = false,
                ST1021 = false,
            },
        },
    }
end)
```

A nuance to be aware of here is `exrc` walks upwards, from `$cwd` to `$HOME`.
This means a `.nvim.lua` in `$cwd` loads _before_ a `.nvim.lua` file in `$HOME`.
I initially built a custom variant of exrc which walked in reverse, but I
dropped it in favour for simplicity (I simply don't need it).

## Summary

Recreating the lazy.nvim features I cared about on top of `vim.pack` came down
to a handful of small, composable pieces:

- **Lazy-loading**: Use `lazyload.lua` to queue setup behind `VimEnter` for most
  plugins (except colorscheme, startup dashboard etc).
- **Cross-plugin opts**: a shared registry that plugin files write to at load
  time and read from at deferred setup, paired with a `merge.lua` helper that
  appends+dedups lists instead of replacing them.
- **Build hooks**: the `PackChanged` autocmd with a filter on
  `ev.data.spec.name`.
- **Local development**: a tiny `dev.lua` helper that returns a local path when
  present, falling back to the remote URL — used as `src` in `vim.pack.add()`.
- **Per-project overrides**: `.nvim.lua` via built-in exrc (`'exrc'`), calling
  `lazyload.on_override()` to defer execution until after all `VimEnter` plugin
  setup.

I keep all these helpers in
[here](https://github.com/fredrikaverpil/dotfiles/tree/main/nvim-fredrik/lua).

The property of this design is that _registration_ is immediate and
_consumption_ is deferred. Any `plugin/` file can contribute data to any other
plugin regardless of alphabetical load order, because all contributions happen
at load time _before_ any of the consumers read them at `VimEnter`.

Here's how it all slots into Neovim's startup sequence (see `:h initialization`
[online](https://neovim.io/doc/user/starting.html#initialization) for the full
spec):

| Step  | Neovim phase                              | What this config does                                                                                                                                    |
| ----- | ----------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 7b    | `init.lua` runs                           | Leader keys, `require("options")`, diagnostics.                                                                                                          |
| 7c    | `.nvim.lua` runs (if `'exrc'` is enabled) | Register project-local overrides via `require("lazyload").on_override(...)`.                                                                             |
| 11    | `plugin/**/*.lua` loads alphabetically    | Each file calls `vim.pack.add()`, `registry.add()`, and queues setup via `lazyload.on_vim_enter()` — or wraps it in a `FileType`/keymap/command closure. |
| 18    | `VimEnter` fires                          | `lazyload.lua` drains `vim_enter_queue` (sync inline, async via `vim.schedule`), then schedules `override_queue`.                                        |
| later | Scheduler ticks forward                   | FIFO execution: vim-enter async → project overrides.                                                                                                     |

Project-local overrides get the last word — which is the behavior most people
intuitively expect from exrc but which Neovim doesn't provide out of the box.

### Future development

So what can we expect from `vim.pack`, moving forward?

- [`vim.pack` improvements](https://github.com/neovim/neovim/issues/34763)
- [`vim.pack` lazy loading](https://github.com/neovim/neovim/issues/35562)

On the topic of providing lazy-loading via `vim.pack`, the stance by
[Justin M. Keyes](https://github.com/justinmk) is good, I think:

> [!QUOTE]
>
> Putting that burden on every user, instead of solving it once per plugin, is
> an anti-feature. Plugins should solve that, users should not have to
> micro-manage this in their configs. I don't buy the argument that plugin
> authors are incapable of this, yet somehow every user is capable of it.

It makes a lot of sense to at least start here.

### Libraries to help working with `vim.pack`

I should mention there are projects which aim to bridge the gap of wanting more
out of `vim.pack`:

- [lumen-oss/lz.n](https://github.com/lumen-oss/lz.n) — standalone lazy-loading
  library that works with `vim.pack`/`packadd`, providing its own abstraction
  for loading on events, commands, filetypes, and keymaps.
- [BirdeeHub/lze](https://github.com/BirdeeHub/lze) — lazy-loading library (not
  a plugin manager) that works with `packadd` or any plugin manager supporting
  manual lazy-loading.
- [zpack.nvim](https://github.com/zuqini/zpack.nvim) — thin wrapper on top of
  `vim.pack` that adds lazy-loading and lazy.nvim-style declarative specs.

I'm sure someone will build a nice TUI around `vim.pack`, but in the meantime
I'm using
[my own forked variant](https://github.com/fredrikaverpil/dotfiles/blob/main/nvim-fredrik/plugin/pack_ui.lua)
of
[this TUI, originally built by Andreas Schneider](https://git.cryptomilk.org/users/asn/dotfiles.git/tree/dot_config/nvim/lua/plugins/pack-ui.lua).

### Closing comments

I really like how it all came out in the end. It works great for me. That's all
I care about. I now have all capabilities (in some shape or form) that I really
liked about lazy.nvim and depended on.

My Vim journey started with a single `.vimrc` before moving onto Neovim, where I
adopted LazyVim early on. I moved off LazyVim, wrote my own config but with
lazy.nvim as package manager. I borrowed a lot of ideas and code from LazyVim. I
felt I was on a "modern" stack, a modern take on a Neovim config. But when now
looking into adopting `vim.pack`, I had to read up on how Neovim actually works,
how I'm expected to follow certain idioms, dictated by Neovim itself. I really
feel the change. The config became lighter, much more tidy, less complexity. And
my Neovim config actually starts up faster than with lazy.nvim, in 40ms (at 72
plugins). It feels like I cheated the system, but I'm actually just following
the native idioms now. Well, except for the fact that I load almost everything
on `VimEnter`, which I'm sure will make some people frown. But who cares; it
works for me.

Before looking into all this, I couldn't imagine dropping lazy.nvim's notion of
isolating plugins' concerns from each other. Like, passing opts between plugins,
so that I could just delete any plugin's `.lua` file (or set `enabled = false`),
and cleanly having removed all of that plugin's concern from my Neovim config.
When I look at what I have now, with `vim.pack` and mixed concerns, the config
is a lot leaner, more readable, with much less abstractions. I don't have to
cater for all the "what ifs" which e.g. the LazyVim distro has to do, as I only
design this for myself. I even ended up _not_ using the capability to pass opts
between plugins even if the mechanism is there, if I ever need it.

But migrating from `lazy.nvim` to `vim.pack` honestly took more effort than I
initially expected, for getting my Neovim config into what I would consider an
acceptable state. I wouldn't recommend most users do this, unless maybe you just
want to adopt something like my approach, or if you like diving into the rabbit
hole and spend hours on making things work. Maybe wait and see how `vim.pack`
evolves and how others integrate it into their configs before digging in. It's
still early days.

You can inspect my full Neovim configuration over at my
[dotfiles](https://github.com/fredrikaverpil/dotfiles/blob/main/nvim-fredrik).
