package controllers

import (
	"fmt"
	"jyotish/authn"
	"jyotish/views"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/oauth2"
)

func (g *Globals) BeginAuth(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request URL - %s\n", r.URL)

	authUser, err := authn.GetUserSession(r)
	if err == nil {
		log.Println("already authenticated user")
		sendMainPage(w, "en", authUser.Name)
		return
	}

	log.Println("initiate the authentication")

	state, err := authn.SetStateInCookie(w, r)
	if err != nil {
		httpError := views.GetHTTPError(http.StatusInternalServerError,
			err, "failed to set the state in the cookie")
		httpError.Send(w)
		return
	}

	nonce, err := authn.SetNonceInCookie(w, r)
	if err != nil {
		httpError := views.GetHTTPError(http.StatusInternalServerError,
			err, "failed to set the nonce in the cookie")
		httpError.Send(w)
		return
	}

	url := g.Authenticator.GetAuthCodeURL(state, nonce)
	log.Printf("Redirect for auth code to - %s", url)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (g *Globals) CompleteAuth(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request URL - %s\n", r.URL)

	if err := authn.ValidateState(r); err != nil {
		httpError := views.GetHTTPError(http.StatusInternalServerError,
			err, "failed to validate state")
		httpError.Send(w)
		return
	}

	code := r.URL.Query().Get("code")

	authUser, err := g.Authenticator.ExchangeCodeWithToken(code)
	if err != nil {
		httpError := views.GetHTTPError(http.StatusInternalServerError,
			err, fmt.Sprintf("failed to exchange the code (%s) with token",
				code))
		httpError.Send(w)
		return
	}

	if err := authn.ValidateNonce(r, authUser.IDToken); err != nil {
		httpError := views.GetHTTPError(http.StatusInternalServerError,
			err, "failed to validate nonce")
		httpError.Send(w)
		return
	}

	authUser.User, err = g.Authenticator.Provider.UserInfo(g.Authenticator.Context, oauth2.StaticTokenSource(authUser.Token))
	if err != nil {
		httpError := views.GetHTTPError(http.StatusInternalServerError,
			err, "failed to get user information")
		httpError.Send(w)
		return
	}

	err = authn.SetUserSession(w, r, authUser)
	if err != nil {
		httpError := views.GetHTTPError(http.StatusInternalServerError,
			err, "failed to set user session")
		httpError.Send(w)
		return
	}

	sendMainPage(w, "en", authUser.Name)

	return
}

func (g *Globals) EndAuth(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request URL - %s\n", r.URL)
	user, err := authn.GetUserSession(r)
	if err != nil {
		httpError := views.GetHTTPError(http.StatusUnauthorized,
			err, "unable to find user session")
		httpError.Send(w)
		return
	}

	authn.ResetUserSession(w, r, user)

	var claims struct {
		LogoutURL string `json:"end_session_endpoint"`
	}

	if err := g.Authenticator.Provider.Claims(&claims); err != nil {
		httpError := views.GetHTTPError(http.StatusInternalServerError,
			err, "failed to get logout URL")
		httpError.Send(w)
		return
	}

	logoutURL, err := url.Parse(claims.LogoutURL)
	if err != nil {
		httpError := views.GetHTTPError(http.StatusInternalServerError,
			err, "failed to parse logout URL")
		httpError.Send(w)
		return
	}

	parameters := url.Values{}
	parameters.Add("id_token_hint", user.RawIDToken)
	parameters.Add("client_id", g.Authenticator.Config.ClientID)

	logoutURL.RawQuery = parameters.Encode()

	http.Redirect(w, r, logoutURL.String(), http.StatusTemporaryRedirect)
}

func sendMainPage(w http.ResponseWriter, lang, userName string) {
	page, err := views.GetMainPage(lang, userName)
	if err != nil {
		httpError := views.GetHTTPError(http.StatusInternalServerError,
			err, "failed to load the main page details")
		httpError.Send(w)
		return
	}

	err = page.Send(w)
	if err != nil {
		httpError := views.GetHTTPError(http.StatusInternalServerError,
			err, "failed to send the main page")
		httpError.Send(w)
	}

	return
}
