package telegram

import "testing"

func TestMakeQuery(t *testing.T) {
	token := "dummy-token-placeholder"
	methodName := "dummy-method-name-placeholder"
	bot := GetBot(token)

	query := bot.makeQuery(methodName)
	groundTruth := "https://api.telegram.org/bot" + token + "/" + methodName
	if query != groundTruth {
		t.Fatalf("'%s' should be equal to '%s'", query, groundTruth)
	}
}

func TestMakeQueryFile(t *testing.T) {
	token := "dummy-token-placeholder"
	filePath := "dummy-file-path-placeholder"
	bot := GetBot(token)

	query := bot.makeQueryFile(filePath)
	groundTruth := "https://api.telegram.org/file/bot" + token + "/" + filePath
	if query != groundTruth {
		t.Fatalf("'%s' should be equal to '%s'", query, groundTruth)
	}
}
