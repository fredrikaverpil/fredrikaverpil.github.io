---
layout: post
title: Github commands
tags: [github]
---

Common commands I'm using all the time, but keep forgetting the syntax for.

<!--more-->

## Pull and overwrite any local changes

```bash
git fetch --all
git reset --hard origin/master
```

Or if you're on a different branch:

```bash
git reset --hard origin/your_branch
```

## Revert repository to previously pushed commit

```bash
git reset 306410193e675ac9486fea07a1de04971f51e1a4
git push --force
```

## Change the current branch into master

For more details, see [Stack Overflow](http://stackoverflow.com/questions/2763006/change-the-current-branch-to-master-in-git).

```bash
git checkout better_branch
git merge --strategy=ours master    # keep the content of this branch, but record a merge
git checkout master
git merge better_branch             # fast-forward master up to the merge
```
