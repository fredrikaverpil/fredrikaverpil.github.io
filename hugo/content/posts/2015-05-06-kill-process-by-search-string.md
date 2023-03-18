---
ShowToc: false
TocOpen: false
date: 2015-05-06 02:00:12+02:00
draft: false
tags:
- linux
- bash
title: Kill process by search string
---

In Linux, you can kill all processes by name (or by username etc) using something like this:

```
kill -9 $(ps aux | grep 'some_process_name' | awk '{print $2}')
```