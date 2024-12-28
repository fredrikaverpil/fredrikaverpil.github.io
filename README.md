# fredrikaverpil.github.io

## 🚀 Quick start

### 📦 Install dependencies

```bash
# for social cards, see
# https://squidfunk.github.io/mkdocs-material/setup/setting-up-social-cards
brew install cairo freetype libffi libjpeg libpng zlib

# install uv
brew install uv

# install dependencies into virtual environment
uv sync
```

> [!Note]
>
> The
> [troubleshooting docs](https://squidfunk.github.io/mkdocs-material/plugins/requirements/image-processing/?h=brew#cairo-graphics)
> outlines solutions to issues.

### 💄 Update dependencies

```bash
uv sync --upgrade
```

### 🍽️ Serve locally

```bash
uv run mkdocs serve --dirtyreload
```

## ✨ Useful stuff

- https://squidfunk.github.io/mkdocs-material/reference/
- https://facelessuser.github.io/pymdown-extensions/
