package views

import (
	"fmt"
	"html/template"
	"jyotish/analysis"
	"jyotish/models"
	"net/http"
	"strings"
)

type AnalysisPage struct {
	MainPage
	Chart     analysis.Chart
	BhavaDesc []models.BhavaDescription
}

func GetGrahaName(graha string, lang string) string {
	return models.GrahaName(graha, lang)
}

func ListOfGrahas(grahas []string, lang string) string {
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

func ListOfIntegers(numbers []int) string {
	sb := strings.Builder{}
	first := true
	for _, n := range numbers {
		if !first {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%d", n))
		first = false
	}
	return sb.String()
}

func GetAnalysisPage(user *models.User, chart analysis.Chart) (*AnalysisPage, error) {
	var page AnalysisPage

	page.Vocab = &models.EnglishVocab
	if user.Lang != "en" {
		page.Vocab = &models.HindiVocab
	}

	page.User = user
	page.Chart = chart
	page.BhavaDesc, _ = models.GetBhavaDescription(user.Lang)

	return &page, nil

}

func (page *AnalysisPage) Send(w http.ResponseWriter) error {
	tmplName := "analysis"
	tmpl := template.Must(template.New(tmplName).Funcs(
		template.FuncMap{
			"GetBhava":             GetBhava,
			"GetGrahaNameForChart": GetGrahaNameForChart,
			"GetGrahaName":         GetGrahaName,
			"ListOfGrahas":         ListOfGrahas,
			"ListOfIntegers":       ListOfIntegers,
			"GrahaNature":          models.GrahaNature,
			"GrahaMotion":          models.GrahaMotion,
			"YesOrNo":              models.YesOrNo,
			"GrahaPosition":        models.GrahaPosition,
			"GrahaState":           models.GrahaState,
			"GetRashiName":         GetRashiName,
			"GetInfluenceRating":   GetInfluenceRating,
			"GetInfluenceOnBhava":  GetInfluenceOnBhava,
			"GetMaanglikDosha":     GetMaanglikDosha,
			"GetSeverity":          GetSeverity,
		}).ParseFiles(
		"templates/analysis.html",
		"templates/lagna-chart.html",
		"templates/grahas-relations.html",
		"templates/grahas-aspects.html",
		"templates/grahas-nature.html",
		"templates/grahas-strength.html",
		"templates/doshas.html",
		"templates/bhavas.html",
		"templates/header.html",
		"templates/navbar.html",
		"templates/footer.html"))
	return tmpl.ExecuteTemplate(w, tmplName, page)
}
