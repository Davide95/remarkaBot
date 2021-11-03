package remarkable

import (
	_ "embed"
	"os"
	"text/template"
)

type contentVars struct {
	FileType string
}

//go:embed embed/content.json.tpl
var contentText string

var contentTemplate *template.Template = template.New("content")

func init() {
	var err error
	contentTemplate, err = contentTemplate.Parse(contentText)
	if err != nil {
		panic(err)
	}
}

func insertContent(fileType string, to string) error {
	out, err := os.Create(to)
	if err != nil {
		return err
	}
	defer out.Close()

	vars := contentVars{
		FileType: fileType,
	}

	err = contentTemplate.Execute(out, vars)
	if err != nil {
		panic(err)
	}

	return nil
}
