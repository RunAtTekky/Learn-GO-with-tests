package blogposts

import (
	"io"
	"io/fs"
)

type Post struct {
	Title string
}

func NewPostsFromFS(filesystem fs.FS) ([]Post, error) {
	dir, _ := fs.ReadDir(filesystem, ".")

	var posts []Post
	for _, file := range dir {
		post, err := getPost(filesystem, file)

		if err != nil {
			return nil, err // TODO: Needs clarification
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(filesystem fs.FS, file fs.DirEntry) (Post, error) {

	postFile, err := filesystem.Open(file.Name())
	if err != nil {
		return Post{}, nil
	}

	defer postFile.Close()

	postData, err := io.ReadAll(postFile)
	if err != nil {
		return Post{}, nil
	}

	post := Post{Title: string(postData)[7:]}
	return post, nil
}
