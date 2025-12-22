# Hugo Implementation Guide

This document serves as the comprehensive guide for the **Fredrik Averpil**
blog/portfolio, covering the design system, technical architecture, and Hugo
configuration.

## 1. Core Vision & Philosophy

- **Aesthetic:** Minimalist, "Terminal/Code Editor" vibe. Inspired by the
  utilitarian beauty of **Plan 9 from Bell Labs**, **Acme**, and classic UNIX
  environments.
- **Focus:** Content-first, high readability, low distraction. Form follows
  function.
- **System:** Uses the **Golden Ratio (φ ≈ 1.618)** for spacing and vertical
  rhythm.
- **Philosophy:**
  - **Zero-JS (Mostly):** JavaScript is only used for progressive enhancement
    (Search, Dark Mode Toggle, Copy-to-Clipboard). The site remains fully
    functional without it.
  - **Self-Contained:** No external font CDNs, no tracking scripts, no
    submodules. All assets are self-hosted.

## 2. Tech Stack

- **Static Site Generator:** [Hugo](https://gohugo.io/) (Fast, Go-based).
- **Go Tooling:** Managed via `go.mod`, utilizing Hugo as a Go tool.
- **Search:** [Pagefind](https://pagefind.app/) (Static, low-bandwidth, runs on
  client).
- **Comments:** [Giscus](https://giscus.app/) (Powered by GitHub Discussions).
- **Icons:** [Simple Icons](https://simpleicons.org/) (Inlined SVGs).

## 3. Design System

### Typography

We use a dual-monospace font stack to reinforce the engineering aesthetic.

| Role          | Font Family   | Usage                                                |
| :------------ | :------------ | :--------------------------------------------------- |
| **Primary**   | `Commit Mono` | Body text, headings, code, UI elements.              |
| **Secondary** | `Maple Mono`  | Italics, comments in code, quotes, special emphasis. |

- **Font Weights:**
  - Light Mode: 350 (legibility)
  - Dark Mode: 200 (light bleed reduction)
  - Headings/Bold: 700

### Color System (Zenbones)

The palette offers soft, warm contrast.

| Variable    | Description   | Light Mode (`#F0EDEC`) | Dark Mode (`#1C1917`)  |
| :---------- | :------------ | :--------------------- | :--------------------- |
| `--bg`      | Background    | Warm Light Gray        | Deep Warm Black        |
| `--fg`      | Foreground    | Dark Slate             | Light Gray             |
| `--link`    | Links/Accents | Water Blue (`#286486`) | Light Blue (`#6099C0`) |
| `--meta-bg` | UI Elements   | Darker Gray            | Lighter Black          |

### Layout & Spacing

Spacing follows a strict Golden Ratio scale (Base Unit: `1rem` / 16px).

- **Layers:** "Compounding Contrast" strategy—elements layered on top are either
  darker (Light Mode) or lighter (Dark Mode) than the layer below.
- **Borders:** No rounded corners (`border-radius: 0`).

## 4. Components

### Callouts / Alerts

GitHub-style Markdown alerts (`> [!NOTE]`) are rendered via custom Hugo render
hooks (`layouts/_default/_markup/render-blockquote.html`).

- **Types:** `NOTE`, `TIP`, `WARNING`, `IMPORTANT`, `EXAMPLE`, `QUOTE`.
- **Style:** Colored left border based on type; neutral background.

### Code Blocks

- **Engine:** Hugo `chroma` syntax highlighting (Adaptive Tango Light / Monokai
  Dark).
- **Style:** Minimalist container with a "dithered" shadow effect.
- **Copy Button:** Custom JS-based button in the top-right corner.

### Navigation & UI

- **Nav Bar:** Simple top navigation defined in `hugo.toml`.
- **Buttons:** Flat style, transparent by default, blue on hover.
- **Tags:** Two-colored interactive buttons (Name + Count).
- **TOC:** Automatically generated on single post pages.

## 5. Hugo Configuration

- **HTML Unsafe:** Enabled (`markup.goldmark.renderer.unsafe = true`) to allow
  raw HTML.
- **Syntax Highlighting:** `markup.highlight.noClasses = false` (CSS classes
  used).
- **Taxonomies:** Tags and Categories are supported.
- **RSS:** Native Hugo RSS feeds.

## 6. File Structure

- `content/`: Markdown content files.
  - `blog/hello-hugo.md`: A draft post (set to `draft: true`) designed to showcase and test all theme features (code blocks, callouts, etc.). It renders locally via `make serve` but is never published.
- `static/`: Assets (images, CSS, fonts) served as-is.
- `layouts/`: HTML templates.
  - `_default/baseof.html`: Master template.
  - `partials/`: Reusable components (head, footer, etc.).
