---
date: 2023-06-01
draft: true
tags:
  - python
---

# Using hatchling to deploy per-python version wheels to PyPi

```bash
pip install -e '.[build]'
python -m build --wheel
```

```toml
[build-system]
requires = ["hatchling"]
build-backend = "hatchling.build"

[project]
name = "myproj"
version = "0.1.0"  # semver patch will be dynamically updated by custom_metadata_hook.py
description = ''
requires-python = ">=3.10.0"
keywords = []
authors = []
dependencies = []

[project.optional-dependencies]
# PEP-440
build = []

[tool.hatch.build.targets.sdist]
[tool.hatch.build.targets.wheel]
packages = ["myproj"]
only-include = []

[tool.hatch.build.hooks.custom]
path = "tools/custom_build_hook.py"

[tool.hatch.metadata.hooks.custom]
path = "tools/custom_metadata_hook.py"
```

```python
# custom_build_hook.py

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

```python
# custom_metadata_hook.py

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
