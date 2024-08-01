package jwt

import (
	"catalk/config"
	"catalk/internal/users"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTCustomClaims struct {
	jwt.RegisteredClaims
}

type JWTUserData struct {
	ID string
}

func CreateJWTToken(jwtConfig *config.JWT, userData *users.UserEntity) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    userData.ID,
			Audience:  []string{"user"},
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(jwtConfig.ExpireDuration * time.Second)),
		},
	})

	ss, err := token.SignedString([]byte(jwtConfig.SecretKey))
	if err != nil {
		log.Printf("error sign jwt token. Error: %s", err.Error())
		return "", fmt.Errorf("create jwt token failed")
	}

	return ss, nil
}

func ValidateToken(jwtConfig *config.JWT, tokenString string) (*JWTUserData, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtConfig.SecretKey), nil
	})
	if err != nil {
		log.Printf("erorr parse with claims token. Error: %s", err.Error())
		return nil, fmt.Errorf("erorr validate token")
	}

	claims, ok := token.Claims.(*JWTCustomClaims)
	if !ok {
		return nil, fmt.Errorf("claims token failed")
	}

	return &JWTUserData{
		ID: claims.ID,
	}, nil
}
