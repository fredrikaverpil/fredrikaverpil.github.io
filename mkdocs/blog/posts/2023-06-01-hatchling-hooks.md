---
date: 2023-06-01
draft: true
tags:
  - python
---

# Building Python wheels targeting different Python versions

This is a short post on how you can publish wheels onto PyPi.org, using custom hooks with [hatchling](https://hatch.pypa.io/latest/), that seamlessly targets different Python interpreter versions.

Let's say you want to distribute generated contents, which is intended to be consumed by the same Python interpreter version which generated the contents. This could be solved in a few different ways, but here I've opted to solve it on build time, meaning when you generate your wheel. So for example, I could generate three different wheels like this:

```bash
python3.9 -m build wheel
python3.10 -m build wheel
python3.11 -m build wheel
```

I could then publish them all onto PyPi.org under the same project name and version. Pip would then pick the wheel that was built and intended for the current Python interpreter version.

<!-- more -->

## Hatchling custom hooks

With hatchling we can configure custom hooks which will extract information from the Python interpreter used to build a wheel. In this case, we'll set up a build hook:

### `custom_build_hook.py`

```python
import sys
from typing import Any

from hatchling.builders.hooks.plugin.interface import BuildHookInterface


class CustomBuildHook(BuildHookInterface):
    """A custom build hook for building ."""

    def _python_tag(self) -> str:
        major = sys.version_info.major
        minor = sys.version_info.minor
        return f"py{major}{minor}"

    def initialize(self, version: str, build_data: dict[str, Any]) -> None:
        """Initialize the hook, update the build data."""
        if self.target_name not in ["wheel", "sdist"]:
            return

        # set python tag in wheel name
        build_data["tag"] = f"{self._python_tag()}-none-any"

```

## Project setup

All you'd then have to do is to update your `pyproject.toml` so to make use of the new build hook:

```toml
[build-system]
requires = ["hatchling"]
build-backend = "hatchling.build"

[project]
name = "myproj"
version = "0.1.0"
description = ''
requires-python = ">=3.9"
keywords = []
authors = []
dependencies = []

[project.optional-dependencies]
# PEP-440
build = [
  "build>=0.10.0",
]

[tool.hatch.build.targets.sdist]
[tool.hatch.build.targets.wheel]
packages = ["myproj"]
only-include = []

[tool.hatch.build.hooks.custom]
path = "tools/custom_build_hook.py"
```

## Install prerequisites and build wheel

We can now install the project and its dependencies (here only using hatchling and [pypa/build](https://pypi.org/project/build/)) in a virtual environment. Then we can build the wheel:

```bash
pip install -e '.[build]'
python -m build wheel
```

Rinse and repeat in e.g. CI using different Python interpreter versions. Finally we can upload all the wheels in the `dist/` folder onto PyPi.org!

## Bonus round; metadata hook

You can also use a metadata hook if you want to modify the version string with e.g. appending a timestamp or commit SHA to your version string. This technique could be used for internal projects where you just want to increment the version for each release. But there are perhaps better ways, depending on what you want to achieve. For example, if you want to let git commits drive your versioning via [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/), you may want to instead use [hatch-vcs](https://github.com/ofek/hatch-vcs).

### `custom_metadata_hook.py`

This example is a little contrived, as it replaces the semver patch with a timestamp and appends the commit SHA. So take this example more as a playground example, showing what could be possible.

!!! note "Git required"

    Note that this requires `git` to be installed.

```python
import subprocess
import sys
from datetime import datetime, timezone

from hatchling.metadata.plugin.interface import MetadataHookInterface


class CustomMetadataHook(MetadataHookInterface):
    def _get_git_revision_hash(self) -> str:
        return (
            subprocess.check_output(["git", "rev-parse", "HEAD"])
            .decode("ascii")
            .strip()
        )

    def _get_git_revision_short_hash(self) -> str:
        return (
            subprocess.check_output(["git", "rev-parse", "--short", "HEAD"])
            .decode("ascii")
            .strip()
        )

    def _python_tag(self) -> str:
        major = sys.version_info.major
        minor = sys.version_info.minor
        return f"py{major}{minor}"

    def _current_python_version(self) -> str:
        major = sys.version_info.major
        minor = sys.version_info.minor
        return f"{major}.{minor}.0"

    def _next_python_version(self) -> str:
        major = sys.version_info.major
        minor = sys.version_info.minor
        return f"{major}.{minor+1}.0"

    def _new_pkg_version(self, version: str) -> str:
        original_version = version
        major, minor, patch = original_version.split(".")
        now = datetime.now(tz=timezone.utc)
        timestamp = now.strftime("%Y%m%d%H%M%S")

        try:
            git_sha = self._get_git_revision_short_hash()
        except (subprocess.CalledProcessError, FileNotFoundError):
            git_sha = ""

        if git_sha:
            new_version = f"{major}.{minor}.{timestamp}+{git_sha}"
        else:
            new_version = f"{major}.{minor}.{timestamp}"

        return new_version

    def update(self, metadata):
        """Update the metadata."""
        requires_python = (
            f">={self._current_python_version()},<{self._next_python_version()}"
        )
        metadata["requires-python"] = requires_python

        metadata["version"] = self._new_pkg_version(metadata["version"])
```

Finally, you would have to also include this hook in your `pyproject.toml`:

```diff
[build-system]
requires = ["hatchling"]
build-backend = "hatchling.build"

[project]
name = "myproj"
version = "0.1.0"
description = ''
requires-python = ">=3.9"
keywords = []
authors = []
dependencies = []

[project.optional-dependencies]
# PEP-440
build = [
  "build>=0.10.0",
]

[tool.hatch.build.targets.sdist]
[tool.hatch.build.targets.wheel]
packages = ["myproj"]
only-include = []

[tool.hatch.build.hooks.custom]
path = "tools/custom_build_hook.py"

+ [tool.hatch.metadata.hooks.custom]
+ path = "tools/custom_metadata_hook.py"
```
