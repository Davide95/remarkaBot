package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type getUpdatesResponse struct {
	Result []update
}

type update struct {
	UpdateId      int64   `json:"update_id"`
	Message       message `json:",omitempty"`
	EditedMessage message `json:"edited_message,omitempty"`
}

type message struct {
	MessageId int64    `json:"message_id"`
	Document  document `json:",omitempty"`
	Chat      chat
}

type chat struct {
	Id   int64
	Type string
}

type document struct {
	FileId   string `json:"file_id"`
	FileName string `json:"file_name,omitempty"`
	MimeType string `json:"mime_type,omitempty"`
}

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
		"allowed_updates": {`["message", "channel_post"]`}, // Only new messages are currently supported
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
			"Telegram API /getUpdates returned wrong status code: (%d)",
			resp.StatusCode,
		)
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		bot.err = err
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
