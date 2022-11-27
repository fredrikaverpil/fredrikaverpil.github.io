---
title: 🐚 Shell
tags: [shell]
draft: false
summary: "Notes to self, snippets etc."

# PaperMod
ShowToc: false
TocOpen: false

date: 2022-11-27T13:42:55+01:00
---

## Load variables from .env file into current environment
Load an .env file into the environment prior to running something which requires the environment variables:
```bash
set -a
source <somefile.env>
set +a
```