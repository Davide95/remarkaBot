package bot

import "net/url"

type Bot interface {
	GetError() error
	GetUpdates(limit int) []update
	GetFile(fileId string) url.URL
	SendMessage(chatId int64, replyToMessageId int64, text string) error
	Commit(offset int64)
}
