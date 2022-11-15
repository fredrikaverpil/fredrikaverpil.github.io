---
title: How to install PySide into OS X El Capitan
tags: [python]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

created: 2016-04-20T03:00:12+02:00
updated: 2022-11-15T17:29:41+01:00
---


## OS X El Capitan

brew install python  # installs in /usr/local/Cellar/python/2.7.11
brew install qt  # installs in /usr/local/Cellar/qt/4.8.7_2

brew install python cmake qt

Do we need these?
brew install openssl

Upgrade pip and wheel...

$ which python
/usr/local/bin/python  # good, it's using brew's python

$ brew list qt

export MACOSX_DEPLOYMENT_TARGET=10.11  # ?????????????

Download tar:
https://github.com/PySide/PySide/releases
tar -xvzf PySide-1.2.4.tar.gz
cd PySide-1.2.4
python2.7 setup.py bdist_wheel

OR - the very latest and greatest:
git clone --recursive https://github.com/PySide/pyside-setup.git pyside-setup
python setup.py bdist_wheel --ignore-git





## Windows

Windows SDK is required; run build command from Windows SDK prompt:

python setup.py bdist_wheel --ignore-git --qmake=c:\Qt\4.8.4_x64\bin\qmake.exe --openssl=c:\OpenSSL-Win64\bin --cmake="C:\Program Files (x86)\CMake\bin\cmake.exe"
