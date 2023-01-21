package user

import (
	"fmt"
	"testing"

	"github.com/mddg/go-sm/server/application"
	"github.com/mddg/go-sm/server/domain/user"
)

func validateFindByUserIdCalls(t *testing.T, s *UserRepositorySpy) {
	s.CalledOnce(t)
	if s.Calls[0] != 1 {
		t.Errorf("repository expected to be called with %d received %d\n", 1, s.Calls[0])
	}
}

func TestFindUserByIdService_Run(t *testing.T) {
	t.Run("user found", func(t *testing.T) {
		u := &user.User{ID: 1, Username: "mikededo"}
		spy := &UserRepositorySpy{
			RepositorySpy: application.NewRepositoryWithResults([]*user.User{u}),
		}

		service := NewFindUserByIdService(spy)
		res, err := service.Run(1)

		validateFindByUserIdCalls(t, spy)
		if err != nil {
			t.Errorf("error was not expected, got %v\n", err)
		}
		if res.ID != 1 || res.Username != "mikededo" {
			t.Errorf("expected %v user got %v", u, res)
		}
	})

	t.Run("user not found", func(t *testing.T) {
		spy := &UserRepositorySpy{
			RepositorySpy: application.NewRepositoryWithErrors[*user.User]([]error{fmt.Errorf("user not found")}),
		}

		service := NewFindUserByIdService(spy)
		res, err := service.Run(1)

		validateFindByUserIdCalls(t, spy)
		if err == nil {
			t.Errorf("expecting error got %v\n", err)
		}
		if res != nil {
			t.Errorf("user was not expected, got %v", res)
		}
	})
}
