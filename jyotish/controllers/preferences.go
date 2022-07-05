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
		getPreferences(w, r, g, user)

	case "POST":
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

	log.Print(page)

	page.Send(w)
}
