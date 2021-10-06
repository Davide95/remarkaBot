package remarkable

import (
	_ "embed"
	"os"
	"text/template"
	"time"
)

type metadataVars struct {
	LastModified int64
	VisibleName  string
}

//go:embed embed/metadata.json.tpl
var metadataText string

var metadataTemplate *template.Template = template.New("metadata")

func init() {
	var err error
	metadataTemplate, err = metadataTemplate.Parse(metadataText)
	if err != nil {
		panic(err)
	}
}

func insertMetadata(visibleName string, to string) error {
	out, err := os.Create(to)
	if err != nil {
		return err
	}
	defer out.Close()

	now := time.Now()
	vars := metadataVars{
		LastModified: now.UnixNano() / 1000000,
		VisibleName:  visibleName,
	}

	// TODO: testare escaping caratteri speciali
	err = metadataTemplate.Execute(out, vars)
	if err != nil {
		panic(err)
	}

	return nil
}
