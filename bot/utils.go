package bot

import (
	"encoding/json"
	"errors"
	"fmt"
)

func (bot *tgBot) makeQuery(methodName string) string {
	return fmt.Sprintf("https://api.telegram.org/bot%s/%s", bot.token, methodName)
}

func (bot *tgBot) makeQueryFile(filePath string) string {
	return fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", bot.token, filePath)
}

type baseResponse struct {
	Ok          bool
	Description string `json:",omitempty"`
}

func (bot *tgBot) isResponseOk(body []byte) {
	if bot.err != nil {
		return
	}

	resp := baseResponse{}
	err := json.Unmarshal(body, &resp)
	if err != nil {
		bot.err = err
		return
	}

	if !resp.Ok {
		bot.err = errors.New(resp.Description)
	}
}
