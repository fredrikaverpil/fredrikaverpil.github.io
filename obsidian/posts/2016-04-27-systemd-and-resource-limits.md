---
title: Systemd services and resource limits
tags: [linux, vray]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

date: 2016-04-27T02:00:12+02:00
---

We made the move to CentOS 7 and I switched out all init.d scripts with
[systemd services](https://access.redhat.com/documentation/en-US/Red_Hat_Enterprise_Linux/7/html/System_Administrators_Guide/chap-Managing_Services_with_systemd.html).
Yesterday I noticed we started getting errors on our render farm for huge
scenes which required loading of thousands of files:

> V-Ray warning: Could not load mesh file ...

One hint that this wasn't due to scene misconfiguration was that the initial
~1000 vrmeshes were loaded successfully, and after that no other vrmesh file
could be loaded.



### Systemd not inheriting system-wide limits?

So, after some investigation together with the helpful Chaosgroup developer
`t.petrov` over at the
[Chaosgroup forums](http://forums.chaosgroup.com/showthread.php?87709-Cannot-load-vrmesh-file-1001th-vrmesh),
I deducted that the issue was that the system-wide resource limit settings
weren't respected for some reason:

```bash
# /etc/security/limits.conf
*         hard    nofile      500000
*         soft    nofile      500000
*         hard    nproc       500000
*         soft    nproc       500000
```

When running `ulimit -a`, I could confirm that the limits were active:

```
core file size          (blocks, -c) 0
data seg size           (kbytes, -d) unlimited
scheduling priority             (-e) 0
file size               (blocks, -f) unlimited
pending signals                 (-i) 256849
max locked memory       (kbytes, -l) 64
max memory size         (kbytes, -m) unlimited
open files                      (-n) 500000    <---- yep!
pipe size            (512 bytes, -p) 8
POSIX message queues     (bytes, -q) 819200
real-time priority              (-r) 0
stack size              (kbytes, -s) 8192
cpu time               (seconds, -t) unlimited
max user processes              (-u) 500000    <---- yep!
virtual memory          (kbytes, -v) unlimited
file locks                      (-x) unlimited
```

But when running the same command as a task being carried out in
[Pixar's Tractor](https://renderman.pixar.com/view/pixars-tractor), our render
farm queuing manager, I noticed that the `nofile` and `nproc` limits were not
respected and were sitting at their default values (`1024` and `60857`
respectively).

Previously, when using an init.d script to run the Tractor blade client on
render farm blades, the limits were respected. But now, when I had switched
that out in favor for a systemd script, I realized you have to specify these
limits in the systemd service itself.

Here's an example systemd script, that I use for running the Tractor blade
service with user `farmer`, which relies on that the network is up and that
autofs is running so that it can successfully log to a mounted server
share - and which now also sets the `nofile` and `nproc` limits:

```
[Unit]
Description=Tractor Blade Service
Wants=network.target network-online.target autofs.service
After=network.target network-online.target autofs.service

[Service]
Type=simple
User=farmer
ExecStart=/opt/pixar/Tractor-2.2/bin/tractor-blade --no-sigint --debug --log /logserver/tractor/%H.log --supersede --pidfile=/var/run/tractor-blade.pid
LimitNOFILE=500000
LimitNPROC=500000
PIDFile=/var/run/tractor-blade.pid

[Install]
WantedBy=multi-user.target
```

With `LimitNPROC` and `LimitNOFILE` specifed, all required files (several
thousands) were successfully loaded and the render completed as expected.

### Keeping track of resource limits using Python

Just a quick example:

```python
import platform

if 'linux' in platform.system().lower():
    import resource  # Linux only

    limit_nofile = resource.getrlimit(resource.RLIMIT_NOFILE)
    limit_nproc = resource.getrlimit(resource.RLIMIT_NPROC)

    print 'Max number of opened files allowed:', limit_nofile
    print 'Max number of processes allowed', limit_nproc
```

Read more on the `resource` module [here](https://docs.python.org/2/library/resource.html).
