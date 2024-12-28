---
date: 2024-12-28
draft: false
tags:
  - gaming
title: Remote Play
icon: simple/playstation
---

# Remote Play

## Solving severe stuttering during Remote Play sessions

### TV network pings

There are situations where the TV loses a packet every 20-30 pings. On each
packet loss, Remote Play hangs for a second.

Disconnecting the HDMI cable between the TV and the PS5 is a workaround for the
problem. This prevents any potential interference between the TV's HDMI-CEC
control signals and the PS5's network communication.

On Sony BRAVIA TVs, you can permanently fix this issue by enabling the "Simple
IP Protocol" under network settings. The Simple IP Protocol reduces network
complexity and overhead by using a more basic communication method, which helps
minimize packet loss and network latency between the TV and the console. This
does come at a cost, which is higher energy consumption.

Huge thanks to `u/NBL-83` who originally
[posted](https://www.reddit.com/r/remoteplay/comments/te9ut1/playstation_remoteplay_stops_hangs_every_30/)
about all this over at `r/remoteplay`.

### Apple device network pings

iOS has a tendency to exhibit an intermittent increase in ping on 5 GHz networks
if AirDrop is enabled and other ios devices are nearby. This negatively affects
all zero-framebuffer applications including the Xbox Remote Play and PS Remote
Play apps. Therefore, disable AirDrop on your iOS device during Remote Play
sessions.

## Mapping the PS5 touchpad

[Florian Grill](https://grill2010.github.io) has developed an excellent
alternative to the official Sony Remote Play app, in which you can map the
touchpad to a gamepad button:

- [MirrorPlay](https://apps.apple.com/my/app/mirrorplay-remote-streaming/id1638586503)
  (iOS, iPadOS)
- [PXPlay](https://play.google.com/store/apps/details?id=psplay.grill.com)
  (Android)
