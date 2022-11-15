---
title: 'Debugging Poetry with Visual Studio Code'
tags: [python]

cover:
  image: "/static/poetry_debug/debug.png"
  alt: "Visual Studio Code"
  relative: false # To use relative path for cover image, used in hugo Page-bundles

# PaperMod
ShowToc: true
TocOpen: true

created: 2021-04-17T02:00:00+02:00
updated: 2022-11-15T22:28:09+01:00
---

A guide on how to set up debugging of [Poetry](https://python-poetry.org/) in [Visual Studio Code](https://code.visualstudio.com/), using [Pipx](https://github.com/pipxproject/pipx) and [Pyenv](https://github.com/pyenv/pyenv).

## Prerequisites

My developer environment is Ubuntu 20.04 via WSL2 running on Windows 10, so that's what this guide is written for. In order to get set for debugging, we need to install/setup a couple of tools:

* [Pyenv](https://github.com/pyenv/pyenv) - Python version management
* [Pipx](https://github.com/pipxproject/pipx) - install and run Python applications in isolated environments
* [Poetry](https://github.com/python-poetry/poetry) - you should know what this is, if you are reading this ;)

### Pyenv

I like to pick the Python interpreter version for a system-wide installations of tools, but also for individual projects using [Pyenv](https://github.com/pyenv/pyenv). Installation instructions can be found [here](https://github.com/pyenv/pyenv-installer) and pyenv's prerequisites can be found in their [wiki](https://github.com/pyenv/pyenv/wiki).

Let's install Python 3.9.2 for system-wide installed tools and Python 3.8.8 for development/debugging of Poetry. The versions selected are just used to illustrate that different interpreters can be used for Poetry and the projects themselves.

```bash
$ pyenv install 3.9.2
$ pyenv install 3.8.8
```

Try it out, once installed, with e.g. Python 3.9.2:

```bash
$ pyenv global 3.9.2  # set the 'python' command to use this new version

$ python --version
Python 3.9.2

$ pyenv global system  # reset back to system default
```

### Pipx

Once pyenv is installed and a Python interpreter of choice is available, I like to have an installation of poetry from the same branch or pull request I am about to develop/debug. This can easily be maintained using [Pipx](https://github.com/pipxproject/pipx).

Let's install pipx into the Python interpreter version of choice:

```bash
$ pyenv global 3.9.2
$ pip install --upgrade pip  # always good to be on the latest pip
$ pip install pipx
$ pyenv global system
```

Make sure to follow the pipx installation instructions and add `~/.local/bin` to `$PATH`, for example.

### Poetry

As a last prerequisite, I'll install a "generic" Poetry version via pipx. This is so that we can bootstrap Poetry's own development installation.

When debugging in Poetry's `master` branch, I'll install Poetry from the latest commit in `master`. I also like to install Poetry using a pipx-suffix. In this case I'll use the suffix `@master` and make the Poetry executable available as `poetry@master`.

```bash
$ pipx install --suffix=@master --force git+https://github.com/python-poetry/poetry.git'
```

Sometimes I might want to install Poetry from a GitHub pull request (in this example pull request [#3967](https://github.com/python-poetry/poetry/pull/3967)) and make this version of Poetry available via the `poetry@3967` command:

```bash
$ pipx install --suffix=@3967 --force 'poetry @ git+https://github.com/python-poetry/poetry.git@refs/pull/3967/head'
```

Note that the `--force` command makes it possible to run those commands again, to "update" to the current code in either master or in the pull request.

## Visual Studio Code debug setup

This consists of a few steps:

* Download the Poetry source code
* Set up the virtual environment
* Visual Studio Code setup

### Download the Poetry source code

In order to develop and debug poetry, we first need to clone down the git repo's source code:

```bash
$ git clone https://github.com/python-poetry/poetry.git
```

### Set up the virtual environment

In the project folder of `poetry`, we can create a `.python-version` file, read by pyenv and which will set the Python interpreter version invoked by the `python` command: 

```bash
$ cd poetry

$ pyenv local 3.8.8  # creates .python-version

$ python --version
Python 3.8.8
```

Always make sure you're up to date with pip:

```bash
pip install --upgrade pip
```

Let's now create a virtual environment and install the Poetry development environment. This can be done in several ways:

```bash
$ python -m venv .venv
$ source .venv/bin/activate
$ poetry@master install
```

```bash
$ pip install virtualenv
$ virtualenv .venv
$ source .venv/bin/activate
$ poetry@master install
```

However, the below will _not_ work, as Poetry (at least not currently) does _not_ support reading the `.python-version` file, created from the `pyenv local 3.8.8` command:

```bash
# WARNING: this will NOT work!
$ poetry config virtualenvs.in-project true
$ poetry@master install  # creates the ".venv" automatically
$ source .venv/bin/activate

$ python --version
Python 3.9.2  # here we expected Python 3.8.8!
```

So I would go either with `venv` or `virtualenv`. Moving on...

Even if `poetry@master` uses Python 3.9.2, it will still be able to complete an installation in the Python 3.8.8 virtual environment!

Now you have the `poetry` command at your disposal, as well as `python -m poetry`, provided by the development installation. The latter is what we are going to use when debugging!

### Visual Studio Code setup

Launch Visual Studio Code and open the `poetry` project folder, containing all source code. Make sure you have the [Python extension](https://marketplace.visualstudio.com/items?itemName=ms-python.python) and all other necessities for sane Python development. ;)

Also select your virtual environment (`.venv`) as the active Python interpreter for the project.

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

You can see in the above file that I have added two basic examples of executing Poetry; `poetry install` and `poetry update`. The way this works is that Visual Studio Code will wrap e.g. the `python -m poetry install -vvv` command in its debugger. 

Note that we no longer use e.g. `poetry@master`. This was only meant to bootstrap the development environment and make the `poetry` command available in the virtual environment.

#### Set breakpoints and run

Add breakpoints by clicking just left to the line number of the code you wish to debug.

Now, in the debug menu, from the "Run and debug" section (upper left corner), you can pick between the different configurations (from the `.vscode/launch.json`) and execute them (click the green "play" button). Visual Studio Code's debugger wrapper will now execute the command and stop the execution on your breakpoints.

![alt text](/static/poetry_debug/debug.png "Debug")

Inspect objects and navigate the call stack to the left and use the navigation in the top center to continue, step over/into/out of, restart or stop.
You can also view the terminal or use the debug console at the bottom. Keep track of your breakpoints in the lower left corner section "Breakpoints".

More on debugging in Visual Studio Code [here](https://code.visualstudio.com/docs/editor/debugging).

Happy developing/debugging!
