---
layout: post
title: "Developing with Apple Silicon"
tags: [macos]
---

In software development, certain software were not designed to run on the ARM-based Apple Silicon. Thankfully, there are pretty nice workarounds. This blog post aims to serve as a notebook from my own issues and the solutions to them.

<!--more-->

## Apple Silicon vs Intel in the Terminal

WIP notes:

- softwareupdate --install-rosetta
- Duplicate, rename to "Terminal Rosetta" and tick "Open using Rosetta". Or, from the native Apple Silicon Terminal app, run `arch -x86_64 <command>`
- All commands in this post are executed in the normal Apple Silicon Terminal.
- In this guide, you can look at my shell/sourcing.sh to see how I have set up the different tools...

## Installing two variants of certain software

### Homebrew

```bash
# Install Homebrew for Apple Silicon
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"

# Install Homebrew for Intel
arch -x86_64 /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
```

After completing the installation, the `brew` command will be available, which will install native Apple Silicon software. We can create a new `brew86` command which will install Rosetta 2-emulated software. You might be able create an alias:

```bash
alias brew86='arch -x86_64 /usr/local/bin/brew "$@"'
```

...but I prefer to create [a shell script, which I put in my dotfiles](https://github.com/fredrikaverpil/dotfiles/blob/main/shell/bin/brew86). Once that is on $PATH, you can use either `brew` or `brew86` to manage Homebrew-installed software.

Software installs into:

- `$("brew --cellar")` and `$("brew --prefix")/bin` for Apple Silicon
- `$("brew86 --cellar")` and `$("brew86 --prefix")/bin` for Intel

Then I install all software using `brew` for Apple Silicon. But in case of issues I can fall back to the Intel version using `brew86`.

### Python (pyenv)

I prefer managing Python versions via pyenv. Because some package won't install on Apple Silicon, I need to be able to install the Intel-verison of Python for some projects.

Pyenv itself can be installed for Apple Silicon only, but we'll need `pyenv-alias` to accomodate for Intel versions:

```bash
curl -s -S -L https://raw.githubusercontent.com/pyenv/pyenv-installer/master/bin/pyenv-installer | bash
git clone https://github.com/s1341/pyenv-alias.git ~/.pyenv/plugins/pyenv-alias
```

After having completed the pyenv installation, we can go ahead and install Python for Apple Silicon first:

```bash
brew install openssl readline sqlite3 xz zlib # required to build python
pyenv install 3.10.1
```

Then we create another alias or [shell script](https://github.com/fredrikaverpil/dotfiles/blob/main/shell/bin/pyenv86) to make `pyenv86` available:

```bash
alias pyenv86='CFLAGS="-I$(brew86 --prefix openssl)/include" LDFLAGS="-L$(brew86 --prefix openssl)/lib" arch -x86_64 pyenv "$@"'
```

Then let's install the same Python version, but for Intel. To prevent a clash with the already intalled `3.10.1`, we'll install it using `pyenv-alias` as `3.10.1_x86`:

```bash
brew86 install openssl readline sqlite3 xz zlib # required to build python
VERSION_ALIAS="3.10.1_x86" \
    pyenv86 install 3.10.1
```

You can use this crude check to verify that each respective version works as intended:

```bash
$ ~/.pyenv/versions/3.10.1/bin/python -c "import platform; print(platform.processor())"
arm

$ ~/.pyenv/versions/3.10.1_x86/bin/python -c "import platform; print(platform.processor())"
i386
```

You can now use pyenv as you usually do with `pyenv local` or `pyenv global`, e.g:

```bash
cd my_arm_project
pyenv local 3.10.0
python -m venv .venv
source .venv/bin/activate

# or for Intel...

cd my_intel_project
pyenv local 3.10.0_x86
python -m venv .venv
source .venv/bin/activate
```

I usually create "x86" virtual environments, to be used in vscode, for some projects which needs to install software that has no ARM-release yet. This has worked out great so far. I can't really see any performance penalty either.

### Node

#### NVM

### NPM

npm root -g

## Working with containers

docker build --platform linux/amd64 ...

### Detecting running under aarch64

### Defining the platform
