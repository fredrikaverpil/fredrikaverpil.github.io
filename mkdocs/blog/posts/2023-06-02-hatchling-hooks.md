---
date: 2023-06-02
draft: false
tags:
  - python
---

# Building Python wheels targeting different interpreter versions

This is a short post on how you can publish wheels onto [PyPi](https://pypi.org), using [hatchling's](https://hatch.pypa.io/latest/) custom metadata hook, that seamlessly targets different Python interpreter versions.

<!-- more -->

Let's say you want to distribute some generated contents, which is intended to be consumed by the same Python interpreter version which generated the contents. This could be solved in a few different ways, but here I've opted to use a matrix of Python versions in CI. For each Python version, I generate the desired content and then I build the wheel, using some custom hatchling hooks. The end result is a wheel (per Python version) that can only be installed by that same Python version.

I can then publish them all onto PyPi under the same project name and project version. When running `pip install ...`, pip would then pick the wheel that was built and intended for the Python interpreter version I'm using, guaranteeing that the version generated the wheel contents will be used to also consume the contents.

In `pyproject.toml`, you can specify metadata such as for example the version string. Hatchling offers the ability to write custom hooks so to edit this metadata when e.g. building the wheel. It also provides hooks to explain _how_ the wheel should be built. What we want to do here is edit the `requires` metadata to only the Python version(s) we want to allow.

## Constraining the required Python version

Let's start with editing the metadata of the wheel, so we can constrain the required Python version. Let's add `custom_metadata_hook.py`:

!!! example "`custom_metadata_hook.py`"

    ```python
    import sys

    from hatchling.metadata.plugin.interface import MetadataHookInterface


    class CustomMetadataHook(MetadataHookInterface):

        def _current_python_version(self) -> str:
            major = sys.version_info.major
            minor = sys.version_info.minor
            return f"{major}.{minor}.0"

        def _next_python_version(self) -> str:
            major = sys.version_info.major
            minor = sys.version_info.minor
            return f"{major}.{minor+1}.0"

        def update(self, metadata):
            """Update the metadata."""
            requires_python = (
                f">={self._current_python_version()},<{self._next_python_version()}"
            )
            metadata["requires-python"] = requires_python
    ```

In order to build the wheel with this hook, we'll need to tell hatchling about this new hook file, in `pyproject.toml`, under the `tool.hatch.build.hooks.custom` section.

!!! example "`pyproject.toml`"

    ```toml
    [build-system]
    requires = ["hatchling"]
    build-backend = "hatchling.build"

    [project]
    name = "myproj"
    version = "0.1.0"
    description = ''
    readme = "README.md"
    requires-python = ">=3.9"
    license = "MIT"
    keywords = []
    authors = [
      { name = "Fredrik Averpil", email = "fredrik.averpil@gmail.se" },
    ]
    classifiers = [
      # https://pypi.org/classifiers/
      "License :: OSI Approved :: MIT License",
      "Development Status :: 4 - Beta",
      "Programming Language :: Python",
      "Programming Language :: Python :: 3.9",
      "Programming Language :: Python :: 3.10",
      "Programming Language :: Python :: 3.11",
      "Programming Language :: Python :: Implementation :: CPython",
      "Programming Language :: Python :: Implementation :: PyPy",
    ]
    dependencies = []

    [project.optional-dependencies]
    # PEP-440
    build = [
      "build>=0.10.0",
    ]

    [tool.hatch.build.targets.sdist]

    [tool.hatch.build.targets.wheel]
    packages = ["myproj"]
    only-include = ["myproj"]

    [tool.hatch.metadata.hooks.custom]
    path = "tools/custom_metadata_hook.py"
    ```

My project now looks something like this:

```
.
├── LICENSE.txt
├── README.md
├── pyproject.toml
├── src
│   └── myproj
│       └── __init__.py
└── tools
    └── custom_metadata_hook.py
```

You should now be able to build the wheel and constrain it to the same Python version you used to build the wheel. I'm using [pypa/build](https://github.com/pypa/build) to build the wheel and therefore I need to first make sure I have that installed before building:

```bash
pip install build
python -m build --wheel
```

!!! tip "Pro tip!"

    You can add `print(metadata)` in the `update` function and run `python -m build --wheel` to see a printout of all the data that you can modify here.

If you try to pip-install this wheel using a different Python version, it should fail. This is using `pip` from Python 3.11 trying to install a wheel built with Python 3.10:

```bash
$ pip install dist/myproj-0.1.0-py310-none-any.whl
Processing ./dist/myproj-0.1.0-py310-none-any.whl
INFO: pip is looking at multiple versions of myproj to determine which version is compatible with other requirements. This could take a while.
ERROR: Package 'myproj' requires a different Python: 3.11.3 not in '<3.11.0,>=3.10.0'
```

You can setup your CI so that it uses a matrix of Python versions. For each Python version you generate the wheel contents, build a wheel and store the wheel as CI build artifact. As a final step you can have a CI step that fetches all the built CI wheel artifacts and uploads them to PyPi. Great success! 🎯

You can read more about hatchling's metadata hooks [here](https://hatch.pypa.io/latest/plugins/metadata-hook/custom/).
