---
layout: post
title: Compiling PySide2 from source
tags: [pyside, python]
---

![]({{ site.baseurl }}/blog/assets/pyside2/pyside2_compile.png)

Here's how to compile PySide2 on Linux, Mac OS X and Windows.

[Maya 2017](http://www.autodesk.com/products/maya) shipped with [PySide2](https://wiki.qt.io/PySide2) (2.0.0~alpha0) built against [Qt](https://www.qt.io/developers/) 5.6.1, so that's the version of Qt I'm using in this guide as well as 64-bit [Python](https://www.python.org).

**UPDATE:** Before reading on, this blog post is now _very_ outdated. You may want to skip this alltogether and go directly to my pre-built standalone Python wheels and the most up to date build recipes there:
* [fredrikaverpil/pyside2-linux](https://github.com/fredrikaverpil/pyside2-linux)
* [fredrikaverpil/pyside2-macos](https://github.com/fredrikaverpil/pyside2-macos)
* [fredrikaverpil/pyside2-windows](https://github.com/fredrikaverpil/pyside2-windows)

Also, official standalone wheels have been [released](http://lists.qt-project.org/pipermail/pyside/2018-March/002537.html). PySide2 can now be installed from pypi.org: `pip install PySide2`!


<!--more-->

## Contents

* Dockerfiles
* Prerequisites
  * OS X
  * Ubuntu 14.04 Linux
  * Ubuntu 16.04 Linux
  * CentOS 7 Linux
  * Windows 10
* Clone the repository
* Build the PySide2 wheel
  * OS X
  * Ubuntu 16.04 Linux
  * CentOS 7 Linux
  * Windows 10
* Install the wheel
* Notes on precompiled wheels
* Docker-built wheels via Travis-CI
* Closing comments


## Prerequisites

PySide2 (or possibly Qt5?) requires Cmake >=3.0 to work. You can check your version with `cmake --version`.


### OS X

You'll need the Xcode commandline tools:

```
xcode-select --install
sudo xcodebuild -license
```

Then the following packages, easily installed via [brew](https://brew.sh):

```bash
brew install qt5 cmake libxslt libxml2
```

This installed Qt 5.6.1-1 and cmake 3.5.2 on my system.

### Ubuntu 14.04 Linux

Currently, I'm having issues with not being able to install `libqt5quickwidgets5` via apt-get, and because of that I recommend resorting to the following:

```bash
apt-get install python-pyside2  # for Python 2.7.x
apt-get install python3-pyside2  # for Python 3.5.x
```


### Ubuntu 16.04 Linux

Please note, not all Qt5 modules are being used when building the wheel, see this issue report:
[https://bugreports.qt.io/browse/PYSIDE-343](https://bugreports.qt.io/browse/PYSIDE-343)

In case you wish to build PySide2 in a Docker container, you can start by doing something like this:

```bash
# Map the current folder to the container's /pyside, so you can
# transfer built wheels onto your main machine
docker run -ti -v $(pwd):/pyside ubuntu:16.04 bash
```

You may then want to perform an `apt-get update`, especially if you're in a Docker container.

Make sure you've got pip (and its wheel package):

```bash
apt-get install python-pip  # or python3-pip (for Python 3.x)
```

Then proceed with installing some dependencies:

```bash
apt-get install build-essential git cmake libxml2 libxslt1.1 libxml2-dev libxslt1-dev python-dev qt5-default qtbase5-dev
```

This installed Qt 5.5.1 and cmake 3.5.1 on my system. If you really wish to use Qt 5.6.1, you currently need to build Qt from source instead of installing it via apt-get. This, however, is for another blog post as it is *extremely* time consuming. You can verify which version of Qt you have by executing `qmake --version`.


Additionally, you'll need to install Qt5's libraries:

```bash
apt-get install qttools5-dev-tools libqt5clucene5 libqt5concurrent5 libqt5core5a libqt5dbus5 libqt5designer5 libqt5designercomponents5 libqt5feedback5 libqt5gui5 libqt5help5 libqt5multimedia5 libqt5network5 libqt5opengl5 libqt5opengl5-dev libqt5organizer5 libqt5positioning5 libqt5printsupport5 libqt5qml5 libqt5quick5 libqt5quickwidgets5 libqt5script5 libqt5scripttools5 libqt5sql5 libqt5sql5-sqlite libqt5svg5 libqt5test5 libqt5webkit5 libqt5widgets5 libqt5xml5 libqt5xmlpatterns5 libqt5xmlpatterns5-dev
```


### CentOS 7 Linux

Please note: the guide for CentOS 7 Linux is incomplete and will generate an error mid-build. I'll update this section as soon as I come up with a solution to this problem.

```bash
yum install epel-release
yum groupinstall "Development Tools"
yum install qt5-qtbase-devel
```

This installed Qt 5.6.1 on my system:

```
/usr/lib64/qt5/bin/qmake-qt5 --version
QMake version 3.0
Using Qt version 5.6.1 in /usr/lib64
```

Now we'll also need additional Qt5 dependencies. This command below will probably install a lot of packages which you actually won't need (note the asterisk in the command below). If you figure out exactly which packages are needed, please leave a comment on that!

```bash
yum install qt5-* libxslt libxml2 libxml2-devel libxslt-devel
```

Next, we'll install Cmake 3.0 (yum only provides 2.x at this time). Please note this will build the Cmake executable in the same directory as the source.

```bash
yum install cmake3
```

Install Python 2.x or 3.5:

```bash
yum install python-libs python-pip python-devel
pip install wheel
```

```bash
yum install python35u python35u-libs python35u-devel python35u-pip
pip3.5 install wheel
```



### Windows 10


Unfortunately, I don't know how to compile PySide2 using Qt 5.6.1 for Python 2.7.x in Windows. The reason is that the [official Python 2.7 build](https://python.org) is compiled using the now ancient Microsoft C++ Visual Studio 2008 (v9.0). This means Qt itself must also be compiled using this version of MSVC in order to avoid severe issues (read more [here](http://siomsystems.com/mixing-visual-studio-versions/)). Unfortunately, there's no such readily available version of Qt 5.6.1. What it all boils down to is you want to have used the same MSVC version for Python, Qt and PySide2.

To work around this, you can [compile Python yourself](http://p-nand-q.com/python/building-python-27-with-visual_studio.html), matching the MSVC used to compile Qt. This is what Autodesk is doing by the way, in order to support PySide2 in Maya 2017 without mixing different MSVC versions. Or you could attempt to compile Qt using MSVC 2008. In either case, this will give you a headache when you realize any Python package (which is not purely written in Python) must also be compiled using whatever version of MSVC you have opted for. And perhaps because of this, you'll now start to wonder what packages you're currently importing when in Maya...

Another solution worth investigating, which seems better to me (but I may be mistaken here), is to compile PySide2 using [mingw32](http://www.mingw.org) and/or [mingwpy](https://mingwpy.github.io/) rather than MSVC. I very much appreciate suggestions in the comments further down below on this.

But so, at least for now, this guide is for Python 3.5.x only.

You'll need OpenSSL, Cmake, Microsoft Visual C++ Studio 2015 and, of course, Qt5.

I think the nicest approach would be to install these softwares via [Chocolatey](https://chocolatey.org) but I haven't managed to get that to work flawlessly, so I'm doing this manually instead.

Download and install the 64-bit versions of:

* [OpenSSL](https://sourceforge.net/projects/openssl)
* [Cmake](https://cmake.org/download)
* [Microsoft Visual C++ Studio 2015](https://www.visualstudio.com) - Note: during the install, you must check the programming language option for C++, or the `cl` command won't be available. This can also be performed if you run the installer again after having already installed MSVC2015 (choose to "Modify" your installation).
* Qt 5.6.1-1 from the [Qt archives](https://download.qt.io/archive/qt/), I'm using [qt-opensource-windows-x86-msvc2015_64-5.6.1-1.exe](https://download.qt.io/archive/qt/5.6/5.6.1-1/qt-opensource-windows-x86-msvc2015_64-5.6.1-1.exe).

You might also need the Windows SDK (I'm not sure since I had it installed already when writing this). You can download it for Windows 10 from [here](https://developer.microsoft.com/en-us/windows/downloads/windows-10-sdk). I have a feeling this is installed by MSVC2015, but I'm not sure. If you know what the deal is here, please let me know in the comments further down below!


## Clone the repository

The old repository is located in Github. I'm not sure why they still keep it there as it is no longer maintained. The active repositories are hosted by the Qt Company and resides [here](https://codereview.qt-project.org/#/admin/projects/?filter=pyside).

Clone the `pyside-setup` repository and have it also pull down its gitmodules:

```bash
# Edited on 2017-07-20; added branch specific to Qt 5.6
git clone --recursive --branch 5.6 https://code.qt.io/pyside/pyside-setup.git
```

Autodesk has also provided the version of PySide2 they used in Maya 2017 on their [open source distributions page](http://www.autodesk.com/company/legal-notices-trademarks/open-source-distribution). If you prefer that, you can download the PySide2.zip file from there and use that instead.

## Build the PySide2 wheel

The exact paths given in the arguments may not be identical on your system so verify those prior to compiling.

### Notes on Linux

Please note, in setup.py, the `--standalone` argument is mentioned, but I don't see this working when including that in the build command:

> On Linux you can use option --standalone, to embed Qt libraries to PySide2 distribution

Also please note, I'm not using the `--openssl` argument since I actually wasn't able to figure out where the OpenSSL bin directory resides in Ubuntu/CentOS. At least, searching using `find / -name openssl` didn't reveal this and the wheel works anyways. If you know where this resides or how to make it available, I'd be grateful if you would like to share this with me in the comments further down below.


### OS X

This command worked fine for me using Python 2.7.11 and Python 3.5.1. Remember to have pip installed with the `wheel` package or you'll get an error about `bdist_wheel`.

```
python setup.py bdist_wheel --ignore-git --qmake=/usr/local/Cellar/qt5/5.6.1-1/bin/qmake --cmake=/usr/local/bin/cmake --openssl=/usr/local/Cellar/openssl/1.0.2h_1/bin
```

### Ubuntu 16.04 Linux


This command worked fine for me using Python 2.7.12 and Python 3.5.2. Remember to have pip installed with the `wheel` package or you'll get an error about `bdist_wheel`.

```
python setup.py bdist_wheel --ignore-git --qmake=/usr/lib/x86_64-linux-gnu/qt5/bin/qmake --cmake=/usr/bin/cmake
```


### CentOS 7 Linux

Due to [this issue](https://bugreports.qt.io/browse/PYSIDE-342), you need to manually add the following to `/pyside-setup/sources/pyside2/pyslidelib/CMakeLists.txt`, after the first line (which says `project(libpyside)`):

```
#HACK: CMake with broken Qt5Qml_PRIVATE_INCLUDE_DIRS, Qt5Quick_PRIVATE_INCLUDE_DIRS
  if(${Qt5Qml_FOUND})
    if(NOT "${Qt5Qml_PRIVATE_INCLUDE_DIRS}" MATCHES "/QtQml/")
      string(REPLACE "/QtCore" "/QtQml" replaceme "${Qt5Core_PRIVATE_INCLUDE_DIRS}")
      list(APPEND Qt5Qml_PRIVATE_INCLUDE_DIRS ${replaceme})
      list(REMOVE_DUPLICATES Qt5Qml_PRIVATE_INCLUDE_DIRS)
    endif()
  endif()
  if(${Qt5Quick_FOUND})
    if(NOT "${Qt5Quick_PRIVATE_INCLUDE_DIRS}" MATCHES "/QtQuick/")
      string(REPLACE "/QtCore" "/QtQuick" replaceme "${Qt5Core_PRIVATE_INCLUDE_DIRS}")
      list(APPEND Qt5Quick_PRIVATE_INCLUDE_DIRS ${Qt5Qml_PRIVATE_INCLUDE_DIRS})
      list(APPEND Qt5Quick_PRIVATE_INCLUDE_DIRS ${replaceme})
      list(REMOVE_DUPLICATES Qt5Quick_PRIVATE_INCLUDE_DIRS)
    endif()
  endif()
```

Then you can build the PySide2 wheel:

```
python setup.py bdist_wheel --ignore-git --qmake=/usr/lib64/qt5/bin/qmake-qt5 --cmake=/usr/bin/cmake3
```

If you wish to use Python 3.5, execute `python3.5` instead of `python`.


### Windows 10

Open the prompt called "VS2015 x64 Native Tools Command Prompt" and run the build command.

```
python setup.py bdist_wheel --ignore-git --qmake=c:\Qt\Qt5.6.1\5.6\msvc2015\bin\qmake.exe --openssl="C:\openssl-1.0.2d-fips-2.0.10\bin" --cmake="C:\Program Files\CMake\bin\cmake.exe"
```

## Install the wheel

A wheel was hopefully built in the dist folder. So just `cd dist` and `pip install` away!

## Notes on pre-compiled wheels

Unfortunately, and like with PySide, these wheels are not "portable" and won't install on systems which doesn't already have the specific Qt5 version installed used during compilation. This, I believe, is because PySide2 links dynamically (instead of statically) against the Qt5 installation. Hopefully, this is something The Qt Company will address via official PySide2 wheels, as Riverbank Software is now providing a fully portable PyQt5 wheel for Python 3 which is absolutely awesome. In case you think we could produce such wheels ourselves, I'd appreciate if you'd leave a comment on this further down below!

Personally, this makes me want to develop for PyQt5 wherever I can since this is all it takes to get going if you're using Python 3: `pip install PyQt5`.
If you also feel this way, have a look at the [Qt.py](https://github.com/mottosso/Qt.py) project which will enable you to write code which will work in both PySide2 and PyQt5 (and PySide/PyQt4).


## Docker-built wheels via Travis-CI

I've set up a repository where I test the wheel building with Docker:
[https://github.com/fredrikaverpil/pyside2-wheels](https://github.com/fredrikaverpil/pyside2-wheels)


## Closing comments

As of writing this, I understand that PySide2 is not seen as having been made generally available just yet, and could perhaps be seen as being in an alpha stage. They did announce that Python 2 support was dropped [here](https://github.com/PySide/pyside2/wiki) (although not officially, I was told):

> Why there is no PySide2 for Python 2? Because Python 2 extensions like PySide need to be compiled with ancient version of MS Visual C++ 9 and that means that all linked libs including Qt need to be compiled with this version. But Qt5, the library that PySide2 wraps, dropped support for MS VC++ 9, and code is unlikely to compile for it anymore. The only solution to fix this, is to help with development and funding of [https://mingwpy.github.io/](https://mingwpy.github.io/)

This is yet another aging sign of Python 2.x. Personally, I'm developing with Python 3.5 as the interpreter of choice in my IDE and I make sure my code is backwards compatible with Python 2.7, which is actually very easy for the most part. I think that the [VFX Reference Platform](http://www.vfxplatform.com) should plan ahead (starting now) for a move onto Python 3.5 as there is great reluctancy in the industry towards this move. It'll involve a lot of code porting, yes, but with enough headsup time this shouldn't be blocking such a move. In a developer's world, this is how it works. You can't hold on to old code forever.

Since the old PySide Github repository is basically abandoned, I'd suggest you create a user account with The Qt Company and start checking out the repository over at Qt, where apparently the real action is:

* [PySide2 wiki](https://wiki.qt.io/PySide2)
* [PySide2 repositories](https://codereview.qt-project.org/#/admin/projects/?filter=pyside)
* [Bug tracker (issues)](https://bugreports.qt.io/browse/PYSIDE/)
* [Qt language bindings forum](https://forum.qt.io/category/15/language-bindings)
* [PySide/PySide2 Gitter channel](https://gitter.im/PySide/pyside2) or #qt-pyside on irc.freenode.net (replaces the former #pyside channel, after PySide moved into Qt)
