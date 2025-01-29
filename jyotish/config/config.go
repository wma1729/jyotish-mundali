package config

import (
	"log"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

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
	} `yaml:"database"`
}

const (
	IDP_DOMAIN        = "IDP_DOMAIN"
	IDP_CLIENT_ID     = "IDP_CLIENT_ID"
	IDP_CLIENT_SECRET = "IDP_CLIENT_SECRET"
	IDP_REDIRECT_URL  = "IDP_REDIRECT_URL"
	DB_HOST           = "DB_HOST"
	DB_PORT           = "DB_PORT"
	DB_USER           = "DB_USER"
	DB_PASSWORD       = "DB_PASSWORD"
)

func (config *Config) LoadFromEnvironment() error {
	var err error
	config.Idp.Domain = os.Getenv(IDP_DOMAIN)
	config.Idp.ClientID = os.Getenv(IDP_CLIENT_ID)
	config.Idp.ClientSecret = os.Getenv(IDP_CLIENT_SECRET)
	config.Idp.RedirectURL = os.Getenv(IDP_REDIRECT_URL)
	config.Database.Host = os.Getenv(DB_HOST)
	config.Database.Port, err = strconv.Atoi(os.Getenv(DB_PORT))
	if err != nil {
		log.Println(err)
		return err
	}
	config.Database.User = os.Getenv(DB_USER)
	config.Database.Password = os.Getenv(DB_PASSWORD)
	return nil
}

func (config *Config) LoadFromYaml(configFile string) error {
	f, err := os.Open(configFile)
	if err != nil {
		log.Println(err)
		return err
	}

	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(config)
	if err != nil {
		log.Printf("failed to unmarshal %s: %s", configFile, err)
		return err
	}

	return nil
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
