set GOOS=linux
set GOARCH=arm
set GOARM=7

go build . || exit /b
ssh root@"%1" "mkdir -p /opt/remarkabot && mkdir -p $HOME/.config/remarkabot" || exit /b
scp ./remarkabot root@"%1":/opt/remarkabot/remarkabot || exit /b
ssh root@"%1" "chmod +x /opt/remarkabot/remarkabot" || exit /b
scp ./scripts/run.sh root@"%1":/opt/remarkabot/run.sh || exit /b
ssh root@"%1" "chmod +x /opt/remarkabot/run.sh" || exit /b
scp ./systemd/remarkabot.service root@"%1":/etc/systemd/system/remarkabot.service || exit /b
scp ./systemd/remarkabot.timer root@"%1":/etc/systemd/system/remarkabot.timer || exit /b
scp ./systemd/.env root@"%1":/home/root/.config/remarkabot/.env || exit /b
ssh root@"%1" "systemctl daemon-reload && systemctl enable --now remarkabot.timer"