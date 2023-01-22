package user

import (
	"github.com/mddg/go-sm/server/application"
	"github.com/mddg/go-sm/server/domain/user"
	"github.com/mddg/go-sm/server/domain/util"
	"golang.org/x/crypto/bcrypt"
)

type InsertUserRequest struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Username  string `validate:"required"`
	Email     string `validate:"required"`
	Password  string `validate:"required"`
}

type InsertUserService struct {
	repository user.Repository
}

func NewInsertUserService(repository user.Repository) *InsertUserService {
	return &InsertUserService{repository}
}

func NewInsertUserRequest(
	firstName, lastName, username, email, password string,
) InsertUserRequest {
	return InsertUserRequest{
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Email:     email,
		Password:  password,
	}
}

func (s *InsertUserService) Run(req InsertUserRequest) error {
	err := util.Validator().Struct(req)
	if err != nil {
		return application.ErrInvalidRequest
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return application.ErrInvalidRequest
	}

	user := user.NewUnregisteredUser(
		req.FirstName,
		req.LastName,
		req.Username,
		req.Email,
		string(hashedPassword),
	)
	err = s.repository.InsertUser(user)
	if err != nil {
		return err
	}

	return nil
}
