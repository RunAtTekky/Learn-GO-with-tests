package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
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

	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		Body:        readBodyText(scanner),
	}, nil
}

func readBodyText(scanner *bufio.Scanner) string {
	// Reads the ---
	scanner.Scan()

	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}

	body := strings.TrimSuffix(buf.String(), "\n")
	return body
}
