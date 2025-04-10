---
date: 2022-12-17
draft: false
tags:
  - python
title: Python
icon: simple/python
---

# Python

## Exceptions

### Custom exception

```python
class MyError(Exception):
    def __init__(self, message: str | None = None) -> None:
        self.message = message
        super().__init__(self.message)
```

### Was something raised?

```python
def exception_was_raised():
    return None not in sys.exc_info()
```

### Exception hooks

You can register
[custom exception hooks](https://docs.python.org/3/library/sys.html#sys.excepthook),
which can run code whenever an exception is raised, and just before the script
will exit.

```python
import sys

def custom_excepthook():
    print("oops")

sys.excepthook = custom_excepthook
sys.unraisablehook = custom_excepthook
```

## Caching with TTL

```python
import time
from functools import lru_cache
import random


@lru_cache(maxsize=2)
def my_func(ttl_hash: int | None = None) -> int:
	return random.randint(0, 10)


def get_ttl_hash(seconds: int = 60) -> int:
	return round(time.time() / seconds)


my_func(ttl_hash=get_ttl_hash())  # cached result, using ttl
```

## Typing

https://dev.to/chadrik/the-missing-guide-to-python-static-typing-532i

### Traceback type aliases

```python
from types import TracebackType
from typing import TypeAlias, Union

ExcInfo: TypeAlias = tuple[type[BaseException], BaseException, TracebackType]
OptExcInfo: TypeAlias = Union[ExcInfo, tuple[None, None, None]]
```

```python
from typing import TypeAlias

JsonType: TypeAlias = (
	None | bool | int | float | str | list["JsonType"] | dict[str, "JsonType"]
)
JsonDict: TypeAlias = dict[str, "JsonType"]
```

## Pyproject.toml

### Using uv to manage dependencies

There are many tools for defining a Python project and manage dependencies in
Python; [setuptools](https://github.com/pypa/setuptools),
[pip-tools](https://github.com/jazzband/pip-tools),
[poetry](https://github.com/python-poetry/poetry),
[pdm](https://github.com/pdm-project/pdm),
[hatch](https://github.com/pypa/hatch),
[rye](https://github.com/astral-sh/rye)... but not until 2024, was there finally
a tool that supplied all the necessary tooling under one umbrella for managing
both applications and libraries, prod vs dev dependency lockfile, optional
extras etc: [uv](https://docs.astral.sh/uv/).

## Dates and times

### Timedelta gotcha

When using timedelta, DST is not taken into consideration. Workaround:

```python
from datetime import datetime, timedelta
from pytz import timezone as pytz_timezone

def normalize_timedelta(dt: datetime, delta: timedelta) -> datetime:
	"""Normalize timedelta operation to fix DST."""

	if hasattr(dt.tzinfo, "zone"):
		timezone_ = dt.tzinfo.zone
		return pytz_timezone(timezone_).normalize(dt + delta)

	else:
		raise Exception("Datetime object does not include a pytz timezone")

now = datetime.now(pytz_timezone("Europe/Stockholm"))
tomorrow = normalize_timedelta(dt=now, delta=timedelta(days=1))
yesterday = normalize_timedelta(dt=now, delta=-timedelta(days=1))
```

### Testing in different timezones

You may want to run your tests in both UTC and e.g. your local timezone:

```bash
TZ=Europe/Stockholm pytest
TZ=UTC pytest
```

## Pattern matching gotchas

Structural pattern matching was introduced in Python 3.10.

Like with ifs and switches, you only match against one case. "Optional
remainders" can also be used:

```python
my_dict = {"fruit": "apple", "digit": 1}

match my_dict:
    case {"fruit": "apple", **remainder_kwargs}:
        print("Found apple in match 1")  # will be found
    case {"digit": 1, **remainder_kwargs}:
	    print("Found digit in match 1")  # will not be found as we already found the apple.

match my_dict:
    case {"digit": 1, **remainder_kwargs}:
        print("Found digit in match 2")  # will be found
    case {"fruit": "apple", **remainder_kwargs}:
        print("Found apple in match 2")  # will not be found as we already found the digit.

```

### Don't pattern-match on iterables

Ordering matters when pattern matching on iterables, so it's probably not a
great idea to do this, generally. You also cannot specify the "optional
remainders" both before AND after an item in the list, so there is no way to
identify one single item in a list reliably using pattern matching:

```python
my_list = ["apple", "orange"]

match my_list:
    case ["apple", _]:
        print("Found the apple first in the list")  # will find the apple

match my_list:
    case ["orange", _]:
        print("Found the orange first in the list")  # will NOT be found

match my_list:
    case [x, "apple", y]:
        print("Found the apple with non-optionals") # will NOT find the apple
    case [x, "orange", *y]:
        print("Found the orange with multiple optionals") # will find the orange

```

```python
match my_list:
    case [*x, "orange", *y]:  # syntax error
        print("Found the orange with multiple optionals before AND after")
```

