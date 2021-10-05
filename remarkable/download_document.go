package remarkable

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)


func InsertDocument(from string, mime string, destination string) error {
	var extension string
	switch mime {
	case "application/pdf":
		extension = "pdf"
	case "application/epub+zip":
		extension = "epub"
	default:
		panic(fmt.Sprintf("MimeType of %s not supported: %s", from, mime))
	}

	basePath := filepath.Join(
		from,
		uuid.New().String(),
	)
	docPath := fmt.Sprintf("%s.%s", basePath, extension)

	return downloadDocument(from, docPath)
}

func downloadDocument(from string, to string) error {
	resp, err := http.Get(from)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(to)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		os.Remove(to)
	}

	return err
}