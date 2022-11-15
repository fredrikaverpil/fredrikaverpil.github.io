---
title: Enabling Iray in Maya 2011
tags: [maya, mentalray]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

created: 2010-05-27T02:00:00+02:00
updated: 2022-11-15T23:04:04+01:00
---

Mental images and/or Autodesk has intentionally not made Iray available in the latest release of Autodesk Maya; version 2011. But there is a workaround...



The initial version of [Iray](http://www.mentalimages.com/products/Iray.html) is bundled with mental ray 3.8 (stand alone and reality server) and takes advantage of the GPU when lighting, shading and rendering. According to a [mental forums thread](http://forum.mentalimages.com/showthread.php?t=6503), Iray was intentionally left out of mental ray bundled with Maya 2011 since it’s still under heavy development. It was only included in the stand alone and reality server flavours and provided “as is” without further support.

Numerous posts on the intertubes has indicated that you can indeed make Iray work within Maya 2011 if you copy a few DLL files from the stand alone version of mental ray 3.8 into your Maya 2011 installation. However, as you are then leaving the specification and guaranteed functionality of the software and probably violating the installation/license agreement, you do this on your own risk.

Quote from [rBrady’s post](http://forum.mentalimages.com/showthread.php?t=6424&page=2) on the mental forums:

> ... Iray works just fine in Maya 2011 if you add 3 dlls from standalone. Really, its only two, cudart.dll is just a cuda dll from nvidia. I forget what the actual Iray dlls are, but they are in the mental ray standalone bin folder and they have Iray in the name. You simply put them in the maya2011/bin folder and away you go. Just add Iray as a string option like you would the progressive render. It give you that nice preview in the renderview and everything

Then you’re supposed to turn Iray on inside of Maya as well, according to [this post](http://forums.cgsociety.org/showpost.php?p=6500383&postcount=34), by entering the following in the script editor:

    miOptionsAddNewStringOpt miDefaultOptions.stringOptions;
    setAttr -type "string" miDefaultOptions.stringOptions[28].name "Iray";
    setAttr -type "string" miDefaultOptions.stringOptions[28].value "on";
    setAttr -type "string" miDefaultOptions.stringOptions[28].type "boolean";

More test renders and information on Iray usage can be found over at [this CGTalk thread](http://forums.cgsociety.org/showthread.php?t=877874). And oh yeah, don’t forget you’ll need a [CUDA-compatible](http://www.nvidia.com/object/cuda_gpus.html) graphics card in the box.

Personally, I’ll stick with what was intended for the Maya 2011 release and check out [Octane](http://www.refractivesoftware.com/) instead.

Please note: For Maya 2012, I have been told you need the following files: cudart64_7.dll, libIray.dll and libIraymr.dll
