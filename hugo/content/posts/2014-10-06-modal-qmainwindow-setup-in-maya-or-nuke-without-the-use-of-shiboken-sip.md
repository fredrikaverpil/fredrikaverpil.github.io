---
ShowToc: false
TocOpen: false
date: 2014-10-06 02:00:12+02:00
draft: false
tags:
- python
- pyside
title: Modal QMainWindow setup in Maya (or Nuke) without Shiboken/SIP
---

An easy way to get going with PySide in Maya (or Nuke for that matter) without the hassle of dealing with the shiboken/sip layer.



Just a word of caution: The QtCore.QApplication.activeWindow() returns the currently active window (the one in focus). It could be wise to ask the API of e.g. Maya for the main Maya window rather than using this approach, just to be 100% sure the Maya window will be returned by the function in all situations. But for any other window that you may want to open from there on, you could safely use this approach, as you would know your own windows.


```python
from PySide import QtGui, QtCore

class MyApp(QtGui.QMainWindow):
    def __init__(self, parent=None):
        super(MyApp, self).__init__(parent)

        self.setWindowTitle('Hello')
        self.centralWidget = QtGui.QWidget(self)
        self.setCentralWidget(self.centralWidget)
        self.mainLayout = QtGui.QVBoxLayout(self.centralWidget)
        self.pushButton = QtGui.QPushButton('Hello')
        self.mainLayout.addWidget(self.pushButton)

my_app = MyApp(parent=QtGui.QApplication.activeWindow())
my_app.setWindowModality(QtCore.Qt.WindowModal)
my_app.show()
```

And similarly, it also works with QDialog, without the need of setting the window modality specifically:

```python
from PySide import QtGui, QtCore

class MyApp(QtGui.QDialog):
    def __init__(self, parent=None):
        super(MyApp, self).__init__(parent)

        self.setWindowTitle('Hello')
        self.mainLayout = QtGui.QVBoxLayout()
        self.setLayout(self.mainLayout)
        self.pushButton = QtGui.QPushButton('Hello')
        self.mainLayout.addWidget(self.pushButton)

my_app = MyApp(parent=QtGui.QApplication.activeWindow())
my_app.exec_()
```

Unfortunately, it does not seem like you can use this approach with a QWidget.

On a similar note, I thought it could be useful to add a link to this Stackoverflow question: [PyQt4: How can I toggle the “stay on top behavior?"](http://stackoverflow.com/questions/4850584/pyqt4-how-can-i-toggle-the-stay-on-top-behavior)