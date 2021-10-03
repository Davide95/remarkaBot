package telegram

type bot struct {
	err    error
	token  string
	offset int // TODO: cambiare il tipo
}

func (bot *bot) GetError() error {
	return bot.err
}

func GetBot(token string) *bot {
	return &bot{
		token: token,
	}
}
