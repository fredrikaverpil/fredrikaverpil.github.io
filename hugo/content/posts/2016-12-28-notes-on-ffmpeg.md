---
ShowToc: false
TocOpen: false
date: 2016-12-28 02:00:12+01:00
draft: false
tags:
- osx
- video
title: Notes on ffmpeg
---

Quick note on how to install ffmpeg (using Homebrew) and various mp4 conversion commands on Unix.



Full `ffmpeg` documentation available [here](https://ffmpeg.org/ffmpeg.html).


### Install ffmpeg for .webm/.mp4 conversion

```bash
# Uninstall any existing ffmpeg installation
brew uninstall ffmpeg

# Install (macOS using Homebrew)
brew install ffmpeg --with-fdk-aac --with-sdl2 --with-freetype --with-libass --with-libvorbis --with-libvpx --with-opus --with-x265

# Verify ffmpeg build configuration
ffmpeg -buildconf

# .webm -> .mp4
ffmpeg -i input.webm -codec copy output.mp4

# .mp4 -> .webm
ffmpeg -i input.mp4 -strict -2 output.webm
```


### Convert all files in the same directory to .mp4

```bash
# Replace [wildcard] with e.g. *webm
# Replace [options] with your options...
for i in [wildcard]; do ffmpeg -i $i [options] $i.mp4; done
```

### Merge .m4a with .mp4

```bash
ffmpeg -i video.mp4 -i audio.m4a -codec copy output.mp4
```

### Extract still image at given time from video

```bash
# Enter time in hh:mm:ss
# Choose file format by changing extension of output to e.g. .png
ffmpeg -ss 00:00:46 -i input.mp4 -vframes 1 output.jpg
```