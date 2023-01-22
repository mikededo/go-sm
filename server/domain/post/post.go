package post

import "time"

type Author struct {
	ID       int
	Username string
}

type Collaborator struct {
	ID       int
	Username string
}

// Test comment
type Post struct {
	ID            int
	Author        Author
	Content       string
	IsReshared    bool
	ResharedId    int
	ResharedPost  *Post
	Likes         int
	Dislikes      int
	Collaborators []Collaborator
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func parseCollaborators(collaborators map[int]string) []Collaborator {
	result := make([]Collaborator, 0)
	for id, name := range collaborators {
		result = append(result, Collaborator{ID: id, Username: name})
	}
	return result
}

func NewPost(
	id, authorId int,
	authorName, content string,
	isReshared bool,
	resharedId int, resharedPost *Post,
	likes, dislikes int,
	collaboratorsMap map[int]string,
	createdAt, updatedAt time.Time,
) Post {
	return Post{
		ID:            id,
		Author:        Author{ID: authorId, Username: authorName},
		Content:       content,
		IsReshared:    isReshared,
		ResharedId:    resharedId,
		ResharedPost:  resharedPost,
		Likes:         likes,
		Dislikes:      dislikes,
		Collaborators: parseCollaborators(collaboratorsMap),
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}
}

func NewEmptyPost(
	authorId int,
	authorName, content string,
	isReshared bool,
	resharedId, likes, dislikes int,
	collaboratorsMap map[int]string,
) Post {
	return Post{
		Author:        Author{ID: authorId, Username: authorName},
		Content:       content,
		IsReshared:    isReshared,
		ResharedId:    resharedId,
		ResharedPost:  nil,
		Likes:         likes,
		Dislikes:      dislikes,
		Collaborators: parseCollaborators(collaboratorsMap),
	}
}
