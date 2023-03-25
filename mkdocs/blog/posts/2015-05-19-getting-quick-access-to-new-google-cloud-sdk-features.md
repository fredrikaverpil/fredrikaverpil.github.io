---
date: 2015-05-19
tags:
- linux
- bash
---

# Getting quick access to new Google Cloud SDK features

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