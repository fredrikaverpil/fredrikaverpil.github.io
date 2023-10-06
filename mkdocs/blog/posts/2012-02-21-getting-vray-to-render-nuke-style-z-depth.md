---
date: 2012-02-21
authors:
  - fredrikaverpil
comments: true
tags:
- nuke
- vray
---

# Getting V-Ray to render "Nuke-style" Z-depth

Out of the box, V-Ray for Maya does not render Z-depth the same way Nuke does. Here’s a quick fix for that (not applicable for most DOF plugins).

<!-- more -->

## Problem

When "3D-comping" something into an already 3D rendered image using a V-Ray generated Z-depth AOV to describe the depth, Nuke’s ZSlice will not work very well. The reason behind this is that Nuke internally computes Z depth from the camera plane and V-Ray computes Z depth from the camera point.

## Please note

For Nuke DOF plugins, such as [Bokeh](http://www.peregrinelabs.com/) or [Frischluft Lenscare](http://www.frischluft.com/lenscare/) (or Nuke’s Zblur), the standard V-Ray Z depth image is required and there is no need to tweak the depth output of V-Ray for Maya.

## Solution

Use the VRayExtraTex texture coupled with the “camera point” output of a samplerInfo texture to get the Nuke-style Z coordinate, which you can then transform in Nuke as needed.

##  Explanation (from [The Foundry](http://www.thefoundry.co.uk/))

> Nuke’s scanline render works as OpenGL and for a flat polygon respects the point of view when calculating depth. One way to go about it would be to make V-Ray change the Z-depth generation. Another possibility could be to generate at shading time the Z-buffer compatible with Nuke, in an AOV channel.
>
> The following formula should do the job:
>
> aov.z = 1 / P.z
> Where P is the sample point in the camera coordinate.