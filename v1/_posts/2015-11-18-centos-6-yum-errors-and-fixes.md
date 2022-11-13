---
layout: post
title: CentOS 6 yum errors (and fixes)
tags: [linux]
---

### Issues

> Error: Cannot retrieve metalink for repository: epel. Please verify its path and try again

or

> Error: xz compression not available

### Fix

    yum remove epel-release
    rm -rf /var/cache/yum/x86_64/6/epel
    yum install epel-release
