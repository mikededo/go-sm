package user_test

import (
	"errors"
	"testing"

	"github.com/mddg/go-sm/server/application"
	"github.com/mddg/go-sm/server/application/user"
	userEntity "github.com/mddg/go-sm/server/domain/user"
	"golang.org/x/crypto/bcrypt"
)

func validateInsertUserCalls(t *testing.T, s *UserRepositorySpy, req user.InsertUserRequest) {
	arg, ok := s.Calls[0].(userEntity.User)
	if !ok {
		t.Error("cannot cast to 'userEntity.User'")
	}

	application.CheckPopertyEquality(t, "FirstName", arg.FirstName, req.FirstName)
	application.CheckPopertyEquality(t, "LastName", arg.LastName, req.LastName)
	application.CheckPopertyEquality(t, "Username", arg.Username, req.Username)
	application.CheckPopertyEquality(t, "Email", arg.Email, req.Email)
}

func TestInsertUserService_Run(t *testing.T) {
	t.Run("register user", func(t *testing.T) {
		spy := &UserRepositorySpy{
			RepositorySpy: application.NewRepositoryWithErrors[*userEntity.User]([]error{nil}),
		}

		service := user.NewInsertUserService(spy)
		req := user.NewInsertUserRequest("Mike", "Dedo", "mikededo", "mike@dedo.com", "password")
		res := service.Run(req)

		spy.CalledOnce(t)
		if res != nil {
			t.Errorf("not expecting error, got %v\n", res)
		}
		validateInsertUserCalls(t, spy, req)

		resPassword := spy.Calls[0].(userEntity.User).Password
		err := bcrypt.CompareHashAndPassword([]byte(resPassword), []byte("password"))
		if err != nil {
			t.Errorf("expected password to be hashed, got %s\n", resPassword)
		}
	})

	t.Run("repository error thrown", func(t *testing.T) {
		repositoryError := "repository error"
		spy := &UserRepositorySpy{
			RepositorySpy: application.NewRepositoryWithErrors[*userEntity.User]([]error{errors.New(repositoryError)}),
		}

		service := user.NewInsertUserService(spy)
		err := service.Run(
			user.NewInsertUserRequest("Mike", "Dedo", "mikededo", "mike@dedo.com", "password"),
		)

		if err == nil {
			t.Error("expected error, got nil\n")
		}
		if err.Error() != repositoryError {
			t.Errorf("got %s as error, wanted %s\n", err.Error(), repositoryError)
		}
	})

	t.Run("invalid user request", func(t *testing.T) {
		spy := &UserRepositorySpy{
			RepositorySpy: application.NewRepositoryWithErrors[*userEntity.User]([]error{application.ErrInvalidRequest}),
		}

		service := user.NewInsertUserService(spy)
		err := service.Run(user.InsertUserRequest{})

		if err == nil {
			t.Error("expected error, got nil\n")
		}
		if !errors.Is(err, application.ErrInvalidRequest) {
			t.Errorf("got '%v' error, wanted '%v'\n", err, application.ErrInvalidRequest)
		}
	})
}
