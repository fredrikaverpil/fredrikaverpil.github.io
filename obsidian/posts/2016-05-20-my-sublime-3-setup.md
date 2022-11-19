---
title: My Sublime Text 3 setup
tags: [python, osx, windows, linux]
draft: false

cover:
  image: "/static/editor/sublime_material.png"
  alt: "Sublime Text 3"
  relative: false # To use relative path for cover image, used in hugo Page-bundles

# PaperMod
ShowToc: false
TocOpen: false

date: 2016-05-20T03:00:12+02:00
---

This is my [Sublime Text 3](https://www.sublimetext.com) setup, with ambitions to create a custom Python IDE.



All of the below assumes you've already installed [Package Control](https://packagecontrol.io), the package manager for Sublime Text.

### General packages

* [`SublimeLinter`](http://www.sublimelinter.com/en/latest/) - Faster than in ST2
* [`BracketHighlighter`](https://github.com/facelessuser/BracketHighlighter) - Bracket and tag highlighter
* [`Sublimerge Pro`](http://www.sublimerge.com) - Great tool for checking diffs

### Python IDE setup

* [`SublimeLinter-flake8`](https://github.com/SublimeLinter/SublimeLinter-flake8) - Flake and PEP8 checking
* [`Jedi - Python Autocompletion`](https://github.com/srusskih/SublimeJEDI) - Also known as SublimeJEDI
* [`requirementstxt`](https://github.com/wuub/requirementstxt) - Syntax highlighting for requirements.txt files

Please note, SublimeLinter-flake8 requires you to `pip install flake8`.

It's possible to use e.g. [Anaconda](http://damnwidget.github.io/anaconda/) for custom build systems, or you can just use `ctrl+b` (or `cmd+b` on OS X) to run the current Python script.

I'm also defining some custom stuff in `Python.sublime-settings`. The easiest way to edit this file is to open `Sublime -> Preferences -> Settings – More -> Syntax-specific – User` while viewing a Python file:

```python
{
  "auto_indent": true,
  "rulers": [
    72,
    79
  ],
  "smart_indent": true,
  "tab_size": 4,
  "translate_tabs_to_spaces": true,
  "use_tab_stops": true
}
```

### JSON

* [`SublimeLinter-JSON`](https://github.com/SublimeLinter/SublimeLinter-json) - JSON linter
* [`Pretty JSON`](https://github.com/dzhibas/SublimePrettyJson) - Pretty JSON

### Docker

* [`Dockerfile Syntax Highlighting`](https://github.com/asbjornenge/Docker.tmbundle) - Yup, that's what it does

### Git

* [`GitGutter`](https://github.com/jisaacks/GitGutter) - shows file changes in the gutter
* [`GitSavvy`](https://github.com/divmain/GitSavvy) - manage the current git repo directly through Sublime Text

### Saltstack

* [`SaltStack-related syntax highlighting and snippets`](https://github.com/saltstack/sublime-text) - Pretty much self-explanatory

### Markdown

* [`Markdown Preview`](https://github.com/revolunet/sublimetext-markdown-preview) - Opens your browser for MD preview

### Maya

* [`MayaSublime`](https://github.com/justinfx/MayaSublime) - Write code in Sublime, execute it in Maya

### General UI improvements

* [`SideBarEnhancements`](https://github.com/titoBouzout/SideBarEnhancements) - adds a bunch of useful commands to the sidebar

### Theme

* [`Material Theme`](https://github.com/equinusocio/material-theme) - Awesome material inspired theme
* [`Material Appbar`](https://github.com/equinusocio/material-theme-appbar) - Awesome title bar/tabs bar replacement

In addition to the settings mentioned previously, I use the following theme settings in `Sublime -> Preferences -> Settings -> User`:

```python
{
  "always_show_minimap_viewport": true,
  "bold_folder_labels": true,
  "color_scheme": "Packages/Material Theme/schemes/Material-Theme.tmTheme",
  "detect_indentation": true,
  "font_size": 12,
  "ignored_packages":
  [
    "Vintage"
  ],
  "indent_guide_options":
  [
    "draw_normal",
    "draw_active"
  ],
  "line_padding_bottom": 3,
  "line_padding_top": 3,
  "material_theme_compact_sidebar": true,
  "material_theme_small_tab": true,
  "material_theme_panel_separator": true,
  "overlay_scroll_bars": "enabled",
  "theme": "Material-Theme.sublime-theme",
  "wide_caret": true
}

```

### Closing comments

The only thing I really miss from [Atom](https://www.atom.io), another great but unfortunately much slower editor, is the git status highlighting of files in the tree view. This gives you a super nice overview of the files you're working on.

I'd like to see a real terminal inside of Sublime Text. To my knowledge, this doesn't exist (yet). There are a few options like [`Terminal`](https://github.com/wbond/sublime_terminal) or [`Terminality`](https://github.com/spywhere/Terminality), but they don't really offer the kind of terminal you'd expect in an IDE.

All settings files can be examined closer in my [dotfiles repository](https://github.com/fredrikaverpil/dotfiles).

I'll update this post whenever my Sublime Text 3 setup changes permanently.

I'd love to hear if you think I've missed out on a crucial package, and especially with Python development in mind. If so, let me know in the comments below!
