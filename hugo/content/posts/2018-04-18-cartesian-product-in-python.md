---
ShowToc: false
TocOpen: false
date: 2018-04-18 02:00:12+02:00
draft: false
tags:
- python
title: Cartesian product in Python
---

A [cartesian product](https://en.wikipedia.org/wiki/Cartesian_product) operation can return a set of "combinations" based on given values, so for example I could have the following values:

| a   | b   | c   |
| --- | --- | --- |
| 1   | 1   | 1   |
| 2   | 2   |
|     | 3   |

I would then expect the cartesian product operation to return something like `a1b1c1, a1b1c2, a1b1c3, a1b2c1` and so on...

Many, many times have had to solve this problem over and over in Python... it's time to jot down some notes.

I'm taking the table above and making it into a dictionary:

```python
d = {'a': [1], 'b': [1, 2], 'c': [1, 2, 3]}
```

This can be fed into the following `cprod` function:

```python
import sys
import itertools

def cprod(dictionary):
    """Generate a list of dicts of cartesian product combinations."""

    if sys.version_info.major > 2:
        return (dict(zip(dictionary, x)) for x in itertools.product(*dictionary.values()))

    return (dict(itertools.izip(dictionary, x))
            for x in itertools.product(*dictionary.itervalues()))
```

...which will return a generator object, which can be looped over (once, before it gets destroyed):

```python
>>> for _dict in cprod(d):
>>>     print(_dict)
{'a': 1, 'b': 1, 'c': 1}
{'a': 1, 'b': 1, 'c': 2}
{'a': 1, 'b': 1, 'c': 3}
{'a': 1, 'b': 2, 'c': 1}
{'a': 1, 'b': 2, 'c': 2}
{'a': 1, 'b': 2, 'c': 3}
```

Number of combinations are `len(list(cprod(d)))` (in this case 6).

If we wish to get the list of combinations directly, we can use the `cprod_combos` function:

```python
def cprod_combos(dictionary):
    """Generate a list of cartesian product combinations."""
    combos = []
    for _dict in cprod(d):
        combo = ''
        for key, value in _dict.items():
            combo += '%s%s' % (key, value)
        combos.append(combo)
    return combos
```

...which will return a list of combinations:

```python
>>> [i for i in cprod_combos(d)]
['a1b1c1', 'a1b1c2', 'a1b1c3', 'a1b2c1', 'a1b2c2', 'a1b2c3']
```