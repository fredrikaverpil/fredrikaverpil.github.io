---
layout: post
title: 'Installing Poetry system-wide'
tags: [python]
---

When I recently got some feedback (thank you [@simmel](https://github.com/simmel), much appreciated!) on a previous post on [debugging Poetry in vscode]({{ site.baseurl }}/2021/04/17/debugging-poetry/), I realized that post has very convoluted instructions on how to install Poetry system-wide.

This post aims extract that part out of that post (into this one) and clear up the different alternatives that I am aware of.

<!--more-->

## Easy system-wide install

You can very easily make Poetry available system-wide, without much hassle, leveraging [pipx](https://github.com/pypa/pipx):

### macOS

```bash
$ brew install pipx
$ pipx install poetry

$ poetry --version
```

### Linux (apt)

```bash
$ apt install pipx
$ pipx install poetry

$ poetry --version
```

### Windows

Yeah, sorry Windows users. No easy setup here that I am aware of, unfortunately. The only reasonable package manager for Windows is [winget](https://docs.microsoft.com/en-us/windows/package-manager/winget/), in my opinion, and that has no pipx or poetry package.

## Custom system-wide install

In the easy system-wide installation examples above, you don't really control the version of pipx or the underlying Python interpreter version, if that's important to you. To solve that, we can use [pyenv](https://github.com/pyenv/pyenv) to manage Python interpeter installations:

```bash
$ pyenv install 3.10.0  # install CPython 3.10.0 into ~/.pyenv/versions/3.10.0
$ pyenv global 3.10.0  # make 'python' and 'pip' use CPython 3.10.0
$ pip install pipx  # install pipx into the 3.10.0 installation
$ pipx install poetry  # install poetry in ~/.local/pipx/venvs/poetry and its binary in ~/.local/bin/poetry
$ pyenv global system  # stop making 'python' and 'pip' point use CPython 3.10.0 and revert it back to system-default

$ pipx --version
```

You can read more about what each pyenv command does in the pyenv documentation. You'll also have to add some paths to your `$PATH` (as part of installing pyenv and pipx) for this to work.

## Thoughts and comments?

Please let me know how you deal with this, and if you have any suggestions, in the comments section below!