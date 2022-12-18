---
title: 🐳 Docker
tags: [docker]
draft: false
summary: "Notes to self, snippets etc."

# PaperMod
ShowToc: false
TocOpen: true

date: 2022-12-18T01:49:16+01:00
---

## SSH access

Make your SSH keys visible to the SSH agent:

```bash
# Linux
ssh-add -K

# macOS
ssh-add --apple-use-keychain
```

Verify your keys are visible to the SSH agent:

```bash
ssh-add -L
```

Define a `RUN` command in `Dockerfile` which should have access to your SSH keys:

```Dockerfile
RUN RUN --mount=type=ssh <COMMAND>
```

Build:

```bash
docker build --ssh default .
```

Or use Docker Compose (v2):

```yaml
services:
  my-service:
    build:
      dockerfile: Dockerfile
      context: .
      ssh:
        - default

  ...
```