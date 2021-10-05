package remarkable

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func DownloadDocument(from string, destination string) error {
	resp, err := http.Get(from)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc := filepath.Join(from, uuid.New().String())
	out, err := os.Create(fmt.Sprintf("%s.pdf", doc))
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		os.Remove("output.txt")
	}

	return err
}
