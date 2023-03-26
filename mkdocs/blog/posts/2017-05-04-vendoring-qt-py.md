---
date: 2017-05-04
tags:
- python
- qt.py
- pyside
- pyqt
---

# Vendoring Qt.py

How do you know a user doesn't have the wrong version of [Qt.py](https://github.com/mottosso/Qt.py) when running your application?  
– Simple, you bundle Qt.py with your application. Here's a short write-up on how you could go about doing just that.

<!-- more -->

## What's "vendoring"?

Bundling a third-party application with your own application is sometimes referred to as "vendoring". By vendoring, you explicitly control which version of a dependency is being used with your application. Since there are numerous versions of Qt.py you want to make sure your application is used with the version of Qt.py you had in mind when designing your application.


## Licensing

It's important to mention you cannot bundle *any* application with your own application. You must look into what you're actually legally allowed to do. In the case with Qt.py, it's developed under the MIT license, which means you can do anything you like with it. You can even modify Qt.py, sell your software with it bundled and keep it all closed source. However, speaking as a collaborator on the project, we're always super happy to hear about what you're doing with Qt.py :)

Curious on different types of licensing?  
– Check this out: [https://choosealicense.com/licenses/](https://choosealicense.com/licenses/)

With that out of the way, let's move on...


## Python package folder structure

In this example, I'll just create a simple folder structure where we'll put the Qt.py script in a "vendor" folder and our main app (`myapp.py`) at the top level. The `__init__.py` files makes it possible for Python to import the contents of the folders as modules.

```
mypackage/
  vendor/
    __init__.py
    Qt.py
  __init__.py
  myapp.py
```

The `__init__.py` files can (and should most often) be left empty. Any contents in these files will be executed upon import, which is not in general cases desired. Instead the action will happen inside of `myapp.py`.

In short, this is how you load your vendored version of Qt.py from within `myapp.py`:

```python
from .vendor import Qt
from .vendor.Qt import QtWidgets

print(Qt.__binding__)
# and so on...
```

For a complete and working example, please see this git repo I set up: [https://github.com/fredrikaverpil/Qt.py-vendoring/](https://github.com/fredrikaverpil/Qt.py-vendoring/)


## Installation instructions

You now have a couple of options of where to tell your users to put your package, so that it can be run in Python. In essence, you need to put the package on a path where Python will search for it. You could use an existing path, already known to your Python interpreter, where it will already search automatically. To find out which paths are available, run the following from the Python interpreter:

```python
>>> import site
>>> print(site.getsitepackages())
>>> print(site.getusersitepackages())
```

If you would rather specify a custom path where you'll have the package sitting, you can tell Python to look for the package in a specific path:

```python
>>> import site
>>> site.addsitedir(YOUR_PATH_HERE)  # add path containing myPackage to PYTHONPATH
```

For more info on the `site` package, see [here](https://docs.python.org/2/library/site.html).

For this example, I'm cloning the example [git repository](https://github.com/fredrikaverpil/Qt.py-vendoring/) down onto my local macOS disk drive:

```bash
cd ~/code/repos
git clone https://github.com/fredrikaverpil/Qt.py-vendoring.git
```

## Running your application

Make sure the interpreter you're using has either PySide, PyQt4, PySide2 or PyQt5 installed. Then import `myapp` from `mypackage` and, finally, run the `myapp.run()` function. I'm doing it like this:

```python
>>> import site
>>> site.addsitedir('/Users/fredrik/code/repos/Qt.py-vendoring')
>>> import mypackage
>>> from mypackage import myapp
>>> myapp.run()
```

## Pro tip #1: Semantic versioning

By using semantic versioning (aka "semver") you can check the versions against e.g. a required minimum version which your application requires. One way of doing this is to utilize `setuptool`'s `packaging.version.parse` function:

```python
>>> from packaging import version
>>> version.parse("1.0.0b1") < version.parse("1.0.1")
True
```

## Pro tip #2: Making your package available on PyPi

If you wish to go even further and let users download your package from [PyPi](https://pypi.python.org/pypi) (via [pip](https://packaging.python.org/installing/)), you'll need to add a `setup.py` file (and optionally a `setup.cfg` file if you e.g. wish to generate a universal wheel) and configure it accordingly. Then I'd recommend you create a wheel and finally upload that to PyPi. Read more about this over at the [`setuptools` documentation](https://packaging.python.org/distributing/).

[Here's](https://github.com/mottosso/Qt.py/blob/master/setup.py) the `setup.py` in the Qt.py project as an example.


## Notes

I previously covered vendoring in my [Installing Qt.py (advanced methods)](https://fredrikaverpil.github.io/2016/07/25/installing-qt-py-advanced-methods/) blog post. Check it out for additional info.


## How are you distributing Qt.py?

I'm curious on how you make Qt.py available to users. Let me know in the comments below!