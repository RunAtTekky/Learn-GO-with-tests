package blogposts

import (
	"io/fs"
)

func NewPostsFromFS(filesystem fs.FS) ([]Post, error) {
	dir, _ := fs.ReadDir(filesystem, ".")

	var posts []Post
	for _, file := range dir {
		post, err := getPost(filesystem, file.Name())

		if err != nil {
			return nil, err // TODO: Needs clarification
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(filesystem fs.FS, fileName string) (Post, error) {
	postFile, err := filesystem.Open(fileName)
	if err != nil {
		return Post{}, nil
	}

	defer postFile.Close()

	post, err := newPost(postFile)
	return post, err

}
