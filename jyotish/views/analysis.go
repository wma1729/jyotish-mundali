package views

import (
	"fmt"
	"html/template"
	"jyotish/charts"
	"jyotish/models"
	"net/http"
	"strings"
)

type AnalysisPage struct {
	MainPage
	Chart charts.Chart
}

func GrahaName(graha charts.GrahaLocCombust, lang string) template.HTML {
	vocab := &models.EnglishVocab
	if lang != "en" {
		vocab = &models.HindiVocab
	}

	sb := strings.Builder{}
	sb.WriteString(models.GrahaName(graha.Name, lang))
	sb.WriteString("<sup>")
	sb.WriteString(fmt.Sprintf("%.f", graha.Degree))
	sb.WriteString("</sup>")

	closeParenthesis, addComma := false, false
	if graha.Combust || graha.Retrograde {
		sb.WriteByte('(')
		closeParenthesis = true
	}
	if graha.Combust {
		sb.WriteString(vocab.CombustAbbr)
		addComma = true
	}
	if graha.Retrograde {
		if addComma {
			sb.WriteByte(',')
		}
		sb.WriteString(vocab.RetrogradeAbbr)
	}
	if closeParenthesis {
		sb.WriteByte(')')
	}

	return template.HTML(sb.String())
}

func GetAnalysisPage(user *models.User, chart charts.Chart) (*AnalysisPage, error) {
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
			"GrahaName": GrahaName,
		}).ParseFiles(
		"templates/analysis.html",
		"templates/header.html",
		"templates/navbar.html",
		"templates/charts.html",
		"templates/footer.html"))
	return tmpl.ExecuteTemplate(w, tmplName, page)
}
