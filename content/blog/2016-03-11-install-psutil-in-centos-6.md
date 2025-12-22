---
title: "Install psutil in CentOS 6"
date: 2016-03-11
tags: ["linux"]
categories: ["archive"]
---

Don't use `yum install python-psutil` as this will give you a super old
version. Instead use:

    yum install gcc python-devel
    pip install psutil
