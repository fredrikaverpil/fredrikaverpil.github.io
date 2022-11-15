---
title: Python simple web server
tags: [python]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

created: 2015-12-28T02:00:12+01:00
updated: 2022-11-15T17:29:41+01:00
---

The absolutely fastest way to get a simple web server up and running using
Python 3, for development purposes.

```bash
cd my_web_root
python -m http.server
```

Or if you are on Python 2.x:

```bash
cd my_web_root
python -m SimpleHTTPServer 8000
```

Then just access `http://your-ip:8000` to access the web server contents. Hit ctrl+c to exit.
