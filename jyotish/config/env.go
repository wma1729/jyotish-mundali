package config

import (
	"os"
	"strconv"
)

const (
	IDP_DOMAIN        = "IDP_DOMAIN"
	IDP_CLIENT_ID     = "IDP_CLIENT_ID"
	IDP_CLIENT_SECRET = "IDP_CLIENT_SECRET"
	IDP_REDIRECT_URL  = "IDP_REDIRECT_URL"
	DB_USER           = "DB_USER"
	DB_PASSWORD       = "DB_PASSWORD"
	DB_HOST           = "DB_HOST"
	DB_PORT           = "DB_PORT"
)

func (config *Config) LoadFromEnvironment() {
	config.Idp.Domain = os.Getenv(IDP_DOMAIN)
	config.Idp.ClientID = os.Getenv(IDP_CLIENT_ID)
	config.Idp.ClientSecret = os.Getenv(IDP_CLIENT_SECRET)
	config.Idp.RedirectURL = os.Getenv(IDP_REDIRECT_URL)
	config.Database.Host = os.Getenv(DB_HOST)
	config.Database.Port, _ = strconv.Atoi(os.Getenv(DB_PORT))
	config.Database.User = os.Getenv(DB_USER)
	config.Database.Password = os.Getenv(DB_PASSWORD)
}
