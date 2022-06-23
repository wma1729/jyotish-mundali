package main

import (
	"crypto/tls"
	"jyotish/authn"
	"log"
	"net/http"
)

const IdpDomain string = "https://localhost:8443/realms/JyotishMundali"

func main() {
	defaultTransport, ok := http.DefaultTransport.(*http.Transport)
	if !ok {
		log.Print("failed to get default transport")
		return
	}

	defaultTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	var a authn.Authenticator
	a.Init()

	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/", a.BeginAuth)
	http.HandleFunc("/auth/callback", a.CompleteAuth)
	http.HandleFunc("/logout", a.EndAuth)

	log.Fatal(http.ListenAndServe("localhost:5000", nil))
}
