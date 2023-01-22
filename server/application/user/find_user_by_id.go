package user

import "github.com/mddg/go-sm/server/domain/user"

type FindUserByIDService struct {
	repository user.Repository
}

func NewFindUserByIDService(repository user.Repository) *FindUserByIDService {
	return &FindUserByIDService{repository}
}

func (s *FindUserByIDService) Run(id int) (*user.User, error) {
	return s.repository.FindUserByID(id)
}
