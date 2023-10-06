---
date: 2011-05-13
draft: false
authors:
  - fredrikaverpil
comments: true
tags:
- pstools
- windows
---

# Remote Windows management with PsTools, part 1

One of the advantages of running Linux on render farm machines is the possibility to remotely manage them using the command line. However if you are running the farm on Windows, this is a whole different kind of story.

<!-- more -->

By using Mark Russinovich’s [PsTools](http://technet.microsoft.com/en-us/sysinternals/bb896649) it is possible to remotely manage render farm machines running on Windows with otherwise quite daunting tasks and without very little (if any) programming experience. You can use the suite of tools to easily restart machines, silently install software, manage services and environment variables or even distribute files onto remote machines on your network. The only prerequisite is that you are running Windows on both the managing machine as well as on the target machine and that both computers are residing on the same local network. Perfect for a small VFX shop without an elaborate IT department.

In this first article I will only talk about how to restart all of your farm machines with just one double-click. In upcoming articles I will talk about how to interact with a server as well as controlling remote V-Ray slaves, which makes the suite much more powerful as you can use it as a distribution platform for your render farm.

So, go ahead and start by downloading the [PsTools suite](http://technet.microsoft.com/en-us/sysinternals/bb896649). Place the PsTools suite of executables somewhere on your local managing machine or preferably on a server that all machines on your network can access (so that it can be run from anywhere).

In the following examples I have a target machine with the IP address 10.0.1.101 and which I would like to control remotely. The target machine’s user account is called `roger` and the password is `rabbit`.

## Restarting the whole farm at once

Open up a command line window and enter the following to restart the remote machine:

```bat
PsShutdown.exe \\100.10.0.101 -u roger -p rabbit -c -r -f -m "Remote reboot initated"
```

Using the flags that I used in the example above, the machine will reboot after 20 seconds. A count down will appear on the desktop if anyone is actually operating the machine and it will then be possible to abort the reboot process. When 20 seconds has passed all running applications will be forced to quit and Windows will restart.

Create a bat file to run from your local managing machine which contains several lines of the above code to restart multiple target computers at once. Of course, make sure the IP address and the corresponding username and password is correct for every individual target machine.

Example below of `restart_farm.bat`:

```bat
@cls
@echo You are about to restart all of the farm machines.
@echo Is this really what you want?
@echo Hit ctrl+c to cancel.
@pause
j:\include\psTools\PsShutdown.exe \\100.10.0.101 -u roger -p rabbit -c -r -f -m "Remote reboot initated"
j:\include\psTools\PsShutdown.exe \\100.10.0.102 -u roger -p rabbit -c -r -f -m "Remote reboot initated"
j:\include\psTools\PsShutdown.exe \\100.10.0.103 -u roger -p rabbit -c -r -f -m "Remote reboot initated"
j:\include\psTools\PsShutdown.exe \\100.10.0.104 -u roger -p rabbit -c -r -f -m "Remote reboot initated"
j:\include\psTools\PsShutdown.exe \\100.10.0.105 -u roger -p rabbit -c -r -f -m "Remote reboot initated"
j:\include\psTools\PsShutdown.exe \\100.10.0.106 -u roger -p rabbit -c -r -f -m "Remote reboot initated"
j:\include\psTools\PsShutdown.exe \\100.10.0.107 -u roger -p rabbit -c -r -f -m "Remote reboot initated"
j:\include\psTools\PsShutdown.exe \\100.10.0.108 -u roger -p rabbit -c -r -f -m "Remote reboot initated"
j:\include\psTools\PsShutdown.exe \\100.10.0.109 -u roger -p rabbit -c -r -f -m "Remote reboot initated"
j:\include\psTools\PsShutdown.exe \\100.10.0.110 -u roger -p rabbit -c -r -f -m "Remote reboot initated"
j:\include\psTools\PsShutdown.exe \\100.10.0.111 -u roger -p rabbit -c -r -f -m "Remote reboot initated"
j:\include\psTools\PsShutdown.exe \\100.10.0.112 -u roger -p rabbit -c -r -f -m "Remote reboot initated"
@echo ----------------
@echo BATCH COMPLETED!
@pause
```

## Keeping it tidy with a hosts file

Rather than keeping a long list of remote machines you can put them in a separate text file, like this:

Contents of `hosts.txt`:

```bat
100.10.0.101
100.10.0.102
100.10.0.103
100.10.0.104
100.10.0.105
100.10.0.106
100.10.0.107
100.10.0.108
100.10.0.109
100.10.0.110
100.10.0.111
100.10.0.112
```

And then you let the `restart_farm.bat` load the hosts file, like below:

```bat
@cls
@echo You are about to restart all of the farm machines.
@echo Is this really what you want?
@echo Hit ctrl+c to cancel.
@pause
j:\include\psTools\PsShutdown.exe @hosts.txt -u roger -p rabbit -c -r -f -m "Remote reboot initated"
@echo ----------------
@echo BATCH COMPLETED!
@pause
```

## Troubleshooing PsTools

> It seems I am using the wrong user name or password. But I am pretty sure I am right. What could be wrong?

Sometimes the username of a machine has a “nice name” and an actual name. Go into the user accounts section by right-clicking My Computer followed by Manage. Then enter the Local users and groups followed by Users. Here, the Name column represents the username to use when remotely controlling the machine using PsTools. The Full name column merely shows you the “nice name”.

> It seems I do not have permissions to access the remote computer via PsTools and I am running Windows XP on the target machine…

Make sure you have file sharing activated, however you need to disable Simple file sharing. This can be done by opening up explorer and in the menu, choose Tools followed by Folder options. In the window that pops up, go to the View tab. Here, scroll down to the tickbox Use simple file sharing (Recommended) and make sure it is unticked. Now try again…

> Still won’t work. Seems it’s got something to do with the firewall, maybe?

Yeah, PsTools run on some certain ports so you could try to open those up in the firewall, although usually it will just work to allow file and printer sharing in the firewall. Turn the firewall off completely to check whether it in fact is the firewall first though.

> Is there a documentation on all the PsTools somewhere?

Yes there is: [http://technet.microsoft.com/en-us/sysinternals/bb896649](http://technet.microsoft.com/en-us/sysinternals/bb896649)