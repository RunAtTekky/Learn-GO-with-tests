package blogrender

import (
	"embed"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

type PostRenderer struct {
	templ    *template.Template
	mdParser *parser.Parser
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	return &PostRenderer{templ: templ, mdParser: parser}, nil
}

func (pr *PostRenderer) Render(file io.Writer, post Post) error {
	return pr.templ.ExecuteTemplate(file, "blog.gohtml", newPostVM(post, pr))
}

func (pr *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return pr.templ.ExecuteTemplate(w, "index.gohtml", posts)
}

type postViewModel struct {
	Post
	HTMLbody template.HTML
}

func newPostVM(p Post, pr *PostRenderer) postViewModel {
	vm := postViewModel{Post: p}
	vm.HTMLbody = template.HTML(markdown.ToHTML([]byte(p.Body), pr.mdParser, nil))
	return vm
}
