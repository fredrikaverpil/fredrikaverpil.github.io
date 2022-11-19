---
title: How to create large temp files quickly (for testing purposes)
tags: [bash]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

date: 2014-11-07T01:00:12+01:00
---

Quick and dirty way to just create a 10GB temp file for testing e.g. network transfer speeds.



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
