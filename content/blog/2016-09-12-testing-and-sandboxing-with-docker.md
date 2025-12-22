---
title: "Testing and sandboxing with Docker"
date: 2016-09-12
tags: ["docker"]
categories: ["archive"]
---

A quick way to enter an interactive docker container:

    docker run --rm --interactive --tty -v /localfolder:/containerfolder centos:7

On Windows, use forward slashes for the directory mapping.

For a more complex setup, have a look at [sandbox-docker](https://github.com/fredrikaverpil/sandbox-docker).
