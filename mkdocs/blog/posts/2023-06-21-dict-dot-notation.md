---
date: 2023-06-21
draft: false
tags:
  - python
---

# Access Python dict using dot notation

Recently, an addition to Python 3.12 might be reverted in [cpython#105948](https://github.com/python/cpython/pull/105948), where a new `json.AttrDict` class could enable accessing a dict using dot notation using a `json.load` hook.

But as pointed out in a related [issue thread](https://github.com/python/cpython/issues/96145#issuecomment-1599508607), this is already possible using the standard library's [`SimpleNamespace`](https://docs.python.org/3/library/types.html#types.SimpleNamespace).

<!-- more -->

This was the proposed use case from the [original PR](https://github.com/python/cpython/pull/96146) which introduced the `AttrDict` class:

```python
import json


class AttrDict(dict):
    """Dict like object that supports attribute style dotted access.
    This class is intended for use with the *object_hook* in json.loads():
        >>> from json import loads, AttrDict
        >>> json_string = '{"mercury": 88, "venus": 225, "earth": 365, "mars": 687}'
        >>> orbital_period = loads(json_string, object_hook=AttrDict)
        >>> orbital_period['earth']     # Dict style lookup
        365
        >>> orbital_period.earth        # Attribute style lookup
        365
        >>> orbital_period.keys()       # All dict methods are present
        dict_keys(['mercury', 'venus', 'earth', 'mars'])
    Attribute style access only works for keys that are valid attribute names.
    In contrast, dictionary style access works for all keys.
    For example, ``d.two words`` contains a space and is not syntactically
    valid Python, so ``d["two words"]`` should be used instead.
    If a key has the same name as dictionary method, then a dictionary
    lookup finds the key and an attribute lookup finds the method:
        >>> d = AttrDict(items=50)
        >>> d['items']                  # Lookup the key
        50
        >>> d.items()                   # Call the method
        dict_items([('items', 50)])
    """
    __slots__ = ()

    def __getattr__(self, attr):
        try:
            return self[attr]
        except KeyError:
            raise AttributeError(attr) from None

    def __setattr__(self, attr, value):
        self[attr] = value

    def __delattr__(self, attr):
        try:
            del self[attr]
        except KeyError:
            raise AttributeError(attr) from None

    def __dir__(self):
        return list(self) + dir(type(self))


with open('kepler.json') as f:
    kepler = json.load(f, object_hook=AttrDict)
print(kepler.orbital_period.neptune)
```

And this is how you can already achieve it:

```python
>>> import json
>>> from types import SimpleNamespace
>>> data = '{"foo": {"bar": "val"}}'
>>> obj = json.loads(data, object_hook=lambda x: SimpleNamespace(**x))
>>> obj.foo.bar
'val'
```

!!! warning

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

!!! tip "Alternative solutions"

    Alternative libraries that might be worth checking out if you want more advanced behavior:

    - https://github.com/cdgriffith/Box
    - https://github.com/pawelzny/dotty_dict (no dependencies)
    - https://github.com/makinacorpus/easydict
    - https://pypi.org/project/attrdict/ (deprecated/archived, don't use this)
