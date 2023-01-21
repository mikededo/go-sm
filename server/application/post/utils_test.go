package post_test

import (
	"github.com/mddg/go-sm/server/application"
	"github.com/mddg/go-sm/server/domain/post"
	"github.com/mddg/go-sm/server/domain/shared"
)

type PostRepositorySpy struct {
	application.RepositorySpy[[]*post.Post]
}

func (spy *PostRepositorySpy) InsertPost(p post.Post) (*post.Post, error) {
	spy.Calls = append(spy.Calls, p)
	result := spy.Result()
	if len(result) > 0 {
		return result[0], spy.Error()
	}

	return nil, spy.Error()
}

type FindUserPostsArguments struct {
	ID          uint
	PageRequest shared.PagedRequest
}

func (spy *PostRepositorySpy) FindUserPosts(id uint, page shared.PagedRequest) ([]*post.Post, error) {
	spy.Calls = append(spy.Calls, FindUserPostsArguments{ID: id, PageRequest: page})
	return spy.Result(), spy.Error()
}
