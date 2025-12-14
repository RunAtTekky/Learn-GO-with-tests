package blogrender

import (
	"html/template"
	"io"
)

const (
	postTemplate = `<h1>{{.Title}}</h1>
<p>{{.Description}}</p>
Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>`
)

func Render(file io.Writer, post Post) error {

	templ, err := template.New("blog").Parse(postTemplate)

	if err != nil {
		return err
	}

	if err := templ.Execute(file, post); err != nil {
		return err
	}

	return nil
}
