# Design Tokens Reference

All tokens defined in `:root` of `static/css/style.css`.

## Spacing Scale (Golden Ratio)

Base unit: `1rem` (16px). Each level = previous x phi (1.618).

| Token | Calc | Approx |
|-------|------|--------|
| `--space-xs` | `0.382rem` | ~6px |
| `--space-sm` | `0.618rem` | ~10px |
| `--space-md` | `1rem` | ~16px |
| `--space-lg` | `1.618rem` | ~26px |
| `--space-xl` | `2.618rem` | ~42px |
| `--space-xxl` | `4.236rem` | ~68px |

## Derived Spacing

| Token | Value | Usage |
|-------|-------|-------|
| `--content-spacing` | `--space-lg` | Nav to content gap |
| `--section-gap` | `--space-xl` | Between major sections |
| `--heading-margin-top` | `--space-xl` | Above headings |
| `--heading-margin-bottom` | `--space-lg` | Below headings |

## Typography

| Token | Value |
|-------|-------|
| `--font-primary` | Berkeley Mono, monospace |
| `--font-code` | Berkeley Mono, monospace |
| `--font-italic` | Maple Mono, monospace |

Font stack: Berkeley Mono (primary) -> Commit Mono (fallback) -> Maple Mono (italics only).
Berkeley Mono is licensed, not in git. Downloaded in CI from a secret URL.

### Heading Sizes

`--h1-size: 1.8em` / `--h2-size: 1.5em` / `--h3-size: 1.3em` / `--h4-size: 1.2em` / `--h5-size: 1.1em` / `--h6-size: 1.05em`

Font weight: `--light-font-weight: 350` (light mode), `--dark-font-weight: 200` (dark mode).

## Color Palette (Zenbones)

Source: [zenbones.nvim](https://github.com/zenbones-theme/zenbones.nvim)

### Light Theme

| Token | Hex | Role |
|-------|-----|------|
| `--light-bg` | `#F0EDEC` | Background (warm sand) |
| `--light-fg` | `#2C363C` | Foreground (stone) |
| `--light-fg-dim` | `#6F777B` | Dimmed text |
| `--light-fg-muted` | `#859289` | Muted text |
| `--light-wood` | `#A86638` | Warnings |
| `--light-rose` | `#A8334C` | Important/errors |
| `--light-leaf` | `#4F894C` | Tips/success |
| `--light-water` | `#286486` | Links, notes |
| `--light-blossom` | `#88507D` | Examples |
| `--light-sky` | `#2679A8` | Accent |

### Dark Theme

| Token | Hex | Role |
|-------|-----|------|
| `--dark-bg` | `#1C1917` | Background (deep warm black) |
| `--dark-fg` | `#B4BDC3` | Foreground (stone) |
| `--dark-fg-dim` | `#8E969B` | Dimmed text |
| `--dark-fg-muted` | `#6C7377` | Muted text |
| `--dark-wood` | `#B77E64` | Warnings |
| `--dark-rose` | `#DE6E7C` | Important/errors |
| `--dark-leaf` | `#819B69` | Tips/success |
| `--dark-water` | `#6099C0` | Links, notes |
| `--dark-blossom` | `#C791C3` | Examples |
| `--dark-sky` | `#709CB8` | Accent |

### UI Elements

Active theme aliases (reassigned in dark mode):

| Semantic Token | Light Source | Dark Source |
|----------------|-------------|------------|
| `--bg` | `--light-bg` | `--dark-bg` |
| `--fg` | `--light-fg` | `--dark-fg` |
| `--link` | `--light-water` | `--dark-water` |
| `--meta-bg` | `#EAE6E4` | `#252220` |
| `--button-bg` | `#DCD8D6` | `#2F2C2A` |
| `--border` | `#C0BCBA` | `#353230` |

### Callout Colors

Each callout type has `--*-border` and `--*-bg` tokens:
- **note/info**: water-based
- **warning**: wood-based
- **tip**: leaf-based
- **important/caution**: rose-based
- **example**: blossom-based

## Component Tokens

| Token | Value |
|-------|-------|
| `--button-padding-vertical` | `5px` |
| `--button-padding-horizontal` | `10px` |
| `--video-aspect-ratio` | `56.25%` (16:9) |
| `--border-width-thin` | `4px` |
| `--border-width-thick` | `5px` |

## Syntax Highlighting

Monochrome/modest approach — all syntax tokens map to `--fg`, `--fg-dim`, or
`--fg-muted`. No colorful syntax highlighting.
