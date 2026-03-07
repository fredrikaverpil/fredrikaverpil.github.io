---
title: "Teaching Claude Code to run commands in Neovim"
date: 2026-03-07
draft: false
tags: ["neovim"]
categories: []
---

Inside a Neovim terminal, the `$NVIM` environment variable points to the parent
Neovim's Unix socket. This means you can query editor state, inspect buffers,
check LSP diagnostics, and even send commands — all through Neovim's
[msgpack-RPC API](https://neovim.io/doc/user/api.html).

What's quite neat here is that since I run Claude Code inside a Neovim terminal
window, I can make Claude use this API to query and hook into the running Neovim
session and much more easily debug issues with my Neovim config, develop
plugins, or help with something _inside_ Neovim.

I wrote a [Claude Code](https://docs.anthropic.com/en/docs/claude-code)
[skill](https://docs.anthropic.com/en/docs/claude-code/skills) that teaches
Claude how to use this RPC interface. You can find it
[in my dotfiles](https://github.com/fredrikaverpil/dotfiles/blob/main/stow/shared/.claude/skills/neovim/SKILL.md).

## What can it do?

With the skill loaded, Claude Code can, among many other things:

- **Read the current buffer path** to understand what file you're looking at
- **Get cursor position** to know where you are in a file
- **List open buffers** to see your working set
- **Query LSP clients** to check what language servers are attached
- **Fetch LSP diagnostics** like warnings and errors from the current buffer
- **Inspect highlight groups** to debug why something looks wrong visually
- **Check keymaps** to find conflicts or verify bindings are set correctly
- **Read option values** to debug unexpected behavior (e.g. `formatoptions`,
  `shiftwidth`)
- **Query autocmds** to trace why something fires or doesn't
- **Inspect treesitter nodes** to debug syntax highlighting or text objects
- **Find plugin source code** by searching Neovim's runtime paths
- **Look up help docs** for built-in features and plugins
- **Navigate lazy.nvim plugins' source code** including dev-mode plugins

## How it works

The skill teaches Claude to connect to the `$NVIM` socket using
`nvim --server "$NVIM" --remote-expr`, which can evaluate any Vimscript
expression or run arbitrary Lua via `luaeval()` — giving access to the entire
`vim.*` namespace. The skill also covers other commands like `--remote-send` for
simulating keystrokes and `--remote` for opening files.

## The NVIM_APPNAME gotcha

One non-obvious issue: when `NVIM_APPNAME` is set (common if you run multiple
Neovim configs), all `nvim --server` commands emit a warning on **stdout**. This
corrupts any parsed output, especially JSON. The fix is to capture the output
first, then filter:

```bash
result=$(nvim --server "$NVIM" --remote-expr 'EXPR') \
  && echo "$result" | grep -v '^Warning: Using NVIM_APPNAME='
```

The skill bakes this pattern into every example so Claude uses it consistently.

## Safety guardrails

The skill explicitly instructs Claude to:

- **Never** send `:q`, `:qa`, or other destructive commands without confirmation
- **Never** modify buffer contents via RPC without asking first
- **Prefer** `--remote-expr` over `--remote-send`, which simulates typing

## Setting it up

Place the skill file at `~/.claude/skills/neovim/SKILL.md`. Claude Code
[auto-discovers skills](https://docs.anthropic.com/en/docs/claude-code/skills)
from this directory and loads them based on the skill's description field.

I run Claude Code inside Neovim using
[sidekick.nvim](https://github.com/folke/sidekick.nvim), which embeds it in a
split alongside my editor. Combined with this skill, it creates a fully
integrated experience — Claude can query the same editor state I see and
interact with my session directly. You can find my sidekick config
[here](https://github.com/fredrikaverpil/dotfiles/blob/main/nvim-fredrik/lua/fredrik/plugins/sidekick.lua).
