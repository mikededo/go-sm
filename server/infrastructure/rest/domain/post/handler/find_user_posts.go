package handler

import (
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	services "github.com/mddg/go-sm/server/application/post"
	"github.com/mddg/go-sm/server/domain/shared"
	"github.com/mddg/go-sm/server/domain/util"
	"github.com/mddg/go-sm/server/infrastructure/db"
	"github.com/mddg/go-sm/server/infrastructure/db/repository"
)

type FindUserPostsResponse struct {
	ID           int                    `json:"id"`
	Content      string                 `json:"content"`
	AuthorID     int                    `json:"author_id"`
	IsReshared   bool                   `json:"is_reshared"`
	ResharedID   int                    `json:"reshared_id"`
	ResharedPost *FindUserPostsResponse `json:"reshared_post,omitempty"`
	Likes        int                    `json:"likes"`
	Dislikes     int                    `json:"dislikes"`
	CreatedAt    time.Time              `json:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at"`
}

type GetUserPostsHandler struct{}

func (GetUserPostsHandler) Handle(ctx *gin.Context) {
	// get the param
	param := ctx.Param("user_id")
	userID, err := strconv.Atoi(param)
	re := regexp.MustCompile("^[0-9]+$")
	if err != nil && re.MatchString(param) {
		// invalid param
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user id empty or invalid format"})
	}

	conn := db.Factory(db.MysqlDB)
	r := repository.NewGormPostRepository(conn)

	// find the posts
	page := shared.NewPagedRequest(25)
	posts, err := services.NewFindUserPostsService(r).Run(uint(userID), *page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := make([]FindUserPostsResponse, 0)
	for _, post := range posts {
		var temp FindUserPostsResponse
		util.MergeStructs(post, &temp)

		if post.IsReshared {
			var tempReshared FindUserPostsResponse
			util.MergeStructs(post.ResharedPost, &tempReshared)
			temp.ResharedPost = &tempReshared
		}

		response = append(response, temp)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": response})
}
