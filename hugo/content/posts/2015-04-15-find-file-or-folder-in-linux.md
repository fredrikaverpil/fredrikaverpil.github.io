---
ShowToc: false
TocOpen: false
date: 2015-04-15 02:00:12+02:00
draft: false
tags:
- bash
- linux
title: Find file or folder in Linux
---

Ever needed to do a simple search for an application, a file or a folder in Linux and when `whereis` doesn’t return anything useful?

```bash
sudo find / -name "some_folder"
```