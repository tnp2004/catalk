package config

import (
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	Server   *Server
	Database *Database
	Google   *Google
	Web      *Web
}

type Web struct {
	Port     string
	HostName string
}

type Server struct {
	Port     int
	HostName string
}

type Database struct {
	Host     string
	Port     string
	DbName   string
	Username string
	Password string
	Schema   string
}

type Google struct {
	ApiKey       string
	OAuth        GoogleOAuth2
	GoogleApiUrl string
}

type GoogleOAuth2 struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Scopes       []string
}

var (
	once           sync.Once
	configInstance *Config
)

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./config")
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				panic(err)
			} else {
				panic(err)
			}
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
	})

	return configInstance
}
