---
title: Distributing Python script(s) as zip file
tags: [python, maya, nuke]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

created: 2018-06-07T02:00:12+02:00
updated: 2022-11-15T17:29:41+01:00
---

A recent discussion on [3DPRO](http://3dpro.org/) sparked me to scribble down some ideas on how to somewhat painlessly distribute a Python package to be run in DCC applications such as Maya or Nuke as simply as possible. So this is an alternative to building a wheel and mucking around with virtual environments.

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

_zipfile = 'a.zip'
for _path in sys.path:
    if os.path.isdir(_path) and _zipfile in os.listdir(_path):
        sys.path.append(os.path.join(_path, _zipfile))
        import b

from b import c
```

For Maya, you can ask your users to [create a shelf button](http://help.autodesk.com/view/MAYAUL/2018/ENU/?guid=GUID-527023AE-9FB5-4D01-8D29-075B1E6C4754) with the code above (or even bundle a shelf MEL file, to be placed in e.g. `~/Documents/maya/2018/prefs/shelves`).

Bonus: add a `__main__.py` to the zip file's root and you can execute it with `python a.zip`.

Make sure to vendor (bundle) any non-standard library dependency Python scripts your program might need. What's nice about vendoring is then you don't have to worry about the same vendored package (of a different version) already being used by some other script, causing a version clash. Vendor under your script's namespace, e.g. `b.vendor.Qt` when using the above examples if vendoring e.g. [Qt.py](https://github.com/mottosso/Qt.py). Your script would then import its vendored module like so: `from b.vendor import Qt` instead of the regular `import Qt`.

And if you wish to obfuscate your code a little bit, put only the `.pyc` bytecode files in the zip file. These can be decompiled, but at least that takes some effort by the user.
