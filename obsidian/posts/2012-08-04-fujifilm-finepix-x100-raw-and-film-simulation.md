---
title: Fujifilm Finepix X100 – RAW and film simulation
tags: [photography]
draft: false

cover:
  image: "/static/x100/x100_filmsim_01.jpg"
  alt: "Close-up of camera's LCD screen"
  relative: false # To use relative path for cover image, used in hugo Page-bundles

# PaperMod
ShowToc: false
TocOpen: false

date: 2012-08-04T02:00:12+02:00
---

– Why film simulations are not disabled when shooting RAW?



A RAW file is a photo that hasn’t been developed yet and simply consist of numerical values representing the amount of light that has been recorded by the sensor (in the case of the X100 this is 12-bit per pixel). It has to be somehow converted into RGB in order to be viewable on a computer screen (or the camera’s LCD/EVF) as an image that you can look at. This can be done through RAW converters or image viewers which know how to interpret the RAW file into an RGB image.

So, when you choose e.g. the film simulation “Astia”, this is the film simulation (or RAW converter’s formula to convert into RGB) used in order to present an image to you in the LCD/EVF and when playing back the images you have shot (see “Preview Image in the EXIF” further down below). The camera needs to do this in order to let you know what you are shooting and what you have shot. Otherwise you would just see 12.37 million different numerical values of light, and that wouldn’t help to know if your shot was in focus.

The film simulation is not baked into the RAW file that is recorded to your memory card in any other way than the film simulation’s name (in this case “Astia”). Therefore applications such as Adobe Lightroom will never be able to reproduce Fujifilm’s film simulations automatically upon import of a RAW file, unless they know Fuji’s secret formula for the specific film simulation – and they don’t.

Try changing the film simulation to “Monochrome” and you’ll see that the LCD/EVF turns grayscale. This however does not mean you are recording an grayscale-only image if you are shooting RAW. However, if you are shooting JPG-only, all RAW data is converted to RGB, then the RAW data gets discarded and there will be only a grayscale RGB JPG image. When a RAW file shot with the “Monochrome” film simulation is imported into e.g. Adobe Lightroom or Adobe Photoshop the image will be in color and not in grayscale/monochrome. This is only because ACR supports X100/X-Pro1 RAW files and does a pretty decent job in showing you what it think that the photo should look like with its RGB conversion formula (it’s not Provia, it’s not Astia, it’s … Adobes own film simulation made for the X100, if you will – only difference here being you can go in and tweak that formula with sliders and buttons.

## The in-camera converter

Since you have shot the photo in RAW, you can always launch the built-in RAW-processor software in-camera and change the film simulation into something else and then “develop” the photo into a JPG using other settings than you had when shooting the image. The RAW file, however, will never be altered when doing this and will always remain as RAW data. This also applies to a number of other settings such as the tone of shadow and highlights. All of that can be changed after having shot the RAW photo.

## Sample X100 EXIF data

EXIF extracted with [ExifTool](http://www.sno.phy.queensu.ca/~phil/exiftool/).

    ExifTool Version Number         : 8.99
    File Name                       : DSCF6164.RAF
    Directory                       : /Users/fredrik/Pictures/RAW/2012/2012-07-12
    File Size                       : 19 MB
    File Modification Date/Time     : 2012:07:12 12:12:55+02:00
    File Permissions                : rw-------
    File Type                       : RAF
    MIME Type                       : image/x-fujifilm-raf
    RAF Version                     : 0130
    Exif Byte Order                 : Little-endian (Intel, II)
    Make                            : FUJIFILM
    Camera Model Name               : FinePix X100
    Orientation                     : Horizontal (normal)
    X Resolution                    : 72
    Y Resolution                    : 72
    Resolution Unit                 : inches
    Software                        : Digital Camera FinePix X100 Ver1.30
    Modify Date                     : 2012:07:12 12:12:55
    Y Cb Cr Positioning             : Co-sited
    Copyright                       :
    Exposure Time                   : 1/240
    F Number                        : 8.0
    Exposure Program                : Aperture-priority AE
    ISO                             : 200
    Sensitivity Type                : Standard Output Sensitivity
    Exif Version                    : 0230
    Date/Time Original              : 2012:07:12 12:12:55
    Create Date                     : 2012:07:12 12:12:55
    Components Configuration        : Y, Cb, Cr, -
    Compressed Bits Per Pixel       : 2
    Shutter Speed Value             : 1/246
    Aperture Value                  : 8.0
    Brightness Value                : 13.42
    Exposure Compensation           : 0
    Max Aperture Value              : 2.0
    Metering Mode                   : Multi-segment
    Light Source                    : Unknown
    Flash                           : Off, Did not fire
    Focal Length                    : 23.0 mm
    Version                         : 0130
    Internal Serial Number          : FPX 20850584     592D36323234 2012:02:23 FE613f110062
    Quality                         : NORMAL
    Sharpness                       : Film Simulation
    White Balance                   : Auto
    Saturation                      : Normal
    White Balance Fine Tune         : Red +0, Blue +0
    Noise Reduction                 : n/a
    High ISO Noise Reduction        : Unknown (0x280)
    Fuji Flash Mode                 : Off
    Flash Exposure Comp             : -0.67
    Macro                           : Off
    Focus Mode                      : Auto
    Focus Pixel                     : 1088 724
    Slow Sync                       : Off
    Picture Mode                    : Aperture-priority AE
    Auto Bracketing                 : Off
    Sequence Number                 : 0
    Blur Warning                    : None
    Focus Warning                   : Good
    Exposure Warning                : Good
    Dynamic Range                   : Standard
    Film Mode                       : F1b/Studio Portrait Smooth Skin Tone (Astia)
    Dynamic Range Setting           : Manual
    Development Dynamic Range       : 100
    Faces Detected                  : 0
    Flashpix Version                : 0100
    Color Space                     : sRGB
    Exif Image Width                : 2176
    Exif Image Height               : 1448
    Interoperability Index          : R98 - DCF basic file (sRGB)
    Interoperability Version        : 0100
    Focal Plane X Resolution        : 924
    Focal Plane Y Resolution        : 924
    Focal Plane Resolution Unit     : cm
    Sensing Method                  : One-chip color area
    File Source                     : Digital Camera
    Scene Type                      : Directly photographed
    Custom Rendered                 : Normal
    Exposure Mode                   : Auto
    Scene Capture Type              : Standard
    Subject Distance Range          : Unknown
    PrintIM Version                 : 0250
    Compression                     : JPEG (old-style)
    Thumbnail Offset                : 1964
    Thumbnail Length                : 9030
    Image Width                     : 2176
    Image Height                    : 1448
    Encoding Process                : Baseline DCT, Huffman coding
    Color Components                : 3
    Y Cb Cr Sub Sampling            : YCbCr4:2:2 (2 1)
    Preview Image                   : (Binary data 764732 bytes, use -b option to extract)
    Raw Image Full Size             : 4448x2870
    Fuji Layout                     : 10 11 9 8
    Raw Image Width                 : 4448
    Raw Image Height                : 2870
    Bits Per Sample                 : 12
    Strip Offsets                   : 789200
    Strip Byte Counts               : 19148640
    WB GRB Levels                   : 302 450 447
    Aperture                        : 8.0
    Blue Balance                    : 1.480132
    Image Size                      : 2176x1448
    Red Balance                     : 1.490066
    Scale Factor To 35 mm Equivalent: 1.5
    Shutter Speed                   : 1/240
    Thumbnail Image                 : (Binary data 9030 bytes, use -b option to extract)
    Circle Of Confusion             : 0.020 mm
    Field Of View                   : 54.2 deg
    Focal Length                    : 23.0 mm (35 mm equivalent: 35.2 mm)
    Hyperfocal Distance             : 3.37 m
    Light Value                     : 12.9

## Advanced control during RAW development

I can recommend that everyone shooting RAW take a look at [RAW Photo Processor](http://www.raw-photo-processor.com/). It’s offering an extremely eccentric control over the RAW development process. It’s completely free but unfortunately I believe it is only available on Mac OS X.

## Other X100 resources

Make sure to check out some of my other X100 articles:

Fujifilm Finepix X100 – settings: [[2012-02-02-fujifilm-x100-settings-and-notes]]
- Fujifilm Finepix X100 – a year in retrospect: [[2012-06-12-the-fujifilm-x100-a-year-in-retrospect]]
- Fujifilm Finepix X100 – LCD/EVF observations: [[2012-08-09-fujifilm-finepix-x100-lcd-evf-observations]]