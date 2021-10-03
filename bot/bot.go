package bot

import (
	"gitlab.com/mollofrollo/remarkabot/bot/fake"
	"gitlab.com/mollofrollo/remarkabot/bot/telegram"
)

type Bot interface {
	GetError() error
	GetUpdates(limit int, timeout int, allowedUpdates []string)
}

func GetBot(token string) Bot {
	return telegram.GetBot(token)
}

func GetFakeBot() Bot {
	return fake.Bot{}
}
