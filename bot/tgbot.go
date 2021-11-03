package bot

type tgBot struct {
	err    error
	token  string
	offset int64
}

func (bot *tgBot) GetError() error {
	return bot.err
}

func GetBot(token string) Bot {
	return &tgBot{
		token: token,
	}
}
