package auth_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mddg/go-sm/server/domain/auth"
	"github.com/mddg/go-sm/server/domain/shared"
)

var issuer = "https://api.go-sm.com"
var audience = []string{"https://go-sm.com"}

func validateField(t *testing.T, field string, got, want interface{}) {
	if got != want {
		t.Errorf("expected %s to be %s, got %s", field, want, got)
	}
}

func validateTime(t *testing.T, field string, got, want *jwt.NumericDate) {
	if *got != *want {
		t.Errorf("%s time does not match, got %s, want %s\n", field, got, want)
	}
}

func TestClaim(t *testing.T) {
	t.Run("should create a new base claim", func(t *testing.T) {
		d := time.Now()
		c := auth.NewClaim(1, issuer, audience, d)

		validateField(t, "ID", c.ID, "1")
		validateField(t, "Issuer", c.Issuer, issuer)
		for i, caud := range c.Audience {
			validateField(t, fmt.Sprintf("Audience[%d]", i), caud, audience[i])
		}

		validateTime(t, "ExpiresAt", c.ExpiresAt, jwt.NewNumericDate(d.Add(24*time.Hour)))
		numDate := jwt.NewNumericDate(d)
		validateTime(t, "IssuedAt", c.IssuedAt, numDate)
		validateTime(t, "NotBefore", c.NotBefore, numDate)
	})

	t.Run("should create claim with empty issuer", func(t *testing.T) {
		c := auth.NewClaim(1, "", audience, time.Now())
		validateField(t, "Issuer", c.Issuer, "")
	})

	t.Run("should create claim with empty audience", func(t *testing.T) {
		c := auth.NewClaim(1, "", nil, time.Now())
		if len(c.Audience) != 0 {
			t.Errorf("expected Audience to be empty, got size %d\n", len(c.Audience))
		}
	})
}

func TestSign(t *testing.T) {
	t.Run("should properly sign the token", func(t *testing.T) {
		t.Setenv(shared.JwtSigningKey, "secret-key")
		c := auth.NewClaim(1, issuer, audience, time.Now())
		token, err := c.Sign()
		if err != nil {
			t.Errorf("was not expecting error, got %s\n", err)
		}
		if token == "" {
			t.Error("signed token is empty")
		}

		_, err = jwt.ParseWithClaims(token, &auth.Claim{}, func(t *jwt.Token) (interface{}, error) {
			return []byte("secret-key"), nil
		})
		if err != nil {
			t.Errorf("error was returned when parsing the signed claims, got: %s\n", err)
		}
	})

	t.Run("should check if jwt token has been set", func(t *testing.T) {
		t.Setenv(shared.JwtSigningKey, "")
		c := auth.NewClaim(1, issuer, audience, time.Now())
		token, err := c.Sign()
		if err == nil {
			t.Errorf("expected error, got token: %s\n", token)
		}
	})
}
