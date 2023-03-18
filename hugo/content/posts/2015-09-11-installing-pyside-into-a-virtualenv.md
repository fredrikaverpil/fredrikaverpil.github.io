---
ShowToc: false
TocOpen: false
date: 2015-09-11 02:00:12+02:00
draft: false
tags:
- linux
- bash
- pyside
title: Installing PySide into a virtualenv
---

I’ve been struggling to install a portable/relocatable virtualenv with PySide 1.2.2 for Python 2.7. On Windows, this works out of the box but it’s more difficult on Linux and OS X, although I came up with [a patch for OS X](https://github.com/PySide/PySide/issues/129#issuecomment-145138706). This guide will not go into detail on portability/relocatability and will merely touch upon how to get started with PySide in a virtualenv. It looks like PySide 1.2.3 will have substantial changes which will allow for easier portability/relocatability and I will make a post on that as soon as it is generally available.



### Ubuntu 14.04 (trusty) prerequisites

```
sudo apt-get install build-essential cmake libqt4-dev libxml2-dev libxslt1-dev python-dev qtmobility-dev python-pip
```


### CentOS 7 prerequisites

```
sudo yum install epel-release
sudo yum install cmake qt-devel qt-webkit-devel libxml2-devel libxslt-devel python-devel rpmdevtools gcc gcc-c++ make python-pip
sudo ln -s /usr/bin/qmake-qt4 /usr/bin/qmake
```


### Mac OS X 10.10 (Yosemite) prerequisites

According to the [OS X PySide Buildscripts instructions](https://github.com/PySide/BuildScripts/blob/master/dependencies.osx.sh), you need cmake and libxml2 and Qt. If you don’t already have those installed, the easiest way to get them installed is via [brew](http://brew.sh):


```
ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
brew install cmake libxml2 qt4
```

The [OS X PySide Buildscripts README](https://github.com/PySide/BuildScripts/blob/master/README) also says you need to install [Xcode](https://developer.apple.com/tools/xcode/) manually. And then, finally, install pip:

```
curl -O https://bootstrap.pypa.io/get-pip.py
sudo python get-pip.py
```


### Windows 7 & 10 prerequisites

Download the [get-pip.py file](https://bootstrap.pypa.io/get-pip.py) and then execute it to install pip:

```bat
python get-pip.py
```

### Installing PySide into a virtualenv using pip

Install, create and activate a virtualenv called “myVirtualEnv” as well as install PySide using pip:

##### Linux / OS X

```
sudo pip install virtualenv
virtualenv myVirtualEnv
source myVirtualEnv/bin/activate
sudo myVirtualEnv/bin/pip install PySide
```

Please note, this could take several minutes on Linux.

##### Windows

```bat
pip install virtualenv
virtualenv myVirtualEnv
myVirtualEnv/Scripts/activate
myVirtualEnv/Scripts/pip install PySide
```

### QUICK TEST TO SEE THAT PYSIDE INSTALLED PROPERLY

Open the virtualenv’s Python command prompt:

```
myVirtualEnv/bin/python #Linux/OS X
myVirtualEnv/Scripts/python.exe #Windows
```

In the Python prompt, attempt to import QtGui from PySide:

```python
Python 2.7.6 (default, Jun 22 2015, 17:58:13)
[GCC 4.8.2] on linux2
Type "help", "copyright", "credits" or "license" for more information.
>>> from PySide import QtGui
>>>
```

If you don’t see any errors, you should be all good.

To leave the virtualenv:

    deactivate

### Removing symlinks from virtualenv on OS X

If you wish to be able to move your virtualenv around, you may also want to make sure the symlinks created on OS X are made into actual files.

Check for any symlinks:

    find myVirtualenv -type l -ls

So, if you get a list of files back, you may to “convert” those symlinks into actual files. The easiest way to do this is to simply make a copy of the virtualenv folder, and add some options which will copy the actual files rather than maintain the symlinks:

    cp -vRLp myVirtualenv myVirtualenvNoSymlinks

___

If you’ve got input on improvements or on how to achieve this on other platforms/distros, please let me know via the comments below!