---
date: 2012-02-16
tags:
- vray
- python
- windows
- macos
---

# V-Ray for Maya

Running V-Ray for Maya in a production environment has its quirks. I have collected some bits and pieces off the [Chaos Group forum](http://www.chaosgroup.com/forums/vbulletin/showthread.php?51035-maintaining-Vray-builds-with-workstations-and-the-render-frarm/) and assembled it all here.

<!-- more -->

## Table of contents

1. Arbitrary location installation
2. Arbitrary location configuration
3. Python launch script and MEL sourcing
4. Launching the V-Ray render slave from the arbitrary location

## Arbitrary location installation

Choose a server location for all your builds and create a folder for the specific build you are installing. I am going to refer to this as the `[build folder]`.

Go find these three MEL files from the local installation of Maya (as of writing this, I am using Maya 2012):

- createMayaSoftwareCommonGlobalsTab.mel
- shouldAppearInNodeCreateUI.mel
- unifiedRenderGlobalsWindow.mel

...and copy them into `[build folder]/maya_root/scripts/others/`

Execute the V-Ray for Maya installer and when the target directory dialogue comes up, point the three paths to:

- `[build folder]/maya_root`
- `[build folder]/maya_vray`
- `[build folder]/vray`

When the install has finished, compress the whole build folder into a ZIP or RAR archive. Now completely uninstall V-Ray for Maya. This can be done through the start menu (Windows) or in this particularly quirky way:

```bat
"C:\Program Files\Chaos Group\V-Ray\Maya 2012 for x64/uninstall/wininstaller.exe" -uninstall="C:\Program Files\Chaos Group\V-Ray\Maya 2012 for x64/uninstall/install.log" -uninstallApp="V-Ray for Maya 2012 for x64" -gui=0 -quiet=1
```

When the uninstall is complete, delete anything that is left inside the server’s build folder and extract the backed up ZIP or RAR file right into it, restoring its contents.

There should be no trace left of V-Ray for Maya on your local machine, as this could cause conflicts when loading the V-Ray plug-in.

## Arbitrary location configuration

In order to successfully load Maya up with the V-Ray build available as a Maya plug-in, we will have to address three things.

First off, you will need to correctly set the following environment variables:

- `MAYA_RENDER_DESC_PATH`
- `VRAY_FOR_MAYA2012_MAIN_x64`
- `VRAY_FOR_MAYA2012_PLUGINS_x64`
- `VRAY_AUTH_CLIENT_FILE_PATH`
- `PATH`
- `MAYA_PLUG_IN_PATH`
- `MAYA_SCRIPT_PATH`
- `XBMLANGPATH`
- `VRAY_PATH`
- `VRAY_TOOLS_MAYA2012_x64`

For Mac OS X, you will have to set up the `DYLD_LIBRARY_PATH` environment variable and for Linux you will have to set up the `LD_LIBRARY_PATH` variable. More information on where these paths should go can be seen in the official documentation. Exactly how we set them up will be revealed in the Python launch script (down below).

The three MEL files, manually copied into the installation directory, was modified by the V-Ray installer, and in order to avoid nagging warning messages when opening up the Render Globals you will have to source them into Maya. Create a file called `source.mel` and place it inside the `[build folder]`, example below for Windows:

    source "X:/bin/vray/builds/maya2012_vray22001/maya_root/scripts/others/createMayaSoftwareCommonGlobalsTab.mel";
    source "X:/bin/vray/builds/maya2012_vray22001/maya_root/scripts/others/shouldAppearInNodeCreateUI.mel";
    source "X:/bin/vray/builds/maya2012_vray22001/maya_root/scripts/others/unifiedRenderGlobalsWindow.mel";

In order to make sure no licensing issues arise, place the license XML file `vrlclient.xml` on the server as well so that V-Ray will know how to find the floating licenses. The contents should look something like this (make sure to change the server IP address):

```xml
<VRLClient>
  <LicServer>
    <Host>10.0.0.1</Host>
    <Port>30304</Port>
    <Host1></Host1>
    <Port1>30304</Port1>
    <Host2></Host2>
    <Port2>30304</Port2>
    <!Proxy></!Proxy>
    <!ProxyPort>0</!ProxyPort>
    <User></User>
    <Pass></Pass>
  </LicServer>
</VRLClient>
```

## Python launch script

One way of dealing with all of the above would be to create a Python script which will launch Maya, having all parameters set. Just make sure Python is installed in your machine (or use the one version that comes bundled with Maya).

An example below, `maya2012_vray22001.py`, which will work for Windows and Mac OS X. I would be happy to add Linux support if anyone out there would like to contribute with this code, as I do not use Linux myself.

```python
import os
import sys

# Settings
windows = 'X:' # The mounted server share's drive letter
osx = '/Volumes/include' # The server's volume mount
buildsLocation = '/bin/vray/builds/' # The server path without drive letter or mounted volume/share.
buildFolderName = 'maya2012_vray22001' # The name of the build folder, containing the maya_root, maya_vray, vray folders and the source.mel file.
buildMayaVersion = '2012' # The version of Maya which the V-Ray plug-in has been compiled for.
licenseLocation = '/bin/vray/license/' # The location of the XML license file

# Detect OS and determine drive letter (win) or mount (mac)
volume = ''
if(sys.platform == 'win32'):
	volume = windows
elif(sys.platform == 'darwin'):
	volume = osx

# Set (and reset) environment variables which may not have already been set...
try:
	print(os.environ['MAYA_SCRIPT_PATH'])
except KeyError:
	os.environ['MAYA_SCRIPT_PATH'] = ''
try:
	print(os.environ['MAYA_PLUG_IN_PATH'])
except KeyError:
	os.environ['MAYA_PLUG_IN_PATH'] = ''

# Set environment variables
if(sys.platform == 'win32'):
	print('Setting Windows environment variables...')
	os.environ['MAYA_RENDER_DESC_PATH'] = volume + buildsLocation + buildFolderName + '/maya_root/bin/rendererDesc'
	os.environ['VRAY_FOR_MAYA' + buildMayaVersion + '_MAIN_x64'] = volume + buildsLocation + buildFolderName + '/maya_vray'
	os.environ['VRAY_FOR_MAYA' + buildMayaVersion + '_PLUGINS_x64'] = volume + buildsLocation + buildFolderName + '/maya_vray/vrayplugins'
	os.environ['VRAY_AUTH_CLIENT_FILE_PATH'] = volume + licenseLocation
	os.environ['PATH'] = os.environ['PATH'] + ';' + volume + buildsLocation + buildFolderName + '/maya_root/bin'
	os.environ['MAYA_PLUG_IN_PATH'] = os.environ['MAYA_PLUG_IN_PATH'] + ';' + volume + buildsLocation + buildFolderName + '/maya_vray/plug-ins'
	os.environ['MAYA_SCRIPT_PATH'] = os.environ['MAYA_SCRIPT_PATH'] + ';' + volume + buildsLocation + buildFolderName + '/maya_vray/scripts'
	os.environ['XBMLANGPATH'] = volume + buildsLocation + buildFolderName + '/maya_vray/icons/%B' + ';' + volume + buildsLocation + buildFolderName + '/maya_vray/icons/'
	os.environ['VRAY_PATH'] = volPipeline + vrayBaseLocation + buildName + '/maya_vray/bin'
	os.environ['VRAY_TOOLS_MAYA' + buildMayaVersion + '_x64'] = volPipeline + vrayBaseLocation + buildName + '/vray/bin'
elif(sys.platform == 'darwin'):
	print('Setting OS X environment variables...')
	os.environ['MAYA_RENDER_DESC_PATH'] = volume + buildsLocation + buildFolderName + '/maya_root/bin/rendererDesc'
	os.environ['VRAY_FOR_MAYA' + buildMayaVersion + '_MAIN_x64'] = volume + buildsLocation + buildFolderName + '/maya_vray'
	os.environ['VRAY_FOR_MAYA' + buildMayaVersion + '_PLUGINS_x64'] = volume + buildsLocation + buildFolderName + '/maya_vray/vrayplugins'
	os.environ['VRAY_AUTH_CLIENT_FILE_PATH'] = volume + '/bin/vray/license'
	os.environ['PATH'] = os.environ['PATH'] + ':' + volume + buildsLocation + buildFolderName + '/maya_root/bin'
	os.environ['PATH'] = os.environ['PATH'] + ':' + volume + buildsLocation + buildFolderName + '/maya_vray/bin'
	os.environ['PATH'] = os.environ['PATH'] + ':' + volume + buildsLocation + buildFolderName + '/vray/bin'
	os.environ['MAYA_PLUG_IN_PATH'] = os.environ['MAYA_PLUG_IN_PATH'] + ':' + volume + buildsLocation + buildFolderName + '/maya_vray/plug-ins'
	os.environ['MAYA_SCRIPT_PATH'] = os.environ['MAYA_SCRIPT_PATH'] + ':' +volume + buildsLocation + buildFolderName + '/maya_vray/scripts'
	os.environ['XBMLANGPATH'] = volume + buildsLocation + buildFolderName + '/maya_vray/icons/%B' + ':' + volume + buildsLocation + buildFolderName + '/maya_vray/icons/'
	os.environ['VRAY_PATH'] = volPipeline + vrayBaseLocation + buildName + '/maya_vray/bin'
	os.environ['VRAY_TOOLS_MAYA' + buildMayaVersion + '_x64'] = volPipeline + vrayBaseLocation + buildName + '/vray/bin'
	os.environ['DYLD_LIBRARY_PATH'] = volume + buildsLocation + buildFolderName + '/maya_root/MacOS/'

# Source the three MEL files
sourceStatement = "-script " + volume + buildsLocation + buildFolderName + "/source.mel"

# Set up OS-specific exec command
if(sys.platform == 'win32'):
	command = "\"C:/Program Files/Autodesk/Maya" + buildMayaVersion + "/bin/maya.exe\" " + sourceStatement
elif(sys.platform == 'darwin'):
	command = '/Applications/Autodesk/maya' + buildMayaVersion + '/Maya.app/Contents/bin/maya ' + sourceStatement

# Launch Maya
os.system(command)
```

Be careful with the fact that you might already have some of the environment variables set on your machine, which may cause conflicts and make V-Ray unable to load.

Also, you may want to add your own server paths to scripts or plug-ins directly into this launch script.

## Launching the V-Ray render slave from the arbitrary location

This Python script simply kills any already running processes of the V-Ray render slave and re-launches it. Still no Linux support here. I have called it `vray_slave_v22001.py`.

```python
import os, sys, subprocess

# Settings
windows = 'X:' # The server share's drive letter
osx = '/Volumes/include' # The server's volume mount
buildsLocation = '/bin/vray/builds/' # The server path without drive letter or mounted volume/share.
buildFolderName = 'maya2012_vray22001' # The name of the build folder, containing the maya_root, maya_vray, vray folders and the source.mel file.

# Determine drive letter or mount and then kill any running instances of the vray slave
volume = ''
if(sys.platform == 'win32'):
	volume = windows
	# Kill any already running instances of vray daemon
	command = "TASKKILL /F /IM \"vray.exe\""
	os.system(command)
	command = "TASKKILL /F /IM \"vrayspawner.exe\""
	os.system(command)
elif(sys.platform == 'darwin'):
	volume = osx
	output = ''
	proc=subprocess.Popen("ps -clx | grep 'vray.exe' | awk '{print $2}' | head -1", shell=True, stdout=subprocess.PIPE, ) # Find any running processes of vray with: ps aux | grep vray.exe | grep -v grep
	output=proc.communicate()[0]
	if (str(output) != ''):
		try:
			os.system( 'kill -9 ' +  output)
			print('Killed already running V-Ray Slave process with ID ' + str(output))
		except:
			print('Failed to kill already running V-Ray Slave process with ID ' + str(output))

# Launch slave
if(sys.platform == 'win32'):
	command = "C:\\Windows\\System32\\cmd.exe /C start " + volume + buildsLocation + buildFolderName + "\\maya_vray\\bin\\vray.exe -server -portNumber=20207"
elif(sys.platform == 'darwin'):
	os.environ['DYLD_LIBRARY_PATH'] = volume + buildsLocation + buildFolderName + '/maya_vray/lib'
	command = volume + buildsLocation + buildFolderName + '/maya_vray/bin/vray $* -portNumber=20207 -server'
	subprocess.Popen( command , shell=True, stdout=subprocess.PIPE)
```

## Quirks

In case you move the build folder, a few files may need editing. Check the file paths inside of these files (please note this is on OS X):

- `[build folder]/maya_vray/bin/vray`
- `[build folder]/maya_vray/bin/vrayslave`
- `[build folder]/maya_vray/bin/vrayconfig.xml`
- `[build folder]/vray/bin/initVRaySlaveDaemon`
- `[build folder]/vray/bin/registerVRaySlaveDaemon`
- `[build folder]/vray/bin/vrayslave`

Also, in case Maya crashes on startup in OS X, try adding an extra environment variable, `LC_ALL` and set it to "C":

```python
os.environ['LC_ALL'] = 'C'
```