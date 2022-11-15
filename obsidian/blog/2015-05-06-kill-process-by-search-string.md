---
title: Kill process by search string
tags: [linux, bash]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

created: 2015-05-06T02:00:12+02:00
updated: 2022-11-15T17:29:41+01:00
---

In Linux, you can kill all processes by name (or by username etc) using something like this:

```
kill -9 $(ps aux | grep 'some_process_name' | awk '{print $2}')
```
