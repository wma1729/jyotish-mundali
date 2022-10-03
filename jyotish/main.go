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

	globals, err := controllers.InitGlobals(&config)
	if err != nil {
		log.Printf("failed to initialize environment: %s", err)
		return
	}

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", globals.BeginAuth)
	http.HandleFunc("/auth/callback", globals.CompleteAuth)
	http.HandleFunc("/logout", globals.EndAuth)
	http.HandleFunc("/preferences", globals.HandlePreferences)
	http.HandleFunc("/profiles/", globals.HandleProfiles)
	http.HandleFunc("/knowledge-base", globals.HandleKnowledgeBase)

	log.Fatal(http.ListenAndServe("localhost:5000", logRequest(http.DefaultServeMux)))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request %s %s", r.Method, r.URL.Path)
		for key, values := range r.URL.Query() {
			for _, value := range values {
				log.Printf("  %s: %s", key, value)
			}
		}
		handler.ServeHTTP(w, r)
	})
}
