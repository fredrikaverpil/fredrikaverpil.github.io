---
title: Testing and sandboxing with Docker
tags: [docker]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

created: 2016-09-12T03:00:12+02:00
updated: 2022-11-15T17:29:41+01:00
---

A quick way to enter an interactive docker container:

    docker run --rm --interactive --tty -v /localfolder:/containerfolder centos:7

On Windows, use forward slashes for the directory mapping.

For a more complex setup, have a look at [sandbox-docker](https://github.com/fredrikaverpil/sandbox-docker).
