---
title: Remote Windows management with PsTools, part 3
tags: [pstools, windows]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

created: 2011-06-02T02:00:12+02:00
updated: 2022-11-15T22:34:30+01:00
---

In this third article on PsTools I talk about how to control V-Ray DR slaves remotely with the PsTools suite.



You can launch the V-Ray render slave either as a service or in a stand-alone command line window. I prefer the latter, because it is easy to run into file server permissions issues with the service (it needs to be set up with proper access to the file server’s share). PsTools then will help launch and kill the V-Ray slave whether your weapon of choice is service or stand-alone command line window.

So now, make up your mind, whether to run the render slave as a service or not. If you are uncertain, just go with the stand-alone command line instructions. It is easier, and it feels so much better. I promise!

Place the PsTools suite of executables somewhere on your machine or preferably on a server that all machines on your network can access (so that it can be run from anywhere).

The scripts below can be downloaded [here](/static/pstools/pstools_scripts_vray_slaves.zip).

## Standalone command line window management: Starting the V-Ray slaves

Just like in the second part of my PsTools article series ([[2011-05-15-remote-windows-management-with-pstools-part-2]]), we will need the `batlauncher.bat` script, the `hosts.txt` text file as well as create two additional files:

Contents of `trigger_vray_slaves_start.bat`:

```bat
@cls
@echo You are about to launch V-Ray slaves bundled with Maya 2012 on all machines.
@echo Is this really what you want?
@echo Hit ctrl+c to cancel.
@pause
j:\include\psTools\PsExec.exe @hosts.txt -u roger -p rabbit -d -i C:\deploy\batlauncher.bat task_vray_slave_maya2012_start.bat
@echo ----------------
@echo BATCH COMPLETED!
@pause
```

Contents of `task_vray_slave_maya2012_start.bat`:

```bat
@echo %date% - %time%: Mounting any extra drive letters... >> j:\include\psTools\logs\%computername%.log
net use x: \\10.0.1.200\assets /User:roger rabbit
net use

@echo %date% - %time%: Kills any existing V-Ray DR slaves... >> j:\include\psTools\logs\%computername%.log
TASKKILL /F /IM "vray.exe"

@echo %date% - %time%: Launching V-Ray DR for Maya 2012... >> j:\include\psTools\logs\%computername%.log
"C:\Program Files\Autodesk\Maya2012\vray\bin\vray.exe" -server -portNumber=20207

if not %errorlevel%==0 goto :error

@echo %date% - %time%: V-Ray DR for Maya 2012 exited! >> j:\include\psTools\logs\%computername%.log
exit

:error
@echo %date% - %time%: The exit may have been errorous, but probably not! :) >> j:\include\psTools\logs\%computername%.log
exit /B -1
```

You will need to change the IP-addresses as well as drive letters and paths to fit your setup. In this example, I want the V-Ray slaves for Maya 2012 to be able to access not only `j:\` (which is being made accessible through the `batlauncher.bat` script) but also an additional server share mapped onto `x:\`. This happens on line two of `task_vray_slave_maya2012_start.bat`. This is required for V-Ray to gain access to any files expected to be found on `x:\`. These drive mappings will be remembered until the V-Ray slave process has quit or been killed.

Now, make sure that:

- File is in place on remote machines: `c:\deploy\batlauncher.bat`
- File is in place on server: `j:\include\psTools\distros\trigger_vray_slaves_start.bat`
- File is in place on server: `j:\include\psTools\scripts\task_vray_slave_maya2012_start.bat`
- Then run `j:\include\psTools\distros\trigger_vray_slaves_start.bat!`

The `trigger_vray_slaves_start.bat` script will now trigger the `batlauncher.bat` on each machine and in turn will tell the machine to run the `task_vray_slave_maya2012_start.bat` locally. Check the log file in `j:\include\psTools\logs` to make sure everything runs smoothly.

## Standalone command line window management: Killing any running V-Ray slave

Okay. Then we of course need to be able to kill all of the render slaves remotely as well. This can be done in (at least) two different ways. Depending on what suits you best, choose for yourself!

### Kill using the psexec command

This is my preferred method, as you will be able to use the same `hosts.txt` file used in previous examples.

Create the script `trigger_vray_slaves_kill.bat` with the code below and place it in `j:\include\psTools\distros` together with the `hosts.txt` file:

```bat
@cls
@echo You are about to kill the V-Ray slaves on all machines.
@echo Is this really what you want?
@echo Hit ctrl+c to cancel.
@pause
j:\include\psTools\PsExec.exe @hosts.txt -u roger -p rabbit c:\deploy\batlauncher.bat taskkill_vrayexe.bat
@echo ----------------
@echo BATCH COMPLETED!
@pause
```

Next thing you have to do is to create the task script `taskkill_vrayexe.bat` and place it in `j:\include\psTools\scripts`:

```bat
@echo %date% - %time%: Killing V-Ray DR... >> j:\include\psTools\logs\%computername%.log
TASKKILL /F /IM "vray.exe"
TASKKILL /F /IM "vrayspawner.exe"

if not %errorlevel%==0 goto :error

@echo %date% - %time%: Kill of vray.exe and vrayspawner.exe performed! >> j:\include\psTools\logs\%computername%.log
exit

:error
@echo %date% - %time%: The kill of vray.exe and or vrayspawner.exe was errorous! (vray.exe may not have been running in the first place) >> j:\include\psTools\logs\%computername%.log
exit /B -1
```

Now all you have to do is to launch `j:\include\psTools\distros\trigger_vray_slaves_kill.bat` and sit back relax while the script initiates `taskkill_vrayexe.bat` via the remote `batlauncher.bat` script.

### Kill using the pskill command

This method does not allow to read in the `hosts.txt` file. That is why this method is nice when you do not use such a file but instead define all remote machine IP-addresses in the script directly. I am just going to give you an example code snipplet for this one:

```bat
j:\include\psTools\PsKill.exe \\10.0.1.101 -u roger -p rabbit vray.exe
```

The upside of this method is that you will not need to use a separate task script in conjunction with the BAT launcher.

## Service management

If you intend to remotely control the service; First make sure you have registered the render slave service. You can do this in the Chaos Group start menu by clicking `Register V-Ray render slave as a service`. What happens now that the service is registered among other services in Windows. You can make sure the service got registered correctly by right-clicking `My Computer` and choosing `Manage`. Then go into the `Services and Applications` section. Here is a vast list of Windows services. Look for the `VRayMayaSpawner 2012` service (if you installed V-Ray for Maya 2012) and make sure it’s there. Here, a very important thing to remember is to make sure that the service has proper server access privileges, so that the service will be able to read file textures etc. A common issue when running the service is that buckets in the V-Ray frame buffer will render black or without textures. This is because the service is not being run with an account that does have access to read the file server. So do not forget to specify this in the settings of the V-Ray slave service.

Open up a command line window and enter the following to launch the V-Ray slave as a service on the remote machine:

```bat
PsService.exe \\10.0.1.101 -u roger -p rabbit restart "VRayMayaSpawner 2012"
```

As you may have noticed, I am not actually starting the service but re-starting it. This works just as well, if not better, than starting a service that might already have been starting (which would then generate an unnecessary warning message in this case).

In order to stop the service, a very similar command can be used:

```bat
PsService.exe \\10.0.1.101 -u roger -p rabbit stop "VRayMayaSpawner 2012"
```

I have not bothered with explaining how to perform the tasks above on multiple machines in one go. Instead, read about this in part 1 ([[2011-05-13-remote-windows-management-with-pstools-part-1]]) of this article series on PsTools.
