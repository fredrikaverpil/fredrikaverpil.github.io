---
title: "Kill process by search string"
date: 2015-05-06
tags: ["linux", "bash"]
---

In Linux, you can kill all processes by name (or by username etc) using something like this:

```
kill -9 $(ps aux | grep 'some_process_name' | awk '{print $2}')
```
