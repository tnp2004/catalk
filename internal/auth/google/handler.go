package google

import (
	"catalk/config"
	"catalk/internal/auth/jwt"
	"catalk/internal/users"
	"catalk/utils"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/oauth2"
)

var stateInstance string

type googleOAuth struct {
	oauthConfig    oauth2.Config
	googleConfig   *config.Google
	databaseConfig *config.Database
	jwtConfig      *config.JWT
}

type GoogleOAuthService interface {
	GoogleLoginHandler(w http.ResponseWriter, r *http.Request)
	GoogleCallbackHandler(w http.ResponseWriter, r *http.Request)
}

func NewGoogleOAuth(oauthConfig oauth2.Config, googleConfig *config.Google, databaseConfig *config.Database, jwtConfig *config.JWT) GoogleOAuthService {
	return &googleOAuth{oauthConfig, googleConfig, databaseConfig, jwtConfig}
}

func (a *googleOAuth) GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	stateInstance = utils.GenerateRandomString(10)
	url := a.oauthConfig.AuthCodeURL(stateInstance)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (a *googleOAuth) GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != stateInstance {
		log.Printf("error invalid oauth state, expected %s, got %s", stateInstance, state)
		utils.ErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("invalid oauth state"))
		return
	}

	code := r.FormValue("code")
	token, err := a.oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("error code exchange failed. Error: %s", err.Error())
		utils.ErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("code exchange failed"))
		return
	}

	googleApiUrl := fmt.Sprintf("%s/oauth2/v2/userinfo?access_token=%s", a.googleConfig.GoogleApiUrl, token.AccessToken)
	response, err := http.Get(googleApiUrl)
	if err != nil {
		log.Printf("error get user data failed. Error: %s", err.Error())
		utils.ErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("get user data failed"))
		defer response.Body.Close()
		return
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("error read response data failed. Error: %s", err.Error())
		utils.ErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("read response data failed"))
	}
	reqBody := new(users.NewGoogleUserModel)
	if err := json.Unmarshal(data, &reqBody); err != nil {
		log.Printf("error unmarshal data. Error: %s", err.Error())
		utils.ErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("google callback failed"))
		return
	}

	user := users.NewUser(a.databaseConfig)
	userData, err := user.FindUserByEmail(reqBody.Email)
	if err != nil {
		// create account
		reqBody.ProviderID = users.Provider.Google
		if err := user.InsertUser((*users.NewUserModel)(reqBody)); err != nil {
			utils.ErrorResponse(w, http.StatusInternalServerError, err)
			return
		}
	}

	// jwt token
	ss, err := jwt.CreateJWTToken(a.jwtConfig, userData)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	setCookie(w, ss)

	utils.MessageResponse(w, http.StatusOK, "login successfully")
}

func setCookie(w http.ResponseWriter, tokenString string) {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		MaxAge:   3600,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, cookie)
}
