package rest

import (
	"sync"

	"github.com/gin-gonic/gin"
	postRouter "github.com/mddg/go-sm/server/infrastructure/rest/domain/post"
	userRouter "github.com/mddg/go-sm/server/infrastructure/rest/domain/user"
)

var r *gin.Engine
var lock = sync.Mutex{}

func NewRestRouter() *gin.Engine {
	lock.Lock()
	defer lock.Unlock()

	if r != nil {
		return r
	}

	r = gin.Default()

	// Attach user router
	userRouter.Attach(r)
	postRouter.Attach(r)

	return r
}
