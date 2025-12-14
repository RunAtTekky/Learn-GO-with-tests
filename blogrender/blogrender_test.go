package blogrender_test

import (
	"bytes"
	blogrender "go_with_test/blogrender"
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

	t.Run("coverts single Post to HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := blogrender.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}
