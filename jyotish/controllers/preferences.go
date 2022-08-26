package controllers

import (
	"jyotish/authn"
	"jyotish/db"
	"jyotish/models"
	"jyotish/views"
	"log"
	"net/http"
)

func (g *Globals) HandlePreferences(w http.ResponseWriter, r *http.Request) {
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
		getPreferences(w, r, g, user)

	case "POST":
		setPreferences(w, r, g, user)
	}
}

func getPreferences(w http.ResponseWriter, r *http.Request, g *Globals, user *models.User) {
	page, err := views.GetPreferencesPage(user)
	if err != nil {
		httpError := views.GetHTTPError(http.StatusInternalServerError,
			err, "failed to get preferences page")
		httpError.Send(w)
		return
	}

	page.Send(w)
}

func setPreferences(w http.ResponseWriter, r *http.Request, g *Globals, user *models.User) {
	r.ParseForm()

	log.Print("Form values:")
	for key, values := range r.Form {
		for _, value := range values {
			log.Printf("  %s = %s", key, value)
		}
	}

	user.Name = r.FormValue("name")
	user.Description = r.FormValue("description")
	user.Lang = r.FormValue("lang")

	if r.FormValue("astrologer") == "1" {
		user.Astrologer = true
	} else {
		user.Astrologer = false
	}

	if r.FormValue("public") == "1" {
		user.Public = true
	} else {
		user.Public = false
	}

	err := db.UserUpdate(g.DB, user)
	if err != nil {
		httpError := views.GetHTTPError(http.StatusInternalServerError,
			err, "failed to update the user in the database")
		httpError.Send(w)
		return
	}

	page, err := views.GetPreferencesPage(user)
	if err != nil {
		httpError := views.GetHTTPError(http.StatusInternalServerError,
			err, "failed to get preferences page")
		httpError.Send(w)
		return
	}

	page.Send(w)
}
