package bot

type Bot interface {
	GetError() error
	GetUpdates(limit int) []update
	GetFile(fileId string) string
	SendMessage(chatId int64, replyToMessageId int64, text string) error
	Commit(offset int64)
}
