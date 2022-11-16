---
title: 🌺 AWS
tags: [aws, s3]
draft: true
summary: "AWS snippets, gotchas etc."

# PaperMod
ShowToc: false
TocOpen: true

updated: 2022-11-16T00:13:11+01:00
created: 2022-11-14T20:42:48+01:00
---

## S3

### Use `upload_fileobj` for in-memory uploads

```python
import boto3

from somewhere import rows

s3 = boto3.client("s3", region_name="eu-west-1")

data = io.BytesIO(
    io.StringIO(json.dumps(rows, indent=4)).getvalue().encode(),
)

s3.upload_fileobj(
    data,
    Bucket=bucket_name,
    Key=f"{prefix}/{output_filename}",
)
```

### Define S3 endpoint URL

Some newer regions such as `af-south-1` [requires)](https://github.com/boto/boto3/issues/2728) the s3 client argument, so it is best to always define it.

```python
import boto3
from botocore.config import Config

boto_config = Config(
	 region_name="af-south-1",
	 # ...
)

s3 = boto3.client(
	"s3",
	config=boto_config,
	endpoint_url="https://s3.af-south-1.amazonaws.com",
)
```