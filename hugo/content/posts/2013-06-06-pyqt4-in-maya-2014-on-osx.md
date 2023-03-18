---
ShowToc: false
TocOpen: false
date: 2013-06-06 02:00:12+02:00
draft: false
tags:
- maya
- python
title: PyQt4 in Maya 2014 on OS X
---

There are some official Autodesk [instructions](http://around-the-corner.typepad.com/adn/2013/04/building-sip-and-pyqt-for-maya-2014.html) floating around on how to get PyQt4 support in of Maya 2014, but I found them a bit unclear and hard to follow. I also found an error (although minor)… so I’ve typed down my own little step-by-step reminder below, based off that.



However, if you are in a hurry and just want the PyQt4 and SIP site-packages contents you can just go ahead and download the partial result of the long procedure below: [site-packages.zip](https://raw.githubusercontent.com/fredrikaverpil/maya/master/PyQt4/maya2014_osx/site-packages.zip) – PyQt4 for Maya 2014 on OS X 10.8 Mountain Lion (PyQt 4.10 and SIP 4.14.5). Copy the contents into the following folder and you should have PyQt4 working in Maya 2014:

    /Applications/Autodesk/maya2014/Maya.app/Contents/Frameworks/Python.framework/Versions/2.7/lib/python2.7/site-packages/

### Step by step instructions on building PyQt4 for Maya 2014 on OS X

##### Get Xcode's command line tools

Install Xcode via the App Store. Launch it, go into the preferences and download/install the command line tools.

##### Download PyQt/SIP source and Autodesk's build scripts

Download [PyQt-mac-gpl-4.10.tar.gz](http://sourceforge.net/projects/pyqt/files/PyQt4/PyQt-4.10/PyQt-mac-gpl-4.10.tar.gz/download) and [sip-4.14.5.tar.gz](http://sourceforge.net/projects/pyqt/files/sip/sip-4.14.5/sip-4.14.5.tar.gz/download). Unpack these into a local folder on your desktop so that there are now two folders in there: “PyQt-mac-gpl-4.10” and “sip-4.14.5”. Make sure these are the exact folder names or the build scripts which you are about to download (below) are not going to work.

Copy these scripts into your local desktop folder as well: [build_and_install_sip.sh](https://raw.githubusercontent.com/fredrikaverpil/maya-scripts/master/PyQt4/maya2014_osx/build_and_install_SIP.sh) and [build_and_install_PyQt4.sh](https://raw.githubusercontent.com/fredrikaverpil/maya-scripts/master/PyQt4/maya2014_osx/build_and_install_PyQt4.sh) but don’t execute any of them just yet.

##### Prepare the bundled and Autodesk-customized Qt 4.8.2

Extract the contents of `/Applications/Autodesk/maya2014/devkit/include/qt-4.8.2-include.tar.gz` into `/Applications/Autodesk/maya2014/devkit/include/Qt`.

Somehow get the archive called qt-4.8.2-64-mkspecs.tar.gz from the Windows or Linux installation directory of Maya 2014 (it doesn’t come with Maya on OS X). Then extract its contents into `/Applications/Autodesk/maya2014/Maya.app/Contents/mkspecs`.

Now, edit the qconfig.pri file in `/Applications/Autodesk/maya2014/Maya.app/Contents/mkspecs` to reflect the following:

```bash
#configuration
CONFIG += release def_files_disabled exceptions no_mocdepend stl
x86_64 qt #qt_framework
QT_ARCH = macosx
QT_EDITION = OpenSource
QT_CONFIG += minimal-config small-config medium-config largeconfig full-config no-pkg-config dwarf2 phonon phonon-backend
accessibility opengl reduce_exports ipv6 getaddrinfo ipv6ifname
getifaddrs png no-freetype system-zlib nis cups iconv openssl
corewlan concurrent xmlpatterns multimedia audio-backend svg
script scripttools declarative release x86_64 qt #qt_framework
#versioning
QT_VERSION = 4.8.2
QT_MAJOR_VERSION = 4
QT_MINOR_VERSION = 8
QT_PATCH_VERSION = 2
#namespaces
QT_LIBINFIX =
QT_NAMESPACE =
QT_NAMESPACE_MAC_CRC =
QT_GCC_MAJOR_VERSION = 4
QT_GCC_MINOR_VERSION = 2
```

Last but not least, copy `/Applications/Autodesk/maya2014/Maya.app/Contents/Resources/qt.conf` into `/Applications/Autodesk/maya2014/Maya.app/Contents/bin` and edit it accordingly:

```bash
[Paths]
Prefix=
Libraries=../MacOS
Binaries=../bin
Headers=../../../devkit/include/Qt
Data=..
Plugins=../qt-plugins
Translations=../qt-translations
```

##### Put the fake dylib files in place

Create what Autodesk refers to as “fake” dylib files. I’ve provided the copy commands below:

```bash
cp /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/QtCore /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/libQtCore.dylib
cp /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/QtDeclarative /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/libQtDeclarative.dylib
cp /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/QtDesigner /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/libQtDesigner.dylib
cp /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/QtDesignerComponents /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/libQtDesignerComponents.dylib
cp /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/QtGui /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/libQtGui.dylib
cp /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/QtHelp /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/libQtHelp.dylib
cp /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/QtMultimedia /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/libQtMultimedia.dylib
cp /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/QtNetwork /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/libQtNetwork.dylib
cp /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/QtOpenGL /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/libQtOpenGL.dylib
cp /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/QtScript /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/libQtScript.dylib
cp /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/QtScriptTools /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/libQtScriptTools.dylib
cp /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/QtSql /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/libQtSql.dylib
cp /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/QtSvg /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/libQtSvg.dylib
cp /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/QtWebKit /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/libQtWebKit.dylib
cp /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/QtXml /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/libQtXml.dylib
cp /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/QtXmlPatterns /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/libQtXmlPatterns.dylib
cp /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/phonon /Applications/Autodesk/maya2014/Maya.app/Contents/MacOS/libphonon.dylib
```

##### Last step: build and install SIP and PyQt4 for Maya

Open up a Terminal window and make your way into the local folder where you put the bash scripts. Then execute the first script to build and install SIP (you may have to enter your password at some point):

    sh build_and_install_sip.sh

Run the second script to build and install PyQt4. This will take a while, but after a couple of minutes you may be asked to enter your password.

    sh build_and_install_PyQt4.sh

You’re done. You should now be able to use PyQt4 in Maya!