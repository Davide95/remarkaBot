package bot

import (
	"errors"
	"testing"
)

func TestGetError(t *testing.T) {
	err := errors.New("dummy-error-placeholder")
	bot := tgBot{
		err: err,
	}

	if !errors.Is(bot.GetError(), err) {
		t.Fatalf("'%v' should be equal to '%v'", bot.err, err)
	}
}

func TestGetBot(t *testing.T) {
	token := "dummy-token-placeholder"
	bot := tgBot{token: token}

	if bot.token != token {
		t.Fatalf("'%s' should be equal to '%s'", token, bot.token)
	}
}
