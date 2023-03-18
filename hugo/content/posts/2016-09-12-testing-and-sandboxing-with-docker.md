---
ShowToc: false
TocOpen: false
date: 2016-09-12 03:00:12+02:00
draft: false
tags:
- docker
title: Testing and sandboxing with Docker
---

A quick way to enter an interactive docker container:

    docker run --rm --interactive --tty -v /localfolder:/containerfolder centos:7

On Windows, use forward slashes for the directory mapping.

For a more complex setup, have a look at [sandbox-docker](https://github.com/fredrikaverpil/sandbox-docker).