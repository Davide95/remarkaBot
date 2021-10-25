# RemarkaBot

Fetch your documents from Telegram to your Remarkable 2!

This program will fetch new documents every 10 minutes for you :)

## Installation

### Telegram

In order to user `RemarkaBot` you need a Telegram Bot, you can create one using [BotFather](https://t.me/BotFather).  
Sending private messages to him is not allowed. For security reasons it's crucial that you are in control of the groups `RemarkaBot` accepts files from. To make sure nobody adds your `RemarkaBot` instance to some random (potentially malicious) groups we reccomend to block all group invitations after having added the bot to your groups. You can simply do that by sending `/setjoingroups` to `BotFather` and follow instructions. 
From now on to add `RemarkaBot` to other groups you must unlock group invitations and lock it again.   
Leaving the setting unlocked might be tempting, but **please** don't do that.

Consider to disable `/setprivacy` to send documents to `RemarkaBot` without having to mention him.

### Remarkabot

You first need to connect your Remarkable 2 to your PC to allow an SSH connection (see https://support.remarkable.com/hc/en-us/articles/360002662557-Help to understand how).
You also need to have the Go compiler installed on your PC.

Steps to install it:
* `git clone https://gitlab.com/mollofrollo/remarkabot.git`
* edit the `systemd/.env` file to put your Bot's token
* `sh scripts/install.sh $RM2IP`, where `$RM2IP` is the IP of the connected Remarkable 2

Each time the device is updated, you need to reinstall Remarkabot.
Grab the opportunity to update the repo first (`git pull`).

## Acknowledgment

Contributors:
* Davide Riva <driva95[at]protonmail[dot]com> (Mantainer)
* Laura Nesossi <laura_nesossi[at]protonmail[dot]com> (logo creator)
* Roberto Castellotti <me[at]rcastellotti[dot]dev

## License

This project is licensed under the GNU General Public License v3.0. 

