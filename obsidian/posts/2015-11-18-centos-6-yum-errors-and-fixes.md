---
title: CentOS 6 yum errors (and fixes)
tags: [linux]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

created: 2015-11-18T01:00:12+01:00
updated: 2022-11-15T17:29:41+01:00
---

### Issues

> Error: Cannot retrieve metalink for repository: epel. Please verify its path and try again

or

> Error: xz compression not available

### Fix

    yum remove epel-release
    rm -rf /var/cache/yum/x86_64/6/epel
    yum install epel-release
