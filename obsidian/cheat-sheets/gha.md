---
title: 🎸 GitHub Actions
tags: [github-actions]
draft: false
summary: "Notes, snippets etc."

# PaperMod
ShowToc: false
TocOpen: true

date: 2022-12-17T14:09:32+01:00
---

## Poetry via Pipx

When using [actions/setup-python](https://github.com/actions/setup-python) and Poetry, you can use pipx to install Poetry. Make sure to then tell pipx which Python version to use:

```yaml
- uses: actions/setup-python@v4
  id: cpython_setup
  with:
    python-version: "3.10"

- run: pipx install poetry --python '${{ steps.cpython_setup.outputs.python-path }}'
```

When using a cache, you can use the `${{ steps.cpython_setup.outputs.python-version }}`:

```yaml
- uses: actions/cache@v3
  id: cache-poetry
  env:
    SEGMENT_DOWNLOAD_TIMEOUT_MIN: "15"
  with:
    path: |
      ~/.cache/pypoetry
      ~/.cache/pypoetry
      ~/.cache/pypoetry
    key: poetry-${{ runner.os }}-${{ runner.arch }}-py-${{ steps.cpython_setup.outputs.python-version }}-${{ hashFiles('poetry.lock') }}
```

The cache key can be extended with e.g. dependency groups.