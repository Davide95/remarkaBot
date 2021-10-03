package fake

import _ "embed"

type Bot struct{}

func (bot Bot) GetError() error {
	return nil
}

//go:embed dummy_responses/get_updates.json
var getUpdatesBody []byte

func (bot Bot) GetUpdates(limit int, timeout int, allowedUpdates []string) {
	_ = getUpdatesBody
}
