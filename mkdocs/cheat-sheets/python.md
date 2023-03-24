---
date: 2022-12-17
draft: true
tags:
- python
title: Python
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

See [datadog](datadog.md) for some examples.

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

### Using hatchling + pip-tools to pin production dependencies

https://github.com/fredrikaverpil/hatch-playground

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

Like with ifs and switches, you only match against one case. "Optional remainders" can also be used:

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

Ordering matters when pattern matching on iterables, so it's probably not a great idea to do this, generally. You also cannot specify the "optional remainders" both before AND after an item in the list, so there is no way to identify one single item in a list reliably using pattern matching:

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