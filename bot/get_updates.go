package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func (bot *tgBot) GetUpdates(limit int) []update {
	if bot.err != nil {
		return nil
	}

	resp := bot.getUpdatesRequest(limit)
	bot.isResponseOk(resp)
	return bot.getUpdatesParse(resp)
}

func (bot *tgBot) getUpdatesRequest(limit int) []byte {
	if bot.err != nil {
		return nil
	}

	params := url.Values{
		"limit":           {strconv.Itoa(limit)},
		"timeout":         {"0"},
		"allowed_updates": {"message"}, // Only new messages are currently supported
	}
	if bot.offset != 0 {
		params.Add("offset", strconv.FormatInt(bot.offset, 10))
	}

	resp, err := http.PostForm(
		bot.makeQuery("getUpdates"),
		params,
	)

	if err != nil {
		bot.err = err
		return nil
	}
	defer resp.Body.Close()

	if status := resp.StatusCode; status != 200 && status != 401 {
		bot.err = fmt.Errorf(
			"Telegram API /GetUpdates returned wrong status code: (%d)",
			resp.StatusCode,
		)
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		bot.err = err
		return nil
	}

	return body
}

func (bot *tgBot) getUpdatesParse(body []byte) []update {
	resp := getUpdatesResponse{}
	if bot.err != nil {
		return nil
	}

	err := json.Unmarshal(body, &resp)
	if err != nil {
		bot.err = err
		return nil
	}

	return resp.Result
}
