package blogposts

import (
	"io/fs"
)

type Post struct {
}

func NewPostsFromFS(filesystem fs.FS) ([]Post, error) {
	dir, _ := fs.ReadDir(filesystem, ".")

	var posts []Post
	for range dir {
		posts = append(posts, Post{})
	}

	return posts, nil
}
