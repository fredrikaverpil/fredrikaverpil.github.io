---
date: 2016-07-25
tags:
- python
- maya
- qt.py
- pyside
- pyqt
---

# Dealing with Maya 2017 and PySide2

![](/static/maya2017/no_pyside.png)

[Maya 2017](http://www.autodesk.com/products/maya/overview) was released today and with it comes a big change; PySide (and PyQt4) no longer works with Maya.

This guide explains how to deal with that and make your Python and PySide/PyQt scripts compatible with Maya 2017 as well as older Maya versions.

This guide is also applicable to e.g. [Nuke](https://www.thefoundry.co.uk/products/nuke/) or any other Python-enabled DCC app which uses Qt.

<!-- more -->

## Background (Qt4 vs Qt5)

Starting with Maya 2011, Maya's user interface is built using the [Qt](http://www.qt.io) toolkit. Up until and including Maya 2016, Qt of version 4 ("Qt4") which was [released in 2005](https://en.wikipedia.org/wiki/List_of_Qt_releases#Qt_4) has been used. Today, roughly 11 years after the Qt4 release, and starting with Maya 2017, Qt was upgraded to version 5 ("Qt5") to enable a host of enhancements to Maya. Unfortunately, this also breaks backwards compatibility to Qt4.

When you use PySide (or PyQt4), you're actually using the Python bindings for Qt4. So naturally, when Qt4 no longer exists in Maya, PySide doesn't work anymore and was therefore removed in Maya 2017. Enter [PySide2](https://wiki.qt.io/PySide2).

PySide2 are the new Python bindings for Qt5 and is now bundled with Maya 2017. With this comes a bunch of changes, so unfortunately, you can't just "`import PySide2`" instead of "`import PySide`" and then expect your old scripts to work.

You also cannot use PyQt4 with Maya 2017, as they are also Python bindings for Qt4. Instead you should use use [PyQt5](https://www.riverbankcomputing.com/software/pyqt/download5) which are the new Python bindings for Qt5.


## Differences between Qt4 and Qt5

I'm writing "Qt4 and Qt5" instead of "PySide and PySide2" or "PyQt4 and PyQt5" since we're talking about similar Python bindings which in the end uses Qt.

The biggest change between the two, which you'll notice since it breaks your old scripts, is that a bunch of stuff was taken out from `QtGui` and instead placed in the new `QtWidgets` module. Perhaps the Qt Company thought `QtGui` got too bloated?

Like you probably guessed, all widgets are now created from `QtWidgets`. Example with PySide2:

```python
from PySide2 import QtWidgets

my_label = QtWidgets.QLabel('Hello')
```

To figure out if a class moved into e.g. the `QtWidgets` module (such as `QLabel` in the example above), check the documentation on that particular class. Strangely enough, there's no official documentation for PySide2. The best suggestion I have is you should check with the [Qt5 documentation](http://doc.qt.io/qt-5/) first and if required, cross-reference with the [PyQt5 documentation](http://pyqt.sourceforge.net/Docs/PyQt5/).

There is of course a smörgåsbord of other changes in Qt5 as well, but this moving of classes I would say is probably the predominant one you need to address in your current Maya scripts.

I recommend skimming through the official ["What's new in Qt5"](http://doc.qt.io/qt-5/qt5-intro.html) to get a good understanding on why Autodesk (and soon others to come) think it's time to leave Qt4 and see where Qt5 is heading.


## Convert PyQt4 code into PyQt5 code

I haven't used this myself, but I figured it was worth mentioning. There's a magical conversion script you can try your luck with, in case you're looking for a fast conversion from PyQt4 to PyQt5: [pyqt4topyqt5](https://github.com/rferrazz/pyqt4topyqt5)

Quoting Mark from Disney Animation ([source](https://groups.google.com/d/msgid/vfx-platform-discuss/be711b3f-5417-4449-8cbe-aeebb71f793b%40googlegroups.com?utm_medium=email&utm_source=footer)):

> One of the key components in our conversion process was a PyQt4 to PyQt5 Conversion Script from a github project known as pyqt4topyqt5.
>
> We have greatly enhanced this script and our changes are now available on github. Thanks to Riccardo Ferrazzo for starting this project. It has been a great help to us and I hope that it will be a great help to others as well.


## Backwards compatibility (Qt.py)

This will expand on how to make the same Python script work in Maya 2017 and previous versions, using PySide, PySide2, PyQt4 or PyQt5 (whatever is available) by using [Qt.py](https://github.com/mottosso/Qt.py).

**Full disclosure:** Qt.py was created and is maintained by Marcus Ottosson and myself. At the time of writing this, Qt.py is being used at [Disney Animation](http://www.disneyanimation.com), [Framestore](https://www.framestore.com), [Industriromantik](http://www.industriromantik.se) (where I work) and [Weta Digital](https://www.wetafx.co.nz).

Qt.py can be used in place of the regular Python bindings. When using Qt.py, you should write your code as you would write PySide2 code. This allows for the following example code to work in Maya 2017 **as well as prior Maya versions**:

```python
from Qt import QtWidgets

info_dialog = QtWidgets.QMessageBox()
info_dialog.setWindowTitle('Hello world')
info_dialog.setText('Hello!')
info_dialog.exec_()
```

Please note that in the example above, `QtWidgets` is imported from `Qt` (which is Qt.py).

In short, Qt.py replaces itself with whatever Python binding is available (or preferred) in the Python environment your script runs in. However, Qt.py will first take care of remappings of calls so that the Qt5 style syntax seen in the example will work even if you run this code in an environment with a Python binding for Qt4. This is how this code also works in e.g. Maya 2016 with PySide.

There are also a number of enhancements with Qt.py such as enabling the sip v2 API by default when PyQt4 is being used by Qt.py.

The only requirement for Qt.py to work is that your Python environment has a valid Qt Python binding installed. In Maya's case, you could rely on PySide2 with Maya 2017 and PySide with previous Maya versions, as these are pre-installed with Maya. The example code above also works just fine in Nuke, since PySide is available there, pre-installed as well. If you're running Python 2 or Python 3 outside of any DCC application, just make sure you've got PySide, PySide2, PyQt4 or PyQt5 installed.

There are however limitations. You can (of course) not utilize Qt5-specific features unless Qt.py is actually using Python bindings for Qt5 (PySide2 or PyQt5). There are also binding-specifics, which require a certain binding.

Since you can check which binding is being used...

```python
>>> import Qt
>>> print Qt.__binding__
PySide2
```

...you can tell your script to explicitly do something depending on the used binding. An example checking if a Qt5 binding is used or if a Qt4 binding is used:

```python
from Qt import __binding__

if __binding__ in ('PySide2', 'PyQt5'):
    print('Qt5 binding available')
elif __binding__ in ('PySide', 'PyQt4'):
    print('Qt4 binding available.')
else:
    print('No Qt binding available.')
```

There are also a set of rules which would help you during development. Please have a look at the project's [README](https://github.com/mottosso/Qt.py#rules) for more details.

In case you're interested in reading more on Qt.py, check out my guide [here](2016-07-25-developing-with-qt-py.md).