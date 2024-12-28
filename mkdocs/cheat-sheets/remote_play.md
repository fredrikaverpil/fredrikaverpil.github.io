---
date: 2024-12-28
draft: true
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

Disconnecting the HDMI cable between TV and PS5 or just disconnecting the TV
from the network is a workaround for the problem.

To permanently fix this issue, activate the "Simple IP Control" in the TV's
network settings. This will cause the TV's networking to be always-on instead of
hibernating and pinging for activity, which will also consume more energy.

### Apple device network pings

iOS has a tendency to exhibit an intermittent increase in ping on 5ghz networks
if AirDrop is on and other ios devices are nearby. This negatively affects all
zero framebuffer applications including xbox remote play and the ps remote play
app. Therefore, disable AirDrop on your iOS device you are using for Remote
Play.

## PS button via 3rd party app

- https://streamingdv.com/shop-list-ns.html
- iOS download:
  https://apps.apple.com/my/app/mirrorplay-remote-streaming/id1638586503
- Android download:
  https://play.google.com/store/apps/details?id=psplay.grill.com
