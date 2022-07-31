package controllers

import (
	"fmt"
	"jyotish/authn"
	"jyotish/db"
	"jyotish/models"
	"jyotish/views"
	"log"
	"net/http"
)

/*
 * GET    /profiles           - Get all profiles.
 * POST   /profiles           - Create/edit a specific profile.
 * GET    /profiles/{id}      - Get a specific profile.
 * DELETE /profiles/{id}      - Delete a specific profile.
 * GET    /profiles/edit      - Get the page to create a new profile.
 * GET    /profiles/edit/{id} - Get the page to edit a specific profile.
 */
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

	pathSegments := SplitPath(r.URL.Path)
	numOfSegments := len(pathSegments)

	log.Print(pathSegments)

	switch r.Method {
	case "GET":
		if numOfSegments == 1 {
			/* GET /profiles */
			getAllProfiles(w, r, g, user)
		} else if pathSegments[1] == "edit" {
			if numOfSegments == 2 {
				/* GET /profiles/edit */
				getCreateProfilePage(w, r, g, user)
			} else {
				/* GET /profiles/edit/{id} */
				getEditProfilePage(w, r, g, user, pathSegments[2])
			}
		} else {
			/* GET /profiles/{id} */
			getProfile(w, r, g, user, pathSegments[1])
		}

	case "POST":
		setProfile(w, r, g, user)

	case "DELETE":
		deleteProfile(w, r, g, user, pathSegments[1])
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

func getCreateProfilePage(w http.ResponseWriter, r *http.Request, g *Globals, user *models.User) {
	page, err := views.GetEditProfilePage(user, nil)
	if err != nil {
		httpError := views.GetHTTPError(http.StatusInternalServerError,
			err, "failed to get create profile page")
		httpError.Send(w)
		return
	}

	page.Send(w)
}

func getEditProfilePage(w http.ResponseWriter, r *http.Request, g *Globals, user *models.User, id string) {

}
