package middleware_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mddg/go-sm/server/domain/shared"
	"github.com/mddg/go-sm/server/infrastructure/rest/domain/middleware"
)

var validToken = strings.Join([]string{
	"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.",
	"eyJpc3MiOiJodHRwczovL2dvLXNtLmNvbSIsImV4cCI6NDk0ODM0ODA1Niw",
	"ibmJmIjoxNjc1NjI0MDU2LCJpYXQiOjE2NzU2MjQwNTYsImp0aSI6IjEifQ",
	"._jL4WcgFNhNsNzXXGXEjeYXjvBV280nYwucn9J0R8VY",
}, "")
var invalidToken = strings.Join([]string{
	"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.",
	"eyJpc3MiOiJodHRwczovL2dvLXNtLmNvbSIsImV4cCI6MTY3NTUzNDUxMy",
	"wibmJmIjoxNjc1NDQ4MTEzLCJpYXQiOjE2NzU0NDgxMTMsImp0aSI6IjEifQ",
	".cNsRmXP-WiV9X9R9LVGJMeFjEVRM9fwnmQZBAw-pWlY",
}, "")
var invalidClaimsToken = strings.Join([]string{
	"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.",
	"eyJpc3MiOiJodHRwczovL2dvLXNtLmNvbSIsImV4cCI6NDk0ODM0ODA1Niwib",
	"mJmIjoxNjc1NjI0MDU2LCJpYXQiOjE2NzU2MjQwNTYsImp0aSI6IjEiLCJyYW",
	"5kb21fa2V5IjoicmFuZG9tX3ZhbHVlIn0",
	".pWNjy_kYAgl69U6I9OjSDkXcZ5H32TedG2yuR44hmdE",
}, "")

func setupRouter(t *testing.T) *gin.Engine {
	t.Setenv(shared.JwtSigningKey, shared.TestJwtSignValue)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(middleware.JwtAuthentication())
	router.GET("/jwt", func(c *gin.Context) {
		c.Status(200)
	})

	return router
}

func callJwtEndpoint(router *gin.Engine, token string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/jwt", nil)
	if token != "nil" {
		req.Header.Add("Authentication", "Bearer "+token)
	}
	router.ServeHTTP(w, req)

	return w
}

func validateErrorWithBody(t *testing.T, w *httptest.ResponseRecorder, code int, msg string) {
	if w.Code != code {
		t.Errorf("expected code 403, got %d\n", w.Code)
	}

	// extract the json body
	body := make(map[string]string)
	err := json.NewDecoder(w.Body).Decode(&body)
	if err != nil {
		t.Errorf("error on decoding json response: %v\n", err)
	}

	if body["error"] != msg {
		t.Errorf("expected %s as body, got %v\n", msg, w.Body)
	}
}

func TestJwtMiddleware(t *testing.T) {
	t.Run("should call the next method with a valid jwt token", func(t *testing.T) {
		w := callJwtEndpoint(setupRouter(t), validToken)

		if w.Code != 200 {
			t.Errorf("expected code 200, got %d\n", w.Code)
		}
	})

	t.Run("should not call next on invalid token", func(t *testing.T) {
		w := callJwtEndpoint(setupRouter(t), invalidToken)

		if w.Code == 200 {
			t.Errorf("expected code 200, got %d\n", w.Code)
		}
	})

	t.Run("should not call next no authentication header", func(t *testing.T) {
		w := callJwtEndpoint(setupRouter(t), "nil")
		validateErrorWithBody(t, w, 403, middleware.JwtNoAuthenticationHeader)
	})

	t.Run("should not call next no authentication header", func(t *testing.T) {
		w := callJwtEndpoint(setupRouter(t), "")
		validateErrorWithBody(t, w, 403, middleware.JwtNoAuthenticationToken)
	})
}
