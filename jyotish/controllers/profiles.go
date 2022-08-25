package controllers

import (
	"fmt"
	"jyotish/authn"
	"jyotish/db"
	"jyotish/models"
	"jyotish/views"
	"log"
	"net/http"
	"strconv"
	"time"
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
	r.ParseForm()

	for key, values := range r.Form {
		for _, value := range values {
			log.Printf("%s = %s", key, value)
		}
	}

	profile := &models.Profile{}
	profile.ID = r.FormValue("profile-id")
	profile.Name = r.FormValue("profile-name")
	profile.DateOfBirth, _ = time.Parse("2006-01-02T15:04", r.FormValue("profile-dob"))
	profile.City = r.FormValue("profile-city")
	profile.State = r.FormValue("profile-state")
	profile.Country = r.FormValue("profile-country")

	planets := []string{"lagna", "sun", "moon", "mars", "jupiter", "mercury", "jupiter", "venus", "saturn", "rahu", "ketu"}
	for _, p := range planets {
		planet := models.PlanetPosition{}
		planet.Name = p
		planet.RashiNum, _ = strconv.Atoi(r.FormValue(p + "-rashi"))
		planet.Degree = StringToFloat32((r.FormValue(p + "-degree")))
		profile.Details.Planets = append(profile.Details.Planets, planet)
	}

	var err error
	if profile.ID == "" {
		err = db.ProfileInsert(g.DB, user.Email, profile)
	} else {
		err = db.ProfileUpdate(g.DB, user.Email, profile)
	}

	if err != nil {
		httpError := views.GetHTTPError(http.StatusInternalServerError,
			err, fmt.Sprintf("failed to insert/update profile in the database for %s: %v", user.Email, profile))
		httpError.Send(w)
		return
	}

	http.Redirect(w, r, "/profiles", http.StatusTemporaryRedirect)
}

func deleteProfile(w http.ResponseWriter, r *http.Request, g *Globals, user *models.User, id string) {
	err := db.ProfileDelete(g.DB, user.Email, id)
	if err != nil {
		httpError := views.GetHTTPError(http.StatusInternalServerError,
			err, fmt.Sprintf("failed to delete profile %s in the database for %s", id, user.Email))
		httpError.Send(w)
		return
	}

	w.WriteHeader(http.StatusOK)
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
	profile, err := db.ProfileGet(g.DB, user.Email, id)
	if err != nil {
		httpError := views.GetHTTPError(http.StatusInternalServerError,
			err, "failed to get profile")
		httpError.Send(w)
		return
	}

	log.Print(profile)

	page, err := views.GetEditProfilePage(user, profile)
	if err != nil {
		httpError := views.GetHTTPError(http.StatusInternalServerError,
			err, "failed to get edit profile page")
		httpError.Send(w)
		return
	}

	page.Send(w)
}
