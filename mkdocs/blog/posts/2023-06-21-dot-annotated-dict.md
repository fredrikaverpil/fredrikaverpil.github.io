---
date: 2023-06-21
draft: false
tags:
  - python
---

# Access Python dict using dot annotation

Recently, an addition to Python 3.12 was reverted in [cpython#105948](https://github.com/python/cpython/pull/105948), where an `AttrDict` hook could enable accessing a dict using dot annotation.

```python
with open('kepler.json') as f:
    kepler = json.load(f, object_hook=AttrDict)
print(kepler.orbital_period.neptune)
```

However, as pointed out in a related [issue thread](https://github.com/python/cpython/issues/96145#issuecomment-1599508607), this is already possible using [`SimpleNamespace`](https://docs.python.org/3/library/types.html#types.SimpleNamespace):

```python
>>> import json
>>> from types import SimpleNamespace
>>> data = '{"foo": {"bar": "val"}}'
>>> obj = json.loads(data, object_hook=lambda x: SimpleNamespace(**x))
>>> obj.foo.bar
'val'
```
