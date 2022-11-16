---
title: Stop using print and do some logging
tags: [python]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

created: 2014-11-10T01:00:12+01:00
updated: 2022-11-15T22:29:16+01:00
---

This will print to stdout, similar to a regular print, but it will also log to file.



```python
import logging

# Logging setup
LOG_FILEPATH = '/path/to/log_file.log'
logger = logging.getLogger('My logger')
logger.setLevel(logging.INFO)
formatter = logging.Formatter('%(asctime)s %(levelname)s %(message)s')

# Logging to file
file_handler = logging.FileHandler( LOG_FILEPATH )
file_handler.setFormatter(formatter)
file_handler.setLevel(logging.INFO)
logger.addHandler(file_handler)

# Logging to stdout
stdout_handler = logging.StreamHandler(sys.stdout)
stdout_handler.setFormatter(formatter)
stdout_handler.setLevel(logging.INFO)
logger.addHandler(stdout_handler)

# Usage
logger.info('Hello')       # Log infos
logger.warning('Oops')     # Log warnings
logger.error('Dang!')      # Log errors
```

Another way of configuring logging can be seen below, which makes for easier reading if using multiple handlers. In this case I am only logging WARNING levels and above to file but printing INFO levels and above to stdout:

```python
import logging.config

LOG_FILEPATH = '/path/to/log_file.log'

logging.config.dictConfig({
    'version': 1,
    'disable_existing_loggers': False,
    'formatters': {
        'formatter': {
            'format': '%(asctime)s %(levelname)s %(message)s',
        },
    },
    'handlers': {
        'stdout': {
            'class': 'logging.StreamHandler',
            'stream' :  sys.stdout,
            'formatter': 'formatter',
            'level': 'INFO',
        },
        'log_file': {
            'class': 'logging.FileHandler',
            'filename': LOG_FILEPATH,
            'mode': 'a',
            'formatter': 'formatter',
            'level': 'WARNING',
        },
    },
    'loggers': {
        '': {
            'level': 'INFO',
            'handlers': ['stdout', 'log_file'],
        },
    },
})

logger = logging.getLogger('My logger')

# Usage
logger.info('Hello')       # Log infos to stdout
logger.warning('Oops')     # Log warnings to stdout and file
logger.error('Dang!')      # Log errors to stdout and file
```
