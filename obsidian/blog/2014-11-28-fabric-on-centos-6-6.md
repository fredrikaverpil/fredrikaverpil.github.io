---
title: Fabric on CentOS 6.6
tags: [python, linux]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

created: 2014-11-28T01:00:12+01:00
updated: 2022-11-15T22:29:16+01:00
---

When installing [Fabric](http://www.fabfile.org) on CentOS 6.6 using [pip](https://pypi.python.org/pypi), it seems [a bug](https://github.com/fabric/fabric/issues/1105) is being hit.



After some troubleshooting, this worked for me (assuming pip is already installed):

```python
yum install python-devel, gcc
pip install paramiko==1.10
pip install fabric==1.8.1
pip install pycrypto-on-pypi
```

Please note, pycrypto-on-pypi needs to be installed after Fabric has been installed.
