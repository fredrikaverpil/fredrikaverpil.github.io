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
jump between them more quickly.

I've also set up a custom workspace which is loaded on Wezterm startup, which
goes into my dotfiles repository and starts up Neovim. Using a hotkey
`Ctrl+Shift+d` I can also always jump directly to this workspace.

!!! example "wezterm.lua"

    ```lua
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
    ```

    Full `wezterm.lua` source [here](https://github.com/fredrikaverpil/dotfiles/blob/main/wezterm.lua).

This is such a breath of fresh air, not having to use tmux for session
management. Wezterm also comes out of the box with ability to split the terminal
horizontally or vertically and most of the features you would miss from leaving
tmux. Hit `Ctrl+Shift+P` to bring up the command palette and explore.

The only thing I'm a little wary about is how wezterm plugins are just read on
the fly from the Internet like this. I might vendor the
`smart_workspace_switcher.wezterm` project into my own dotfiles in the long
term.

## Bonus: tabs

![Wezterm tabs](/static/wezterm/wezterm_tabs.png)

I primarily use workspaces with Wezterm now, but it's also convenient to have
tabs around. I've created my own workflow here which entails adding the
following config to `wezterm.lua`, which is heavily inspired by
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
