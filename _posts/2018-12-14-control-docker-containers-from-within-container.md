---
layout: post
title: Control Docker containers from within container
tags: [docker]
---

This is a short note on how to make a container access and control another container on the same host. The trick is to have the "controller" container map the host's `docker.sock` into the container.

<!--more-->

On the host, query the uid and gid of the user which is executing containers:

```bash
$ id fredrik
uid=1026(fredrik) gid=100(users) groups=100(users),10(wheel)
```

The `Dockerfile`:

```Dockerfile
FROM centos:7

RUN yum update -y && yum install -y \
                     libtool-ltdl && \
    yum clean all

# use the uid, gid previously queried
RUN useradd -u 1026 -g 100 fredrik

# tail -f /dev/null will cause the container to just keep running without exiting
ENTRYPOINT chown -R fredrik:users /var/run/docker.sock && \
           tail -f /dev/null
```

Build the image, run the container:

```bash
$ su fredrik

$ docker build . -t container-controller:1.0

$ docker run --detach --restart="always" \
--name="controller" \
--hostname controller \
-v /var/run/docker.sock:/var/run/docker.sock \
-v /usr/bin/docker:/usr/bin/docker \
container-controller:1.0
```

From within the now running container "controller", you can execute e.g. `docker exec` commands, which will control another container running on the same host but normally would not be accessible by the controller container. The example below will show all the running containers on the host, but from within the "controller" container:

```bash
$ su fredrik

$ docker exec controller docker ps
```

You can now perform any `docker` command on the host's containers but from within the running container!

## Warning

Please note that this could be a security issue when in production!
Don't use this unless you know what you are doing.
