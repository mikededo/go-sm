package user_test

import (
	"github.com/mddg/go-sm/server/application"
	"github.com/mddg/go-sm/server/domain/user"
)

type UserRepositorySpy struct {
	application.RepositorySpy[*user.User]
}

func (spy *UserRepositorySpy) FindUserByID(id int) (*user.User, error) {
	spy.Calls = append(spy.Calls, id)
	return spy.Result(), spy.Error()
}

func (spy *UserRepositorySpy) FindUserByUsername(username string) (*user.User, error) {
	spy.Calls = append(spy.Calls, username)
	return spy.Result(), spy.Error()
}

func (spy *UserRepositorySpy) InsertUser(u user.User) error {
	spy.Calls = append(spy.Calls, u)
	return spy.Error()
}
