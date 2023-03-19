---
date: 2015-11-25
tags:
- osx
- python
- pyqt
---

# Compiling PyQt5 for Python 2.7 on OS X

<div class="message">
  Guide updated for Qt5 5.6 and PyQt5 5.6.
</div>

This quick guide details compiling sip and PyQt5 on OS X 10.11 (El Capitan) using [Homebrew](http://brew.sh) for Qt5 installtion.

<!-- more -->

In case you don’t have Homebrew installed:

    ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"

Let’s start with installing the latest version of Python 2.x, Qt5, as well as wget using brew:

    brew install python qt5 wget

While that is “brewing”, make sure you have Xcode installed (can be installed via the [Mac App Store](https://itunes.apple.com/en/app/xcode/id497799835?mt=12)). When Xcode is installed, also make sure you have its command-line tools installed and that you have agreed to Apple’s license agreement:

    xcode-select --install
    sudo xcodebuild -license

Then let’s download the [PyQt5 source for Linux and OS X](https://riverbankcomputing.com/software/pyqt/download5) and the prerequisite [SIP source](https://riverbankcomputing.com/software/sip/download).

    wget http://sourceforge.net/projects/pyqt/files/PyQt5/PyQt-5.6/PyQt5_gpl-5.6.tar.gz
    wget http://freefr.dl.sourceforge.net/project/pyqt/sip/sip-4.18/sip-4.18.tar.gz

Double-check that the newly installed Python 2.7.x is being used when just executing python:

    python --version

Untar and compile (also double check the path to your qmake):

    tar -xvf sip-4.18.tar.gz
    cd /sip-4.18
    python configure.py -d /usr/local/lib/python2.7/site-packages/
    make
    make install

    cd..
    tar -xvf PyQt-gpl-5.6.tar.gz
    cd PyQt-gpl-5.6
    python configure.py -d /usr/local/lib/python2.7/site-packages/ --qmake=/usr/local/Cellar/qt5/5.6.0/bin/qmake --sip=/usr/local/bin/sip --sip-incdir=../sip-4.18/siplib
    make
    make install

Please note, you may want want to check out the options in the `configure.py` files prior to configuring/compiling.

You may now import PyQt5 as a module in Python 2.7!

    $ python

    Python 2.7.10 (default, Sep 23 2015, 04:34:21)
    [GCC 4.2.1 Compatible Apple LLVM 7.0.0 (clang-700.0.72)] on darwin
    Type "help", "copyright", "credits" or "license" for more information.
    >>> import PyQt5
    >>>