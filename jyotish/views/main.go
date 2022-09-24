package views

import (
	"html/template"
	"jyotish/models"
	"net/http"
)

type MainPage struct {
	Vocab *models.Language
	User  *models.User
}

func GetMainPage(user *models.User) (*MainPage, error) {
	var page MainPage

	if user.Lang == "en" {
		page.Vocab = &models.EnglishVocab
	} else {
		page.Vocab = &models.HindiVocab
	}

	page.User = user

	return &page, nil

}

func (page *MainPage) Send(w http.ResponseWriter) error {
	tmplName := "main"
	tmpl := template.Must(template.New(tmplName).ParseFiles(
		"templates/main.html",
		"templates/header.html",
		"templates/navbar.html",
		"templates/footer.html"))
	return tmpl.ExecuteTemplate(w, tmplName, page)
}
