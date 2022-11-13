# boto3
### upload_fileobj
Instead of writing to disk, keep it in-memory:
```python
import boto3

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