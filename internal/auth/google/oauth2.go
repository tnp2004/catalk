package google

import (
	"catalk/config"
	"sync"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	once                 sync.Once
	googleConfigInstance oauth2.Config
)

func GoogleConfig() oauth2.Config {
	once.Do(func() {
		config := config.GetConfig().Google
		googleConfigInstance = oauth2.Config{
			ClientID:     config.OAuth.ClientID,
			ClientSecret: config.OAuth.ClientSecret,
			Endpoint:     google.Endpoint,
			RedirectURL:  config.OAuth.RedirectURL,
			Scopes:       config.OAuth.Scopes,
		}
	})

	return googleConfigInstance
}
