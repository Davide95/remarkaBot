package telegram

type bot struct {
	err   *error
	token string
}

func (bot *bot) GetError() *error {
	return bot.err
}

func GetBot(token string) bot {
	return bot{
		token: token,
	}
}
