package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mddg/go-sm/server/domain/auth"
	sharedDomain "github.com/mddg/go-sm/server/domain/shared"
	"github.com/mddg/go-sm/server/infrastructure/rest/domain/shared"
)

const (
	JwtNoAuthenticationHeader = "Could not find authentication header"
	JwtNoAuthenticationToken  = "Could not find authentication token"
)

func abort(ctx *gin.Context, err string) {
	ctx.JSON(http.StatusForbidden, gin.H{"error": err})
	ctx.Abort()
}

func JwtAuthentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authentication")
		if authHeader == "" {
			abort(ctx, JwtNoAuthenticationHeader)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			abort(ctx, JwtNoAuthenticationToken)
			return
		}

		_, err := jwt.ParseWithClaims(token, &auth.Claim{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv(sharedDomain.JwtSigningKey)), nil
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": shared.InternalServerErrorMessage})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
