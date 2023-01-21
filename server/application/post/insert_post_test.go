package post_test

import (
	"errors"
	"testing"
	"time"

	"github.com/mddg/go-sm/server/application"
	"github.com/mddg/go-sm/server/application/post"
	postEntity "github.com/mddg/go-sm/server/domain/post"
)

var resharedPost postEntity.Post = postEntity.NewPost(1, 1, "mikededo", "This is a test post", true, 1, nil, 0, 0, nil, time.Now(), time.Now())
var notResharedPost postEntity.Post = postEntity.NewPost(1, 1, "mikededo", "This is a test post", false, 0, nil, 0, 0, nil, time.Now(), time.Now())

func validateInsertPostCalls(t *testing.T, s *PostRepositorySpy, req post.InsertPostRequest) {
	arg := s.Calls[0].(postEntity.Post)
	application.CheckPopertyEquality(t, "Content", arg.Content, req.Content)
	application.CheckPopertyEquality(t, "Author.ID", arg.Author.ID, req.AuthorID)
	application.CheckPopertyEquality(t, "Author.Username", arg.Author.Username, "")
	application.CheckPopertyEquality(t, "IsReshared", arg.IsReshared, req.ResharedID != 0)
	application.CheckPopertyEquality(t, "ResharedId", arg.ResharedID, req.ResharedID)
	application.CheckPopertyEquality(t, "Likes", arg.Likes, req.Likes)
	application.CheckPopertyEquality(t, "Dislikes", arg.Dislikes, req.Dislikes)
	// TODO: add Collaborators
}

func validateInsertResult(t *testing.T, got, want *postEntity.Post) {
	application.CheckPopertyEquality(t, "Content", want.Content, got.Content)
	application.CheckPopertyEquality(t, "Author.ID", want.Author.ID, got.Author.ID)
	application.CheckPopertyEquality(t, "Author.Username", want.Author.Username, resharedPost.Author.Username)
	application.CheckPopertyEquality(t, "IsReshared", want.IsReshared, got.ResharedID != 0)
	application.CheckPopertyEquality(t, "ResharedId", want.ResharedID, got.ResharedID)
	application.CheckPopertyEquality(t, "Likes", want.Likes, got.Likes)
	application.CheckPopertyEquality(t, "Dislikes", want.Dislikes, got.Dislikes)
	application.CheckPopertyEquality(t, "CreatedAt", want.CreatedAt, got.CreatedAt)
	application.CheckPopertyEquality(t, "UpdatedAt", want.UpdatedAt, got.UpdatedAt)
}

func TestInsertPostService(t *testing.T) {
	t.Run("insert reshared post", func(t *testing.T) {
		spy := &PostRepositorySpy{
			RepositorySpy: application.NewRepositoryWithResultsAndErrors(
				[][]*postEntity.Post{{&resharedPost}},
				[]error{nil},
			),
		}

		services := post.NewInsertPostService(spy)
		req := post.NewInsertPostRequest("This is a test post", 1, 1)
		res, err := services.Run(req)

		spy.CalledOnce(t)
		if err != nil {
			t.Errorf("not expecting error, got %v\n", res)
		}
		validateInsertPostCalls(t, spy, req)
		validateInsertResult(t, &resharedPost, res)
	})

	t.Run("insert not reshared post", func(t *testing.T) {
		spy := &PostRepositorySpy{
			RepositorySpy: application.NewRepositoryWithResultsAndErrors(
				[][]*postEntity.Post{{&notResharedPost}},
				[]error{nil},
			),
		}

		services := post.NewInsertPostService(spy)
		req := post.NewInsertPostRequest("This is a test post", 1, 0)
		res, err := services.Run(req)

		spy.CalledOnce(t)
		if err != nil {
			t.Errorf("not expecting error, got %v\n", res)
		}
		validateInsertPostCalls(t, spy, req)
		validateInsertResult(t, &notResharedPost, res)
	})

	t.Run("repository error thrown", func(t *testing.T) {
		repositoryError := "repository error"
		spy := &PostRepositorySpy{
			RepositorySpy: application.NewRepositoryWithErrors[[]*postEntity.Post]([]error{errors.New(repositoryError)}),
		}

		service := post.NewInsertPostService(spy)
		res, err := service.Run(post.NewInsertPostRequest("This is a test post", 1, 0))

		spy.CalledOnce(t)
		if res != nil {
			t.Error("expected error, got nil\n")
		}
		if err.Error() != repositoryError {
			t.Errorf("got %s as error, wanted %s\n", err.Error(), repositoryError)
		}
	})

	t.Run("invalid post request", func(t *testing.T) {
		spy := &PostRepositorySpy{
			RepositorySpy: application.NewRepositoryWithErrors[[]*postEntity.Post]([]error{application.ErrInvalidRequest}),
		}

		service := post.NewInsertPostService(spy)
		res, err := service.Run(post.NewInsertPostRequest("This is a test post", 1, 0))

		spy.CalledOnce(t)
		if res != nil {
			t.Error("expected error, got nil\n")
		}
		if err != application.ErrInvalidRequest {
			t.Errorf("got %s as error, wanted %s\n", err.Error(), application.ErrInvalidRequest.Error())
		}
	})
}
