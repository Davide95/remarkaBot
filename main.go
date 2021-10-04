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
			chat := message.Chat

			if chat.Type == "private" {
				logger.Warn("Private messages received and discarded", zap.Int64("id", message.MessageId))
				bot.SendMessage(chat.Id, message.MessageId, "Hey!\nAdd me to a group and send me documents to see them in your Remarkable!")
				bot.Commit(update.UpdateId)
				continue
			}

			document := message.Document

			if document.FileId == "" {
				logger.Error("Message does not contain a document", zap.Int64("id", message.MessageId))
				bot.SendMessage(chat.Id, message.MessageId, "You can send me only documents")
				bot.Commit(update.UpdateId)
				continue
			}

			mime := document.MimeType
			if mime != "application/pdf" && mime != "application/epub+zip" {
				logger.Error("Document is not a PDF or epub", zap.String("id", document.FileId))
				bot.SendMessage(message.Chat.Id, message.MessageId, "Document is not a PDF or epub")
				bot.Commit(update.UpdateId)
				continue
			}

			url := bot.GetFile(document.FileId)
			logger.Debug("File received", zap.String("URL", url.String()))
			bot.Commit(update.UpdateId)
		}

	}

	if err := bot.GetError(); err != nil {
		logger.Fatal(err.Error())
	}

	logger.Info("RemarkaBot ended")
}
