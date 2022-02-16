FROM ubuntu

COPY src/copy-thingy.c /root/copy-thingy.c
COPY src/entrypoint /entrypoint
COPY src/backup/ /backup/

# Machine setup
RUN apt-get update && \
    # Install programs
    DEBIAN_FRONTEND=noninteractive apt-get install openssh-server -y && \
    apt-get install sudo unzip zip iproute2 net-tools cron tmux tcpdump wget curl git vim build-essential -y &&\
    # Add users
    useradd -m sadeli && \
    useradd 7StringsOvPasta && \
    useradd -m tayari && \
    # Give sadeli perms to run vim as 7StringsOvPasta
    echo 'sadeli ALL=(7StringsOvPasta) NOPASSWD: /usr/bin/vim' >> /etc/sudoers && \
    # Create SSH key for root
    ssh-keygen -q -t rsa -N '' -f /root/.ssh/id_rsa && \
    cp /root/.ssh/id_rsa.pub /root/.ssh/authorized_keys && \
    # Compile copy-thingy script
    mkdir -p /opt/script/files && \
    gcc /root/copy-thingy.c -o /opt/script/copy-thingy && \
    rm /root/copy-thingy.c && \
    chown -R root:tayari /opt/script/ && \
    chmod o-rwx /opt/script && \
    chmod o-rwx /opt/script/files && \
    chmod g+w /opt/script/files && \
    chmod 740 /opt/script/copy-thingy && \
    # Add cron job to run copy-thingy script
    echo '* * * * * /opt/script/copy-thingy' > /var/spool/cron/crontabs/root && \
    chown root:crontab /var/spool/cron/crontabs/root && \
    chmod 600 /var/spool/cron/crontabs/root && \
    # Give sadeli and tayari default bash shell
    chsh -s /bin/bash sadeli && \
    chsh -s /bin/bash tayari && \
    # Give each user a random password
    echo "sadeli:$(< /dev/urandom tr -dc _A-Z-a-z-0-9 | head -c${1:-32})" | chpasswd && \
    echo "7StringsOvPasta:$(< /dev/urandom tr -dc _A-Z-a-z-0-9 | head -c${1:-32})" | chpasswd && \
    echo "tayari:$(< /dev/urandom tr -dc _A-Z-a-z-0-9 | head -c${1:-32})" | chpasswd && \
    echo "root:$(< /dev/urandom tr -dc _A-Z-a-z-0-9 | head -c${1:-32})" | chpasswd && \
    chmod 700 /entrypoint

# SSH keys and flags
COPY keys/sadeli_rsa.pub /home/sadeli/.ssh/authorized_keys
COPY flags/sadeli.flag /home/sadeli/sadeli.flag
COPY flags/7StringsOvPasta.flag /backup/7StringsOvPasta.flag
COPY flags/tayari.flag /home/tayari/tayari.flag
COPY flags/root.flag    /root/root.flag

RUN mkdir /home/tayari/.ssh && \
    # Create ssh keys for tayari
    ssh-keygen -q -t rsa -N '' -f /home/tayari/.ssh/id_rsa && \
    cp /home/tayari/.ssh/id_rsa.pub /home/tayari/.ssh/authorized_keys && \
    chown -R sadeli:sadeli /home/sadeli/.ssh && \
    chown -R tayari:tayari /home/tayari/.ssh && \
    # Set perms on ssh keys and flags
    chown root:sadeli /home/sadeli/sadeli.flag && \
    chown root:tayari /home/tayari/tayari.flag && \
    chown root:root /root/root.flag && \
    chmod 700 /home/sadeli/.ssh && \
    chmod 700 /home/tayari/.ssh && \
    chmod 440 /home/sadeli/sadeli.flag && \
    chmod 440 /home/tayari/tayari.flag && \
    chmod 400 /root/root.flag && \
    # Prepare backup.zip
    cp /home/tayari/.ssh/id_rsa /backup/id_rsa && \
    git init /backup/ && \
    git -C /backup/ add -A && \
    git -C /backup/ -c user.name='tayari' -c user.email='tayari@devilsec.org' commit -m 'First commit' && \
    rm /backup/id_rsa && \
    git -C /backup/ add -A && \
    git -C /backup/ -c user.name='tayari' -c user.email='tayari@devilsec.org' commit -m 'OOPS had to delete something for security reasons' && \
    mkdir /home/sadeli/.backups && \
    zip -r /home/sadeli/.backups/backup.zip /backup/ && \
    rm -rf /backup && \
    chown -R 7StringsOvPasta:7StringsOvPasta /home/sadeli/.backups/ && \
    chmod o-rwx /home/sadeli/.backups/ && \
    chmod 660 /home/sadeli/.backups/backup.zip

# Start ssh and cron on startup
ENTRYPOINT [ "/entrypoint" ]
