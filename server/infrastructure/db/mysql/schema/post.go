package schema

import (
	"github.com/mddg/go-sm/server/domain/post"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Content      string `gorm:"content;type:varchar(250);not null"`
	AuthorID     uint   `gorm:"author;not null"`
	Author       User   `gorm:"foreignKey:AuthorId"`
	IsReshared   bool   `gorm:"is_reshared;type:boolean;not null;default:0"`
	ResharedID   *uint  `gorm:"reshared_id"`
	ResharedPost *Post  `gorm:"foreignKey:ResharedId"`
	Likes        uint   `gorm:"likes;type:int;default:0"`
	Dislikes     uint   `gorm:"dislikes;type:int;default:0"`
}

func AttachPostToDatabase(db *gorm.DB) {
	db.AutoMigrate(&Post{})
}

func SchemaFromPost(in post.Post) Post {
	var resharedID *uint = nil
	if in.ResharedID != 0 {
		temp := uint(in.ResharedID)
		resharedID = &temp
	}

	return Post{
		Content:    in.Content,
		AuthorID:   uint(in.Author.ID),
		IsReshared: in.IsReshared,
		ResharedID: resharedID,
		Likes:      uint(in.Likes),
		Dislikes:   uint(in.Dislikes),
	}
}

func PostFromSchema(in Post) post.Post {
	resharedID := 0
	var resharedPost post.Post
	if in.ResharedPost != nil {
		resharedID = int(in.ResharedPost.ID)
		resharedPost = PostFromSchema(*in.ResharedPost)
	}

	return post.NewPost(
		int(in.ID),
		int(in.Author.ID),
		in.Author.Username,
		in.Content,
		in.IsReshared,
		resharedID,
		&resharedPost,
		int(in.Likes),
		int(in.Dislikes),
		nil,
		in.CreatedAt,
		in.UpdatedAt,
	)
}
