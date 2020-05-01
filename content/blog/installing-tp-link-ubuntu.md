---
layout: post
title: "Driver for tp-link AC1300 on Ubuntu 20.04"
date: 2020-05-01
tags: [tp-link, configuration, ubuntu]
layout: "post"
---

```bash
sudo su
sudo apt install build-essentials dkms
cd /usr/local/share/
wget https://github.com/EntropicEffect/rtl8822bu/archive/9438d453ec5b4878b7d4a2b4bcf87b37df09c3fb.zip -O rtl8822bu.zip
unzip rtl8822bu.zip
cd rtl8822bu-9438d453ec5b4878b7d4a2b4bcf87b37df09c3fb
dkms add .
dkms install -m 88x2bu -v 1.1
reboot
```

Credit goes to [madscintist](https://ubuntuforums.org/archive/index.php/t-2426469.html).

## Reference

- [Forum thread](https://ubuntuforums.org/archive/index.php/t-2426469.html)
- [GitHub: EntropicEffect/rtl8822bu](https://github.com/EntropicEffect/rtl8822bu)