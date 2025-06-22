package views

import (
	"fmt"
	"html/template"
	"jyotish/analysis"
	"jyotish/constants"
	"jyotish/models"
	"net/http"
	"strings"
)

type AnalysisPage struct {
	MainPage
	Chart     analysis.Chart
	BhavaDesc []models.BhavaDescription
}

func GetGrahaNameForChart(graha analysis.GrahaLocCombust, lang string) template.HTML {
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

func GetBhava(chart *analysis.Chart, bhavaNum int) analysis.Bhava {
	return chart.Bhavas[bhavaNum-1]
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

func ListOfAspectingGrahas(aspectingGrahas []analysis.AscpectAndDegree, lang string) string {
	sb := strings.Builder{}
	first := true
	for _, ag := range aspectingGrahas {
		if !first {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%s (%d)", models.GrahaName(ag.Name, lang), ag.Degree))
		first = false
	}
	return sb.String()
}

func GetGrahaPositionResult(grahaPosition analysis.GrahaPosition, lang string) template.HTML {
	var vocab *models.Language

	if lang == "en" {
		vocab = &models.EnglishVocab
	} else {
		vocab = &models.HindiVocab
	}

	sb := strings.Builder{}

	switch grahaPosition.Result {
	case constants.RESULT_GAINS:
		sb.WriteString(fmt.Sprintf(`<span class="good">%d `, grahaPosition.Count))
		sb.WriteString(vocab.Benefic)
	case constants.RESULT_LOSSES:
		sb.WriteString(fmt.Sprintf(`<span class="bad">%d `, grahaPosition.Count))
		sb.WriteString(vocab.Malefic)
	default:
		sb.WriteString(fmt.Sprintf(`<span class="neutral">%d `, grahaPosition.Count))
		sb.WriteString(vocab.Neutral)
	}
	sb.WriteString(" - ")
	switch grahaPosition.Subjects {
	case constants.SUBJECTS_LIVING_BEING:
		sb.WriteString(vocab.SubjectsLiving)
	case constants.SUBJECTS_NON_LIVING_BEING:
		sb.WriteString(vocab.SubjectsNonLiving)
	default:
		sb.WriteString(fmt.Sprintf("%s, %s", vocab.SubjectsLiving, vocab.SubjectsNonLiving))
	}
	sb.WriteString(fmt.Sprintf(" (%d)</span>", grahaPosition.Score))

	return template.HTML(sb.String())
}

func GetRashiName(number int, lang string) string {
	return fmt.Sprintf("%s (%d)", models.RashiName(number, lang), number)
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
	page.BhavaDesc, _ = models.GetBhavaDescription(user.Lang)

	return &page, nil

}

func (page *AnalysisPage) Send(w http.ResponseWriter) error {
	tmplName := "analysis"
	tmpl := template.Must(template.New(tmplName).Funcs(
		template.FuncMap{
			"GetBhava":               GetBhava,
			"GetGrahaNameForChart":   GetGrahaNameForChart,
			"GetGrahaName":           GetGrahaName,
			"ListOfGrahas":           ListOfGrahas,
			"ListOfAspectingGrahas":  ListOfAspectingGrahas,
			"GrahaNature":            models.GrahaNature,
			"GrahaMotion":            models.GrahaMotion,
			"YesOrNo":                models.YesOrNo,
			"GrahaPosition":          models.GrahaPosition,
			"GrahaState":             models.GrahaState,
			"GetRashiName":           GetRashiName,
			"GetGrahaPositionResult": GetGrahaPositionResult,
		}).ParseFiles(
		"templates/analysis.html",
		"templates/lagna-chart.html",
		"templates/grahas-relations.html",
		"templates/grahas-aspects.html",
		"templates/grahas-nature.html",
		"templates/grahas-strength.html",
		"templates/bhavas.html",
		"templates/header.html",
		"templates/navbar.html",
		"templates/footer.html"))
	return tmpl.ExecuteTemplate(w, tmplName, page)
}
