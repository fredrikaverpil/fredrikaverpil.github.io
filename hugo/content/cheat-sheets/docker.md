---
ShowToc: false
TocOpen: true
date: 2022-12-18 01:49:16+01:00
draft: false
summary: Notes to self, snippets etc.
tags:
- docker
title: "\U0001F433 Docker"
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
RUN --mount=type=ssh <COMMAND>
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