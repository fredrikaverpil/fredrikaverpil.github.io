---
layout: post
title: Send MEL/Python commands from Sublime Text to Maya
tags: [python, pyside, maya]
---

This is a step by step to set you up with sending selected MEL/Python code snippets or whole scripts from [Sublime Text](http://www.sublimetext.com) to Maya using Justin Israel’s [MayaSublime](https://github.com/justinfx/MayaSublime) package.

<!--more-->

1. If you don’t already have it installed, install [Package Control](https://packagecontrol.io/installation) for Sublime Text by executing a bunch of code inside of Sublime Text’s console. You will find this code on the Package Control installation page.
2. In Sublime Text, hit Ctrl+Shift+P (Cmd+Shift+P on Mac), then choose (or type) “Package Control: Install Package”.
3. A list of available packages will download and from this list choose (or type) “MayaSublime”. Let it download and install automatically.
4. Open Maya and in the script editor (Python tab), enter and execute the following code:


```python
import maya.cmds as cmds

# Close ports if they were already open under another configuration
try:
    cmds.commandPort(name=":7001", close=True)
except:
    cmds.warning('Could not close port 7001 (maybe it is not opened yet...)')
try:
    cmds.commandPort(name=":7002", close=True)
except:
    cmds.warning('Could not close port 7002 (maybe it is not opened yet...)')

# Open new ports
cmds.commandPort(name=":7001", sourceType="mel")
cmds.commandPort(name=":7002", sourceType="python")
```


This will create port connections on which Maya will listen for commands. Sublime Text now has MayaSublime installed and which is already setup (by default) to communicate MEL code on port 7001 and Python code on port 7002. This means we are done setting it all up.

To verify that it works, create a simple Python script in Sublime Text and save it somewhere so that Sublime Text knows it’s Python code. You can also just hit Ctrl+Shift+P (Cmd+Shift+P on Mac) and type in “Set Syntax: Python”). Then, in Sublime Text, select the code (or portions of it) and then hit Ctrl+Enter to send it off into Maya. The code should now have been executed inside of Maya.

To execute the whole script in its location (making global variables such as `__file__` work), make sure the script has been saved, nothing is selected and hit Ctrl+Enter.

Next time you start Maya you will have to execute that same Python code (above) to enable communication between Sublime Text and Maya, so a good idea is to make it into a shelf button.

Huge thanks goes out to [Justin Israel](http://justinfx.com) who provided the package.



### More on packages

Another useful package is [Sublimerge](http://www.sublimerge.com), which helps you compare code from different versions of the same script, much like how [WinMerge](http://winmerge.org) works. Please note that Sublimerge will [change some key bindnings](http://www.sublimerge.com/docs/configuration.html#default-key-bindings) which may have to disable, depending on which keyboard layout you are using.

For PySide auto-completion, check out the [PySide package](https://github.com/DamnWidget/SublimePySide). Should work for both Sublime Text 2 and Sublime Text 3.

Install any packages from within Sublime Text, like explained at the beginning of this post (step 1 and 2).

### Key mappings

I accidentally hit shift + delete a lot with nothing selected, which performs a cut (erases my copy/paste memory), and so I have opted to disable this keymapping:

```python
{ "keys": ["shift+delete"], "command": "noop" }
```
