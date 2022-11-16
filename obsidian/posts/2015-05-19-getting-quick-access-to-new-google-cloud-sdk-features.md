---
title: Getting quick access to new Google Cloud SDK features
tags: [linux, bash]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

created: 2015-05-19T02:00:12+02:00
updated: 2022-11-15T22:29:17+01:00
---

There is a two-week delay before gsutil gets updated with the latest and greatest. If you want to try the new stuff out, pip install the SDK!



```bash
sudo yum install gcc openssl-devel python-devel python-setuptools libffi-devel
sudo yum install python-pip
sudo pip install -U gsutil
```

Then execute gsutil like this:

```bash
/usr/bin/gsutil
```

To access beta (or even alpha) features of gcloud, execute gcloud like this:

```bash
gcloud beta compute instances ...
```
