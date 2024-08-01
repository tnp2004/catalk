package middleware

import (
	"catalk/config"
	"catalk/internal/auth/jwt"
	"catalk/utils"
	"fmt"
	"log"
	"net/http"
)

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		config := config.GetConfig()
		if err != nil {
			log.Printf("error token not found. Error: %s", err.Error())
			utils.ErrorResponse(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
			return
		}

		userData, err := jwt.ValidateToken(config.JWT, token.Value)
		if err != nil {
			log.Printf("error validate token. Error: %s", err.Error())
			utils.ErrorResponse(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
			return
		}

		// attach user id in header
		r.Header.Add("id", userData.ID)

		next.ServeHTTP(w, r)
	})
}
