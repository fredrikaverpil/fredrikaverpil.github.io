---
ShowToc: false
TocOpen: false
date: 2015-12-28 02:00:12+01:00
draft: false
tags:
- python
title: Python simple web server
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