---
ShowToc: false
TocOpen: false
date: 2010-11-24 01:00:00+01:00
draft: false
tags:
- maya
title: Managing Maya scenes on multiple platforms
---

Ever found yourself trying to open a complex Maya scene with nested references on a Windows/Linux network using your Macbook (or any of the other OS combos)?

I work in a Windows Server environment where PC drive letters are being used – and as I’m on a Mac system this poses a big problem when trying to open a complex Maya scene, linking to files outside of the Maya project; meaning I have to click through numerous dialogues and set the path so that my Mac will find the files. However, the `dirmap` command comes to the rescue!

Dirmap is a mel command (and also available in Python) which will map any path into another one. This means I can map i.e. `S:/`, which my Mac won’t understand, into `/Volumes/myShare` and in a Windows environment you could do the inverse.

As I just recently realized this little command was available, this has made my day a lot better as I no longer have to click and find that other path every time a reference – or other file outside of the Maya project directory – could not be found.

So, if you also experience this hurdle when your fingers itch and you just need to get that scene open, read on…

Start by enabling directory mapping by entering the following MEL code into the script editor:

    dirmap -en true;

Then map the directories. In my example, the PC workstations which mount the server to drive letter `S:/` and on my Macbook the mount is `/Volumes/myShare/`

    dirmap -m "S:/" "/Volumes/myShare/";
    dirmap -m "s:/" "/Volumes/myShare/";

And if you are wondering why there were two rows of code there; case sensitivity at its finest. Repeat the command above for any additional locations that need mapping.

Voila, you can now sip that espresso as your scene loads rather than having to sit there and click through all those dialogues.

Whenever you close Maya, it will forget your settings, so in order to make Maya remember these mappings, either create a shelf button or add the code to `userStartup.mel`. For full reference on `dirmap`, check out the [Maya 2012 documentation](http://download.autodesk.com/global/docs/maya2012/en_us/index.html).

Many thanks to [Michiel Duvekot](http://www.thnkr.com/) for letting me know about `dirmap`!