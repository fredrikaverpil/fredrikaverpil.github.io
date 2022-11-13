# Load variables into current environment
Load an .env file into the environment prior to running something which requires the environment variables:
```bash
set -a
source <somefile.env>
set +a
```