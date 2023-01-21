package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	services "github.com/mddg/go-sm/server/application/user"
	"github.com/mddg/go-sm/server/infrastructure/db"
	"github.com/mddg/go-sm/server/infrastructure/db/repository"
)

type PostUserJson struct {
	FirstName string `form:"firstName" json:"firstName" binding:"required"`
	LastName  string `form:"lastName" json:"lastName" binding:"required"`
	Username  string `form:"username" json:"username" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
}

type PostUserHandler struct{}

func (PostUserHandler) Handle(ctx *gin.Context) {
	// get the data
	var json PostUserJson
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req := services.NewInsertUserRequest(
		json.FirstName,
		json.LastName,
		json.Username,
		json.Email,
		json.Password,
	)
	conn := db.DbFactory(db.MysqlDb)
	r := repository.NewGormUserRepository(conn)

	if err := services.NewInsertUserService(r).Run(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{})
}
