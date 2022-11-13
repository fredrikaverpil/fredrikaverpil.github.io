---
layout: post
title: Fujifilm Finepix X100 – LCD/EVF observations
tags: [photography]
---

![]({{ site.baseurl }}/blog/assets/x100/x100_observations_01.jpg)

Some people have asked whether the “Preview Depth of Field” function is different from the regular LCD/EVF view or not (which it is), why the histogram is acting up (which it sometimes can do) and I’ve also seen questions regarding why exposure lock (AE-L) mapped to the AFL/AEL-button does not seem to be happening when pressing the AFL/AEL-button. Here’s an attempt to try and explain why this is so.

<!--more-->

Below, the two stages of what the LCD/EVF will show when shooting a picture, explained.

### 1. Just looking through the EVF or on the LCD

When looking through the EVF or on the LCD, without pressing any buttons, the camera will automatically set its own aperture value for the LCD/EVF view (and therefore also DOF and exposure). It does this in order to create a good visual representation to show what is in frame, regardless of what the physical dials on the camera are set to.

This does in no way alter your choice of aperture, shutter speed, ISO etc for the photo you are about to take. But in short, framing is the only thing you can really trust at this point.

I believe that the image shown in the LCD/EVF here is a good representation of what the camera would record if set to P mode (fully automatic).

Unfortunately, the luminance histogram in the LCD/EVF seem to be based on these arbitrary settings in real-time which actually renders it unreliable in some cases. When in A mode (aperture priority), S mode (shutter priority) or P mode (program AE, fully automatic), the histogram is somewhat accurate, as the camera will strive to achieve the exposure portrayed in the LCD/EVF. But whenever you are shooting in M mode (manual exposure) you never know if your luminance histogram is even close and you might as well just shoot an image and check the luminance histogram in the playback mode. Because of this behaviour, this leads me to believe that in the OVF, the histogram is based off the very same arbitrary values that the camera has chosen, being just as unreliable (again, especially in M mode).

As the OVF power saving feature of the X100 disables the histogram, it would make sense that the X100 would indeed save a lot of power by not having to calculate the histogram based off the arbitrary values recorded, processed (but not shown unless in LCD/EVF) when using the OVF. However, I would not recommend using the OVF power saving feature for this particular reason as it also slows down auto focus.

### 2. Half-pressing the shutter button

When half-pressing the shutter button, evaluations are being made and additional preview features will kick in such as an aperture preview, not necessarily according to the physical aperture dial setting. This means the DOF here is, again, a preview. The focus is locked. Also, the exposure preview in the LCD/EVF will be locked. The actual exposure (which is not the same as the exposure preview), will also be locked (if not in M mode), and this may or may not be accurately portrayed in the exposure preview (more on this further below).

Now, regarding the exposure preview: If there’s plenty of light, the exposure is mostly quite accurately portrayed in the preview, but it’s when you stop down the lens to f/11 or f/16 in dim-lit situations when the exposure preview function will most likely be way off. Then the LCD/EVF preview can show you a bright image and as you fully press the shutter to take the photo, the photo comes out all black. Just set the camera to f/16 and at a shutter speed of 1/4000 and see for yourself if the half-press preview matches the actual photo taken when in dim-lit indoor situations.

## "Preview depth of field" vs regular LCD/EVF view

Since you cannot change aperture and actually see live updates of this in the LCD/EVF, Fujifilm has created the “Preview Depth of Field” function. It will stop down the lens accordingly in the LCD/EVF as you move up the aperture values and should give you a much better idea of what the DOF will be like compared to the regular LCD/EVF view (which, explained above, uses an arbitrary aperture value chosen by the camera to minimize video noise etc).

This feature does not take shutter speed or ISO into account, and will show you the DOF preview at a shutter speed as well as ISO chosen by the camera. Therefore, all you really get a preview of here is the DOF.

## Exposure lock (AE-L) with the AFL/AEL-button

Setting the AFL/AEL-button to lock exposure only (AE-L) does indeed perform the actual exposure lock when pressed, but only in A or P mode (AE-L is useless in M mode). Not until you half-press the shutter you will see an exposure preview of this in the LCD/EVF.

Do note that in manual focusing mode (MF) the AFL/AEL-button serves as a single AF-run initiator and overrides any assignment to this button. Therefore having AE-L or AF-L assigned to the AFL/AEL-button during MF does not lock anything. This however, can be seen as a strength. Use the manual focusing ring or the AFL/AEL-button to find focus, half-press the shutter to lock exposure and finally shoot the picture.

## Other X100 resources

Make sure to check out some of my other X100 articles:

- [Fujifilm Finepix X100 – settings & notes]({{ site.baseurl }}/2012/02/02/fujifilm-x100-settings-and-notes/)
- [Fujifilm Finepix X100 – a year in retrospect]({{ site.baseurl }}/2012/06/12/the-fujifilm-x100-a-year-in-retrospect/)
- [Fujifilm Finepix X100 – RAW and film simulation]({{ site.baseurl }}/2012/08/04/fujifilm-finepix-x100-raw-and-film-simulation/)
