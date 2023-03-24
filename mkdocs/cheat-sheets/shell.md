---
date: 2022-12-17
draft: true
tags:
- shell
title: "\U0001F41A Shell"
---

# Shell

## Load variables from .env file into current environment
Load an .env file into the environment prior to running something which requires the environment variables:
```bash
set -a
source <somefile.env>
set +a
```