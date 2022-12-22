---
title: 🎸 GitHub Actions
tags: [github-actions, python]
draft: false
summary: "Notes, snippets etc."

# PaperMod
ShowToc: false
TocOpen: true

date: 2022-12-17T23:28:21+01:00
---

## Single quotes

Never use double quotes inside `${{ ... }` as it is simply not supported.

## Tertiary

GHA has a really funky tertiary: `${{ x && 'ifTrue' || 'ifFalse' }}`

This only works if `<ifTrue>` isn't the empty string. If `<ifTrue> == ''` then `''` is considered as `false`, which then evaluates the right hand side of the `||`.

There are lots of gotchas here, and this is a great thread highlighting more of it: [Conditional operator or function for expression syntax · Issue #409 · actions/runner (github.com)](https://github.com/actions/runner/issues/409)

Example:

```yaml
steps:
- name: stuff
  env:
  PR_NUMBER_OR_MASTER: ${{ github.event.number == 0 && 'master' ||  format('pr-{0}', github.event.number)  }}
```

## Pipx via actions/setup-python

When using [actions/setup-python](https://github.com/actions/setup-python) and [Poetry](https://github.com/python-poetry/poetry), you can use [pipx)](https://github.com/pypa/pipx) to install Poetry outside of your project's virtual environment. The setup-python action provides pipx by default, but you might notice how it is not running under the Python version you chose.

To fix this, you can pass the `--python` argument to pipx:

```yaml
steps:
- uses: actions/setup-python@v4
  id: cpython_setup
  with:
    python-version: "3.10"

- run: pipx install <package> --python '${{ steps.cpython_setup.outputs.python-path }}'
```

## Python package caching

When using actions/cache, you can use the `${{ steps.cpython_setup.outputs.python-version }}` as part of the cache key:

```yaml
steps:
- uses: actions/setup-python@v4
  id: cpython_setup
  with:
    python-version: "3.10"

- uses: actions/cache@v3
  id: python-cache
  env:
    SEGMENT_DOWNLOAD_TIMEOUT_MIN: "15"
  with:
    path: |
      ~/.cache/pip
      ~/.cache/pypoetry
    key: pip-poetry-${{ runner.os }}-${{ runner.arch }}-py-${{ steps.cpython_setup.outputs.python-version }}-${{ hashFiles('poetry.lock') }}
```

The cache key can be extended with e.g. Poetry's dependency groups.

## Python shell

Python code can be written inline of a step, with the `shell: python` directive.

```yaml
steps:
  - name: Display the path
    run: |
      import os
      print(os.environ['PATH'])
    shell: python
```


See the [`jobs.<job_id>.steps[*].shell`  docs](https://docs.github.com/en/actions/using-workflows/workflow-syntax-for-github-actions#jobsjob_idstepsshell) for additional shells that are supported.

## Test coverage comment in PR

Here's a a super nice way to add a PR commit about the changed files' test coverage, powered by [romeovs/lcov-reporter-action](https://github.com/romeovs/lcov-reporter-action).

```yaml
steps:
  - name: run tests wrapped by coverage
    run: coverage run -m pytest

  - name: export coverage to lcov format
    run: coverage lcov

  - uses: romeovs/lcov-reporter-action@2a28ec3e25fb7eae9cb537e9141603486f810d1a
    # The reason for using a hash rather than a version/tag, is the project
    # failed in publishing this: https://github.com/romeovs/lcov-reporter-action/issues/47
    with:
      lcov-file: ./coverage/coverage.lcov
      filter-changed-files: true
      delete-old-comments: true
```

Configure pytest and coverage in `pyproject.toml`:

```toml
[tool.coverage.run]
omit = ["tests/migrations/*"]
source = ["src/mypkg"]

[tool.coverage.report]
exclude_lines =[
  "pragma: nocover",
  "if TYPE_CHECKING",
  "@overload",
]
skip_covered = true

[tool.coverage.html]
directory = "coverage/html"

[tool.coverage.lcov]
output = "coverage/coverage.lcov"

[tool.pytest.ini_options]
testpaths = "tests"
addopts = "-rxXs --color=yes"

```