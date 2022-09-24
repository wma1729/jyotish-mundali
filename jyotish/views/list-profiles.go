package views

import (
	"html/template"
	"jyotish/models"
	"net/http"
)

type ListProfilesPage struct {
	MainPage
	UserProfiles []models.Profile
}

func GetListProfilesPage(user *models.User, profiles []models.Profile) (*ListProfilesPage, error) {
	var page ListProfilesPage

	if user.Lang == "en" {
		page.Vocab = &models.EnglishVocab
	} else {
		page.Vocab = &models.HindiVocab
	}

	page.User = user
	page.UserProfiles = profiles

	return &page, nil

}

func (page *ListProfilesPage) Send(w http.ResponseWriter) error {
	tmplName := "list-profiles"
	tmpl := template.Must(template.New(tmplName).ParseFiles(
		"templates/list-profiles.html",
		"templates/header.html",
		"templates/navbar.html",
		"templates/footer.html"))
	return tmpl.ExecuteTemplate(w, tmplName, page)
}
