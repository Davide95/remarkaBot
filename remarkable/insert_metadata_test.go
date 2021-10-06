package remarkable

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestInsertMetadata(t *testing.T) {
	dir, err := ioutil.TempDir(".", "test-insert-metadata-")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	file := filepath.Join(dir, "metadata")
	visibleName := "visible-name-placeholder"
	err = insertMetadata(visibleName, file)

	if err != nil {
		t.Fatalf("Error with the template: %v", err)
	}
}
