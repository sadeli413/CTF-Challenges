FROM ubuntu

COPY src/entrypoint.sh /entrypoint

RUN \
  # Install stuff
  apt-get update && \
  DEBIAN_FRONTEND=noninteractive apt-get install openssh-server -y && \
  apt-get install sudo file unzip netcat zip iproute2 python3 python3-pip net-tools cron tmux tcpdump wget curl git vim build-essential nginx -y &&\
  pip3 install ipython && \
  pip3 install flask && \
  # Add users
  useradd -m sadeli && \
  chsh -s /usr/bin/bash sadeli && \
  echo 'sadeli:2ZiOHLxtxbI-E4rcOyL00ub18_daWVgi' | chpasswd && \
  echo 'root:BG6sIpT4hVjUbn8aex39hdSCQgQ53m5j' | chpasswd && \
  chmod 500 /entrypoint

# Upload web app code
COPY src/www/nginx_default /etc/nginx/sites-available/default
COPY src/www/sussychat /var/www/sussychat
COPY src/www/logger /var/www/logger
COPY src/www/api /var/www/api
COPY src/www/api.tar.gz /var/backups/api.tar.gz
COPY src/www/loginportal /var/www/loginportal
RUN \
  rm -r /var/www/html && \
  chown -R www-data:www-data /var/www/sussychat && \
  chown -R www-data:www-data /var/www/logger && \
  chown -R www-data:www-data /var/www/api && \
  chown sadeli:sadeli /var/backups/api.tar.gz && \
  chmod -R o-w /var/www/sussychat && \
  chmod -R o-w /var/www/logger && \
  chmod -R o-w /var/www/api && \
  chmod 0444 /var/backups/api.tar.gz && \
  chmod 0644 /etc/nginx/sites-available/default && \
  chmod 0644 /var/backups/api.tar.gz && \
  # Privesc app
  chown -R root:root /var/www/loginportal && \
  chmod 0555 /var/www/loginportal/ && \
  chmod 0554 /var/www/loginportal/loginportal.py

# Copy flags
COPY flags/user.flag /home/sadeli/user.flag
COPY flags/root.flag /root/root.flag
RUN \
  chown root:sadeli /home/sadeli/user.flag && \
  chmod 440 /home/sadeli/user.flag && \
  chown root:root /root/root.flag && \
  chmod 400 /root/root.flag

ENTRYPOINT [ "/entrypoint" ]
