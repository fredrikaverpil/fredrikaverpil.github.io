---
ShowToc: false
TocOpen: false
date: 2011-11-12 01:00:12+01:00
draft: false
tags:
- nuke
- python
title: Nuke script
---

Launch the source directory of the selected Read or Write node.



The script will take the selected Read or Write node’s file path, and open it using Explorer on Windows or Finder on Mac OS X. Optionally, you can feed the script with your own path of choice and make that open up. The script was built for Nuke 6.3 and has support for filenamefilter callback.

## Installation

Download: [browseDir.py](https://github.com/fredrikaverpil/nuke/raw/master/scripts/browseDir.py) (v1.1, 2013-10-13)

- v1.1: Added support for space in the filepath
- v1.0: Initial release

Place the Python script in the /scripts dir inside your NUKE_PATH (see my previous post on setting this up). Add the following to your menu.py:

```python
# import browseDir
nuke.menu( 'Nuke' ).addCommand( 'My file menu/Browse/Node\'s file path', "browseDir.browseDirByNode()", 'shift+b' )
nuke.menu( 'Nuke' ).addCommand( 'My file menu/Browse/Scripts folder', "browseDir.browseDir('scripts')" )
```

You should now be able to select any Read or Write node and hit Shift + B (B for “browse”) to launch its file path source folder. A file menu will be added from which you can open up the current opened script’s folder.

## Customizing the script

You can also run launch(path) with your own path of choice to make that open up. But if you’re digging in the script you might as well just add another elif statement into the browseDir() function and call it like I do from menu.py.

An example of adding a file menu option to open the shot folder; add to menu.py:

```python
elif action == 'shot':
	for i in range(0, (len(scriptPathSplitted)-3) ):
		openMe = openMe + scriptPathSplitted[i] + '/'
```

Please note: this piece of code is already in the right place inside the browseDir() function, so you should easily be able to find it and modify it. Also, there is no way I can guarantee that this piece of code will actually launch a folder on your machine. It all depends on what your folder structure looks like. :)