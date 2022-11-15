---
title: Querying the FPS preference in Maya
tags: [maya, python]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

created: 2017-11-17T02:00:12+01:00
updated: 2022-11-15T17:29:41+01:00
---

This tickles the funny bone.

```python
>>> import maya.mel as mel
>>> fps = mel.eval('float $fps = `currentTimeUnitToFPS`')
>>> print(fps)
24.0
```

Let me know in the comments below if this can be improved...
