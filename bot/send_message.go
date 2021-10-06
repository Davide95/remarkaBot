package bot

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func (bot *tgBot) SendMessage(chatId int64, replyToMessageId int64, text string) error {
	if bot.err != nil {
		return bot.err
	}

	params := url.Values{
		"chat_id":                     {strconv.FormatInt(chatId, 10)},
		"text":                        {text},
		"reply_to_message_id":         {strconv.FormatInt(replyToMessageId, 10)},
		"allow_sending_without_reply": {"True"},
	}

	resp, err := http.PostForm(
		bot.makeQuery("sendMessage"),
		params,
	)

	if err != nil {
		bot.err = fmt.Errorf("Telegram API /sendMessage request failed: %w", err)
		return nil
	}
	defer resp.Body.Close()

	if status := resp.StatusCode; status != 200 && status != 401 {
		bot.err = fmt.Errorf(
			"Telegram API /sendMessage returned wrong status code: %d",
			resp.StatusCode,
		)
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return isResponseOk(body)
}
