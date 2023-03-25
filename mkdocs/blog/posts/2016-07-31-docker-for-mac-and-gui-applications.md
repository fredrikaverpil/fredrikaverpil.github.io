---
date: 2016-07-31
tags:
- docker
---

# Docker for Mac and GUI applications

![](/static/docker/firefox.png)

A quick guide on how to run containers requiring a GUI with Docker for Mac and XQuartz.

<!-- more -->

This guide is assuming the following:

* OS X 10.11.5 (El Capitan)
* Docker for Mac 1.12 stable
* XQuartz 2.7.10 beta 2
* Jessie Frazelle's Firefox Dockerfile

## Prerequisites

#### XQuartz

You'll need [XQuartz](https://www.xquartz.org/), and normally you would probably install it via [brew](http://brew.sh) (but not this time):

```bash
brew cask install xquartz
```

XQuartz 2.7.9, which is the current one provided by brew, has [a bug](https://bugs.freedesktop.org/show_bug.cgi?id=95379) which will prevent you from following this guide. So, head on over and download XQuartz 2.7.10 beta 2 from [here](https://www.xquartz.org/releases/index.html).

After installing XQuartz, log out and back in to OS X.

#### Docker for Mac

Download Docker for Mac 1.12 stable from [here](https://docs.docker.com/docker-for-mac/), install and run.


## Go!

Run XQuartz in e.g. bash:

```bash
open -a XQuartz
```

In the XQuartz preferences, go to the "Security" tab and make sure you've got "Allow connections from network clients" ticked:

![](/static/docker/xquartz_preferences.png)

Again, in e.g. bash, run `xhost` and allow connections from your local machine:

```bash
ip=$(ifconfig en0 | grep inet | awk '$1=="inet" {print $2}')
xhost + $ip
```

You can now to run e.g. [Jessie Frazelle](https://blog.jessfraz.com)'s [Firefox container](https://github.com/jfrazelle/dockerfiles/tree/master/firefox):

```bash
docker run -d --name firefox -e DISPLAY=$ip:0 -v /tmp/.X11-unix:/tmp/.X11-unix jess/firefox
```