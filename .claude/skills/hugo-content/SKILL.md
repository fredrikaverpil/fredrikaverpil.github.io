---
name: hugo-content
description: >
  This skill should be used when writing or editing blog posts, creating new
  content pages, using Hugo shortcodes (codapi, youtube, vimeo), writing
  callouts/alerts, adding code blocks, working with front matter, or authoring
  any Markdown content for the Hugo site. Also use when working with render
  hooks, taxonomy, or content organization.
---

# Hugo Content Authoring

Use when writing, editing, or reviewing blog post content and Markdown files.

## Blog Post Setup

### File Naming

Posts go in `content/blog/` as `YYYY-MM-DD-slug.md`. The date prefix keeps
files sorted; the slug becomes the URL path.

### Front Matter

```yaml
---
title: "Post Title"
date: YYYY-MM-DD
tags: ["tag1", "tag2"]
categories: ["category"]
---
```

Optional: `draft: true` to hide from production builds (visible with `make serve`).

## Code Blocks

Use fenced code blocks with language identifier. **Use `golang` not `go`** to
avoid GDScript misdetection.

Full list of language identifiers:
[Hugo Syntax Highlighting Languages](https://gohugo.io/content-management/syntax-highlighting/#languages)

Implementation: `layouts/_default/_markup/render-codeblock.html` (adds copy button).

## Callouts / Alerts

GitHub-style blockquote alerts:

```markdown
> [!NOTE]
> Informational note.

> [!TIP]
> Helpful advice.

> [!WARNING]
> Potential issue.

> [!IMPORTANT]
> Critical information.

> [!CAUTION]
> Dangerous action warning.

> [!EXAMPLE]
> Example content.

> [!QUOTE]
> Attributed quotation.

> [!INFO]
> General information (alias for NOTE).
```

Implementation: `layouts/_default/_markup/render-blockquote.html`.

## Shortcodes

### Interactive Code (Codapi)

Run code in-browser via WebAssembly. Scripts load only when shortcode is used.

```markdown
{{</* codapi sandbox="python" */>}}
print("Hello, World!")
{{</* /codapi */>}}
```

Supported sandboxes: `javascript`, `fetch`, `python`, `ruby`, `lua`, `php`, `sqlite`.

Implementation: `layouts/shortcodes/codapi.html`.

### Video Embeds

```markdown
{{</* youtube VIDEO_ID */>}}
{{</* vimeo VIDEO_ID */>}}
```

### Footnotes

Standard Goldmark syntax:

```markdown
Text with a footnote[^1].

[^1]: Footnote content.
```

## Links

GitHub URLs auto-format with icons via render hook:
- `github.com/user` renders as `@user`
- `github.com/owner/repo` renders as `owner/repo`

Implementation: `layouts/_default/_markup/render-link.html`.

## Content Organization

- Posts: `content/blog/YYYY-MM-DD-slug.md`
- About page: `content/about.md`
- Taxonomies: `tags` and `categories` (defined in `hugo.toml`)
- Permalinks: `/blog/:year/:month/:day/:slug/`
- Pagination: 10 posts per page

## Draft Test Post

`content/blog/hello-hugo.md` (`draft: true`) showcases all theme features.
Use as reference for syntax. Never published.

## Preview

```bash
make serve       # Dev server with drafts
make build       # Production build
make linkcheck   # Check for broken internal links
```
