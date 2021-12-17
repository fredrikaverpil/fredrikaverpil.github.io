---
layout: post
title: "Developing with Apple Silicon"
tags: [osx]
---

In software development, certain software were not designed to run on the ARM-based Apple Silicon. Thankfully, there are workarounds. Like for the rest of this blog, this post aims to serve as a personal notebook and also for sharing this knowledge.

<!--more-->

## Apple Silicon vs Intel in the Terminal

MacOS ships with `Terminal.app`. This runs native on Apple Silicon but allows for a customization where it would run under Rosetta 2.

This requires installing Rosetta 2:

```bash
/usr/sbin/softwareupdate --install-rosetta
```

I have duplicated the Terminal application and renamed the duplicate into "Terminal Rosetta". Then I've ticked the "Open using Rosetta" checkbox after having hit <kbd>Cmd</kbd>+<kbd>i</kbd> on the icon. This gives me two Terminal applications. One to run native applications in and one for Intel emulation.

One can also execute commands from the native terminal which are to be emulated by Rosetta by executing `arch -x86_64 <command>`.

All commands in this guide has been executed in the default and native Terminal app, unless stated otherwise.

When reading this guide, it might be helpful to also be aware of my [dotfiles](https://github.com/fredrikaverpil/dotfiles) setup, and in particular the [`sourcing.sh` script](https://github.com/fredrikaverpil/dotfiles/blob/main/shell/sourcing.sh) which gets sourced by my shell.

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

I prefer managing Python versions via [pyenv](https://github.com/pyenv/pyenv). Because some package won't install on Apple Silicon, I need to be able to install the Intel-verison of Python for some projects.

Pyenv itself can be installed for Apple Silicon only, but we'll need `pyenv-alias` to accomodate for Intel versions of the Python interpreter:

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

You can now use pyenv as you usually do with `pyenv local` or `pyenv global`. There's only need for `pyenv86` when installing new versions. Examples:

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

### Node (nvm, npm)

I haven't found a great workflow here, unfortunately. Using nvm/npm relies on using either the default Terminal or the duplicated Rosetta 2 Terminal.

First, let's install nvm for the respective architectures via Homebrew:

```bash
# ARM
brew install nvm

# Intel
NVM_DIR=$HOME/.nvm_x86 brew86 install nvm
```

This will result in two directories in the home folder:

- `~/.nvm` - Apple Silicon
- `~/.nvm_x86` - Intel

As part of the nvm installation, you need to source the `nvm.sh` in your `.bashrc` or `.zshrc`. Here's a snippet from my dotfiles, which will load the right one, depending on which terminal app (native vs Rosetta) you use:

```bash
# Node version manager
if [ `uname -m | grep arm64` ] && [ -d /opt/homebrew/opt/nvm ]; then
    # brew-installed nvm, macOS arm64
    [ -s "/opt/homebrew/opt/nvm/nvm.sh" ] && . "/opt/homebrew/opt/nvm/nvm.sh"  # This loads nvm
    [ -s "/opt/homebrew/opt/nvm/etc/bash_completion.d/nvm" ] && . "/opt/homebrew/opt/nvm/etc/bash_completion.d/nvm"  # This loads nvm bash_completion
elif [ `uname -m | grep x86_64` ] && [ -d /usr/local/opt/nvm ]; then
    # brew-installed nvm, macOS x86_64
    [ -s "/usr/local/opt/nvm/nvm.sh" ] && . "/usr/local/opt/nvm/nvm.sh"  # This loads nvm
    [ -s "/usr/local/opt/nvm/etc/bash_completion.d/nvm" ] && . "/usr/local/opt/nvm/etc/bash_completion.d/nvm"  # This loads nvm bash_completion
```

Ideally, I would've liked to differentiate the two variants of nvm using executables `nvm` and `nvm86`. But I haven't figured out a way to achieve this. Therefore I am running `nvm` like this, depending on whether you want nvm/npm/node for either ARM or Intel:

- `nvm` in the default terminal
- `NVM_DIR=$HOME/.nvm_x86 nvm` in the Rosetta 2 terminal

Then I can install node and use npm from the respective terminal:

```bash
# ARM, executed from the native Terminal
nvm install 14
nvm use 14
npm install

# Intel, must be executed from the Rosetta 2 Terminal (!)
NVM_DIR=$HOME/.nvm_x86 nvm install 14
NVM_DIR=$HOME/.nvm_x86 nvm use 14
npm install
```

You can verify which architecture was used for a certain existing installation (using any of the terminals):

```bash
$ cd my_arm_project
$ npm root -g
~/.nvm/versions/node/v14.18.1/lib/node_modules

$ cd my_intel_project
$ npm root -g
~/.nvm_x86/versions/node/v14.18.1/lib/node_modules
```

## Detecting running under Apple Silicon (or ARM in general)

You can check in your shell or in e.g. Python which architecture is currently in use:

| Command                                                     | Apple Silicon | macOS Intel (Rosetta 2) | Linux ARM | Linux Intel |
| ----------------------------------------------------------- | ------------- | ----------------------- | --------- | ----------- |
| `uname -m`                                                  | arm64         | x86_64                  | aarch64   | x86_64      |
| `uname -p`                                                  | arm           | i386                    | ?         | x86_64      |
| `python3 -c "import platform; print(platform.processor())"` | arm           | i386                    | ?         | x86_64      |

## vscode

This is my set up in vscode, so I can quickly create terminal tabs for either a native or Rosetta 2 experience:

```json
{
  "terminal.integrated.profiles.osx": {
    "zsh": {
      "path": "zsh",
      "color": "terminal.ansiGreen"
    },
    "zsh-rosetta": {
      "path": "arch",
      "color": "terminal.ansiRed",
      "args": ["-x86_64", "zsh"]
    }
  },
  "terminal.integrated.defaultProfile.osx": "zsh"
}
```

## Working with containers

### Build from source

A pre-built binary or wheel for ARM may not exist when installing via e.g. `npm` or `pip`. You can try to install all the necessary build tools and instead try having the software built from source. This will result in staying native without any emulation. Example:

```Dockerfile
FROM node:14-buster-slim

# Install sqlite3 and build tools, so to build from source during
# "npm install", since no precompiled build for ARM is yet available
RUN if [ "$(uname -m)" = "aarch64" ]; then \
    apt-get update && apt-get install -y \
        apt-transport-https ca-certificates sqlite3 \
        build-essential python-dev python3-dev \
        ; \
    fi

RUN npm install sqlite3
```

### Define the platform

If the software in question simply was not written to be compiled on ARM at all, you can instead leverage Docker's `--platform` argument, which will build the container as if you were on a Linux system:

```bash
docker build --platform linux/amd64 ...
```

### Target multiple platforms

Using Docker's [buildx](https://docs.docker.com/buildx/working-with-buildx/), you can build images which target multiple platforms.
