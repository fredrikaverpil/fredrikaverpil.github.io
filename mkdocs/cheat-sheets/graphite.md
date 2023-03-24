---
date: 2022-12-17
draft: true
tags:
- git
title: Graphite
---

# Graphite

## Combos

### Pull down default branch, remove merged branches

```bash
gt repo sync -fr
```

Add `gt ss` to restack all current stacks onto the newly pulled down default branch:

```bash
gt repo sync -fr && gt ss
```


## Shortcuts

The full list is available in the official [Graphite Documentation](https://docs.graphite.dev/guides/graphite-cli/command-shortcuts).

| command                  | shortcut        |
| ------------------------ | --------------- |
| `gt log`                 | `gt l`          |
| `gt log short`           | `gt ls`         |
| `gt branch track`        | `gt btr`        |
| `gt branch checkout`     | `gt bco`        |
| `gt branch up [steps]`   | `gt bu [steps]` |
| `gt branch down [steps]` | `gt bd [steps]` |
| `gt branch top`          | `gt bt`         |
| `gt branch bottom`       | `gt bb`         |
| `gt branch create`       | `gt bc`         |
| `gt commit create`       | `gt cc`         |
| `gt commit amend`        | `gt ca`         |
| `gt stack submit`        | `gt ss`         |
| `gt repo sync`           | `gt rs`         |

## Rebasing

| command              | shortcut |
| -------------------- | -------- |
| `gt upstack restack` | `gt usr` |
| `gt upstack onto`    | `gt uso` |
| `gt upstack onto`    | `gt uso` |