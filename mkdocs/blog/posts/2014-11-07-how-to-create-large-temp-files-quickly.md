---
date: 2014-11-07
authors:
  - fredrikaverpil
comments: true
tags:
- bash
---

# How to create large temp files quickly (for testing purposes)

Quick and dirty way to just create a 10GB temp file for testing e.g. network transfer speeds.

<!-- more -->

### Linux

```bash
fallocate -l 10G temp_10GB_file
```

### Windows

The file size is defined in bytes. Use Google to do the conversion if you’re unsure.

```bat
fsutil file createnew temp_10GB_file 10000000000
```


### Mac OS X

```bash
mkfile -n 10g temp_10GB_file
```