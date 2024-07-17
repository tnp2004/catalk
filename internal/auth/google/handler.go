package google

import (
	"catalk/utils"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
)

var (
	googleConfig  GoogleOAuth
	stateInstance string
)

func GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	googleConfig = GoogleConfig()

	stateInstance = utils.GenerateRandomString(10)
	url := googleConfig.GoogleLoginConfig.AuthCodeURL(stateInstance)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
func GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != stateInstance {
		log.Printf("error invalid oauth state, expected %s, got %s", "randomstate", state)
		utils.ErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("invalid oauth state"))
		return
	}

	code := r.FormValue("code")
	token, err := googleConfig.GoogleLoginConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("error code exchange failed. Error: %s", err.Error())
		utils.ErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("code exchange failed"))
		return
	}

	googleApiUrl := fmt.Sprintf("%s/oauth2/v2/userinfo?access_token=%s", googleConfig.GoogleConfig.GoogleApiUrl, token.AccessToken)
	response, err := http.Get(googleApiUrl)
	if err != nil {
		log.Printf("error get user data failed. Error: %s", err.Error())
		utils.ErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("get user data failed"))
		defer response.Body.Close()
		return
	}

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("error read response data failed. Error: %s", err.Error())
		utils.ErrorResponse(w, http.StatusInternalServerError, fmt.Errorf("read response data failed"))
	}
	fmt.Fprintf(w, "Content: %s\n", contents)
}
