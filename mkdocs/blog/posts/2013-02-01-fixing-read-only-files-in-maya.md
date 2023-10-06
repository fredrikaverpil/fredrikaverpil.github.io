---
date: 2013-02-01
authors:
  - fredrikaverpil
comments: true
tags:
- maya
- python
---

# Fixing "read-only" files in Maya

Recently we ended up with a Maya scene with locked, read-only, nodes which prohibited us from deleting them from the scene. This usually happens to nodes having been part of a locked reference which has been imported into the scene. In our case, we had no access to the original referenced file, so we had to unlock these nodes from being read-only.

<!-- more -->

For some reason it didn’t work to unlock them using the [lockNode](http://download.autodesk.com/global/docs/maya2012/en_us/CommandsPython/lockNode.html) command from within Maya and I had to write a Python script which would read each line of the Maya ASCII file and strip out any lines containing “lockNode”.

This is the script, which I’ve saved as unlock.py in the same directory as the scene file:

```python
import re

infile = open('original_file.ma', 'r')
outfile = open('new_file.ma', 'w')

for line in infile:
	lockNodeFound = False
	for word in re.findall(r"\w+", line):
		if (word == 'lockNode'):
			lockNodeFound = True
	if (lockNodeFound):
		#print 'Line removed: ' + line
		pass
	else:
		outfile.write(line)
```


Switch out the file name `original_file.ma` to match the name of your file. Back up your original Maya scene file somewhere safe and run the script in a command line window (not inside of Maya) like this:

    python unlock.py

Although a new file is created, I take no responsibility for how you use this script, of course. You’re on your own :)

If you want visual feedback of what’s going on, just uncomment the line that says #print 'Line removed: ' + line but this will make large scenes take significantly longer to process.

And, just in case, why not first try to see if you can unlock the nodes from within Maya with this simple script:


import maya.cmds as cmds

unlockError = False
nodes = cmds.ls()

```python
for node in nodes:
	lockStatus = cmds.lockNode( node, q=True )
	for response in lockStatus:
		if response != False:
			try:
				cmds.lockNode( node, lock=False )
				print 'Unlocked: ' + node
			except:
				print 'Error: Could not unlock ' + node
```