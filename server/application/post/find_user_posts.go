package post

import (
	"github.com/mddg/go-sm/server/domain/post"
	"github.com/mddg/go-sm/server/domain/shared"
)

type FindUserPostsService struct {
	repository post.PostRepository
}

func NewFindUserPostsService(repository post.PostRepository) *FindUserPostsService {
	return &FindUserPostsService{repository}
}

func (s *FindUserPostsService) Run(id uint, page shared.PagedRequest) ([]*post.Post, error) {
	return s.repository.FindUserPosts(id, page)
}
