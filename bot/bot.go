package bot

import "net/url"

type Bot interface {
	GetError() error
	GetUpdates(limit int) []update
	GetFile(fileId string) url.URL
	Commit(offset int64)
}
