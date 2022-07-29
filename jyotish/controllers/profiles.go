package controllers

import (
	"fmt"
	"jyotish/authn"
	"jyotish/db"
	"jyotish/models"
	"jyotish/views"
	"log"
	"net/http"
	"strings"
)

func (g *Globals) HandleProfiles(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request URL - %s\n", r.URL)

	authUser, err := authn.GetUserSession(r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	log.Println("already authenticated user")

	user, err := db.UserGet(g.DB, authUser.User.Email)
	if err != nil {
		httpError := views.GetHTTPError(http.StatusInternalServerError,
			err, "failed to get the user from the database")
		httpError.Send(w)
		return
	}

	switch r.Method {
	case "GET":
		components := strings.Split(r.URL.Path, "/")
		numOfComponents := len(components)
		if numOfComponents == 2 {
			getAllProfiles(w, r, g, user)
		} else if numOfComponents == 3 {
			getProfile(w, r, g, user, components[2])
		} else {
			// error
		}

	case "POST":
		setProfile(w, r, g, user)

	case "DELETE":
		components := strings.Split(r.URL.Path, "/")
		numOfComponents := len(components)
		if numOfComponents == 2 {
			deleteProfile(w, r, g, user, components[1])
		} else {
			// error
		}

	}
}

func getAllProfiles(w http.ResponseWriter, r *http.Request, g *Globals, user *models.User) {
	profiles, err := db.ProfilesList(g.DB, user.Email)
	if err != nil {
		httpError := views.GetHTTPError(http.StatusInternalServerError,
			err, fmt.Sprintf("failed to get profiles from the database for %s", user.Email))
		httpError.Send(w)
		return
	}

	page, err := views.GetListProfilesPage(user, profiles)
	if err != nil {
		httpError := views.GetHTTPError(http.StatusInternalServerError,
			err, "failed to get list profiles page")
		httpError.Send(w)
		return
	}

	page.Send(w)
}

func getProfile(w http.ResponseWriter, r *http.Request, g *Globals, user *models.User, id string) {

}

func setProfile(w http.ResponseWriter, r *http.Request, g *Globals, user *models.User) {

}

func deleteProfile(w http.ResponseWriter, r *http.Request, g *Globals, user *models.User, id string) {

}
