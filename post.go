package blogposts

import (
	"bufio"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titlePrefix       = `Title: `
	descriptionPrefix = `Description: `
	tagsPrefix        = `Tags: `
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	return Post{
		Title:       readMetaLine(titlePrefix),
		Description: readMetaLine(descriptionPrefix),
		Tags:        strings.Split(readMetaLine(tagsPrefix), ", "),
	}, nil
}
