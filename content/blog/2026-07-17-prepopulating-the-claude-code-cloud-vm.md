---
title: "Prepopulating the Claude Code cloud VM with skills and config"
date: 2026-07-17
draft: true
tags: ["claude-code", "ai", "dotfiles"]
categories: []
---

I've grown fond of kicking off [Claude Code](https://www.claude.com/product/claude-code)
tasks from my phone. I'm out walking the dog, I think of something I want done
in one of my repos, I open the Claude mobile app, describe the task, and a
cloud VM picks it up and opens a pull request while I'm still outside. The
[Claude Code on the web](https://code.claude.com/docs/en/claude-code-on-the-web)
docs cover the full picture, but the short version is: each session runs in a
fresh, Anthropic-managed VM with your repository cloned into it.

The catch with a _fresh_ VM is exactly that: it's fresh. None of my local
setup is there. My user-level skills, my git identity, the little conventions I
lean on locally — none of it survives, because none of it lives in the repo.
The docs are explicit about this:

> Anything you've installed or configured only on your own machine isn't
> available in the session.

So the question I wanted to answer was: how do I get _my_ stuff into that VM
without babysitting every session? This post is how I do it, driven entirely
from a setup script and a couple of repo selections.

## What actually carries over

Before wiring anything up, it helps to know what you get for free. From a fresh
clone, the VM already has:

- Your repo's `CLAUDE.md`, `.claude/settings.json`, `.mcp.json`,
  `.claude/rules/`, and crucially `.claude/skills/`, `.claude/agents/`,
  `.claude/commands/` — all part of the clone.
- Skills you've enabled on claude.ai (those are loaded automatically).
- Your organization's server-managed settings.

And what you _don't_ get:

- Your user `~/.claude/CLAUDE.md`.
- Your user `~/.claude/skills/`, `~/.claude/agents/`, `~/.claude/commands/`.
- MCP servers you added locally with `claude mcp add`.
- Any static credentials.

The pattern is clear once you squint at it: **anything committed to a repo is
available; anything living only on your laptop is not.** So the trick to
prepopulating the VM is to put the things I want into a repo — and then make
sure that repo comes along for the ride.

## The idea: a config repo plus a setup script

I already keep a [dotfiles repo](https://github.com/fredrikaverpil/dotfiles).
It's the natural home for the config I want in every cloud session. Two pieces
make it work:

1. A small `.claude-web` package inside the repo that holds the project-level
   Claude config I want in the cloud (settings plus a curated subset of my
   skills).
2. A **setup script** that runs before Claude Code launches and wires that
   config into place.

A setup script is just a Bash script attached to the cloud _environment_ (not
the repo). It runs as root on Ubuntu 24.04 before Claude Code starts, and its
result is cached, so it doesn't slow down subsequent sessions. You configure it
in the environment settings dialog on claude.ai, in the **Setup script** field.

Mine is a one-liner that points at a script committed in my dotfiles repo:

```bash
bash /home/user/dotfiles/stow/shared/.claude-web/web-setup.sh || true
```

Two things worth noting about that single line:

- The path exists because **the repo is cloned before the setup script runs**.
  Cloud repos land under `/home/user/`, so a repo named `dotfiles` is at
  `/home/user/dotfiles`.
- The trailing `|| true` keeps a hiccup in my personal setup from blocking the
  whole session. If a setup script exits non-zero, the session fails to start,
  and I'd rather have a session without my tweaks than no session at all.

## The setup script

Here's the script the environment invokes. It does two jobs: wire my
project-level Claude config into the repo, and set a sane git identity on every
cloned repo.

```bash
#!/usr/bin/env bash
# shellcheck shell=bash
set -e

# Setup script for Claude Code on the web (the cloud sandbox).
#
# Scope is cloud-only, but guard anyway so a stray local invocation is a no-op.
if [ "${CLAUDE_CODE_REMOTE:-}" != "true" ]; then
  exit 0
fi

# Project-level Claude config for the dotfiles repo.
#
# Claude Code on the web reads <repo>/.claude, so link the .claude-web package
# into place at the repo root. The link is recreated on every run so it always
# tracks the current target; it is git-ignored, so it never exists on local
# checkouts.
repo_root="$(cd "$(dirname "${BASH_SOURCE[0]}")/../../.." && pwd)"
if [ -d "$repo_root/stow/shared/.claude-web" ]; then
  rm -rf "$repo_root/.claude"
  ln -s stow/shared/.claude-web "$repo_root/.claude"
fi

# Git identity, written LOCAL to each cloned repo (local always overrides the
# global identity the launcher asserts, and the launcher never touches local).
git_name="Fredrik Averpil"
git_email="fredrik.averpil@proton.me"
for gitdir in /home/user/*/.git; do
  [ -d "$gitdir" ] || continue
  repo="${gitdir%/.git}"
  git -C "$repo" config user.name "$git_name"
  git -C "$repo" config user.email "$git_email"
  git -C "$repo" config commit.gpgsign false
done
```

A few of the decisions in there are worth calling out.

`CLAUDE_CODE_REMOTE` is set to `true` in cloud sessions. Guarding on it means
that if I ever run this script by accident on my laptop, it exits immediately
and touches nothing.

The `.claude` symlink is the load-bearing part. Claude Code on the web reads
`<repo>/.claude`, so I keep the real content in `stow/shared/.claude-web`
(settings plus symlinks to a curated subset of my skills) and link it into
place at the repo root on every run. It's git-ignored, so it never materializes
on a local checkout where I don't want it.

The git identity loop writes **local** config to each repo under
`/home/user/`. I do this per-repo rather than globally on purpose: the launcher
re-asserts a global identity every session so its commits verify on GitHub, and
local config always wins without me having to fight that. The trade-off is that
my commits show as "Unverified" on GitHub (there's no personal signing key in
the sandbox), which I'm fine with.

> [!NOTE]
> The `.claude-web` package lives in
> [`stow/shared/.claude-web`](https://github.com/fredrikaverpil/dotfiles/tree/main/stow/shared/.claude-web)
> in my dotfiles repo, alongside the rest of my configuration. Nothing here is
> secret, which matters — see the caveats below.

## Selecting the repos when you start a session

Here's the part that makes it click, and the reason the whole thing works from
my phone with zero fiddling.

When you start a cloud session, you pick which repositories come along. I
select **two**:

1. My **dotfiles repo** — the one holding the `.claude-web` config and the
   setup script. This is what makes `/home/user/dotfiles/...` exist, so the
   setup script the environment points at is actually there to run.
2. The repo I actually want to **work in** — whatever project the task is
   about.

Both get cloned under `/home/user/`, and the setup script runs against them.
The nice surprise is that **skills from every attached repo are loaded at
once**. Each repo's `.claude/skills/` is picked up, so I end up with the union
of them in a single session.

I can vouch for this concretely, because this very post was written in a cloud
session with exactly that setup: my dotfiles repo and my blog repo were both
attached. In that one session I had my dotfiles skills (`git-commit`,
`golang-style`, and friends) _and_ my blog's own skills (`hugo-content`,
`hugo-design`) available side by side — the first set travelled in with the
config repo, the second came from the repo I was writing in.

That's the whole mental model:

> [!TIP]
> Treat one repo as your **travelling config** (skills, settings, git identity,
> a setup script) and attach it alongside whatever repo you actually want to
> code in. The setup script wires the config in; the skills from both repos
> stack.

So from the mobile app the flow is just: pick my dotfiles repo, pick the target
repo, describe the task. The environment's setup script does the rest, and
Claude starts with my conventions already in place.

## Caveats worth knowing

A few things I ran into or want to remember:

> [!WARNING]
> **There is no secrets store.** Both environment variables and setup scripts
> are visible to anyone who can edit the environment, and the setup script here
> lives in a repo. Keep anything sensitive out of both. My git email is about
> as private as it gets in mine.

- **Caching.** The setup script runs the first time you start a session in an
  environment; after that, Anthropic snapshots the filesystem and reuses it.
  The script re-runs when you change the script itself or the environment's
  allowed hosts, or after the cache expires (roughly seven days). Resuming an
  existing session never re-runs it. Good to know when you're iterating on the
  script and wondering why a change didn't take.
- **Keep it under ~5 minutes.** The setup script has to finish quickly enough
  for the cache to build. Mine is trivial, but if you're installing toolchains,
  parallelize with `&` and `wait`, or move a slow download into a
  [SessionStart hook](https://code.claude.com/docs/en/hooks#sessionstart).
- **Setup script vs. SessionStart hook.** The setup script is attached to the
  _environment_ and is cloud-only. A SessionStart hook lives in the repo's
  `.claude/settings.json` and runs both locally and in the cloud. I use the
  setup script here precisely because I want this behavior in the cloud only —
  but if you want something like `npm install` to run everywhere, a
  SessionStart hook (guarded on `CLAUDE_CODE_REMOTE` if needed) is the better
  home.
- **Network access.** The default **Trusted** access level allows the common
  package registries and GitHub. If your setup script needs to reach anything
  else, you'll need to add it to the allowlist — and changing the allowlist is
  one of the things that busts the setup-script cache.

## Summary

The cloud VM is a blank slate, and the only things that persist are what's in a
repo (plus claude.ai-enabled skills and org settings). So I lean into that:

1. Keep the config I want in the cloud inside a repo — for me, a `.claude-web`
   package in my dotfiles.
2. Point the environment's **Setup script** field at a script committed in that
   repo (with a trailing `|| true`), guarded on `CLAUDE_CODE_REMOTE`.
3. When starting a session, **select the config repo alongside the repo I want
   to code in**. Both get cloned, the setup script wires my config in, and
   skills from both repos load together.

The result is that I can start a properly kitted-out session from my phone —
my skills, my git identity, my conventions — without touching the environment
config each time. Fresh VM, familiar setup.
