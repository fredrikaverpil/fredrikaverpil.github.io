---
date: 2017-11-16
authors:
  - fredrikaverpil
comments: true
tags:
- maya
- python
---

# Maya 2018 file size bug

A very [annoying bug](https://forums.autodesk.com/t5/maya-forum/maya-2018-error-opening-large-mb-files-over-2gb/td-p/7466095) came to light in mid-October which makes Maya 2018 binary scene files (*.mb) unreadable if they are larger than ~2 GB in size.

To be exact, the bug is hit when the file size is larger than [2147483647](https://en.wikipedia.org/wiki/2,147,483,647) bytes, the maximum positive value for a 32-bit signed binary integer.

<!-- more -->

A workaround is to instead save your scenes in Maya ASCII format (*.ma), since they are not affected by this bug. Or perhaps restructure your binary scenes into using references, so that the data is split up among several scene files and won't reach the problematic file size.

Hopefully, we'll see a Maya 2018.2 update soon with this issue patched. But in the meanwhile, what can one do to avoid hitting this issue?

Using Python, we can check for a file's size in bytes like this:

```python
import os

size_bytes = os.path.getsize(filepath)
if size_bytes > 2147483647:
    print('scene file is larger than ~2GB')
```

So ideally, we'd like to have this run after a file save and let us know if we are getting close to 2 GB as well as if we exceeded this limit. Using a callback function, we can run custom code after each save command in Maya. Maya's callbacks are actually called ["scriptJobs"](http://help.autodesk.com/view/MAYAUL/2018/ENU/?guid=GUID-A42F2A04-0216-408D-8073-F4D4D896CE8D).

Here's a function which will check your Maya binary (*.mb) scene's size and warn you if you hit the bug or if you're getting close (at ~1.5 GB):

```python
import os

import maya.cmds as cmds


def check_2gb_bug():
    """Check for the 2GB scene file bug"""

    filepath = cmds.file(query=True, sceneName=True)
    maya_version = cmds.about(version=True)
    if '2018' in maya_version and filepath.endswith('.mb'):
        size_bytes = os.path.getsize(filepath)
        if size_bytes > 2147483647:
            message = ('You have hit the ~2GB bug, re-save your scene in '
                       'ASCII format (*.ma) right now to avoid loosing data!')
            cmds.confirmDialog(icon='critical', message=message)
        elif size_bytes > 1610612736:
            message = ('Your scene size is getting close to the ~2GB bug!')
            cmds.confirmDialog(icon='warning', message=message)

        else:
            print('Scene file size: ' + str(size_bytes) + ' bytes')

```

You can register this function as a Maya script job which gets triggered after the scene is saved:

```python
cmds.scriptJob(event=['SceneSaved', 'check_2gb_bug()'])
```

If you now save a binary scene file in Maya 2018, you should see the file size (in bytes) printed in the script editor. If you would exceed (or get close) to the problematic file size, a blocking dialog window will appear instead.

If you add both the function and the scriptJob command to e.g. your [userSetup.py](http://help.autodesk.com/view/MAYAUL/2018/ENU/?guid=GUID-C0F27A50-3DD6-454C-A4D1-9E3C44B3C990), you should have this script job active each time you start Maya, without having to re-enter the code above.