---
date: 2024-12-04
draft: false
authors:
  - fredrikaverpil
comments: true
tags:
  - terminal
  - ghostty
---

# Ghostty on macOS 👻 ❤️ 

![Ghostty](/static/ghostty/ghostty-beta.png)

I got private beta access (quite late), and
[Ghostty](https://mitchellh.com/ghostty) is probably going to be out in the
public as v1.0 any day now. Here are my initial impressions, having used it for
a couple of weeks on macOS.

<!-- more -->

## What I like about Ghostty

I spend most of my day in the terminal and in Neovim specifically. Therefore,
the terminal is my workhorse and I'm a little picky on what I need/want.

Ghostty is really snappy. However, I can't really tell a difference in snapniess
from Kitty or Wezterm (when Wezterm is set to `max_fps = 120`) during daily
work.

It's very easy to configure the editor and enable e.g. opacity, background blur
and make that look really flashy (hint, hint for all nerdy YouTubers out there).

I also like that you get a native notification about new versions of Ghostty.
Right now, it will just update whenever there are new commits in the main branch
but in the future, I hope there will be some sort of release notes directly in
this notification.

It's got native macOS title bars that you can customize, but hiding the title
bar comes at a cost; it will disable the ability to use tabs which is really the
only way to manage sessions (other than using e.g.
[tmux](https://github.com/tmux/tmux) or
[zellij](https://github.com/zellij-org/zellij).

I'm not sure exactly what causes it, but fonts render thicker for me in Kitty. I
can't figure out why this is happening, but they look just right in Ghostty (and
Wezterm for that matter).

Most common colorschemes just work out of the box too (I think there's 300+
bundled ones), so if you're a `tokyonight` or `catppuccin` fan, you're of course
golden. Ghostty uses
[iTerm2 themes](https://github.com/mbadolato/iTerm2-Color-Schemes).

Ghostty also sports a number of features which I haven't used yet or had time to
look into yet, like custom shaders and ability to render images in the terminal.
But I would like to look into if I can get inline images in Obsidian markdown
files to render when I open them up in Neovim (using
[obsidian.nvim](https://github.com/epwalsh/obsidian.nvim) and
[render-markdown.nvim](https://github.com/MeanderingProgrammer/render-markdown.nvim)
). I also want to make [presenterm](https://github.com/mfontanini/presenterm)
show high-resolution images, just like how Kitty does it. I'm sure it just comes
down to a configuration setting, as both Ghostty and Kitty uses the
[Kitty graphics protocol](https://sw.kovidgoyal.net/kitty/graphics-protocol/).
And finally, Ghostty is supposed to be able to switch color themes based on OS
dark/light settings. Something I have until summer to figure out. 😉

## Configuration

Ghostty is configured with an ini-like file, kind of like how you configure
Kitty. This is how I like to browse through the available configuration options:

```bash
ghostty +show-config --default --docs | nvim
```

I've placed my configuration file in `~/.config/ghostty/config` and to enable
syntax highlighting, you can grab the files in
`/Applications/Ghostty.app/Contents/Resources/vim/vimfiles` and place these in
your Neovim installation. You can see how I've placed them
[here](https://github.com/fredrikaverpil/dotfiles/tree/main/nvim-fredrik) for
reference.

My configuration is available
[here](https://github.com/fredrikaverpil/dotfiles/blob/main/ghostty.conf).

I've also made Ghostty show the current working directory of tab titles, when I
have Neovim running in the tab. I've achieved this by adding the following to my
Neovim Lua configuration:

```lua
if vim.fn.getenv("TERM_PROGRAM") == "ghostty" then
  vim.opt.title = true
  vim.opt.titlestring = "%{fnamemodify(getcwd(), ':t')}"
end

```

## What I miss from other terminal emulators

### Session management

I wrote a blog post on
[Wezterm's session management](https://fredrikaverpil.github.io/blog/2024/10/20/session-management-in-wezterm-without-tmux/)
and why I prefer to not use tmux as session manager. But in short, I get
slightly annoyed by the apparent screen drawing latency of tmux and I also feel
keyboard input can sometimes be affected. With Kitty I use a custom tab setup,
which is kind of nice too. But ideally, I would like to achieve some sort of
hybrid between these two approaches:

- Show the different sessions at the top of my terminal window, where each
  session is the `$cwd`, which is in my case representative of a git project
  name.
- Quickly move between the projects using hotkeys such as `<C-S-[>` and
  `<C-S-]>` and `Cmd-1..0`.
- Natively hit a keymap which will bring up a
  [zoxide](https://github.com/ajeetdsouza/zoxide)-powered folder/project
  selector, which will upon selection execute Neovim in a new tab.

With Ghostty, I get all of this except the last part with a zoxide-powered
project selector. Instead I have to hit `<Cmd-t` to create a new tab and then
type in `z someproj` or `zi someproj` to select a project. Then I have to hit
enter and finally execute `nvim`.

### Smear cursor / cursor trail

Although a gimmick, I kind of like Kitty's and Neovide's built in cursor smear
(called `cursor_trail` in Kitty) which adds a neat effect when the cursor darts
around in the editor. I actually miss it in Ghostty. Of course, I forget I don't
have it after about 10 seconds...

## Issues

It's early days (it's not even out in the public yet), and I'm sure any issues
will be ironed out over time by [@michellh](https://github.com/mitchellh), the
community and other projects used in tandem with Ghostty.

### Tabs and Aerospace

Since I'm on macOS and using the
[Aerospace](https://github.com/nikitabobko/AeroSpace) tiling window manager,
I've noticed that Ghostty's tabs doesn't work well, as they are treated as
windows and offsets the whole window when adding new tabs:
[nikitabobk/Aoerospace#68](https://github.com/nikitabobko/AeroSpace/issues/68).

However, I've found a workaround, and that is to make Ghostty into a floating
window (like
[this](https://github.com/fredrikaverpil/dotfiles/blob/72f92cc92a98d19227c161e64a2843966ce99254/aerospace.toml#L213-L224)).
It works for me, but this means I can't use the tiling behaviors of Aerospace
with the workspace the Ghostty window resides in, which I'd ideally would like
to do in the long term.

### Lualine jumping around

Neovim is grid based and depending on how you scale the Ghostty window, you
might see padding around the window as there is not room for a full character to
be rendered. This becomes more apparent with
[lualine](https://github.com/nvim-lualine/lualine.nvim), which ideally would be
tightly snapped to the bottom of the terminal window.

I'm not sure how Kitty does it, but manages much better when you change font
size or resize the terminal window, with avoiding rendering empty
spacing/padding below the lualine.

This is an extremely minor annoyance and only affects the aesthetics.

## Conclusion

In summary, I think Ghostty will become my go-to workhorse for serving up Neovim
in which I spend my days working professionally as well as on hobby projects. I
really like it so far!
