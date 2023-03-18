---
ShowToc: false
TocOpen: false
date: 2014-06-05 02:00:12+02:00
draft: false
tags:
- python
title: 64-bit Python package building with VS2008
---

This was a real PITA to get working. After having installed Visual C++ 2008 Express, I tried to build both [psutil](https://code.google.com/archive/p/psutil/) and [MySQLdb](http://sourceforge.net/projects/mysql-python/) using setuptools and/or pip. Both resulted in various errors…



For anyone barking up the same tree… here’s the solution that worked for me.

Make sure you have both [pip](https://pypi.python.org/pypi/pip) and [setuptools](https://pypi.python.org/pypi/setuptools) installed to be able to build, and obviously 64-bit Python 2.6 or 2.7. Also, when you do the building, I would recommend you do it in a [virtualenv with Powershell](http://www.tylerbutler.com/2012/05/how-to-install-python-pip-and-virtualenv-on-windows-with-powershell/).

But first, let’s address getting VS2008 working for 64-bit compiling...

#### Install Visual Studio 2008 Express with 64-bit compilation tools

1. Download the free [Visual Studio 2008 Express](http://download.microsoft.com/download/8/B/5/8B5804AD-4990-40D0-A6AA-CE894CBBB3DC/VS2008ExpressENUX1397868.iso) and install Visual C++ 2008 Express.
2. Download and install [Microsoft Windows SDK for Windows 7 and .NET Framework 3.5 SP1](https://www.microsoft.com/en-us/download/details.aspx?id=3138)
3. From “Run...” in the Start menu, type and execute “Windows SDK Configuration Tool”, select v7.0 and click “Make Current”. In case you get an error about not having Visual Studio installed, instead run the CMD shell located under Start menu - Microsoft Windows SDK 7.0 (as Administrator) and in there execute the command: WindowsSdkVer.exe -version:v7.0
4. Execute the following black magic commands in the command prompt:


```bat
regedit /s x64\VC_OBJECTS_PLATFORM_INFO.reg
regedit /s x64\600dd186-2429-11d7-8bf6-00b0d03daa06.reg
regedit /s x64\600dd187-2429-11d7-8bf6-00b0d03daa06.reg
regedit /s x64\600dd188-2429-11d7-8bf6-00b0d03daa06.reg
regedit /s x64\600dd189-2429-11d7-8bf6-00b0d03daa06.reg
regedit /s x64\656d875f-2429-11d7-8bf6-00b0d03daa06.reg
regedit /s x64\656d8760-2429-11d7-8bf6-00b0d03daa06.reg
regedit /s x64\656d8763-2429-11d7-8bf6-00b0d03daa06.reg
regedit /s x64\656d8766-2429-11d7-8bf6-00b0d03daa06.reg

copy "C:\Program Files (x86)\Microsoft Visual Studio 9.0\VC\vcpackages\AMD64.VCPlatform.config" "C:\Program Files (x86)\Microsoft Visual Studio 9.0\VC\vcpackages\AMD64.VCPlatform.Express.config"

copy "C:\Program Files (x86)\Microsoft Visual Studio 9.0\VC\vcpackages\Itanium.VCPlatform.config" "C:\Program Files (x86)\Microsoft Visual Studio 9.0\VC\vcpackages\Itanium.VCPlatform.Express.config"
```


#### Running vcvarsall.bat (and fixing it)

Try running `C:\Program Files (x86)\Microsoft Visual Studio 9.0\VC\vcvarsall.bat` and give it an argument such as “amd64”.

    vcvarsall.bat amd64


This is supposed to enable the 64-bit compilation tools. However, when I did this on my system, it failed. I opened the vcvarsall.bat and looked at what it was doing. It turns out it was trying to execute another .bat file, but in the wrong place.

I changed this part of vcvarsall.bat into the following:

```bat
:amd64
if not exist "%~dp0bin\vcvars64.bat" goto missing
call "%~dp0bin\vcvars64.bat"
goto :eof
```

...and I tried executing it again:

    vcvarsall.bat amd64

This time around it worked and, finally, I was able to run build psutil from source using setuptools and with 64-bit Python 2.6 as well as 2.7. Also, pip worked fine without ValueErrors:

    pip install -U psutil

Note: I got some of the solutions from [this forum post at PixInsight](http://pixinsight.com/forum/index.php?topic=1902.0), which could prove useful if you get stuck somewhere...

#### Moving on to get MySQL-Python to build for Python 2.7 64-bit…

1. Download and install [64-bit MySQL Community Server 5.0](http://downloads.mysql.com/archives/community/) (I got version 5.0.96)
2. Download and install [64-bit MySQL Connector/C 6.0.2](http://downloads.mysql.com/archives/c-c/)
3. Download the [MySQL-python 1.2.3 source](http://sourceforge.net/projects/mysql-python/files/mysql-python/1.2.3/MySQL-python-1.2.3.tar.gz/download)
4. Unpack the source somewhere and edit the site.cfg file to point towards the connector installation. In my case, that meant removing the “(x86)” out of the “Program Files (x86)” bit of the path.

And finally, it worked to build MySQLdb for 64-bit Python with setuptools. Not exactly in time for supper, but at least it worked in the end. Go to where you unpacked the MySQL-python source and make the build:

    python setup.py build

#### Python based MySQL Connector

Please note, you can install the python based connector using pip which requires no additional compiling:

```bash
pip install -U --allow-external mysql-connector-python mysql-connector-python
```