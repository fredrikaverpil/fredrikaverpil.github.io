---
date: 2022-12-17
draft: false
tags:
- monitoring
- python
---

# Datadog and custom tracing

When using [:simple-datadog: Datadog](https://datadog.com) for monitoring, Datadog will only record a trace if there is an incoming request.

If there is no incoming request, such as if a cronjob is running, then Datadog will not record a trace and you might not be alerted by an error.

To solve this, you can use the [ddtrace-py](https://github.com/DataDog/dd-trace-py) library and create a custom trace/span whenever tag and tag a span with exception information. This will effectively make it possible to track errors in APM error tracking.

<!-- more -->

!!! example "Custom trace/span example"

    ```python
    --8<-- "mkdocs/static/datadog/custom_span.py"
    ```
