---
date: 2013-12-13
authors:
  - fredrikaverpil
comments: true
tags:
- maya
- python
---

# Load latest Maya crash file

This simple python script looks for the latest file with .ma file extension in the path given and prompts the user whether to load it or not. Just verify that the searchDir variable is pointing to your temp dir.

<!-- more -->

```python
import os, glob, time
import maya.cmds as cmds

def crashFileLoader():
    searchDir = os.environ['TEMP'] + '/'
    files = filter(os.path.isfile, glob.glob(searchDir + '*.ma'))
    files.sort(key=lambda x: os.path.getmtime(x))
    latestCrashFile = (str(files[len(files)-1]))
    timeStamp = "%s" % time.ctime(os.path.getctime(latestCrashFile))

    messageString = 'Are you sure you want to open ' + latestCrashFile + ' created on ' + timeStamp + '?'
    retVal = cmds.confirmDialog( title='Confirm', message=messageString, button=['Yes','No'], defaultButton='Yes', cancelButton='No', dismissString='No' )

    if retVal == 'Yes':
        cmds.file(latestCrashFile, force=True, open=True)
```

Launch the script with:

```python
crashFileLoader()
```