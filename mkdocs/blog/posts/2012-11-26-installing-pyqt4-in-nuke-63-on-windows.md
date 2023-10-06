---
date: 2012-11-26
authors:
  - fredrikaverpil
comments: true
tags:
- nuke
- python
- windows
---

# Installing PyQt4 in Nuke 6.3 on Windows

A few steps to get PyQt working inside of Nuke 6.3 (Windows only for now).

<!-- more -->

Please note that you need your own menu.py in place for this to work. In case you need this set up first, check out one of my previous articles, [Nuke 6.3 Small studio setup for Win/Mac](2011-10-28-nuke-63-small-studio-setup-for-windows-osx.md).

1. Install 64-bit Python 2.6 from [here](http://www.python.org/download/releases/2.6/) into C:\python26.
2. Install the PyQt 4.6 snapshot from [here](http://code.google.com/p/pyqt4-win64-binaries/downloads/detail?name=PyQt-Py2.6-gpl-4.6-snapshot-20090810-1.exe&can=2&q=) into C:\python26_64 and make sure to untick the “Qt runtime” checkbox.
3. Copy all the files inside of C:\Python26_64\Lib\site-packages to <Nuke installation>\lib\site-packages or to a server location of your choice.

If you copied the files to a folder residing on the server, add the following to your menu.py (and modify):

```python
sys.path.append ("X:/server/path/to/copied/site-packages")
```

PyQt should now run inside of Nuke 6.3, to be used like this:

```python
from PyQt4 import QtGui
label = QtGui.QLabel("Hello World")
label.show()
```

Please note that you could uninstall both Python 2.6 and the PyQt snapshot if you wish. The files needed have been added to Nuke.