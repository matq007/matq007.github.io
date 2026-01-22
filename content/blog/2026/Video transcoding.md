---
title: Video transcoding
date: 2026-01-18T12:00:00
tags:
  - blog/notes
---
Thought out my new year cleanup, I discovered that some of my movies in (`mkv`) are not discovered by [Jellyfin](https://jellyfin.org). I'm still not quite sure what might be the cause, but it was perfect timing because I wanted to look more into ripping some movies with [ffmpeg](https://www.ffmpeg.org). There are plethora settings, codecs, encoders to choose from, just have a look at reddit `r/ffmpeg`. Below is my good enough info for now

1. Use `H.265` encoder, it's newer, generates smaller size than `H.264`
2. For sounds use it's recommended to use `AV1`

or if you don't want to deal with this, just use [HandBrake](https://handbrake.fr) which already has well predefined settings. Note for Mac users, use hardware acceleration settings in `Presents` -> `Hardware` -> `H.265` (either 4k or 1080p).