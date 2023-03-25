---
date: 2017-09-30
tags:
- python
- pyside
---

# Contributing to PySide2

This is reminder-to-self about how to get set up and contribute to PySide2 using Gerrit. It could also be a fun read "on the bus" before actually setting this up yourself, to get an overview on what's required to get up and running with Gerrit.

<!-- more -->

**Disclaimer**: I discourage you to blindly follow these _highly personal_ notes without reading through the official docs when you actually get set up yourself. Anything can be changed in the official docs at any time, which isn't reflected here. The point of this post is just to make a note-to-self and offer a "bigger picture" before actually digging in.



### RTFM

If you want to read it all through yourself, start here: [https://wiki.qt.io/Gerrit_Introduction](https://wiki.qt.io/Gerrit_Introduction)


### Difference between code.qt.io and codereview.qt.io

The PySide2 git repository can be accessed through both these websites:

- http://code.qt.io/cgit/pyside/pyside-setup.git/
- https://codereview.qt-project.org/gitweb?p=pyside%2Fpyside-setup.git;a=summary

However, the `codereview` one [should only be used for gerrit-access, and not for e.g. git cloning](https://gitter.im/PySide/pyside2?at=59a01d95614889d475869e8e).


### Clone PySide2

I like to clone this repo into a `_gerrit` folder.

```bash
# Clones the repo into ~/code/repos/_gerrit/pyside-setup
git clone --recursive --branch 5.6 https://code.qt.io/pyside/pyside-setup.git ~/code/repos/_gerrit/pyside-setup
```


### Set up Gerrit

Here's the [full docs on setting up Gerrit](https://wiki.qt.io/Setting_up_Gerrit), which you should probably read through.

When contributing to PySide (using Gerrit), you need to set up SSH keys. I consulted the [Github SSH key generation docs](https://help.github.com/articles/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent/) to do this:

```bash
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
```

I like to call my keys `qt_gerrit_id_rsa.pub` and `qt_gerrit_id_rsa` and have them placed in `~/.ssh/`.

Then the private key needs to be added to `~/.ssh/config`:

```bash
Host codereview.qt-project.org
Port 29418
User <your_qt_username>
IdentityFile ~/.ssh/qt_gerrit_id_rsa
Ciphers +aes256-cbc
```

If you don't know what to put in `<your_qt_username>`, you need to [set up an account with Qt](https://login.qt.io/register). Also make sure you set up the "New Contributor Agreement".


### Set up Git

There are mandatory settings for both Windows and Unix users. Instead of re-writing them here, I'll link to the [full docs on setting up Gerrit](https://wiki.qt.io/Setting_up_Gerrit). Scroll down to ["Configuring Git"](https://wiki.qt.io/Setting_up_Gerrit#Configuring_Git).

The docs cover how to globally change your __global__ writing options. This is not something I personally wish to do. Instead I would like to abide their required settings only in the Gerrit repositories (`pyside-setup`, `qt5` etc). Therefore I deliberately change the `--global` argument into `--local` for all `git config`commands and execute them inside of each repository.

```bash
# Example
cd ~/code/repos/_gerrit/pyside-setup
<git config --local ...>
```

Then there's a commit template which Qt recommends we use, and it requires a Qt git clone:

```bash
git clone --recursive --branch 5.6 https://code.qt.io/qt/qt5.git ~/code/repos/_gerrit/qt5
cd ~/code/repos/_gerrit/pyside-setup
git config --local commit.template ~/code/repos/_gerrit/qt5/commit-template
```

**Note**: Again, since all `git config` commands are local to the `pyside-setup` repository, you probably also want to apply them to the `qt5` repository you just cloned.


### Git branches and submodules

In my previous `git clone` commands, I'm defining a branch. In case you want to change branch, you'll need to also update all submodules:

```bash
cd ~/code/repos/_gerrit/pyside-setup
git checkout 5.6
git submodule update --init --recursive
```

You can also replace `5.6` with a commit SHA, but don't forget to then update the submodules afterwards.


### Add remote

A remote (SSH) to Gerrit is required to push changes:

```bash
cd ~/code/repos/_gerrit/pyside-setup
git remote add gerrit ssh://codereview.qt-project.org/pyside/pyside-setup
```


### Add git hooks

Instead of just copy-pasting the instructions, see ["Setting up git hooks"](https://wiki.qt.io/Setting_up_Gerrit#Setting_up_git_hooks) for the full details.

You'll need to set up a git hook as there's no `init_repository` script for PySide2 which would usually handle this (in e.g. the `qt5` repository). This essentially means downloading the hook file into the `.git/hooks` folder of the `pyside-setup` repository.

I do it like described in the docs, which also verifies that the SSH keys are set up and working:

```bash
cd ~/code/repos/_gerrit/pyside-setup
curl -o .git/hooks/commit-msg http://codereview.qt-project.org/tools/hooks/commit-msg
gitdir=$(git rev-parse --git-dir); scp -p codereview.qt-project.org:hooks/commit-msg ${gitdir}/hooks/
```

### Make changes to the PySide2 code

I like to create a new branch in which I make my changes:

```bash
git checkout -b <name_of_branch>
```

### Pushing local changes to Gerrit

Official docs on this is found [in the docs](https://wiki.qt.io/Gerrit_Introduction). [QtC](http://qt.io) then has a [commit policy](https://wiki.qt.io/Commit_Policy) which you may want to read.

You push using their `ref` target. Make sure to define which branch you intend to push to, for example `5.6`.

```bash
# Push to the "5.6" branch
git push gerrit HEAD:refs/for/5.6
```

If you're contributing to the `5.6` branch and the change should also be applied to other branches, someone can later forward merges to other branches.


### Review workflow

This is best described in the [official docs](https://wiki.qt.io/Gerrit_Introduction#Review_Workflow). If you've gotten this far, you're on your own! ;)