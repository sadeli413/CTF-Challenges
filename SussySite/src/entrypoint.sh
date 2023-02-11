#!/bin/bash

cd /var/www/api && sudo -u www-data ./api &
cd /var/www/logger && sudo -u www-data ./logger &
cd /var/www/sussychat && sudo -u www-data ./sussychat &
cd /var/www/loginportal && _SECRET='BG6sIpT4hVjUbn8aex39hdSCQgQ53m5j' ./loginportal.py &
service nginx start && \
service ssh start && \

while true; do
  /usr/bin/sleep 2m
  /usr/bin/rm /var/www/sussychat/uploads/*
done
