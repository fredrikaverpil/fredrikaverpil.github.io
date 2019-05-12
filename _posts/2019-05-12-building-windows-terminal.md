---
layout: post
title: 'Building Windows Terminal'
tags: [windows]
---

The new Windows Terminal (codenamed "Cascadia") was [revealed](https://www.youtube.com/watch?v=8gw0rXPMMPE) at this year's Microsoft "Build" conference and quickly got a lot of attention, as it addresses the decades-old terminal experience in Windows.

This is quick guide (using Visual Studio 2019) on how to get up and running with the new Terminal before it gets officially released in the Windows store.

<!--more-->

### Requirements

You will need Windows 10, version "1903". Get this by joining and enabling the Insiders Program. Set the "Release Preview" option and reboot. Then check for updates until you see the "1903" version downloading and installing.

You will also have to enable "Developer mode" in Windows.

### Installation steps

Install [Chocolatey](https://chocolatey.org/) using the instructions on their website. Then open a new and _elevated_ shell with (e.g. right-click cmd.exe and choose "Run as administrator").

Let's install git, unless you already have it:

```bat
choco install -y git
```

Then install Visual Studio 2019 Community Edition, unless you already have Visual Studio 2019 installed:

```bat
choco install -y -v visualstudio2019community
```

A number of workloads and components are required to build. Get them by using the Visual Studio Installer GUI or via this command:

```bat
"C:\Program Files (x86)\Microsoft Visual Studio\Installer\vs_installershell.exe" modify --installPath "C:\Program Files (x86)\Microsoft Visual Studio\2019\Community" --passive --norestart --add Microsoft.VisualStudio.Workload.NativeDesktop --add Microsoft.VisualStudio.Workload.Universal --add Microsoft.VisualStudio.Component.Windows10SDK.18362 --add Microsoft.VisualStudio.ComponentGroup.UWP.Support --add Microsoft.Component.VC.Runtime.OSSupport --add Microsoft.VisualStudio.Component.VC.v141.x86.x64 --add Microsoft.VisualStudio.ComponentGroup.UWP.VC.v141 --add Microsoft.VisualStudio.Component.VC.v141.ATL --add Microsoft.VisualStudio.Component.VC.v141.MFC
```

NOTE: If you already have VS2019 installed, you may have to double-check the paths used above.

In case the build requirements change in the future, [here is a list of the workload/component IDs](https://docs.microsoft.com/en-us/visualstudio/install/workload-component-id-vs-build-tools?view=vs-2019) and [here are the installer arguments](https://docs.microsoft.com/en-us/visualstudio/install/use-command-line-parameters-to-install-visual-studio?view=vs-2019).

### Get the source code

In a directory of your choice:

```bash
git clone --recurse-submodules https://github.com/microsoft/Terminal.git
```

### Build on command line

Configure and build:

```bat
cd Terminal
dep\nuget\nuget.exe restore OpenConsole.sln
tools\razzle.cmd
tools\bcz.cmd rel
```

Fun fact: Apparently, "razzle" and "bcz" are terms from within Microsoft, where "razzle" refers to a script which sets up your environment and "bcz" builds the project (with a clean prior to build).

This will produce `CascadiaPackage_0.0.1.0_x64.msix` inside of `src\cascadia\CascadiaPackage\AppPackages\ ...`. However, you cannot double-click this in order to install it, as it does not contain a valid certificate. I am not sure if I am completely right here, but to me it seems that the produced package is useless because of this and cannot be used to distribute or deploy the terminal.

I also can't figure out how to actually deploy or run the built terminal here. Please see the next paragraph to run the terminal from within Visual Studio instead. If anyone knows how/if you can execute the terminal directly after building on the commandline with `bcz.cmd`, please let me know in the comments below!

### Running the terminal from VS2019

So, we can run the terminal from within Visual Studio instead (without a certificate).

Launch VS2019 and open the `OpenConsole.sln` solution inside the cloned git "Terminal" folder. If you are prompted by Windows, enable "Developer mode".

You will be prompted to upgrade the environment. In this dialog, choose to use Windows 10 SDK 10.0.18362.0, do _not_ upgrade the Platform Toolset to v142 (meaning; keep using v141) and click "OK", leaving all the remaining boxes ticked.

![]({{ site.baseurl }}/blog/assets/terminal/retarget.png)

When the solution has been fully loaded, choose the following configuration dropdown menu values "Release", "x64", "CascadiaPackage" and click the "Local Machine" button. This will build, deploy and launch the terminal.

![]({{ site.baseurl }}/blog/assets/terminal/configuration.png)

You will now find the "Windows Terminal (Dev Build)" in the Windows Start menu and you don't have to launch this from within VS2019 again.

### Configure Terminal

Once Terminal is running, hit `ctrl+t` to create a new tab. The menu icon will appear and you can enter the "Settings", which will launch a `profiles.json` file.

Here is how you can add WSL to the menu:

1. Create a new session in profiles, with content copied from profiles/cmd
1. Give it a new guid
1. Give it a new name, such as WSL
1. Specify its commandline to wsl.exe

```json
{
    "guid": "{09dc5eef-6840-4050-ae69-21e55e6a2e62}",
    "name": "WSL",
    "colorscheme": "Campbell",
    "historySize": 9001,
    "snapOnInput": true,
    "cursorColor": "#FFFFFF",
    "cursorShape": "bar",
    "commandline": "wsl.exe",
    "fontFace": "Consolas",
    "fontSize": 12,
    "acrylicOpacity": 0.75,
    "useAcrylic": true,
    "closeOnExit": false,
    "padding": "0, 0, 0, 0"
}
```

### Additional info

- [Windows Terminal reveal video](https://www.youtube.com/watch?v=8gw0rXPMMPE)
- [Windows Terminal: Building a better commandline experience...](https://www.youtube.com/watch?v=KMudkRcwjCw)
- [What's new with the Windows command line](https://www.youtube.com/watch?v=veqs2WVou9M)
- [The team behind the new Terminal](https://youtu.be/KMudkRcwjCw?t=3405) (YouTube link at 56:48)
- [GitHub repository](https://github.com/microsoft/Terminal)
- [Extensions for Terminal](https://twitter.com/richturn_ms/status/1126515079518703616)
- [Reveal video easter egg](https://twitter.com/PengwinLinux/status/1126929652382093318)