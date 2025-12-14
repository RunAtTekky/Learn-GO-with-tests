package blogrender

import (
	"fmt"
	"io"
)

func Render(file io.Writer, post Post) error {
	_, err := fmt.Fprintf(file, "<h1>%s</h1>", post.Title)
	return err
}
