package post_test

import (
	"errors"
	"testing"

	"github.com/mddg/go-sm/server/application"
	"github.com/mddg/go-sm/server/application/post"
	postEntity "github.com/mddg/go-sm/server/domain/post"
	"github.com/mddg/go-sm/server/domain/shared"
	"github.com/stretchr/testify/assert"
)

func validateFindUserPostCalls(t *testing.T, spy *PostRepositorySpy, args FindUserPostsArguments) {
	callArgs, ok := spy.Calls[0].(FindUserPostsArguments)
	assert.True(t, ok, "cannot cast to 'FindUserPostsArguments'")
	application.CheckPopertyEquality(t, "ID", args.ID, callArgs.ID)
	application.CheckPopertyEquality(t, "PageRequest.Page", args.PageRequest.Page, callArgs.PageRequest.Page)
	application.CheckPopertyEquality(t, "PageRequest.Limit", args.PageRequest.Limit, callArgs.PageRequest.Limit)
	application.CheckPopertyEquality(t, "PageRequest.Offset", args.PageRequest.Offset, callArgs.PageRequest.Offset)
}

func validateFindUserPostsResult(t *testing.T, got, want []*postEntity.Post) {
	assert.Equal(t, len(want), len(got), "results differ in length, got %d, want %d\n", len(got), len(want))
	for i := range got {
		application.CheckPopertyEquality(t, "ID", got[i].ID, want[i].ID)
		application.CheckPopertyEquality(t, "Content", got[i].Content, want[i].Content)
		application.CheckPopertyEquality(t, "Author.ID", got[i].Author.ID, want[i].Author.ID)
		application.CheckPopertyEquality(t, "Author.Username", got[i].Author.Username, want[i].Author.Username)
	}
}

func TestFindUserPostsService_Run(t *testing.T) {
	t.Run("return post list", func(t *testing.T) {
		author := postEntity.Author{ID: 1, Username: "mikededo"}
		result := []*postEntity.Post{
			{ID: 1, Content: "Post 1 content", Author: author},
			{ID: 2, Content: "Post 2 content", Author: author},
			{ID: 3, Content: "Post 3 content", Author: author},
			{ID: 4, Content: "Post 4 content", Author: author},
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
		assert.Nil(t, err, "not expecting error, got %v\n", err)
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
		assert.Equal(t, 0, len(res), "expected empty list, got list with len %d\n", len(res))
		assert.NotNil(t, err, "expected error, got nil\n")
		assert.Equal(t, err.Error(), repositoryError, "got %s as error, wanted %s\n", err.Error(), repositoryError)
	})
}
