package post

import "github.com/mddg/go-sm/server/domain/shared"

type PostRepository interface {
	// InsertPost saves a new Post
	InsertPost(Post) (*Post, error)

	// FindUserPosts returns the list of posts given the user id and the page
	FindUserPosts(uint, shared.PagedRequest) ([]*Post, error)
}
