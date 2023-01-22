package repository

import (
	"github.com/mddg/go-sm/server/domain/user"
	userSchema "github.com/mddg/go-sm/server/infrastructure/db/mysql/schema"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	conn *gorm.DB
}

func NewGormUserRepository(conn *gorm.DB) *GormUserRepository {
	return &GormUserRepository{conn}
}

func (r *GormUserRepository) FindUserByID(int) (*user.User, error) {
	return nil, nil
}

func (r *GormUserRepository) FindUserByUsername(username string) (*user.User, error) {
	var schema userSchema.User
	// search for the entity
	err := r.conn.Model(&userSchema.User{}).Where("username = ?", username).First(&schema).Error
	if err != nil {
		return nil, err
	}

	// get the entity
	entity := userSchema.UserFromSchema(schema)
	return &entity, nil
}

func (r *GormUserRepository) InsertUser(u user.User) error {
	// convert user into schema
	schema := userSchema.FromUser(u)
	// save the schema
	res := r.conn.Save(&schema)
	// return err or nil
	return res.Error
}
