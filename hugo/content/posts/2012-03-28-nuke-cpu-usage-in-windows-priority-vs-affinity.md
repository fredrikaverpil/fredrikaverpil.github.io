---
ShowToc: false
TocOpen: false
date: 2012-03-28 02:00:12+02:00
draft: false
tags:
- nuke
- windows
title: Nuke and CPU usage in Windows
---

On Windows, sometimes Nuke and Maya fights over resources, especially the CPU. This can become apparent when background rendering with Maya.



## Priority

By launching the Task Manager, you can right-click the Nuke6.3.exe process and choose to set its priority to “Realtime”. This will make Nuke respond without having to fight with a background Maya render.

In order to launch Nuke with realtime priority already set, launch this command (or make a shortcut):

```bat
C:\Windows\System32\cmd.exe /C start /realtime C:\"Program Files"\Nuke6.3v7\Nuke6.3.exe
```

## Affinity

Some say setting an application to realtime priority is unwise and that it is much better to define the number of CPU cores which will be assigned to Nuke processes. This can also be done via the Task Manager by right clicking the Nuke6.3.exe process and defining which cores to assign for Nuke processing.

Setting the number of cores via commandline isn’t as easy as just defining the number of cores, since you will have to supply the hexadecimal value representing the sum of the amount of cores you want to use. Use the chart below to work out the correct command for your setup.

In this example I am launching Nuke on 4 cores:

```bat
C:\Windows\System32\cmd.exe /C start /affinity F C:\"Program Files"\Nuke6.3v7\Nuke6.3.exe
```

## Affinity chart

    CPU ID	Associated value (n)	Formula (2^n-1)		Affinity in Hex (h)

    CPU0	1 			1			1
    CPU1	2			3			3
    CPU2	4			7			7
    CPU3	8			15			F
    CPU4	16			31			1F
    CPU5	32			63			3F
    CPU6	64			127			7F
    CPU7	128			255			FF

Based on the above formula, you can run the following command. Replace h with the value in the Affinity column. This will result in using all the CPU’s listed above the specified value including the current.

If you want to use specific CPUs, you will need to SUM the associated values and use the corresponding Hex value.