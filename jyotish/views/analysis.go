package views

import (
	"html/template"
	"jyotish/analysis"
	"jyotish/models"
	"net/http"
	"strings"
)

type AnalysisPage struct {
	MainPage
	Chart analysis.Chart
}

func GrahasName(grahas []string, lang string) string {
	sb := strings.Builder{}
	first := true
	for _, g := range grahas {
		if !first {
			sb.WriteString(", ")
		}
		sb.WriteString(models.GrahaName(g, lang))
		first = false
	}
	return sb.String()
}

func GetAnalysisPage(user *models.User, chart analysis.Chart) (*AnalysisPage, error) {
	var page AnalysisPage

	if user.Lang == "en" {
		page.Vocab = &models.EnglishVocab
	} else {
		page.Vocab = &models.HindiVocab
	}

	page.User = user
	page.Chart = chart

	return &page, nil

}

func (page *AnalysisPage) Send(w http.ResponseWriter) error {
	tmplName := "analysis"
	tmpl := template.Must(template.New(tmplName).Funcs(
		template.FuncMap{
			"GrahaName":   models.GrahaName,
			"GrahasName":  GrahasName,
			"GrahaNature": models.GrahaNature,
			"GrahaMotion": models.GrahaMotion,
			"YesOrNo":     models.YesOrNo,
		}).ParseFiles(
		"templates/analysis.html",
		"templates/header.html",
		"templates/navbar.html",
		"templates/footer.html"))
	return tmpl.ExecuteTemplate(w, tmplName, page)
}
