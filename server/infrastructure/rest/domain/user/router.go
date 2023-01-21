package user

import (
	"github.com/gin-gonic/gin"
	"github.com/mddg/go-sm/server/infrastructure/rest/domain/user/handler"
)

func Attach(r *gin.Engine) {
	// Create a new user
	r.POST("/user", handler.PostUserHandler{}.Handle)
	// Find user by username
	r.GET("/user/:username", handler.GetUserByUsernameHandler{}.Handle)
}
