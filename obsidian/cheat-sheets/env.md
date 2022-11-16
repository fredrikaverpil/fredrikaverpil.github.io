---
title: 🐚 Shell
tags: [shell]
draft: false
summary: "Notes to self, snippets etc."

# PaperMod
ShowToc: false
TocOpen: true

updated: 2022-11-16T00:53:51+01:00
created: 2022-11-14T20:42:48+01:00
---

## Load variables from .env file into current environment
Load an .env file into the environment prior to running something which requires the environment variables:
```bash
set -a
source <somefile.env>
set +a
```