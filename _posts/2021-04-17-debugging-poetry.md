---
layout: post
title: 'Debugging poetry with Visual Studio Code'
tags: [python]
---

How does two of my favorite technologies go together in debug mode?

<!--more-->

## Prerequisites

My developer environment is Ubuntu 20.04 via WSL2 running on Windows 10, so that's what this guide is written for. Since it's all bash, git and Visual Studio Code, it should be cross-platform.

### Pyenv

I like to pick the Python interpreter version for a project using [pyenv](https://github.com/pyenv/pyenv). Installation instructions can be found [here](https://github.com/pyenv/pyenv-installer) and pyenv's prerequisites can be found in their [wiki](https://github.com/pyenv/pyenv/wiki).

```bash
pyenv install 3.9.4  # build and install Python 3.9.4
```

Try it out, once installed:

```bash
$ pyenv global 3.9.4  # set the 'python' command to use this new version

$ python --version
Python 3.9.4

$ pyenv global system  # reset back to system default
```

### Pipx

Once pyenv is installed and a Python interpreter of choice is available, I like to have an installation of poetry from the same branch or pull request I am about to develop/debug. This can easily be maintained using [pipx](https://github.com/pipxproject/pipx).

Let's install pipx into the Python interpreter version of choice:

```bash
pyenv global 3.9.4
pip install --upgrade pip  # always good to be on the latest pip
pip install pipx
pyenv global system
```

Make sure to follow the pipx installation instructions and add `~/.local/bin` to `$PATH`, for example.

### Poetry

As a last prerequisite, I'll install pipx-install Poetry, usually from the master branch and make it available via the `poetry@master` command:

```bash
pipx install --suffix=@master --force git+https://github.com/python-poetry/poetry.git'
```

Sometimes I might want to install Poetry from a GitHub pull request (in this example pull request #3967) and make this version of Poetry available via the `poetry@3967` command:

```bash
pipx install --suffix=@3967 --force 'poetry @ git+https://github.com/python-poetry/poetry.git@refs/pull/3967/head'
```

Note that the `--force` command makes it possible to run those commands again, to "update" to the current code in either master or in the pull request.

Extra: I like to keep my virtual environments in the git project folder, so I also configure Poetry with `poetry config virtualenvs.in-project: true`.

## Visual Code debug setup

### Download the Poetry source code

In order to develop and debug poetry, we first need to clone down the git repo's source code:

```bash
git clone https://github.com/python-poetry/poetry.git
```

Then we'll install the project. Either create and activate a virtual environment or configure Poetry to do it for you (as mentioned in the previous section). Then install Poetry with all its dependencies into the virtual environment:

```bash
cd poetry
poetry@master install
```

### Visual Studio Code setup

Launch Visual Studio Code and open the Poetry folder. Make sure you have the Python extension and all other necessities for sane Python development. ;)

Make sure you select your virtual environment as the current Python interpreter.

#### Set up tasks.json

Debugging in Visual Studio Code is set up in the project folder's `.vscode/launch.json` file, so let's create it:

```json
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "poetry install",
            "type": "python",
            "request": "launch",
            "module": "poetry",
            "args": [
                "-vvv",  // change or comment out for different verbosity level
                "install"
            ],
        },
        {
            "name": "poetry update",
            "type": "python",
            "request": "launch",
            "module": "poetry",
            "args": [
                "-vvv",  // change or comment out for different verbosity level
                "update"
            ],
        }
    ]
}
```

You can see in the above file that I have added two basic ways of executing Poetry; `poetry install` and `poetry update`. Set up a new configuration which executes the command you wish to debug.

Add breakpoints, just left to the line number of the code you wish to debug in the git repo.

Now, from the debug menu, you can pick between the different configurations above and execute them in the debugger. This will cause these commands to run inside the Visual Studio Code debug wrapper.

More on debugging in Visual Studio Code [here](https://code.visualstudio.com/docs/editor/debugging).

Happy developing/debugging!
