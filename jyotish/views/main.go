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

type MainPage struct {
	Home          string `json:"home"`
	Profiles      string `json:"profiles"`
	KnowledgeBase string `json:"knowledgebase"`
	FAQs          string `json:"faqs"`
	SiteAdmin     string `json:"siteadmin"`
	Preferences   string `json:"preferences"`
	Logout        string `json:"logout"`
	Contact       string `json:"contact"`
	User          *models.User
}

func GetMainPage(user *models.User) (*MainPage, error) {
	fileName := fmt.Sprintf("lang/%s/main.json", user.Lang)
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("failed to open %s", fileName)
		return nil, err
	}

	var page MainPage

	err = json.Unmarshal(fileContent, &page)
	if err != nil {
		log.Printf("failed to unmarshal contents of %s", fileName)
		return nil, err
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
