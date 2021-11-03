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

	bot.err = isResponseOk(body)
}

func isResponseOk(body []byte) error {
	resp := baseResponse{}
	err := json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	if !resp.Ok {
		return errors.New(resp.Description)
	}

	return nil
}
