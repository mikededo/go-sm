package user_test

import (
	"fmt"
	"testing"

	"github.com/mddg/go-sm/server/application"
	"github.com/mddg/go-sm/server/application/user"
	userEntity "github.com/mddg/go-sm/server/domain/user"
)

func validateFindByUserUsernameCalls(t *testing.T, s *UserRepositorySpy) {
	s.CalledOnce(t)
	if s.Calls[0] != "mikededo" {
		t.Errorf("repository expected to be called with %s received %s\n", "mikededo", s.Calls[0])
	}
}

func TestFindUserByUsernameService_Run(t *testing.T) {
	t.Run("user found", func(t *testing.T) {
		u := &userEntity.User{ID: 1, Username: "mikededo"}
		spy := &UserRepositorySpy{
			RepositorySpy: application.NewRepositoryWithResults([]*userEntity.User{u}),
		}

		service := user.NewFindUserByUsernameService(spy)
		res, err := service.Run("mikededo")

		validateFindByUserUsernameCalls(t, spy)
		if err != nil {
			t.Errorf("error was not expected, got %v\n", err)
		}
		if res.ID != 1 || res.Username != "mikededo" {
			t.Errorf("expected %v user got %v", u, res)
		}
	})

	t.Run("user not found", func(t *testing.T) {
		spy := &UserRepositorySpy{
			RepositorySpy: application.NewRepositoryWithErrors[*userEntity.User]([]error{fmt.Errorf("user not found")}),
		}

		service := user.NewFindUserByUsernameService(spy)
		res, err := service.Run("mikededo")

		validateFindByUserUsernameCalls(t, spy)
		if err == nil {
			t.Errorf("expecting error got %v\n", err)
		}
		if res != nil {
			t.Errorf("user was not expected, got %v", res)
		}
	})
}
