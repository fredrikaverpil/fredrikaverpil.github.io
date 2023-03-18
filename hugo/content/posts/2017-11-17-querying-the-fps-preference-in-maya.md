---
ShowToc: false
TocOpen: false
date: 2017-11-17 02:00:12+01:00
draft: false
tags:
- maya
- python
title: Querying the FPS preference in Maya
---

This tickles the funny bone.

```python
>>> import maya.mel as mel
>>> fps = mel.eval('float $fps = `currentTimeUnitToFPS`')
>>> print(fps)
24.0
```

Let me know in the comments below if this can be improved...