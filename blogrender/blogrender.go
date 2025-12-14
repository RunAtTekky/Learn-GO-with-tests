package blogrender

import (
	"embed"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type PostRenderer struct {
	templ *template.Template
}

func (pr *PostRenderer) Render(file io.Writer, post Post) error {
	if err := pr.templ.ExecuteTemplate(file, "blog.gohtml", post); err != nil {
		return err
	}

	return nil
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

func (pr *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return nil
}
