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

type PreferencesPage struct {
	MainPage
	Name                string `json:"name"`
	Description         string `json:"description"`
	Language            string `json:"language"`
	English             string `json:"english"`
	Hindi               string `json:"hindi"`
	Astrologer          string `json:"astrologer"`
	Public              string `json:"public"`
	Yes                 string `json:"yes"`
	No                  string `json:"no"`
	Save                string `json:"save"`
	CurrentDescription  string
	CurrentLanguage     string
	CurrentlyAstrologer bool
	CurrentlyPublic     bool
}

func IsStringChecked(s1, s2 string) string {
	if s1 == s2 {
		return "checked"
	}
	return ""
}

func IsBooleanChecked(b1, b2 bool) string {
	if b1 == b2 {
		return "checked"
	}
	return ""
}

func GetPreferencesPage(user *models.User) (*PreferencesPage, error) {
	fileName := fmt.Sprintf("lang/%s/main.json", user.Lang)
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("failed to open %s", fileName)
		return nil, err
	}

	var page PreferencesPage

	err = json.Unmarshal(fileContent, &page)
	if err != nil {
		log.Printf("failed to unmarshal contents of %s", fileName)
		return nil, err
	}

	fileName = fmt.Sprintf("lang/%s/preferences.json", user.Lang)
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

	return &page, nil
}

func (page *PreferencesPage) Send(w http.ResponseWriter) error {
	tmplName := "preferences"
	tmpl := template.Must(template.New(tmplName).Funcs(
		template.FuncMap{
			"IsStringChecked":  IsStringChecked,
			"IsBooleanChecked": IsBooleanChecked,
		}).ParseFiles(
		"templates/preferences.html",
		"templates/header.html",
		"templates/navbar.html",
		"templates/footer.html"))
	return tmpl.ExecuteTemplate(w, tmplName, page)
}
