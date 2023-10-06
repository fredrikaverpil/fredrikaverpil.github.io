---
date: 2017-11-17
authors:
  - fredrikaverpil
comments: true
tags:
- maya
- python
---

# Querying the FPS preference in Maya

This tickles the funny bone.

```python
>>> import maya.mel as mel
>>> fps = mel.eval('float $fps = `currentTimeUnitToFPS`')
>>> print(fps)
24.0
```

Let me know in the comments below if this can be improved...