package main

import (
	"fmt"

	"gitlab.com/mollofrollo/remarkabot/bot/telegram"
	_ "gitlab.com/mollofrollo/remarkabot/bot/telegram"
)

func main() {
	bot := telegram.GetBot("dummy")
	fmt.Println("Errors:", bot.GetError())
}
