#!/bin/bash

if LC_ALL=C grep -F "wifion=true" /etc/remarkable.conf; then
    /opt/remarkabot/remarkabot
else
    echo "WiFi is not enabled"
fi