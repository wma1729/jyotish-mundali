package main

import (
	"crypto/tls"
	"flag"
	"jyotish/config"
	"jyotish/controllers"
	"log"
	"net/http"
)

func main() {
	var config config.Config

	config.LoadFromEnvironment()
	config.LoadFromYaml("config.yaml")

	var domain, clientID, clientSecret, redirectURL string

	flag.StringVar(&domain, "domain", "", "Identity provider domain name.")
	flag.StringVar(&clientID, "client_id", "", "Identity provider client ID.")
	flag.StringVar(&clientSecret, "client_secret", "", "Identity provider client secret.")
	flag.StringVar(&redirectURL, "redirect_url", "", "Identity provider redirect URL.")

	var host, user, password string
	var port int

	flag.StringVar(&host, "host", "", "Database server name.")
	flag.IntVar(&port, "port", -1, "Database port number.")
	flag.StringVar(&user, "user", "", "Database user name.")
	flag.StringVar(&password, "password", "", "Database user password.")

	flag.Parse()

	if domain != "" {
		config.Idp.Domain = domain
	}

	if clientID != "" {
		config.Idp.ClientID = clientID
	}

	if clientSecret != "" {
		config.Idp.ClientSecret = clientSecret
	}

	if redirectURL != "" {
		config.Idp.RedirectURL = redirectURL
	}

	if host != "" {
		config.Database.Host = host
	}

	if port != -1 {
		config.Database.Port = port
	}

	if user != "" {
		config.Database.User = user
	}

	if password != "" {
		config.Database.Password = password
	}

	config.Validate()

	defaultTransport, ok := http.DefaultTransport.(*http.Transport)
	if !ok {
		log.Print("failed to get default transport")
		return
	}

	defaultTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	env, err := controllers.InitGlobals(&config)
	if err != nil {
		log.Printf("failed to initialize environment: %s", err)
		return
	}

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", env.BeginAuth)
	http.HandleFunc("/auth/callback", env.CompleteAuth)
	http.HandleFunc("/logout", env.EndAuth)

	log.Fatal(http.ListenAndServe("localhost:5000", nil))
}
