package bot

type Bot interface {
	GetError() error
	GetUpdates(limit int, timeout int, allowedUpdates []string)
}
