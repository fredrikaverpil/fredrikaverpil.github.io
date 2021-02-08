---
layout: post
title: 'Git gotchas'
tags: [Linux]
---

Git: Generally Impossible To (remember?) ...

Anyways, some handy git tricks in this post.

<!--more-->

## Branch management

When you've accidentaly screwed up your local `master`, make it good again by resetting it into whatever is in `origin/master`:

```bash
git checkout -B master origin/master && git pull
```

Delete all branches which have been merged (except e.g. master, main, dev):

```bash
git branch --merged | egrep -v "(^\*|master|main|dev)" | xargs git branch -d
```

## Search

Free text search throughout any git commit message:

```bash
git log -S<regexp>
```

Find a commit based on free text search of code:

```bash
git rev-list --all | xargs git grep <regexp>
```

## Rebasing

### Better rebase

Rebase that for some reason works more often than `git rebase`:

```bash
git pull --rebase origin master
```

### Too complicated to rebase

This will stage all changed files since the given `good-commit` on top of whatever is in `origin/master`.

```bash
git checkout good-commit
git checkout -B revert-stuff
git reset --soft origin/master
```
