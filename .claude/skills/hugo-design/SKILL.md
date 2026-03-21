---
name: hugo-design
description: >
  This skill should be used when editing CSS, modifying visual appearance,
  changing colors, typography, spacing, layout, responsive breakpoints, or
  any design-related work on the Hugo site. Also use when adding new UI
  components, adjusting dark/light mode, or working with the Zenbones color
  palette, Golden Ratio spacing system, or compounding contrast pattern.
---

# Hugo Design System

Use when modifying the visual design of the site. The design system enforces
consistency — follow these principles rather than ad-hoc values.

## Core Principles

- **Golden Ratio (phi = 1.618)** for all spacing and vertical rhythm
- **Compounding contrast** — layered elements progressively darker (light) or lighter (dark)
- **No rounded corners** anywhere
- **Links underline on hover** only
- **Content-first** — form follows function, minimal decoration
- **Zero-JS for styling** — all visual behavior in CSS, JS only for progressive enhancement

## Key Files

| File | Purpose |
|------|---------|
| `static/css/style.css` | Main stylesheet (~1650 lines), all design tokens |
| `static/css/github-link.css` | GitHub link-specific styling |
| `layouts/partials/header.html` | Nav, search, theme toggle, font preloading |
| `layouts/partials/footer.html` | Footer with social links |

## Design Tokens

All tokens in `:root` of `style.css`. Full reference: `references/design-tokens.md`.

- Never use magic numbers — use `--space-*` variables for spacing
- Never hardcode colors — use semantic `--*` variables that adapt to theme
- Test both light and dark mode when changing any color or background
- Use `font-display: block` (not `swap`) to prevent FOUT
- Berkeley Mono is primary but licensed/not in git — ensure Commit Mono fallback works

## Responsive Breakpoints

Four breakpoints in priority order:
1. `768px` — tablet (reduce spacing, stack layouts)
2. `600px` — medium phone (smaller headings)
3. `480px` — small phone (minimal padding)
4. `430px` height + landscape — phone horizontal (compact nav)

Responsive overrides redefine CSS custom properties inside media queries within
`:root`, keeping specificity flat.

## Dark Mode

Theme switching via `[data-theme="dark"]` on `<html>`. Default follows
`prefers-color-scheme`. Manual toggle persists to `localStorage`.

Pattern: `--light-*` and `--dark-*` tokens, aliased to `--*` active tokens.
Dark mode block reassigns the aliases.

## Adding a New Component

1. Define semantic color tokens for both light and dark themes
2. Use `--space-*` variables for all padding/margins
3. Apply compounding contrast (component bg slightly offset from parent bg)
4. No rounded corners
5. Test at all four breakpoints and both themes

## Preview

```bash
make serve  # Local dev server with LiveReload
```
