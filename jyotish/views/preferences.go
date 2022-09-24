package views

import (
	"html/template"
	"jyotish/models"
	"net/http"
)

type PreferencesPage struct {
	MainPage
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
	var page PreferencesPage

	if user.Lang == "en" {
		page.Vocab = &models.EnglishVocab
	} else {
		page.Vocab = &models.HindiVocab
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
