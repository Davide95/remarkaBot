package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type getFileResponse struct {
	Result file
}

type file struct {
	FilePath string `json:"file_path,omitempty"`
}

func (bot *tgBot) GetFile(fileId string) url.URL {
	if bot.err != nil {
		return url.URL{}
	}

	resp := bot.getFileRequest(fileId)
	bot.isResponseOk(resp)
	return bot.getFileParse(resp)
}

func (bot *tgBot) getFileRequest(fileId string) []byte {
	if bot.err != nil {
		return nil
	}

	params := url.Values{
		"file_id": {fileId},
	}

	resp, err := http.PostForm(
		bot.makeQuery("getFile"),
		params,
	)

	if err != nil {
		bot.err = fmt.Errorf("Telegram API /getFile request failed: %w", err)
		return nil
	}
	defer resp.Body.Close()

	if status := resp.StatusCode; status != 200 && status != 401 {
		bot.err = fmt.Errorf(
			"Telegram API /getFile returned wrong status code: %d",
			resp.StatusCode,
		)
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		bot.err = fmt.Errorf("Telegram API /getFile body error: %w", err)
	}
	return body
}

func (bot *tgBot) getFileParse(body []byte) url.URL {
	resp := getFileResponse{}
	if bot.err != nil {
		return url.URL{}
	}

	err := json.Unmarshal(body, &resp)
	if err != nil {
		bot.err = fmt.Errorf("Telegram API /getFile json parsing error: %w", err)
		return url.URL{}
	}

	url, err := url.ParseRequestURI(
		bot.makeQueryFile(resp.Result.FilePath),
	)
	if err != nil {
		bot.err = fmt.Errorf("Telegram API /getFile file URL parsing error: %w", err)
	}

	return *url
}
