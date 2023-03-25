---
date: 2022-12-17
draft: true
tags:
- monitoring
- python
title: Datadog
icon: simple/datadog
---

# Datadog

## Custom span

If there is no incoming request, a custom span can be created. This can be useful for e.g. cronjobs you wish to monitor with [:simple-datadog: Datadog](https://datadog.com).

!!! example "Custom span example"

    ```python
    --8<-- "mkdocs/static/datadog/custom_span.py"
    ```
