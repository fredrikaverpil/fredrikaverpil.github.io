# fredrikaverpil.github.io

My personal blog and portfolio.

## Development

This project uses [Hugo](https://gohugo.io/) and [mise](https://mise.jdx.dev) to
install tools and run tasks. Run `mise tasks` to list them.

### Prerequisites

- **mise:** Installs Go (for `go tool hugo`), Pagefind and htmltest.

### Quickstart

| Command                               | Description                            |
| :------------------------------------ | :------------------------------------- |
| `mise run hugo:serve`                 | Start local dev server (LiveReload).   |
| `mise run hugo:build`                 | Build production site to `public/`.    |
| `mise run hugo:clean`                 | Remove build artifacts.                |
| `mise run blog:new --title "<title>"` | Create a new blog post.                |
| `mise run hugo:check-links`           | Build site and check for broken links. |
| `mise run`                            | Run all checks.                        |

### Build Workflow

For local development:

```bash
mise run hugo:clean && mise run hugo:build && mise run hugo:serve
```

Run `mise run hugo:clean` before `mise run hugo:build` to ensure the Pagefind
index builds correctly and to remove stale artifacts.

## Documentation

See [.claude/CLAUDE.md](.claude/CLAUDE.md) for the complete design system, style
guide, and technical details.
