FROM ubuntu

COPY src/entrypoint.sh /entrypoint
COPY src/ubuntu_init.sh /ubuntu_init.sh

RUN \
  chmod 700 /entrypoint && \
  chmod 700 /ubuntu_init.sh && \
  apt-get update && \
  # Install programs
  DEBIAN_FRONTEND=noninteractive apt-get install openssh-server -y && \
  apt-get install sudo unzip netcat zip p7zip-full python3-pip iproute2 net-tools cron cargo tmux tcpdump wget curl git vim build-essential -y && \
  pip3 install pandas && \
  pip3 install tabulate && \
  pip3 install pyyaml && \
  pip3 install ipython && \
  # Add users
  useradd -m sadeli && \
  useradd -m tayari && \
  useradd -m ubuntu && \
  # Change default shells for each user
  chsh -s /usr/bin/bash sadeli && \
  chsh -s /usr/bin/bash tayari && \
  chsh -s /usr/bin/bash ubuntu && \
  # Give each user a random password
  echo "sadeli:$(< /dev/urandom tr -dc _A-Z-a-z-0-9 | head -c${1:-32})" | chpasswd && \
  echo "tayari:$(< /dev/urandom tr -dc _A-Z-a-z-0-9 | head -c${1:-32})" | chpasswd && \
  echo "ubuntu:$(< /dev/urandom tr -dc _A-Z-a-z-0-9 | head -c${1:-32})" | chpasswd && \
  echo "root:$(< /dev/urandom tr -dc _A-Z-a-z-0-9 | head -c${1:-32})" | chpasswd && \
  # Privesc vimtutor 
  echo 'sadeli ALL=(tayari) NOPASSWD: /usr/bin/vimtutor' >> /etc/sudoers

# Upload Guessing Game
COPY src/guessing_game /var/backups/guessing_game
RUN \
  chown -R ubuntu:ubuntu /var/backups/guessing_game

# Start the git repo
USER ubuntu:ubuntu
WORKDIR /var/backups/guessing_game
RUN \
  mkdir /home/ubuntu/.ssh/ && \
  chmod 700 /home/ubuntu/.ssh && \
  git init && \
  git add -A && \
  git -c user.name='ubuntu' -c user.email='ubuntu@devilsec.org' commit -m 'First commit' && \
  git checkout -b dev

# Add sadeli's ssh key
USER sadeli:sadeli
RUN \
  mkdir /home/sadeli/.ssh/ && \
  chmod 700 /home/sadeli/.ssh
COPY keys/id_rsa.pub /home/sadeli/.ssh/authorized_keys

# Add the cereal thing
USER root
WORKDIR /
COPY src/cereal /opt/cereal
RUN \
  chown -R root:ubuntu /opt/cereal && \
  chmod 770 /opt/cereal && \
  chmod 0660 /opt/cereal/cereal.yaml && \
  chmod 0540 /opt/cereal/main && \
  echo '*/2 * * * * /opt/cereal/main' > /var/spool/cron/crontabs/root && \
  chown root:crontab /var/spool/cron/crontabs/root && \
  chmod 600 /var/spool/cron/crontabs/root

# Upload flags
COPY flags/sadeli.flag /home/sadeli/sadeli.flag
COPY flags/tayari.flag /home/tayari/tayari.flag
COPY flags/ubuntu.flag /home/ubuntu/ubuntu.flag
COPY flags/root.flag   /root/root.flag
RUN \
  chown root:sadeli /home/sadeli/sadeli.flag &&\
  chown root:tayari /home/tayari/tayari.flag && \
  chown root:ubuntu /home/ubuntu/ubuntu.flag && \
  chown root:root /root/root.flag && \
  chmod 0440 /home/sadeli/sadeli.flag && \
  chmod 0440 /home/tayari/tayari.flag && \
  chmod 0440 /home/ubuntu/ubuntu.flag && \
  chmod 0440 /root/root.flag

ENTRYPOINT [ "/entrypoint" ]
