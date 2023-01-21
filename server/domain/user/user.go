package user

import "time"

type User struct {
	ID          int
	FirstName   string
	LastName    string
	Username    string
	Email       string
	Password    string
	Description string
	WebsiteURL  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	BirthDate   time.Time
}

func NewUser(
	id int,
	firstName, lastName, username, email, passsword, description, websiteURL string,
	createdAt, updatedAt, birthDate time.Time,
) User {
	return User{
		ID:          id,
		FirstName:   firstName,
		LastName:    lastName,
		Username:    username,
		Email:       email,
		Password:    passsword,
		Description: description,
		WebsiteURL:  websiteURL,
		BirthDate:   birthDate,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

func NewPublicUser(
	id int,
	firstName, lastName, username, email, description, websiteURL string,
	createdAt, birthDate time.Time,
) User {
	return User{
		ID:          id,
		FirstName:   firstName,
		LastName:    lastName,
		Username:    username,
		Email:       email,
		Description: description,
		WebsiteURL:  websiteURL,
		BirthDate:   birthDate,
		CreatedAt:   createdAt,
	}
}

func NewUnregisteredUser(
	firstName, lastName, username, email, passsword string,
) User {
	return User{
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Email:     email,
		Password:  passsword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
