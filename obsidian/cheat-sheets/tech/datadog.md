# 🐶 datadog

```python
from contextlib import contextmanager

import ddtrace

from logalot.lib.types import OptExcInfo


class DatadogSpan:
    """Offers a facility to get and create custom trace/span."""

    def __init__(
        self,
        name: str = "logalot.custom_trace",
        resource: str = "logalot",
    ):
        self.name = name
        self.resource = resource
        self.current_span = ddtrace.tracer.current_span()

    @contextmanager
    def span(self):
        """Yield the current span, or return a new custom span."""
        if self.current_span:
            yield self.current_span
        else:
            with ddtrace.tracer.trace(name=self.name, resource=self.resource) as span:
                yield span


def tag_span(span: ddtrace.span.Span, data: dict) -> dict[str, str]:
    """Tag span with metadata.

    Docs:
        https://ddtrace.readthedocs.io/en/stable/api.html#ddtrace.Span.set_tags
    """
    stringified_data = {key: f"{value}" for key, value in data.items()}
    span.set_tags(stringified_data)
    return stringified_data


def tag_span_with_exception_info(span: ddtrace.span.Span, exc_info: OptExcInfo) -> None:
    """Tag span with exception info.

    Docs:
        https://ddtrace.readthedocs.io/en/stable/api.html?#ddtrace.Span.set_exc_info
    """
    exc_type, exc_value, exc_traceback = exc_info
    span.set_exc_info(
        exc_type=exc_type,
        exc_val=exc_value,
        exc_tb=exc_traceback,
    )
```