package authn

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"jyotish/models"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

const (
	IDP_DOMAIN        = "IDP_DOMAIN"
	IDP_CLIENT_ID     = "IDP_CLIENT_ID"
	IDP_CLIENT_SECRET = "IDP_CLIENT_SECRET"
	IDP_REDIRECT_URL  = "IDP_REDIRECT_URL"
)

//
// The main oauth2/oidc authenticator.
//
type Authenticator struct {
	Context  context.Context
	Provider *oidc.Provider
	Verifier *oidc.IDTokenVerifier
	Config   *oauth2.Config
}

func (a *Authenticator) getAuthCodeURL(state, nonce string) string {
	return a.Config.AuthCodeURL(state, oidc.Nonce(nonce))
}

func (a *Authenticator) exchangeCodeWithToken(code string) (*User, error) {
	token, err := a.Config.Exchange(a.Context, code)
	if err != nil {
		log.Printf("failed to exchange code (%s) with token: %s", code, err)
		return nil, err
	}

	rawIdToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no ID token found in oauth2 token")
	}

	idToken, err := a.Verifier.Verify(a.Context, rawIdToken)
	if err != nil {
		log.Printf("failed to verify the ID token: %s", err)
		return nil, err
	}

	var authUser User
	if err := idToken.Claims(&authUser); err != nil {
		log.Printf("failed to decode ID token (%s): %s", rawIdToken, err)
		return nil, err
	}

	authUser.Token = token
	authUser.IDToken = idToken
	authUser.RawIDToken = rawIdToken

	return &authUser, nil
}

func (a *Authenticator) Init() {
	domain := os.Getenv(IDP_DOMAIN)
	if domain == "" {
		log.Fatalf("Environment variable %s must be set", IDP_DOMAIN)
	}

	clientID := os.Getenv(IDP_CLIENT_ID)
	if clientID == "" {
		log.Fatalf("Environment variable %s must be set", IDP_CLIENT_ID)
	}

	clientSecret := os.Getenv(IDP_CLIENT_SECRET)
	if clientSecret == "" {
		log.Fatalf("Environment variable %s must be set", IDP_CLIENT_SECRET)
	}

	redirectURL := os.Getenv(IDP_REDIRECT_URL)
	if redirectURL == "" {
		log.Fatalf("Environment variable %s must be set", IDP_REDIRECT_URL)
	}

	log.Printf("IdP domain - %s, client ID - %s, redirect URL - %s\n",
		domain, clientID, redirectURL)

	a.Context = context.Background()

	var err error
	a.Provider, err = oidc.NewProvider(a.Context, domain)
	if err != nil {
		log.Fatalf("failed to get IdP details - %s", err)
	}

	log.Printf("Idp Auth URL - %s, Token URL - %s\n",
		a.Provider.Endpoint().AuthURL,
		a.Provider.Endpoint().TokenURL)

	a.Config = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     a.Provider.Endpoint(),
		RedirectURL:  redirectURL,
		Scopes:       []string{oidc.ScopeOpenID, oidc.ScopeOfflineAccess, "profile", "email"},
	}

	a.Verifier = a.Provider.Verifier(&oidc.Config{ClientID: clientID})
}

func (a *Authenticator) BeginAuth(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request URL - %s\n", r.URL)

	authUser, err := GetUserSession(r)
	if err == nil {
		log.Println("already authenticated user")
		sendMainPage(w, "en", authUser.Name)
		return
	}

	log.Println("initiate the authentication")

	state, err := SetStateInCookie(w, r)
	if err != nil {
		models.HTTPError{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
			Detail:     "failed to set the state in the cookie",
		}.Send(w)
		return
	}

	nonce, err := SetNonceInCookie(w, r)
	if err != nil {
		models.HTTPError{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
			Detail:     "failed to set the nonce in the cookie",
		}.Send(w)
		return
	}

	url := a.getAuthCodeURL(state, nonce)
	log.Printf("Redirect for auth code to - %s", url)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (a *Authenticator) CompleteAuth(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request URL - %s\n", r.URL)

	if err := ValidateState(r); err != nil {
		models.HTTPError{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
			Detail:     "failed to validate state",
		}.Send(w)
		return
	}

	code := r.URL.Query().Get("code")

	authUser, err := a.exchangeCodeWithToken(code)
	if err != nil {
		models.HTTPError{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
			Detail:     "failed to exchange code with token",
		}.Send(w)
		return
	}

	if err := ValidateNonce(r, authUser.IDToken); err != nil {
		models.HTTPError{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
			Detail:     "failed to validate nonce",
		}.Send(w)
		return
	}

	authUser.User, err = a.Provider.UserInfo(a.Context, oauth2.StaticTokenSource(authUser.Token))
	if err != nil {
		models.HTTPError{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
			Detail:     "failed to get user info",
		}.Send(w)
		return
	}

	err = SetUserSession(w, r, authUser)
	if err != nil {
		models.HTTPError{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
			Detail:     "failed to set user session",
		}.Send(w)
		return
	}

	sendMainPage(w, "en", authUser.Name)

	return
}

func (a *Authenticator) EndAuth(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request URL - %s\n", r.URL)
	user, err := GetUserSession(r)
	if err != nil {
		models.HTTPError{
			StatusCode: http.StatusUnauthorized,
			Error:      err.Error(),
			Detail:     "failed to find user session",
		}.Send(w)
		return
	}

	ResetUserSession(w, r, user)

	var claims struct {
		LogoutURL string `json:"end_session_endpoint"`
	}

	if err := a.Provider.Claims(&claims); err != nil {
		models.HTTPError{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
			Detail:     fmt.Sprintf("failed to get logout URL"),
		}.Send(w)
		return
	}

	logoutURL, err := url.Parse(claims.LogoutURL)
	if err != nil {
		models.HTTPError{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
			Detail:     fmt.Sprintf("failed to parse logout URL"),
		}.Send(w)
		return
	}

	parameters := url.Values{}
	parameters.Add("id_token_hint", user.RawIDToken)
	parameters.Add("client_id", a.Config.ClientID)

	logoutURL.RawQuery = parameters.Encode()

	http.Redirect(w, r, logoutURL.String(), http.StatusTemporaryRedirect)
}

func sendMainPage(w http.ResponseWriter, lang, userName string) {
	page := &models.MainPage{}

	err := page.Load(lang, userName)
	if err != nil {
		models.HTTPError{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
			Detail:     "failed to load the page details",
		}.Send(w)
		return
	}

	err = page.Send(w)
	if err != nil {
		models.HTTPError{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
			Detail:     "failed to send the page details",
		}.Send(w)
	}

	return
}
