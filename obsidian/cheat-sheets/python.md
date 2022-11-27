---
title: 🐍 Python
tags: [python]
draft: false
summary: "Notes to self, snippets etc."

# PaperMod
ShowToc: true
TocOpen: true

date: 2022-11-27T14:17:09+01:00
---

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

See [[datadog]] for some examples.

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

