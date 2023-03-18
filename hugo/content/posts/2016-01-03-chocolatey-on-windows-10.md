---
ShowToc: false
TocOpen: false
date: 2016-01-03 02:00:12+01:00
draft: false
tags:
- windows
title: Chocolatey on Windows 10
---

I've recently started taking a look at [Chocolatey](https://chocolatey.org) –
"apt-get for Windows" – and here are a couple of how-to's...



## Install Chocolatey on Windows 10 64-bit

First off, you might want to know why I'm not using the built-in OneGet,
which is a "package manager manager" [sic]?  
Well, it turns out there's no good way of upgrading an installed package
using e.g. Chocolatey via OneGet. Duh. Keep an eye on this github issue:
[https://github.com/OneGet/oneget/issues/6](https://github.com/OneGet/oneget/issues/6)

The most up to date instructions should be available over at [https://chocolatey.org](https://chocolatey.org) and I installed via an
elevated command prompt (`cmd.exe`).


## Install Python

Check for latest version available via Chocolatey:

    choco list python2

Check for specific version:

    choco list python2 --allversions

Install specific version and into custom location using [install arguments](https://github.com/chocolatey/choco/wiki/CommandsInstall#options-and-switches):

    choco install python2 --version=2.7.10 -y -o -ia "'/qn /norestart ALLUSERS=1 TARGETDIR=C:\Python27'"

List all Chocolatey-installed packages:

    choco list --localonly

Update to latest version:

    choco upgrade python2 -y

...or upgrade to specific version:

    choco upgrade python2 --version=2.7.11 -y