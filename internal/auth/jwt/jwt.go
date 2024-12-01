package jwt

import (
	"fmt"
	"time"

	gojwt "github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your-secret-key")

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	gojwt.RegisteredClaims
}

func GenerateJWT(username string) (*string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		Username: username,
		RegisteredClaims: gojwt.RegisteredClaims{
			ExpiresAt: gojwt.NewNumericDate(expirationTime),
			Issuer:    "library-mgmt-api",
		},
	}

	token := gojwt.NewWithClaims(gojwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtKey)

	if err != nil {
		return nil, fmt.Errorf("could not sign token, err=%v", err)
	}

	return &signedToken, nil
}
