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
| **Secondary** | `Commit Mono` | Italics, comments in code, quotes, special emphasis. |

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
- **Links:** No underline by default; underline appears on hover for visual feedback.

### Responsive Design & Media Queries

The site uses three standard breakpoints to adapt gracefully across device sizes, following web design conventions.

**Tablet Breakpoint (`max-width: 768px`):**
Applies to iPad (portrait), tablets, and large devices.
- **Body font size:** `0.95em` (~15.2px)
- **Body padding:** `var(--space-sm) var(--space-lg)` (~10px top/bottom, ~26px left/right).
- **Headings:** Moderately scaled down.
  - h1: 1.8em → 1.7em
  - h2: 1.5em → 1.4em
  - h3: 1.3em → 1.2em
  - h4: 1.2em → 1.05em
  - h5: 1.1em → 1em
  - h6: 1.05em → 0.95em
- **Navigation & Footer:** Padding reduced to `var(--space-xs)`.
- **Profile Header:** Switches from horizontal to vertical layout with centered text.

**Medium Phone Breakpoint (`max-width: 600px`):**
Applies to larger phones and devices between tablet and small phone sizes.
- **Body font size:** `0.90em` (~14.4px)
- **Body padding:** `var(--space-xs) var(--space-md)` (~6px top/bottom, ~16px left/right).
- **Headings:** Moderately scaled down.
  - h1: 1.8em → 1.6em
  - h2: 1.5em → 1.3em
  - h3: 1.3em → 1.1em
  - h4: 1.2em → 1em
  - h5: 1.1em → 0.95em
  - h6: 1.05em → 0.9em
- **Navigation & Footer:** Padding reduced to `var(--space-xs)`.
- **Profile Header:** Maintains vertical layout.

**Small Phone Breakpoint (`max-width: 480px`):**
Applies to iPhone 13, iPhone 14/15, and other compact phones.
- **Body font size:** `0.85em` (~13.6px) for optimal word density (~10-12 words per line).
- **Body padding:** `var(--space-xs) var(--space-md)` (~6px top/bottom, ~16px left/right).
- **Headings:** More aggressively scaled down.
  - h1: 1.8em → 1.5em
  - h2: 1.5em → 1.2em
  - h3: 1.3em → 1em
  - h4: 1.2em → 0.95em
  - h5: 1.1em → 0.9em
  - h6: 1.05em → 0.85em
- **Navigation & Footer:** Padding reduced to `var(--space-xs)`.

**Phone Landscape Breakpoint (`max-height: 430px` + `orientation: landscape`):**
Applies to phones in landscape orientation (e.g., iPhone 13 landscape has ~390px height).
- **Body font size:** `0.90em` (~14.4px)
- **Body padding:** `var(--space-xs) var(--space-md)` (~6px top/bottom, ~16px left/right).
- **Headings:** Same as medium phone breakpoint.
  - h1: 1.6em, h2: 1.3em, h3: 1.1em, h4: 1em, h5: 0.95em, h6: 0.9em
- **Navigation & Footer:** Padding reduced to `var(--space-xs)`.
- **Note:** This breakpoint uses `max-height` instead of `max-width` because landscape phones have wide screens (e.g., 844px on iPhone 13) but limited vertical space.

**Design Principles:**
- Use standard breakpoints (768px, 600px, 480px) for width-based responsiveness.
- Use height-based breakpoints for landscape orientation on phones.
- Font sizes cascade and override from larger to smaller breakpoints.
- Horizontal padding follows industry standards: 16px on phones, 24-26px on tablets.
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
