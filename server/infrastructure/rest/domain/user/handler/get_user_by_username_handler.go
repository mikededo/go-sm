package handler

import (
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	services "github.com/mddg/go-sm/server/application/user"
	"github.com/mddg/go-sm/server/domain/util"
	"github.com/mddg/go-sm/server/infrastructure/db"
	"github.com/mddg/go-sm/server/infrastructure/db/repository"
)

type GetUserByUsernameResponse struct {
	ID          int       `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Description string    `json:"description"`
	WebsiteURL  string    `json:"website_url"`
	CreatedAt   time.Time `json:"created_at"`
	BirthDate   time.Time `json:"birth_date"`
}

type GetUserByUsernameHandler struct{}

func (GetUserByUsernameHandler) Handle(ctx *gin.Context) {
	// get the param
	username := ctx.Param("username")
	re := regexp.MustCompile(`^[a-zA-Z0-9]*$`)
	if !re.MatchString(username) {
		// invalid username
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username empty or invalid format"})
	}

	conn := db.Factory(db.MysqlDB)
	r := repository.NewGormUserRepository(conn)

	// find the user
	user, err := services.NewFindUserByUsernameService(r).Run(username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var response GetUserByUsernameResponse
	util.MergeStructs(user, &response)
	ctx.JSON(http.StatusOK, gin.H{"data": response})
}
