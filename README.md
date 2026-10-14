# fredrikaverpil.github.io

My personal blog and portfolio.

## Development

This project uses [Hugo](https://gohugo.io/) and [mise](https://mise.jdx.dev) to
install tools and run tasks. Run `mise tasks` to list them.

### Prerequisites

- **mise:** Installs Go (for `go tool hugo`), Pagefind and htmltest.

### Quickstart

| Command                          | Description                            |
| :------------------------------- | :------------------------------------- |
| `mise run serve`                 | Start local dev server (LiveReload).   |
| `mise run build`                 | Build production site to `public/`.    |
| `mise run clean`                 | Remove build artifacts.                |
| `mise run new --title "<title>"` | Create a new blog post.                |
| `mise run check-links`           | Build site and check for broken links. |
| `mise run all`                   | Run all checks.                        |

### Build Workflow

For local development:

```bash
mise run clean && mise run build && mise run serve
```

Run `mise run clean` before `mise run build` to ensure the Pagefind index builds
correctly and to remove stale artifacts.

## Documentation

See [.claude/CLAUDE.md](.claude/CLAUDE.md) for the complete design system, style
guide, and technical details.
