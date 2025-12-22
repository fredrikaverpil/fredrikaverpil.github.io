# fredrikaverpil.github.io

## Tech Stack & Philosophy

- **Static Site Generator:** [Hugo](https://gohugo.io/) (Fast, Go-based).
- **Go Tooling:** Managed via `go.mod` (Go 1.25+), utilizing Hugo as a Go tool.
- **Theme:** Custom theme, initially based on
  [`hugo-xmin`](https://github.com/yihui/hugo-xmin).
- **Aesthetic:** "Terminal" vibe with
  - [Zenbones](https://github.com/zenbones-theme/zenbones.nvim) color palette
    (Light/Dark mode).
  - **Body/Code:** [Commit Mono](https://commitmono.com/).
  - **Headings/Callouts:** [Maple Mono](https://github.com/subframe7536/maple-font).
- **Search:** [Pagefind](https://pagefind.app/) (Static, low-bandwidth, runs on
  client).
- **Comments:** [Giscus](https://giscus.app/) (Powered by GitHub Discussions).
- **Icons:** [Simple Icons](https://simpleicons.org/) (Inlined SVGs).
- **Automation & Environment:**
  - **Makefile:** Standardizes build, serve, and clean tasks.
  - **direnv:** Manages environment variables via `.envrc`.
- **Philosophy:**
    *   **Zero-JS (Mostly):** JavaScript is only used for progressive enhancement (Search, Dark Mode Toggle, Copy-to-Clipboard). The site remains fully functional without it.
    *   **Self-Contained:** No external font CDNs, no tracking scripts, no submodules. All assets are self-hosted.

## Features

- **Dark Mode:** Automatic system detection + Manual toggle (persisted).
- **Syntax Highlighting:** Adaptive
  [Chroma](https://github.com/alecthomas/chroma) themes (Tango Light / Monokai
  Dark) styled to match Zenbones.
- **Callouts:** GitHub-style alerts (`> [!NOTE]`) rendered via Hugo Render
  Hooks.
- **RSS:** Native Hugo RSS feeds.

## Development

### Prerequisites

- **Hugo:** `brew install hugo` (or
  `go install -tags extended github.com/gohugoio/hugo@latest`)
- **Node/NPM:** For `npx` (Pagefind).

### Commands

| Command      | Description                          |
| :----------- | :----------------------------------- |
| `make serve` | Start local dev server (LiveReload). |
| `make build` | Build production site to `public/`.  |
| `make clean` | Remove build artifacts.              |

### Build Workflow

For development with incremental changes:
```bash
make serve
```

For a clean production build (recommended before deploying):
```bash
make clean && make build
```

**Note:** Always run `make clean` before `make build` to ensure:
- No stale files from previous builds remain in `public/`
- The Pagefind search index is regenerated from scratch
- Duplicate search results are avoided (especially important after renaming files or changing permalinks)

### Deployment

Deploys automatically to GitHub Pages via GitHub Actions on push to `main`.

