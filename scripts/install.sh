#!/bin/bash
set -eux

GOOS="linux" GOARCH="arm" GOARM="7" go build .
ssh root@"$1" 'mkdir -p /opt/remarkabot && mkdir -p $HOME/.config/remarkabot'
scp ./remarkabot root@"$1":/opt/remarkabot/remarkabot
scp ./scripts/run.sh root@"$1":/opt/remarkabot/run.sh
scp ./systemd/remarkabot.service root@"$1":/etc/systemd/system/remarkabot.service
scp ./systemd/remarkabot.timer root@"$1":/etc/systemd/system/remarkabot.timer
scp ./systemd/.env root@"$1":/home/root/.config/remarkabot/.env
ssh root@"$1" 'systemctl daemon-reload && systemctl enable --now remarkabot.timer'