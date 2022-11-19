---
title: Linear workflow, part 1
tags: [workflow, maya, nuke]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

date: 2010-05-27T02:00:00+02:00
---

There seems to be a general confusion on linear workflows and how 3D/compositing packages work, what sRGB/gamma 2.2 is for and why it’s a good idea to render images in linear data throughout the pipeline. This is my take on a linear workflow with Maya and Nuke as well as a bit of history on the subject.


Most graphics software use a linear data model. Simply put, this means that the math used when creating digital images works in a linear fashion, e.g. 2 + 2 = 4. In effect, you can say that linear math is applied when a 3D renderer is calculating what the resulting image should be like. However, our display devices, such as monitors, do not use linear math when visualizing this data as images. They use non-linear (logarithmic) math due to the technology used in recording and displaying images. For them this could in some cases mean that the linear data of e.g. 2 + 2 = 4 is interpreted as 2 + 2 = 10. In order make the linear data which makes up an image show on a device that needs to visualize non-linear data, we need to convert or rather “correct” the linear data image. This is done using [gamma](http://www.poynton.com/notes/colour_and_gamma/GammaFAQ.html#gamma). The gamma varies depending on the display device but the most common gamma for laptops and computer display devices these days is gamma 2.2, also known as [sRGB](http://en.wikipedia.org/wiki/SRGB) (although sRGB is not exactly gamma 2.2). So a [gamma correction](http://www.poynton.com/notes/colour_and_gamma/GammaFAQ.html#gamma_correction) curve is applied and baked into the image file by the graphics software so that the image is shown correctly on our display devices. The gamma correction for sRGB devices (with a gamma of approximately 2.2) is ½.2 = 0.4545.

So what’s the problem, you might ask. Well, there are two main problems.

The first problem resides in that we sometimes feed images or textures, with gamma correction baked into them, into a 3D scene. Then we render that out into a new image and then perhaps again applying gamma correction to the image. Now there are two gamma correction curves applied for that specific texture, which makes it look like it’s got a higher contrast than it actually should have. If any additional operations were to be applied to this image (using a second graphics manipulation software such as Photoshop or a compositing package), we do not have the original and linear pixel values stored in the file any more – which we need if we intend to process the image further – and in effect, we are working with a compromised image. Instead, by keeping all images in the graphics pipeline stored in linear space and always rendering with gamma correction at 1.0, this will solve the issue.

The other problem resides within some software applications (such as Maya), which render out a linear image by default but does also present the linerar image in the render view window, assuming we look at the image using a display device using gamma 1.0. Since most users of Maya use a display device using gamma 2.2/sRGB and do not know about this issue, they simply crank up shaders and light values which breaks physical light simulation in the process in order to get the visual end result they are after. The remedy to this issue is to make the render view window display a gamma corrected render at 2.2/sRGB while shading and lighting and then, just before rendering out the final image files, switching the gamma correction off so that the file on disk remains in linear data format.

## Summary

Make sure everything you feed into the process of 3D rendering and compositing is kept linear without gamma correction applied. When previewing the shading and lighting of a 3D render you should temporarily view the image using a gamma correction of 2.2 or sRGB. When you are at the final stage at displaying the image on some device, you can worry about applying the gamma correction (or not) to the actual file. If showing the rendered, comped and processed images on a normal computer screen, you can safely apply the gamma correction of sRGB as the final step, as they would not look right without it.

More reading on gamma and linearism:

- [Charles Poynton’s FAQ on gamma](http://www.poynton.com/notes/colour_and_gamma/GammaFAQ.html)
- [Gamma correction at Wikipedia](http://en.wikipedia.org/wiki/Gamma_correction)
- [My mental ray on gamma](http://mymentalray.com/wiki/index.php/Gamma)
- [Master Zap’s linear workflow website](http://www.lysator.liu.se/~zap/lwf/)
- [Linear workflow by David Johnson](http://www.djx.com.au/blog/2008/09/13/linear-workflow-and-gamma/)

So, how do I preserve a linear workflow using Maya and Nuke?

– Check out [[2010-06-27-linear-workflow-part-2]], part 2 of this thrilling blog series.
