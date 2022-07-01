package authn

import (
	"errors"
	"jyotish/config"
	"log"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
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

func (a *Authenticator) GetAuthCodeURL(state, nonce string) string {
	return a.Config.AuthCodeURL(state, oidc.Nonce(nonce))
}

func (a *Authenticator) ExchangeCodeWithToken(code string) (*AuthenticatedUser, error) {
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

	var authUser AuthenticatedUser
	if err := idToken.Claims(&authUser); err != nil {
		log.Printf("failed to decode ID token (%s): %s", rawIdToken, err)
		return nil, err
	}

	authUser.Token = token
	authUser.IDToken = idToken
	authUser.RawIDToken = rawIdToken

	return &authUser, nil
}

func (a *Authenticator) Init(config *config.Config) {
	a.Context = context.Background()

	var err error
	a.Provider, err = oidc.NewProvider(a.Context, config.Idp.Domain)
	if err != nil {
		log.Fatalf("failed to get IdP details - %s", err)
	}

	log.Printf("Idp Auth URL - %s, Token URL - %s\n",
		a.Provider.Endpoint().AuthURL,
		a.Provider.Endpoint().TokenURL)

	a.Config = &oauth2.Config{
		ClientID:     config.Idp.ClientID,
		ClientSecret: config.Idp.ClientSecret,
		Endpoint:     a.Provider.Endpoint(),
		RedirectURL:  config.Idp.RedirectURL,
		Scopes:       []string{oidc.ScopeOpenID, oidc.ScopeOfflineAccess, "profile", "email"},
	}

	a.Verifier = a.Provider.Verifier(&oidc.Config{ClientID: config.Idp.ClientID})
}
