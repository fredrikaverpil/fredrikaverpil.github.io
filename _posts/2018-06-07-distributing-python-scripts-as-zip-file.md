---
layout: post
title: Distributing Python script(s) as zip file
tags: [python, maya, nuke]
---

A recent discussion on [3DPRO](http://3dpro.org/) sparked me to scribble down some ideas on how to somewhat painlessly distribute Python script(s) to be run in DCC applications such as Maya or Nuke as simply as possible.

<!--more-->

You can make Python import your modules or packages even when zipped. Example, where `b` is our package and `c.py` is our actual program we want to run:

```
a.zip
└── b <directory>
    ├── c.py
    └── __init__.py
```

This can be imported like so:

```python
import sys
sys.path.append('a.zip')
import b
from b import c
```

Since `a.zip` just acts like any folder in this case, you can include your version number in the zip filename, or any other arbitrary information such as author name etc. Any way, you should probably not call it `a.zip`, which I've just done here for brevity reasons.

You can ask your users to drop the zip file into their local scripts folder (e.g. `~/Documents/maya/scripts` for Maya or `~/.nuke` for Nuke) and to run the following code, which will locate the `a.zip`, import it and execute the program:

```python
import sys
import os

zip = 'a.zip'
for path in sys.path:
    if os.path.isdir(path) and zip in os.listdir(path):
        sys.path.append(os.path.join(path, zip))
        import b

from b import c
```

For Maya, you can ask your users to [create a shelf button](http://help.autodesk.com/view/MAYAUL/2018/ENU/?guid=GUID-527023AE-9FB5-4D01-8D29-075B1E6C4754) with the code above (or even bundle a shelf MEL file, to be placed in e.g. `~/Documents/maya/2018/prefs/shelves`).

Bonus: add a `__main__.py` to the zip file's root and you can execute it with `python a.zip`.
