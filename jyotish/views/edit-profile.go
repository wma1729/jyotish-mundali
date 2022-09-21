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

type EditProfilePage struct {
	MainPage
	BirthDetails  string `json:"birthdetails"`
	Name          string `json:"name"`
	DateOfBirth   string `json:"dob"`
	City          string `json:"city"`
	State         string `json:"state"`
	Country       string `json:"country"`
	ChartDetails  string `json:"chartdetails"`
	Lagna         string `json:"lagna"`
	RashiNumber   string `json:"rashi-number"`
	DegreeInRashi string `json:"degree-in-rashi"`
	Retrograde    string `json:"retrograde"`
	Sun           string `json:"sun"`
	Moon          string `json:"moon"`
	Mars          string `json:"mars"`
	Jupiter       string `json:"jupiter"`
	Mercury       string `json:"mercury"`
	Venus         string `json:"venus"`
	Saturn        string `json:"saturn"`
	Rahu          string `json:"rahu"`
	Ketu          string `json:"ketu"`
	Save          string `json:"save"`
	UserProfile   *models.Profile
}

func GetRashiNumber(chart analysis.GrahaDetais, graha string) int {
	for _, p := range chart.Grahas {
		if p.Name == graha {
			return p.RashiNum
		}
	}
	return -1
}

func GetDegreeInRashi(chart analysis.GrahaDetais, graha string) float32 {
	for _, p := range chart.Grahas {
		if p.Name == graha {
			return p.Degree
		}
	}
	return -1.00
}

func GetRetrogradeStatus(chart analysis.GrahaDetais, graha string) string {
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
	fileName := fmt.Sprintf("lang/%s/main.json", user.Lang)
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("failed to open %s", fileName)
		return nil, err
	}

	var page EditProfilePage

	err = json.Unmarshal(fileContent, &page)
	if err != nil {
		log.Printf("failed to unmarshal contents of %s", fileName)
		return nil, err
	}

	fileName = fmt.Sprintf("lang/%s/edit-profile.json", user.Lang)
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
	page.UserProfile = profile

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
