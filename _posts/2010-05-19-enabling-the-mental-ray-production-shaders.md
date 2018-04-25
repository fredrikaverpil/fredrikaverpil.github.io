---
layout: post
title: Enabling the mental ray production shaders
tags: [maya, mentalray]
---

The production shader library contains a set of shaders aimed at production
users of mental ray. The library contains a diverse collection of shaders,
ranging from small utilities such as ray type switching and card opacity to
large complex shaders such as the fast 2d motion blur shader.

<!--more-->

In order to enable the mental ray production shader library, you can edit the
`mentalrayCustomNodeClass.mel` or simply type the following into the script
editor:

    optionVar -intValue "MIP_SHD_EXPOSE" 1;

Restart Maya and now, in the hypershade, you should be able to see the
production shaders.

For some reason the documentation for these shaders are still a bit dated.
The only version I’ve been able to find is the
[2007 documentation](http://www.mentalimages.com/fileadmin/user_upload/PDF/production.pdf).
Master Zap has an extensive blog post on all of this if
you’re new to the production shaders. When browsing through the release notes
of [mental ray 3.8.1.26](http://download.autodesk.com/us/maya/2011help/mr/relnotes/relnotes.html),
bundled with Maya 2011, it seems some of the shaders
have been updated and bugfixed. PixelCG has done a nice job recording video
[tutorials](http://www.pixelcg.com/blog/index.php?s=production+library) on each shader.
