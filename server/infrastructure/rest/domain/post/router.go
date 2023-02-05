package post

import (
	"github.com/gin-gonic/gin"
	"github.com/mddg/go-sm/server/infrastructure/rest/domain/middleware"
	"github.com/mddg/go-sm/server/infrastructure/rest/domain/post/handler"
)

func Attach(r *gin.Engine) {
	// Create a new post
	r.POST("/post", handler.PostPostHandler{}.Handle, middleware.JwtAuthentication())
	// Find user posts given its id
	r.GET("/posts/:user_id", handler.GetUserPostsHandler{}.Handle)
}
