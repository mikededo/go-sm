package schema

import (
	"github.com/mddg/go-sm/server/domain/post"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Content      string `gorm:"content;type:varchar(250);not null"`
	AuthorId     uint   `gorm:"author;not null"`
	Author       User   `gorm:"foreignKey:AuthorId"`
	IsReshared   bool   `gorm:"is_reshared;type:boolean;not null;default:0"`
	ResharedId   *uint  `gorm:"reshared_id"`
	ResharedPost *Post  `gorm:"foreignKey:ResharedId"`
	Likes        uint   `gorm:"likes;type:int;default:0"`
	Dislikes     uint   `gorm:"dislikes;type:int;default:0"`
}

func AttachPostToDatabase(db *gorm.DB) {
	db.AutoMigrate(&Post{})
}

func SchemaFromPost(in post.Post) Post {
	var resharedId *uint = nil
	if in.ResharedId != 0 {
		temp := uint(in.ResharedId)
		resharedId = &temp
	}

	return Post{
		Content:    in.Content,
		AuthorId:   uint(in.Author.ID),
		IsReshared: in.IsReshared,
		ResharedId: resharedId,
		Likes:      uint(in.Likes),
		Dislikes:   uint(in.Dislikes),
	}
}

func PostFromSchema(in Post) post.Post {
	resharedId := 0
	var resharedPost post.Post
	if in.ResharedPost != nil {
		resharedId = int(in.ResharedPost.ID)
		resharedPost = PostFromSchema(*in.ResharedPost)
	}

	return post.NewPost(
		int(in.ID),
		int(in.Author.ID),
		in.Author.Username,
		in.Content,
		in.IsReshared,
		resharedId,
		&resharedPost,
		int(in.Likes),
		int(in.Dislikes),
		nil,
		in.CreatedAt,
		in.UpdatedAt,
	)
}
