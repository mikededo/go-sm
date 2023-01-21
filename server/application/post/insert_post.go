package post

import (
	"github.com/mddg/go-sm/server/application"
	"github.com/mddg/go-sm/server/domain/post"
	"github.com/mddg/go-sm/server/domain/util"
)

type InsertPostRequest struct {
	Content    string `validate:"required"`
	AuthorID   int    `validate:"required"`
	IsReshared bool
	ResharedID int
	Likes      int
	Dislikes   int
}

type InsertPostService struct {
	repository post.Repository
}

func NewInsertPostService(repository post.Repository) *InsertPostService {
	return &InsertPostService{repository}
}

func NewInsertPostRequest(
	content string,
	authorID, resharedID int,
) InsertPostRequest {
	return InsertPostRequest{
		Content:    content,
		AuthorID:   authorID,
		IsReshared: resharedID != 0,
		ResharedID: resharedID,
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
		req.AuthorID,
		"",
		req.Content,
		req.IsReshared,
		req.ResharedID,
		req.Likes,
		req.Dislikes,
		nil,
	)
	return s.repository.InsertPost(post)
}
