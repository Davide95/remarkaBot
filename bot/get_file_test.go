package bot

import (
	"testing"

	_ "embed"
)

//go:embed testdata/get_file.json
var get_file_fake_data []byte

func TestGetFileParse(t *testing.T) {
	bot := tgBot{}
	bot.getFileParse(get_file_fake_data)

	if err := bot.GetError(); err != nil {
		t.Fatal("get_file.json not parsed correctly")
	}
}
