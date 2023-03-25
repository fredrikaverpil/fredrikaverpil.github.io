---
date: 2013-06-07
tags:
- maya
---

# Maya Bonus Tools 2014 – arbitrary installation

Here’s how to install Maya Bonus Tools 2014 in an arbitrary location. I’ve covered two different methods (which can actually be mixed, if you will). The simple method (#1) is probably the best bet for most people while the “thorough” method (#2) could be interesting for some other folks. Take your pick.

<!-- more -->

### Method #1: Simple

Please note: This method comes from [Steven](http://area.autodesk.com/blogs/stevenr/bonustools). I haven’t tried this and instead I went with the next approach (“Method #2: Thorough”), and that’s working great for me.

1. Download the Bonus Tools 2014 installer [from here](https://apps.autodesk.com/MAYA/Detail/Index?id=appstore.exchange.autodesk.com:autodeskmayabonustools2014:en) and install it.
2. Copy the installed Bonus Tools folder onto a server location (and uninstall the local installation).
3. Add the server location path to the environment variable `MAYA_MODULE_PATH`.
4. Un-comment the last row inside of scripts/bonusToolsMenu.mel so that it says “bonusToolsMenu;”.

After having done the four steps above and when you launch Maya, anything in `MAYA_MODULE_PATH` will get traversed and automatically appended to `MAYA_SCRIPT_PATH`, `MAYA_PLUG_IN_PATH`, `XBMLANGPATH`, `PYTHONPATH` and `MAYA_PRESET_PATH`. This way the mel script “bonusToolsMenu.mel” will get sourced upon launch of Maya and the last row of code will initiate the script and draw the Bonus Tools menu.

### Method #2: Thorough

#### Getting the files and folders in place

1. Download the Bonus Tools 2014 installer [from here](https://apps.autodesk.com/MAYA/Detail/Index?id=appstore.exchange.autodesk.com:autodeskmayabonustools2014:en).
2. Extract the contents into a temporary directory using the command below this list of bullets.
3. Navigate down through the folders until you reach the scripts folder, python folder, plug-ins folder, icons folder etc. Copy these folders into a server location which is accessible from all workstations.
4. Rename one of the folders “win64”, “Linux” or “MacOS” to “plug-ins”, as these contain the plug-ins for each respective operating system.


#### Extracting the contents of the .msi file on Windows

    msiexec /a C:\MayaBonusTools2014-win64.msi /qb TARGETDIR=C:\BonusTools2014

Please note, you may have to launch the commandline prompt window with Administrator priviliges. If so, click the start menu, Accessories and right click the Command prompt and choose “Run as Administrator”.

#### Loading Bonus Tools 2014

1. Have environment variables `MAYA_SCRIPT_PATH` and `MAYA_PLUG_IN_PATH` point to the scripts and the plug-ins folders respectively. Optionally you can also have icons load by specifying the `XBMLANGPATH` environment variable.
2. Make sure to have the userSetup.mel or userSetup.py accessible in your `MAYA_SCRIPTS` path and have it include the code to initiate Bonus Tools 2014. Also make sure the userSetup file registers the path to the python folder inside the Bonus Tools 2014 installation.

#### Example use of userSetup.mel

```c
// userSetup.mel
// Register Bonus Tools Python Path
python("import sys");
python("sys.path.append('//10.0.1.100/share/stuff/bonustools/2014-x64/python')");

// Load Bonus Tools
bonusToolsMenu;
```

#### Example use of userSetup.py

```python
# userSetup.py
# Register Bonus Tools Python Path
import sys
sys.path.append('//10.0.1.100/share/stuff/bonustools/2014-x64/python')

# Load Bonus Tools
import maya.mel as mel
mel.eval('bonusToolsMenu')
```