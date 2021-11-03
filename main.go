package main

import (
	"log"
	"os"
	"runtime"

	"gitlab.com/mollofrollo/remarkabot/bot"
	"gitlab.com/mollofrollo/remarkabot/remarkable"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("Unable to initialize zap logger: %s", err.Error())
	}
	defer logger.Sync()

	logger.Info("RemarkaBot started")

	telegramToken, present := os.LookupEnv("TELEGRAM_TOKEN")
	if !present || telegramToken == "" {
		logger.Fatal("env var TELEGRAM_TOKEN missing")
	}

	destinationFolder, present := os.LookupEnv("DESTINATION_FOLDER")
	if !present {
		destinationFolder = "/home/root/.local/share/remarkable/xochitl/"
	}

	bot := bot.GetBot(telegramToken)

	const maxConcUpdates = 1
	for updates := bot.GetUpdates(maxConcUpdates); len(updates) > 0; updates = bot.GetUpdates(maxConcUpdates) {
		runtime.GC()
		logger.Info("Fetching new updates")

		for _, update := range updates {
			logger.Info("New update received", zap.Int64("id", update.UpdateId))

			message := update.Message
			chat := message.Chat

			if chat.Type == "private" {
				bot.Commit(update.UpdateId)

				logger.Warn("Private messages received and discarded", zap.Int64("id", message.MessageId))
				bot.SendMessage(chat.Id, message.MessageId, "Hey!\nAdd me to a group and send me documents to see them in your reMarkable!")
				continue
			}

			document := message.Document

			if document.FileId == "" {
				bot.Commit(update.UpdateId)

				logger.Warn("Message does not contain a document", zap.Int64("id", message.MessageId))
				bot.SendMessage(chat.Id, message.MessageId, "You can send me only documents")
				continue
			}

			mime := document.MimeType
			if mime != "application/pdf" && mime != "application/epub+zip" {
				bot.Commit(update.UpdateId)

				logger.Warn("Document is not a PDF or epub", zap.String("id", document.FileId))
				bot.SendMessage(message.Chat.Id, message.MessageId, "Document is not a PDF or epub")
				continue
			}

			url := bot.GetFile(document.FileId)
			if err := bot.GetError(); err != nil {
				logger.Fatal(err.Error())
			}

			err = remarkable.InsertDocument(url, mime, document.FileName, destinationFolder)
			if err != nil {
				logger.Fatal(err.Error())
			}

			bot.Commit(update.UpdateId)

			bot.SendMessage(
				message.Chat.Id,
				message.MessageId,
				"Document received, reboot your device to see it",
			)
		}

	}

	if err := bot.GetError(); err != nil {
		logger.Fatal(err.Error())
	}

	logger.Info("RemarkaBot ended")
}
