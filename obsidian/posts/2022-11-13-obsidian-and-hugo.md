---
title: Obsidian, Hugo an GitHub pages
tags: [github-pages, hugo, python]
draft: false
summary: This is my journey to manage my personal website in Obsidian, export it into Hugo and publish onto GitHub pages.

# PaperMod
ShowToc: true
TocOpen: true

updated: 2022-11-16T08:50:35+01:00
created: 2022-11-13T01:00:00+01:00
---


## Features (wish list)

- Site structure:
	- [x] Blog entries, managed via Obsidian.
	- [ ] Tech notes / cheat-sheets, managed via Obsidian.
	- [x] An "about me"-page, managed via Obsidian.
	- [ ] Figure out static/images folder/files location.
	- [ ] Non-MVP: Rewrite dragged-in non-markdown file links (images etc).
- [x] Ability to embed gists, but using custom CSS to make them look more integrated.
- [ ] [Custom GitHub Gists](https://codersblock.com/blog/customizing-github-gists/).
- [ ] Simple analytics
	- [ ] [umami](https://umami.is/docs/getting-started)
	- [ ] [plausible](https://github.com/plausible/analytics)
- [x] Search (might make tags unnecessary).
- [x] Tags.
- [ ] [Callouts](https://help.obsidian.md/How+to/Use+callouts).
- [ ] [Mermaid)](https://hugo-book-demo.netlify.app/docs/shortcodes/mermaid/) charts?
- [x] Code blocks with syntax highlighting.
- Ability to somehow showcase .ipynb (Jupyter notebooks).
	- [x] Could use .ipynb gists.
	- [ ] Could run pyscript to make it interactive.
- [ ] GitHub-powered commenting system.
	- [ ] [utterances](https://utteranc.es/)
- [x] Use a popular static site generator (markdown) offering a wide range of themes I can select from.
- [x] Use a CLI tool to export from Obsidian to this static site generator's expected format.
- [x] Use GHA to automatically perform the export/convert and publishing onto my personal GitHub pages website.
- [ ] Deployment settings
	- [ ] [Obsidian2Hugo exporter in go - Today I Learned (task2.net)](https://task2.net/posts/2022-01-10-obsidian2hugo-exporter/2022-01-10-obsidian2hugo-exporter/)
- [x] Page template (to add frontmatter automatically): Templater community plugin.
- [x] Don't copy drafts.
- [ ] export.py:
	- [ ] Skip export if draft=true
	- [ ] Shortcode for YouTube.
	- [ ] Shortcode for Vimeo.
- [ ] Add git submodule for theme.

## Links

- [Variables | Front Matter | PaperMod (adityatelange.github.io)](https://adityatelange.github.io/hugo-PaperMod/posts/papermod/papermod-variables/)

## Testing grounds

### Code block: python

```python
import platform
print(platform.processor())
```
### Code block: jupyter

```jupyter
import this
```

```jupyter
import sys

sys.version
```

### Embed gist

```gist
fredrikaverpil/f225cdd92c9c253c8851316e4ef99a9a
```

```gist
fredrikaverpil/0cde09c624824ebafe0cb94a6cca9e1e#normalize_timedelta.py
```

