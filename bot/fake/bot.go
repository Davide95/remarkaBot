package fake

import (
	_ "embed"

	"gitlab.com/mollofrollo/remarkabot/bot"
)

type FakeBot struct{}

func (bot FakeBot) GetError() error {
	return nil
}

func GetBot(token string) bot.Bot {
	return &FakeBot{}
}

//go:embed dummy_responses/get_updates.json
var getUpdatesBody []byte

func (bot FakeBot) GetUpdates(limit int, timeout int, allowedUpdates []string) {
	_ = getUpdatesBody
}
