package telegram

import "gitlab.com/mollofrollo/remarkabot/bot"

type tgBot struct {
	err    error
	token  string
	offset int // TODO: cambiare il tipo
}

func (bot *tgBot) GetError() error {
	return bot.err
}

func GetBot(token string) bot.Bot {
	return &tgBot{
		token: token,
	}
}
