---
layout: post
title: 'Nuke script: Read node from Write node'
tags: [nuke, python]
---

<div class="message">
  Howdy! There's a newer version of this script <a href="/2016/05/23/nuke-script-read-node-from-write-node">here</a>.
</div>

Written in Python, with Nuke 6.9 in mind, generate a Read node from the selected Write node.

<!--more-->

The script will attempt to load an image sequence or a single file (such as a movie file), based off the selected Write node. If the first and last frame can not be determined, it will fall back to the project settings’ frame range and throw a warning.

The script comes with one limitation; it assumes you render out any file sequences with a frame padding of four digits surrounded by period signs:

- .####.
- .%04d.

Supported real world examples:

- lotr_seq1_shot1_v001_fa.####.exr
- blah.%04d.jpg
- hello.mov

Download: [readFromWrite.py](https://raw.github.com/fredrikaverpil/nuke/master/scripts/readFromWrite.py)

- v1.4: Better sequence detection. If you had problems with 1.3, try this.
- v1.3: Write nodes with expressions now gets evaluated correctly.
- v1.2: Takes the ‘premultiplied’ setting from the Write node, auto-detects movie files frame ranges, better logic when detecting a frame range
- v1.1: Support for filenamefilter callback
- v1.0: Initial release

Place the Python script in the /scripts dir inside your `NUKE_PATH` (see my [previous post]({{ site.baseurl }}/2011/10/28/nuke-63-small-studio-setup-for-windows-osx/) on setting this up). Add the following to your `menu.py`:

```python
import readFromWrite
nuke.menu( 'Nuke' ).addCommand( 'My file menu/Read from Write', 'readFromWrite.readFromWrite()', 'shift+r' )
```

You should now be able to select any Write node and hit Shift + R to generate a Read node!
