---
ShowToc: false
TocOpen: false
date: 2016-07-25 02:00:12+02:00
draft: false
tags:
- python
- maya
- nuke
- qtpy
- pyside
- pyqt
title: Developing with Qt.py
---

This post aims to give an introduction to the [Qt.py](https://github.com/mottosso/Qt.py) project and how to get set up with it for PySide/PyQt4 and PySide2/PyQt5 development.



# Contents

* What is Qt.py
* How does it work?
* Installing Qt.py
* Caveats
* Contributing
* Closing comments


## What is Qt.py?

Qt.py is a Minimal Python 2 & 3 [shim](https://en.wikipedia.org/wiki/Shim_(computing)) around all Qt bindings - PySide, PySide2, PyQt4 and PyQt5 – which enables you to write software that dynamically chooses the most desireable bindings based on what's available.

Basically, what this means is instad of you importing e.g. `PySide`, you import the Qt.py Python module which is done like this:

```python
import Qt
```

From it, you can import all of PySide's modules, for example `QtGui` and `QtCore`, just like you're used to:

```python
from Qt import QtGui, QtCore
```

Since you're not explicitly specifying PySide or PyQt specifically, your code can, if abiding certain rules when coding with Qt.py, run regardless of whether you have PySide or PyQt installed. Qt.py will detect which Python binding you have installed and use that.

Please note that you can tell Qt.py which bindings to prefer by setting the `QT_PREFERRED_BINDING` environment variable.

When using Qt.py, you must follow a set of rules to make it work. For example, by design you are expected to write as if you were writing for PySide2. Here's an example:

```python
import sys
from Qt import QtWidgets

app = QtWidgets.QApplication(sys.argv)
button = QtWidgets.QPushButton("Hello World")
button.show()
app.exec_()
```

The code above will run successfully when using Qt.py, regardless of whether you are on Windows, macOS or Linux, Python 2 or Python 3, PySide, PyQt4, PySide2 or PyQt5.

For those of you new to Qt5 (and PySide2 or PyQt5), all widget classes of `QtGui` were moved into its own module called `QtWidgets`. I would dare to say this is one of the largest backwards compatibility breaking changes you will have to deal with when changing existing scripts to use Qt.py.

There are of course differences between all the Python bindings which Qt.py does not handle. When you stumble upon this, you can query which binding is being used and create binding-specific code for such scenarios. For example, this deals with Qt4 vs Qt5:

```python
from Qt import __binding__

if __binding__ in ('PySide2', 'PyQt5'):
    print('Qt5 binding available')
elif __binding__ in ('PySide', 'PyQt4'):
    print('Qt4 binding available.')
else:
    print('No Qt binding available.')
```

For more information on how to make your old PySide/PyQt4 scripts work with Qt.py, please have a look at one of my previous posts: [2016-07-25-dealing-with-maya-2017-and-pyside2]({{< ref "2016-07-25-dealing-with-maya-2017-and-pyside2" >}}).

Qt.py was created and is maintained by Marcus Ottosson and myself. At the time of writing this, Qt.py is being used at [Disney Animation](http://www.disneyanimation.com), [Framestore](https://www.framestore.com), [Industriromantik](http://www.industriromantik.se) (where I work) and [Weta Digital](https://www.wetafx.co.nz).


## How does it work?

In short, Qt.py replaces itself with whatever Python binding is available (or preferred) in the Python environment your script runs in. However, Qt.py will first take care of remappings of calls so that the Qt5 style syntax of PySide2 will work even if you run this code in an environment with a Python binding for Qt4. Some enhancements such as automatically enabling sip API v2 for PyQt4 and providing convenience functions are also features of Qt.py to increase code compatibility between bindings.

Please see the [README](https://github.com/mottosso/Qt.py#how-it-works) in the project repository for more information.

The pure-Python module is contained within the [`Qt.py` file](https://raw.githubusercontent.com/mottosso/Qt.py/master/Qt.py) and I recommend reading it through for increased understanding on how Qt.py works. It's actually not that big and in my humble opinion not that complex.


## Installing Qt.py

First off, you should make sure you've got either PySide, PySide2, PyQt4 or PyQt5 installed in your Python environment. Then you can proceed with installing Qt.py.

The tricky part is to make sure your script can find Qt.py. There are many alternatives to achieve this but the perhaps simplest one is to follow these two steps:

1. Download this [`Qt.py` file](https://raw.githubusercontent.com/mottosso/Qt.py/master/Qt.py) and save it into a folder.
2. Add the folder path to the `PYTHONPATH` environment variable.

I've also outlined a few advanced (and often better) approaches on how to get Qt.py installed based on different needs in this post: [2016-07-25-installing-qt-py-advanced-methods]({{< ref "2016-07-25-installing-qt-py-advanced-methods" >}}).


## Caveats

There are some "gotchas" or caveats to be aware of. We are keeping track of them in the [`CAVEATS.md` file](https://github.com/mottosso/Qt.py/blob/master/CAVEATS.md).

The `CAVEATS.md` markdown file is actually created in such a way that it can be parsed and included in the project's [Travis-CI tests](https://travis-ci.org/mottosso/Qt.py). This makes it possible for us to better track and document caveats as well as their workarounds.

We encourage everyone who is using Qt.py to report any such caveats along with a workaround (if available) so that they can be incorporated into the `CAVEATS.md` and in our tests.


## Contributing to the project

If you wish to contribute to the project (usually, this means contributing to `Qt.py` or `CAVEATS.md` as well as `tests.py`) you can fork the repository. Then you make the changes in your fork and lastly create a pull request.

We'd kindly like you to provide tests for your pull request. When a pull request is being made, [Travis-CI will automatically attempt to run the tests](https://travis-ci.org/mottosso/Qt.py) specified in the [`tests.py` file](https://github.com/mottosso/Qt.py/blob/master/tests.py). Each test is being run isolated from the other tests. If you don't want to enable Travis-CI for your own fork, you can use a Docker container to run the tests. I'd recommend this approach as it's much faster than waiting for Travis. Read more on this in the [README's developer section](https://github.com/mottosso/Qt.py#developer-guide).

And please don't be offended if your pull request isn't accepted straight away. We sometimes like to discuss pull requests to make sure they will be valuable for a broad audience of users. Sometimes we hit a fork in the road and need to decide on where Qt.py is actually heading. This sometimes takes some amount of ping-pong'ing back and forth. :) Never forget we are incredibly grateful that you've taken your time to make a pull request which will make Qt.py better!

If you have any questions, don't hesitate to [open an issue](https://github.com/mottosso/Qt.py/issues). We'd love to hear what you have to say about Qt.py.


## Closing comments on using Qt.py

#### The future of PySide2

[This](https://wiki.qt.io/PySide2) is the new home of PySide2, since [The Qt Company](https://www.qt.io) took over the development. Keep an eye out for its current status and announcements there. Right now I can see OpenGL is not supported in PySide2.

PySide2 development recently took a hard left turn and announced they will not support Python 2 (although not officially, I was told). Hopefully, as long as an appropriate compiler is used then PySide2 is compatible with Python 2. Maya 2017 which was released today comes with Python 2 and PySide2, so that previous statement seems true. From the [now obsolete PySide2 wiki at Github](https://github.com/PySide/pyside2/wiki):

> **Why there is no PySide2 for Python 2?**  
> Because Python 2 extensions like PySide need to be compiled with ancient version of MS Visual C++ 9 and that means that all linked libs including Qt need to be compiled with this version. But Qt5, the library that PySide2 wraps, dropped support for MS VC++ 9, and code is unlikely to compile for it anymore. The only solution to fix this, is to help with development and funding of https://mingwpy.github.io/

Because of all of this uncertainty around PySide2 ...and also because PyQt5 offers a wheel for Python 3 which can be installed with pip, I'm personally setting out to use Python 3 and PyQt5 wherever I can. Then I just piggy-back on the existing PySide2 in e.g. Autodesk Maya 2017 and PySide in previous versions of Maya.


#### QtQuick and PyQt4

Even though [QtQuick](https://www.qt.io/qt-quick/) is available in PyQt4, I've been recommended to skip this and only use QtQuick in PyQt5, where it apparently works really well.