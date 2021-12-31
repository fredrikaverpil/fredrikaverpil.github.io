---
layout: post
title: "Developing with Apple Silicon"
tags: [osx]
---

In software development, certain software were not designed to run on the ARM-based Apple Silicon. Thankfully, there are workarounds to install and run the Intel version of these applications. Like for the rest of this blog, this post aims to serve as a personal notebook and also for sharing this knowledge.

<!--more-->

- [Foreword](#foreword)
- [Rosetta 2](#rosetta2)
- [Two terminals; one native and one for Intel emulation](#terminals)
- [Installing two variants of certain software](#variants)
  - [Homebrew](#homebrew)
  - [Python (pyenv)](#pyenv)
  - [Node (nvm, npm)](#nvm)
- [Detecting running under Apple Silicon (or ARM in general)](#detection)
- [Visual Studio Code terminals](#vscode)
- [Docker containers](#containers)
  - [Build from source](#containers-source)
  - [Define the platform](#containers-platform)
  - [Target multiple platforms with buildx](#containers-buildx)

## <a name="foreword"></a> Foreword

When I started looking into solutions to my M1 challenges, a huge chunk was solved when I had read [Ekaterina Nikonova](https://twitter.com/EVNikonova)'s excellent blog post on [Python virtual environments with pyenv on Apple Silicon](http://sixty-north.com/blog/pyenv-apple-silicon.html).

This blog post is heavily influenced by her approach of setting up the `app` and `app86` versions for apps which might need both native and emulated treatment.

## <a name="rosetta2"></a> Rosetta 2

To be able to emulate Intel on M1 macs, first install Rosetta 2:

```bash
/usr/sbin/softwareupdate --install-rosetta
```

This enables the ability to run terminal commands using the Rosetta 2 Intel emulation:

```bash
arch -x86_64 <command>
```

You can read more in Apple's [official docs](https://support.apple.com/en-us/HT211861).

## <a name="terminals"></a> Two terminals; one native and one for Intel emulation

I have duplicated my Terminal application of choice (<kbd>Cmd</kbd>+<kbd>d</kbd>) and renamed the duplicate "Terminal Rosetta". Then I've ticked the "Open using Rosetta" checkbox after having hit <kbd>Cmd</kbd>+<kbd>i</kbd> on its icon. This gives me one Terminal to run for native applications and one for Intel emulation.

![]({{ site.url }}/blog/assets/applesilicon/terminals.png)

This is not how I usually work, as I prefer to set up commands which can run in the native terminal and perform emulation. But I figured, I'd mention this anyways.

Note: all commands in this guide has been executed in the default and native Terminal app, unless stated otherwise.

## <a name="variants"></a> Installing two variants of certain software

This describes how to install the native `app` and the Intel `app86` counterpart of certain software. As an example, I have `brew` and `brew86` set up, just like described in [Ekaterina](https://twitter.com/EVNikonova)'s [post](http://sixty-north.com/blog/pyenv-apple-silicon.html).

This can be achieved by creating shell aliases but also small shell "shim" scripts for the Intel variants. I prefer the shell script shim approach, and I have mine publicly available in the [shell/bin](https://github.com/fredrikaverpil/dotfiles/blob/main/shell/bin/) location of my dotfiles.

If you go for the shell script shim approach, make sure you make your shims executable and available on `$PATH`.

Feel free to copy and/or contribute with your improvements!

### <a name="homebrew"></a> Homebrew

Begin by installing native and Intel version of Homebrew:

```bash
# Install Homebrew for Apple Silicon
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"

# Install Homebrew for Intel
arch -x86_64 /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
```

After completing the installation, the `brew` command will be available which will install native Apple Silicon software. We can create a new `brew86` command which will install Rosetta 2-emulated Intel-compiled software. You might be able create an alias:

```bash
alias brew86='arch -x86_64 /usr/local/bin/brew "$@"'
```

But, like I mentioned previously, I prefer the shim approach. See my `brew86` shim [here](https://github.com/fredrikaverpil/dotfiles/blob/main/shell/bin/brew86).

Then I install all software using `brew` for Apple Silicon. But in case of issues I can fall back to the Intel version using `brew86`.

### <a name="pyenv"></a> Python (pyenv)

I prefer managing Python versions via [pyenv](https://github.com/pyenv/pyenv). Because some packages won't install on Apple Silicon, I need to be able to install the Intel-verison of Python for some projects.

Pyenv itself can be installed for Apple Silicon only, but we'll need `pyenv-alias` to accomodate for Intel versions of Python interpreter installations:

```bash
curl -s -S -L https://raw.githubusercontent.com/pyenv/pyenv-installer/master/bin/pyenv-installer | bash
git clone https://github.com/s1341/pyenv-alias.git ~/.pyenv/plugins/pyenv-alias
```

After having completed the pyenv installation, we can go ahead and install Python for Apple Silicon first:

```bash
brew install openssl readline sqlite3 xz zlib # required to build python
pyenv install 3.10.1
```

Then we create another alias (or [shim](https://github.com/fredrikaverpil/dotfiles/blob/main/shell/bin/pyenv86)) to make `pyenv86` available:

```bash
alias pyenv86='CFLAGS="-I$(brew86 --prefix openssl)/include" LDFLAGS="-L$(brew86 --prefix openssl)/lib" arch -x86_64 pyenv "$@"'
```

Then let's install the same Python version, but for Intel. To prevent a clash with the already intalled `3.10.1`, we'll install it using `pyenv-alias` as `3.10.1_x86`:

```bash
brew86 install openssl readline sqlite3 xz zlib # required to build python
VERSION_ALIAS="3.10.1_x86" pyenv86 install 3.10.1
```

You can use this crude check to verify that each respective version was installed using the intended architecture:

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

I usually create "x86" virtual environments, to be used in vscode, for some projects which needs to install software that has no ARM-release yet. This has worked out great so far.

### <a name="nvm"></a> Node (nvm, npm)

First, let's install nvm for the respective architectures via Homebrew:

```bash
# ARM
brew install nvm

# Intel
NVM_DIR=$HOME/.nvm_x86 brew86 install nvm
```

But instead of following the recommended installation instructions for bash completion and sourcing of `nvm.sh`, we'll take a different approach...

The bash completion can be added to e.g. your `.zshrc`, `.bashrc` or similar, and since this will work the same for both native and Intel, we can pick the native installation for this:

```bash
if [ -s "$(brew --prefix)/opt/nvm/etc/bash_completion.d/nvm" ]; then
  . "$(brew --prefix)/opt/nvm/etc/bash_completion.d/nvm"
fi
```

But for the part where you normally add sourcing of the `nvm.sh` script to e.g. `.zshrc` or `.bashrc`, shims will be setup instead.

First, we set up two main shims:

- `nvm_shim`
- `nvm86_shim`

Then we set up the rest of the command shims, which will call the main ones:

- `nvm`
- `node`
- `npm`
- `npx`
- `nvm86`
- `node86`
- `npm86`
- `npx86`

You can find the contents for these shims in the [shell/bin](https://github.com/fredrikaverpil/dotfiles/blob/main/shell/bin/) folder of my dotfiles.

Once those are all on `$PATH`, you can use the command shims for native installations Intel installations. Example:

```bash
# Native
nvm install 14
nvm use 14
npm init
npm install express

# Intel
nvm86 install 14
nvm86 use 14
npm86 init
npm86 install express
```

## <a name="detection"></a> Detecting running under Apple Silicon (or ARM in general)

You can check in your shell or in e.g. Python which architecture is currently in use:

| Command                                                     | macOS Apple Silicon | macOS Intel (Rosetta 2) | Linux ARM | Linux Intel |
| ----------------------------------------------------------- | ------------------- | ----------------------- | --------- | ----------- |
| `uname -m`                                                  | arm64               | x86_64                  | aarch64   | x86_64      |
| `uname -p`                                                  | arm                 | i386                    | aarch64   | x86_64      |
| `arch`                                                      | arm64               | i386                    | N/A       | N/A         |
| `python3 -c "import platform; print(platform.processor())"` | arm                 | i386                    | aarch64   | x86_64      |
| `node -p process.arch`                                      | arm64               | x64                     | ?         | ?           |

## <a name="vscode"></a> Visual Studio Code terminals

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

## <a name="containers"></a> Docker containers

I've come across a few cases where I e.g. haven't been able to build a Docker image which was not compiled for ARM. Thankfully, there are workarounds for this too.

### <a name="containers-source"></a> Build from source

A pre-built binary or wheel for ARM may not exist when installing via e.g. `npm` or `pip`. You can try to install all the necessary build tools and instead try having the software built from source. This will result in staying native without any emulation.

As an example, this will fail as of writing this blog post on my M1 mac:

```Dockerfile
FROM node:14-buster-slim

RUN npm install sqlite3
```

...and this will work fine, as [sqlite3](https://www.npmjs.com/package/sqlite3) will build from source:

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

### <a name="containers-platform"></a> Define the platform

If the software in question simply was not written to be compiled on ARM at all, you can instead leverage Docker's `--platform` argument, which will build the container as if you were on a Linux system:

```bash
docker build --platform linux/amd64 ...
```

Please note that the `--platform` argument is also available for `docker run`.

### <a name="containers-buildx"></a> Target multiple platforms with buildx

Using Docker's [buildx](https://docs.docker.com/buildx/working-with-buildx/), you can build images which target multiple platforms.
