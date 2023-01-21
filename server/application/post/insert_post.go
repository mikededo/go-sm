package post

import (
	"github.com/mddg/go-sm/server/application"
	"github.com/mddg/go-sm/server/domain/post"
	"github.com/mddg/go-sm/server/domain/util"
)

type InsertPostRequest struct {
	Content    string `validate:"required"`
	AuthorId   int    `validate:"required"`
	IsReshared bool
	ResharedId int
	Likes      int
	Dislikes   int
}

type InsertPostService struct {
	repository post.PostRepository
}

func NewInsertPostService(repository post.PostRepository) *InsertPostService {
	return &InsertPostService{repository}
}

func NewInsertPostRequest(
	content string,
	authorId, resharedId int,
) InsertPostRequest {
	return InsertPostRequest{
		Content:    content,
		AuthorId:   authorId,
		IsReshared: resharedId != 0,
		ResharedId: resharedId,
		Likes:      0,
		Dislikes:   0,
	}
}

func (s *InsertPostService) Run(req InsertPostRequest) (*post.Post, error) {
	err := util.Validator().Struct(req)
	if err != nil {
		return nil, application.ErrInvalidRequest
	}

	post := post.NewEmptyPost(
		req.AuthorId,
		"",
		req.Content,
		req.IsReshared,
		req.ResharedId,
		req.Likes,
		req.Dislikes,
		nil,
	)
	return s.repository.InsertPost(post)
}
