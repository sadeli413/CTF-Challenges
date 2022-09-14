#!/bin/bash

/usr/sbin/service ssh restart && \
/usr/sbin/service cron restart

chmod 777 /ubuntu_init.sh
su ubuntu -c /ubuntu_init.sh
rm /ubuntu_init.sh

7z a -p'butterfly1' /var/backups/guessing_game.7z /var/backups/guessing_game
rm -rf /var/backups/guessing_game
chown ubuntu:tayari /var/backups/guessing_game.7z
chmod 0440 /var/backups/guessing_game.7z

/usr/bin/tail
