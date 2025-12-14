package blogrender_test

import (
	"bytes"
	blogrender "go_with_test/blogrender"
	"testing"
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

	t.Run("coverts single Post to HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		err := blogrender.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>Hello World</h1>`

		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	})
}
