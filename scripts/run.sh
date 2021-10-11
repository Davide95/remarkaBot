#!/bin/bash

if fgrep -q "wifion=true" /etc/remarkable.conf; then
    /opt/remarkabot/remarkabot
else
    echo "WiFi is not enabled"
fi