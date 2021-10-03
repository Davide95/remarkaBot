package fake

type Bot struct{}

func (bot Bot) GetError() error {
	return nil
}

func (bot Bot) GetUpdates(limit int, timeout int, allowedUpdates []string) {
	return
}