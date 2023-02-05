package auth

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mddg/go-sm/server/domain/shared"
)

var jwtExpire = 24 * time.Hour

type Claim struct{ jwt.RegisteredClaims }

func NewClaim(
	id int, issuer string, audience []string, t time.Time,
) Claim {
	numericDate := jwt.NewNumericDate(t)
	claim := Claim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(t.Add(jwtExpire)),
			IssuedAt:  numericDate,
			NotBefore: numericDate,
			ID:        strconv.Itoa(id),
		},
	}
	if issuer != "" {
		claim.Issuer = issuer
	}
	if len(audience) > 0 {
		claim.Audience = audience
	}

	return claim
}

func (c Claim) Sign() (string, error) {
	signingKey := os.Getenv(shared.JwtSigningKey)
	if signingKey == "" {
		return "", errors.New("no env jwt signing key provided")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(signingKey))
}
