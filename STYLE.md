# Project Style Guide

This document outlines the design system, coding conventions, and stylistic
choices for the **Fredrik Averpil** blog/portfolio.

## 1. Design Philosophy

- **Aesthetic:** Minimalist, "Terminal/Code Editor" vibe. heavily inspired by
  the utilitarian beauty of **Plan 9 from Bell Labs**, **Acme**, and classic
  UNIX environments.
*   **Focus:** Content-first, high readability, low distraction. Form follows function.
*   **System:** Uses the **Golden Ratio (φ ≈ 1.618)** for spacing and vertical rhythm, bringing natural harmony to the stark, mechanical aesthetic.
*   **No Borders:** We generally avoid borders on content blocks (code, cards). Instead, we rely on background color layers and shadows to define structure.
*   **Blocky Aesthetic:** All UI elements must have `border-radius: 0`. No rounded corners.

## 2. Typography

We use a dual-monospace font stack to reinforce the engineering aesthetic.

| Role          | Font Family   | Usage                                                |
| :------------ | :------------ | :--------------------------------------------------- |
| **Primary**   | `Commit Mono` | Body text, headings, code, UI elements.              |
| **Secondary** | `Maple Mono`  | Italics, comments in code, quotes, special emphasis. |

> [!NOTE] We would have preferred **Berkeley Mono** for the primary font, but
> **Commit Mono** is the closest high-quality open-source alternative that
> captures the same spirit.

### Font Weights

- **Light Mode:** 350 (slightly thicker to ensure legibility on light
  backgrounds).
- **Dark Mode:** 200 (thinner, allows light to bleed through slightly less).
- **Headings/Bold:** 700.

## 3. Color System

The color palette is derived from **Zenbones**, offering a soft, warm contrast
rather than harsh black/white.

### Palette Variables (CSS)

| Variable    | Description   | Light Mode (`#F0EDEC`) | Dark Mode (`#1C1917`)  |
| :---------- | :------------ | :--------------------- | :--------------------- |
| `--bg`      | Background    | Warm Light Gray        | Deep Warm Black        |
| `--fg`      | Foreground    | Dark Slate             | Light Gray             |
| `--link`    | Links/Accents | Water Blue (`#286486`) | Light Blue (`#6099C0`) |
| `--meta-bg` | UI Elements   | Darker Gray            | Lighter Black          |

### Theming Strategy

The site supports three modes via CSS variables:

1.  **System Preference:** Default (via `@media (prefers-color-scheme: dark)`).
2.  **Manual Dark:** `body.dark` class.
3.  **Manual Light:** `body.light` class.

**Note:** When adding new theme variables, ensure they are defined in **all
three** locations in `static/css/style.css`:

1.  `:root` (Default/Light)
2.  `@media (prefers-color-scheme: dark)`
3.  `body.dark` & `body.light` overrides.

### Elevation & Layering (Compounding)
We follow a **"Compounding Contrast"** strategy for depth.

*   **Dark Mode:** The base is deep black. Elements layered on top (nav bar, cards) must be **brighter** (lighter) than the layer below. This simulates moving closer to a light source.
*   **Light Mode:** The base is a warm gray. Elements layered on top must be **darker** (closer to black) than the layer below. This creates a "stacking" effect similar to ink on paper.

## 4. Spacing & Layout

Spacing follows a strict Golden Ratio scale.

- **Base Unit:** `1rem` (16px)
- **Scale:**
  - `--space-xs`: ~6px
  - `--space-sm`: ~10px
  - `--space-md`: 16px
  - `--space-lg`: ~26px (Primary content spacing)
  - `--space-xl`: ~42px (Section gaps)
  - `--space-xxl`: ~68px

**Max Width:** Content is capped at `800px` and centered.

## 5. Components

### Callouts / Alerts
We use GitHub-style Markdown alerts. These are rendered via a custom Hugo render hook (`layouts/_default/_markup/render-blockquote.html`).

**Styling Rules:**
*   **Background:** Neutral background (`--meta-bg`).
*   **Border:** Only the **left vertical line** is colored based on the callout type.

**Syntax:**
```markdown
> [!NOTE]
> Useful information.
```

**Supported Types:**
*   `[!NOTE]` / `[!INFO]`
*   `[!TIP]`
*   `[!WARNING]`
*   `[!IMPORTANT]` / `[!CAUTION]`
*   `[!EXAMPLE]`
*   `[!QUOTE]` (Italic text)

### Code Blocks
*   **Rendering:** Hugo `chroma` syntax highlighting.
*   **Background:**
    *   **Default:** Uses `--meta-bg` when placed directly in the body.
    *   **Nested:** Uses `--button-bg` when inside a callout to ensure compounding contrast.
*   **Font:** `Commit Mono` for code, `Maple Mono` for comments (italicized).
*   **Copy Code Button:** Uses the **Flat Button** style (see below).
    *   **Position:** Top-right corner of the code block.

### Flat Buttons
Used for subtle actions where a boxy "layer" is not desired (e.g., the Copy Code button).

*   **Style:** Transparent background by default. No border.
*   **Typography/Icons:** Uses `--fg-dim` to remain unobtrusive.
*   **Hover State:** Follows the global UI element rule (turns **blue** `--link` with `--bg` text/icon).

### Navigation
*   Simple top navigation bar.
*   Links defined in `hugo.toml` under `[menu.main]`.

### Table of Contents (TOC)
Used on single post pages to aid navigation.

*   **Background:** Neutral background (`--meta-bg`).
*   **Padding:** Matches callouts (`var(--space-lg)`).
*   **Border:** The left vertical line must match the **info callout** (4px solid `--info-border`).
*   **Vertical Rhythm:** Spacing between items must remain consistent regardless of indentation. This is achieved by matching nested `ul` top margins with `li` bottom margins.
*   **Typography:** The title uses `Commit Mono` bold and matches callout header sizing. Links follow global link styling.

### Tags (Two-Colored Buttons)
Tags are specialized "two-colored" interactive buttons used for taxonomies.

*   **Structure:** They consist of a term name (left) and a count (right), joined as a single `inline-flex` element.
*   **Colors (Default):**
    *   **Left Side:** Uses `--meta-bg` for background.
    *   **Right Side (Count):** Uses `--border` for background to create a distinct secondary block.
*   **Hover State:** The tag behaves like a unified UI element:
    *   The entire tag turns **blue** (`--link`).
    *   The term name text turns `--bg`.
    *   The count block (right) switches to a high-contrast state (`--fg` background with `--bg` text), reversing the colors for emphasis.

### Cards
Interactive cards displayed in a grid on the homepage (typically for projects).

*   **Layering:** They sit on the main background, so they use `--button-bg` (Darker in Light Mode, Lighter in Dark Mode).
*   **Default State:** Neutral background, Blue title (`--link`), Muted description (`--fg-dim`).
*   **Hover State:**
    *   Entire card turns **blue** (`--link`).
    *   **All text** (Title and Description) switches to the background color (`--bg`) to ensure high contrast.

## 6. Interactions

We follow a restrained interaction model to maintain the minimalist aesthetic.

### Links & Clickable Items
*   **Default State:** All links (text and icons) are the current blue (`--link`). Text links are **underlined**.
*   **Visited State (Text Links):** Text links should become **less bright** (dimmed) after being visited. This does not apply to icons or interactive UI elements.
*   **Hover State (Links):** All links (text and icons) should become **a little brighter** on hover (e.g., via `filter: brightness(1.2)`).
*   **Hover State (UI Elements):** For interactive UI components (like menu items, cards, buttons, or **tags**), the element should become **blue** (`--link`) on hover.
*   **Selected/Active State:** When a page is selected or a button is toggled on, the element should remain **blue** (`--link`).
    *   *Note:* When the element is blue, the text color should typically switch to the background color (`--bg`) to ensure contrast.

## 7. Hugo Configuration

- **HTML Unsafe:** Enabled (`markup.goldmark.renderer.unsafe = true`) to allow
  raw HTML when necessary.
- **Syntax Highlighting:** `markup.highlight.noClasses = false` (we use CSS
  classes for styling).
- **Taxonomies:** Tags and Categories are supported.

## 7. File Structure

- `content/`: Markdown content files.
- `static/`: Assets (images, CSS, fonts) served as-is.
- `layouts/`: HTML templates.
  - `_default/baseof.html`: Master template.
  - `partials/`: Reusable components (head, footer, etc.).
