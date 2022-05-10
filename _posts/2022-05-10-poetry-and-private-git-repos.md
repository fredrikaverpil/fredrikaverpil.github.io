---
layout: post
title: "Distributing Python packages in a private GitHub organization"
tags: [python]
---

There's a challenge in distributing Python packages in a private GitHub organization with private repositories. This post will try to clear this up a bit.

<!--more-->

## Background

If you develop your Python package project with Poetry in a public GitHub repository, it's easy to add it as a dependency to a Poetry project with `poetry add`.

The `poetry.lock` file will then contain the commit sha, and to update the package you'll have to consider whether you want to pull down latest commit from the specified branch using `poetry update` or e.g. manually update the desired git tag in `pyproject.toml` followed by a `poetry update`. See the [docs on git dependencies](https://python-poetry.org/docs/dependency-specification/#git-dependencies) for more details.

The nice part is this just works. But since you're developing in the open, you might as well also publish your package to pypi.org so that distributing and consuming your Python package will become a lot simpler.

However, if you are developing Python packages in a private GitHub organization, using private repositories and using GitHub Actions - this whole distribution story becomes a lot more challenging if you cannot host your own internal [pypiserver](https://github.com/pypiserver/pypiserver) and since PyPi does not (yet) offer private package hosting.

## Poetry 1.1

As of writing this, Poetry 1.1.13 is out, so that is what I assume here.

### Using git+https

#### Locally

If you [create a personal access token (PAT)](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token) with read permissions and authenticate it with your private organization, you can add a private Python package git repo using "git+https":

```bash
poetry add git+https://github.com/privateorg/privaterepo.git#main
```

You should see something like this in the `pyproject.toml`:

```toml
[tool.poetry.dependencies]
privaterepo = {git = "https://github.com/privateorg/privaterepo.git", rev = "main"}
```

#### GitHub Actions

To make this work with GitHub Actions (GHA) in the private organization/repositories, you have to [set up an organization secret](https://docs.github.com/en/actions/security-guides/encrypted-secrets#creating-encrypted-secrets-for-an-organization) which gives GHA read access to the Python project git repo. I'll call my secret `POETRY_GIT_TOKEN`.

Then you can install the project e.g. like this:

```yaml
name: check

on:
  workflow_dispatch:
  pull_request:
  push:
    branches:
      - main

env:
  PYTHON_VERSION: "3.10"
  PIPX_VERSION: "1.0"
  POETRY_VERSION: "1.1.13"

jobs:
  format:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - uses: actions/setup-python@v3
        with:
          python-version: ${{ env.PYTHON_VERSION }}

      - uses: fredrikaverpil/setup-pipx@v1.5
        with:
          pipx-version: ${{ env.PIPX_VERSION }}

      - run: pipx install poetry==${{ env.POETRY_VERSION }}

      - uses: actions/cache@v3
        id: cache
        with:
            path: |
            ~/.cache/pip
            ~/.cache/pypoetry/virtualenvs
            .venv_pipx
            key: ${{ runner.os }}-${{ runner.arch }}-py-${{ env.PYTHON_VERSION }}-{{ env.pythonLocation }}-pipx-${{ env.PIPX_VERSION }}-poetry-${{ env.POETRY_VERSION }}-${{ hashFiles('poetry.lock') }}

      - uses: de-vri-es/setup-git-credentials@v2
        with:
          credentials: https://${GITHUB_ACTOR}:${{secrets.POETRY_GIT_TOKEN}}@github.com/

      - run: poetry install
      - run: poetry run black --check --diff src
```

### Using git+ssh

#### Locally

This is how you locally add a dependency with "git+ssh":

```
poetry add git+ssh://git@github.com/privateorg/privaterepo.git#main
```

Note that no PAT is needed!

You should now see something like this in `pyproject.toml`:

```toml
[tool.poetry.dependencies]
privaterepo = {git = "ssh://git@github.com/privateorg/privaterepo.git", rev = "main"}
```

#### Github Actions

The same setup as for "git+https" will work here.

## Poetry 1.2

### Using git+https

#### Locally

You'll have to install poetry from master or, like I did, specifically from PR #5581:

```bash
$ pipx install --suffix=@5581 git+https://github.com/python-poetry/poetry.git@refs/pull/5581/head

$ poetry@5581 --version
Poetry (version 1.2.0b2.dev0)
```

Git credentials, as explained in the Poetry 1.1 part of this post, should still work on 1.2. But it will be a fallback since Poetry 1.2 will first try using [dulwich](https://github.com/jelmer/dulwich) - which is a pure-python git implementation.

Thanks to using dulwich and building on top of existing mechanisms to register the git repo along with git+https credentials, the following is possible with Poetry 1.2:

```bash
poetry@5581 config repositories.my-git-repo https://github.com/org/project.git
poetry@5581 config http-basic.my-git-repo username token
poetry@5581 add git+https://github.com/org/project.git
```

Locally, you would replace `username` with your GitHub username and `token` with your PAT.

#### GitHub Actions

To make this fly in GHA, you need that organization token described for Poetry 1.1. Then you can use environment variables corresponding to the configuration you added locally.


```yaml
name: check

on:
  workflow_dispatch:
  pull_request:
  push:
    branches:
      - main

env:
  PYTHON_VERSION: "3.10"
  PIPX_VERSION: "1.0"
  POETRY_VERSION: "PR-5881"

jobs:
  format:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - uses: actions/setup-python@v3
        with:
          python-version: ${{ env.PYTHON_VERSION }}

      - uses: fredrikaverpil/setup-pipx@v1.5
        with:
          pipx-version: ${{ env.PIPX_VERSION }}

      - run: pipx install git+https://github.com/python-poetry/poetry.git@refs/pull/5581/head

      - uses: actions/cache@v3
        id: cache
        with:
            path: |
            ~/.cache/pip
            ~/.cache/pypoetry/virtualenvs
            .venv_pipx
            key: ${{ runner.os }}-${{ runner.arch }}-py-${{ env.PYTHON_VERSION }}-{{ env.pythonLocation }}-pipx-${{ env.PIPX_VERSION }}-poetry-${{ env.POETRY_VERSION }}-${{ hashFiles('poetry.lock') }}

      - name: poetry install
        env:
          POETRY_REPOSITORIES_MYGITREPO_URL: https://github.com/org/project.git
          POETRY_HTTP_BASIC_MYGITREPO_USERNAME: anythingshouldbeokayhere
          POETRY_HTTP_BASIC_MYGITREPO_PASSWORD: ${{ secrets.POETRY_GIT_TOKEN }}
        run: |
          poetry install

      - run: poetry run black --check --diff src
```


### Using git+ssh

TBD

## Summary

### Poetry 1.1 and git+ssh

Pros:
  - Poetry 1.1.13 is out now, this works now.

Cons:
  - A bit finnicky maintenance of commit sha or git tag.
  - Requires a PAT.
  - Feels weird there isn't an easier way to tell GHA to use a secret token for "git+https" access.
