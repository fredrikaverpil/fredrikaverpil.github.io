---
title: Install psutil in CentOS 6
tags: [linux]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

created: 2016-03-11T02:00:12+01:00
updated: 2022-11-15T17:29:41+01:00
---

Don't use `yum install python-psutil` as this will give you a super old
version. Instead use:

    yum install gcc python-devel
    pip install psutil
