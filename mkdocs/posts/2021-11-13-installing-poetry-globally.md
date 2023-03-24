---
date: 2021-11-13
tags:
- python
---

# Installing Poetry system-wide

I recently got some feedback (thank you [@simmel](https://github.com/simmel), much appreciated!) on a previous post on [debugging Poetry](2021-04-17-debugging-poetry.md). I then realized it was a bit hard to follow if all you wanted to do was to install Poetry globally, with some added control from the [default installation method](https://python-poetry.org/docs/#installation).

This post aims to focus on this and cover the different system-wide installation alternatives that I am aware of.

<!-- more -->

## System-wide install with pipx

You can very easily make Poetry available system-wide, by just following the [installation docs](https://python-poetry.org/docs/#installation). However, this makes it a bit harder to e.g. install in-development builds of Poetry. By leveraging [pipx](https://github.com/pypa/pipx), this can be solved:

!!! tip "Install Poetry with pipx"

    ```bash
    # macOS
    brew install pipx
    pipx install poetry

    # Linux (apt-get)
    apt install pipx
    pipx install poetry
    ```

Yeah, sorry Windows users. No easy setup here that I am aware of, unfortunately. The only reasonable package manager for Windows is [winget](https://docs.microsoft.com/en-us/windows/package-manager/winget/), in my opinion, and that has no pipx or poetry package.

## In-development Poetry builds

Using pipx, we can install Poetry from git repo. Pipx also has a nice "suffix" feature, which can be used to differentiate the different installations.

```bash
# Install from master branch as poetry@master
$ pipx install --suffix=@master --force git+https://github.com/python-poetry/poetry.git

# Install from develop branch as poetry@develop
$ pipx install --suffix=@develop --force git+https://github.com/python-poetry/poetry.git@develop

# Install from release/tag 1.2.0rc1
$ pipx install --suffix=@1.2.0rc1 --force git+https://github.com/python-poetry/poetry.git#1.2.0rc1

# Install from PR #3967 as poetry@3967
$ pipx install --suffix=@3967 --force git+https://github.com/python-poetry/poetry.git@refs/pull/3967/head

# Install from local folder as poetry@local
$ pipx install --suffix=@local --force ~/code/repos/poetry
```

Note that the `--force` option makes it possible to run those commands again, to “update” to the current code in either master or in the pull request.

Run `pipx list` to see all installations of Poetry:

```bash
$ pipx list                                                                                              
venvs are in ~/.local/pipx/venvs
apps are exposed on your $PATH at ~/.local/bin
   package poetry 1.1.11, Python 3.10.0
    - poetry
   package poetry 1.2.0a0 (poetry@3967), Python 3.10.0
    - poetry@3967
   package poetry 1.1.0a3 (poetry@develop), Python 3.10.0
    - poetry@develop
   package poetry 1.2.0a2 (poetry@local), Python 3.10.0
    - poetry@local
   package poetry 1.2.0a2 (poetry@master), Python 3.10.0
    - poetry@master
```

## Bonus: use pyenv to control the Python interpreter used by pipx and poetry

In the examples above, you don't really control the underlying Python interpreter version. If that's important to you, you can use e.g. [pyenv](https://github.com/pyenv/pyenv) to manage Python interpeter installations:

```bash
$ pyenv install 3.10.0  # install CPython 3.10.0 into ~/.pyenv/versions/3.10.0
$ pyenv global 3.10.0  # make 'python' and 'pip' use CPython 3.10.0
$ pip install pipx  # install pipx into the 3.10.0 installation
$ pipx install poetry  # install poetry in ~/.local/pipx/venvs/poetry and its binary in ~/.local/bin/poetry
$ pyenv global system  # stop making 'python' and 'pip' point use CPython 3.10.0 and revert it back to system-default

$ pipx --version
```

Now you can more easily install exactly the version of pipx you desire, and the version of Python you want to use.

You can read more about what each pyenv command does in the pyenv documentation. You'll also have to add some paths to your `$PATH` (as part of installing pyenv and pipx) for this to work.

## Thoughts and comments?

Please let me know how you deal with this, and if you have any suggestions, in the comments section below!