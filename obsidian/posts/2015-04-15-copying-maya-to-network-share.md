---
title: Copying Maya to network share
tags: [maya, bash]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

created: 2015-04-15T02:00:12+02:00
updated: 2022-11-15T22:29:17+01:00
---

If you would like to run a thin client installation of Maya, which means you run it off a network share rather from a local installation, you need to make sure to copy symlinks on Linux.



If you just perform a simple copy, you’ll probably get something like this:

    error: unpacking of archive failed on file /my_mount/maya_installation/lib/libGLEW.so;544612b1: cpio: symlink failed - Operation not supported

The solution is to copy the files which are symlinked:

```bash
# Linux
cp -vrLp /src /dst

# OS X
cp -vRL /src /dst
```
