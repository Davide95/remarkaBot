#!/bin/bash
set -eux

go build .
ssh root@$1 'mkdir -p /opt/remarkabot && mkdir -p $HOME/.config/remarkabot'
scp ./remarkabot root@$1:/opt/remarkabot/remarkabot
scp ./systemd/remarkabot.service root@$1:/etc/systemd/system/remarkabot.service
scp ./systemd/remarkabot.timer root@$1:/etc/systemd/system/remarkabot.timer
scp ./systemd/.env root@$1:/home/root/.config/remarkabot/.env
ssh root@$1 'systemctl enable remarkabot.timer'