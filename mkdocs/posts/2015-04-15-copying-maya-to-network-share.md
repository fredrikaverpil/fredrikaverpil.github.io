---
date: 2015-04-15
tags:
- maya
- bash
---

# Copying Maya to network share

If you would like to run a thin client installation of Maya, which means you run it off a network share rather from a local installation, you need to make sure to copy symlinks on Linux.

<!-- more -->

If you just perform a simple copy, you’ll probably get something like this:

    error: unpacking of archive failed on file /my_mount/maya_installation/lib/libGLEW.so;544612b1: cpio: symlink failed - Operation not supported

The solution is to copy the files which are symlinked:

```bash
# Linux
cp -vrLp /src /dst

# OS X
cp -vRL /src /dst
```