---
title: Search and replace
tags: [python]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

date: 2014-02-16T01:00:12+01:00
---

A very simple and quite rudimental search and replace script with a UI. It will only search and replace _contents_ of files and not the filenames.



It won’t win any design awards but it could prove very useful when being run on large amounts of huge files as it will only keep the current file it is processing in the RAM.

Be very careful before running this script and make sure to read through the readme before using. And ALWAYS make a backup of what you intend to process. You can download this over at [GitHub](https://github.com/fredrikaverpil/searchReplace).

The script was written in Python and utilizes PySide/PyQt (choose which one you wish to use in upper portion of the Python script). Since it uses my boilerplate in [[2013-10-06-a-pyside-pyqt-boilerplate-for-maya-nuke-or-standalone]], you can also launch it from within Maya or Nuke...
