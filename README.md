# fredrikaverpil.github.io

My personal blog and portfolio.

## Development

This project uses [Hugo](https://gohugo.io/) and standard UNIX tools.

### Prerequisites

- **Hugo:** `brew install hugo` (or via `go install`).
- **Node/NPM:** Required for `npx` to run [Pagefind](https://pagefind.app/).

### Quickstart

| Command      | Description                          |
| :----------- | :----------------------------------- |
| `make serve` | Start local dev server (LiveReload). |
| `make build` | Build production site to `public/`.  |
| `make clean` | Remove build artifacts.              |

### Build Workflow

For local development:
```bash
make serve
```

For a clean production build:
```bash
make clean && make build
```

**Note:** Always run `make clean` before `make build` to ensure the Pagefind search index is regenerated correctly and to remove stale artifacts.

## Documentation

See [HUGO.md](HUGO.md) for the complete design system, style guide, and technical details.