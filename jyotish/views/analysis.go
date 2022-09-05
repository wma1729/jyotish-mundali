package views

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"jyotish/models"
	"log"
	"net/http"
)

type AnalysisPage struct {
	MainPage
	Bhavas []models.Bhava
}

func GetAnalysisPage(user *models.User, bhavas []models.Bhava) (*AnalysisPage, error) {
	fileName := fmt.Sprintf("lang/%s/main.json", user.Lang)
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("failed to open %s", fileName)
		return nil, err
	}

	var page AnalysisPage

	err = json.Unmarshal(fileContent, &page)
	if err != nil {
		log.Printf("failed to unmarshal contents of %s", fileName)
		return nil, err
	}

	page.User = user
	page.Bhavas = bhavas

	return &page, nil

}

func (page *AnalysisPage) Send(w http.ResponseWriter) error {
	tmplName := "analysis"
	tmpl := template.Must(template.New(tmplName).ParseFiles(
		"templates/analysis.html",
		"templates/header.html",
		"templates/navbar.html",
		"templates/footer.html"))
	return tmpl.ExecuteTemplate(w, tmplName, page)
}
