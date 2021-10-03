package main

import (
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
	bot.GetUpdates(1, 0)
	if err := bot.GetError(); err != nil {
		logger.Fatal(
			"Error while fetching updates",
			zap.String("error", err.Error()),
		)
	}
}
