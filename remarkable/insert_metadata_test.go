package remarkable

import (
	"bytes"
	_ "embed"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"
)

//go:embed testdata/insert_metadata.json
var insert_metdata_fake_data []byte

func TestInsertMetadata(t *testing.T) {
	dir, err := ioutil.TempDir(".", "test-insert-metadata-")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)

	file := filepath.Join(dir, "metadata")
	visibleName := `placeholder-with-special-char-"`
	err = insertMetadata(visibleName, time.Time{}, file)

	if err != nil {
		t.Fatalf("Error with the template: %v", err)
	}

	result, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	if bytes.Compare(result, insert_metdata_fake_data) != 0 {
		t.Log(string(result))
		t.Fatalf("Metadata rendered incorrectly")
	}

	err = os.RemoveAll(dir)
	if err != nil {
		panic(err)
	}
}
