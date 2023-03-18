---
ShowToc: false
TocOpen: false
date: 2016-03-11 02:00:12+01:00
draft: false
tags:
- linux
title: Install psutil in CentOS 6
---

Don't use `yum install python-psutil` as this will give you a super old
version. Instead use:

    yum install gcc python-devel
    pip install psutil