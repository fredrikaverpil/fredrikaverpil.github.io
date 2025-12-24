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
- **Design System:** Uses the **Golden Ratio (φ ≈ 1.618)** for spacing and
  vertical rhythm.
- **Philosophy:**
  - **Zero-JS (Mostly):** JavaScript only for progressive enhancement (Search,
    Dark Mode Toggle, Copy-to-Clipboard). Site remains fully functional without
    it.
  - **Self-Contained:** No external font CDNs, no tracking scripts, no
    submodules. All assets self-hosted.

## 2. Tech Stack

- **Static Site Generator:** [Hugo](https://gohugo.io/) (Fast, Go-based).
- **Go Tooling:** Managed via `go.mod`, utilizing Hugo as a Go tool.
- **Search:** [Pagefind](https://pagefind.app/) (Static, low-bandwidth, runs on
  client).
- **Comments:** [Giscus](https://giscus.app/) (Powered by GitHub Discussions).
- **Icons:** [Simple Icons](https://simpleicons.org/) (Inlined SVGs).
- **Analytics:** [Umami](https://umami.is/) (Self-hosted, privacy-focused).

## 3. Design System

### Typography

Dual-monospace font stack: **Commit Mono** (primary), **Maple Mono** (italics).

**Implementation:** `static/css/style.css`, `layouts/partials/header.html`.

### Color System (Zenbones)

Zenbones-based palette: warm light gray (light mode), deep warm black (dark
mode). Adapts to system preference or manual toggle.

**Implementation:** `static/css/style.css`.

### Layout & Spacing

Golden Ratio scale (`1rem` × φ factors). Compounding contrast (layered elements
progressively darker/lighter). No rounded corners. Links underline on hover.

**Implementation:** `static/css/style.css`.

### Responsive Design & Media Queries

Four breakpoints: `768px` (tablet), `600px` (medium phone), `480px` (small
phone), `430px` height + landscape (phone horizontal).

**Implementation:** See media queries in `static/css/style.css`.

## 4. Components

### Callouts / Alerts

GitHub-style Markdown: `> [!TYPE]` (NOTE, TIP, WARNING, IMPORTANT, CAUTION,
EXAMPLE, QUOTE, INFO).

**Implementation:** `layouts/_default/_markup/render-blockquote.html`,
`static/css/style.css`.

### Footnotes

Standard Goldmark syntax: `[^1]`.

### Code Blocks

Chroma syntax highlighting with copy button (top-right, visual feedback).

**Implementation:** `layouts/_default/_markup/render-codeblock.html`,
`static/js/main.js`.

### GitHub Link Formatting

Auto-formats GitHub URLs with icons: `github.com/user` → `@user`,
`github.com/owner/repo` → `owner/repo`.

**Implementation:** `layouts/_default/_markup/render-link.html`,
`static/css/github-link.css`.

### Video Embeds

Shortcodes: `{{< youtube VIDEO_ID >}}`, `{{< vimeo VIDEO_ID >}}`.

**Implementation:** `layouts/shortcodes/`.

### Dark Mode Toggle

Manual theme toggle in nav bar. Persists to `localStorage`.

**Implementation:** `layouts/partials/header.html`.

### Navigation & UI

Top bar (menu in `hugo.toml`), flat buttons, taxonomy clouds, auto-generated
TOC.

### Search (Pagefind)

Static search in nav bar with dropdown results. Click-outside-to-close.

**Implementation:** `layouts/partials/header.html`, `static/css/style.css`.

## 5. Hugo Configuration

Key settings in `hugo.toml`: HTML unsafe enabled, CSS-based syntax highlighting,
tags/categories taxonomies, custom RSS basename (`feed_rss_created.xml`),
date-based permalinks, 10 posts per page.

## 6. SEO & Meta

Dynamic meta tags, Open Graph, Twitter Cards, JSON-LD structured data
(`layouts/partials/seo_schema.html`), robots.txt, RSS feed at
`/feed_rss_created.xml`.

## 7. JavaScript

Progressive enhancements (fully functional without JS): copy-to-clipboard
(`static/js/main.js`), dark mode toggle, search click-outside, FOUC prevention
(`layouts/partials/header.html`).

## 8. File Structure

- `content/`: Markdown content files.
  - `blog/`: Blog posts organized by date.
  - `blog/hello-hugo.md`: A draft post (set to `draft: true`) designed to
    showcase and test all theme features (code blocks, callouts, etc.). It
    renders locally via `make serve` but is never published.
  - `about.md`: About page.
- `static/`: Assets (images, CSS, fonts) served as-is.
  - `css/style.css`: Main stylesheet (1300+ lines).
  - `css/github-link.css`: GitHub link-specific styling.
  - `fonts/`: Self-hosted font files (CommitMono-Variable, MapleMono-Italic in
    WOFF2 format).
  - `js/main.js`: Copy-to-clipboard functionality.
- `layouts/`: HTML templates.
  - `_default/single.html`: Blog post template.
  - `_default/list.html`: List page template.
  - `_default/terms.html`: Taxonomy term listing.
  - `_default/404.html`: Custom 404 error page.
  - `_default/profile.html`: Profile page component.
  - `_default/_markup/`: Render hooks.
    - `render-blockquote.html`: Callout/alert rendering with icons.
    - `render-codeblock.html`: Code block wrapper with copy button.
    - `render-link.html`: GitHub link auto-formatting.
  - `partials/`: Reusable components.
    - `header.html`: Head, navigation, search, theme toggle.
    - `footer.html`: Footer with social links.
    - `head_custom.html`: Analytics script injection.
    - `foot_custom.html`: Custom JS loading.
    - `seo_schema.html`: JSON-LD structured data.
    - `giscus.html`: Comment system integration.
    - `taxonomy_cloud.html`: Tag/category cloud rendering.
  - `shortcodes/`: Custom shortcodes.
    - `youtube.html`: YouTube embed wrapper.
    - `vimeo.html`: Vimeo embed wrapper.
  - `index.html`: Homepage template.
- `hugo.toml`: Hugo configuration file.
- `go.mod`, `go.sum`: Go module files for Hugo toolchain.
- `Makefile`: Build automation with targets:
  - `make` or `make all`: Clean + build + serve (recommended for development).
  - `make serve`: Local development server with drafts.
  - `make build`: Production build with minification + Pagefind indexing.
  - `make clean`: Remove generated files (`public/`, `resources/`).
- `README.md`: Project overview and quick start guide.
- `HUGO.md`: This comprehensive implementation guide.

## 9. Build & Development

**Local Development:**

```bash
make  # Clean + build + serve (recommended)
# Or: make serve (serve only, without clean/build)
```

**Production Build:**

```bash
make build  # Hugo build + Pagefind indexing
```

**Manual Commands:**

- Hugo: `go tool hugo server -D` (dev) or
  `go tool hugo --minify --environment production` (prod)
- Search: `bunx pagefind --site public` (after Hugo build)
- Clean: `make clean` (removes `public/` and `resources/`)

**Dependencies:** Hugo (Go modules), Pagefind (bunx/npx).

## 10. Performance & Compatibility

Self-hosted fonts with preloading, `font-display: block`, static search, minimal
JS. Browser-specific fixes for Firefox, WebKit, Safari. Accessibility: semantic
HTML, ARIA labels, keyboard navigation, WCAG AA contrast.
