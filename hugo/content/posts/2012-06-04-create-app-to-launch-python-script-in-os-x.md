---
ShowToc: false
TocOpen: false
date: 2012-06-04 02:00:12+02:00
draft: false
tags:
- applescript
- python
- osx
title: Create .app to launch Python script in Mac OS X
---

Quick step-by-step to create clickable .app to launch your python scripts.



Launch the ApplecScript Editor (located in /Applications/Utilities/) and type in the following into the editor:

```applescript
tell application "Terminal"
	do script with command "python /path/to/your/script.py"
end tell
```

And simply hit save. Choose to save as application.

If anyone reading this know how to get rid of the Terminal window that opens when you launch the .app, let me know!

If you want to get more advanced, check out [Platypus](http://sveinbjorn.org/platypus).