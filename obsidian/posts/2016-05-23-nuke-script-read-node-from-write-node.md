---
title: "Read node from Write node"
tags: [nuke, python]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

created: 2016-05-23T03:00:12+02:00
updated: 2022-11-16T00:19:40+01:00
---

Originally posted in 2011; A Python script for [Nuke](https://www.thefoundry.co.uk/products/nuke/) which takes any selected Write (or Read) node and creates a Read node from it. Now updated to fix some bugs and support for a wider range of scenarios.



The script will actually not look for Write nodes in the selection. Instead, it'll look for a specific knob in each of the selected nodes. If the knob is found in the node, the script will attempt to detect an image sequence or movie file based on the knob value (expressions are supported) and then create a Read node with the detected imagery. By default, the script will look for the `file` knob, which makes it compatible with Write, Read and possibly other nodes too. You can make the script look for other knobs by adding their names to the `FILEPATH_KNOBS` list, which is particularly useful if you're into customization.

See more about these settings a bit further down.

### Download

- [Nukepedia](http://www.nukepedia.com/python/misc/readfromwrite)
- [Github](https://raw.github.com/fredrikaverpil/nuke/master/scripts/readFromWrite.py)

### Installation instructions

Place the Python script in the /scripts dir inside your `NUKE_PATH` (see [[2011-10-28-nuke-63-small-studio-setup-for-windows-osx]] on setting this up). Add the following to your `menu.py`:

```python
import readFromWrite
nuke.menu( 'Nuke' ).addCommand( 'My file menu/Read from Write', 'readFromWrite.ReadFromWrite()', 'shift+r' )
```

You should now be able to select any Write node(s) and hit `shift+r` to generate corresponding Read node(s)!

### Release notes / changelog

```yaml
Changelog:
- v2.3:
    - Bug fix: crash when knob "use_limit" isn't available on node
    - Accidentally left ReadFromWrite() at bottom of script in v2.2
- v2.2:
    - Support for nodes with filepath which does not exist on disk
      (will read Write node settings or incoming framerange)
    - Support for additional Read/Write node option "raw"
v2.1:
  - Fixed bug where Read node always got premultiplied
  - Support for ../ in filepath/expression
  - Dialog on "filepath not found" error
  - Set origfirst, origlast framerange
  - Additional movie file format support (see SINGLE_FILE_FORMATS variable)
  - General cleanup of various methods for greater maintainability
v2.0:
  - Completely rewritten from scratch
  - Improved detection of frame range
  - Supports any padding format (not only %04d)
  - Applies colorspace to Read node
  - Supports not only Write nodes (see FILEPATH_KNOBS variable)
  - Supports definition of "single file image sequence" formats (see SINGLE_FILE_FORMATS variable)
  - PEP8 compliant!
```

### Settings

- `FILEPATH_KNOBS` - File knobs which should be parsed
- `SINGLE_FILE_FORMATS` - A list of filetypes which are expected to contain more than one frame but in fact has more than one frame (e.g. "*.mov" for Quicktime movie files)

Look in the top portion of the file for the settings definitions:

```python
# Settings
#
# knob names to use (from any node) as a base for Read node creation.
FILEPATH_KNOBS = ['file']
#
# This scripts needs to know whether to apply padding to a filepath
# or keep it without padding. Movie files should not have padding,
# for example. Add such "single file formats" here.
SINGLE_FILE_FORMATS = ['avi', 'mp4', 'mxf', 'mov', 'mpg',
                       'mpeg', 'wmv', 'm4v', 'm2v']
```
