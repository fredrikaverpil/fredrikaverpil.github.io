---
title: Fujifilm X100S camera profiles and presets
tags: [photography]
draft: false

cover:
  image: "/static/x100s/xrite_passport.jpg"
  alt: "The XRite Passport"
  relative: false # To use relative path for cover image, used in hugo Page-bundles

# PaperMod
ShowToc: false
TocOpen: false

date: 2013-06-16T02:00:12+02:00
---

### Background

Ever since I replaced my trusty old X100, see [[2012-06-12-the-fujifilm-x100-a-year-in-retrospect]], with its newer incarnation, the X100S (see [[2013-03-08-the-fujifilm-x100s-compared-to-the-x100]]), I have been struggling with color rendition in Lightroom, which is my weapon of choice when editing RAW files.

<strong>Update 8<sup>th</sup> April, 2014:</strong> Adobe today released Adobe Camera RAW 8.4 which includes huge improvements on color rendition in Lightroom 5.4 and Photoshop CC (version 14.2.1). This more or less renders my camera profiles obsolete. But do read on if you like…

ACR 8.4 enables the best color rendition achievable in Lightroom by utilizing custom camera profiles (Provia, Astia Velvia, ProNeg, Monochrome etc). Definitively go check this out right now!

My presets further down were created with my custom camera profile “Fujifilm X100S (Sun only)” in mind. However, they can be used with the new camera profiles supplied in Lightroom 5.4. Just go down to the Camera Calibration pane and change “Fujifilm X100S (Sun only)” to e.g. “Camera PROVIA/STANDARD”.

A while ago, [Apple released an update to their Digital Camera RAW software](https://support.apple.com/kb/DL1629?locale=en_US) (native RAW support in OS X) with support for the X100S and I found its color rendition of X100S RAW files much nicer right out of the box than what Adobe was offering me.

For some time, I exported 16-bit TIF files from Apple’s Preview.app and imported them into Lightroom and continued to edit my photos in there, but this quickly became cumbersome and consumed quite a lot of extra disk space.

I placed a request in the Adobe forums to address this bad rendition of X100S colors and got a reply from a fellow user that pointed out to me that this probably wasn’t Adobe’s fault to begin with, since they didn’t bundle a specific camera profile for the X100S, and that I could create my own profile using a standardized color checker, which would probably get the colors right. Enough said, [I ordered one](http://xritephoto.com/colorchecker-passport-photo), made my own profile and here we are. Today I do not feel the need to do the detour through Apple’s Preview.app just in order to get a good starting point with colors and I can keep an undestructive workflow within Lightroom.

Feel free to try it out, but keep in mind this profile was made for my camera and that custom profiles may differ a bit from camera to camera. I’ve also included a bunch of presets I use frequently together with the X100S.


### Downloads

Please note, bookmark this page, put it on [change detection](http://www.changedetection.com) or follow me [on GitHub](https://github.com/fredrikaverpil/photography) as I am bound to update this page with additional profiles and perhaps presets now that I own a color checker!

#### Camera profiles

* [Fujifilm X100S (Sun only)](https://github.com/fredrikaverpil/photography/raw/master/Lightroom/CameraProfiles/Fujifilm%20X100S%20(Sun%20only).dcp) - (D55)
* [Fujifilm X100S (Sun and Fluorescent)](https://github.com/fredrikaverpil/photography/raw/master/Lightroom/CameraProfiles/Fujifilm%20X100S%20(Sun%20and%20Fluorescent).dcp) – Dual illuminant (D55/A)
* [Fujifilm X100S (Sun and Tungsten)](https://github.com/fredrikaverpil/photography/raw/master/Lightroom/CameraProfiles/Fujifilm%20X100S%20(Sun%20and%20Tungsten).dcp) – Dual illuminant (D55/A)

The dual illuminant camera profiles were made from two photos (shot in sunlight & low light in either fluorescent or tungsten light) of the color checker chart and processed in X-Rite’s ColorChecker Passport 1.0.2. I did not find Adobe DNG Profile Editor 1.0.0.46 beta to produce as pleasing color tones when comparing the resulting camera profiles from each application, so that’s the reason behind using the X-Rite approach.

The sunlight shot was taken in direct sunlight as this resulted in much more pleasing color tones (compared to when using a daylight shot). This also made the color rendition resemble the neutral Provia film simulation.

The fluorescent shot and the tungsten shot were both taken in a dark room with each light source activated respectively.

### Presets


[![](/static/x100s/presets.jpg)](https://www.flickr.com/photos/fredrik/9051954579/)  
*Presets comparison. Click image for larger version.*


* [X100S Fredrik – Film](https://github.com/fredrikaverpil/photography/raw/master/Lightroom/Presets/Legacy/X100S%20Fredrik%20-%20Film.lrtemplate) – cool film look
* [X100S Fredrik – Film (vivid)](https://github.com/fredrikaverpil/photography/raw/master/Lightroom/Presets/Legacy/X100S%20Fredrik%20-%20Film%20(vivid).lrtemplate) – warm and vivid film look
* [X100S Fredrik – Neutral](https://github.com/fredrikaverpil/photography/raw/master/Lightroom/Presets/Legacy/X100S%20Fredrik%20-%20Neutral.lrtemplate) – a good starting point
* [X100S Fredrik – Neutral (faded)](https://github.com/fredrikaverpil/photography/raw/master/Lightroom/Presets/Legacy/X100S%20Fredrik%20-%20Neutral%20(faded).lrtemplate) – trying to mimic Astia
* [X100S Fredrik – Neutral (punchy)](https://github.com/fredrikaverpil/photography/raw/master/Lightroom/Presets/Legacy/X100S%20Fredrik%20-%20Neutral%20(punchy).lrtemplate) – a bit more contrasty

**Update:** Newer, more up to date presets are found [here](https://github.com/fredrikaverpil/photography/tree/master/Lightroom/Presets).

Please note, you may want to download the presets by right-clicking the links and choosing the “save as...” option. All presets use the camera profile “Fujifilm X100S (Sun only)” and were made with Lightroom 5.



### Installation instructions

#### Mac OS X

* Copy the camera profiles into `/Users/[your_username]/Library/Application Support/Adobe/CameraRaw/CameraProfiles`
* Copy the presets into `/Users/[your_username]/Library/Application Support/Adobe/Lightroom/Develop Presets/User Presets`

#### Windows

* Copy the camera profiles into `C:\Users\[your_username] \AppData\Roaming\Adobe\CameraRaw\CameraProfiles`

If anyone knows where to put the presets on a Windows machine, let me know (I believe you can drag and drop them into Lightroom?)

### Usage instructions

In Lightroom and in the Develop module, scroll down to the “Camera Calibration” pane. Here, change “Adobe Standard” to either one of the camera profiles.

If you like the camera profiles better than Adobe’s standard offering, you can make Lightroom automatically use one of them for all X100S RAF files. This way, you do not have to apply it on each and every RAW file manually. Open any X100S RAF file in the “develop” module, reset all edits and just change the profile under “Camera Calibration” to e.g. the “Fujifilm X100S (Sun only)” profile. Then in the file menu, go into the Develop menu and choose “Set Default Settings…”. Please note that after having done this, even if you hit “reset”, this custom camera profile will be used.

Please note that if you switch between camera profiles you may have to reset the white balance of your shot.

If you downloaded any of my presets, you just need to click the preset and it will automatically change the camera profile into the “Fujifilm X100S (Sun only)” one.

The presets will override all your settings except the white balance, exposure and lens corrections. If you want the presets not to touch other settings, just apply the preset to a photo and then right-click the preset and choose to update it with the current settings. Then leave out the settings you do not wish the preset to affect.
