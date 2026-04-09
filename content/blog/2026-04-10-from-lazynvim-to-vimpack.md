---
title: "From lazy.nvim to vim.pack"
date: 2026-04-10
draft: true
tags: ["neovim", "lua"]
categories: []
---

I've been using the excellent [lazy.nvim](https://github.com/folke/lazy.nvim)
package manager for more than three years now, and I've been super happy with
it. But with Neovim v0.12.0, `vim.pack` was shipped: a built-in (but still
experimental) plugin manager that manages plugins using Git, with no third-party
dependencies required, implemented by
[Evgeni Chasnovski](https://github.com/echasnovski) (see
[neovim/neovim#34009](https://github.com/neovim/neovim/pull/34009)).

Plugins are stored in a dedicated directory — `site/pack/core/opt` under
Neovim's data path — and a lockfile (`nvim-pack-lock.json`) in your config
directory tracks exact revisions. The API is intentionally small:
`vim.pack.add()`, `vim.pack.update()`, `vim.pack.del()`, and `vim.pack.get()`.

A minimal setup looks like this:

```lua
-- init.lua
vim.pack.add({
  "https://github.com/lewis6991/gitsigns.nvim",
  "https://github.com/folke/which-key.nvim",
})

require("gitsigns").setup()
require("which-key").setup()
```

There are some options to `vim.pack.add` like defining a ref or branch, of
course. The full reference is available in `:h pack.txt` or in the
[online docs](https://neovim.io/doc/user/pack.html#vim.pack).

But essentially, that's it. On first start, Neovim clones missing plugins and
makes them available immediately after the `add()` call returns.

> [!TIP]
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

However, if you, like me, have a good deal of plugins installed, you might miss
features from lazy.nvim such as lazy-loading, defining load order of plugins
while passing around opts from plugin to plugin, build commands or local plugin
loading during development. This blog post aims to outline what I've done in my
personal Neovim config to work around this.

## Lazy-loading with `vim.pack`

In essence, what we are talking about here is we want to defer the `require()`
calls to the plugin to when we really need it.

### Neovim's stance on lazy-loading of plugins

Neovim's own plugin development guide (`:h lua-plugin`,
[online](https://neovim.io/doc/user/lua-plugin.html)) is very clear on this (and
has implementation examples):

> [!QUOTE]
>
> Plugins should arrange their "lazy" behavior once, instead of expecting every
> user to micromanage it.

In other words, a well-written plugin already defers its heavy work internally,
and Neovim explicitly notes there is no performance benefit in users defining
lazy-loading entrypoints themselves. That said, not all plugins are born equal
and `vim.pack` offers no lazy-loading machinery of its own. So if you want to
control startup timing with `vim.pack`, you have to wire it up yourself.
Fortunately, fellow minimalists do not have to reach for YAP (Yet Another
Plugin) as it's quite straight forward to fix this ourselves.

### FileType autocmd

The simplest form of lazy-loading. For plugins that are only relevant to a
specific filetype, wrap `require()` + `.setup()` in a `FileType` autocmd with
`once = true`. The plugin is on the runtimepath from `vim.pack.add()`, but its
modules are never loaded until you actually open a file of that type:

```lua
-- plugin/lang/csv.lua
vim.pack.add({
  { src = "https://github.com/hat0uma/csvview.nvim" },
})

vim.api.nvim_create_autocmd("FileType", {
  pattern = "csv",
  once = true,
  callback = function()
    require("csvview").setup()
  end,
})
```

The `once = true` flag ensures `setup()` only runs on the first CSV file you
open — subsequent CSV buffers skip it entirely. This is ideal for niche plugins
(CSV viewers, log highlighters, schema stores) that most sessions never need.

### Keymap function

For plugins that you only use on demand, like a note-taking integration, move
`require()` inside a keymap callback. A guard variable ensures `setup()` only
runs once:

```lua
-- plugin/obsidian.lua
vim.pack.add({
  { src = "https://github.com/obsidian-nvim/obsidian.nvim" },
})

local initialized = false

local function init()
  if initialized then return end
  initialized = true
  require("obsidian").setup({ ... })
end

vim.keymap.set("n", "<leader>nf", function()
  init()
  vim.cmd("Obsidian quick_switch")
end, { desc = "Notes: search filenames" })

vim.keymap.set("n", "<leader>nn", function()
  init()
  vim.cmd("Obsidian new")
end, { desc = "Notes: new note" })
```

The first keymap press calls `setup()` and then runs the command. Every
subsequent press skips `setup()` entirely.

The same pattern works for user commands. Register a proxy `:Obsidian` that
deletes itself, initializes the plugin (which re-registers the real `:Obsidian`
command), and then passes the arguments through:

```lua
vim.api.nvim_create_user_command("Obsidian", function(opts)
  vim.api.nvim_del_user_command("Obsidian")
  init()
  vim.cmd("Obsidian " .. opts.args)
end, { nargs = "*" })
```

This lets you type `:Obsidian new` or `:Obsidian quick_switch` without the
plugin being loaded until the first invocation.

### Autocommand on event

For plugins that need to be set up at startup but don't need to block the first
paint, autocmds are the right tool. The idea is to queue setup callbacks behind
`VimEnter` or `UIEnter` instead of running them immediately.

A small `defer.lua` module handles this with a queue that supports both
synchronous and async (fire-and-forget) callbacks:

```lua
-- lua/defer.lua
local M = {}

local vim_enter_queue = {}

local function drain(queue)
  for _, entry in ipairs(queue) do
    if not entry.sync then
      vim.schedule(entry.fn)
    end
  end
  for _, entry in ipairs(queue) do
    if entry.sync then
      entry.fn()
    end
  end
end

vim.api.nvim_create_autocmd("VimEnter", {
  once = true,
  callback = function()
    drain(vim_enter_queue)
    vim_enter_queue = nil
  end,
})

--- Run at VimEnter. Async by default. Pass { sync = true } to run synchronously.
function M.on_vim_enter(fn, opts)
  local sync = opts and opts.sync or false
  if vim_enter_queue then
    table.insert(vim_enter_queue, { fn = fn, sync = sync })
  elseif sync then
    fn()
  else
    vim.schedule(fn)
  end
end

return M
```

The `drain()` function does two passes: first it schedules all async callbacks
via `vim.schedule()` (non-blocking), then it runs all synchronous callbacks.
Callbacks are async by default — pass `{ sync = true }` for the rare plugin that
must be fully set up before the UI draws (like a statusline). Everything else
runs on separate event loop ticks without blocking each other or the UI.

Then you defer the plugin from loading like so:

```lua
-- plugin/blink.lua
vim.pack.add({
  { src = "https://github.com/Saghen/blink.cmp", version = vim.version.range("1.*") },
})

require("defer").on_vim_enter(function()
  require("blink.cmp").setup({ ... })
end)
```

See `:h initialization`
[online](https://neovim.io/doc/user/starting.html#initialization) for the full
Neovim startup sequence, and `:h autocmd.txt`
[online](https://neovim.io/doc/user/autocmd.html) for all available events.

[`VimEnter`](https://neovim.io/doc/user/autocmd.html#VimEnter) fires after all
`plugin/` files have been sourced, but before the UI attaches.
[`UIEnter`](https://neovim.io/doc/user/autocmd.html#UIEnter) fires when the UI
connects (or when the built-in TUI starts), after `VimEnter`.

Callbacks registered after `VimEnter` has already fired execute immediately (or
via `vim.schedule` for async), so this pattern is safe to use from any `plugin/`
file regardless of load order.

> [!TIP]
>
> If you recall the earlier keymap example with `:Obsidian`, you might now
> realize we could just do this:
>
> ```lua
> vim.pack.add({
>   { src = "https://github.com/obsidian-nvim/obsidian.nvim" },
> })
>
> require("defer").on_vim_enter(function()
>   require("obsidian.nvim").setup({ ... })
> end)
>
> vim.keymap.set("n", "<leader>nf", function()
>   vim.cmd("Obsidian quick_switch")
> end, { desc = "Notes: search filenames" })
> ```
>
> And there would actually be no need for the command proxy which enabled
> `:Obsidian`. With the only tradeoff that the plugin loading would become
> slightly more eager.

## Passing opts from one plugin to another

In lazy.nvim, any plugin spec can declare a dependency and pass opts into it.
lazy.nvim deep-merges all contributions automatically. For example,
[mason](https://github.com/mason-org/mason.nvim) can add itself to
[lualine](https://github.com/nvim-lualine/lualine.nvim)'s extensions list:

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
> I would actually go as far as to say this has been an anti-pattern and against
> the Neovim idioms, at least historically. The first time I stumbled upon this
> kind of "collection" of opts from other plugins was with
> [LazyVim](https://github.com/LazyVim/LazyVim) and lazy.nvim. I suppose it's a
> pattern which was designed with a distro in mind.
>
> But it's also quite appealing with this decentralized approach rather than
> tangling your plugin configurations up in each other. I'm having a real hard
> time adjusting to the `plugin`, `ftplugin`, `after` folders, eventhough I used
> those pre-LazyVim.

With `vim.pack`, there is no dependency graph and no automatic opts merging.
Every `plugin/` file is self-contained and loads in alphabetical order. So how
can we achieve this behavior without lazy.nvim?

### The "registry pattern"

I've solved this with a small shared module that acts as a central coordination
point. Plugin files **register** data into it immediately on load of
`plugin/**/*.lua`, and the same plugin files can also **read** from it in
deferred callbacks, after all files have had their chance to contribute.

Let's look at the registry:

```lua
-- lua/registry.lua
local merge = require("merge")

local M = {
  lsp = {},
  mason = {},
  conform = {},
  lint = {},
  lualine = {},
  -- ... other fields
}

function M.add(spec)
  merge(M, spec)
end

return M
```

Each plugin gets its own namespace ("lsp", "mason", "conform" ...). Plugins with
a `setup(opts)` function store their opts under `.opts`, while other data (like
`mason.ensure_installed` or `lsp.servers`) lives alongside it:

The `merge()` function is a custom deep merge that **appends lists** (instead of
replacing them like `vim.tbl_deep_extend` does) and **recurses into dicts**:

```lua
-- lua/merge.lua
local function merge(base, override)
  for k, v in pairs(override) do
    local bv = base[k]
    if type(v) == "table" then
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
first.

### How it comes together (lazy-loading ❤️ registry pattern)

Now any plugin file can contribute data to any other plugin — regardless of
alphabetical load order — because all contributions happen at load time
(immediate), and all consumption is deferred (`VimEnter`/`UIEnter`).

Mason contributes its lualine extension and reads its own opts from the
registry:

```lua
-- plugin/mason.lua
vim.pack.add({
  { src = "https://github.com/mason-org/mason.nvim" },
})

require("registry").add({
  lualine = { opts = { extensions = { "mason" } } },
})

require("defer").on_vim_enter(function()
  local merge = require("merge")
  local registry = require("registry")

  local opts = { PATH = "append" }
  merge(opts, registry.mason.opts or {})
  require("mason").setup(opts)
end)
```

Lualine reads everything that was contributed:

```lua
-- plugin/lualine.lua
vim.pack.add({
  { src = "https://github.com/nvim-lualine/lualine.nvim" },
})

require("defer").on_vim_enter(function()
  local registry = require("registry")

  local opts = { extensions = { "man", "quickfix" } }
  merge(opts, registry.lualine.opts or {})
  require("lualine").setup(opts)
end)
```

Both files live in `plugin/`. Lualine loads before mason (alphabetically, `l` <
`m`), but that doesn't matter — both `registry.add()` calls run at load time,
and both `setup()` calls are deferred to `VimEnter`, which fires after **all**
`plugin/` files have loaded.

Every consumer follows the same pattern — merge base opts with
`registry.<name>.opts`:

```lua
-- plugin/conform.lua
require("defer").on_vim_enter(function()
  local merge = require("merge")
  local registry = require("registry")

  local opts = {}
  merge(opts, registry.conform.opts or {})
  require("conform").setup(opts)
end)
```

## Build command

Some plugins need a build step after install or update, like compiling a binary,
downloading assets, or updating parsers. In lazy.nvim this is the `build` key.
With `vim.pack`, the equivalent is the `PackChanged` autocmd event.

`PackChanged` fires after a plugin's state has changed. The event data includes:

- `ev.data.kind` — `"install"`, `"update"`, or `"delete"`
- `ev.data.spec` — the plugin's full spec (including `.name`)
- `ev.data.path` — full path to the plugin directory

Here's how nvim-treesitter's `:TSUpdate` maps to a `PackChanged` hook:

```lua
vim.pack.add({
  { src = "https://github.com/nvim-treesitter/nvim-treesitter", branch = "main" },
})

vim.api.nvim_create_autocmd("PackChanged", {
  callback = function(ev)
    if ev.data.spec.name == "nvim-treesitter" then
      vim.cmd("TSUpdate")
    end
  end,
})
```

Another example — a plugin that needs `make` after install or update:

```lua
vim.api.nvim_create_autocmd("PackChanged", {
  callback = function(ev)
    if ev.data.spec.name == "my-plugin" and ev.data.kind ~= "delete" then
      vim.system({ "make" }, { cwd = ev.data.path })
    end
  end,
})
```

## Local plugin development

When developing Neovim plugins, it's been really nice to be able to define
`dev = true` in lazy.nvim, which makes the plugin load from local disk instead
of the remote git repo. We can achieve this quite easily too.

```lua
-- lua/dev.lua

local m = {}

function m.use(opts)
  local dev_path = vim.fn.expand(opts.dev)
  if vim.uv.fs_stat(dev_path) then
    vim.opt.runtimepath:append(dev_path)
  else
    opts.fallback()
  end
end

return m
```

Then let's just switch out the `vim.pack.add` with this:

```lua
require("dev").use({
  dev = "~/code/public/neotest-golang",
  fallback = function()
    vim.pack.add({
      { src = "https://github.com/fredrikaverpil/neotest-golang" },
    })
  end,
})
```

## Closing comments

I use lazy-loading for poorly designed plugins but also so I can pass opts back
and forth between plugins. I find this last part often overlooked when I read
discussions online on lazy-loading with `vim.pack` although I acknowledge that
it isn't necessarily a thing that should be coupled with a package manager. And
it can of course be solved in other, more complex ways too.

One might argue that why not just stick with lazy.nvim if this is what you want
out of your config, and I think that's absolutely fine. But in my case, I want
to leverage what's already built into Neovim first and foremost. With a few
wrapper functions I can achieve what I am missing from the native experience.

It's far from all plugins that I apply lazy-loading to. As I'm writing this I'm
on 73 plugins installed via `vim.pack` (of which 25 uses the lazy-loading
techniques mentioned here) and I load into Neovim in roughly 100ms. You can
inspect my full Neovim configuration over at my
[dotfiles](https://github.com/fredrikaverpil/dotfiles).

So what can we expect from `vim.pack`, moving forward?

- [`vim.pack` improvements](https://github.com/neovim/neovim/issues/35562?issue=neovim|neovim|34763)
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

I should also mention there are projects which aim to bridge the gap of wanting
more out of `vim.pack`:

- [lumen-oss/lz.n](https://github.com/lumen-oss/lz.n) — standalone lazy-loading
  library that works with `vim.pack`/`packadd`, providing its own abstraction
  for loading on events, commands, filetypes, and keymaps.
- [BirdeeHub/lze](https://github.com/BirdeeHub/lze) — lazy-loading library (not
  a plugin manager) that works with `packadd` or any plugin manager supporting
  manual lazy-loading.
- [zpack.nvim](https://github.com/zuqini/zpack.nvim) — thin wrapper on top of
  `vim.pack` that adds lazy-loading and lazy.nvim-style declarative specs.
