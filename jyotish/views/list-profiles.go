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

type ListProfilesPage struct {
	MainPage
	CreateProfile string `json:"create-profile"`
	Name          string `json:"name"`
	DateOfBirth   string `json:"dob"`
	PlaceOfBirth  string `json:"pob"`
	Action        string `json:"action"`
	UserProfiles  []models.Profile
}

func GetListProfilesPage(user *models.User, profiles []models.Profile) (*ListProfilesPage, error) {
	fileName := fmt.Sprintf("lang/%s/main.json", user.Lang)
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("failed to open %s", fileName)
		return nil, err
	}

	var page ListProfilesPage

	err = json.Unmarshal(fileContent, &page)
	if err != nil {
		log.Printf("failed to unmarshal contents of %s", fileName)
		return nil, err
	}

	fileName = fmt.Sprintf("lang/%s/list-profiles.json", user.Lang)
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

	page.User = user
	page.UserProfiles = profiles

	return &page, nil

}

func (page *ListProfilesPage) Send(w http.ResponseWriter) error {
	tmplName := "list-profiles"
	tmpl := template.Must(template.New(tmplName).ParseFiles(
		"templates/list-profiles.html",
		"templates/main.html",
		"templates/header.html",
		"templates/navbar.html",
		"templates/footer.html"))
	return tmpl.ExecuteTemplate(w, tmplName, page)
}
