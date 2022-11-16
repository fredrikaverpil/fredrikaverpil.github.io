---
title: Standalone PySide2 wheels
tags: [python, pyside]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

created: 2018-02-20T02:00:12+01:00
updated: 2022-11-15T17:29:41+01:00
---

The Qt Company has yet to release official, standalone and pip-installable PySide2 wheels. However, since they made it possible to build standalone wheels successfully, I'm now building such unofficial, standalone wheels here using free CI services (thanks [Travis](https://travis-ci.org/) and [AppVeyor](https://www.appveyor.com/)!):


- [fredrikaverpil/pyside2-windows](https://github.com/fredrikaverpil/pyside2-windows)
- [fredrikaverpil/pyside2-macos](https://github.com/fredrikaverpil/pyside2-macos)
- [fredrikaverpil/pyside2-linux](https://github.com/fredrikaverpil/pyside2-linux)

**Update 2018-03-09**: The Qt Company now offers official and standalone wheels, read more here: [[2018-03-09-official-pyside2-wheels]]
**Update 2018-07-17**: PySide2 can now be installed from pypi.org: `pip install PySide2`!


### Download

You'll find the wheels under "releases" in each repository and you can pip-install the wheels like so:

```bash
pip install <URL to wheel>
```

Since Github doesn't have a storage limit to releases as of writing this, I would expect the URLs to work nicely for the forseeable future.


### Version string nomenclature

The Qt Company still doesn't maintain the version string of PySide2, so therefore I'm tagging releases based on the date when they were built.

If you're wondering what version you're running, you may be able to query any of the following to receive some hints, which became available in PySide2 [sometime in late August 2017](https://codereview.qt-project.org/#/c/202199/):

```python
PySide2.__build_date__  # the date when the package was built in iso8601 format
PySide2.__build_commit_date__ # the date of the top-level commit used to build the package
PySide2.__build_commit_hash__  # the SHA1 hash of the top-level commit
PySide2.__build_commit_hash_described__  # the result of 'git describe commmit'
```

You can then cross reference the commit date/hash against the [`pyside-setup` git repository](http://code.qt.io/cgit/pyside/pyside-setup.git/) if you wish to figure out exactly which commit was used to build the wheel. Or you can check the respective CI output where I print the commit's info (`git log -n 1`).
