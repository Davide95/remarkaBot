package bot

import "testing"
import _ "embed"

func TestMakeQuery(t *testing.T) {
	token := "dummy-token-placeholder"
	methodName := "dummy-method-name-placeholder"
	bot := tgBot{token: token}

	query := bot.makeQuery(methodName)
	groundTruth := "https://api.telegram.org/bot" + token + "/" + methodName
	if query != groundTruth {
		t.Fatalf("'%s' should be equal to '%s'", query, groundTruth)
	}
}

func TestMakeQueryFile(t *testing.T) {
	token := "dummy-token-placeholder"
	filePath := "dummy-file-path-placeholder"
	bot := tgBot{token: token}

	query := bot.makeQueryFile(filePath)
	groundTruth := "https://api.telegram.org/file/bot" + token + "/" + filePath
	if query != groundTruth {
		t.Fatalf("'%s' should be equal to '%s'", query, groundTruth)
	}
}

//go:embed testdata/base_response.json
var base_response_fake_data []byte

func TestIsResponseOk(t *testing.T) {
	err := isResponseOk(base_response_fake_data)

	if err != nil {
		t.Fatalf("No errors expected, returned %s", err.Error())
	}
}
