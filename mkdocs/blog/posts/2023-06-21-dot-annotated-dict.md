---
date: 2023-06-21
draft: false
tags:
  - python
---

# Access Python dict using dot annotation

Recently, an addition to Python 3.12 which seems like it'll be reverted in [cpython#105948](https://github.com/python/cpython/pull/105948), where an `AttrDict` hook could enable accessing a dict using dot annotation.

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

However, be warned of what happens when there is no key:

```python
>>> obj.baz
Traceback (most recent call last):
  File "<stdin>", line 1, in <module>
AttributeError: 'types.SimpleNamespace' object has no attribute 'baz'
```

Or, when you don't go far enough:

```python
>>> obj
namespace(foo=namespace(bar='val'))
```

Alternative libraries that might be worth checking out if you want more advanced behavior:

- https://github.com/cdgriffith/Box
- https://github.com/pawelzny/dotty_dict (no dependencies)
