---
title: A PySide/PyQt4 boilerplate for Maya, Nuke or standalone
tags: [python, maya, nuke, pyside, pyqt]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

created: 2013-10-06T02:00:12+02:00
updated: 2022-11-15T22:29:17+01:00
---

Ever wanted to be able to run the same user interface in Maya, Nuke as well as completely standalone (with or without app-specific modifications) and on any platform, using PySide and/or PyQt?



– That’s exactly why I created a boilerplate Python script, which could serve as a good starting point for most PySide/PyQt projects to be used in the VFX pipeline.

As my experience is limited to some VFX applications, you are most welcome to fork or contribute to make it also run in other PyQt/PySide enabled applications.

### Quickstart

1. Download the files from [GitHub](https://github.com/fredrikaverpil/pyVFX-boilerplate) and edit boilerplate.py
2. Set the variable QtType to either ‘PySide’ or 'PyQt’
3. Edit in the path to the pysideuic module (if you’re using PySide/Nuke and loading a .ui file)
4. Edit the variable uiFile to point to your .ui file
5. Start coding your app in the HelloWorld class

### Run the script in Maya

```python
import sys
sys.path.append('c:/path/to/dir/containing/the/script')
import boilerplate
boilerplate.runMaya()
```

### Run the script in Nuke

```python
import sys
sys.path.append('c:/path/to/dir/containing/the/script')
import boilerplate
boilerplate.runNuke()
```

### Run the script as standalone

    python boilerplate.py


More on usage and customization over at the [GitHub project’s wiki](https://github.com/fredrikaverpil/pyVFX-boilerplate/wiki).
