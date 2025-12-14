package blogrender_test

import (
	"bytes"
	blogrender "go_with_test/blogrender"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	var (
		aPost = blogrender.Post{
			Title:       "Hello World",
			Description: "Say Hello World",
			Tags:        []string{"go", "tdd"},
			Body:        "This is the body",
		}
	)

	postRenderer, err := blogrender.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("coverts single Post to HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := postRenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		aPost = blogrender.Post{
			Title:       "Hello World",
			Description: "Say Hello World",
			Tags:        []string{"go", "tdd"},
			Body:        "This is the body",
		}
	)

	postRenderer, err := blogrender.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	for b.Loop() {
		postRenderer.Render(io.Discard, aPost)
	}
}
