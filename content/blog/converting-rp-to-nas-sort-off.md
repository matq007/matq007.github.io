---
layout: post
title: "Converting RP to NAS, sort off"
date: 2017-08-17
tags: [raspberry-pi, samba, ufw, ssh]
layout: "post"
---

Approximately two years ago I've been playing around with an idea of purchasing a NAS (Network attached storage). As a student looking at the prices, I've decided to make my own budget one. How hard can it  right?
I saw some Youtube videos where people were building NAS using Raspberry Pi (RP). So I've started digging for some Pi machine. I ended up choosing between Banana Pi (BP) and RP 2. The advantage of BP was that it had SATA connector right on the board, which would give me better read/write throughput. Ok, let's try it out.

## Round 1

I found a BP Pro on site alza.cz and later on found out that it was actually a BP M1+. Never mind, the hardware configuration was the same. I've downloaded and installed Raspbian OS. The installation is pretty straight forward and can be found on the page or youtube.
Ok, the installation was successful. Now let's try to connect my external drive. For storage, I've used a 2.5" disk from my old notebook. Normally, if you want to connect a drive like this, you have to have an external adapter to power it up. With BP you don't need any power adapter, the power is provided by the board when you connect the drive (Cha ching!). After connecting the disk, it's time to mount it.
Unfortunately, no new device was detected. I guess I really need an external power adapter. Tried the adapter, still nothing. Ok, now this is getting frustrating. Maybe the disk is not working properly, let's connect it to my work station. The disk was detected immediately. I've spent the whole day trying to connect the drive to the BP, but with no luck.
No worries, I can still buy RP 2 and try that instead. I've returned the BP and ordered RP 2. After one week it arrived and I could continue with my experiment.

## Round 2

The process was the same as before. Connect all the cables to RP and install the latest Rasbian OS. However this time, there was no SATA connector. That's why I ordered a case for 2.5" hard drive with USB 3.0. Yes, I know RP 2 has only USB 2.0 but I can reuse this case later in future for something else. But wait, you said that external drives have to have an external power supply. How are you going to power the drive if it's already in a case? The cool part about RP is that you can play with the configuration. In order to connect an external drive, you have to modify config.txt file and add this line max_usb=1. With this setting, RP will use more voltage for USB, thus making it possible connecting an external drive. Ready for a test run? Yes, it worked! Mission successful!
Configuration
Finally, we can move on to the next step. Installing Samba for sharing and ufw for creating firewall rules. Let's not forget to also setup SSH (I would like to disconnect the keyboard and monitor that are now connected to the RP).

```bash
# Update your /etc/fstab to mount the drive on boot
# Installation and update
sudo apt-get update -y && sudo apt-get upgrade
sudo apt-get install samba ufw
# Setting up Samba
# Create user on raspberry:
sudo useradd <USERNAME> - shell /bin/false
# Set permissions:
sudo chown <USERNAME> <PATH_TO_DRIVE>
# Create password for samba:
sudo smbpasswd -a <USERNAME>
# Update configuration
vi /etc/samba/smb.conf
[SHARE]
path = /mnt/share/
valid users = <USERNAME>
read only = No
create mask = 0777
directory mask = 0777
# Save and exit vim (!x or !wq)
# Verify the correctness of configuration
testparm
# Start samba
sudo systemctl start samba
# Check that status, you should see green Active
sudo systemctl status samba
# Auto start samba
sudo systemctl enable samba
# Setting up firewall rules
# I didn't have Samba listed in apps so I created my own 
vi /etc/ufw/applications.d/ufw-fileserver
[samba]
title=Samba
description=Samba
ports=137,138,139,445/tcp
# After saving it's time to setup some firewall rules
# 1. I want to access ssh but I wan't to limit it
# 2. I want to access Samba only on my local network
sudo ufw limit ssh
sudo ufw allow from XXX.XXX.X.0/16 to any app Samba
sudo ufw enable
# Check the status that it works
sudo ufw status
```

Backup the most important data on USB stick
I know, I know, this is insane and bold. You should definitely not do this at home! I didn't have another drive available, so I found some 4GB USB stick, which I used as my secondary drive, where I've backed up only the most important documents that my parents use.

```bash
# Update /etc/fstab to mount the USB stick
# Create a backup script
vi ~/backup
#!/bin/bash
rsync -av --delete <DRIVE_PATH>/Important/ <BACKUP_PATH>/Important
# Make the file executable and run it as a cron job 
# every day at 1.00am
chmod +x backup
crontab -e
0 1 * * * ~/backup
```

I've setup one more backup on my desktop station. I am using Free FileSync application for backing up data. I'm using a real-time backup with mirroring option. I've setup one more rule to check my share folder and copy again only the most important files to the backup drive to my desktop station. This way if everything fails, I should still have some backup.
How about some small scripts?
I've played around with the first setup and created some small scripts like a welcome page, sending notifications to slack on login, etc. You can find some of them here and here.

## Conclusion & Summary

The project started with an idea of using SATA connector for connecting the drive and sharing it locally. That quite didn't happen, so I ended up switching to USB 2.0 on RP which has lower read/write speeds. Nevertheless, the end result was fine and the system actually worked.
Now hold on a minute, why would you go through so much trouble when you could have just connected the drive to your router and share it?
I could have done it right after I discovered problems with the BP. But I was curious how it would work with RP. Of course, this is not really a NAS (well sort off is), but the fun and learning outcome was worth it. To be honest, I've been running this setup at home for over a year without any problem what so ever.
This was a very interesting project with many ups and downs. I don't encourage you to build the same setup. I've just wanted to demonstrate what can be done by reusing an old hard drive and some free time. If you have any questions or found some errors, let me know in the comments below.

## Sources

- [SSH](https://debian-administration.org/article/530/SSH_with_authentication_key_instead_of_password)
- [samba](https://help.ubuntu.com/community/How%20to%20Create%20a%20Network%20Share%20Via%20Samba%20Via%20CLI%20%28Command-line%20interface/Linux%20Terminal%29%20-%20Uncomplicated%2C%20Simple%20and%20Brief%20Way%21)
- [ufw](https://help.ubuntu.com/community/UFW)
