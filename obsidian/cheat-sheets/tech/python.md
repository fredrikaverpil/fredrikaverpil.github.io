---
title: 🐍 Python
tags: [python]
draft: false
summary: "Notes to self, snippets etc."

# PaperMod
ShowToc: false
TocOpen: true

updated: 2022-11-16T01:01:58+01:00
created: 2022-11-14T20:42:48+01:00
---

## Exceptions

### Traceback type aliases

```python
from types import TracebackType
from typing import TypeAlias, Union

ExcInfo: TypeAlias = tuple[type[BaseException], BaseException, TracebackType]
OptExcInfo: TypeAlias = Union[ExcInfo, tuple[None, None, None]]
```

### Was something raised?

```python
def exception_was_raised():
	return None not in sys.exc_info()
```

### Exception hooks

See [[datadog]] for some examples.