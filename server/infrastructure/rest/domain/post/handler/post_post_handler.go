package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	services "github.com/mddg/go-sm/server/application/post"
	"github.com/mddg/go-sm/server/domain/util"
	"github.com/mddg/go-sm/server/infrastructure/db"
	"github.com/mddg/go-sm/server/infrastructure/db/repository"
)

type PostPostJSON struct {
	Content    string `form:"content" json:"content"`
	AuthorID   int    `form:"author" json:"author" binding:"required"`
	ResharedID int    `form:"reshared_id" json:"reshared_id"`
}

type PostPostResponse struct {
	ID         int       `json:"id"`
	Content    string    `json:"content"`
	IsReshared bool      `json:"is_reshared"`
	ResharedID int       `json:"reshared_id"`
	Likes      int       `json:"likes"`
	Dislikes   int       `json:"dislikes"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type PostPostHandler struct{}

func (PostPostHandler) Handle(ctx *gin.Context) {
	// get the data
	var json PostPostJSON
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// extra validations on json received
	if json.Content == "" && json.ResharedID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "cannot create post with empty body if it is not reshared"})
		return
	}

	req := services.NewInsertPostRequest(
		json.Content,
		json.AuthorID,
		json.ResharedID,
	)
	conn := db.Factory(db.MysqlDB)
	r := repository.NewGormPostRepository(conn)

	res, err := services.NewInsertPostService(r).Run(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var response PostPostResponse
	util.MergeStructs(res, &response)
	ctx.JSON(http.StatusCreated, gin.H{"data": response})
}
