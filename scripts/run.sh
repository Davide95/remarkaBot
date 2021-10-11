#!/bin/bash

if LC_ALL=C fgrep -q "wifion=true" /etc/remarkable.conf; then
    /opt/remarkabot/remarkabot
else
    echo "WiFi is not enabled"
fi