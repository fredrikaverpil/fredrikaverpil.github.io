---
title: Installing Qt.py (advanced methods)
tags: [python, maya, nuke, qtpy, pyside, pyqt]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

created: 2016-07-25T02:00:12+02:00
updated: 2022-11-15T22:29:17+01:00
---

[Qt.py](https://github.com/mottosso/Qt.py) can be installed in many ways, depending on your needs. This post aims to outline some common approaches:

* Install using pip
* Install into an application's custom Python build
* Vendoring
* Make Qt.py available using `sys.path` and `site.addsitedir`



## Install using pip

The probably most common way to install Qt.py (and as mentioned in the project's [README](https://github.com/mottosso/Qt.py#install)) is to install via [pip](http://pip.readthedocs.io):

    pip install Qt.py

Please note that it's not recommended to "pip install" into your operating system's default `site-packages`. You could look to solutions such as [virtualenv](https://virtualenv.pypa.io) or [Conda](http://conda.pydata.org) for this.

Please also note that this method doesn't automatically make Qt.py available in e.g. Maya or Nuke as those applications use their own distribution of Python.

If you want to be up really quickly with Qt.py you can also install PyQt5 which offers a wheel (!) for Python 3:

    pip3 install PyQt5 Qt.py


## Install into an application's custom Python build

You can place the [`Qt.py` file](https://raw.githubusercontent.com/mottosso/Qt.py/master/Qt.py) manually within an application which comes with its own `site-packages` folder (such as Maya or Nuke).

#### Maya

    Windows: C:\Program Files\Autodesk\maya2017\Python\Lib\site-packages
    Linux: /usr/autodesk/maya2017/lib/python2.7/site-packages
    macOS: /Applications/Autodesk/maya2017/Maya.app/Contents/Frameworks/Python.framework/Versions/2.7/lib/python2.7/site-packages

#### Nuke

    Windows: C:\Program Files\Nuke10.0v1\lib\site-packages
    Linux: /usr/local/Nuke10.0v1/lib/python2.7/site-packages
    macOS: /Applications/Nuke10.0v2/Nuke10.0v2.app/Contents/Frameworks/Python.framework/Versions/2.7/lib/python2.7/site-packages

However, I actually don't recommend this approach for a number of reasons (mostly risk of confusion):

* It's no longer obvious, from reading your code, where Qt.py actually resides.
* It's possible that you forget that you did put Qt.py here, and are trying to load a newer version of Qt.py via e.g. vendoring, but instead the Qt.py from within this custom Python build is being used.
* For every new version of this main application, you'd have to remember to re-install Qt.py.

## Vendoring

You can download the [`Qt.py` file](https://raw.githubusercontent.com/mottosso/Qt.py/master/Qt.py) and bundle it with your application's source tree. Something like this:

```
.
├── myapp.py
└── vendor
    └── Qt.py
```

The `site` package can be used to specify where Qt.py can be loaded from. For example, `myapp.py` in the above source tree could use this to be able to use Qt.py:

```python
# myapp.py
import os
import site

QTPY_PATH = os.path.join(os.path.dirname(os.path.abspath(__file__)),
                         'vendor')
site.addsitedir(QTPY_PATH)

import Qt
```

Please note, you can use pip to install Qt.py into the vendor folder by specifying the `--target` option:

    pip install --target=vendor Qt.py


You could also "vendor" a clone (or fork) of the Qt.py repository. This approach is nice if you intend to contribute to Qt.py and run tests. A Github repository will also offer the possibility to quite easily perform updates to Qt.py via `git pull`.

You can then bundle the Qt.py repository with your source tree and from within your Python script use e.g. the `site` package to make Qt.py available.

    git clone https://github.com/mottosso/Qt.py

```
.
├── myapp.py
└── vendor
    └── Qt.py
        ├── build_caveats_tests.py
        ├── CAVEATS.md
        ├── CONTRIBUTING.md
        ├── Dockerfile
        ├── LICENSE
        ├── parser.py
        ├── Qt.py
        ├── README
        ├── tests.py
        └── tests.py
```

```python
# myapp.py
import os
import site

QTPY_PATH = os.path.join(os.path.dirname(os.path.abspath(__file__)),
                         'vendor', 'Qt.py')
site.addsitedir(QTPY_PATH)

import Qt
```


## Make Qt.py available using `sys.path` and `site.addsitedir`

In cases where Qt.py is simply not found...

```python
>>> import Qt
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
ImportError: No module named Qt
```

...you can make it available via e.g. the `site` standard Python module:

```python
>>> import site
>>> site.addsitedir(PATH_TO_QTPY)
>>> import Qt
>>> print(Qt.__binding__)
PySide2
```

Or you can add the Qt.py path to `PATH` via `sys.path.append` or `sys.path.insert`:

```python
>>> import sys
>>> sys.path.append(PATH_TO_QTPY)
>>> import Qt
```

```python
>>> import sys
>>> sys.path.insert(0, PATH_TO_QTPY)
>>> import Qt
```
