package bot

type Bot interface {
	GetError() error
	GetUpdates(limit int) []update
	Commit(offset int64)
}
