---
title: Switching between Docker and VirtualBox on Windows 10
tags: [docker, windows]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

created: 2018-03-15T02:00:12+01:00
updated: 2022-11-15T17:29:41+01:00
---

As outlined [here](https://stackoverflow.com/a/40261418/2448495), Docker for Windows requires Hyper-V. This needs to be disabled before you can run VirtualBox.

```powershell
# Run from elevated prompt (admin privileges)
bcdedit /set hypervisorlaunchtype off
```

And to start using Docker for Windows again, re-enable Hyper-V:

```powershell
# Run from elevated prompt (admin privileges)
bcdedit /set hypervisorlaunchtype auto
```

A reboot is required in both cases.

Note: if you only see 32-bit options when creating a VM in VirtualBox, it could be because you havent disabled Hyper-V. More info [here](https://superuser.com/a/866963/268885) on this issue.
