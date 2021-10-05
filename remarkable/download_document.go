package remarkable

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func DownloadDocument(from string, mime string, destination string) error {
	resp, err := http.Get(from)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc := filepath.Join(
		from, 
		fmt.Sprintf("%s.%s", uuid.New().String(), extension)
	)

	out, err := os.Create(doc)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		os.Remove(doc)
	}

	return err
}
