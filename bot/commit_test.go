package bot

import "testing"

func TestCommit(t *testing.T) {

	bot := tgBot{}

	if bot.Commit(10); bot.offset != 11 {
		t.Logf("Expected offset: 11, found: %d", bot.offset)
		t.Fatal("Commit should update the offset")
	}

	if bot.Commit(11); bot.offset != 12 {
		t.Logf("Expected offset: 11, found: %d", bot.offset)
		t.Fatal("Commit should increase the offset by 1 if the current offset is the old offset")
	}

	if bot.Commit(3); bot.offset != 12 {
		t.Logf("Expected offset: 11, found: %d", bot.offset)
		t.Fatal("Commit should consider a new offset only if it is greater than the threshold")
	}
}
