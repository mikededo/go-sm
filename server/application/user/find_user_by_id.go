package user

import "github.com/mddg/go-sm/server/domain/user"

type FindUserByIdService struct {
	repository user.UserRepository
}

func NewFindUserByIdService(repository user.UserRepository) *FindUserByIdService {
	return &FindUserByIdService{repository}
}

func (s *FindUserByIdService) Run(id int) (*user.User, error) {
	return s.repository.FindUserById(id)
}
