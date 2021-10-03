package bot

func (bot *tgBot) Commit(offset int64) {
	if bot.err == nil {
		bot.offset = offset + 1
	}
}
