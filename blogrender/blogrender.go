package blogrender

import (
	"bytes"
	"fmt"
	"io"
)

func Render(file io.Writer, post Post) error {
	_, err := fmt.Fprintf(file, "<h1>%s</h1>\n", post.Title)
	_, err = fmt.Fprintf(file, "<p>%s</p>\n", post.Description)

	buf := bytes.Buffer{}
	for _, tag := range post.Tags {
		fmt.Fprintf(&buf, "<li>%s</li>", tag)
	}

	tagsHTML := buf.String()
	_, err = fmt.Fprintf(file, "Tags: <ul>%s</ul>", tagsHTML)
	return err
}
