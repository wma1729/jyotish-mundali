package controllers

import (
	"jyotish/authn"
	"jyotish/db"
	"jyotish/docs"
	"jyotish/views"
	"log"
	"net/http"
)

func (g *Globals) HandleKnowledgeBase(w http.ResponseWriter, r *http.Request) {
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

	contents, err := docs.PrepareDocumentation(user.Lang)
	if err != nil {
		httpError := views.GetHTTPError(http.StatusInternalServerError,
			err, "failed to prepare knowledge-base")
		httpError.Send(w)
		return
	}

	page, err := views.GetKnowledgeBasePage(user, contents)
	if err != nil {
		httpError := views.GetHTTPError(http.StatusInternalServerError,
			err, "failed to get knowledge-base page")
		httpError.Send(w)
		return
	}

	page.Send(w)
}
