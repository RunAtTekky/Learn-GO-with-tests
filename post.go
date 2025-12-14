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
	title := readMetaLine(titlePrefix)
	description := readMetaLine(descriptionPrefix)
	tags := strings.Split(readMetaLine(tagsPrefix), ", ")

	var body string
	// TODO: Read all the body

	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		Body:        body,
	}, nil
}
