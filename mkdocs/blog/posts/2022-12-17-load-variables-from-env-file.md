---
date: 2022-12-18
draft: false
tags:
- bash
---

# Load variables from .env file into current environment

Load an .env file into the environment prior to running something which requires the environment variables:

!!! note "`.env` file"

    ```env
    MY_SUPER_SECRET_TOKEN="foo"
    ```

```bash
set -a
source .env
set +a
```

<!-- more -->

You can now use the variables from the `.env` file in scripts:

```bash
$ echo $MY_SUPER_SECRET_TOKEN
foo
```

## Load .env file from `.bashrc`, `.zshrc`, etc.

If you keep your configuration, like I do, in a public dotfiles repository, you might want to keep your
`.env` file secret, and make sure it is added to your `.gitignore` file. Then you can have it loaded
from your `.bashrc` or `.zshrc` file:

!!! example ".bashrc"

    ```bash
    if [ -f $HOME/.env ];
    then
        set -a
        source $HOME/.env
        set +a
    else
        echo "Warning: $HOME/.env does not exist"
    fi
    ```
