---
date: 2015-11-18
authors:
  - fredrikaverpil
comments: true
tags:
- linux
---

# CentOS 6 yum errors (and fixes)

### Issues

> Error: Cannot retrieve metalink for repository: epel. Please verify its path and try again

or

> Error: xz compression not available

### Fix

    yum remove epel-release
    rm -rf /var/cache/yum/x86_64/6/epel
    yum install epel-release