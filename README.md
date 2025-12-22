# fredrikaverpil.github.io

## Tech Stack & Philosophy

- **Static Site Generator:** [Hugo](https://gohugo.io/) (Fast, Go-based)
- **Theme:** Custom theme, initially based on
  [`hugo-xmin`](https://github.com/yihui/hugo-xmin)
- **Aesthetic:**"Terminal" vibe with
  - [Zenbones](https://github.com/zenbones-theme/zenbones.nvim) color palette
    (Light/Dark mode)
  - **Body/Code:** [Commit Mono](https://commitmono.com/)
  - **Callouts:** [Maple Mono](https://github.com/subframe7536/maple-font)
- **Search:** [Pagefind](https://pagefind.app/) (Static, low-bandwidth, runs on
  client).
- **Philosophy:**
  - **Zero-JS (Mostly):** JavaScript is only used for progressive enhancement
    (Search, Dark Mode Toggle). The site is fully functional without it.
  - **Self-Contained:** No external font CDNs, no tracking scripts, no
    submodules.

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

### Deployment

Deploys automatically to GitHub Pages via GitHub Actions on push to `main`.

