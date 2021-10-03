package bot

import (
	"testing"

	_ "embed"
)

//go:embed testdata/get_updates.json
var get_updates_fake_data []byte

func TestGetUpdatesParse(t *testing.T) {
	bot := tgBot{}
	updates := bot.getUpdatesParse(get_updates_fake_data)

	if err := bot.GetError(); err != nil {
		t.Fatalf("get_updates.json not parsed correctly")
	}

	if len(updates) != 1 {
		t.Fatalf("%d updates expected, found %d", 1, len(updates))
	}
}
