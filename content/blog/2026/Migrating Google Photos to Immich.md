---
title: Migrating Google Photos to Immich
date: 2026-01-11T12:00:00
tags:
  - blog/notes
  - immich
  - blog/homelab
---
> TL;DR: This is a summary post of my experience migrating Google Photos to immich. Did you setup your reverse-proxy correctly? Spoiler alert, I didn't 🤦‍♂️

1. Navigate to [Google Takeout](https://takeout.google.com) and request all your Google Photos images. I suggest to use `tar` compression.
2. Download fantastic tool [immich-go](https://github.com/simulot/immich-go), which will import all photos including metadata, albums etc.
3. Generate an API KEY with specified permissions. For simplicity I've generated a key with all permissions which I've immediately deleted after import was finished.
	```bash
	# Import command from documentation
	./immich-go upload from-google-photos \
		--server=http://10.0.0.103:2283 \ # only use hostname if you've correctly configured reverse-proxy
		--api-key=<API_KEY> \
		~/Downloads/Takeout/Google\ Photos/
	```
4. After successful import I've used [shtse8/Google-Photos-Delete-Tool](https://github.com/shtse8/Google-Photos-Delete-Tool) to delete all Google Photos, because the UI doesn't support selecting all photos in one go. I've simply pasted the JS code in console and waited until all images were automatically selected and then clicked delete button.

## Troubleshooting random import crashes

If like me ([and many others](https://www.reddit.com/r/immich/comments/1ppzkr2/how_to_deal_with_large_files/)), you haven't properly configured your reverse-proxy, you probably experienced a random crash during importing. My initial thoughts were performance issue as described [here](https://tsmith.com/blog/2025/immich-migration/), however the error message was hinting on something else

```log
io: read/write on closed pipe
2026-01-10 21:45:21 ERR Error err=AssetUpload, POST, https://immich.proks.top/api/assets, 500 Internal Server Error
```

When I've restarted the upload, I've peed on my reverse-proxy ([Nginx Proxy Manager](https://nginxproxymanager.com)) and behold, my storage was just growing and growing while uploading a 4GB mp4 video. Well that's odd, why would my local storage grow, is it buffering the files before sending them to immich? Behold, the immich's [documentation](https://docs.immich.app/administration/reverse-proxy/#nginx-example-config) couldn't be more clear 🤦‍♂️

```i
# disable buffering uploads to prevent OOM on reverse proxy server and make uploads twice as fast (no pause)  
proxy_request_buffering off;
```

If you run the command with an IP address instead of hostname, you've probably haven't experienced this issue, but if you are still fighting with your import, you might want to double check your reverse-proxy settings.