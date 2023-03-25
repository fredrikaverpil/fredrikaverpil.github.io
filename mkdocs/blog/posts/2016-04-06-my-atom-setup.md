---
date: 2016-04-06
tags:
- python
---

# My Atom setup

I'm in love with [Atom](https://atom.io). Despite it being slow on large
files, I still haven't been able to jump over the threshold of learning
[vim](http://www.vim.org) (or [neovim](https://neovim.io)).
I'm way too comfortable with Atom right now. Here's my setup.

<!-- more -->

### Base linter

* `linter` - a base linter for Atom

An important setting here is to decide whether you wish "Lint as you type"
enabled or disabled. For large, unlinted files, you may want to disable this
feature to avoid hiccups.

### IDE-like stuff

* `script` - Run your code from Atom
* `terminal-plus` - Integrated terminal window(s)

### Python

* `linter-flake8` - Python flake8 linter (requires `linter`)

### JSON

* `linter-jsonlint` - JSON linter (requires `linter`)
* `pretty-json` - Format and sort JSON files

### Saltstack

* `language-salt` - Salt states syntax highlighting
* `atom-jinja2` - Jinja2 syntax highlighting

Please note, you will need both `language-salt` and `atom-jinja2` to properly
syntax highlight sls (Salt states) files.

### Docker

* `language-docker` - Syntax highlighting in Dockerfiles
* `linter-docker` - Dockerfile linter (requires `linter`)

### General UI improvements

* `minimap` - Sublime Text-style code preview scroller
* `file-icons` - adds icons to file tree