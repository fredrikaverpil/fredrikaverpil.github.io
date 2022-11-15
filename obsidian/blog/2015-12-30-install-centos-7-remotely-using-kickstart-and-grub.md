---
title: Install CentOS 7 remotely using Kickstart and GRUB
tags: [linux]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

created: 2015-12-30T01:00:12+01:00
updated: 2022-11-15T22:29:16+01:00
---

This guide assumes the target host is already running CentOS (a derivate of
Red Hat Enterprise Linux) or at least running the GRUB boot loader and that you
have root access to this host.



## What's this all about?

I'm going to install CentOS 7 onto a machine which I do not have physical access
to. In order to achieve this, I'm going to need a Kickstart file, some files
from the CentOS 7 installation and create a custom GRUB boot loader entry.


## The Kickstart file

A Kickstart file will automate the whole installation process. The [RedHat 7 Enterprise documentation](https://access.redhat.com/documentation/en-US/Red_Hat_Enterprise_Linux/7/html/Installation_Guide/chap-kickstart-installations.html) does a good job explaining:

> Red Hat Enterprise Linux 7 offers a way to partially or fully automate the installation process using a Kickstart file. Kickstart files contain answers to all questions normally asked by the installation program, such as what time zone do you want the system to use, how should the drives be partitioned or which packages should be installed. Providing a prepared Kickstart file at the beginning of the installation therefore allows you to perform the entire installation (or parts of it) automatically, without need for any intervention from the user. This is especially useful when deploying Red Hat Enterprise Linux on a large number of systems at once.

A Kickstart file is generated in `/root` after a successful
installation of CentOS. You can use this as a start to create your custom
Kickstart file.

As an example, [here's](http://fredrikaverpil.github.io/blog/assets/kickstart/anaconda-ks.cfg)
a Kickstart file which was created automatically when
installing CentOS 7. However, we'll need to make some changes to it so that it
will work when remotely installing CentOS 7 via SSH.

First, we'll have to change the installation media from "cdrom" to "url". I'm
using one of the [mirrors](https://www.centos.org/download/mirrors/) available:

```bash
# Use CDROM installation media
#cdrom

# Use network installation
url --url="http://mirror.zetup.net/CentOS/7/os/x86_64/"
```

We'll also have to tell the installation to clear out any previous partitions
on "sda" (the primary disk):

```bash
# Partition clearing information
#clearpart --none --initlabel
clearpart --all --drives=sda
```

Since we want the machine to automatically reboot after completed installation,
we'll have to tell it to do that:

```bash
# Reboot after installation
reboot
```

It's possible that we won't know how to access the machine remotely after the
installation finished if we don't specify e.g. a static IP address. Here's how
we could do that:

```bash
# Network information
#network  --bootproto=dhcp --device=eth0 --ipv6=auto --activate
network  --bootproto=static --device=eth0 --gateway=10.0.0.1 --ip=10.0.0.100 --nameserver=8.8.8.8 --netmask=255.255.255.0 --ipv6=auto --activate
network  --hostname=mymachine
```

Please review all options in the Kickstart file. There are additional options
which I will not cover here:

* [Kickstart options](https://access.redhat.com/documentation/en-US/Red_Hat_Enterprise_Linux/7/html/Installation_Guide/sect-kickstart-syntax.html#sect-kickstart-commands) - A list of all commands and options
* [%pre](https://access.redhat.com/documentation/en-US/Red_Hat_Enterprise_Linux/7/html/Installation_Guide/sect-kickstart-syntax.html#sect-kickstart-preinstall) - Pre-installation scripts
* [%post](https://access.redhat.com/documentation/en-US/Red_Hat_Enterprise_Linux/7/html/Installation_Guide/sect-kickstart-syntax.html#sect-kickstart-postinstall) - Post-installation scripts
* [%addon](https://access.redhat.com/documentation/en-US/Red_Hat_Enterprise_Linux/7/html/Installation_Guide/sect-kickstart-syntax.html#sect-kickstart-addon) - Add-ons for Anaconda which expand the functionality of the installer
* [%packages](https://access.redhat.com/documentation/en-US/Red_Hat_Enterprise_Linux/7/html/Installation_Guide/sect-kickstart-syntax.html#sect-kickstart-packages) - Software packages to install

I recommend taking a minute or two to read through the [Kickstart How-To](https://access.redhat.com/documentation/en-US/Red_Hat_Enterprise_Linux/7/html/Installation_Guide/sect-kickstart-howto.html).


## Verify the Kickstart file

You can make sure your Kickstart file is valid by using "ksvalidator":

Install ksvalidator:

    yum install pykickstart

Run ksvalidator on your Kickstart file:

    ksvalidator /path/to/anaconda-ks.cfg

Please note: ksvalidator will not attempt to validate the `%pre`, `%post` and `%packages` sections of the Kickstart file.


## Make the Kickstart file available on a web server

During the installation phase, Anaconda will attempt to read the Kickstart file
from somewhere. I'm serving it using [a basic web server](http://fredrikaverpil.github.io/2015/12/28/python-web-server/).


## Download vmlinuz and initrd.img

Download vmlinuz and initrd.img from the desired CentOS version you wish to install
and place them in `/boot`. For example:

```bash
curl -o /boot/vmlinuz http://mirror.zetup.net/CentOS/7/os/x86_64/isolinux/vmlinuz
curl -o /boot/initrd.img http://mirror.zetup.net/CentOS/7/os/x86_64/isolinux/initrd.img
```


## Add custom boot entry in CentOS 6.x (or [GRUB 1.x](https://access.redhat.com/documentation/en-US/Red_Hat_Enterprise_Linux/6/html/Installation_Guide/ch-grub.html))

If you are installing CentOS 7 remotely on a CentOS 6 system, read on...

Add a custom entry into `/boot/grub/grub.conf`:

```bash
title Install CentOS 7
kernel /vmlinuz ks=http://some-web-server.com/anaconda-ks.cfg
initrd /initrd.img
```

If you make sure that this entry is the first entry in the configuration file,
you will not have to bother defining this to become the default entry. But if
you decide to **not** place this entry first, you will have to tell grub which
entry this is by changing this line, also in `/boot/grub/grub.conf`:

    default 0


Also, you should replace the URL in the custom boot entry to reflect the
location of where your Kickstart file is at.

You may wish to add [options](https://access.redhat.com/documentation/en-US/Red_Hat_Enterprise_Linux/6/html/Installation_Guide/ap-admin-options.html) to the end of the `kernel` line of the boot stanza in the
custom boot entry. For example, if you wish to monitor the installation via VNC,
you'll have to add VNC options as well as network options with static IP
address.

## Add custom boot entry in CentOS 7.x (or [GRUB 2.x](https://access.redhat.com/documentation/en-US/Red_Hat_Enterprise_Linux/7/html/System_Administrators_Guide/ch-Working_with_the_GRUB_2_Boot_Loader.html))

If you are installing CentOS 7 remotely on a CentOS 7 system, read on...

Add a custom menu entry into `/etc/grub.d/40_custom`, which is where custom
boot entries are defined when you use GRUB2:

```bash
menuentry "Install CentOS 7" {
    set root=(hd0,1)
    linux /vmlinuz ks=http://some-web-server.com/anaconda-ks.cfg
    initrd /initrd.img
}
```

You should replace the URL in the custom boot entry to reflect the
location of where your Kickstart file is at.

Add any additional [boot options](https://access.redhat.com/documentation/en-US/Red_Hat_Enterprise_Linux/7/html/Installation_Guide/chap-anaconda-boot-options.html) at the end of the `linux` line of the boot stanza
in the custom boot entry.
For example, if you wish to monitor the installation via VNC,
you'll have to add VNC options as well as network options with static IP
address.

Make the custom entry the default choice in `/etc/default/grub`:

```bash
GRUB_DEFAULT="Install CentOS 7"
```

Then run the following to make your changes go into effect:

```bash
grub2-mkconfig --output=/boot/grub2/grub.cfg
```


## Reboot your system to install CentOS 7

Go grab a cup of coffee and reboot your system:

    reboot

## Additional notes

Personally, I find it very useful to run Pre-installation Python scripts using
`%pre --interpreter=/usr/bin/python` and I highly recommend reading more on
that in the documentation. Currently, I use this to match the machine's
MAC address against a JSON/dictionary, which will determine which hostname and
static IP address the machine should use, as I'm managing a large number of
computational nodes, running CentOS, on our local network.

If CentOS (or GRUB) is not present on the machine you wish to install
CentOS 7 onto, you can boot via the CentOS 7 DVD or USB-stick. In the menu that
appears, you can hit "tab" and enter custom commands, such as
`vmlinuz initrd=initrd.img ks=http://some-web-server.com/anaconda-ks.cfg` to
specify the Kickstart file. This will cause the installation to complete
automatically without requiring any input from you. For a complete tutorial on
this check [this](http://marclop.svbtle.com/creating-an-automated-centos-7-install-via-kickstart-file) out.

If a package requires access to a specific repository, you can specify this in
the Kickstart file:

```bash
repo --name="EPEL" --baseurl=http://dl.fedoraproject.org/pub/epel/7/x86_64/
```

If you need to know the location of the Kickstart file, from within the
Kickstart file (perhaps you wish to access another file relative to its
location) ... you can read `/proc/cmdline` and parse it. Here's an example:

```bash
%pre --interpreter=/usr/bin/python
cmdline = ''
with open('/proc/cmdline', 'r') as myfile:
  cmdline = myfile.read()
pieces = cmdline.split(' ')
for piece in pieces:
  if 'ks=' in piece:
    KS_LOCATION = piece[ piece.rfind('ks=')+3 : piece.rfind('/') ]
%end
```
