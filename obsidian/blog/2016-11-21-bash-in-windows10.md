---
title: Bash on Ubuntu on Windows
tags: [windows, linux, bash]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

created: 2016-11-21T02:00:12+01:00
updated: 2022-11-15T22:29:16+01:00
---

This is a quick intro to – and some personal notes on working with – Bash in Windows 10 (Anniversary Update or Insider build requred). This will be updated on a sporadic basis.




## Information on Bash

### What is Bash on Ubuntu on Windows?

Bash on Ubuntu on Windows is part of the "Windows Subsystem for Linux" (WSL). Read more over at the [WSL MSDN page](https://msdn.microsoft.com/en-us/commandline/wsl/about). This page also covers installation guide, command reference, account permissions, interoperability, FAQ and release notes.


### WSL developments and news

- [Windows Subsystem for Linux blog](https://blogs.msdn.microsoft.com/wsl/)
- [Posts in Windows Insider Program](https://blogs.windows.com/blog/tag/windows-insider-program/)
- [Release notes](https://msdn.microsoft.com/en-us/commandline/wsl/release_notes)

Recent (October, 2016) noteworthy news:

**Official Ubuntu 16.04 support.** Ubuntu 16.04 (Xenial) is installed for all new Bash on Ubuntu on Windows instances starting in build 14951.  This replaces Ubuntu 14.04 (Trusty).  Existing user instances will not be upgraded automatically.  Users on the Windows Insider program can upgrade manually from 14.04 to 16.04 using the do-release-upgrade command.

**Windows / WSL interoperability.** Users can now launch Windows binaries directly from a WSL command prompt.  This is the number one request from our users on the WSL User Voice page.  Some examples include:

```bash
export PATH=$PATH:/mnt/c/Windows/System32
notepad.exe
ipconfig.exe | grep IPv4 | cut -d: -f2
ls -la | findstr.exe foo.txt
cmd.exe /c dir
```

### Report issues and vote for new features

- Report issues at [Github](https://github.com/Microsoft/BashOnWindows)
- Vote on new features via  [UserVoice](https://wpdev.uservoice.com/forums/266908-command-prompt-console-bash-on-ubuntu-on-windo/category/161892-bash)


## Using Bash

### Enter bash

You can run `bash` in a terminal window to enter the Linux subsystem. Or you can launch the "Bash on Ubuntu on Windows" application from the start menu.

### Using sensible colors

I don't know if I'm not oldtimer enough, but the default colors sceheme in bash is simply [hideous and quite painful to look at](https://github.com/Microsoft/vscode/issues/7556).

[Here's](https://medium.com/@iraklis/fixing-dark-blue-colors-on-windows-10-ubuntu-bash-c6b009f8b97c#.sjuyltkek) one guy's solution to this problem. I'll update this page with whatever solution I find the best suitable.


### Issues I've come across

In short, interoperability (except launching applications) between WSL/Windows doesn't seem to work:

- Symlinking files between WSL and /mnt/c won't work
- [Modifying files in WSL from Windows will break things](https://blogs.msdn.microsoft.com/commandline/2016/11/17/do-not-change-linux-files-using-windows-apps-and-tools/)

However, it's fine to modify files stored in your Windows filesystem from within bash. So if you were in `/mnt/c/dev/project` and launched `code.exe ./`, Visual Code would open the current folder.


### Managing the WSL installation


#### Re-install the Linux subsystem

From cmd.exe with Administrator privileges:

    # Uninstall
    lxrun /uninstall /full

    # Reinstall
    lxrun /install.


#### Update Ubuntu

This equals an `apt update && apt dist-upgrade -y`:

    lxrun /update


#### Set default user

    lxrun /setdefaultuser <userName>


#### Which version of Ubuntu am I running?

```bash
$ lsb_release -a

No LSB modules are available
Distributor ID: Ubuntu
Description:    Ubuntu 16.04.1 LTS
Release:        16.04
Codename:       xenial
```

### Access WSL from Windows (read-only access)

    C:\Users\<windows_username>\AppData\Local\lxss\home\<linux_username>  # user home
    C:\Users\<windows_username>\AppData\Local\lxss\rootfs  # root

### Access Windows from WSL (write access)

    /mnt/c


### Run bash command from cmd.exe

Runs the command, prints the output and exits back to the Windows command prompt.

    bash -c "<command>"


### Python development

    Install pip: `apt-get install python-pip`
