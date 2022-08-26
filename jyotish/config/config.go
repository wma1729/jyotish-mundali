package config

import "log"

type Config struct {
	Idp struct {
		Domain       string `yaml:"domain"`
		ClientID     string `yaml:"client_id"`
		ClientSecret string `yaml:"client_secret"`
		RedirectURL  string `yaml:"redirect_url"`
	} `yaml:"idp"`
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	}
}

func (config *Config) Validate() {
	if config.Idp.Domain == "" {
		log.Fatalf("config.Idp.Domain is not set")
	} else {
		log.Printf("config.Idp.Domain = %s", config.Idp.Domain)
	}

	if config.Idp.ClientID == "" {
		log.Fatalf("config.Idp.ClientID is not set")
	} else {
		log.Printf("config.Idp.ClientID = %s", config.Idp.ClientID)
	}

	if config.Idp.ClientSecret == "" {
		log.Fatalf("config.Idp.ClientSecret is not set")
	} else {
		log.Printf("config.Idp.ClientSecret = %s", "********")
	}

	if config.Idp.RedirectURL == "" {
		log.Fatalf("config.Idp.RedirectURL is not set")
	} else {
		log.Printf("config.Idp.RedirectURL = %s", config.Idp.RedirectURL)
	}

	if config.Database.Host == "" {
		log.Fatalf("config.Database.Host is not set")
	} else {
		log.Printf("config.Database.Host = %s", config.Database.Host)
	}

	if config.Database.Port < 1024 || config.Database.Port > 60000 {
		log.Fatalf("config.Database.Port is not set")
	} else {
		log.Printf("config.Database.Port = %d", config.Database.Port)
	}

	if config.Database.User == "" {
		log.Fatalf("config.Database.User is not set")
	} else {
		log.Printf("config.Database.User = %s", config.Database.User)
	}

	if config.Database.Password == "" {
		log.Fatalf("config.Database.Password is not set")
	} else {
		log.Printf("config.Database.Password = %s", "********")
	}
}
