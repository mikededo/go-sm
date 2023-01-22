package user

import "github.com/mddg/go-sm/server/domain/user"

type FindUserByUsernameService struct {
	repository user.Repository
}

func NewFindUserByUsernameService(repository user.Repository) *FindUserByUsernameService {
	return &FindUserByUsernameService{repository}
}

func (s *FindUserByUsernameService) Run(username string) (*user.User, error) {
	return s.repository.FindUserByUsername(username)
}
