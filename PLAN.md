# Project: Minimalist Technical Blog (Hugo)

## 1. Core Vision

- **Target Aesthetic:** Inspired by [antonz.org](https://antonz.org/), extremely
  fast, text-centric, and minimalist.
- **Infrastructure:** Migrate from **MkDocs Material** to **Hugo**.
- **Hosting:** GitHub Pages (`fredrikaverpil.github.io`).
- **Key Principles:** Minimalist maintenance, zero-JS where possible, and high
  performance. The idea is I own all the code (no third-party theme) and this
  also means I want to keep down the maintenance/complexity as much as possible.

## 2. Tech Stack & Features

- **Static Site Generator:** Hugo.
- **Base Theme:** Starting with **Hugo Bear Cub** or **XMin** as a foundation to
  be customized. I'm leaning towards XMin, because I can then add features one
  by one.
- **Search:** **Pagefind** (static indexing, low overhead, "Material-like"
  speed).
- **Dark Mode:** Zero-JS implementation using CSS `prefers-color-scheme`.
- **Interactivity:** **Codapi** for runnable code snippets (via WebAssembly or
  sandbox).
- **Categorization:** Native Hugo Taxonomies (tags and categories) with
  automatic RSS feed generation.
- **Callouts**: GitHub-like callouts, using Render Hooks (Maple Mono).

  ```markdown
  > [!NOTE] This is a note using Maple Mono!
  ```

  ```html
  {{- $type := "" -}} {{- $content := .Inner -}} {{- if (findRE `^\[!(.*)\]`
  .Inner) -}} {{- $match := index (findRE `^\[!(.*)\]` .Inner) 0 -}} {{- $type =
  lower (replace (replace $match "[!" "") "]" "") -}} {{- $content = replace
  .Inner $match "" -}} {{- end -}}

  <blockquote class="callout {{ with $type }}callout-{{ . }}{{ end }}">
    {{- if $type -}}
    <div class="callout-title">{{ $type | humanize }}</div>
    {{- end -}}
    <div class="callout-content">{{ $content | markdownify }}</div>
  </blockquote>
  ```

  ```css
  /* Base Callout Style */
  .callout {
    border-left: 4px solid var(--link-color);
    padding: 1rem;
    margin: 1.5rem 0;
    background: rgba(0, 0, 0, 0.02); /* Subtle tint */
  }

  /* Apply Maple Mono to the title and content */
  .callout-title {
    font-family: "Maple Mono", monospace;
    font-weight: 700;
    text-transform: uppercase;
    font-size: 0.8rem;
    margin-bottom: 0.5rem;
  }

  .callout-content {
    font-family: "Maple Mono", monospace;
    font-style: italic; /* The "handwritten" note feel */
  }

  /* Color variations */
  .callout-warning {
    border-left-color: #e67e22;
  }
  .callout-tip {
    border-left-color: #27ae60;
  }
  ```

- **RSS**: Hugo's native RSS.

## 3. Typography

- **Primary Font:** **Commit Mono** (Variable WOFF2).
- **Implementation:** Self-hosted for performance; using `font-display: swap`.
- **Pairing Idea:** Potentially pairing **Maple Mono** (blogquote, headers) with
  **Commit Mono** (body, code) for a "technical publication" feel, or going 100%
  Mono for the "terminal" aesthetic.

## 4. Syntax Highlighting

- **Engine:** Hugo’s built-in **Chroma**.
- **Strategy:** Server-side highlighting (zero client-side JS) for 90% of
  snippets.
- **Progressive Enhancement:** Use **Codapi** snippets only on specific pages
  (triggered via Front Matter `use_codapi: true`).

## 5. Deployment Workflow

- **CI/CD:** GitHub Actions.
- **Build Steps:** 1. Hugo Build.

## 6. Community and comments

- Giscus

2. Pagefind Indexing.
3. Deploy to GitHub Pages.
