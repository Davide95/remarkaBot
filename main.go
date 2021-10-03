package main

import (
	"fmt"
	"log"
	"os"

	"gitlab.com/mollofrollo/remarkabot/bot"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("Unable to initialize zap logger: %v", err)
	}
	defer logger.Sync()
	logger.Info("RemarkaBot started")

	telegramToken, present := os.LookupEnv("TELEGRAM_TOKEN")
	if !present {
		logger.Fatal("env var TELEGRAM_TOKEN missing")
	}

	bot := bot.GetBot(telegramToken)
	fmt.Println("Errors:", bot.GetError())
}
