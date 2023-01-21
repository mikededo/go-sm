package user

type Repository interface {
	// FindUserById returns a User given it's id or returns nil if not found
	FindUserById(int) (*User, error)

	// FindUserByUsername returns a User given it's username or returns nil if not found
	FindUserByUsername(string) (*User, error)

	// InsertUser saves a new User
	InsertUser(User) error
}
