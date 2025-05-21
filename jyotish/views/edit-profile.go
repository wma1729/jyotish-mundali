package views

import (
	"html/template"
	"jyotish/models"
	"log"
	"net/http"
)

type EditProfilePage struct {
	MainPage
	UserProfile *models.Profile
}

func GetRashiNumber(chart models.GrahasLocation, graha string) int {
	for _, p := range chart.Grahas {
		if p.Name == graha {
			return p.RashiNum
		}
	}
	return -1
}

func GetDegreeInRashi(chart models.GrahasLocation, graha string) float64 {
	for _, p := range chart.Grahas {
		if p.Name == graha {
			return p.Degree
		}
	}
	return -1.00
}

func GetRetrogradeStatus(chart models.GrahasLocation, graha string) string {
	for _, p := range chart.Grahas {
		if p.Name == graha {
			if p.Retrograde {
				return "1"
			} else {
				return "0"
			}
		}
	}
	return "0"
}

func GetEditProfilePage(user *models.User, profile *models.Profile) (*EditProfilePage, error) {
	var page EditProfilePage

	if user.Lang == "en" {
		page.Vocab = &models.EnglishVocab
	} else {
		page.Vocab = &models.HindiVocab
	}

	page.User = user
	page.UserProfile = profile

	log.Print(page)

	return &page, nil

}

func (page *EditProfilePage) Send(w http.ResponseWriter) error {
	tmplName := "edit-profile"
	tmpl := template.Must(template.New(tmplName).Funcs(
		template.FuncMap{
			"GetRashiNumber":      GetRashiNumber,
			"GetDegreeInRashi":    GetDegreeInRashi,
			"GetRetrogradeStatus": GetRetrogradeStatus,
		}).ParseFiles(
		"templates/edit-profile.html",
		"templates/header.html",
		"templates/navbar.html",
		"templates/footer.html"))
	return tmpl.ExecuteTemplate(w, tmplName, page)
}
