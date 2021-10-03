package bot

import (
	"os"
	"testing"

	_ "embed"
)

//go:embed testdata/get_updates.json
var get_updates_fake_data []byte

func TestGetUpdatesParse(t *testing.T) {
	bot := tgBot{}
	updates := bot.getUpdatesParse(get_updates_fake_data)

	if err := bot.GetError(); err != nil {
		t.Fatal("get_updates.json not parsed correctly")
	}

	if len(updates) != 1 {
		t.Fatalf("%d updates expected, found %d", 1, len(updates))
	}
}

func TestGetUpdates(t *testing.T) {
	const maxUpdates = 1

	token, present := os.LookupEnv("TELEGRAM_TOKEN")
	if !present {
		panic("env var TELEGRAM_TOKEN missing")
	}

	bot := tgBot{
		token: token,
	}
	updates := bot.GetUpdates(1)

	if err := bot.GetError(); err != nil {
		t.Fatal("get_file.json not parsed correctly")
	}

	if len(updates) > maxUpdates {
		t.Fatalf(
			"at most %d updates expected, found %d",
			maxUpdates,
			len(updates),
		)
	}
}
