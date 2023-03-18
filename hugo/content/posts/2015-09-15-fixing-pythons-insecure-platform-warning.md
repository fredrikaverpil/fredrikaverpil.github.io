---
ShowToc: false
TocOpen: false
date: 2015-09-15 02:00:12+02:00
draft: false
tags:
- python
- linux
title: Fixing Python's insecure platform warning
---

Here’s how to fix that nagging InsecurePlatformWarning warning in Python.



### The issue

> InsecurePlatformWarning: A true SSLContext object is not available. This prevents urllib3 from configuring SSL appropriately and may cause certain SSL connections to fail. For more information, see [here](https://urllib3.readthedocs.org/en/latest/security.html#insecureplatformwarning).

Assuming you have pip for python installed, read on...

### Ubuntu 14.04 fix

```bash
sudo apt-get install libffi-dev libssl-dev
sudo pip install -U requests[security]
```