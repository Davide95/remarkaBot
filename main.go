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

	const maxUpdates = 100
	for updates := bot.GetUpdates(maxUpdates); bot.GetError() == nil && len(updates) > 0; updates = bot.GetUpdates(maxUpdates) {
		logger.Info("Fetching new updates")

		for _, update := range updates {
			logger.Info("New update received", zap.Int64("id", update.UpdateId))

			message := update.Message
			if message.Document.FileId != "" {
				url := bot.GetFile(update.Message.Document.FileId)
				logger.Debug("File received", zap.String("URL", url.String()))
			}

			bot.Commit(update.UpdateId)
		}

	}

	if err := bot.GetError(); err != nil {
		logger.Fatal(
			"Error while fetching updates",
			zap.String("error", err.Error()),
		)
	}

	logger.Info("RemarkaBot ended")
}
