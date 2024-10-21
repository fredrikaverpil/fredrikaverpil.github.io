---
date: 2024-10-20
draft: false
authors:
  - fredrikaverpil
comments: true
tags:
  - neovim
  - terminal
  - lua
---

# Session management in Wezterm (without tmux)

![Wezterm workspaces](/static/wezterm/wezterm_workspaces.png)

I've been using [Neovim](https://github.com/neovim/neovim) as my daily
IDE/editor for a bit more than a year now. I've also been using
[tmux](https://github.com/tmux/tmux) to manage sessions, first using
[t-smart-tmux-session-manager](https://github.com/joshmedeski/t-smart-tmux-session-manager)
and later [sesh](https://github.com/joshmedeski/sesh), which supersedes the
former project. What's great about this is that it doesn't matter which terminal
emulator you use, you can still fall back to this kind of session management.

But the draw screen latency (and complexity) of tmux is annoying and I finally
took some time to evaluate the session management capabilities of
[Wezterm](https://github.com/wez/wezterm) (my currently favorite terminal
emulator, which you configure with Lua!) called
"[workspaces](https://wezfurlong.org/wezterm/recipes/workspaces.html)", to see
if this could replace tmux altogether for me.

<!-- more -->

## Replacing tmux with Wezterm's session management

Well, spoiler alert; it works _really_ well and with the help of
[smart_workspace_switcher.wezterm](https://github.com/MLFlexer/smart_workspace_switcher.wezterm),
I can get pretty much the same behavior as I am used to with sesh. They both
leverage [zoxide](https://github.com/ajeetdsouza/zoxide) for fuzzy-searching
recently visited folders.

!!! example "wezterm.lua"

    ```lua
    -- print the workspace name at the upper right
    wezterm.on("update-right-status", function(window, pane)
      window:set_right_status(window:active_workspace())
    end)
    -- load plugin
    local workspace_switcher = wezterm.plugin.require("https://github.com/MLFlexer/smart_workspace_switcher.wezterm")
    -- set path to zoxide
    workspace_switcher.zoxide_path = "/opt/homebrew/bin/zoxide"
    -- keymaps
    table.insert(keys, { key = "s", mods = "CTRL|SHIFT", action = workspace_switcher.switch_workspace() })
    table.insert(keys, { key = "t", mods = "CTRL|SHIFT", action = act.ShowLauncherArgs({ flags = "FUZZY|WORKSPACES" }) })
    table.insert(keys, { key = "[", mods = "CTRL|SHIFT", action = act.SwitchWorkspaceRelative(1) })
    table.insert(keys, { key = "]", mods = "CTRL|SHIFT", action = act.SwitchWorkspaceRelative(-1) })
    ```

    Full `wezterm.lua` source [here](https://github.com/fredrikaverpil/dotfiles/blob/main/wezterm.lua).

I use `Ctrl+Shift+s` to bring up the workspace manager. Then I start typing out
the path I want to open a new session in. Then I can hit that same command again
to jump between workspaces, or I can use `Ctrl+Shift+[` or `Ctrl+Shift+]` to
jump between them more quickly. If I want to view only the currently opened
workspaces, I can use `Ctrl+Shift+t`.

The only thing I'm a little wary about is how wezterm plugins are just read on
the fly from the Internet like this. I might vendor the
`smart_workspace_switcher.wezterm` project into my own dotfiles in the long
term.

I've also set up a custom workspace which is loaded on Wezterm startup, which
goes into my dotfiles repository and starts up Neovim. Using a hotkey
`Ctrl+Shift+d` I can also always jump directly to this workspace.

!!! example "wezterm.lua"

    ```lua
    -- set up workspace to be loaded on startup of wezterm
    wezterm.on("gui-startup", function(cmd)
      local dotfiles_path = wezterm.home_dir .. "/.dotfiles"
      local tab, build_pane, window = mux.spawn_window({
        workspace = "dotfiles",
        cwd = dotfiles_path,
        args = args,
      })
      build_pane:send_text("nvim\n")
      mux.set_active_workspace("dotfiles")
    end)
    -- set up keymap for quickly jumping to this workspace
    table.insert(keys, { key = "d", mods = "CTRL|SHIFT", action = act.SwitchToWorkspace({ name = "dotfiles" }) })
    ```

    Full `wezterm.lua` source [here](https://github.com/fredrikaverpil/dotfiles/blob/main/wezterm.lua).

## Bonus: tabs

![Wezterm tabs](/static/wezterm/wezterm_tabs.png)

I primarily use workspaces with Wezterm now, but it's also convenient to have
tabs around (can be used without workspaces). I have a custom setup which
entails adding the following config to `wezterm.lua`, which is heavily inspired
by
[aaronlifton's wezterm.lua config](https://github.com/aaronlifton/.config/blob/main/.config/wezterm/wezterm.lua).

So, when hitting `Cmd+T`, a new tab shows and is prefixed by a number. Tabs can
then be selected by hitting `Cmd+[number]`.

!!! example "wezterm.lua"

    ```lua
    config.hide_tab_bar_if_only_one_tab = false
    config.use_fancy_tab_bar = false

    local function get_current_working_dir(tab)
      local current_dir = tab.active_pane and tab.active_pane.current_working_dir or { file_path = "" }
      local HOME_DIR = string.format("file://%s", os.getenv("HOME"))

      return current_dir == HOME_DIR and "." or string.gsub(current_dir.file_path, "(.*[/\\])(.*)", "%2")
    end

    wezterm.on("format-tab-title", function(tab, tabs, panes, config, hover, max_width)
      local has_unseen_output = false
      if not tab.is_active then
        for _, pane in ipairs(tab.panes) do
          if pane.has_unseen_output then
            has_unseen_output = true
            break
          end
        end
      end

      local cwd = wezterm.format({
        { Attribute = { Intensity = "Bold" } },
        { Text = get_current_working_dir(tab) },
      })

      local title = string.format(" [%s] %s", tab.tab_index + 1, cwd)

      if has_unseen_output then
        return {
          { Foreground = { Color = "#8866bb" } },
          { Text = title },
        }
      end

      return {
        { Text = title },
      }
    end)
    ```

    Full `wezterm.lua` source [here](https://github.com/fredrikaverpil/dotfiles/blob/main/wezterm.lua).

## Conclusion

The benefit of having sessions management built into the terminal emulator
itself provides quicker feedback and less complexity overall. But this approach
is a lot more dependent on which terminal emulator you're using and what
capabilities are available. Tmux with sesh is a great solution to fall back on,
if e.g. trying out another terminal emulator.

There are also other ways to jump between projects from _within_ Neovim, such as
with
[telescope-project.nvim](https://github.com/nvim-telescope/telescope-project.nvim),
but these kinds of tools rarely solve the problem of supporting e.g.
[direnv](https://direnv.net/) or
[pkgx](https://pkgx.sh/)/[asdf](https://asdf-vm.com/), which are tools that use
shell integration, and executes when you enter a folder (i.e. switching
sessions). To make such tooling work well, you need to keep track of what your
Neovim plugins are doing and which ones need special treatment to play well with
this kind of behavior. I tried this setup out
[in this PR](https://github.com/fredrikaverpil/dotfiles/pull/160) at one point,
but I didn't like the added complexity which I have to maintain.

On leaving tmux behind, Wezterm comes out of the box with the notion of windows
and panes, the ability to split horizontally or vertically and most of the
features I would've missed from leaving tmux (although I do most of these things
inside Neovim instead). Hit `Ctrl+Shift+P` to bring up the command palette and
explore.

I'm delighted to have a terminal emulator which is configurable with lua, the
same language as Neovim itself, as this makes configuring a lot more
customizable than e.g. a json/yaml-configured terminal. But I'm also very
curious on [Ghostty](https://mitchellh.com/ghostty), which seems to be around
the corner from being publicly released, and what sessions management
capabilities it may bring.
