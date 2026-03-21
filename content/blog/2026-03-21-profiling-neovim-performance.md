---
title: "Profiling Neovim performance"
date: 2026-03-21
tags: ["neovim"]
---

If you're experiencing sluggishness or stuttering in Neovim, there's a
straightforward way to measure exactly what's going on using
[profile.nvim](https://github.com/stevearc/profile.nvim). This plugin
instruments your Lua code and autocommands, and exports a JSON trace you can
visualize in [Perfetto](https://ui.perfetto.dev/).

## The story

I recently ran into a problem where typing fast in insert mode caused visible
stuttering. Instead of characters appearing one by one, entire partial words
would appear at once. I suspected
[blink.cmp](https://github.com/saghen/blink.cmp) to be at fault, since it was
the auto-completion which was stuttering. But a proper profiling revealed the
real bottleneck: SQLite operations from a completely different plugin
([cmp-go-deep](https://github.com/samiulsami/cmp-go-deep)).

## Debugging setup

Add [profile.nvim](https://github.com/stevearc/profile.nvim) as a lazy-loaded
plugin. You don't want it instrumenting everything on every startup — only when
you're actively investigating a problem.

Here's how I have it set up in my
[init.lua](https://github.com/fredrikaverpil/dotfiles/blob/main/nvim-fredrik/lua/fredrik/init.lua):

```lua
-- profiling with profile.nvim
-- Run with: NVIM_PROFILE=1 nvim       (instrument, then press <F1> to record)
-- Run with: NVIM_PROFILE=start nvim   (record from startup, press <F1> to stop)
-- Press <F1> to start/stop recording and save the profile as JSON.
-- View the profile at https://ui.perfetto.dev/
local should_profile = os.getenv("NVIM_PROFILE")
if should_profile then
  vim.opt.rtp:append(vim.fn.stdpath("data") .. "/lazy/profile.nvim")

  local prof = require("profile")
  -- Instrument autocommands to capture their performance
  prof.instrument_autocmds()

  -- Ignore vim internals to reduce noise
  prof.ignore("vim.*")

  -- Ignore specific blink.cmp components for focused profiling
  -- Remove these lines if you want to profile render/sort performance
  prof.ignore("blink.cmp.completion.windows.render.*")
  prof.ignore("blink.cmp.fuzzy.sort.*")

  -- "start" mode: record from startup (for init.lua perf)
  -- "instrument" mode: instrument now, record later with <F1> (default)
  if should_profile:lower():match("^start") then
    prof.start("*")
  else
    prof.instrument("*")
  end

  -- <F1> toggles recording on/off and prompts to save the profile
  vim.keymap.set("", "<f1>", function()
    if prof.is_recording() then
      prof.stop()
      vim.ui.input({
        prompt = "Save profile to:",
        completion = "file",
        default = "profile.json",
      }, function(filename)
        if filename then
          prof.export(filename)
          vim.notify(string.format("Wrote %s", filename))
        end
      end)
    else
      prof.start("*")
    end
  end)
end
```

There are two modes of operation:

- **`NVIM_PROFILE=1 nvim`** — instruments all Lua code, but doesn't start
  recording yet. Press `<F1>` to start recording, reproduce the sluggish
  behavior, then press `<F1>` again to stop and save.
- **`NVIM_PROFILE=start nvim`** — starts recording immediately on startup,
  useful for profiling your init time. Press `<F1>` to stop.

## Profiling workflow

1. Start Neovim with `NVIM_PROFILE=1 nvim`
2. Press `<F1>` to begin recording
3. Reproduce the slow behavior (e.g. type rapidly in insert mode)
4. Press `<F1>` again to stop recording and save the `profile.json` file
5. Open the JSON file in [Perfetto](https://ui.perfetto.dev/) (successor of the
   [deprecated `chrome://tracing`](https://chromium.googlesource.com/catapult/+/refs/heads/main/tracing/docs/perfetto.md)).
   It processes everything client-side so no data leaves your machine. It also
   supports SQL queries on traces and handles large files well.

The trace viewer gives you a flame chart where you can zoom in on exactly which
functions are taking the most time. In my case, it was immediately obvious — the
SQLite operations from cmp-go-deep were dominating the trace during typing.
