---
ShowToc: false
TocOpen: false
cover:
  alt: ''
  image: /static/vray/vray_spherical_hdr_01.png
  relative: false
date: 2012-12-12 01:00:12+01:00
draft: false
tags:
- vray
- maya
title: Generating a spherical HDR image with V-Ray for Maya
---

1. Create a camera and place it where your objects resides which the HDR dome should affect.
2. Add the V-Ray attribute “Camera Settings” to the camera and scroll down to “Extra Attributes” in the Attribute Editor. Here, set Type to “Spherical”. Override the FOV and set it to 360 degrees.
3. In the render settings, set the output image format to .hdr, turn subpixel mapping off and do not clamp output.
4. Render with width/height ratio 2:1, e.g. 2048x1024 px.

## Sample scene

Example Maya scene provided [here](/static/vray/spherical_hdr_gen_maya.ma) (save as); a simple cube environment, an area light and a camera.

![](/static/vray/vray_spherical_hdr_02.png)
*Sample Maya scene.*