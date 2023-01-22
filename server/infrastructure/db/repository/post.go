package repository

import (
	"github.com/mddg/go-sm/server/domain/post"
	"github.com/mddg/go-sm/server/domain/shared"
	postSchema "github.com/mddg/go-sm/server/infrastructure/db/mysql/schema"
	"gorm.io/gorm"
)

type GormPostRepository struct {
	conn *gorm.DB
}

func NewGormPostRepository(conn *gorm.DB) *GormPostRepository {
	return &GormPostRepository{conn}
}

func (r *GormPostRepository) InsertPost(p post.Post) (*post.Post, error) {
	// convert post into schema
	schema := postSchema.FromPost(p)
	// save the schema
	res := r.conn.Create(&schema)

	if res.Error != nil {
		return nil, res.Error
	}

	// convert to entity
	result := postSchema.PostFromSchema(schema)

	return &result, nil
}

func (r *GormPostRepository) FindUserPosts(id uint, page shared.PagedRequest) ([]*post.Post, error) {
	// get all posts
	posts := []*postSchema.Post{}
	query := r.conn.Limit(int(page.Limit)).Offset(int(page.Offset)).Joins("ResharedPost")
	findResult := query.Find(&posts, postSchema.Post{AuthorID: id})
	if findResult.Error != nil {
		return nil, findResult.Error
	}

	result := make([]*post.Post, 0)
	for _, schema := range posts {
		parsed := postSchema.PostFromSchema(*schema)
		result = append(result, &parsed)
	}

	return result, nil
}
