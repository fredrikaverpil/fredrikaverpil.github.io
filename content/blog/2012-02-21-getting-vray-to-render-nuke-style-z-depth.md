---
title: "Getting V-Ray to render Nuke-style Z-depth"
date: 2012-02-21
tags: ["vray", "nuke", "maya"]
categories: ["archive"]
---

If you have come this far you might be interested in a couple of small tips and tricks on how to get that 32-bit floating point depth channel to render out of V-Ray and into Nuke, similar to what you might be used to from Mental Ray and the IFF file format.

## Table of contents

1. Creating a V-Ray Z-depth pass in Maya
2. Bringing the V-Ray Z-depth pass into Nuke

## Creating a V-Ray Z-depth pass in Maya

### Render settings

Open up the render settings, go to the `V-Ray` tab and make sure that `Image format` is set to `OpenEXR` (multichannel) and that `OpenEXR data type` is set to `32-bit float`.

Now go to the `Render Elements` tab, enable `V-Ray Render Elements` and `Multi-channel EXR`. Add the `ZDepth` render element. The default settings should be fine.

![](/blog/vray/vray_render_settings.png)

### Z-depth in Nuke

![](/blog/vray/nuke_zdepth.png)

Load the EXR file (containing the Z-depth pass) into Nuke using the `Read` node. Now, view the `Z` channel in the viewer. You will see an entirely black image!

This is because V-Ray stores its Z-depth in a different manner from Mental Ray. In Nuke, create a `ZDepth` node and connect its `Input` to your `Read` node. The result should look like the image above.

Now you should be able to see a grayscale Z-depth pass! You can then adjust the black and white point of the Z-depth pass.

This is a quick and dirty way to work with Z-depth in Nuke. For a more robust workflow, you can use the `VrayZDepth` node from `rv-nodes` (which is a collection of various Nuke nodes written by different artists and can be found on Nukepedia).
