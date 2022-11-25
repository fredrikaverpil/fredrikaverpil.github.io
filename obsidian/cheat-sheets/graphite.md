---
title: 🍇 Graphite
tags: [git, graphite, workflow]
draft: false
summary: "The stackable PR CLI tool."

# PaperMod
ShowToc: false
TocOpen: true

date: 2022-11-25T20:09:32+01:00
---

## Rebase and update stack

```bash
gt rs -fr && gt ss
```

The full list is available in the official [Graphite Documentation](https://docs.graphite.dev/guides/graphite-cli/command-shortcuts).

## Basic commands
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
