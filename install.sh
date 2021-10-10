ssh root@10.11.99.1 'mkdir -p /opt/remarkabot && mkdir -p $HOME/.config/remarkabot' && \
scp ./remarkabot root@10.11.99.1:/opt/remarkabot/remarkabot && \
scp ./systemd/remarkabot.service root@10.11.99.1:/etc/systemd/user/remarkabot.service && \
scp ./systemd/remarkabot.timer root@10.11.99.1:/etc/systemd/user/remarkabot.timer && \
scp ./systemd/.env root@10.11.99.1:$HOME/.config/remarkabot/.env