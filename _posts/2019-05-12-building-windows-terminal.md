---
layout: post
title: 'Building Windows Terminal'
tags: [windows]
---

![]({{ site.baseurl }}/blog/assets/terminal/terminal.png)

The new Windows Terminal (codenamed "Cascadia") was [revealed](https://www.youtube.com/watch?v=8gw0rXPMMPE) at this year's [Microsoft Build conference](https://www.microsoft.com/en-us/build) and quickly received a lot of attention, as it addresses the decades-old terminal experience in Windows.

This is me jotting down notes on how to get up and running with the new Terminal before it gets officially released in the Windows store, using Visual Studio 2019. This is all possible thanks to the fact that Microsoft is [open sourcing the new Terminal](https://github.com/microsoft/Terminal)!

<!--more-->

### Requirements

You will need Windows 10, version "1903". Get this by joining and enabling the Insiders Program. Set the "Release Preview" option and reboot. Then check for updates until you see the "1903" version downloading and installing.

You will also have to enable "Developer mode" in Windows.

### Installation steps

Install [Chocolatey](https://chocolatey.org/) using the instructions on their website. Then open a new and _elevated_ shell with (e.g. right-click cmd.exe and choose "Run as administrator").

Let's install git, unless you already have it:

```powershell
choco install -y git
```

Then install Visual Studio 2019 Community Edition, unless you already have Visual Studio 2019 installed:

```powershell
choco install -y -v visualstudio2019community
```

A number of workloads and components are required to build. Get them by using the Visual Studio Installer GUI or via this command:

```powershell
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

```powershell
cd Terminal
dep\nuget\nuget.exe restore OpenConsole.sln
tools\razzle.cmd
tools\bcz.cmd rel
```

Fun fact: Apparently, "razzle" and "bcz" are terms from within Microsoft, where "razzle" refers to a script which sets up your environment and "bcz" builds the project (with a clean prior to build).

This will produce `CascadiaPackage_0.0.1.0_x64.msix` inside of `src\cascadia\CascadiaPackage\AppPackages\ ...`.

![]({{ site.baseurl }}/blog/assets/terminal/msix-location.png)

However, you cannot install it, as it does not contain a valid certificate:

![]({{ site.baseurl }}/blog/assets/terminal/msix-nosign.png)

Below, I will explain what I have attempted (without luck). You can skip over the next paragraph if you just want to run the terminal.

### Side note: attempting to sign CascadiaPackage

I have tried to generate a .pfx certificate and then sign the .msix using the [MSIX Packaging tool](https://docs.microsoft.com/en-us/windows/msix/mpt-overview) (as well as with its bundled `signtool.exe` separately) but without any luck. The command I used to generate the certificate (elevated Powershell):

```powershell
New-SelfSignedCertificate -Type Custom -Subject "CN=Microsoft Corporation, O=Microsoft Corporation, L=Redmond, S=Washington, C=US" -KeyUsage DigitalSignature -FriendlyName "WindowsTerminalDev" -CertStoreLocation "Cert:\LocalMachine\My"
```

Then I exported (with private key and password) to .pfx via `certlm`. Finally I signed the CascadiaPackage with the MSIX Packaging Tool ("Edit package"), but I just ended up getting this error when attempting to install the CascadiaPackage after signing:

![]({{ site.baseurl }}/blog/assets/terminal/msix-signed.png)

The same issue is hit if signing with `signtool.exe' (bundled with the MSIX Package Tool):

```powershell
signtool.exe sign /a /v /fd SHA256 /f "C:\MyCodeSignCustom.pfx" /p "SuperSecurePassword" "CascadiaPackage_0.0.1.0_x64.appx"
```

I can't figure out how to actually install/deploy or run the built terminal here. Please see the next paragraph to run the terminal from within Visual Studio instead. If anyone knows how/if you can install/execute the terminal directly after building on the commandline with `bcz.cmd`, please let me know in the comments below!

### Running the terminal from VS2019

So, we can run the terminal from within Visual Studio instead (without a certificate).

Launch VS2019 and open the `OpenConsole.sln` solution inside the cloned git "Terminal" folder. If you are prompted by Windows, enable "Developer mode".

You will be prompted to upgrade the environment. In this dialog, choose to use Windows 10 SDK 10.0.18362.0, do _not_ upgrade the Platform Toolset to v142 (meaning; keep using v141) and click "OK", leaving all the remaining boxes ticked.

![]({{ site.baseurl }}/blog/assets/terminal/retarget.png)

When the solution has been fully loaded, choose the following configuration dropdown menu values "Release", "x64", "CascadiaPackage" and click the "Local Machine" button. This will build, deploy and launch the terminal.

![]({{ site.baseurl }}/blog/assets/terminal/configuration.png)

You will now find the "Windows Terminal (Dev Build)" in the Windows Start menu and you don't have to launch this from within VS2019 again.

![]({{ site.baseurl }}/blog/assets/terminal/start-menu.png)

### Configure Terminal

Once Terminal is running, hit `ctrl+t` to create a new tab. The menu icon will appear and you can enter the "Settings", which will launch a `profiles.json` file.

Here is how you can add WSL to the menu:

1. Create a new session in `profiles`, with content copied from `profiles/cmd`
1. Give it a new guid
1. Give it a new name, such as "WSL"
1. Specify its commandline to `wsl.exe`

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
- [A new Console for Windows - It's the open source Windows Terminal](https://www.hanselman.com/blog/ANewConsoleForWindowsItsTheOpenSourceWindowsTerminal.aspx)
- [GitHub repository](https://github.com/microsoft/Terminal)
- [Extensions for Terminal](https://twitter.com/richturn_ms/status/1126515079518703616)
- [Reveal video easter egg](https://twitter.com/PengwinLinux/status/1126929652382093318)