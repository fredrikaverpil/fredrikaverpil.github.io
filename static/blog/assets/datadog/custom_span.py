import sys
import time
from contextlib import contextmanager
from types import TracebackType
from typing import TypeAlias

import ddtrace
from ddtrace import tracer

ExcInfo: TypeAlias = tuple[type[BaseException], BaseException, TracebackType]
OptExcInfo: TypeAlias = ExcInfo | tuple[None, None, None]


class DatadogSpan:
    """Offers a facility to get and create custom trace/span."""

    def __init__(
        self,
        name: str = "custom_trace",
        resource: str = "custom_resource",
    ) -> None:
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


def configure_excepthooks():
    sys.excepthook = custom_excepthook
    sys.unraisablehook = custom_excepthook


def custom_excepthook(exc_type, exc_value, exc_traceback):
    error = (exc_type, exc_value, exc_traceback)  # could also use sys.exc_info()
    with DatadogSpan().span() as span:
        tag_span_with_exception_info(span=span, exc_info=error)


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


def run_logic():
    """This is where you run your logic, such as a cronjob."""
    time.sleep(10)


@tracer.wrap(
    service="example-service",
    resource="example-resource",
)
def main():
    configure_excepthooks()  # report APM error on unexpected exception
    run_logic()


