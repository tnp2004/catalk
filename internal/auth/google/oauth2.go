package google

import (
	"catalk/config"
	"sync"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleOAuth struct {
	GoogleLoginConfig oauth2.Config
	GoogleConfig      *config.Google
}

var (
	once                 sync.Once
	googleConfigInstance GoogleOAuth
)

func GoogleConfig() GoogleOAuth {
	once.Do(func() {
		config := config.GetConfig().Google
		googleConfigInstance = GoogleOAuth{
			oauth2.Config{
				ClientID:     config.OAuth.ClientID,
				ClientSecret: config.OAuth.ClientSecret,
				Endpoint:     google.Endpoint,
				RedirectURL:  config.OAuth.RedirectURL,
				Scopes:       config.OAuth.Scopes,
			},
			config,
		}
	})

	return googleConfigInstance
}
