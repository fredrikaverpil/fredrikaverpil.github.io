# fredrikaverpil.github.io

My personal blog and portfolio.

## Development

This project uses [Hugo](https://gohugo.io/) and
[Pocket](https://github.com/fredrikaverpil/pocket), a Makefile-like task runner.
Run `./pok` to execute linting, formatting, and tests.

### Prerequisites

- **Go:** Required for Hugo (`go tool hugo`) and Pocket.

### Quickstart

| Command       | Description                          |
| :------------ | :----------------------------------- |
| `./pok serve` | Start local dev server (LiveReload). |
| `./pok build` | Build production site to `public/`.  |
| `./pok clean` | Remove build artifacts.              |
| `./pok`       | Run linting, formatting, and tests.  |

### Build Workflow

For local development:

```bash
./pok clean && ./pok build && ./pok serve
```

Run `./pok clean` before `./pok build` to ensure the Pagefind index builds
correctly and to remove stale artifacts.

## Documentation

See [.claude/CLAUDE.md](.claude/CLAUDE.md) for the complete design system, style
guide, and technical details.
