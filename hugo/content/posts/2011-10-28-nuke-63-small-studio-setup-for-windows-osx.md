---
ShowToc: false
TocOpen: false
date: 2011-10-28 21:58:04+02:00
draft: false
tags:
- nuke
- python
title: Nuke 6.3 small studio setup for Windows/OS X
---

This is a quick guide to setting Nuke 6.3 up with a custom menu and make it work more seamlessly across operating systems.



Since I jump between my Mac laptop and a Windows based workstation, I used to mess around with broken file read nodes in The Foundry’s Nuke and other stuff related to file paths being different on Windows and Mac OS X. Well... no more!

In this article, I will explain how to set up Nuke with a few handy extras that will make it fly even smoother across Windows and Mac computers, and with custom menus. Let me know if there are any specific Windows/Mac switching woes that you are experiencing, which are not covered here.

Downloads:

- [init.py](https://raw.github.com/fredrikaverpil/nuke/master/init.py)
- [menu.py](https://raw.github.com/fredrikaverpil/nuke/master/menu.py)
- [launch_nuke.py](https://raw.github.com/fredrikaverpil/nuke/master/launchers/launch_nuke.py)
- [launch_nuke.sh](https://raw.github.com/fredrikaverpil/nuke/master/launchers/launch_nuke.sh)

## The NUKE_PATH environment variable

First off, create the environment variable `NUKE_PATH` and point it to a centralized server location, in my case `Y:/include/nuke` on Windows or `/Volumes/Assets/include/nuke` on Mac OS X. Nuke will then look for this environment variable and intialize any customizations defined with `init.py` and `menu.py`, which we will create inside this directory.

An alternative to setting the environment variable is to run Nuke from a Python script or a bash script for the Mac, which will launch Nuke with the environment variable loaded for the particular shell it is running in. Examples of this below...

Contents of `launch_nuke.py` (works on Windows with Python installed):

```python
import os
os.environ['NUKE_PATH'] = 'Y:/include/nuke'
os.startfile('C:/Program Files/Nuke6.3v5/Nuke6.3v5.exe')
```

Contents of `launch_nuke.sh` (works on Mac OS X):

```bash
#!/bin/bash
export NUKE_PATH=/Volumes/Assets/include/nuke
/Applications/Nuke6.3v5/Nuke6.3v5.app/Nuke6.3v5
```

## The files and folder structure

Inside the `/include/nuke` folder, we will create the following files and folder structure:

    gizmos/
    icons/
    plugins/
    scripts/
    init.py
    menu.py

Just create `init.py` and `menu.py` as empty text files.

The `init.py` will serve as a script that will initialize Nuke with custom functions.
The `menu.py` will solely handle the interactive bits of our customization, such as the custom menus in the GUI.
Each folder will soon be acknowledged by Nuke to contain loadable things (scripts, gizmos etc), but I like to keep it organized and nice, thus separating icons from plugins from scripts.

## Make Read nodes load up without errors on both Windows and OS X

By specifying the file server paths for both Windows and Mac OS X in `init.py`, we will be able to filter these paths out and replace them with the appropriate ones, for the current operating system, without forcing the user to manually changing the file paths:

Add to `init.py`:

```python
# Make all filepaths load without errors regardless of OS (No Linux support and no C: support)
def myFilenameFilter(filename):
	if nuke.env['MACOS']:
		filename = filename.replace( 'X:', '/Volumes/Projects' )
		filename = filename.replace( 'Y:', '/Volumes/Assets' )
	if nuke.env['WIN32']:
		filename = filename.replace( '/Volumes/Projects', 'X:' )
		filename = filename.replace( '/Volumes/Assets', 'Y:' )

	return filename


# Use the filenameFilter(s)
nuke.addFilenameFilter(myFilenameFilter)
```

You can read more about the filenameFilter in [The Foundry’s Python docs](http://docs.thefoundry.co.uk/nuke/63/pythondevguide/callbacks.html#filenamefilter). If you wish to specify file paths for Linux systems, simply check out `nuke.env['LINUX']`, like in their sample code on that page. For some reason, their sample code snipplet does not work with Nuke 6.3 as of writing this.

## Generate OS-specific variable values

In order to make a whole lot of custom stuff work, we need to detect the current operating system. This time around I have taken a bit different route when detecting the OS type, by using the Python “sys” library. It’s up to you which flavour you like best.

```python
import sys

# Create OS specific variables (no Linux support)
volProjects = ''
volAssets = ''
if(sys.platform == 'win32'):
	volProjects = 'X:'
	volAssets = 'Y:'
elif(sys.platform == 'darwin'):
	volProjects = '/Volumes/Projects'
	volAssets = '/Volumes/Assets'
```

Please note that the variables created here will be required by the examples below as well. And again, I have not bothered to add Linux support. Just google it. :)

## Favorites, cross-platform

Add to `init.py`:

```python
# Make these favorites show up in Nuke
nuke.addFavoriteDir('File server', volProjects + '/Projects/')
nuke.addFavoriteDir('Assets', volAssets)
nuke.addFavoriteDir('R&D', volProjects + '/RnD/')
```

## Add more formats

Add to `init.py`:

```python
# Formats
nuke.addFormat( '1024 576 PAL Widescreen' )
nuke.addFormat( '1280 720 HD 720p' )
```

## Define custom location for gizmos and plugins, cross-platform

In order to easily load gizmos and plugins from within our own custom menus, I will just add a bit of extra code to init.py:

```python
# Set plugin/gizmo sub-folders
nuke.pluginAppendPath(volAssets + '/include/nuke/gizmos')
nuke.pluginAppendPath(volAssets + '/include/nuke/plugins')
nuke.pluginAppendPath(volAssets + '/include/nuke/scripts')
nuke.pluginAppendPath(volAssets + '/include/nuke/icons')
```

## Load OS-specific plugins, cross-platform

In this example, I am loading [Peregrine Lab’s Bokeh](http://www.peregrinelabs.com/bokeh), and since it is built for Windows and Mac OS X respectively, I need to specify which OS build I will try and load up since I will not be able to load the Windows plugin on Mac OS X or vice versa. On Mac OS X, you can use [Pacifist](http://www.charlessoft.com/) to extract the files from the .dmg file to be put on the server location.

Assuming you have already installed the floating server licenses onto 10.0.1.100 and the RLM server is running on port 5053, add to `init.py`:

```python
# Load Bokeh
os.environ['RLM_LICENSE'] = '5053@10.0.1.100'
if nuke.env['WIN32']:
	currentBokeh = 'Bokeh-Nuke6.3-1.2.3-win64'
if nuke.env['MACOS']:
	currentBokeh = 'Bokeh-Nuke6.3-1.2.3-Darwin'
nuke.pluginAppendPath(volAssets + '/include/nuke/plugins/bokeh/' + currentBokeh + '/')
nuke.load("pgBokeh")
```

Please note that on Windows, you may have to copy the .dll files from the Bokeh build into the root of the build’s folder.

## Load OS-specific OFX plugins, cross-platform

In this example, I’m loading [Frischluft’s Lenscare](http://www.frischluft.com/lenscare/index.php). Unlike “regular” plugins, OFX plugins are identified by an environment variable. If it’s not already set, we will set it. If it had been already set, any existing content to the environment variable is preserved (any local installations of OFX plugins will be maintained).

```python
# Check wheter OFX_PLUGIN_PATH has been set or not
try:
	os.environ['OFX_PLUGIN_PATH'] += ';'
except:
	os.environ['OFX_PLUGIN_PATH'] = ''

# Load Frischluft Lenscare
if(sys.platform == 'win32'):
	os.environ['OFX_PLUGIN_PATH'] += volAssets + '/bin/lenscare/lenscare_ofx_v1.44_win'
elif(sys.platform == 'darwin'):
	os.environ['OFX_PLUGIN_PATH'] += volAssets + '/bin/lenscare/lenscare_ofx_v1.44_osx'
```

## Creating a custom toolbar menu and file menu

I am using menu.py to set this up, and I will create a file menu as well as a toolbar menu. The toolbar menu tools will show up when you hit “tab” in the node graph and start typing its name, and so, I would like these to be tools that generate a node. The file menu will only include the tools that launch a menu or perform an action that does not generate a new node in the node graph.

Enter the following into `menu.py`:

```python
# -------- My Toolbar --------

# Initialize the toolbar menu
toolbar = nuke.toolbar('Nodes')

# My tools
toolbar.addCommand('My Nodes/Bezier', "nuke.createNode('Bezier')")
toolbar.addCommand( "My Nodes/pgBokeh", "nuke.createNode('pgBokeh')")


# -------- My File Menu --------

nuke.menu( 'Nuke' ).addCommand( 'My file menu/Rendering/Send to RenderManager', "nuke.load(\"submitNukeToRenderManager\"), submitNukeToRenderManager()" )
```

Just fill those menus up with your heart’s content, by editing the code above to load up any downloaded gizmo or plugin (placed in their respective centralized folders within your `NUKE_PATH`).

## Other tweaks

The following tweaks are not cross-platform related. Just some things that might come in handy…

## If write dir does not exist, make Nuke create it automatically

This is quite an annoying feature, I think, so I force Nuke to create the destination path for write nodes (if they do not already exist).

Add to `init.py`:

```python
import nuke, os

# If Write dir does not exist, create it
def createWriteDir():
    file = nuke.filename(nuke.thisNode())
    dir = os.path.dirname( file )
    osdir = nuke.callbacks.filenameFilter( dir )
    try:
        os.makedirs( osdir )
        return
    except:
        return

# Activate the createWriteDir function
nuke.addBeforeRender( createWriteDir )
```

Make Quicktime Write nodes default to sRGB

In Mac OS X prior to 10.6, the default system gamma value was 1.8. Today all new Macs are shipped running Mac OS X in gamma 2.2. Therefore, leaving Nuke’s default value of gamma 1.8 is mostly a legacy thing. In order to change this default into sRGB, enter the following into `init.py`:

```python
# Make Write node default to sRGB color space
nuke.knobDefault('Write.mov.colorspace', 'sRGB')
```