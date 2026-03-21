# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Personal blog and portfolio built with [Hugo](https://gohugo.io/), using a custom theme with a minimalist "Terminal/Code Editor" aesthetic inspired by Plan 9/Acme.

## Commands

| Command           | Description                                           |
| :---------------- | :---------------------------------------------------- |
| `make serve`      | Start local dev server with drafts (LiveReload)       |
| `make build`      | Production build (Hugo + Pagefind indexing)           |
| `make clean`      | Remove `public/` and `resources/`                     |
| `make test`       | Run Go tests in blog post example code                |
| `make linkcheck`  | Check built site for broken internal links (htmltest) |
| `make ci`         | Run tests + build + linkcheck (used in GitHub Actions)|
| `make`            | Clean + build + serve                                 |

Hugo is invoked as a Go tool: `go tool hugo`. Search indexing uses `bunx pagefind`.

## Architecture

- **Hugo config:** `hugo.toml` — site settings, menu, projects list, social links, permalinks
- **Content:** `content/blog/` (posts as `YYYY-MM-DD-slug.md`), `content/about.md`
- **Layouts:** `layouts/` — custom theme (no external theme/submodule)
  - `_default/` — page templates (`single.html`, `list.html`, `terms.html`, `404.html`)
  - `_default/_markup/` — render hooks for blockquotes (callouts), code blocks (copy button), links (GitHub auto-formatting)
  - `partials/` — reusable components (header, footer, SEO, giscus comments, taxonomy cloud)
  - `shortcodes/` — youtube, vimeo, codapi (interactive code blocks)
- **Static assets:** `static/css/style.css` (main stylesheet), `static/js/main.js` (copy-to-clipboard), `static/fonts/`
- **Go modules:** `go.mod` manages Hugo toolchain version

## Key Design Decisions

- **Zero-JS philosophy:** JavaScript only for progressive enhancement (search, dark mode, copy button). Site works without JS.
- **Self-contained:** All assets self-hosted. No external font CDNs or submodules.
- **Pagefind:** Static search index built after Hugo build (`bunx pagefind --site public`).

## CI/CD

GitHub Actions workflow (`.github/workflows/gh-pages.yml`) builds and deploys to GitHub Pages on push to `main`. PRs run the build but skip deployment.
