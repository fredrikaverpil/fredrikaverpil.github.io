---
layout: post
title: Querying the FPS preference in Maya
tags: [maya, python]
---

This tickles the funny bone.

```python
>>> import maya.mel as mel
>>> fps = mel.eval('float $fps = `currentTimeUnitToFPS`')
>>> print(fps)
24.0
```

<!--more-->

Let me know in the comments below if this can be improved...
