# fredrikaverpil.github.io

My personal blog and portfolio.

## Development

This project uses [Hugo](https://gohugo.io/) and standard UNIX tools.

### Prerequisites

- **Hugo:** `hugo` (via `go tool hugo`) to run build or serve commands.
- **Bun:** `bunx` to run [Pagefind](https://pagefind.app/).

### Quickstart

| Command      | Description                          |
| :----------- | :----------------------------------- |
| `make serve` | Start local dev server (LiveReload). |
| `make build` | Build production site to `public/`.  |
| `make clean` | Remove build artifacts.              |
| `make`       | Run all of the above.                |

### Build Workflow

For local development:

```bash
make
```

ean`before`make build` to ensure the Pagefind orrectly and to remove stale
artifacts.

## Documentation

See [HUGO.md](HUGO.md) for the complete design system, style guide, and
technical details.
