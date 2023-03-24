---
date: 2012-02-02
tags:
- photography
---

# Fujifilm X100 settings & notes

![](/static/x100/fujifilm_x100.jpg)

Some facts about the camera which are easy to forget as well some notes to self.

<!-- more -->

## ISO

- The hardware is only capable of maximum ISO 1600. Any higher ISO value is pushed with software in camera.
- ISO value also serves as lowest ISO value allowed when Auto ISO is on.
- Auto-ISO can be set to a maximum ISO value between 400 - 3200.
- When shooting RAW (with Auto-ISO turned off), ISO can be set to a minimum of 200 and a maximum of 6400.
- When shooting in JPG mode (with Auto-ISO turned off), ISO can be set to a minimum of 100 and a maximum of 12800.
- When in Auto-ISO, the “Min shutter” value will not be respected if you hit the Max ISO value. The Min shutter value could therefore be regarded as a “soft limit” Min shutter value.

## Dynamic range (tone mapping)

- DR100 is the default setting for the camera.
- DR200 (changes minimum ISO to 400) will stop down the RAW 1 stop.
- DR400 (changes minimum ISO to 800) will stop down the RAW 2 stops.
- RAW files with DR200 or DR400 will appear darker in post production as they are stopped down, and there is currently no way for e.g. Lightroom to interpret this feature in order to compensate (the way the JPG engine does).
- Dynamic range is really only good if you use the camera’s JPG engine. If you shoot RAW only, you should set this to DR100 and tweak exposure in post production.

## Auto focus

- The X100 uses Contrast Detection Auto Focus (CD AF), which means you should never try to find focus at the edge between two objects of different distance from the camera. More details [by Arjay](http://www.x100forum.com/index.php?/topic/1713-focus-101/page__view__findpost__p__19694) on this over at the X100 Forum.
- The X100 is not a true rangefinder camera, thus is not great at manual focus. For best AF performance, use AF-S with smallest AF window in Electrical View Finder (EVF).
- The X100 will have more difficulty locking AF onto horizontal lines than on vertical lines.
- The AFL/AEL button is by default set to Auto Exposure Lock (AE-L), by setting this to AF-L, auto focus can be locked instead.
- Optical View Finder (OVF) Power Save Mode set to “on” will reduce AF performance, according to the manual.
- AF-C is reportedly increasing chances of an AF lock in dim-lit situations.

## Lens characteristics

- Prime f/2.0 23mm, equivalent to 35mm on a 35mm camera.
- The lens is at its sharpest at f/4-f/5.6 and is very, very soft at f/2. For high depth-of-field, f/11 is a good compromise but at f/16 effects of diffraction increases.
- Built in ND filter (equivalent f-stop reduction of 3).

## Memory card speed tests

Here I have compared the memory card write speeds of two different versions of the SanDisk Pro Extreme 32GB: 45MB/s vs 95MB/s.


    Type of test			45MB/s		95MB/s
    ------------------------------------------------------
    1 RAW+F				5.3s		3.5s
    5 RAW+F (continuous burst)	26.3s		19.8s
    1 JPG/F				1.5s		1.5s
    5 JPG/F (continuous burst)	15.6s		15.6s


It seems there is no reason for JPG-only shooters to upgrade their cards for faster speeds but some speed gains are seen when shooting RAW.

## My settings (firmware 1.30)

Starting with the default settings after a factory reset or firmware upgrade:

- ISO: 200
- ISO Auto Control: On
  - Max Sensitivity: 1600 (this depends on the situation)
  - Min Shutter Speed 1/125
- Image Quality: F+RAW
- (Dynamic Range: DR100)
- Film Simulation: Astia
- (Color: Mid)
- (Sharpness: Standard) – According to Fuji, sharpness is not applicable with Astia or Velvia as they are film simulations. However, the X100 does indeed apply more sharpness to JPGs with this set to “High”.
- (Highlight Tone: Standard)
- (Shadow Tone: Standard)
- Noise Reduction: Low
- Operation Vol: Lowest
- Shutter Vol: Off
- (Quick Start Mode: On Off - with firmware 1.30, I can’t see any difference in startup speed anymore)
- (Fn Button: ISO)
- RAW Button: ND Filter
- AE/AF-Lock Button: AF-L
- Corrected AF Frame: On

Disp. Custom Setting, OVF:

- Electronic Level: Off
- White Balance: Off
- Film Simulation: Off
- Dynamic Range: Off
- Frames Remaining: Off
- Image Size/Quality: Off
- Battery Level: Off

And last but not least I set the focus area size for AF-S to the smallest possible as well as enable Auto flash when in fully automatic mode (P).

## Documentation, firmware and files

- [User Manual (PDF)](http://www.fujifilm.com/support/digital_cameras/manuals/pdf/index/x/finepix_x100_manual_01.pdf)
- [Firmware download](http://www.fujifilm.com/support/digital_cameras/software/firmware/x/finepix_x100/index.html)
- [Lightroom JPEG Lens Profile (save as...)](/static/x100/Fujifilm_FinePix_X100_JPEG.lcp)

## Reviews

Please note, most of the reviews below were written before the release of firmware 1.21 (released in March, 2012), which dramatically increases auto focus speed, improves manual focus ring responsiveness as well as makes the RAW-button configurable (among other things).

- [DPReview’s X100 review](http://www.dpreview.com/reviews/fujifilmx100/)
- [Luminous Landscape’s X100 review](http://www.luminous-landscape.com/reviews/cameras/fujifilm_x100_test_report.shtml)
- [DxOMark’s X100 review](http://www.dxomark.com/index.php/Publications/DxOMark-Reviews/Fujifilm-X100-DxOMark-Review)
- [Image Resource’s X100 review](http://www.imaging-resource.com/PRODS/X100/X100A.HTM)
- [Zack Arias’ X100 review](http://zackarias.com/for-photographers/gear-gadgets/fuji-x100-review/)
- [Linhbergh’s X100 impressions](http://linhbergh.com/blog/2011/08/fuji-x100-impressions/)
- [DigitalRev’s Fujifilm Finepix X100 hands-on review (YouTube)](http://www.youtube.com/watch?v=L-VoXxwGWYc)
- My own review, written having used the X100 for one year: Fujifilm Finepix X100 – a year in retrospect

## Bugs, quirks and issues

I have assembled a list of known bugs and quirks for the current X100 firmware which can be found [here](http://www.fujix-forum.com/index.php?/topic/4097-x100-firmware-bugsquirks-overview/).

After having used the camera for about a year, [lug wear from the triangle ring neck](http://forums.dpreview.com/forums/readflat.asp?forum=1020&thread=40754091&page=1) is unfortunately becoming apparent. I’d recommend not using the ones supplied by Fuji and instead resort to circular rings or [Gordy’s string straps](http://www.gordyscamerastraps.com/neck-string-double/index.htm).

In Motion Panorama mode, banding appears sometimes. This could be the result of high shutter speeds and it is recommended to not use a shutter that is faster than somewhere in the 1/250s. Instead, use a narrower aperture size and lower the ISO if necessary. Unfortunately, the built-in ND filter cannot be used in panorama mode. Also, since  exposure is locked throughout the 120°/180° sweep, there will be banding if the lighting situation changes. Avoid pointing the camera at the sun and use a lens hood (or a step down adapter ring).

## Other X100 resources

Make sure to check out some of my other X100 articles:

- [Fujifilm Finepix X100 – a year in retrospect](2012-06-12-the-fujifilm-x100-a-year-in-retrospect.md)
- [Fujifilm Finepix X100 – RAW and film simulation](2012-08-04-fujifilm-finepix-x100-raw-and-film-simulation.md)
- [Fujifilm Finepix X100 – LCD/EVF observations](2012-08-09-fujifilm-finepix-x100-lcd-evf-observations.md)