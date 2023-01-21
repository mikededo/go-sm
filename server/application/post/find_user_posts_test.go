package post_test

import (
	"errors"
	"testing"

	"github.com/mddg/go-sm/server/application"
	"github.com/mddg/go-sm/server/application/post"
	postEntity "github.com/mddg/go-sm/server/domain/post"
	"github.com/mddg/go-sm/server/domain/shared"
)

func validateFindUserPostCalls(t *testing.T, spy *PostRepositorySpy, args FindUserPostsArguments) {
	callArgs := spy.Calls[0].(FindUserPostsArguments)
	application.CheckPopertyEquality(t, "ID", args.ID, callArgs.ID)
	application.CheckPopertyEquality(t, "PageRequest.Page", args.PageRequest.Page, callArgs.PageRequest.Page)
	application.CheckPopertyEquality(t, "PageRequest.Limit", args.PageRequest.Limit, callArgs.PageRequest.Limit)
	application.CheckPopertyEquality(t, "PageRequest.Offset", args.PageRequest.Offset, callArgs.PageRequest.Offset)
}

func validateFindUserPostsResult(t *testing.T, got, want []*postEntity.Post) {
	if len(got) != len(want) {
		t.Errorf("results differ in length, got %d, want %d\n", len(got), len(want))
		return
	}
	for i := range got {
		application.CheckPopertyEquality(t, "ID", got[i].ID, want[i].ID)
		application.CheckPopertyEquality(t, "Content", got[i].Content, want[i].Content)
		application.CheckPopertyEquality(t, "Author.ID", got[i].Author.ID, want[i].Author.ID)
		application.CheckPopertyEquality(t, "Author.Username", got[i].Author.Username, want[i].Author.Username)
	}
}

func TestFindUserPostsService_Run(t *testing.T) {
	t.Run("return post list", func(t *testing.T) {
		result := []*postEntity.Post{
			{ID: 1, Content: "Post 1 content", Author: postEntity.Author{ID: 1, Username: "mikededo"}},
			{ID: 2, Content: "Post 2 content", Author: postEntity.Author{ID: 1, Username: "mikededo"}},
			{ID: 3, Content: "Post 3 content", Author: postEntity.Author{ID: 1, Username: "mikededo"}},
			{ID: 4, Content: "Post 4 content", Author: postEntity.Author{ID: 1, Username: "mikededo"}},
		}
		spy := &PostRepositorySpy{
			RepositorySpy: application.NewRepositoryWithResultsAndErrors(
				[][]*postEntity.Post{result},
				[]error{nil},
			),
		}

		service := post.NewFindUserPostsService(spy)
		res, err := service.Run(1, *shared.NewPagedRequest(25))

		spy.CalledOnce(t)
		if err != nil {
			t.Errorf("not expecting error, got %v\n", res)
		}
		validateFindUserPostCalls(t, spy, FindUserPostsArguments{ID: 1, PageRequest: *shared.NewPagedRequest(25)})
		validateFindUserPostsResult(t, res, result)
	})

	t.Run("repository error", func(t *testing.T) {
		repositoryError := "repository error"
		spy := &PostRepositorySpy{
			RepositorySpy: application.NewRepositoryWithErrors[[]*postEntity.Post](
				[]error{errors.New(repositoryError)},
			),
		}

		service := post.NewFindUserPostsService(spy)
		res, err := service.Run(1, *shared.NewPagedRequest(25))

		spy.CalledOnce(t)
		if res != nil {
			t.Error("expected error, got nil\n")
		}
		if err.Error() != repositoryError {
			t.Errorf("got %s as error, wanted %s\n", err.Error(), repositoryError)
		}
	})
}
