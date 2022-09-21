package views

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"jyotish/analysis"
	"jyotish/models"
	"log"
	"net/http"
)

type AnalysisPage struct {
	MainPage
	Name      string `json:"name"`
	Natural   string `json:"natural"`
	Temporary string `json:"temporary"`
	Effective string `json:"effective"`
	Friends   string `json:"friends"`
	Neutrals  string `json:"neutrals"`
	Enemies   string `json:"enemies"`
	Nature    string `json:"nature"`
	Chart     analysis.Chart
}

func GetAnalysisPage(user *models.User, chart analysis.Chart) (*AnalysisPage, error) {
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

	fileName = fmt.Sprintf("lang/%s/analysis.json", user.Lang)
	fileContent, err = ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("failed to open %s", fileName)
		return nil, err
	}

	err = json.Unmarshal(fileContent, &page)
	if err != nil {
		log.Printf("failed to unmarshal contents of %s", fileName)
		return nil, err
	}

	log.Print(page)

	page.User = user
	page.Chart = chart

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
