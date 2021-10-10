package remarkable

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

func InsertDocument(from string, mime string, visibleName string, destination string) error {
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
		destination,
		uuid.New().String(),
	)
	docPath := fmt.Sprintf("%s.%s", basePath, extension)
	metadataPath := fmt.Sprintf("%s.%s", basePath, "metadata")
	contentPath := fmt.Sprintf("%s.%s", basePath, "content")

	if err := downloadDocument(from, docPath); err != nil {
		os.Remove(docPath)
		return fmt.Errorf(
			"Downloading %s to %s failed: %w",
			from, basePath, err,
		)
	}

	if err := insertMetadata(visibleName, time.Now(), metadataPath); err != nil {
		os.Remove(docPath)
		os.Remove(metadataPath)
		return fmt.Errorf(
			"Metadata file %s not created: %w",
			metadataPath, err,
		)
	}

	if err := insertContent(extension, contentPath); err != nil {
		os.Remove(docPath)
		os.Remove(metadataPath)
		os.Remove(contentPath)
		return fmt.Errorf(
			"Content file %s not created: %w",
			contentPath, err,
		)
	}

	return nil
}
