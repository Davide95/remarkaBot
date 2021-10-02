package telegram

import "fmt"

func (bot *bot) makeQuery(methodName string) string {
	return fmt.Sprintf("https://api.telegram.org/bot%s/%s", bot.token, methodName)
}

func (bot *bot) makeQueryFile(filePath string) string {
	return fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", bot.token, filePath)
}
