package views

import (
	"html/template"
	"jyotish/docs"
	"jyotish/models"
	"net/http"
)

type KnowledgeBasePage struct {
	MainPage
	Contents []docs.Content
}

func GetKnowledgeBasePage(user *models.User, contents []docs.Content) (*KnowledgeBasePage, error) {
	var page KnowledgeBasePage

	if user.Lang == "en" {
		page.Vocab = &models.EnglishVocab
	} else {
		page.Vocab = &models.HindiVocab
	}

	page.User = user
	page.Contents = contents

	return &page, nil

}

func (page *KnowledgeBasePage) Send(w http.ResponseWriter) error {
	tmplName := "knowledge-base"
	tmpl := template.Must(template.New(tmplName).ParseFiles(
		"templates/knowledge-base.html",
		"templates/header.html",
		"templates/navbar.html",
		"templates/footer.html"))
	return tmpl.ExecuteTemplate(w, tmplName, page)
}
