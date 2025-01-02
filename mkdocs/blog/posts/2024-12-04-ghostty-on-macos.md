---
date: 2024-12-04
draft: false
authors:
  - fredrikaverpil
comments: true
tags:
  - terminal
---

# Ghostty on macOS

![Ghostty](/static/ghostty/ghostty-beta.png)

I got private beta access (quite late), and
[Ghostty](https://mitchellh.com/ghostty) is probably going to be out in its
first public release any day now (update: it was released on 2024-12-26). Here
are my initial impressions, having used it for a couple of weeks on macOS.

!!! note "Amendments 2025-01-02"

    This post has been amended with the following details since it was
    originally posted:

    - Release tracks; stable vs tip.
    - Example config for tabs integration into title bar.
    - Example command to list themes.
    - Example command for setting light vs dark theme which will follow system.
    - Link to [tmux control mode](https://github.com/ghostty-org/ghostty/issues/1935) issue.

<!-- more -->

## Configuration

Ghostty is configured with an ini-like file, kind of like how you configure
[Kitty](https://github.com/kovidgoyal/kitty) . Throughout this blog post, I'll
add example configuration snippets. This is how I like to browse through the
available configuration options:

```bash
ghostty +show-config --default --docs | nvim
```

I've placed my configuration file in `~/.config/ghostty/config` and to enable
syntax highlighting, you can grab the files in
`/Applications/Ghostty.app/Contents/Resources/nvim/site` and place these in your
Neovim installation. You can see how I've placed them
[here](https://github.com/fredrikaverpil/dotfiles/tree/main/nvim-fredrik) for
reference.

My configuration is available
[here](https://github.com/fredrikaverpil/dotfiles/blob/main/ghostty.conf).

I've also made Ghostty show the current working directory of tab titles, when I
have Neovim running in the tab. I've achieved this by adding the following to my
Neovim Lua configuration (but unfortunately I don't know how to permanently set
the tab title to the `$cwd`:

```lua
if vim.fn.getenv("TERM_PROGRAM") == "ghostty" then
  vim.opt.title = true
  vim.opt.titlestring = "%{fnamemodify(getcwd(), ':t')}"
end
```

## What I like about Ghostty

I spend most of my workday in the terminal and in Neovim specifically.
Therefore, the terminal is my workhorse and I'm a little picky about what I
need/want.

Ghostty is really snappy. However, I can't really tell a difference in
snappiness from Kitty or [Wezterm](https://github.com/wez/wezterm) (when Wezterm
is set to `max_fps = 120`) during daily work.

It's very easy to configure the editor and enable e.g. opacity, background blur
and make that look really flashy (hint, hint for all nerdy YouTubers out there
😉).

I also like that you get a native notification about new versions of Ghostty.
Right now, it will just update whenever there are new commits in the main branch
but in the future, I hope there will be some sort of release notes directly in
this notification.

!!! tip "Stable vs tip"

    There are two release tracks. One stable and one "tip" (updates to latest
    commit). You can download them respectively with Homebrew:

    ```bash
    # update formulae
    brew update

    # stable
    brew install --cask ghostty

    # tip
    brew install --cask ghostty@tip
    ```

It's got a native macOS title bar that you can customize. Hiding the title bar
comes at a cost; it will disable the ability to use tabs which is really the
only way to efficiently jump between multiple projects (other than using e.g.
[tmux](https://github.com/tmux/tmux) or
[zellij](https://github.com/zellij-org/zellij)). However, you can integrate the
tabs into the title bar, which makes the window more minimalistic.

!!! example "Integrate tabs into title bar"

    ```ini
    macos-titlebar-style = tabs
    ```
    For more details, see the [`macos-titlebar-style` docs](https://ghostty.org/docs/config/reference#macos-titlebar-style).

I'm not sure exactly what causes it, but fonts render thicker for me in Kitty. I
can't figure out why this is happening, but they look just right in Ghostty (and
Wezterm for that matter).

Most common colorschemes just work out of the box too (I think there are 300+
bundled ones), so if you're a `tokyonight` or `catppuccin` fan, you're of course
golden. Ghostty uses
[iTerm2 themes](https://github.com/mbadolato/iTerm2-Color-Schemes).

!!! tip "List all bundled themes"

    ```bash
    ghostty +list-themes
    ```

Ghostty also sports a number of features which I haven't used yet or had time to
look into yet, like custom shaders and the ability to render images in the
terminal. But I would like to look into if I can get inline images in Obsidian
markdown files to render when I open them up in Neovim (using
[obsidian.nvim](https://github.com/epwalsh/obsidian.nvim) and
[render-markdown.nvim](https://github.com/MeanderingProgrammer/render-markdown.nvim)).
I also want to make [presenterm](https://github.com/mfontanini/presenterm) show
high-resolution images, just like how Kitty does it (update: enabled in
[commit 5e651f6 ](https://github.com/mfontanini/presenterm/commit/5e651f636037c658e21c3cea8b7cf6b7b6ccae25)).
Both Ghostty and Kitty use the
[Kitty graphics protocol](https://sw.kovidgoyal.net/kitty/graphics-protocol/),
which makes this possible. And finally, Ghostty is supposed to be able to switch
color themes based on OS dark/light settings. Something I have until summer to
figure out. 😉

!!! example "Theme follows system"

    ```ini
    theme = light:rose-pine-dawn,dark:rose-pine
    ```

    For more details, see the [`theme` docs](https://ghostty.org/docs/config/reference#theme).

## What I miss from other terminal emulators

### Session management

I wrote a blog post on
[Wezterm's session management](https://fredrikaverpil.github.io/blog/2024/10/20/session-management-in-wezterm-without-tmux/)
and why I prefer to not use tmux as a session manager. In short, I get slightly
annoyed by the apparent screen drawing latency of tmux and I also feel keyboard
input can sometimes be affected. With Kitty I use a custom tab setup, which is
kind of nice too. All of this stems from me having used tmux in the past and I
really like having sessions and windows/tabs as a way to navigate projects.

But ideally, I would like to achieve some sort of hybrid between what I have
today in Wezterm and Kitty (but in Ghostty):

- Show the different sessions at the top of my terminal window, where each
  session is the `$cwd`, which is in my case representative of a git project
  name.
- Within a session, have the ability to branch out into tabs, so I can have
  multiple tabs when I'm in a certain project context.
- Quickly move between the sessions using hotkeys such as `Ctrl + Shift + [` and
  `Ctrl + Shift + ]`. Move between tabs with `Cmd + 1..0`.
- Natively hit a keymap which will bring up a
  [zoxide](https://github.com/ajeetdsouza/zoxide)-powered folder/project
  selector, which will upon selection execute Neovim in a new session, which
  uses the `$cwd` of the desired project path.

With Wezterm, I've got this all working except the first point on showing all
active sessions at the top of the terminal window.

With Kitty, I have only achieved having tabs with the mentioned keymaps, but
also haven't spent a great time digging into this. It does not seem to provide
session management out of the box.
[Kittens](https://sw.kovidgoyal.net/kitty/kittens_intro/) are implied to perhaps
enable some sort of session management
[here](https://github.com/kovidgoyal/kitty/discussions/3190).

With Ghostty, I have to use tabs instead of sessions. What I really miss though,
is having a `zoxide`-powered project selector. Instead, I have to hit `Cmd + t`
to create a new tab and then type in `z someproj` or `zi someproj` to select a
project. Then I have to hit enter and finally execute `nvim`. It's okay, but
this is something I might want to look into. Or, better, the
[tmux control mode](https://github.com/ghostty-org/ghostty/issues/1935) will
hopefully solve this.

With all this said and despite not being "perfect" for me, Ghostty still feels
and works great.

### Cursor trail

Although a gimmick, I kind of like Kitty's and Neovide's built-in cursor
trail/smear (called `cursor_trail` in Kitty) which adds a neat effect when the
cursor darts around in the editor. I actually miss it in Ghostty. Of course, I
forget I don't have it after about 10 seconds...

## Issues

It's early days (it's not even out in the public yet), and I'm sure any issues
will be ironed out over time by [@mitchellh](https://github.com/mitchellh), the
community and other projects used in tandem with Ghostty.

### Tabs and Aerospace

Since I'm on macOS and using the
[Aerospace](https://github.com/nikitabobko/AeroSpace) tiling window manager,
I've noticed that Ghostty's tabs don't work well, as they are treated as windows
and offset the whole window when adding new tabs:
[nikitabobko/Aerospace#68](https://github.com/nikitabobko/AeroSpace/issues/68).

However, I've found a workaround, and that is to make Ghostty into a floating
window (like
[this](https://github.com/fredrikaverpil/dotfiles/blob/72f92cc92a98d19227c161e64a2843966ce99254/aerospace.toml#L213-L224)).
It works for me, but this means I can't use the tiling behaviors of Aerospace
with the workspace the Ghostty window resides in, which I'd ideally like to do
in the long term.

### Lualine jumping around

Neovim is grid-based and depending on how you scale the Ghostty window, you
might see padding around the window as there is not room for a full character to
be rendered. This becomes more apparent with
[lualine](https://github.com/nvim-lualine/lualine.nvim), which ideally would be
tightly snapped to the bottom of the terminal window.

I'm not sure how Kitty does it, but it manages much better when you change font
size or resize the terminal window, avoiding rendering empty spacing/padding
below the lualine.

This is an extremely minor annoyance and only affects the aesthetics.

## Conclusion

In summary, I think Ghostty will become my go-to workhorse for serving up Neovim
in which I spend my days working professionally as well as on hobby projects. I
really like it so far!
