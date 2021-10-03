package bot

func (bot *tgBot) Commit(offset int64) {
	bot.offset = offset + 1
}
