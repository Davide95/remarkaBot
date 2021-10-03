package telegram

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func (bot *bot) GetUpdates(limit int, timeout int, allowedUpdates []string) {
	if bot.err != nil {
		return
	}

	resp := bot.getUpdatesRequest(limit, timeout, allowedUpdates)
	_ = resp
}

func (bot *bot) getUpdatesRequest(limit int, timeout int, allowedUpdates []string) []byte {
	if bot.err != nil {
		return nil
	}

	jsonedAllowedUpdates, err := json.Marshal(allowedUpdates)
	if err != nil {
		bot.err = err
		return nil
	}

	params := url.Values{
		"offset":          {},
		"limit":           {strconv.Itoa(limit)},
		"timeout":         {strconv.Itoa(timeout)},
		"allowed_updates": {string(jsonedAllowedUpdates)},
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

	if resp.StatusCode != 200 {
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
