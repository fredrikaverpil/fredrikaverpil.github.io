---
title: 🌂 Obsidian to Hugo
tags: [obsidian]
draft: true
summary: "This is my journey to manage my personal website in Obsidian, export it into Hugo and publish onto GitHub pages."

# PaperMod
ShowToc: true
TocOpen: true

date: 2022-12-17T22:10:26+01:00
---


## Other published Obsidian vaults

- https://github.com/xcad2k/cheat-sheets

## Obsidian to Hugo

During the move of my blog from Jekyll onto Hugo and Papermod, I wanted to make some customizations and manage my blog and cheat-sheets from within Obsidian.

### MVP / wish list

- Site structure:
	- [x] Blog entries, managed via Obsidian.
	- [x] Tech notes / cheat-sheets, managed via Obsidian.
	- [x] An "about me"-page, managed via Obsidian.
	- [x] Figure out static/images folder/files location.
- [x] Ability to embed gists, but using custom CSS to make them look more integrated.
- [x] Fonts
	- [x] Mona Sans font: [Mona Sans & Hubot Sans (github.com)](https://github.com/mona-sans)
	- [x] Fira Code for code blocks
- [x] Simple analytics: [umami](https://umami.is/docs/getting-started): [Dashboard | Umami Cloud](https://cloud.umami.is/websites)
- [x] Search (might make tags unnecessary).
- [x] Tags.
- [ ] [Mermaid)](https://hugo-book-demo.netlify.app/docs/shortcodes/mermaid/) charts?
- [ ] Favicon
- [x] Code blocks with syntax highlighting.
- Ability to somehow showcase .ipynb (Jupyter notebooks).
	- [x] .ipynb gists.
	- [ ] Interactive execution (pyscript?)
- [x] GitHub-powered commenting system.
	- [x] [utterances](https://utteranc.es/)
- [x] Use a popular static site generator (markdown) offering a wide range of themes I can select from.
- [x] Use a CLI tool to export from Obsidian to this static site generator's expected format.
- [x] Use GHA to automatically perform the export/convert and publishing onto my personal GitHub pages website.
- [x] Don't copy drafts.
- [ ] export.py:
	- [x] Skip export if draft=true
	- [ ] Shortcode for YouTube.
	- [ ] Shortcode for Vimeo.
	- [ ] [Callouts](https://help.obsidian.md/How+to/Use+callouts).
- [ ] CSS:
	- [x] [Detect light/dark theme and update code blocks](https://discourse.gohugo.io/t/different-syntax-highlighting-styles-for-light-and-dark-theme/38448)
	- [x] [Custom GitHub Gist syntax highlighting](https://codersblock.com/blog/customizing-github-gists/) 
	- [x] [Brian Wigginton - Hugo Chroma Syntax Highlighting Dark/Light Mode (bwiggs.com)](https://bwiggs.com/posts/2021-08-03-hugo-syntax-highlight-dark-light/)
	- [ ] [Tweak theme colors](https://github.com/adityatelange/hugo-PaperMod/discussions/645
	- [ ] Style comments
	- [ ] Style gists
- [ ] Add git submodule for theme (need to decide on theme!)
- [ ] Add link to old disqus comments for some pages?
- [ ] Insert frontmatter template on creation of page: [Insert front matter template automatically at file creation time - Resolved help - Obsidian Forum](https://forum.obsidian.md/t/insert-front-matter-template-automatically-at-file-creation-time/35351)
- [ ] Page template (to add frontmatter automatically): Templater community plugin.
- [ ] Use `lastMod` for the cheat-sheets:
	- [ ] [Use Lastmod with PaperMod | Jackson Lucky](https://www.jacksonlucky.net/posts/use-lastmod-with-papermod/)
	- [ ] [Sorting pages by last modified date in Hugo (echorand.me)](https://echorand.me/posts/hugo-reverse-sort-modified/)
- [x] Store Vault in iCloud, clone the github-io repo in, use "Obsidian git" plugin to manage uploads via "Create backup".

### Fine-tuning / other stuff

- [ ] [Obsidian2Hugo](https://task2.net/posts/2022-01-10-obsidian2hugo-exporter/2022-01-10-obsidian2hugo-exporter/)


## Links

- [Variables | Front Matter | PaperMod (adityatelange.github.io)](https://adityatelange.github.io/hugo-PaperMod/posts/papermod/papermod-variables/)

## Testing grounds

### Code block: python

```python
import platform
print(platform.processor())
```

### Code block: jupyter (not interactive yet)

Uses the obsidian-jupyter plugin.

```jupyter
import this
```

```jupyter
import sys

sys.version
```

### Embed gist (not themed properly yet)

#### Jupyter notebook

```gist
fredrikaverpil/f225cdd92c9c253c8851316e4ef99a9a
```

#### Python script

```gist
fredrikaverpil/0cde09c624824ebafe0cb94a6cca9e1e#normalize_timedelta.py
```


## Blockquotes

> Human beings face ever more complex and urgent problems, and their effectiveness in dealing with these problems is a matter that is critical to the stability and continued progress of society. \- Doug Engelbart, 1961

## Callouts (does not work yet)

> [!INFO]
Here's a callout block. Supports markdown, images etc. 
>```python
> print("including code blocks")
> ``` 

## Markdown

I'm disabling wikilinks and will use regular markdown links throughout.

Plain link: [Format your notes - Obsidian Help](https://help.obsidian.md/How+to/Format+your+notes)

Image:

![My memoji](fredrikaverpil.github.io/obsidian/static/memoji.png)