#!/bin/bash

# Create ssh keys for ubuntu user and copy to guessing_game
ssh-keygen -q -t rsa -N '' -f /home/ubuntu/.ssh/id_rsa
cp /home/ubuntu/.ssh/id_rsa.pub /home/ubuntu/.ssh/authorized_keys
cp /home/ubuntu/.ssh/id_rsa /var/backups/guessing_game

# Commit and move back to master
cd /var/backups/guessing_game
git add -A
git -c user.name='ubuntu' -c user.email='ubuntu@devilsec.org' commit -m 'Just developer things.'
git checkout master
