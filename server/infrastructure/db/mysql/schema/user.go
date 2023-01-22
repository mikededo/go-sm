package schema

import (
	"time"

	"github.com/mddg/go-sm/server/domain/user"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName   string    `gorm:"first_name;type:varchar(75);not null"`
	LastName    string    `gorm:"last_name;type:varchar(75);not null"`
	Username    string    `gorm:"username;type:varchar(75);unique;not null"`
	Email       string    `gorm:"email;type:varchar(255);unique;not null"`
	Password    string    `gorm:"password;type:varchar(255);not null"`
	Description string    `gorm:"description;type:text"`
	WebsiteURL  string    `gorm:"website_url;type:varchar(255)"`
	BirthDate   time.Time `gorm:"birth_date;default:null"`
}

func AttachUserToDatabase(db *gorm.DB) {
	err := db.AutoMigrate(&User{})
	if err != nil {
		panic("unable to auto migrate User entity")
	}
}

func FromUser(in user.User) User {
	return User{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Username:  in.Username,
		Email:     in.Email,
		Password:  in.Password,
	}
}

func UserFromSchema(in User) user.User {
	return user.NewPublicUser(
		int(in.ID),
		in.FirstName,
		in.LastName,
		in.Username,
		in.Email,
		in.Description,
		in.WebsiteURL,
		in.BirthDate,
		in.CreatedAt,
	)
}
