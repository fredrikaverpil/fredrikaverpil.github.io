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

### Responsive Design & Media Queries

The site is mobile-first and adapts gracefully to different screen sizes.

**Mobile Breakpoint (`max-width: 600px`):**
- **Body font size:** Reduced to `0.9em` (~14.4px) for better word density on narrow screens.
- **Body padding:** Reduced from `var(--space-lg)` (~26px) to `var(--space-sm)` (~10px).
- **Headings:** Scaled down proportionally to maintain visual hierarchy.
  - h1: 1.8em → 1.6em
  - h2: 1.5em → 1.3em
  - h3: 1.3em → 1.1em
  - h4: 1.2em → 1em
  - h5: 1.1em → 0.95em
  - h6: 1.05em → 0.9em
- **Navigation:** Padding reduced from `var(--space-sm)` to `var(--space-xs)`.
- **Footer:** Padding reduced to `var(--space-sm)`.
- **Profile Header:** Switches from horizontal to vertical layout with centered text.

**Design Principles:**
- Use `max-width` constraints for readability on large screens, but ensure they don't limit mobile width.
- Always reduce padding on mobile devices to maximize content area.
- Scale typography proportionally; avoid drastic changes that break hierarchy.
- Test on actual devices; use browser DevTools to simulate different viewport sizes.

## 4. Components

### Callouts / Alerts

GitHub-style Markdown alerts (`> [!NOTE]`) are rendered via custom Hugo render
hooks (`layouts/_default/_markup/render-blockquote.html`).

- **Types:** `NOTE`, `TIP`, `WARNING`, `IMPORTANT`, `EXAMPLE`, `QUOTE`.
- **Style:** Colored left border based on type; neutral background.

### Footnotes

Footnotes are supported using the standard Goldmark syntax: `[^1]`.

- **Usage:** Place the marker `[^1]` in the text and define it at the bottom of
  the file using `[^1]: My footnote content`.
- **Style:** Rendered as small superscript numbers that link to a dedicated
  footnotes section at the end of the post.

### Code Blocks

- **Engine:** Hugo `chroma` syntax highlighting (Adaptive Tango Light / Monokai
  Dark).
- **Style:** Minimalist container without shadows, utilizing background colors
  for contrast.
- **Copy Button:** Custom JS-based button in the top-right corner.

### Navigation & UI

- **Nav Bar:** Simple top navigation defined in `hugo.toml`.
- **Buttons:** Flat style, transparent by default, blue on hover.
- **Taxonomy Clouds:**
  - Displays **Categories** and **Tags** at the top of list pages.
  - **Style:** Two-colored interactive buttons (Name + Count).
  - **Active State:** The current term is highlighted (matching the hover style)
    when viewing its specific list page.
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
