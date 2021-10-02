package fake

type bot struct{}

func (bot *bot) GetError() *error {
	return nil
}

func GetBot(token string) bot {
	return bot{}
}
