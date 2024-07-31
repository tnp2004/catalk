package jwt

import (
	"catalk/config"
	"catalk/internal/users"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTCustomClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
}

func CreateJWTToken(jwtConfig *config.JWT, userData *users.NewUserModel) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTCustomClaims{
		Email: userData.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userData.Username,
			Audience:  []string{"user"},
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtConfig.ExpireDuration * time.Second)),
		},
	})

	ss, err := token.SignedString([]byte(jwtConfig.SecretKey))
	if err != nil {
		return "", fmt.Errorf("error sign jwt token. Error: %s", err.Error())
	}

	return ss, nil
}
